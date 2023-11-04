package types

type (
	BaseEventData struct {
		Typ string `json:"typ"`
	}
	ClickEventData struct {
		BaseEventData
		CtrlKey bool `json:"ctrlKey"`
	}
	OtherEventData struct {
		BaseEventData
		Name string `json:"name"`
	}
	EventData interface {
		ClickEventData | OtherEventData
	}
	Action interface {
		GetType() string
		GetPayload() any
	}
	Renderer interface {
		Render() string
	}
	Component interface {
		Renderer
	}
	Page interface {
		GetPath() string
		Renderer
	}
)
