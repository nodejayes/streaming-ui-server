package server

import (
	"fmt"
	"github.com/nodejayes/streaming-ui-server/pkg/server/socket"
	"github.com/nodejayes/streaming-ui-server/pkg/server/ui/types"
)

func NewReplaceHtmlAction[TContext any](selector string, template types.Component, ctx TContext) socket.Action[string, TContext] {
	return socket.Action[string, TContext]{
		Type:    fmt.Sprintf("replaceHtml::%s", selector),
		Payload: template.Render(),
		Context: ctx,
	}
}

func NewRedirectAction(url string, ctx ClientIdentiy) socket.Action[string, ClientIdentiy] {
	return socket.Action[string, ClientIdentiy]{
		Type:    "redirect::",
		Payload: url,
		Context: ctx,
	}
}