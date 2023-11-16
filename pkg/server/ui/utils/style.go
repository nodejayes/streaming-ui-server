package utils

import (
	"bytes"
	"fmt"
)

type Style struct {
	Padding         string
	Margin          string
	Color           string
	BackgroundColor string
	Font            string
	FontWeight      string
	FontSize        string
	FontFamily      string
	Border          string
	BorderRadius    string
	Width           string
	Height          string
	Display         string
}

func (ctx *Style) GetString() string {
	buf := bytes.NewBuffer([]byte{})

	if len(ctx.Display) > 0 {
		buf.WriteString(fmt.Sprintf("display:%s;", ctx.Display))
	}

	if len(ctx.Padding) > 0 {
		buf.WriteString(fmt.Sprintf("padding:%s;", ctx.Padding))
	}

	if len(ctx.Margin) > 0 {
		buf.WriteString(fmt.Sprintf("margin:%s;", ctx.Margin))
	}

	if len(ctx.Color) > 0 {
		buf.WriteString(fmt.Sprintf("color:%s;", ctx.Color))
	}

	if len(ctx.BackgroundColor) > 0 {
		buf.WriteString(fmt.Sprintf("background-color:%s;", ctx.BackgroundColor))
	}

	if len(ctx.Font) > 0 {
		buf.WriteString(fmt.Sprintf("font:%s;", ctx.Font))
	}

	if len(ctx.FontWeight) > 0 {
		buf.WriteString(fmt.Sprintf("font-weight:%s;", ctx.FontWeight))
	}

	if len(ctx.FontSize) > 0 {
		buf.WriteString(fmt.Sprintf("font-size:%s;", ctx.FontSize))
	}

	if len(ctx.FontFamily) > 0 {
		buf.WriteString(fmt.Sprintf("font-family:%s;", ctx.FontFamily))
	}

	if len(ctx.Border) > 0 {
		buf.WriteString(fmt.Sprintf("border:%s;", ctx.Border))
	}

	if len(ctx.BorderRadius) > 0 {
		buf.WriteString(fmt.Sprintf("border-radius:%s;", ctx.BorderRadius))
	}

	if len(ctx.Width) > 0 {
		buf.WriteString(fmt.Sprintf("width:%s;", ctx.Width))
	}

	if len(ctx.Height) > 0 {
		buf.WriteString(fmt.Sprintf("height:%s;", ctx.Height))
	}

	return buf.String()
}
