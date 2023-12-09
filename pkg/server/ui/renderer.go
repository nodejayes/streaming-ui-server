package ui

import (
	"bytes"
	"context"

	"github.com/a-h/templ"
)

func RenderComponent(component templ.Component) string {
	buf := bytes.NewBuffer([]byte{})
	err := component.Render(context.Background(), buf)
	if err != nil {
		return err.Error()
	}
	return buf.String()
}
