package components

import (
	"github.com/nodejayes/streaming-ui-server/server/ui"
	"github.com/nodejayes/streaming-ui-server/server/ui/types"
	"github.com/nodejayes/streaming-ui-server/server/ui/utils"
)

type (
	ButtonOptions struct {
		OnClick     types.Action
		OnStrgClick types.Action
	}
	Button struct {
		utils.ViewHelper
		Id      string
		Content types.Component
		Options ButtonOptions
	}
)

func NewButton(id string, content types.Component, options ButtonOptions) *Button {
	return &Button{
		Id:      id,
		Content: content,
		Options: options,
	}
}

func (ctx *Button) Render() string {
	return ui.Render(`
{{ if .Options.OnStrgClick }}
	<script>
		function strgFilter(e) {
			return e.ctrlKey;
		}
	</script>
{{ end }}
<button
		id="{{ .Id }}"
		{{ if .Options.OnClick }}
			lrClickAction="{{ .EventType .Options.OnClick }}"
			lrClickPayload="{{ .EventPayload .Options.OnClick }}"
		{{ end }}
		lrClickDelay="500"
		{{ if .Options.OnStrgClick }}
			lrClickFilter="strgFilter"
			{{ if not (eq .Options.OnStrgClick.GetType .Options.OnClick.GetType) }}
			lrClickFilterAction="{{ .EventType .Options.OnStrgClick }}"
			{{ end }}
			lrClickFilterPayload="{{ .EventPayload .Options.OnStrgClick }}"
		{{ end }}
	>{{ .Component .Content }}</button>`, ctx)
}
