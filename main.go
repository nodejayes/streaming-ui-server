package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nodejayes/streaming-ui-server/server"
	"github.com/nodejayes/streaming-ui-server/server/socket"
	"github.com/nodejayes/streaming-ui-server/timetracker"
)

type ActionContext struct {
	ID string
}

func (ctx ActionContext) GetClientId() string {
	return ctx.ID
}

func main() {
	server.CreateActionContext(func(clientId string, ctx *gin.Context) (ActionContext, error) {
		return ActionContext{
			ID: clientId,
		}, nil
	})

	server.OnAction[string, ActionContext]("ping", func(params string, ctx ActionContext) {
		server.SendCaller(socket.Action[string, ActionContext]{
			Type:    "replaceHtml::header",
			Payload: "<h1>Pong</h1>",
			Context: ctx,
		})
	})

	timetracker.NewHandler()
	server.Run(":40000")
}
