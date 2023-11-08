package example

type (
	PingAction struct {
		Type string `json:"type"`
	}
	CounterAction struct {
		Type string `json:"type"`
	}
)

func NewPingAction() *PingAction {
	return &PingAction{
		Type: "ping",
	}
}

func (ctx *PingAction) GetType() string {
	return ctx.Type
}

func NewCounterAction() *CounterAction {
	return &CounterAction{
		Type: "count increase",
	}
}

func (ctx *CounterAction) GetType() string {
	return ctx.Type
}
