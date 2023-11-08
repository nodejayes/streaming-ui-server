package components

import "github.com/nodejayes/streaming-ui-server/pkg/server/ui"

type Text struct {
	Value string
}

func NewText(value string) *Text {
	return &Text{
		Value: value,
	}
}

func (ctx *Text) Render() string {
	return ui.Render(`{{ .Value }}`, ctx)
}
