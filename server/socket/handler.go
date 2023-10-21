package socket

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	event_emitter "github.com/nodejayes/event-emitter"
	di "github.com/nodejayes/generic-di"
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

func Handle(contextCreator func(clientId string, ctx *gin.Context) (any, error)) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			log.Printf("abort connection invalid upgrade %s", err.Error())
			return
		}
		defer func() {
			_ = conn.Close()
		}()
		clientId := ctx.Query("clientId")
		if len(clientId) < 1 {
			log.Printf("abort connection invalid clientId %s", clientId)
			return
		}
		actionContext, err := contextCreator(clientId, ctx)
		if err != nil {
			log.Printf("abort connection invalid contextCreator %s", err.Error())
			return
		}
		socket := NewSocket(conn, clientId)
		sessionFactory := di.Inject[factory]()
		sessionFactory.AddSession(clientId, socket)
		for {
			_, p, err := conn.ReadMessage()
			if err != nil {
				sessionFactory.RemoveSession(clientId)
				return
			}
			event_emitter.Emit(parseSocketMessageEvent, ParseSocketMessageArguments{
				Message: p,
				Context: actionContext,
			})
			time.Sleep(time.Second)
		}
	}
}
