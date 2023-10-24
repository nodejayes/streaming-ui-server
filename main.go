package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	di "github.com/nodejayes/generic-di"
	"github.com/nodejayes/streaming-ui-server/example"
	"github.com/nodejayes/streaming-ui-server/server"
	"github.com/nodejayes/streaming-ui-server/server/socket"
	"github.com/nodejayes/streaming-ui-server/timetracker"
)

func main() {
	server.CreateActionContext(func(clientId string, ctx *gin.Context) (example.ActionContext, error) {
		return example.ActionContext{
			ID:    clientId,
			State: di.Inject[example.AppState](clientId),
		}, nil
	})

	server.OnAction(example.NewPingAction(), func(action example.PingAction, ctx example.ActionContext) {
		server.SendCaller(socket.Action[string, example.ActionContext]{
			Type:    "replaceHtml::#header",
			Payload: "<h1>Pong</h1>",
			Context: ctx,
		})
	})

	server.OnAction(example.NewCounterAction(), func(action example.CounterAction, ctx example.ActionContext) {
		ctx.State.Counter += action.Payload
		server.SendCaller(socket.Action[string, example.ActionContext]{
			Type:    "replaceHtml::#counter li",
			Payload: fmt.Sprintf("<p>%v</p>", ctx.State.Counter),
			Context: ctx,
		})
	})

	timetracker.NewHandler()
	server.Run(":40000")
}
