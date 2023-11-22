package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	event_emitter "github.com/nodejayes/event-emitter"
	di "github.com/nodejayes/generic-di"
	livereplacer "github.com/nodejayes/streaming-ui-server/pkg/live-replacer"
	"github.com/nodejayes/streaming-ui-server/pkg/server/identity"
	"github.com/nodejayes/streaming-ui-server/pkg/server/socket"
	"github.com/nodejayes/streaming-ui-server/pkg/server/ui/types"
	"github.com/nodejayes/streaming-ui-server/pkg/server/utils"
)

func init() {
	di.Injectable(New)
}

var contextCreator func(clientID string, ctx *gin.Context) (any, error)
var stateCleaner func(clientID string)
var actionHandlerMutex = &sync.Mutex{}
var actionHandlers = make(map[string][]func(action types.Action, ctx any, elementID string, inputs map[string]map[string]string, eventData map[string]any))

type (
	ClientIdentiy interface {
		GetClientId() string
	}
	ClientMessage[T any] struct {
		ElementID string                       `json:"elementId"`
		Action    T                            `json:"action"`
		Inputs    map[string]map[string]string `json:"inputs"`
		EventData map[string]any               `json:"eventData"`
	}
)

func New() *gin.Engine {
	router := gin.Default()
	router.StaticFS("/live-replacer", http.FS(livereplacer.Files))
	router.GET("/identity", identity.Handle)
	router.GET("/ws", socket.Handle(contextCreator, stateCleaner))
	return router
}

func getActionKey(action types.Action, elementID *string) string {
	if (*elementID) == "" {
		*elementID = uuid.New().String()
	}
	return fmt.Sprintf("%s_%s", action.GetType(), *elementID)
}

func Engine() *gin.Engine {
	return di.Inject[gin.Engine]()
}

func CreateActionContext[T any](creator func(clientID string, ctx *gin.Context) (T, error)) {
	contextCreator = func(clientID string, ctx *gin.Context) (any, error) {
		return creator(clientID, ctx)
	}
}

func CreateCleanup(cleaner func(clientID string)) {
	stateCleaner = cleaner
}

func SendCaller[TPayload any, TContext ClientIdentiy](action socket.Action[TPayload, TContext]) {
	clientID := action.Context.GetClientId()
	for _, session := range socket.Factory().GetSessions(func(socket socket.Instance) bool { return socket.GetClientId() == clientID }) {
		session.Send(socket.Action[any, any]{
			Type:    action.Type,
			Payload: action.Payload,
		})
	}
}

func SendAll[TPayload, TContext any](action socket.Action[TPayload, TContext]) {
	for _, session := range socket.Factory().GetSessions(func(socket socket.Instance) bool { return true }) {
		session.Send(socket.Action[any, any]{
			Type:    action.Type,
			Payload: action.Payload,
		})
	}
}

func SendTo[TPayload, TContext any](socketSelector func(socket socket.Instance) bool, action socket.Action[TPayload, TContext]) {
	for _, session := range socket.Factory().GetSessions(socketSelector) {
		session.Send(socket.Action[any, any]{
			Type:    action.Type,
			Payload: action.Payload,
		})
	}
}

func OnAction[TAction types.Action, TContext any](actionInstance TAction, execution func(action TAction, ctx TContext, clientID, elementID string, inputs map[string]map[string]string, eventData map[string]any)) {
	event_emitter.Subscribe(socket.ParseSocketMessageEvent, func(params socket.ParseSocketMessageArguments, _ socket.ParseSocketMessageArguments) {
		var msg ClientMessage[TAction]
		err := json.Unmarshal(params.Message, &msg)
		if err != nil {
			return
		}
		if msg.Action.GetType() != actionInstance.GetType() {
			return
		}
		c, ok := params.Context.(TContext)
		if !ok {
			return
		}
		execution(msg.Action, c, params.ClientID, msg.ElementID, msg.Inputs, msg.EventData)
	})
}

func AddPage[T types.Page](path string, pageCreator func(clientID string) T) {
	di.Inject[gin.Engine]().GET(path, func(ctx *gin.Context) {
		clientId, err := ctx.Cookie("ClientId")
		if err != nil {
			clientId = uuid.NewString()
			ctx.SetCookie("ClientId", clientId, 3600, "/", "localhost", false, true)
		}
		page := pageCreator(clientId)
		pageStr := page.Render()
		ctx.Data(http.StatusOK, "text/html", []byte(pageStr))
	})
}

func RegisterAction[TAction types.Action, TContext any](action TAction) {
	OnAction[TAction, TContext](action, func(action TAction, ctx TContext, clientID, elementID string, inputs map[string]map[string]string, eventData map[string]any) {
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

func CreateEventHandler[TContext any, TEventData types.EventData](action types.Action, handler func(action types.Action, ctx TContext, elementID string, inputs map[string]map[string]string, eventData TEventData)) func(eID string) types.Action {
	return func(eID string) types.Action {
		actionHandlerMutex.Lock()
		defer actionHandlerMutex.Unlock()
		actionKey := getActionKey(action, &eID)
		handlers, ok := actionHandlers[actionKey]
		if !ok {
			handlers = make([]func(action types.Action, ctx any, elementID string, inputs map[string]map[string]string, eventData map[string]any), 0)
		}
		handlers = append(handlers, func(action types.Action, ctx any, elementID string, inputs map[string]map[string]string, eventData map[string]any) {
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
