package utils

import (
	"encoding/json"

	event_emitter "github.com/nodejayes/event-emitter"
	"github.com/nodejayes/streaming-ui-server/server/socket"
	"github.com/nodejayes/streaming-ui-server/server/ui/types"
)

func CreateHandler[TAction types.Action, TContext any](actionInstance TAction, handler func(action TAction, ctx TContext)) {
	event_emitter.Subscribe(socket.ParseSocketMessageEvent, func(params socket.ParseSocketMessageArguments, _ socket.ParseSocketMessageArguments) {
		var action TAction
		err := json.Unmarshal(params.Message, &action)
		if err != nil {
			return
		}
		if action.GetType() != actionInstance.GetType() {
			return
		}
		c, ok := params.Context.(TContext)
		if !ok {
			return
		}
		handler(action, c)
	})
}
