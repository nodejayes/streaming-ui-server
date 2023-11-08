package socket

import (
	"encoding/json"

	"github.com/gorilla/websocket"
	event_emitter "github.com/nodejayes/event-emitter"
)

type instance struct {
	id         string
	connection *websocket.Conn
	reciever   event_emitter.Event[ParseSocketMessageArguments, ParseSocketMessageArguments]
}

func NewSocket(connection *websocket.Conn, id string) *instance {
	return &instance{
		connection: connection,
		id:         id,
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
	return ctx.id
}
