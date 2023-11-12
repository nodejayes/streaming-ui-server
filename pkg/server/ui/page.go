package ui

type BasePage struct {
	ID       string
	ClientID string
}

func (ctx BasePage) GetID() string {
	return ctx.ID
}

func (ctx BasePage) GetClientID() string {
	return ctx.ClientID
}
