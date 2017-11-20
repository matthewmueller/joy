package htmlbodyelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/focusevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/uievent"
	"github.com/matthewmueller/golly/js"
)

type HTMLBodyElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLBodyElement) GetALink() (aLink interface{}) {
	js.Rewrite("$<.aLink")
	return aLink
}

func (*HTMLBodyElement) SetALink(aLink interface{}) {
	js.Rewrite("$<.aLink = $1", aLink)
}

func (*HTMLBodyElement) GetBackground() (background string) {
	js.Rewrite("$<.background")
	return background
}

func (*HTMLBodyElement) SetBackground(background string) {
	js.Rewrite("$<.background = $1", background)
}

func (*HTMLBodyElement) GetBgColor() (bgColor interface{}) {
	js.Rewrite("$<.bgColor")
	return bgColor
}

func (*HTMLBodyElement) SetBgColor(bgColor interface{}) {
	js.Rewrite("$<.bgColor = $1", bgColor)
}

func (*HTMLBodyElement) GetBgProperties() (bgProperties string) {
	js.Rewrite("$<.bgProperties")
	return bgProperties
}

func (*HTMLBodyElement) SetBgProperties(bgProperties string) {
	js.Rewrite("$<.bgProperties = $1", bgProperties)
}

func (*HTMLBodyElement) GetLink() (link interface{}) {
	js.Rewrite("$<.link")
	return link
}

func (*HTMLBodyElement) SetLink(link interface{}) {
	js.Rewrite("$<.link = $1", link)
}

func (*HTMLBodyElement) GetNoWrap() (noWrap bool) {
	js.Rewrite("$<.noWrap")
	return noWrap
}

func (*HTMLBodyElement) SetNoWrap(noWrap bool) {
	js.Rewrite("$<.noWrap = $1", noWrap)
}

func (*HTMLBodyElement) GetOnafterprint() (e *event.Event) {
	js.Rewrite("$<.onafterprint")
	return e
}

func (*HTMLBodyElement) SetOnafterprint(e *event.Event) {
	js.Rewrite("$<.onafterprint = $1", e)
}

func (*HTMLBodyElement) GetOnbeforeprint() (e *event.Event) {
	js.Rewrite("$<.onbeforeprint")
	return e
}

func (*HTMLBodyElement) SetOnbeforeprint(e *event.Event) {
	js.Rewrite("$<.onbeforeprint = $1", e)
}

func (*HTMLBodyElement) GetOnbeforeunload() (e *event.Event) {
	js.Rewrite("$<.onbeforeunload")
	return e
}

func (*HTMLBodyElement) SetOnbeforeunload(e *event.Event) {
	js.Rewrite("$<.onbeforeunload = $1", e)
}

func (*HTMLBodyElement) GetOnblur() (blur *focusevent.FocusEvent) {
	js.Rewrite("$<.onblur")
	return blur
}

func (*HTMLBodyElement) SetOnblur(blur *focusevent.FocusEvent) {
	js.Rewrite("$<.onblur = $1", blur)
}

func (*HTMLBodyElement) GetOnerror() (e *event.Event) {
	js.Rewrite("$<.onerror")
	return e
}

func (*HTMLBodyElement) SetOnerror(e *event.Event) {
	js.Rewrite("$<.onerror = $1", e)
}

func (*HTMLBodyElement) GetOnfocus() (focus *focusevent.FocusEvent) {
	js.Rewrite("$<.onfocus")
	return focus
}

func (*HTMLBodyElement) SetOnfocus(focus *focusevent.FocusEvent) {
	js.Rewrite("$<.onfocus = $1", focus)
}

func (*HTMLBodyElement) GetOnhashchange() (e *event.Event) {
	js.Rewrite("$<.onhashchange")
	return e
}

func (*HTMLBodyElement) SetOnhashchange(e *event.Event) {
	js.Rewrite("$<.onhashchange = $1", e)
}

func (*HTMLBodyElement) GetOnload() (e *event.Event) {
	js.Rewrite("$<.onload")
	return e
}

func (*HTMLBodyElement) SetOnload(e *event.Event) {
	js.Rewrite("$<.onload = $1", e)
}

func (*HTMLBodyElement) GetOnmessage() (e *event.Event) {
	js.Rewrite("$<.onmessage")
	return e
}

func (*HTMLBodyElement) SetOnmessage(e *event.Event) {
	js.Rewrite("$<.onmessage = $1", e)
}

func (*HTMLBodyElement) GetOnoffline() (offline *event.Event) {
	js.Rewrite("$<.onoffline")
	return offline
}

func (*HTMLBodyElement) SetOnoffline(offline *event.Event) {
	js.Rewrite("$<.onoffline = $1", offline)
}

func (*HTMLBodyElement) GetOnonline() (online *event.Event) {
	js.Rewrite("$<.ononline")
	return online
}

func (*HTMLBodyElement) SetOnonline(online *event.Event) {
	js.Rewrite("$<.ononline = $1", online)
}

func (*HTMLBodyElement) GetOnorientationchange() (e *event.Event) {
	js.Rewrite("$<.onorientationchange")
	return e
}

func (*HTMLBodyElement) SetOnorientationchange(e *event.Event) {
	js.Rewrite("$<.onorientationchange = $1", e)
}

func (*HTMLBodyElement) GetOnpagehide() (e *event.Event) {
	js.Rewrite("$<.onpagehide")
	return e
}

func (*HTMLBodyElement) SetOnpagehide(e *event.Event) {
	js.Rewrite("$<.onpagehide = $1", e)
}

func (*HTMLBodyElement) GetOnpageshow() (e *event.Event) {
	js.Rewrite("$<.onpageshow")
	return e
}

func (*HTMLBodyElement) SetOnpageshow(e *event.Event) {
	js.Rewrite("$<.onpageshow = $1", e)
}

func (*HTMLBodyElement) GetOnpopstate() (e *event.Event) {
	js.Rewrite("$<.onpopstate")
	return e
}

func (*HTMLBodyElement) SetOnpopstate(e *event.Event) {
	js.Rewrite("$<.onpopstate = $1", e)
}

func (*HTMLBodyElement) GetOnresize() (e *event.Event) {
	js.Rewrite("$<.onresize")
	return e
}

func (*HTMLBodyElement) SetOnresize(e *event.Event) {
	js.Rewrite("$<.onresize = $1", e)
}

func (*HTMLBodyElement) GetOnscroll() (scroll *uievent.UIEvent) {
	js.Rewrite("$<.onscroll")
	return scroll
}

func (*HTMLBodyElement) SetOnscroll(scroll *uievent.UIEvent) {
	js.Rewrite("$<.onscroll = $1", scroll)
}

func (*HTMLBodyElement) GetOnstorage() (e *event.Event) {
	js.Rewrite("$<.onstorage")
	return e
}

func (*HTMLBodyElement) SetOnstorage(e *event.Event) {
	js.Rewrite("$<.onstorage = $1", e)
}

func (*HTMLBodyElement) GetOnunload() (e *event.Event) {
	js.Rewrite("$<.onunload")
	return e
}

func (*HTMLBodyElement) SetOnunload(e *event.Event) {
	js.Rewrite("$<.onunload = $1", e)
}

func (*HTMLBodyElement) GetText() (text interface{}) {
	js.Rewrite("$<.text")
	return text
}

func (*HTMLBodyElement) SetText(text interface{}) {
	js.Rewrite("$<.text = $1", text)
}

func (*HTMLBodyElement) GetVLink() (vLink interface{}) {
	js.Rewrite("$<.vLink")
	return vLink
}

func (*HTMLBodyElement) SetVLink(vLink interface{}) {
	js.Rewrite("$<.vLink = $1", vLink)
}
