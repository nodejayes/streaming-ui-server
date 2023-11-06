package components

import (
	"github.com/google/uuid"
	"github.com/nodejayes/streaming-ui-server/server/ui"
	"github.com/nodejayes/streaming-ui-server/server/ui/types"
	"github.com/nodejayes/streaming-ui-server/server/ui/utils"
)

type (
	ButtonOptions struct {
		OnClick func(eID string) types.Action
	}
	Button struct {
		utils.ViewHelper
		Id      string
		Content types.Component
		Options struct {
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
			OnClick types.Action
		}{
			OnClick: options.OnClick(id),
		},
	}
}

func (ctx *Button) Render() string {
	return ui.Render(`
<button
		id="{{ .Id }}"
		{{ if .Options.OnClick }}
			lrClickAction="{{ .EventType .Options.OnClick }}"
		{{ end }}
		lrClickDelay="250"
	>{{ .Component .Content }}</button>`, ctx)
}
