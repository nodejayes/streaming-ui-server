package utils

import (
	"html/template"

	"github.com/nodejayes/streaming-ui-server/server/ui/ui_types"
)

type ViewHelper struct{}

func (ctx ViewHelper) Component(component ui_types.Component) template.HTML {
	return template.HTML(component.Render())
}
