package utils

import (
	"bytes"
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

func (ctx ViewHelper) GetStyle(style *Style) template.CSS {
	if style == nil {
		return ""
	}
	styleStr := style.GetString()
	if len(styleStr) > 0 {
		buf := bytes.NewBuffer([]byte{})
		err := template.Must(template.New("").Parse(styleStr)).Execute(buf, nil)
		if err != nil {
			return template.CSS(err.Error())
		}
		return template.CSS(buf.String())
	}
	return ""
}
