package example

type ActionContext struct {
	ID    string
	State *AppState
}

func (ctx ActionContext) GetClientId() string {
	return ctx.ID
}
