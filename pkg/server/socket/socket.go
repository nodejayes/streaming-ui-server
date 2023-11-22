package socket

import (
	"encoding/json"
	"github.com/google/uuid"

	"github.com/gorilla/websocket"
	event_emitter "github.com/nodejayes/event-emitter"
)

type instance struct {
	id         string
	clientID   string
	connection *websocket.Conn
	reciever   event_emitter.Event[ParseSocketMessageArguments, ParseSocketMessageArguments]
}

func NewSocket(connection *websocket.Conn, clientID string) *instance {
	return &instance{
		connection: connection,
		id:         uuid.NewString(),
		clientID:   clientID,
		reciever:   ParseSocketMessageEvent,
	}
}

func (ctx *instance) Send(action Action[any, any]) error {
	msg, err := json.Marshal(action)
	if err != nil {
		return err
	}
	return ctx.connection.WriteMessage(1, msg)
}

func (ctx *instance) Reciever() event_emitter.Event[ParseSocketMessageArguments, ParseSocketMessageArguments] {
	return ctx.reciever
}

func (ctx *instance) GetClientId() string {
	return ctx.clientID
}

func (ctx *instance) GetSocketID() string {
	return ctx.id
}
