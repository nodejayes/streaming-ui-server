package example

import (
	"fmt"
	"github.com/nodejayes/streaming-ui-server/server"
	"github.com/nodejayes/streaming-ui-server/server/socket"
	"github.com/nodejayes/streaming-ui-server/server/ui"
	"github.com/nodejayes/streaming-ui-server/server/ui/components"
	"github.com/nodejayes/streaming-ui-server/server/ui/types"
	"github.com/nodejayes/streaming-ui-server/server/ui/utils"
	serverutils "github.com/nodejayes/streaming-ui-server/server/utils"
)

type IndexPage struct {
	utils.ViewHelper
	Title                 string
	IncreaseCounterButton types.Component
	DecreaseCounterButton types.Component
}

func NewIndexPage() *IndexPage {
	increaseCounter := NewCounterAction(1)
	decreaseCounter := NewCounterAction(-1)
	return &IndexPage{
		Title: "Index Page",
		IncreaseCounterButton: components.NewButton(increaseCounter.GetElementID(), components.NewText("+"), components.ButtonOptions{
			OnClick: server.CreateEventHandler(increaseCounter, func(action types.Action, ctx ActionContext) {
				ctx.State.Counter += serverutils.ReadPayload[int](action)
				server.SendCaller(socket.Action[string, ActionContext]{
					Type:    "replaceHtml::#counter li",
					Payload: fmt.Sprintf("<p>%v</p>", ctx.State.Counter),
					Context: ctx,
				})
			}),
		}),
		DecreaseCounterButton: components.NewButton(decreaseCounter.GetElementID(), components.NewText("-"), components.ButtonOptions{
			OnClick: server.CreateEventHandler(decreaseCounter, func(action types.Action, ctx ActionContext) {
				ctx.State.Counter += serverutils.ReadPayload[int](action)
				server.SendCaller(socket.Action[string, ActionContext]{
					Type:    "replaceHtml::#counter li",
					Payload: fmt.Sprintf("<p>%v</p>", ctx.State.Counter),
					Context: ctx,
				})
			}),
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
