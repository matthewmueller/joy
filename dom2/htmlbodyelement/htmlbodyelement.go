package htmlbodyelement

import "github.com/matthewmueller/golly/js"

// js:"HTMLBodyElement,omit"
type HTMLBodyElement struct {
	window.HTMLElement
}

// ALink
func (*HTMLBodyElement) ALink() (aLink interface{}) {
	js.Rewrite("$<.ALink")
	return aLink
}

// ALink
func (*HTMLBodyElement) SetALink(aLink interface{}) {
	js.Rewrite("$<.ALink = $1", aLink)
}

// Background
func (*HTMLBodyElement) Background() (background string) {
	js.Rewrite("$<.Background")
	return background
}

// Background
func (*HTMLBodyElement) SetBackground(background string) {
	js.Rewrite("$<.Background = $1", background)
}

// BgColor
func (*HTMLBodyElement) BgColor() (bgColor interface{}) {
	js.Rewrite("$<.BgColor")
	return bgColor
}

// BgColor
func (*HTMLBodyElement) SetBgColor(bgColor interface{}) {
	js.Rewrite("$<.BgColor = $1", bgColor)
}

// BgProperties
func (*HTMLBodyElement) BgProperties() (bgProperties string) {
	js.Rewrite("$<.BgProperties")
	return bgProperties
}

// BgProperties
func (*HTMLBodyElement) SetBgProperties(bgProperties string) {
	js.Rewrite("$<.BgProperties = $1", bgProperties)
}

// Link
func (*HTMLBodyElement) Link() (link interface{}) {
	js.Rewrite("$<.Link")
	return link
}

// Link
func (*HTMLBodyElement) SetLink(link interface{}) {
	js.Rewrite("$<.Link = $1", link)
}

// NoWrap
func (*HTMLBodyElement) NoWrap() (noWrap bool) {
	js.Rewrite("$<.NoWrap")
	return noWrap
}

// NoWrap
func (*HTMLBodyElement) SetNoWrap(noWrap bool) {
	js.Rewrite("$<.NoWrap = $1", noWrap)
}

// Onafterprint
func (*HTMLBodyElement) Onafterprint() (onafterprint func(window.Event)) {
	js.Rewrite("$<.Onafterprint")
	return onafterprint
}

// Onafterprint
func (*HTMLBodyElement) SetOnafterprint(onafterprint func(window.Event)) {
	js.Rewrite("$<.Onafterprint = $1", onafterprint)
}

// Onbeforeprint
func (*HTMLBodyElement) Onbeforeprint() (onbeforeprint func(window.Event)) {
	js.Rewrite("$<.Onbeforeprint")
	return onbeforeprint
}

// Onbeforeprint
func (*HTMLBodyElement) SetOnbeforeprint(onbeforeprint func(window.Event)) {
	js.Rewrite("$<.Onbeforeprint = $1", onbeforeprint)
}

// Onbeforeunload
func (*HTMLBodyElement) Onbeforeunload() (onbeforeunload func(window.Event)) {
	js.Rewrite("$<.Onbeforeunload")
	return onbeforeunload
}

// Onbeforeunload
func (*HTMLBodyElement) SetOnbeforeunload(onbeforeunload func(window.Event)) {
	js.Rewrite("$<.Onbeforeunload = $1", onbeforeunload)
}

// Onblur
func (*HTMLBodyElement) Onblur() (onblur func(*window.FocusEvent)) {
	js.Rewrite("$<.Onblur")
	return onblur
}

// Onblur
func (*HTMLBodyElement) SetOnblur(onblur func(*window.FocusEvent)) {
	js.Rewrite("$<.Onblur = $1", onblur)
}

// Onerror
func (*HTMLBodyElement) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$<.Onerror")
	return onerror
}

// Onerror
func (*HTMLBodyElement) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$<.Onerror = $1", onerror)
}

// Onfocus
func (*HTMLBodyElement) Onfocus() (onfocus func(*window.FocusEvent)) {
	js.Rewrite("$<.Onfocus")
	return onfocus
}

// Onfocus
func (*HTMLBodyElement) SetOnfocus(onfocus func(*window.FocusEvent)) {
	js.Rewrite("$<.Onfocus = $1", onfocus)
}

// Onhashchange
func (*HTMLBodyElement) Onhashchange() (onhashchange func(window.Event)) {
	js.Rewrite("$<.Onhashchange")
	return onhashchange
}

// Onhashchange
func (*HTMLBodyElement) SetOnhashchange(onhashchange func(window.Event)) {
	js.Rewrite("$<.Onhashchange = $1", onhashchange)
}

// Onload
func (*HTMLBodyElement) Onload() (onload func(window.Event)) {
	js.Rewrite("$<.Onload")
	return onload
}

// Onload
func (*HTMLBodyElement) SetOnload(onload func(window.Event)) {
	js.Rewrite("$<.Onload = $1", onload)
}

// Onmessage
func (*HTMLBodyElement) Onmessage() (onmessage func(window.Event)) {
	js.Rewrite("$<.Onmessage")
	return onmessage
}

// Onmessage
func (*HTMLBodyElement) SetOnmessage(onmessage func(window.Event)) {
	js.Rewrite("$<.Onmessage = $1", onmessage)
}

// Onoffline
func (*HTMLBodyElement) Onoffline() (onoffline func(window.Event)) {
	js.Rewrite("$<.Onoffline")
	return onoffline
}

// Onoffline
func (*HTMLBodyElement) SetOnoffline(onoffline func(window.Event)) {
	js.Rewrite("$<.Onoffline = $1", onoffline)
}

// Ononline
func (*HTMLBodyElement) Ononline() (ononline func(window.Event)) {
	js.Rewrite("$<.Ononline")
	return ononline
}

// Ononline
func (*HTMLBodyElement) SetOnonline(ononline func(window.Event)) {
	js.Rewrite("$<.Ononline = $1", ononline)
}

// Onorientationchange
func (*HTMLBodyElement) Onorientationchange() (onorientationchange func(window.Event)) {
	js.Rewrite("$<.Onorientationchange")
	return onorientationchange
}

// Onorientationchange
func (*HTMLBodyElement) SetOnorientationchange(onorientationchange func(window.Event)) {
	js.Rewrite("$<.Onorientationchange = $1", onorientationchange)
}

// Onpagehide
func (*HTMLBodyElement) Onpagehide() (onpagehide func(window.Event)) {
	js.Rewrite("$<.Onpagehide")
	return onpagehide
}

// Onpagehide
func (*HTMLBodyElement) SetOnpagehide(onpagehide func(window.Event)) {
	js.Rewrite("$<.Onpagehide = $1", onpagehide)
}

// Onpageshow
func (*HTMLBodyElement) Onpageshow() (onpageshow func(window.Event)) {
	js.Rewrite("$<.Onpageshow")
	return onpageshow
}

// Onpageshow
func (*HTMLBodyElement) SetOnpageshow(onpageshow func(window.Event)) {
	js.Rewrite("$<.Onpageshow = $1", onpageshow)
}

// Onpopstate
func (*HTMLBodyElement) Onpopstate() (onpopstate func(window.Event)) {
	js.Rewrite("$<.Onpopstate")
	return onpopstate
}

// Onpopstate
func (*HTMLBodyElement) SetOnpopstate(onpopstate func(window.Event)) {
	js.Rewrite("$<.Onpopstate = $1", onpopstate)
}

// Onresize
func (*HTMLBodyElement) Onresize() (onresize func(window.Event)) {
	js.Rewrite("$<.Onresize")
	return onresize
}

// Onresize
func (*HTMLBodyElement) SetOnresize(onresize func(window.Event)) {
	js.Rewrite("$<.Onresize = $1", onresize)
}

// Onscroll
func (*HTMLBodyElement) Onscroll() (onscroll func(window.UIEvent)) {
	js.Rewrite("$<.Onscroll")
	return onscroll
}

// Onscroll
func (*HTMLBodyElement) SetOnscroll(onscroll func(window.UIEvent)) {
	js.Rewrite("$<.Onscroll = $1", onscroll)
}

// Onstorage
func (*HTMLBodyElement) Onstorage() (onstorage func(window.Event)) {
	js.Rewrite("$<.Onstorage")
	return onstorage
}

// Onstorage
func (*HTMLBodyElement) SetOnstorage(onstorage func(window.Event)) {
	js.Rewrite("$<.Onstorage = $1", onstorage)
}

// Onunload
func (*HTMLBodyElement) Onunload() (onunload func(window.Event)) {
	js.Rewrite("$<.Onunload")
	return onunload
}

// Onunload
func (*HTMLBodyElement) SetOnunload(onunload func(window.Event)) {
	js.Rewrite("$<.Onunload = $1", onunload)
}

// Text
func (*HTMLBodyElement) Text() (text interface{}) {
	js.Rewrite("$<.Text")
	return text
}

// Text
func (*HTMLBodyElement) SetText(text interface{}) {
	js.Rewrite("$<.Text = $1", text)
}

// VLink
func (*HTMLBodyElement) VLink() (vLink interface{}) {
	js.Rewrite("$<.VLink")
	return vLink
}

// VLink
func (*HTMLBodyElement) SetVLink(vLink interface{}) {
	js.Rewrite("$<.VLink = $1", vLink)
}
