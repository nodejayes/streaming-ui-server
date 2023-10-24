package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	event_emitter "github.com/nodejayes/event-emitter"
	di "github.com/nodejayes/generic-di"
	livereplacer "github.com/nodejayes/streaming-ui-server/live-replacer"
	"github.com/nodejayes/streaming-ui-server/server/identity"
	"github.com/nodejayes/streaming-ui-server/server/socket"
)

func init() {
	di.Injectable(New)
}

var contextCreator func(clientId string, ctx *gin.Context) (any, error)

type (
	ClientIdentiy interface {
		GetClientId() string
	}
	Action interface {
		GetType() string
	}
)

func New() *gin.Engine {
	router := gin.Default()
	router.StaticFS("/live-replacer", http.FS(livereplacer.Files))
	router.GET("/identity", identity.Handle)
	router.GET("/ws", socket.Handle(contextCreator))
	return router
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

func OnAction[TAction Action, TContext any](actionInstance TAction, execution func(action TAction, ctx TContext)) {
	event_emitter.Subscribe(socket.ParseSocketMessageEvent, func(params socket.ParseSocketMessageArguments, _ socket.ParseSocketMessageArguments) {
		var action TAction
		err := json.Unmarshal(params.Message, &action)
		if err != nil {
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

func Run(addr ...string) {
	log.Fatal(di.Inject[gin.Engine]().Run(addr...))
}
