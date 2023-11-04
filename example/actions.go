package example

import (
	"github.com/google/uuid"
)

type (
	PingAction struct {
		ElementID string                       `json:"elementId"`
		Type      string                       `json:"type"`
		Payload   string                       `json:"payload"`
		Inputs    map[string]map[string]string `json:"inputs"`
	}
	CounterAction struct {
		ElementID string                       `json:"elementId"`
		Type      string                       `json:"type"`
		Payload   int                          `json:"payload"`
		Inputs    map[string]map[string]string `json:"inputs"`
	}
)

func NewPingAction(payload string) *PingAction {
	return &PingAction{
		ElementID: uuid.New().String(),
		Type:      "ping",
		Payload:   payload,
		Inputs:    nil,
	}
}

func (ctx *PingAction) GetElementID() string {
	return ctx.ElementID
}

func (ctx *PingAction) SetElementID(value string) {
	ctx.ElementID = value
}

func (ctx *PingAction) GetType() string {
	return ctx.Type
}

func (ctx *PingAction) GetPayload() any {
	return ctx.Payload
}

func (ctx *PingAction) GetInputs() map[string]map[string]string {
	return ctx.Inputs
}

func NewCounterAction(payload int) *CounterAction {
	return &CounterAction{
		ElementID: uuid.New().String(),
		Type:      "count increase",
		Payload:   payload,
		Inputs:    nil,
	}
}

func (ctx *CounterAction) GetElementID() string {
	return ctx.ElementID
}

func (ctx *CounterAction) SetElementID(value string) {
	ctx.ElementID = value
}

func (ctx *CounterAction) GetType() string {
	return ctx.Type
}

func (ctx *CounterAction) GetPayload() any {
	return ctx.Payload
}

func (ctx *CounterAction) GetInputs() map[string]map[string]string {
	return ctx.Inputs
}
