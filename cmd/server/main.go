package main

import (
	"github.com/gin-gonic/gin"
	di "github.com/nodejayes/generic-di"
	"github.com/nodejayes/streaming-ui-server/cmd/server/example"
	"github.com/nodejayes/streaming-ui-server/pkg/server"
	"net/http"
)

func main() {
	// use this first!
	server.CreateActionContext(func(clientID string, ctx *gin.Context) (example.ActionContext, error) {
		return example.ActionContext{
			ClientID: clientID,
			State:    di.Inject[example.AppState](clientID),
		}, nil
	})

	server.CreateCleanup(func(clientID string) {
		di.Destroy[example.AppState](clientID)
	})

	server.RegisterAction[*example.CounterAction, example.ActionContext](example.NewCounterAction())
	server.RegisterAction[*example.ReloadAction, example.ActionContext](example.NewReloadAction())
	server.RegisterAction[*example.PingAction, example.ActionContext](example.NewPingAction())

	server.AddPage("/", example.NewIndexPage)

	server.Engine().GET("test", func(ctx *gin.Context) {
		ctx.Data(http.StatusOK, "application/json", []byte("Test"))
	})

	server.Run(":40000")
}
