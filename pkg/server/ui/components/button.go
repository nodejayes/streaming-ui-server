package components

import (
	"github.com/google/uuid"
	"github.com/nodejayes/streaming-ui-server/pkg/server/ui"
	"github.com/nodejayes/streaming-ui-server/pkg/server/ui/types"
	"github.com/nodejayes/streaming-ui-server/pkg/server/ui/utils"
)

type (
	ButtonOptions struct {
		Class   string
		Style   *utils.Style
		OnClick func(eID string) types.Action
	}
	Button struct {
		utils.ViewHelper
		Id      string
		Content types.Component
		Options ButtonOptions
	}
)

func NewButton(content types.Component, options ButtonOptions) *Button {
	id := uuid.New().String()
	return &Button{
		Id:      id,
		Content: content,
		Options: options,
	}
}

func (ctx *Button) Render() string {
	return ui.Render(`
<button
		id="{{ .Id }}"
		{{ if .Options.Class }}
			class="{{ .Options.Class }}"
		{{ end }}		
		{{ if .Options.Style }}
			style="{{ .GetStyle .Options.Style }}"
		{{ end }}
		{{ if .Options.OnClick }}
			lrClickAction="{{ .EventType .Options.OnClick .Id }}"
		{{ end }}
		lrClickDelay="250"
	>{{ .Component .Content }}</button>`, ctx)
}
