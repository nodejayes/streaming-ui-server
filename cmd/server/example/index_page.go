package example

import (
	"fmt"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	di "github.com/nodejayes/generic-di"
	"github.com/nodejayes/streaming-ui-server/cmd/server/example/components"
	"github.com/nodejayes/streaming-ui-server/pkg/server"
	"github.com/nodejayes/streaming-ui-server/pkg/server/ui"
	"github.com/nodejayes/streaming-ui-server/pkg/server/ui/types"
	"github.com/nodejayes/streaming-ui-server/pkg/server/ui/utils"
)

type IndexPage struct {
	Context         *gin.Context
	Title           string
	BlueButtonStyle *utils.Style
	RedButtonStyle  *utils.Style
}

func NewIndexPage(ctx *gin.Context) *IndexPage {
	return &IndexPage{
		Title:   "Index Page",
		Context: ctx,
		BlueButtonStyle: &utils.Style{
			BackgroundColor: "blue",
			Color:           "grey",
		},
		RedButtonStyle: &utils.Style{
			BackgroundColor: "red",
			Color:           "grey",
		},
	}
}

func (ctx *IndexPage) Render() string {
	return ui.RenderComponent(components.Index("Index Page", components.IndexOptions{
		ReloadButton:   ctx.GetReloadButton(),
		IncreaseButton: ctx.GetIncreaseCounterButton(),
		DecreaseButton: ctx.GetDecreaseCounterButton(),
		CounterDisplay: ctx.GetCounterDisplay(),
	}))
}

func (ctx *IndexPage) GetReloadButton() templ.Component {
	elementID := uuid.NewString()
	onClickHandler := server.CreateEventHandler(NewReloadAction(), func(action types.Action, actx ActionContext, elementID string, inputs map[string]map[string]string, eventData types.ClickEventData) {
		// server.SendCaller(server.NewRedirectAction("/", actx))
		server.SendCaller(server.NewAlertAction("Reload", actx))
	})
	onDoubleClickHandler := server.CreateEventHandler(NewDoubleClickNoticeAction(), func(action types.Action, actx ActionContext, elementID string, inputs map[string]map[string]string, eventData types.ClickEventData) {
		server.SendCaller(server.NewAlertAction("Double Click Reload", actx))
	})
	onContextMenuOpenHandler := server.CreateEventHandler(NewContextMenuNoticeAction(), func(action types.Action, actx ActionContext, elementID string, inputs map[string]map[string]string, eventData types.ClickEventData) {
		server.SendCaller(server.NewAlertAction("Context Menu open", actx))
	})
	return components.Button(components.ButtonOptions{
		ID:            elementID,
		Style:         ctx.RedButtonStyle.GetString(),
		OnClick:       onClickHandler(elementID).GetType(),
		OnDoubleClick: onDoubleClickHandler(elementID).GetType(),
		OnContextMenu: onContextMenuOpenHandler(elementID).GetType(),
	}, "Reload")
}

func (ctx *IndexPage) GetIncreaseCounterButton() templ.Component {
	elementID := uuid.NewString()
	onClickHandler := server.CreateEventHandler(NewCounterAction(), func(action types.Action, actx ActionContext, elementID string, inputs map[string]map[string]string, eventData types.ClickEventData) {
		if eventData.CtrlKey {
			actx.State.Counter += 10
		} else {
			actx.State.Counter += 1
		}
		server.SendCaller(server.NewReplaceHtmlAction("#counters > ul li", components.Text(fmt.Sprintf("%v", actx.State.Counter)), actx, utils.NewStartAnimation("fadeInDown", 250)))
	})
	return components.Button(components.ButtonOptions{
		ID:      elementID,
		Class:   "stdWidth",
		Style:   ctx.BlueButtonStyle.GetString(),
		OnClick: onClickHandler(elementID).GetType(),
	}, "+")
}

func (ctx *IndexPage) GetDecreaseCounterButton() templ.Component {
	elementID := uuid.NewString()
	onClickHandler := server.CreateEventHandler(NewCounterAction(), func(action types.Action, actx ActionContext, elementID string, inputs map[string]map[string]string, eventData types.ClickEventData) {
		if eventData.CtrlKey {
			actx.State.Counter -= 10
		} else {
			actx.State.Counter -= 1
		}
		server.SendCaller(server.NewReplaceHtmlAction("#counters > ul li", components.Text(fmt.Sprintf("%v", actx.State.Counter)), actx, utils.NewStartAnimation("fadeOutDown", 250)))
	})
	return components.Button(components.ButtonOptions{
		ID:      elementID,
		Class:   "stdWidth",
		Style:   ctx.BlueButtonStyle.GetString(),
		OnClick: onClickHandler(elementID).GetType(),
	}, "-")
}

func (ctx *IndexPage) GetCounterDisplay() templ.Component {
	clientId, err := ctx.Context.Cookie("ClientId")
	if err != nil {
		return components.Text("Error: " + err.Error())
	}
	state := di.Inject[AppState](clientId)
	return components.CounterDisplay("counter", state.Counter)
}
