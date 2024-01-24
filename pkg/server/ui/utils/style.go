package utils

import (
	"bytes"
	"fmt"
)

type Style struct {
	AlignContent             string
	AlignSelf                string
	Animation                string
	AnimationDelay           string
	AnimationDirection       string
	AnimationDuration        string
	AnimationFillMode        string
	AnimationIterationCount  string
	AnimationName            string
	AnimationPlayState       string
	AnimationTimingFunction  string
	BackfaceVisibility       string
	Background               string
	BackgroundAttachment     string
	BackgroundBlendMode      string
	BackgroundClip           string
	BackgroundImage          string
	BackgroundOrigin         string
	BackgroundPosition       string
	BackgroundRepeat         string
	BackgroundSize           string
	BorderBottom             string
	BorderBottomColor        string
	BorderBottomLeftRadius   string
	BorderBottomRightRadius  string
	BorderBottomStyle        string
	BorderBottomWidth        string
	BorderCollapse           string
	BorderColor              string
	BorderImage              string
	BorderImageOutset        string
	BorderImageRepeat        string
	BorderImageSlice         string
	BorderImageSource        string
	BorderImageWidth         string
	BorderLeft               string
	BorderLeftColor          string
	BorderLeftStyle          string
	BorderLeftWidth          string
	BorderRight              string
	BorderRightColor         string
	BorderRightStyle         string
	BorderRightWidth         string
	BorderSpacing            string
	BorderStyle              string
	BorderTop                string
	BorderTopColor           string
	BorderTopLeftRadius      string
	BorderTopRightRadius     string
	BorderTopStyle           string
	BorderTopWidth           string
	BorderWidth              string
	Bottom                   string
	BoxShadow                string
	CaptionSide              string
	Clear                    string
	Clip                     string
	ColumnCount              string
	ColumnFill               string
	ColumnGap                string
	ColumnRule               string
	ColumnRuleColor          string
	ColumnRuleStyle          string
	ColumnRuleWidth          string
	ColumnSpan               string
	ColumnWidth              string
	Columns                  string
	Content                  string
	CounterIncrement         string
	CounterReset             string
	Direction                string
	EmptyCells               string
	Filter                   string
	Flex                     string
	FlexBasis                string
	FlexDirection            string
	FlexFlow                 string
	FlexGrow                 string
	FlexShrink               string
	FlexWrap                 string
	Float                    string
	FontSizeAdjust           string
	HangingPunctuation       string
	JustifyContent           string
	Left                     string
	ListStyle                string
	ListStyleImage           string
	ListStylePosition        string
	ListStyleType            string
	MarginBottom             string
	MarginTop                string
	MarginLeft               string
	MarginRight              string
	MaxHeight                string
	MaxWidth                 string
	MinHeight                string
	MinWidth                 string
	NavDown                  string
	NavIndex                 string
	NavLeft                  string
	NavRight                 string
	NavUp                    string
	Opacity                  string
	Order                    string
	Outline                  string
	OutlineColor             string
	OutlineOffset            string
	OutlineStyle             string
	OutlineWidth             string
	Overflow                 string
	OverflowX                string
	OverflowY                string
	PaddingBottom            string
	PaddingLeft              string
	PaddingRight             string
	PaddingTop               string
	PageBreakAfter           string
	PageBreakBefore          string
	PageBreakInside          string
	Perspective              string
	PerspectiveOrigin        string
	Position                 string
	Quotes                   string
	Resize                   string
	Right                    string
	TabSize                  string
	TableLayout              string
	TextAlignLast            string
	TextDecoration           string
	TextOverflow             string
	Top                      string
	Transform                string
	TransformOrigin          string
	TransformStyle           string
	Transition               string
	TransitionDelay          string
	TransitionDuration       string
	TransitionProperty       string
	TransitionTimingFunction string
	UnicodeBidi              string
	UserSelect               string
	VerticalAlign            string
	Visibility               string
	WhiteSpace               string
	WordBreak                string
	WordWrap                 string
	ZIndex                   string
	BackgroundColor          string
	Border                   string
	BorderRadius             string
	Color                    string
	Display                  string
	Font                     string
	FontFamily               string
	FontSize                 string
	FontStyle                string
	FontVariantLigatures     string
	FontVariantCaps          string
	FontVariantNumeric       string
	FontVariantEastAsian     string
	FontVariantAlternates    string
	FontVariantPosition      string
	FontStretch              string
	FontOpticalSizing        string
	FontKerning              string
	FontFeatureSettings      string
	FontVariationSettings    string
	FontWeight               string
	TextRendering            string
	LetterSpacing            string
	WordSpacing              string
	LineHeight               string
	TextTransform            string
	TextIndent               string
	TextShadow               string
	TextAlign                string
	AlignItems               string
	Cursor                   string
	BoxSizing                string
	PaddingBlock             string
	PaddingInline            string
	Height                   string
	Margin                   string
	Padding                  string
	Width                    string
}

func (ctx *Style) GetString() string {
	buf := bytes.NewBuffer([]byte{})

	writeBuffer(buf, ctx.AlignContent, "align-content")
	writeBuffer(buf, ctx.AlignItems, "align-items")
	writeBuffer(buf, ctx.AlignSelf, "align-self")
	writeBuffer(buf, ctx.Animation, "animation")
	writeBuffer(buf, ctx.AnimationDelay, "animation-delay")
	writeBuffer(buf, ctx.AnimationDirection, "animation-direction")
	writeBuffer(buf, ctx.AnimationDuration, "animation-duration")
	writeBuffer(buf, ctx.AnimationFillMode, "animation-fill-mode")
	writeBuffer(buf, ctx.AnimationIterationCount, "animation-iteration-count")
	writeBuffer(buf, ctx.AnimationName, "animation-name")
	writeBuffer(buf, ctx.AnimationPlayState, "animation-play-state")
	writeBuffer(buf, ctx.AnimationTimingFunction, "animation-timing-function")
	writeBuffer(buf, ctx.BackfaceVisibility, "backface-visibility")
	writeBuffer(buf, ctx.Background, "background")
	writeBuffer(buf, ctx.BackgroundAttachment, "background-attachment")
	writeBuffer(buf, ctx.BackgroundBlendMode, "background-blend-mode")
	writeBuffer(buf, ctx.BackgroundClip, "background-clip")
	writeBuffer(buf, ctx.BackgroundColor, "background-color")
	writeBuffer(buf, ctx.BackgroundImage, "background-image")
	writeBuffer(buf, ctx.BackgroundOrigin, "background-origin")
	writeBuffer(buf, ctx.BackgroundPosition, "background-position")
	writeBuffer(buf, ctx.BackgroundRepeat, "background-repeat")
	writeBuffer(buf, ctx.BackgroundSize, "background-size")
	writeBuffer(buf, ctx.BorderBottom, "border-bottom")
	writeBuffer(buf, ctx.BorderBottomColor, "border-bottom-color")
	writeBuffer(buf, ctx.BorderBottomLeftRadius, "border-bottom-left-radius")
	writeBuffer(buf, ctx.BorderBottomRightRadius, "border-bottom-right-radius")
	writeBuffer(buf, ctx.BorderBottomStyle, "border-bottom-style")
	writeBuffer(buf, ctx.BorderBottomWidth, "border-bottom-width")
	writeBuffer(buf, ctx.BorderCollapse, "border-collapse")
	writeBuffer(buf, ctx.BorderColor, "border-color")
	writeBuffer(buf, ctx.BorderImage, "border-image")
	writeBuffer(buf, ctx.BorderImageOutset, "border-image-outset")
	writeBuffer(buf, ctx.BorderImageRepeat, "border-image-repeat")
	writeBuffer(buf, ctx.BorderImageSlice, "border-image-slice")
	writeBuffer(buf, ctx.BorderImageSource, "border-image-source")
	writeBuffer(buf, ctx.BorderImageWidth, "border-image-width")
	writeBuffer(buf, ctx.BorderLeft, "border-left")
	writeBuffer(buf, ctx.BorderLeftColor, "border-left-color")
	writeBuffer(buf, ctx.BorderLeftStyle, "border-left-style")
	writeBuffer(buf, ctx.BorderLeftWidth, "border-left-width")
	writeBuffer(buf, ctx.BorderRight, "border-right")
	writeBuffer(buf, ctx.BorderRightColor, "border-right-color")
	writeBuffer(buf, ctx.BorderRightStyle, "border-right-style")
	writeBuffer(buf, ctx.BorderRightWidth, "border-right-width")
	writeBuffer(buf, ctx.BorderSpacing, "border-spacing")
	writeBuffer(buf, ctx.BorderStyle, "border-style")
	writeBuffer(buf, ctx.BorderTop, "border-top")
	writeBuffer(buf, ctx.BorderTopColor, "border-top-color")
	writeBuffer(buf, ctx.BorderTopLeftRadius, "border-top-left-radius")
	writeBuffer(buf, ctx.BorderTopRightRadius, "border-top-right-radius")
	writeBuffer(buf, ctx.BorderTopStyle, "border-top-style")
	writeBuffer(buf, ctx.BorderTopWidth, "border-top-width")
	writeBuffer(buf, ctx.BorderWidth, "border-width")
	writeBuffer(buf, ctx.Border, "border")
	writeBuffer(buf, ctx.BoxSizing, "box-sizing")
	writeBuffer(buf, ctx.BorderRadius, "border-radius")
	writeBuffer(buf, ctx.Bottom, "bottom")
	writeBuffer(buf, ctx.BoxShadow, "box-shadow")
	writeBuffer(buf, ctx.CaptionSide, "caption-side")
	writeBuffer(buf, ctx.Clear, "clear")
	writeBuffer(buf, ctx.Clip, "clip")
	writeBuffer(buf, ctx.Color, "color")
	writeBuffer(buf, ctx.ColumnCount, "column-count")
	writeBuffer(buf, ctx.ColumnFill, "column-fill")
	writeBuffer(buf, ctx.ColumnGap, "column-gap")
	writeBuffer(buf, ctx.ColumnRule, "column-rule")
	writeBuffer(buf, ctx.ColumnRuleColor, "column-rule-color")
	writeBuffer(buf, ctx.ColumnRuleStyle, "column-rule-style")
	writeBuffer(buf, ctx.ColumnRuleWidth, "column-rule-width")
	writeBuffer(buf, ctx.ColumnSpan, "column-span")
	writeBuffer(buf, ctx.ColumnWidth, "column-width")
	writeBuffer(buf, ctx.Columns, "columns")
	writeBuffer(buf, ctx.Content, "content")
	writeBuffer(buf, ctx.CounterIncrement, "counter-increment")
	writeBuffer(buf, ctx.CounterReset, "counter-reset")
	writeBuffer(buf, ctx.Cursor, "cursor")
	writeBuffer(buf, ctx.Direction, "direction")
	writeBuffer(buf, ctx.Display, "display")
	writeBuffer(buf, ctx.EmptyCells, "empty-cells")
	writeBuffer(buf, ctx.Filter, "filter")
	writeBuffer(buf, ctx.Flex, "flex")
	writeBuffer(buf, ctx.FlexBasis, "flex-basis")
	writeBuffer(buf, ctx.FlexDirection, "flex-direction")
	writeBuffer(buf, ctx.FlexFlow, "flex-flow")
	writeBuffer(buf, ctx.FlexGrow, "flex-grow")
	writeBuffer(buf, ctx.FlexShrink, "flex-shrink")
	writeBuffer(buf, ctx.FlexWrap, "flex-wrap")
	writeBuffer(buf, ctx.Float, "float")
	writeBuffer(buf, ctx.Font, "font")
	writeBuffer(buf, ctx.FontStyle, "font-style")
	writeBuffer(buf, ctx.FontWeight, "font-weight")
	writeBuffer(buf, ctx.FontSize, "font-size")
	writeBuffer(buf, ctx.FontFamily, "font-family")
	writeBuffer(buf, ctx.FontVariantLigatures, "font-variant-ligatures")
	writeBuffer(buf, ctx.FontVariantCaps, "font-variant-caps")
	writeBuffer(buf, ctx.FontVariantNumeric, "font-variant-numeric")
	writeBuffer(buf, ctx.FontVariantEastAsian, "font-variant-east-asian")
	writeBuffer(buf, ctx.FontVariantAlternates, "font-variant-alternates")
	writeBuffer(buf, ctx.FontVariantPosition, "font-variant-position")
	writeBuffer(buf, ctx.FontStretch, "font-stretch")
	writeBuffer(buf, ctx.FontOpticalSizing, "font-optical-sizing")
	writeBuffer(buf, ctx.FontKerning, "font-kerning")
	writeBuffer(buf, ctx.FontFeatureSettings, "font-feature-settings")
	writeBuffer(buf, ctx.FontVariationSettings, "font-variation-settings")
	writeBuffer(buf, ctx.FontSizeAdjust, "font-size-adjust")
	writeBuffer(buf, ctx.HangingPunctuation, "hanging-punctuation")
	writeBuffer(buf, ctx.Height, "height")
	writeBuffer(buf, ctx.JustifyContent, "justify-content")
	writeBuffer(buf, ctx.LetterSpacing, "letter-spacing")
	writeBuffer(buf, ctx.Left, "left")
	writeBuffer(buf, ctx.ListStyle, "list-style")
	writeBuffer(buf, ctx.ListStyleImage, "list-style-image")
	writeBuffer(buf, ctx.ListStylePosition, "list-style-position")
	writeBuffer(buf, ctx.ListStyleType, "list-style-type")
	writeBuffer(buf, ctx.LineHeight, "line-height")
	writeBuffer(buf, ctx.Margin, "margin")
	writeBuffer(buf, ctx.MarginBottom, "margin-bottom")
	writeBuffer(buf, ctx.MarginTop, "margin-top")
	writeBuffer(buf, ctx.MarginLeft, "margin-left")
	writeBuffer(buf, ctx.MarginRight, "margin-right")
	writeBuffer(buf, ctx.MaxHeight, "max-height")
	writeBuffer(buf, ctx.MaxWidth, "max-width")
	writeBuffer(buf, ctx.MinHeight, "min-height")
	writeBuffer(buf, ctx.MinWidth, "min-width")
	writeBuffer(buf, ctx.NavDown, "nav-down")
	writeBuffer(buf, ctx.NavIndex, "nav-index")
	writeBuffer(buf, ctx.NavLeft, "nav-left")
	writeBuffer(buf, ctx.NavRight, "nav-right")
	writeBuffer(buf, ctx.NavUp, "nav-up")
	writeBuffer(buf, ctx.Opacity, "opacity")
	writeBuffer(buf, ctx.Order, "order")
	writeBuffer(buf, ctx.Outline, "outline")
	writeBuffer(buf, ctx.OutlineColor, "outline-color")
	writeBuffer(buf, ctx.OutlineOffset, "outline-offset")
	writeBuffer(buf, ctx.OutlineStyle, "outline-style")
	writeBuffer(buf, ctx.OutlineWidth, "outline-width")
	writeBuffer(buf, ctx.Overflow, "overflow")
	writeBuffer(buf, ctx.OverflowX, "overflow-x")
	writeBuffer(buf, ctx.OverflowY, "overflow-y")
	writeBuffer(buf, ctx.Padding, "padding")
	writeBuffer(buf, ctx.PaddingBlock, "padding-block")
	writeBuffer(buf, ctx.PaddingInline, "padding-inline")
	writeBuffer(buf, ctx.PaddingBottom, "padding-bottom")
	writeBuffer(buf, ctx.PaddingLeft, "padding-left")
	writeBuffer(buf, ctx.PaddingRight, "padding-right")
	writeBuffer(buf, ctx.PaddingTop, "padding-top")
	writeBuffer(buf, ctx.PageBreakAfter, "page-break-after")
	writeBuffer(buf, ctx.PageBreakBefore, "page-break-before")
	writeBuffer(buf, ctx.PageBreakInside, "page-break-inside")
	writeBuffer(buf, ctx.Perspective, "perspective")
	writeBuffer(buf, ctx.PerspectiveOrigin, "perspective-origin")
	writeBuffer(buf, ctx.Position, "position")
	writeBuffer(buf, ctx.Quotes, "quotes")
	writeBuffer(buf, ctx.Resize, "resize")
	writeBuffer(buf, ctx.Right, "right")
	writeBuffer(buf, ctx.TextTransform, "text-transform")
	writeBuffer(buf, ctx.TextIndent, "text-indent")
	writeBuffer(buf, ctx.TextShadow, "text-shadow")
	writeBuffer(buf, ctx.TextAlign, "text-align")
	writeBuffer(buf, ctx.TextRendering, "text-rendering")
	writeBuffer(buf, ctx.TabSize, "tab-size")
	writeBuffer(buf, ctx.TableLayout, "table-layout")
	writeBuffer(buf, ctx.TextAlignLast, "text-align-last")
	writeBuffer(buf, ctx.TextDecoration, "text-decoration")
	writeBuffer(buf, ctx.TextOverflow, "text-overflow")
	writeBuffer(buf, ctx.Top, "top")
	writeBuffer(buf, ctx.Transform, "transform")
	writeBuffer(buf, ctx.TransformOrigin, "transform-origin")
	writeBuffer(buf, ctx.TransformStyle, "transform-style")
	writeBuffer(buf, ctx.Transition, "transition")
	writeBuffer(buf, ctx.TransitionDelay, "transition-delay")
	writeBuffer(buf, ctx.TransitionDuration, "transition-duration")
	writeBuffer(buf, ctx.TransitionProperty, "transition-property")
	writeBuffer(buf, ctx.TransitionTimingFunction, "transition-timing-function")
	writeBuffer(buf, ctx.UnicodeBidi, "unicode-bidi")
	writeBuffer(buf, ctx.UserSelect, "user-select")
	writeBuffer(buf, ctx.VerticalAlign, "vertical-align")
	writeBuffer(buf, ctx.Visibility, "visibility")
	writeBuffer(buf, ctx.WhiteSpace, "white-space")
	writeBuffer(buf, ctx.WordBreak, "word-break")
	writeBuffer(buf, ctx.WordWrap, "word-wrap")
	writeBuffer(buf, ctx.WordSpacing, "word-spacing")
	writeBuffer(buf, ctx.Width, "width")
	writeBuffer(buf, ctx.ZIndex, "z-index")

	return buf.String()
}

func writeBuffer(buf *bytes.Buffer, value, key string) {
	if len(value) > 0 {
		buf.WriteString(fmt.Sprintf("%s:%s;", key, value))
	}
}
