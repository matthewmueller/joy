package htmliframeelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/document"
	"github.com/matthewmueller/golly/internal/gendom/dom/domsettabletokenlist"
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/getsvgdocument"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/window"
	"github.com/matthewmueller/golly/js"
)

type HTMLIFrameElement struct {
	*htmlelement.HTMLElement
	*getsvgdocument.GetSVGDocument
}

func (*HTMLIFrameElement) GetAlign() (align string) {
	js.Rewrite("$<.align")
	return align
}

func (*HTMLIFrameElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

func (*HTMLIFrameElement) GetAllowFullscreen() (allowFullscreen bool) {
	js.Rewrite("$<.allowFullscreen")
	return allowFullscreen
}

func (*HTMLIFrameElement) SetAllowFullscreen(allowFullscreen bool) {
	js.Rewrite("$<.allowFullscreen = $1", allowFullscreen)
}

func (*HTMLIFrameElement) GetAllowPaymentRequest() (allowPaymentRequest bool) {
	js.Rewrite("$<.allowPaymentRequest")
	return allowPaymentRequest
}

func (*HTMLIFrameElement) SetAllowPaymentRequest(allowPaymentRequest bool) {
	js.Rewrite("$<.allowPaymentRequest = $1", allowPaymentRequest)
}

func (*HTMLIFrameElement) GetBorder() (border string) {
	js.Rewrite("$<.border")
	return border
}

func (*HTMLIFrameElement) SetBorder(border string) {
	js.Rewrite("$<.border = $1", border)
}

func (*HTMLIFrameElement) GetContentDocument() (contentDocument *document.Document) {
	js.Rewrite("$<.contentDocument")
	return contentDocument
}

func (*HTMLIFrameElement) GetContentWindow() (contentWindow *window.Window) {
	js.Rewrite("$<.contentWindow")
	return contentWindow
}

func (*HTMLIFrameElement) GetFrameBorder() (frameBorder string) {
	js.Rewrite("$<.frameBorder")
	return frameBorder
}

func (*HTMLIFrameElement) SetFrameBorder(frameBorder string) {
	js.Rewrite("$<.frameBorder = $1", frameBorder)
}

func (*HTMLIFrameElement) GetFrameSpacing() (frameSpacing interface{}) {
	js.Rewrite("$<.frameSpacing")
	return frameSpacing
}

func (*HTMLIFrameElement) SetFrameSpacing(frameSpacing interface{}) {
	js.Rewrite("$<.frameSpacing = $1", frameSpacing)
}

func (*HTMLIFrameElement) GetHeight() (height string) {
	js.Rewrite("$<.height")
	return height
}

func (*HTMLIFrameElement) SetHeight(height string) {
	js.Rewrite("$<.height = $1", height)
}

func (*HTMLIFrameElement) GetHspace() (hspace int) {
	js.Rewrite("$<.hspace")
	return hspace
}

func (*HTMLIFrameElement) SetHspace(hspace int) {
	js.Rewrite("$<.hspace = $1", hspace)
}

func (*HTMLIFrameElement) GetLongDesc() (longDesc string) {
	js.Rewrite("$<.longDesc")
	return longDesc
}

func (*HTMLIFrameElement) SetLongDesc(longDesc string) {
	js.Rewrite("$<.longDesc = $1", longDesc)
}

func (*HTMLIFrameElement) GetMarginHeight() (marginHeight string) {
	js.Rewrite("$<.marginHeight")
	return marginHeight
}

func (*HTMLIFrameElement) SetMarginHeight(marginHeight string) {
	js.Rewrite("$<.marginHeight = $1", marginHeight)
}

func (*HTMLIFrameElement) GetMarginWidth() (marginWidth string) {
	js.Rewrite("$<.marginWidth")
	return marginWidth
}

func (*HTMLIFrameElement) SetMarginWidth(marginWidth string) {
	js.Rewrite("$<.marginWidth = $1", marginWidth)
}

func (*HTMLIFrameElement) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*HTMLIFrameElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

func (*HTMLIFrameElement) GetNoResize() (noResize bool) {
	js.Rewrite("$<.noResize")
	return noResize
}

func (*HTMLIFrameElement) SetNoResize(noResize bool) {
	js.Rewrite("$<.noResize = $1", noResize)
}

func (*HTMLIFrameElement) GetOnload() (load *event.Event) {
	js.Rewrite("$<.onload")
	return load
}

func (*HTMLIFrameElement) SetOnload(load *event.Event) {
	js.Rewrite("$<.onload = $1", load)
}

func (*HTMLIFrameElement) GetSandbox() (sandbox *domsettabletokenlist.DOMSettableTokenList) {
	js.Rewrite("$<.sandbox")
	return sandbox
}

func (*HTMLIFrameElement) GetScrolling() (scrolling string) {
	js.Rewrite("$<.scrolling")
	return scrolling
}

func (*HTMLIFrameElement) SetScrolling(scrolling string) {
	js.Rewrite("$<.scrolling = $1", scrolling)
}

func (*HTMLIFrameElement) GetSrc() (src string) {
	js.Rewrite("$<.src")
	return src
}

func (*HTMLIFrameElement) SetSrc(src string) {
	js.Rewrite("$<.src = $1", src)
}

func (*HTMLIFrameElement) GetVspace() (vspace int) {
	js.Rewrite("$<.vspace")
	return vspace
}

func (*HTMLIFrameElement) SetVspace(vspace int) {
	js.Rewrite("$<.vspace = $1", vspace)
}

func (*HTMLIFrameElement) GetWidth() (width string) {
	js.Rewrite("$<.width")
	return width
}

func (*HTMLIFrameElement) SetWidth(width string) {
	js.Rewrite("$<.width = $1", width)
}
