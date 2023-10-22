package socket

import (
	"encoding/json"
	"log"

	event_emitter "github.com/nodejayes/event-emitter"
)

type ParseSocketMessageArguments struct {
	Message []byte
	Context any
}

var parseSocketMessageEvent = event_emitter.Event[ParseSocketMessageArguments, Action[any, any]]{
	Token: "recieveMessageEvent",
	Handler: func(event string, params ParseSocketMessageArguments) (Action[any, any], error) {
		var action Action[any, any]
		err := json.Unmarshal(params.Message, &action)
		action.Context = params.Context
		return action, err
	},
	OnError: func(event string, err error) {
		log.Printf("[Error %s]: %s", event, err.Error())
	},
}

func OnAction[TPayload, TContext any](typ string, execution func(action Action[TPayload, TContext])) event_emitter.Subscription {
	return event_emitter.Subscribe(parseSocketMessageEvent, func(params ParseSocketMessageArguments, action Action[any, any]) {
		if action.Type != typ {
			return
		}
		// TODO: better concept for parsing int
		convertedPayload, ok := action.Payload.(TPayload)
		if !ok {
			log.Printf("skip action %s payload no type match %v", action.Type, action.Payload)
			return
		}
		convertedContext, ok := action.Context.(TContext)
		if !ok {
			log.Printf("skip action %s context no type match %v", action.Type, action.Context)
			return
		}
		execution(Action[TPayload, TContext]{
			Type:    action.Type,
			Payload: convertedPayload,
			Context: convertedContext,
		})
	})
}
