package example

type (
	PingAction struct {
		Type    string `json:"type"`
		Payload string `json:"payload"`
	}
	CounterAction struct {
		Type    string `json:"type"`
		Payload int    `json:"payload"`
	}
)

func NewPingAction(payload string) *PingAction {
	return &PingAction{
		Type:    "ping",
		Payload: payload,
	}
}

func (ctx *PingAction) GetType() string {
	return ctx.Type
}

func (ctx *PingAction) GetPayload() any {
	return ctx.Payload
}

func NewCounterAction(payload int) *CounterAction {
	return &CounterAction{
		Type:    "count increase",
		Payload: payload,
	}
}

func (ctx *CounterAction) GetType() string {
	return ctx.Type
}

func (ctx *CounterAction) GetPayload() any {
	return ctx.Payload
}
