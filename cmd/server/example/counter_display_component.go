package example

import (
	"github.com/nodejayes/streaming-ui-server/pkg/server/ui"
	"github.com/nodejayes/streaming-ui-server/pkg/server/ui/utils"
)

type CounterDisplayComponent struct {
	utils.ViewHelper
	ID      string
	Counter int
}

func NewCounterDisplayComponent(id string, counter int) *CounterDisplayComponent {
	return &CounterDisplayComponent{
		ID:      id,
		Counter: counter,
	}
}

func (ctx *CounterDisplayComponent) Render() string {
	return ui.Render(`
	<ul id="{{ .ID }}">
		<li>{{ .Counter }}</li>
		<li>{{ .Counter }}</li>
		<li>{{ .Counter }}</li>
	</ul>
	`, ctx)
}
