package server

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	event_emitter "github.com/nodejayes/event-emitter"
	di "github.com/nodejayes/generic-di"
	livereplacer "github.com/nodejayes/streaming-ui-server/live-replacer"
	"github.com/nodejayes/streaming-ui-server/server/identity"
	"github.com/nodejayes/streaming-ui-server/server/socket"
	"github.com/nodejayes/streaming-ui-server/server/ui/types"
)

func init() {
	di.Injectable(New)
}

var contextCreator func(clientId string, ctx *gin.Context) (any, error)
var actionHandlerMutex = &sync.Mutex{}
var actionHandlers = make(map[string][]func(action types.Action, ctx any))

type (
	ClientIdentiy interface {
		GetClientId() string
	}
)

func New() *gin.Engine {
	router := gin.Default()
	router.StaticFS("/live-replacer", http.FS(livereplacer.Files))
	router.GET("/identity", identity.Handle)
	router.GET("/ws", socket.Handle(contextCreator))
	return router
}

func getActionKey(action types.Action) string {
	if action.GetElementID() == "" {
		action.SetElementID(uuid.New().String())
	}
	return fmt.Sprintf("%s_%s", action.GetType(), action.GetElementID())
}

func CreateActionContext[T any](creator func(clientId string, ctx *gin.Context) (T, error)) {
	contextCreator = func(clientId string, ctx *gin.Context) (any, error) {
		return creator(clientId, ctx)
	}
}

func SendCaller[TPayload any, TContext ClientIdentiy](action socket.Action[TPayload, TContext]) {
	for _, session := range socket.Factory().GetSessions(func(socket socket.Instance) bool { return socket.GetClientId() == action.Context.GetClientId() }) {
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

func OnAction[TAction types.Action, TContext any](actionInstance TAction, execution func(action TAction, ctx TContext)) {
	event_emitter.Subscribe(socket.ParseSocketMessageEvent, func(params socket.ParseSocketMessageArguments, _ socket.ParseSocketMessageArguments) {
		var action TAction
		err := json.Unmarshal(params.Message, &action)
		if err != nil {
			println(err.Error())
			return
		}
		if action.GetType() != actionInstance.GetType() {
			return
		}
		c, ok := params.Context.(TContext)
		if !ok {
			return
		}
		execution(action, c)
	})
}

func AddPage(page types.Page) {
	di.Inject[gin.Engine]().GET(page.GetPath(), func(ctx *gin.Context) {
		ctx.Data(http.StatusOK, "text/html", []byte(page.Render()))
	})
}

func RegisterAction[TAction types.Action, TContext any](action TAction) {
	OnAction[TAction, TContext](action, func(action TAction, ctx TContext) {
		actionKey := getActionKey(action)
		handlers, ok := actionHandlers[actionKey]
		if !ok {
			return
		}
		for _, handler := range handlers {
			handler(action, ctx)
		}
	})
}

func CreateEventHandler[TContext any](action types.Action, handler func(action types.Action, ctx TContext)) types.Action {
	actionHandlerMutex.Lock()
	defer actionHandlerMutex.Unlock()
	actionKey := getActionKey(action)
	handlers, ok := actionHandlers[actionKey]
	if !ok {
		handlers = make([]func(action types.Action, ctx any), 0)
	}
	handlers = append(handlers, func(action types.Action, ctx any) {
		ctxConverted, ok := ctx.(TContext)
		if !ok {
			fmt.Println(fmt.Sprintf("Error on convert Context: %v", ok))
			return
		}
		handler(action, ctxConverted)
	})
	actionHandlers[actionKey] = handlers
	return action
}

func Run(addr ...string) {
	log.Fatal(di.Inject[gin.Engine]().Run(addr...))
}
