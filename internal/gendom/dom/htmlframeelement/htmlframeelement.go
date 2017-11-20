package htmlframeelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/document"
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/getsvgdocument"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/window"
	"github.com/matthewmueller/golly/js"
)

type HTMLFrameElement struct {
	*htmlelement.HTMLElement
	*getsvgdocument.GetSVGDocument
}

func (*HTMLFrameElement) GetBorder() (border string) {
	js.Rewrite("$<.border")
	return border
}

func (*HTMLFrameElement) SetBorder(border string) {
	js.Rewrite("$<.border = $1", border)
}

func (*HTMLFrameElement) GetBorderColor() (borderColor interface{}) {
	js.Rewrite("$<.borderColor")
	return borderColor
}

func (*HTMLFrameElement) SetBorderColor(borderColor interface{}) {
	js.Rewrite("$<.borderColor = $1", borderColor)
}

func (*HTMLFrameElement) GetContentDocument() (contentDocument *document.Document) {
	js.Rewrite("$<.contentDocument")
	return contentDocument
}

func (*HTMLFrameElement) GetContentWindow() (contentWindow *window.Window) {
	js.Rewrite("$<.contentWindow")
	return contentWindow
}

func (*HTMLFrameElement) GetFrameBorder() (frameBorder string) {
	js.Rewrite("$<.frameBorder")
	return frameBorder
}

func (*HTMLFrameElement) SetFrameBorder(frameBorder string) {
	js.Rewrite("$<.frameBorder = $1", frameBorder)
}

func (*HTMLFrameElement) GetFrameSpacing() (frameSpacing interface{}) {
	js.Rewrite("$<.frameSpacing")
	return frameSpacing
}

func (*HTMLFrameElement) SetFrameSpacing(frameSpacing interface{}) {
	js.Rewrite("$<.frameSpacing = $1", frameSpacing)
}

func (*HTMLFrameElement) GetHeight() (height interface{}) {
	js.Rewrite("$<.height")
	return height
}

func (*HTMLFrameElement) SetHeight(height interface{}) {
	js.Rewrite("$<.height = $1", height)
}

func (*HTMLFrameElement) GetLongDesc() (longDesc string) {
	js.Rewrite("$<.longDesc")
	return longDesc
}

func (*HTMLFrameElement) SetLongDesc(longDesc string) {
	js.Rewrite("$<.longDesc = $1", longDesc)
}

func (*HTMLFrameElement) GetMarginHeight() (marginHeight string) {
	js.Rewrite("$<.marginHeight")
	return marginHeight
}

func (*HTMLFrameElement) SetMarginHeight(marginHeight string) {
	js.Rewrite("$<.marginHeight = $1", marginHeight)
}

func (*HTMLFrameElement) GetMarginWidth() (marginWidth string) {
	js.Rewrite("$<.marginWidth")
	return marginWidth
}

func (*HTMLFrameElement) SetMarginWidth(marginWidth string) {
	js.Rewrite("$<.marginWidth = $1", marginWidth)
}

func (*HTMLFrameElement) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*HTMLFrameElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

func (*HTMLFrameElement) GetNoResize() (noResize bool) {
	js.Rewrite("$<.noResize")
	return noResize
}

func (*HTMLFrameElement) SetNoResize(noResize bool) {
	js.Rewrite("$<.noResize = $1", noResize)
}

func (*HTMLFrameElement) GetOnload() (load *event.Event) {
	js.Rewrite("$<.onload")
	return load
}

func (*HTMLFrameElement) SetOnload(load *event.Event) {
	js.Rewrite("$<.onload = $1", load)
}

func (*HTMLFrameElement) GetScrolling() (scrolling string) {
	js.Rewrite("$<.scrolling")
	return scrolling
}

func (*HTMLFrameElement) SetScrolling(scrolling string) {
	js.Rewrite("$<.scrolling = $1", scrolling)
}

func (*HTMLFrameElement) GetSrc() (src string) {
	js.Rewrite("$<.src")
	return src
}

func (*HTMLFrameElement) SetSrc(src string) {
	js.Rewrite("$<.src = $1", src)
}

func (*HTMLFrameElement) GetWidth() (width interface{}) {
	js.Rewrite("$<.width")
	return width
}

func (*HTMLFrameElement) SetWidth(width interface{}) {
	js.Rewrite("$<.width = $1", width)
}
