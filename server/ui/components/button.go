package components

import (
	"github.com/nodejayes/streaming-ui-server/server/ui"
	"github.com/nodejayes/streaming-ui-server/server/ui/ui_types"
	"github.com/nodejayes/streaming-ui-server/server/ui/utils"
)

type (
	ButtonOptions struct {
		OnClick string
		ClickPayload string
	}
	Button struct {
		utils.ViewHelper
		Content ui_types.Component
		Options ButtonOptions
	}
)

func NewButton(content ui_types.Component, options ButtonOptions) *Button {
	return &Button{
		Content: content,
		Options: options,
	}
}

func (ctx *Button) Render() string {
	return ui.Render(`<button {{if .Options.OnClick }} lrClickAction="{{ .Options.OnClick }}" lrClickPayload="{{ .Options.ClickPayload }}" {{ end }}>{{ .Component .Content }}</button>`, ctx)
}