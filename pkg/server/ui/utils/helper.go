package utils

import (
	"html/template"

	"github.com/nodejayes/streaming-ui-server/pkg/server/ui/types"
)

type ViewHelper struct{}

func (ctx ViewHelper) Component(component types.Component) template.HTML {
	return template.HTML(component.Render())
}

func (ctx ViewHelper) EventType(action types.Action) string {
	return action.GetType()
}
