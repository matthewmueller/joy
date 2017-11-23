package htmlbodyelement

import (
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLBodyElement struct
// js:"HTMLBodyElement,omit"
type HTMLBodyElement struct {
	window.HTMLElement
}

// ALink prop
func (*HTMLBodyElement) ALink() (aLink interface{}) {
	js.Rewrite("$<.aLink")
	return aLink
}

// ALink prop
func (*HTMLBodyElement) SetALink(aLink interface{}) {
	js.Rewrite("$<.aLink = $1", aLink)
}

// Background prop
func (*HTMLBodyElement) Background() (background string) {
	js.Rewrite("$<.background")
	return background
}

// Background prop
func (*HTMLBodyElement) SetBackground(background string) {
	js.Rewrite("$<.background = $1", background)
}

// BgColor prop
func (*HTMLBodyElement) BgColor() (bgColor interface{}) {
	js.Rewrite("$<.bgColor")
	return bgColor
}

// BgColor prop
func (*HTMLBodyElement) SetBgColor(bgColor interface{}) {
	js.Rewrite("$<.bgColor = $1", bgColor)
}

// BgProperties prop
func (*HTMLBodyElement) BgProperties() (bgProperties string) {
	js.Rewrite("$<.bgProperties")
	return bgProperties
}

// BgProperties prop
func (*HTMLBodyElement) SetBgProperties(bgProperties string) {
	js.Rewrite("$<.bgProperties = $1", bgProperties)
}

// Link prop
func (*HTMLBodyElement) Link() (link interface{}) {
	js.Rewrite("$<.link")
	return link
}

// Link prop
func (*HTMLBodyElement) SetLink(link interface{}) {
	js.Rewrite("$<.link = $1", link)
}

// NoWrap prop
func (*HTMLBodyElement) NoWrap() (noWrap bool) {
	js.Rewrite("$<.noWrap")
	return noWrap
}

// NoWrap prop
func (*HTMLBodyElement) SetNoWrap(noWrap bool) {
	js.Rewrite("$<.noWrap = $1", noWrap)
}

// Onafterprint prop
func (*HTMLBodyElement) Onafterprint() (onafterprint func(window.Event)) {
	js.Rewrite("$<.onafterprint")
	return onafterprint
}

// Onafterprint prop
func (*HTMLBodyElement) SetOnafterprint(onafterprint func(window.Event)) {
	js.Rewrite("$<.onafterprint = $1", onafterprint)
}

// Onbeforeprint prop
func (*HTMLBodyElement) Onbeforeprint() (onbeforeprint func(window.Event)) {
	js.Rewrite("$<.onbeforeprint")
	return onbeforeprint
}

// Onbeforeprint prop
func (*HTMLBodyElement) SetOnbeforeprint(onbeforeprint func(window.Event)) {
	js.Rewrite("$<.onbeforeprint = $1", onbeforeprint)
}

// Onbeforeunload prop
func (*HTMLBodyElement) Onbeforeunload() (onbeforeunload func(window.Event)) {
	js.Rewrite("$<.onbeforeunload")
	return onbeforeunload
}

// Onbeforeunload prop
func (*HTMLBodyElement) SetOnbeforeunload(onbeforeunload func(window.Event)) {
	js.Rewrite("$<.onbeforeunload = $1", onbeforeunload)
}

// Onblur prop
func (*HTMLBodyElement) Onblur() (onblur func(*window.FocusEvent)) {
	js.Rewrite("$<.onblur")
	return onblur
}

// Onblur prop
func (*HTMLBodyElement) SetOnblur(onblur func(*window.FocusEvent)) {
	js.Rewrite("$<.onblur = $1", onblur)
}

// Onerror prop
func (*HTMLBodyElement) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$<.onerror")
	return onerror
}

// Onerror prop
func (*HTMLBodyElement) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$<.onerror = $1", onerror)
}

// Onfocus prop
func (*HTMLBodyElement) Onfocus() (onfocus func(*window.FocusEvent)) {
	js.Rewrite("$<.onfocus")
	return onfocus
}

// Onfocus prop
func (*HTMLBodyElement) SetOnfocus(onfocus func(*window.FocusEvent)) {
	js.Rewrite("$<.onfocus = $1", onfocus)
}

// Onhashchange prop
func (*HTMLBodyElement) Onhashchange() (onhashchange func(window.Event)) {
	js.Rewrite("$<.onhashchange")
	return onhashchange
}

// Onhashchange prop
func (*HTMLBodyElement) SetOnhashchange(onhashchange func(window.Event)) {
	js.Rewrite("$<.onhashchange = $1", onhashchange)
}

// Onload prop
func (*HTMLBodyElement) Onload() (onload func(window.Event)) {
	js.Rewrite("$<.onload")
	return onload
}

// Onload prop
func (*HTMLBodyElement) SetOnload(onload func(window.Event)) {
	js.Rewrite("$<.onload = $1", onload)
}

// Onmessage prop
func (*HTMLBodyElement) Onmessage() (onmessage func(window.Event)) {
	js.Rewrite("$<.onmessage")
	return onmessage
}

// Onmessage prop
func (*HTMLBodyElement) SetOnmessage(onmessage func(window.Event)) {
	js.Rewrite("$<.onmessage = $1", onmessage)
}

// Onoffline prop
func (*HTMLBodyElement) Onoffline() (onoffline func(window.Event)) {
	js.Rewrite("$<.onoffline")
	return onoffline
}

// Onoffline prop
func (*HTMLBodyElement) SetOnoffline(onoffline func(window.Event)) {
	js.Rewrite("$<.onoffline = $1", onoffline)
}

// Ononline prop
func (*HTMLBodyElement) Ononline() (ononline func(window.Event)) {
	js.Rewrite("$<.ononline")
	return ononline
}

// Ononline prop
func (*HTMLBodyElement) SetOnonline(ononline func(window.Event)) {
	js.Rewrite("$<.ononline = $1", ononline)
}

// Onorientationchange prop
func (*HTMLBodyElement) Onorientationchange() (onorientationchange func(window.Event)) {
	js.Rewrite("$<.onorientationchange")
	return onorientationchange
}

// Onorientationchange prop
func (*HTMLBodyElement) SetOnorientationchange(onorientationchange func(window.Event)) {
	js.Rewrite("$<.onorientationchange = $1", onorientationchange)
}

// Onpagehide prop
func (*HTMLBodyElement) Onpagehide() (onpagehide func(window.Event)) {
	js.Rewrite("$<.onpagehide")
	return onpagehide
}

// Onpagehide prop
func (*HTMLBodyElement) SetOnpagehide(onpagehide func(window.Event)) {
	js.Rewrite("$<.onpagehide = $1", onpagehide)
}

// Onpageshow prop
func (*HTMLBodyElement) Onpageshow() (onpageshow func(window.Event)) {
	js.Rewrite("$<.onpageshow")
	return onpageshow
}

// Onpageshow prop
func (*HTMLBodyElement) SetOnpageshow(onpageshow func(window.Event)) {
	js.Rewrite("$<.onpageshow = $1", onpageshow)
}

// Onpopstate prop
func (*HTMLBodyElement) Onpopstate() (onpopstate func(window.Event)) {
	js.Rewrite("$<.onpopstate")
	return onpopstate
}

// Onpopstate prop
func (*HTMLBodyElement) SetOnpopstate(onpopstate func(window.Event)) {
	js.Rewrite("$<.onpopstate = $1", onpopstate)
}

// Onresize prop
func (*HTMLBodyElement) Onresize() (onresize func(window.Event)) {
	js.Rewrite("$<.onresize")
	return onresize
}

// Onresize prop
func (*HTMLBodyElement) SetOnresize(onresize func(window.Event)) {
	js.Rewrite("$<.onresize = $1", onresize)
}

// Onscroll prop
func (*HTMLBodyElement) Onscroll() (onscroll func(window.UIEvent)) {
	js.Rewrite("$<.onscroll")
	return onscroll
}

// Onscroll prop
func (*HTMLBodyElement) SetOnscroll(onscroll func(window.UIEvent)) {
	js.Rewrite("$<.onscroll = $1", onscroll)
}

// Onstorage prop
func (*HTMLBodyElement) Onstorage() (onstorage func(window.Event)) {
	js.Rewrite("$<.onstorage")
	return onstorage
}

// Onstorage prop
func (*HTMLBodyElement) SetOnstorage(onstorage func(window.Event)) {
	js.Rewrite("$<.onstorage = $1", onstorage)
}

// Onunload prop
func (*HTMLBodyElement) Onunload() (onunload func(window.Event)) {
	js.Rewrite("$<.onunload")
	return onunload
}

// Onunload prop
func (*HTMLBodyElement) SetOnunload(onunload func(window.Event)) {
	js.Rewrite("$<.onunload = $1", onunload)
}

// Text prop
func (*HTMLBodyElement) Text() (text interface{}) {
	js.Rewrite("$<.text")
	return text
}

// Text prop
func (*HTMLBodyElement) SetText(text interface{}) {
	js.Rewrite("$<.text = $1", text)
}

// VLink prop
func (*HTMLBodyElement) VLink() (vLink interface{}) {
	js.Rewrite("$<.vLink")
	return vLink
}

// VLink prop
func (*HTMLBodyElement) SetVLink(vLink interface{}) {
	js.Rewrite("$<.vLink = $1", vLink)
}
