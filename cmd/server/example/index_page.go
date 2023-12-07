package example

import (
	"github.com/gin-gonic/gin"
	di "github.com/nodejayes/generic-di"
	"github.com/nodejayes/streaming-ui-server/pkg/server"
	"github.com/nodejayes/streaming-ui-server/pkg/server/ui"
	"github.com/nodejayes/streaming-ui-server/pkg/server/ui/components"
	"github.com/nodejayes/streaming-ui-server/pkg/server/ui/types"
	"github.com/nodejayes/streaming-ui-server/pkg/server/ui/utils"
)

type IndexPage struct {
	utils.ViewHelper
	Context               *gin.Context
	Title                 string
	OnLoad                func(eID string) types.Action
	ReloadButton          types.Component
	IncreaseCounterButton types.Component
	DecreaseCounterButton types.Component
	CounterDisplay        types.Component
}

func NewIndexPage(ctx *gin.Context) *IndexPage {
	page := &IndexPage{}
	page.Context = ctx
	page.Title = "Index Page"
	page.ReloadButton = page.GetReloadButton()
	page.IncreaseCounterButton = page.GetIncreaseCounterButton()
	page.DecreaseCounterButton = page.GetDecreaseCounterButton()
	page.CounterDisplay = page.GetCounterDisplay()
	return page
}

func (ctx *IndexPage) GetReloadButton() types.Component {
	return components.NewButton(components.NewText("+"), components.ButtonOptions{
		OnClick: server.CreateEventHandler(NewReloadAction(), func(action types.Action, actx ActionContext, elementID string, inputs map[string]map[string]string, eventData types.ClickEventData) {
			server.SendCaller(server.NewRedirectAction("/", actx))
		}),
	})
}

func (ctx *IndexPage) GetIncreaseCounterButton() types.Component {
	return components.NewButton(components.NewText("+"), components.ButtonOptions{
		Style: &utils.Style{
			BackgroundColor: "blue",
			Color:           "grey",
		},
		OnClick: server.CreateEventHandler(NewCounterAction(), func(action types.Action, actx ActionContext, elementID string, inputs map[string]map[string]string, eventData types.ClickEventData) {
			if eventData.CtrlKey {
				actx.State.Counter += 10
			} else {
				actx.State.Counter += 1
			}
			server.SendCaller(server.NewReplaceHtmlAction("#counters", ctx.GetCounterDisplay(), actx))
		}),
	})
}

func (ctx *IndexPage) GetDecreaseCounterButton() types.Component {
	return components.NewButton(components.NewText("-"), components.ButtonOptions{
		Class: "stdWidth",
		Style: &utils.Style{
			BackgroundColor: "red",
			Color:           "grey",
		},
		OnClick: server.CreateEventHandler(NewCounterAction(), func(action types.Action, actx ActionContext, elementID string, inputs map[string]map[string]string, eventData types.ClickEventData) {
			if eventData.CtrlKey {
				actx.State.Counter -= 10
			} else {
				actx.State.Counter -= 1
			}
			server.SendCaller(server.NewReplaceHtmlAction("#counters", ctx.GetCounterDisplay(), actx))
		}),
	})
}

func (ctx *IndexPage) GetCounterDisplay() types.Component {
	clientId, err := ctx.Context.Cookie("ClientId")
	if err != nil {
		return components.NewText("Error: " + err.Error())
	}
	state := di.Inject[AppState](clientId)
	return NewCounterDisplayComponent("counter", state.Counter)
}

func (ctx *IndexPage) Render() string {
	return ui.Render(`
	<!DOCTYPE html>
	<html>
		<head>
			<title>{{ .Title }}</title>
			<script src="/live-replacer/lib/bundle.js"></script>
			<style>
				.stdWidth {
					width: 150px;
				}
			</style>
		</head>
		<body>
			<div id="header">
				<h1>{{ .Title }}</h1>
				{{ .Component .ReloadButton }}
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
