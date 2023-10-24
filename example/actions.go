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

func NewPingAction() PingAction {
	return PingAction{
		Type:    "ping",
		Payload: "",
	}
}

func (ctx PingAction) GetType() string {
	return ctx.Type
}

func NewCounterAction() CounterAction {
	return CounterAction{
		Type:    "count increase",
		Payload: 0,
	}
}

func (ctx CounterAction) GetType() string {
	return ctx.Type
}
