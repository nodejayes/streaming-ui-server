package ui

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/nodejayes/streaming-ui-server/pkg/server/ui/types"
)

func RenderComponent(component types.Component) string {
	return component.Render()
}

func Render[TData any](tmplStr string, data TData, errorTmpl ...string) string {
	htmlTemplate := template.Must(template.New("").Parse(tmplStr))
	buf := bytes.NewBuffer([]byte{})
	err := htmlTemplate.Execute(buf, data)
	if err != nil {
		if len(errorTmpl) > 0 {
			return fmt.Sprintf(errorTmpl[0], err.Error())
		}
		return fmt.Sprintf("<p>[Error render Template]: %s</p>", err.Error())
	}
	return buf.String()
}
