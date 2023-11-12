package socket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	event_emitter "github.com/nodejayes/event-emitter"
	di "github.com/nodejayes/generic-di"
	"github.com/nodejayes/streaming-ui-server/pkg/server/utils"
)

var upgrader = websocket.Upgrader{
	HandshakeTimeout: 30000,
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	EnableCompression: true,
}

func Handle(contextCreator func(clientID, pageID string, ctx *gin.Context) (any, error), stateCleaner func(clientID string)) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			log.Printf("abort connection invalid upgrade %s", err.Error())
			return
		}
		defer func() {
			_ = conn.Close()
		}()
		clientID := ctx.Query("clientId")
		if len(clientID) < 1 {
			log.Printf("abort connection invalid clientId %s", clientID)
			return
		}
		pageID := ctx.Query("pageId")
		actionContext, err := contextCreator(clientID, pageID, ctx)
		if err != nil {
			log.Printf("abort connection invalid contextCreator %s", err.Error())
			return
		}
		socket := NewSocket(conn, clientID)
		sessionFactory := di.Inject[factory]()
		sessionFactory.AddSession(clientID, socket)
		for {
			_, p, err := conn.ReadMessage()
			if err != nil {
				println(fmt.Sprintf("error in socket handler: %s", err.Error()))
				utils.EmitCleanupHandler(pageID)
				stateCleaner(clientID)
				sessionFactory.RemoveSession(clientID)
				return
			}
			event_emitter.Emit(ParseSocketMessageEvent, ParseSocketMessageArguments{
				Message:  p,
				Context:  actionContext,
				ClientID: clientID,
				PageID:   pageID,
			})
		}
	}
}
