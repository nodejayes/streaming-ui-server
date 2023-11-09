package main

import (
	"github.com/gin-gonic/gin"
	di "github.com/nodejayes/generic-di"
	"github.com/nodejayes/streaming-ui-server/cmd/server/example"
	"github.com/nodejayes/streaming-ui-server/pkg/server"
)

func main() {
	server.CreateActionContext(func(clientId string, ctx *gin.Context) (example.ActionContext, error) {
		return example.ActionContext{
			ID:    clientId,
			State: di.Inject[example.AppState](clientId),
		}, nil
	})

	server.RegisterAction[*example.CounterAction, example.ActionContext](example.NewCounterAction())
	server.RegisterAction[*example.PingAction, example.ActionContext](example.NewPingAction())

	server.AddPage(example.NewIndexPage())

	server.Run(":40000")
}