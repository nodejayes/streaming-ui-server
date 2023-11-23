package socket

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

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
var cleanupMutex = &sync.Mutex{}
var cleanupSessions = make(map[string]bool)

func markForCleanup(clientID string) {
	cleanupMutex.Lock()
	defer cleanupMutex.Unlock()
	cleanupSessions[clientID] = false
}

func unmarkCleanup(clientID string) {
	cleanupMutex.Lock()
	defer cleanupMutex.Unlock()
	delete(cleanupSessions, clientID)
}

func existInCleanup(clientID string) bool {
	cleanupMutex.Lock()
	defer cleanupMutex.Unlock()
	_, ok := cleanupSessions[clientID]
	return ok
}

func Handle(contextCreator func(clientID string, ctx *gin.Context) (any, error), stateCleaner func(clientID string), stateCleanupTime time.Duration) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			log.Printf("abort connection invalid upgrade %s", err.Error())
			return
		}
		defer func() {
			_ = conn.Close()
		}()
		clientID, err := ctx.Cookie("ClientId")
		if len(clientID) < 1 || err != nil {
			log.Printf("abort connection invalid clientId %s", clientID)
			return
		}
		if existInCleanup(clientID) {
			unmarkCleanup(clientID)
		}
		actionContext, err := contextCreator(clientID, ctx)
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
				utils.EmitCleanupHandler(clientID)
				markForCleanup(clientID)
				time.AfterFunc(stateCleanupTime, func() {
					if stateCleaner != nil && existInCleanup(clientID) {
						stateCleaner(clientID)
					}
				})
				sessionFactory.RemoveSession(socket.GetSocketID())
				return
			}
			event_emitter.Emit(ParseSocketMessageEvent, ParseSocketMessageArguments{
				Message:  p,
				Context:  actionContext,
				ClientID: clientID,
			})
		}
	}
}
