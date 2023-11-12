package example

import (
	"github.com/google/uuid"
	di "github.com/nodejayes/generic-di"
	"github.com/nodejayes/streaming-ui-server/pkg/server"
	"github.com/nodejayes/streaming-ui-server/pkg/server/socket"
	"github.com/nodejayes/streaming-ui-server/pkg/server/ui"
	"github.com/nodejayes/streaming-ui-server/pkg/server/ui/components"
	"github.com/nodejayes/streaming-ui-server/pkg/server/ui/types"
	"github.com/nodejayes/streaming-ui-server/pkg/server/ui/utils"
)

type IndexPage struct {
	ui.BasePage
	utils.ViewHelper
	Title                 string
	OnLoad                func(eID string) types.Action
	IncreaseCounterButton types.Component
	DecreaseCounterButton types.Component
	CounterDisplay        types.Component
}

func NewIndexPage() *IndexPage {
	page := &IndexPage{}
	page.ID = uuid.NewString()
	page.Title = "Index Page"
	page.IncreaseCounterButton = page.GetIncreaseCounterButton()
	page.DecreaseCounterButton = page.GetDecreaseCounterButton()
	page.CounterDisplay = page.GetCounterDisplay()
	return page
}

func (ctx *IndexPage) GetPath() string {
	return "/"
}

func (ctx *IndexPage) GetIncreaseCounterButton() types.Component {
	increaseCounter := NewCounterAction()
	return components.NewButton(components.NewText("+"), components.ButtonOptions{
		OnClick: server.CreateEventHandler(increaseCounter, ctx.ID, func(action types.Action, actx ActionContext, elementID string, inputs map[string]map[string]string, eventData types.ClickEventData) {
			if eventData.CtrlKey {
				actx.State.Counter += 10
			} else {
				actx.State.Counter += 1
			}
			server.SendCaller(socket.Action[string, ActionContext]{
				Type:    "replaceHtml::#counters",
				Payload: ctx.GetCounterDisplay().Render(),
				Context: actx,
			})
		}),
	})
}

func (ctx *IndexPage) GetDecreaseCounterButton() types.Component {
	decreaseCounter := NewCounterAction()
	return components.NewButton(components.NewText("-"), components.ButtonOptions{
		OnClick: server.CreateEventHandler(decreaseCounter, ctx.ID, func(action types.Action, actx ActionContext, elementID string, inputs map[string]map[string]string, eventData types.ClickEventData) {
			if eventData.CtrlKey {
				actx.State.Counter -= 10
			} else {
				actx.State.Counter -= 1
			}
			server.SendCaller(socket.Action[string, ActionContext]{
				Type:    "replaceHtml::#counters",
				Payload: ctx.GetCounterDisplay().Render(),
				Context: actx,
			})
		}),
	})
}

func (ctx *IndexPage) GetCounterDisplay() types.Component {
	state := di.Inject[AppState](ctx.ID)
	return NewCounterDisplayComponent("counter", state.Counter)
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
			<div id="counters">
				{{ .Component .CounterDisplay }}
			</div>
			{{ .Component .IncreaseCounterButton }}
			{{ .Component .DecreaseCounterButton }}
			<div style="width:150px;height:150px;background-color:green;" lrmouseupaction="count increase">
			</div>
		</body>
	</html>`, ctx)
}
