package server

import (
	"bytes"
	"fmt"

	"github.com/a-h/templ"
	"github.com/nodejayes/streaming-ui-server/pkg/server/socket"
	"github.com/nodejayes/streaming-ui-server/pkg/server/ui"
)

func NewReplaceHtmlAction[TContext any](selector string, template templ.Component, ctx TContext, animation string) socket.Action[string, TContext] {
	buf := bytes.NewBuffer([]byte(fmt.Sprintf("replaceHtml::%s", selector)))
	if animation != "" {
		buf.WriteString(fmt.Sprintf("::%v", animation))
	}
	return socket.Action[string, TContext]{
		Type:    buf.String(),
		Payload: ui.RenderComponent(template),
		Context: ctx,
	}
}

func NewRedirectAction(url string, ctx ClientIdentity) socket.Action[string, ClientIdentity] {
	return socket.Action[string, ClientIdentity]{
		Type:    "redirect::",
		Payload: url,
		Context: ctx,
	}
}

func NewAlertAction(message string, ctx ClientIdentity) socket.Action[string, ClientIdentity] {
	return socket.Action[string, ClientIdentity]{
		Type:    "alert::",
		Payload: message,
		Context: ctx,
	}
}
