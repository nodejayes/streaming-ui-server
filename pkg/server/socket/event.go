package socket

import (
	"log"

	event_emitter "github.com/nodejayes/event-emitter"
)

type ParseSocketMessageArguments struct {
	Message  []byte
	Context  any
	ClientID string
	PageID   string
}

var ParseSocketMessageEvent = event_emitter.Event[ParseSocketMessageArguments, ParseSocketMessageArguments]{
	Token: "recieveMessageEvent",
	Handler: func(event string, params ParseSocketMessageArguments) (ParseSocketMessageArguments, error) {
		return params, nil
	},
	OnError: func(event string, err error) {
		log.Printf("[Error %s]: %s", event, err.Error())
	},
}
