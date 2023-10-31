package example

import (
	"github.com/nodejayes/streaming-ui-server/server/ui"
	"github.com/nodejayes/streaming-ui-server/server/ui/components"
	"github.com/nodejayes/streaming-ui-server/server/ui/ui_types"
	"github.com/nodejayes/streaming-ui-server/server/ui/utils"
)

type IndexPage struct {
	utils.ViewHelper
	Title                 string
	IncreaseCounterButton ui_types.Component
	DecreaseCounterButton ui_types.Component
}

func NewIndexPage() *IndexPage {
	return &IndexPage{
		Title: "Index Page",
		IncreaseCounterButton: components.NewButton(components.NewText("+"), components.ButtonOptions{
			OnClick:      "count increase",
			ClickPayload: "1",
		}),
		DecreaseCounterButton: components.NewButton(components.NewText("-"), components.ButtonOptions{
			OnClick:      "count increase",
			ClickPayload: "-1",
		}),
	}
}

func (ctx *IndexPage) GetPath() string {
	return "/"
}

func (ctx *IndexPage) Render() string {
	return ui.Render(`
	<!DOCTYPE html>
	<html>
		<head>
			<title>{{ .Title }}</title>
			<script src="/live-replacer/lib/bundle.js"></script>
		</head>
		<body>
			<div id="header">
				<h1>{{ .Title }}</h1>
			</div>
			<ul id="counter">
				<li>
					<p>0</p>
				</li>
				<li>
					<p>0</p>
				</li>
				<li>
					<p>0</p>
				</li>
			</ul>
			{{ .Component .IncreaseCounterButton }}
			{{ .Component .DecreaseCounterButton }}
		</body>
	</html>`, ctx)
}
