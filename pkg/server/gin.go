package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	event_emitter "github.com/nodejayes/event-emitter"
	di "github.com/nodejayes/generic-di"
	livereplacer "github.com/nodejayes/streaming-ui-server/pkg/live-replacer"
	"github.com/nodejayes/streaming-ui-server/pkg/server/socket"
	"github.com/nodejayes/streaming-ui-server/pkg/server/ui/types"
	"github.com/nodejayes/streaming-ui-server/pkg/server/utils"
)

func init() {
	di.Injectable(New)
}

type (
	Action struct {
		Type string `json:"type"`
	}
	ServerOptions struct {
		StateCleanupTime time.Duration
		SessionLifetime  int
		SecureCookie     bool
	}
)

var options = ServerOptions{
	StateCleanupTime: 60 * time.Second,
	SessionLifetime:  3600,
	SecureCookie:     true,
}
var contextCreator func(clientID string, ctx *gin.Context) (any, error)
var stateCleaner func(clientID string)
var actionHandlerMutex = &sync.Mutex{}
var actionHandlers = make(map[string][]func(action string, ctx any, elementID string, inputs map[string]map[string]string, eventData map[string]any))

type (
	ClientIdentity interface {
		GetClientId() string
	}
	ClientMessage[T any] struct {
		ElementID string                       `json:"elementId"`
		Action    T                            `json:"action"`
		Inputs    map[string]map[string]string `json:"inputs"`
		EventData map[string]any               `json:"eventData"`
	}
)

func Configure(serverOptions ServerOptions) {
	options = serverOptions
}

func New() *gin.Engine {
	router := gin.Default()
	router.StaticFS("/live-replacer", http.FS(livereplacer.Files))
	return router
}

func getActionKey(action string, elementID *string) string {
	if (*elementID) == "" {
		*elementID = uuid.New().String()
	}
	return fmt.Sprintf("%s_%s", action, *elementID)
}

func Engine() *gin.Engine {
	return di.Inject[gin.Engine]()
}

func BuildActionContext[T any](creator func(clientID string, ctx *gin.Context) (T, error)) {
	contextCreator = func(clientID string, ctx *gin.Context) (any, error) {
		return creator(clientID, ctx)
	}
}

func CleanupSession(cleaner func(clientID string)) {
	stateCleaner = cleaner
}

func SendCaller[TPayload any, TContext ClientIdentity](action socket.Action[TPayload, TContext]) {
	clientID := action.Context.GetClientId()
	for _, session := range socket.Factory().GetSessions(func(socket socket.Instance) bool { return socket.GetClientId() == clientID }) {
		_ = session.Send(socket.Action[any, any]{
			Type:    action.Type,
			Payload: action.Payload,
		})
	}
}

func SendAll[TPayload, TContext any](action socket.Action[TPayload, TContext]) {
	for _, session := range socket.Factory().GetSessions(func(socket socket.Instance) bool { return true }) {
		_ = session.Send(socket.Action[any, any]{
			Type:    action.Type,
			Payload: action.Payload,
		})
	}
}

func SendTo[TPayload, TContext any](socketSelector func(socket socket.Instance) bool, action socket.Action[TPayload, TContext]) {
	for _, session := range socket.Factory().GetSessions(socketSelector) {
		_ = session.Send(socket.Action[any, any]{
			Type:    action.Type,
			Payload: action.Payload,
		})
	}
}

func OnAction[TContext any](actionInstance string, execution func(action string, ctx TContext, clientID, elementID string, inputs map[string]map[string]string, eventData map[string]any)) {
	event_emitter.Subscribe(socket.ParseSocketMessageEvent, func(params socket.ParseSocketMessageArguments, _ socket.ParseSocketMessageArguments) {
		var msg ClientMessage[Action]
		err := json.Unmarshal(params.Message, &msg)
		if err != nil {
			return
		}
		if msg.Action.Type != actionInstance {
			return
		}
		c, ok := params.Context.(TContext)
		if !ok {
			return
		}
		execution(msg.Action.Type, c, params.ClientID, msg.ElementID, msg.Inputs, msg.EventData)
	})
}

func AddPage[T types.Page, TContext any](path string, pageCreator func(ctx *gin.Context) T, pageActions []string) {
	router := di.Inject[gin.Engine]()
	router.GET(fmt.Sprintf("%s/ws", path), socket.Handle(contextCreator, stateCleaner, options.StateCleanupTime))
	for _, pageAction := range pageActions {
		RegisterAction[TContext](pageAction)
	}
	router.GET(path, func(ctx *gin.Context) {
		clientId, err := ctx.Cookie("ClientId")
		if err != nil {
			clientId = uuid.NewString()
			ctx.SetCookie("ClientId", clientId, options.SessionLifetime, ctx.Request.URL.Path, ctx.Request.URL.Host, options.SecureCookie, !options.SecureCookie)
		}
		page := pageCreator(ctx)
		pageStr := page.Render()
		ctx.Data(http.StatusOK, "text/html", []byte(pageStr))
	})
}

func RegisterAction[TContext any](action string) {
	OnAction[TContext](action, func(action string, ctx TContext, clientID, elementID string, inputs map[string]map[string]string, eventData map[string]any) {
		actionKey := getActionKey(action, &elementID)
		handlers, ok := actionHandlers[actionKey]
		if !ok {
			return
		}
		for _, handler := range handlers {
			handler(action, ctx, elementID, inputs, eventData)
		}
	})
}

func cleanupHandler(clientID string) {
	actionHandlerMutex.Lock()
	defer actionHandlerMutex.Unlock()
	for key := range actionHandlers {
		if !strings.HasPrefix(key, clientID) {
			continue
		}
		delete(actionHandlers, key)
	}
}

func CreateEventHandler[TContext any, TEventData types.EventData](action string, handler func(action string, ctx TContext, elementID string, inputs map[string]map[string]string, eventData TEventData)) func(eID string) string {
	return func(eID string) string {
		actionHandlerMutex.Lock()
		defer actionHandlerMutex.Unlock()
		actionKey := getActionKey(action, &eID)
		handlers, ok := actionHandlers[actionKey]
		if !ok {
			handlers = make([]func(action string, ctx any, elementID string, inputs map[string]map[string]string, eventData map[string]any), 0)
		}
		handlers = append(handlers, func(action string, ctx any, elementID string, inputs map[string]map[string]string, eventData map[string]any) {
			ctxConverted, ok := ctx.(TContext)
			if !ok {
				fmt.Printf("Error on convert Context: %v", ok)
				return
			}
			var convertedEventData TEventData
			str, err := json.Marshal(eventData)
			if err != nil {
				fmt.Printf("Error on convert EventData: %v", err.Error())
				return
			}
			err = json.Unmarshal(str, &convertedEventData)
			if err != nil {
				fmt.Printf("Error on convert EventData: %v", err.Error())
				return
			}
			handler(action, ctxConverted, elementID, inputs, convertedEventData)
		})
		actionHandlers[actionKey] = handlers
		return action
	}
}

func Run(addr ...string) {
	event_emitter.Subscribe(utils.CleanupEventHandler, func(args string, _ string) {
		cleanupHandler(args)
	})
	log.Fatal(di.Inject[gin.Engine]().Run(addr...))
}
