package htmlframesetelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLFrameSetElement struct
// js:"HTMLFrameSetElement,omit"
type HTMLFrameSetElement struct {
	window.HTMLElement
}

// Border prop
func (*HTMLFrameSetElement) Border() (border string) {
	js.Rewrite("$<.border")
	return border
}

// Border prop
func (*HTMLFrameSetElement) SetBorder(border string) {
	js.Rewrite("$<.border = $1", border)
}

// BorderColor prop Sets or retrieves the border color of the object.
func (*HTMLFrameSetElement) BorderColor() (borderColor interface{}) {
	js.Rewrite("$<.borderColor")
	return borderColor
}

// BorderColor prop Sets or retrieves the border color of the object.
func (*HTMLFrameSetElement) SetBorderColor(borderColor interface{}) {
	js.Rewrite("$<.borderColor = $1", borderColor)
}

// Cols prop Sets or retrieves the frame widths of the object.
func (*HTMLFrameSetElement) Cols() (cols string) {
	js.Rewrite("$<.cols")
	return cols
}

// Cols prop Sets or retrieves the frame widths of the object.
func (*HTMLFrameSetElement) SetCols(cols string) {
	js.Rewrite("$<.cols = $1", cols)
}

// FrameBorder prop Sets or retrieves whether to display a border for the frame.
func (*HTMLFrameSetElement) FrameBorder() (frameBorder string) {
	js.Rewrite("$<.frameBorder")
	return frameBorder
}

// FrameBorder prop Sets or retrieves whether to display a border for the frame.
func (*HTMLFrameSetElement) SetFrameBorder(frameBorder string) {
	js.Rewrite("$<.frameBorder = $1", frameBorder)
}

// FrameSpacing prop Sets or retrieves the amount of additional space between the frames.
func (*HTMLFrameSetElement) FrameSpacing() (frameSpacing interface{}) {
	js.Rewrite("$<.frameSpacing")
	return frameSpacing
}

// FrameSpacing prop Sets or retrieves the amount of additional space between the frames.
func (*HTMLFrameSetElement) SetFrameSpacing(frameSpacing interface{}) {
	js.Rewrite("$<.frameSpacing = $1", frameSpacing)
}

// Name prop
func (*HTMLFrameSetElement) Name() (name string) {
	js.Rewrite("$<.name")
	return name
}

// Name prop
func (*HTMLFrameSetElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

// Onafterprint prop
func (*HTMLFrameSetElement) Onafterprint() (onafterprint func(window.Event)) {
	js.Rewrite("$<.onafterprint")
	return onafterprint
}

// Onafterprint prop
func (*HTMLFrameSetElement) SetOnafterprint(onafterprint func(window.Event)) {
	js.Rewrite("$<.onafterprint = $1", onafterprint)
}

// Onbeforeprint prop
func (*HTMLFrameSetElement) Onbeforeprint() (onbeforeprint func(window.Event)) {
	js.Rewrite("$<.onbeforeprint")
	return onbeforeprint
}

// Onbeforeprint prop
func (*HTMLFrameSetElement) SetOnbeforeprint(onbeforeprint func(window.Event)) {
	js.Rewrite("$<.onbeforeprint = $1", onbeforeprint)
}

// Onbeforeunload prop
func (*HTMLFrameSetElement) Onbeforeunload() (onbeforeunload func(window.Event)) {
	js.Rewrite("$<.onbeforeunload")
	return onbeforeunload
}

// Onbeforeunload prop
func (*HTMLFrameSetElement) SetOnbeforeunload(onbeforeunload func(window.Event)) {
	js.Rewrite("$<.onbeforeunload = $1", onbeforeunload)
}

// Onblur prop Fires when the object loses the input focus.
func (*HTMLFrameSetElement) Onblur() (onblur func(*window.FocusEvent)) {
	js.Rewrite("$<.onblur")
	return onblur
}

// Onblur prop Fires when the object loses the input focus.
func (*HTMLFrameSetElement) SetOnblur(onblur func(*window.FocusEvent)) {
	js.Rewrite("$<.onblur = $1", onblur)
}

// Onerror prop
func (*HTMLFrameSetElement) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$<.onerror")
	return onerror
}

// Onerror prop
func (*HTMLFrameSetElement) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$<.onerror = $1", onerror)
}

// Onfocus prop Fires when the object receives focus.
func (*HTMLFrameSetElement) Onfocus() (onfocus func(*window.FocusEvent)) {
	js.Rewrite("$<.onfocus")
	return onfocus
}

// Onfocus prop Fires when the object receives focus.
func (*HTMLFrameSetElement) SetOnfocus(onfocus func(*window.FocusEvent)) {
	js.Rewrite("$<.onfocus = $1", onfocus)
}

// Onhashchange prop
func (*HTMLFrameSetElement) Onhashchange() (onhashchange func(window.Event)) {
	js.Rewrite("$<.onhashchange")
	return onhashchange
}

// Onhashchange prop
func (*HTMLFrameSetElement) SetOnhashchange(onhashchange func(window.Event)) {
	js.Rewrite("$<.onhashchange = $1", onhashchange)
}

// Onload prop
func (*HTMLFrameSetElement) Onload() (onload func(window.Event)) {
	js.Rewrite("$<.onload")
	return onload
}

// Onload prop
func (*HTMLFrameSetElement) SetOnload(onload func(window.Event)) {
	js.Rewrite("$<.onload = $1", onload)
}

// Onmessage prop
func (*HTMLFrameSetElement) Onmessage() (onmessage func(window.Event)) {
	js.Rewrite("$<.onmessage")
	return onmessage
}

// Onmessage prop
func (*HTMLFrameSetElement) SetOnmessage(onmessage func(window.Event)) {
	js.Rewrite("$<.onmessage = $1", onmessage)
}

// Onoffline prop
func (*HTMLFrameSetElement) Onoffline() (onoffline func(window.Event)) {
	js.Rewrite("$<.onoffline")
	return onoffline
}

// Onoffline prop
func (*HTMLFrameSetElement) SetOnoffline(onoffline func(window.Event)) {
	js.Rewrite("$<.onoffline = $1", onoffline)
}

// Ononline prop
func (*HTMLFrameSetElement) Ononline() (ononline func(window.Event)) {
	js.Rewrite("$<.ononline")
	return ononline
}

// Ononline prop
func (*HTMLFrameSetElement) SetOnonline(ononline func(window.Event)) {
	js.Rewrite("$<.ononline = $1", ononline)
}

// Onorientationchange prop
func (*HTMLFrameSetElement) Onorientationchange() (onorientationchange func(window.Event)) {
	js.Rewrite("$<.onorientationchange")
	return onorientationchange
}

// Onorientationchange prop
func (*HTMLFrameSetElement) SetOnorientationchange(onorientationchange func(window.Event)) {
	js.Rewrite("$<.onorientationchange = $1", onorientationchange)
}

// Onpagehide prop
func (*HTMLFrameSetElement) Onpagehide() (onpagehide func(window.Event)) {
	js.Rewrite("$<.onpagehide")
	return onpagehide
}

// Onpagehide prop
func (*HTMLFrameSetElement) SetOnpagehide(onpagehide func(window.Event)) {
	js.Rewrite("$<.onpagehide = $1", onpagehide)
}

// Onpageshow prop
func (*HTMLFrameSetElement) Onpageshow() (onpageshow func(window.Event)) {
	js.Rewrite("$<.onpageshow")
	return onpageshow
}

// Onpageshow prop
func (*HTMLFrameSetElement) SetOnpageshow(onpageshow func(window.Event)) {
	js.Rewrite("$<.onpageshow = $1", onpageshow)
}

// Onpopstate prop
func (*HTMLFrameSetElement) Onpopstate() (onpopstate func(window.Event)) {
	js.Rewrite("$<.onpopstate")
	return onpopstate
}

// Onpopstate prop
func (*HTMLFrameSetElement) SetOnpopstate(onpopstate func(window.Event)) {
	js.Rewrite("$<.onpopstate = $1", onpopstate)
}

// Onresize prop
func (*HTMLFrameSetElement) Onresize() (onresize func(window.Event)) {
	js.Rewrite("$<.onresize")
	return onresize
}

// Onresize prop
func (*HTMLFrameSetElement) SetOnresize(onresize func(window.Event)) {
	js.Rewrite("$<.onresize = $1", onresize)
}

// Onscroll prop
func (*HTMLFrameSetElement) Onscroll() (onscroll func(window.UIEvent)) {
	js.Rewrite("$<.onscroll")
	return onscroll
}

// Onscroll prop
func (*HTMLFrameSetElement) SetOnscroll(onscroll func(window.UIEvent)) {
	js.Rewrite("$<.onscroll = $1", onscroll)
}

// Onstorage prop
func (*HTMLFrameSetElement) Onstorage() (onstorage func(window.Event)) {
	js.Rewrite("$<.onstorage")
	return onstorage
}

// Onstorage prop
func (*HTMLFrameSetElement) SetOnstorage(onstorage func(window.Event)) {
	js.Rewrite("$<.onstorage = $1", onstorage)
}

// Onunload prop
func (*HTMLFrameSetElement) Onunload() (onunload func(window.Event)) {
	js.Rewrite("$<.onunload")
	return onunload
}

// Onunload prop
func (*HTMLFrameSetElement) SetOnunload(onunload func(window.Event)) {
	js.Rewrite("$<.onunload = $1", onunload)
}

// Rows prop Sets or retrieves the frame heights of the object.
func (*HTMLFrameSetElement) Rows() (rows string) {
	js.Rewrite("$<.rows")
	return rows
}

// Rows prop Sets or retrieves the frame heights of the object.
func (*HTMLFrameSetElement) SetRows(rows string) {
	js.Rewrite("$<.rows = $1", rows)
}
