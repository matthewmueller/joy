package htmlframesetelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/focusevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/uievent"
	"github.com/matthewmueller/golly/js"
)

type HTMLFrameSetElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLFrameSetElement) GetBorder() (border string) {
	js.Rewrite("$<.border")
	return border
}

func (*HTMLFrameSetElement) SetBorder(border string) {
	js.Rewrite("$<.border = $1", border)
}

func (*HTMLFrameSetElement) GetBorderColor() (borderColor interface{}) {
	js.Rewrite("$<.borderColor")
	return borderColor
}

func (*HTMLFrameSetElement) SetBorderColor(borderColor interface{}) {
	js.Rewrite("$<.borderColor = $1", borderColor)
}

func (*HTMLFrameSetElement) GetCols() (cols string) {
	js.Rewrite("$<.cols")
	return cols
}

func (*HTMLFrameSetElement) SetCols(cols string) {
	js.Rewrite("$<.cols = $1", cols)
}

func (*HTMLFrameSetElement) GetFrameBorder() (frameBorder string) {
	js.Rewrite("$<.frameBorder")
	return frameBorder
}

func (*HTMLFrameSetElement) SetFrameBorder(frameBorder string) {
	js.Rewrite("$<.frameBorder = $1", frameBorder)
}

func (*HTMLFrameSetElement) GetFrameSpacing() (frameSpacing interface{}) {
	js.Rewrite("$<.frameSpacing")
	return frameSpacing
}

func (*HTMLFrameSetElement) SetFrameSpacing(frameSpacing interface{}) {
	js.Rewrite("$<.frameSpacing = $1", frameSpacing)
}

func (*HTMLFrameSetElement) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*HTMLFrameSetElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

func (*HTMLFrameSetElement) GetOnafterprint() (e *event.Event) {
	js.Rewrite("$<.onafterprint")
	return e
}

func (*HTMLFrameSetElement) SetOnafterprint(e *event.Event) {
	js.Rewrite("$<.onafterprint = $1", e)
}

func (*HTMLFrameSetElement) GetOnbeforeprint() (e *event.Event) {
	js.Rewrite("$<.onbeforeprint")
	return e
}

func (*HTMLFrameSetElement) SetOnbeforeprint(e *event.Event) {
	js.Rewrite("$<.onbeforeprint = $1", e)
}

func (*HTMLFrameSetElement) GetOnbeforeunload() (e *event.Event) {
	js.Rewrite("$<.onbeforeunload")
	return e
}

func (*HTMLFrameSetElement) SetOnbeforeunload(e *event.Event) {
	js.Rewrite("$<.onbeforeunload = $1", e)
}

func (*HTMLFrameSetElement) GetOnblur() (blur *focusevent.FocusEvent) {
	js.Rewrite("$<.onblur")
	return blur
}

func (*HTMLFrameSetElement) SetOnblur(blur *focusevent.FocusEvent) {
	js.Rewrite("$<.onblur = $1", blur)
}

func (*HTMLFrameSetElement) GetOnerror() (e *event.Event) {
	js.Rewrite("$<.onerror")
	return e
}

func (*HTMLFrameSetElement) SetOnerror(e *event.Event) {
	js.Rewrite("$<.onerror = $1", e)
}

func (*HTMLFrameSetElement) GetOnfocus() (focus *focusevent.FocusEvent) {
	js.Rewrite("$<.onfocus")
	return focus
}

func (*HTMLFrameSetElement) SetOnfocus(focus *focusevent.FocusEvent) {
	js.Rewrite("$<.onfocus = $1", focus)
}

func (*HTMLFrameSetElement) GetOnhashchange() (e *event.Event) {
	js.Rewrite("$<.onhashchange")
	return e
}

func (*HTMLFrameSetElement) SetOnhashchange(e *event.Event) {
	js.Rewrite("$<.onhashchange = $1", e)
}

func (*HTMLFrameSetElement) GetOnload() (e *event.Event) {
	js.Rewrite("$<.onload")
	return e
}

func (*HTMLFrameSetElement) SetOnload(e *event.Event) {
	js.Rewrite("$<.onload = $1", e)
}

func (*HTMLFrameSetElement) GetOnmessage() (e *event.Event) {
	js.Rewrite("$<.onmessage")
	return e
}

func (*HTMLFrameSetElement) SetOnmessage(e *event.Event) {
	js.Rewrite("$<.onmessage = $1", e)
}

func (*HTMLFrameSetElement) GetOnoffline() (e *event.Event) {
	js.Rewrite("$<.onoffline")
	return e
}

func (*HTMLFrameSetElement) SetOnoffline(e *event.Event) {
	js.Rewrite("$<.onoffline = $1", e)
}

func (*HTMLFrameSetElement) GetOnonline() (e *event.Event) {
	js.Rewrite("$<.ononline")
	return e
}

func (*HTMLFrameSetElement) SetOnonline(e *event.Event) {
	js.Rewrite("$<.ononline = $1", e)
}

func (*HTMLFrameSetElement) GetOnorientationchange() (e *event.Event) {
	js.Rewrite("$<.onorientationchange")
	return e
}

func (*HTMLFrameSetElement) SetOnorientationchange(e *event.Event) {
	js.Rewrite("$<.onorientationchange = $1", e)
}

func (*HTMLFrameSetElement) GetOnpagehide() (e *event.Event) {
	js.Rewrite("$<.onpagehide")
	return e
}

func (*HTMLFrameSetElement) SetOnpagehide(e *event.Event) {
	js.Rewrite("$<.onpagehide = $1", e)
}

func (*HTMLFrameSetElement) GetOnpageshow() (e *event.Event) {
	js.Rewrite("$<.onpageshow")
	return e
}

func (*HTMLFrameSetElement) SetOnpageshow(e *event.Event) {
	js.Rewrite("$<.onpageshow = $1", e)
}

func (*HTMLFrameSetElement) GetOnpopstate() (e *event.Event) {
	js.Rewrite("$<.onpopstate")
	return e
}

func (*HTMLFrameSetElement) SetOnpopstate(e *event.Event) {
	js.Rewrite("$<.onpopstate = $1", e)
}

func (*HTMLFrameSetElement) GetOnresize() (e *event.Event) {
	js.Rewrite("$<.onresize")
	return e
}

func (*HTMLFrameSetElement) SetOnresize(e *event.Event) {
	js.Rewrite("$<.onresize = $1", e)
}

func (*HTMLFrameSetElement) GetOnscroll() (scroll *uievent.UIEvent) {
	js.Rewrite("$<.onscroll")
	return scroll
}

func (*HTMLFrameSetElement) SetOnscroll(scroll *uievent.UIEvent) {
	js.Rewrite("$<.onscroll = $1", scroll)
}

func (*HTMLFrameSetElement) GetOnstorage() (e *event.Event) {
	js.Rewrite("$<.onstorage")
	return e
}

func (*HTMLFrameSetElement) SetOnstorage(e *event.Event) {
	js.Rewrite("$<.onstorage = $1", e)
}

func (*HTMLFrameSetElement) GetOnunload() (e *event.Event) {
	js.Rewrite("$<.onunload")
	return e
}

func (*HTMLFrameSetElement) SetOnunload(e *event.Event) {
	js.Rewrite("$<.onunload = $1", e)
}

func (*HTMLFrameSetElement) GetRows() (rows string) {
	js.Rewrite("$<.rows")
	return rows
}

func (*HTMLFrameSetElement) SetRows(rows string) {
	js.Rewrite("$<.rows = $1", rows)
}
