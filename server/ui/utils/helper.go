package utils

import (
	"encoding/json"
	"html/template"

	"github.com/nodejayes/streaming-ui-server/server/ui/types"
)

type ViewHelper struct{}

func (ctx ViewHelper) Component(component types.Component) template.HTML {
	return template.HTML(component.Render())
}

func (ctx ViewHelper) EventType(action types.Action) string {
	return action.GetType()
}

func (ctx ViewHelper) EventPayload(action types.Action) string {
	payload := action.GetPayload()
	str, err := json.Marshal(payload)
	if err != nil {
		return ""
	}
	return string(str)
}
