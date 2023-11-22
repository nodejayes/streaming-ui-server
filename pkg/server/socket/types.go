package socket

import (
	event_emitter "github.com/nodejayes/event-emitter"
)

type (
	Action[TPayload, TContext any] struct {
		Type      string                       `json:"type"`
		Payload   TPayload                     `json:"payload"`
		Inputs    map[string]map[string]string `json:"inputs"`
		EventData map[string]any               `json:"eventData"`
		Context   TContext                     `json:"-"`
	}

	SessionFactory interface {
		AddSession(id string, socket Instance)
		GetSession(id string) (Instance, error)
		GetSessions(selector func(socket Instance) bool) []Instance
		RemoveSession(id string)
	}

	Instance interface {
		GetClientId() string
		GetSocketID() string
		Send(action Action[any, any]) error
		Reciever() event_emitter.Event[ParseSocketMessageArguments, ParseSocketMessageArguments]
	}
)
