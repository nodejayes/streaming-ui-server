package components

import (
	"github.com/google/uuid"
	"github.com/nodejayes/streaming-ui-server/pkg/server/ui"
	"github.com/nodejayes/streaming-ui-server/pkg/server/ui/types"
	"github.com/nodejayes/streaming-ui-server/pkg/server/ui/utils"
)

type (
	ButtonOptions struct {
		Style   *utils.Style
		OnClick func(eID string) types.Action
	}
	Button struct {
		utils.ViewHelper
		Id      string
		Content types.Component
		Options struct {
			Style   *utils.Style
			OnClick types.Action
		}
	}
)

func NewButton(content types.Component, options ButtonOptions) *Button {
	id := uuid.New().String()
	return &Button{
		Id:      id,
		Content: content,
		Options: struct {
			Style   *utils.Style
			OnClick types.Action
		}{
			Style:   options.Style,
			OnClick: options.OnClick(id),
		},
	}
}

func (ctx *Button) Render() string {
	return ui.Render(`
<button
		id="{{ .Id }}"
		{{ if .Options.Style }}
			style="{{ .GetStyle .Options.Style }}"
		{{ end }}
		{{ if .Options.OnClick }}
			lrClickAction="{{ .EventType .Options.OnClick }}"
		{{ end }}
		lrClickDelay="250"
	>{{ .Component .Content }}</button>`, ctx)
}
