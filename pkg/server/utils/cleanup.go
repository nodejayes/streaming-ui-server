package utils

import event_emitter "github.com/nodejayes/event-emitter"

var CleanupEventHandler = event_emitter.Event[string, string]{
	Token: "cleanupHandler",
	Handler: func(event string, params string) (string, error) {
		return params, nil
	},
	OnError: func(event string, err error) {},
}

func EmitCleanupHandler(pageID string) {
	event_emitter.Emit(CleanupEventHandler, pageID)
}
