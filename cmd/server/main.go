package main

import (
	"github.com/gin-gonic/gin"
	di "github.com/nodejayes/generic-di"
	"github.com/nodejayes/streaming-ui-server/cmd/server/example"
	"github.com/nodejayes/streaming-ui-server/pkg/server"
	"net/http"
	"time"
)

func main() {
	server.Configure(server.ServerOptions{
		StateCleanupTime: 60 * time.Second,
		SessionLifetime:  3600,
		SecureCookie:     true,
	})

	server.BuildActionContext(func(clientID string, ctx *gin.Context) (example.ActionContext, error) {
		return example.ActionContext{
			ClientID: clientID,
			State:    di.Inject[example.AppState](clientID),
		}, nil
	})

	server.CleanupSession(func(clientID string) {
		di.Destroy[example.AppState](clientID)
	})

	server.AddPage[*example.IndexPage, example.ActionContext]("/", example.NewIndexPage, []string{
		example.PingAction,
		example.CounterAction,
		example.ReloadAction,
		example.DoubleClickNoticeAction,
		example.ContextMenuNoticeAction,
	})

	server.Engine().GET("test", func(ctx *gin.Context) {
		ctx.Data(http.StatusOK, "application/json", []byte("Test"))
	})

	server.Run(":40000")
}
