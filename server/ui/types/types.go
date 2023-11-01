package types

type (
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
