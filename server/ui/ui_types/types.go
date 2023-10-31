package ui_types

type (
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