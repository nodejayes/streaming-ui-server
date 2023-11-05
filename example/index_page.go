package example

import (
	"fmt"
	"github.com/google/uuid"
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
	increase10Counter := NewCounterAction(10)
	increaseCounterButtonID := uuid.New().String()
	decreaseCounter := NewCounterAction(-1)
	decreaseCounterButtonID := uuid.New().String()
	return &IndexPage{
		Title: "Index Page",
		IncreaseCounterButton: components.NewButton(increaseCounterButtonID, components.NewText("+"), components.ButtonOptions{
			OnClick: server.CreateEventHandler(increaseCounter, increaseCounterButtonID, func(action types.Action, ctx ActionContext, elementID string, inputs map[string]map[string]string, eventData types.ClickEventData) {
				ctx.State.Counter += serverutils.ReadPayload[int](action)
				server.SendCaller(socket.Action[string, ActionContext]{
					Type:    "replaceHtml::#counter li",
					Payload: fmt.Sprintf("<p>%v</p>", ctx.State.Counter),
					Context: ctx,
				})
			}),
			OnStrgClick: server.CreateEventHandler(increase10Counter, increaseCounterButtonID, func(action types.Action, ctx ActionContext, elementID string, inputs map[string]map[string]string, eventData types.ClickEventData) {
				ctx.State.Counter += serverutils.ReadPayload[int](action)
				server.SendCaller(socket.Action[string, ActionContext]{
					Type:    "replaceHtml::#counter li",
					Payload: fmt.Sprintf("<p>%v</p>", ctx.State.Counter),
					Context: ctx,
				})
			}),
		}),
		DecreaseCounterButton: components.NewButton(decreaseCounterButtonID, components.NewText("-"), components.ButtonOptions{
			OnClick: server.CreateEventHandler(decreaseCounter, decreaseCounterButtonID, func(action types.Action, ctx ActionContext, elementID string, inputs map[string]map[string]string, eventData types.ClickEventData) {
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
		<body lrloadaction="ping" lrloadpayload="pong">
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
