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
		GetElementID() string
		SetElementID(value string)
		GetType() string
		GetPayload() any
		GetInputs() map[string]map[string]string
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
