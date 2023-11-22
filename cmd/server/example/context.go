package example

type ActionContext struct {
	ClientID string
	State    *AppState
}

func (ctx ActionContext) GetClientId() string {
	return ctx.ClientID
}
