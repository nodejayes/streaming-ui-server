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
	"github.com/nodejayes/streaming-ui-server/pkg/ui_components/base"
)

type IndexPage struct {
	Context         *gin.Context
	Title           string
	BlueButtonStyle *utils.Style
	RedButtonStyle  *utils.Style
	HeadlineStyle   *utils.Style
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
		HeadlineStyle: &utils.Style{
			Cursor: "pointer",
		},
	}
}

func (ctx *IndexPage) Render() string {
	return ui.RenderComponent(components.Index("Index Page", components.IndexOptions{
		HeaderSection: ctx.GetHeaderSection(),
		MainSection:   ctx.GetMainSection(),
	}))
}

func (ctx *IndexPage) GetHeaderSection() templ.Component {
	elementID := uuid.NewString()
	onClickHandler := server.CreateEventHandler(ReloadAction, func(action string, actx ActionContext, elementID string, inputs map[string]map[string]string, eventData types.ClickEventData) {
		server.SendCaller(server.NewRedirectAction("/", actx))
	})
	onDoubleClickHandler := server.CreateEventHandler(DoubleClickNoticeAction, func(action string, actx ActionContext, elementID string, inputs map[string]map[string]string, eventData types.ClickEventData) {
		server.SendCaller(server.NewAlertAction("Double Click Reload", actx))
	})
	onContextMenuOpenHandler := server.CreateEventHandler(ContextMenuNoticeAction, func(action string, actx ActionContext, elementID string, inputs map[string]map[string]string, eventData types.ClickEventData) {
		server.SendCaller(server.NewAlertAction("Context Menu open", actx))
	})
	return base.Section(
		base.SectionOptions{
			ID: uuid.NewString(),
		},
		base.Headline(base.HeadlineOptions{
			ID:            elementID,
			Style:         ctx.HeadlineStyle.GetString(),
			OnClick:       onClickHandler(elementID),
			OnDoubleClick: onDoubleClickHandler(elementID),
			OnContextMenu: onContextMenuOpenHandler(elementID),
		}, base.Text(ctx.Title)))
}

func (ctx *IndexPage) GetMainSection() templ.Component {
	return base.Section(base.SectionOptions{
		ID: "counters",
	},
		components.Main(ctx.GetCounterDisplay(), ctx.GetIncreaseCounterButton(), ctx.GetDecreaseCounterButton()))
}

func (ctx *IndexPage) GetIncreaseCounterButton() templ.Component {
	elementID := uuid.NewString()
	onClickHandler := server.CreateEventHandler(CounterAction, func(action string, actx ActionContext, elementID string, inputs map[string]map[string]string, eventData types.ClickEventData) {
		if eventData.CtrlKey {
			actx.State.Counter += 10
		} else {
			actx.State.Counter += 1
		}
		server.SendCaller(server.NewReplaceHtmlAction("#counters > ul li", base.Text(fmt.Sprintf("%v", actx.State.Counter)), actx, utils.NewStartAnimation("fadeInDown", 250)))
	})
	return base.Button(base.ButtonOptions{
		ID:      elementID,
		Class:   "stdWidth",
		Style:   ctx.BlueButtonStyle.GetString(),
		OnClick: onClickHandler(elementID),
	}, base.Text("+"))
}

func (ctx *IndexPage) GetDecreaseCounterButton() templ.Component {
	elementID := uuid.NewString()
	onClickHandler := server.CreateEventHandler(CounterAction, func(action string, actx ActionContext, elementID string, inputs map[string]map[string]string, eventData types.ClickEventData) {
		if eventData.CtrlKey {
			actx.State.Counter -= 10
		} else {
			actx.State.Counter -= 1
		}
		server.SendCaller(server.NewReplaceHtmlAction("#counters > ul li", base.Text(fmt.Sprintf("%v", actx.State.Counter)), actx, utils.NewStartAnimation("fadeOutDown", 250)))
	})
	return base.Button(base.ButtonOptions{
		ID:      elementID,
		Class:   "stdWidth",
		Style:   ctx.BlueButtonStyle.GetString(),
		OnClick: onClickHandler(elementID),
	}, base.Text("-"))
}

func (ctx *IndexPage) GetCounterDisplay() templ.Component {
	clientId, err := ctx.Context.Cookie("ClientId")
	if err != nil {
		return base.Text("Error: " + err.Error())
	}
	state := di.Inject[AppState](clientId)
	return components.CounterDisplay("counter", state.Counter)
}
