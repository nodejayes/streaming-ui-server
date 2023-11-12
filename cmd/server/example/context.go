package example

type ActionContext struct {
	PageID   string
	ClientID string
	State    *AppState
}

func (ctx ActionContext) GetClientId() string {
	return ctx.ClientID
}

func (ctx ActionContext) GetPageId() string {
	return ctx.PageID
}
