package example

import (
	"fmt"
	"github.com/nodejayes/streaming-ui-server/server"
	"github.com/nodejayes/streaming-ui-server/server/socket"
	"github.com/nodejayes/streaming-ui-server/server/ui"
	"github.com/nodejayes/streaming-ui-server/server/ui/components"
	"github.com/nodejayes/streaming-ui-server/server/ui/types"
	"github.com/nodejayes/streaming-ui-server/server/ui/utils"
)

type IndexPage struct {
	utils.ViewHelper
	Title                 string
	IncreaseCounterButton types.Component
	DecreaseCounterButton types.Component
}

func NewIndexPage() *IndexPage {
	increaseCounter := NewCounterAction()
	decreaseCounter := NewCounterAction()
	return &IndexPage{
		Title: "Index Page",
		IncreaseCounterButton: components.NewButton(components.NewText("+"), components.ButtonOptions{
			OnClick: server.CreateEventHandler(increaseCounter, func(action types.Action, ctx ActionContext, elementID string, inputs map[string]map[string]string, eventData types.ClickEventData) {
				if eventData.CtrlKey {
					ctx.State.Counter += 10
				} else {
					ctx.State.Counter += 1
				}
				server.SendCaller(socket.Action[string, ActionContext]{
					Type:    "replaceHtml::#counter li",
					Payload: fmt.Sprintf("<p>%v</p>", ctx.State.Counter),
					Context: ctx,
				})
			}),
		}),
		DecreaseCounterButton: components.NewButton(components.NewText("-"), components.ButtonOptions{
			OnClick: server.CreateEventHandler(decreaseCounter, func(action types.Action, ctx ActionContext, elementID string, inputs map[string]map[string]string, eventData types.ClickEventData) {
				if eventData.CtrlKey {
					ctx.State.Counter -= 10
				} else {
					ctx.State.Counter -= 1
				}
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
