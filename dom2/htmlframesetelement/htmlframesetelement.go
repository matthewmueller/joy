package htmlframesetelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"HTMLFrameSetElement,omit"
type HTMLFrameSetElement struct {
	window.HTMLElement
}

// Border
func (*HTMLFrameSetElement) Border() (border string) {
	js.Rewrite("$<.Border")
	return border
}

// Border
func (*HTMLFrameSetElement) SetBorder(border string) {
	js.Rewrite("$<.Border = $1", border)
}

// BorderColor Sets or retrieves the border color of the object.
func (*HTMLFrameSetElement) BorderColor() (borderColor interface{}) {
	js.Rewrite("$<.BorderColor")
	return borderColor
}

// BorderColor Sets or retrieves the border color of the object.
func (*HTMLFrameSetElement) SetBorderColor(borderColor interface{}) {
	js.Rewrite("$<.BorderColor = $1", borderColor)
}

// Cols Sets or retrieves the frame widths of the object.
func (*HTMLFrameSetElement) Cols() (cols string) {
	js.Rewrite("$<.Cols")
	return cols
}

// Cols Sets or retrieves the frame widths of the object.
func (*HTMLFrameSetElement) SetCols(cols string) {
	js.Rewrite("$<.Cols = $1", cols)
}

// FrameBorder Sets or retrieves whether to display a border for the frame.
func (*HTMLFrameSetElement) FrameBorder() (frameBorder string) {
	js.Rewrite("$<.FrameBorder")
	return frameBorder
}

// FrameBorder Sets or retrieves whether to display a border for the frame.
func (*HTMLFrameSetElement) SetFrameBorder(frameBorder string) {
	js.Rewrite("$<.FrameBorder = $1", frameBorder)
}

// FrameSpacing Sets or retrieves the amount of additional space between the frames.
func (*HTMLFrameSetElement) FrameSpacing() (frameSpacing interface{}) {
	js.Rewrite("$<.FrameSpacing")
	return frameSpacing
}

// FrameSpacing Sets or retrieves the amount of additional space between the frames.
func (*HTMLFrameSetElement) SetFrameSpacing(frameSpacing interface{}) {
	js.Rewrite("$<.FrameSpacing = $1", frameSpacing)
}

// Name
func (*HTMLFrameSetElement) Name() (name string) {
	js.Rewrite("$<.Name")
	return name
}

// Name
func (*HTMLFrameSetElement) SetName(name string) {
	js.Rewrite("$<.Name = $1", name)
}

// Onafterprint
func (*HTMLFrameSetElement) Onafterprint() (onafterprint func(window.Event)) {
	js.Rewrite("$<.Onafterprint")
	return onafterprint
}

// Onafterprint
func (*HTMLFrameSetElement) SetOnafterprint(onafterprint func(window.Event)) {
	js.Rewrite("$<.Onafterprint = $1", onafterprint)
}

// Onbeforeprint
func (*HTMLFrameSetElement) Onbeforeprint() (onbeforeprint func(window.Event)) {
	js.Rewrite("$<.Onbeforeprint")
	return onbeforeprint
}

// Onbeforeprint
func (*HTMLFrameSetElement) SetOnbeforeprint(onbeforeprint func(window.Event)) {
	js.Rewrite("$<.Onbeforeprint = $1", onbeforeprint)
}

// Onbeforeunload
func (*HTMLFrameSetElement) Onbeforeunload() (onbeforeunload func(window.Event)) {
	js.Rewrite("$<.Onbeforeunload")
	return onbeforeunload
}

// Onbeforeunload
func (*HTMLFrameSetElement) SetOnbeforeunload(onbeforeunload func(window.Event)) {
	js.Rewrite("$<.Onbeforeunload = $1", onbeforeunload)
}

// Onblur Fires when the object loses the input focus.
func (*HTMLFrameSetElement) Onblur() (onblur func(*window.FocusEvent)) {
	js.Rewrite("$<.Onblur")
	return onblur
}

// Onblur Fires when the object loses the input focus.
func (*HTMLFrameSetElement) SetOnblur(onblur func(*window.FocusEvent)) {
	js.Rewrite("$<.Onblur = $1", onblur)
}

// Onerror
func (*HTMLFrameSetElement) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$<.Onerror")
	return onerror
}

// Onerror
func (*HTMLFrameSetElement) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$<.Onerror = $1", onerror)
}

// Onfocus Fires when the object receives focus.
func (*HTMLFrameSetElement) Onfocus() (onfocus func(*window.FocusEvent)) {
	js.Rewrite("$<.Onfocus")
	return onfocus
}

// Onfocus Fires when the object receives focus.
func (*HTMLFrameSetElement) SetOnfocus(onfocus func(*window.FocusEvent)) {
	js.Rewrite("$<.Onfocus = $1", onfocus)
}

// Onhashchange
func (*HTMLFrameSetElement) Onhashchange() (onhashchange func(window.Event)) {
	js.Rewrite("$<.Onhashchange")
	return onhashchange
}

// Onhashchange
func (*HTMLFrameSetElement) SetOnhashchange(onhashchange func(window.Event)) {
	js.Rewrite("$<.Onhashchange = $1", onhashchange)
}

// Onload
func (*HTMLFrameSetElement) Onload() (onload func(window.Event)) {
	js.Rewrite("$<.Onload")
	return onload
}

// Onload
func (*HTMLFrameSetElement) SetOnload(onload func(window.Event)) {
	js.Rewrite("$<.Onload = $1", onload)
}

// Onmessage
func (*HTMLFrameSetElement) Onmessage() (onmessage func(window.Event)) {
	js.Rewrite("$<.Onmessage")
	return onmessage
}

// Onmessage
func (*HTMLFrameSetElement) SetOnmessage(onmessage func(window.Event)) {
	js.Rewrite("$<.Onmessage = $1", onmessage)
}

// Onoffline
func (*HTMLFrameSetElement) Onoffline() (onoffline func(window.Event)) {
	js.Rewrite("$<.Onoffline")
	return onoffline
}

// Onoffline
func (*HTMLFrameSetElement) SetOnoffline(onoffline func(window.Event)) {
	js.Rewrite("$<.Onoffline = $1", onoffline)
}

// Ononline
func (*HTMLFrameSetElement) Ononline() (ononline func(window.Event)) {
	js.Rewrite("$<.Ononline")
	return ononline
}

// Ononline
func (*HTMLFrameSetElement) SetOnonline(ononline func(window.Event)) {
	js.Rewrite("$<.Ononline = $1", ononline)
}

// Onorientationchange
func (*HTMLFrameSetElement) Onorientationchange() (onorientationchange func(window.Event)) {
	js.Rewrite("$<.Onorientationchange")
	return onorientationchange
}

// Onorientationchange
func (*HTMLFrameSetElement) SetOnorientationchange(onorientationchange func(window.Event)) {
	js.Rewrite("$<.Onorientationchange = $1", onorientationchange)
}

// Onpagehide
func (*HTMLFrameSetElement) Onpagehide() (onpagehide func(window.Event)) {
	js.Rewrite("$<.Onpagehide")
	return onpagehide
}

// Onpagehide
func (*HTMLFrameSetElement) SetOnpagehide(onpagehide func(window.Event)) {
	js.Rewrite("$<.Onpagehide = $1", onpagehide)
}

// Onpageshow
func (*HTMLFrameSetElement) Onpageshow() (onpageshow func(window.Event)) {
	js.Rewrite("$<.Onpageshow")
	return onpageshow
}

// Onpageshow
func (*HTMLFrameSetElement) SetOnpageshow(onpageshow func(window.Event)) {
	js.Rewrite("$<.Onpageshow = $1", onpageshow)
}

// Onpopstate
func (*HTMLFrameSetElement) Onpopstate() (onpopstate func(window.Event)) {
	js.Rewrite("$<.Onpopstate")
	return onpopstate
}

// Onpopstate
func (*HTMLFrameSetElement) SetOnpopstate(onpopstate func(window.Event)) {
	js.Rewrite("$<.Onpopstate = $1", onpopstate)
}

// Onresize
func (*HTMLFrameSetElement) Onresize() (onresize func(window.Event)) {
	js.Rewrite("$<.Onresize")
	return onresize
}

// Onresize
func (*HTMLFrameSetElement) SetOnresize(onresize func(window.Event)) {
	js.Rewrite("$<.Onresize = $1", onresize)
}

// Onscroll
func (*HTMLFrameSetElement) Onscroll() (onscroll func(window.UIEvent)) {
	js.Rewrite("$<.Onscroll")
	return onscroll
}

// Onscroll
func (*HTMLFrameSetElement) SetOnscroll(onscroll func(window.UIEvent)) {
	js.Rewrite("$<.Onscroll = $1", onscroll)
}

// Onstorage
func (*HTMLFrameSetElement) Onstorage() (onstorage func(window.Event)) {
	js.Rewrite("$<.Onstorage")
	return onstorage
}

// Onstorage
func (*HTMLFrameSetElement) SetOnstorage(onstorage func(window.Event)) {
	js.Rewrite("$<.Onstorage = $1", onstorage)
}

// Onunload
func (*HTMLFrameSetElement) Onunload() (onunload func(window.Event)) {
	js.Rewrite("$<.Onunload")
	return onunload
}

// Onunload
func (*HTMLFrameSetElement) SetOnunload(onunload func(window.Event)) {
	js.Rewrite("$<.Onunload = $1", onunload)
}

// Rows Sets or retrieves the frame heights of the object.
func (*HTMLFrameSetElement) Rows() (rows string) {
	js.Rewrite("$<.Rows")
	return rows
}

// Rows Sets or retrieves the frame heights of the object.
func (*HTMLFrameSetElement) SetRows(rows string) {
	js.Rewrite("$<.Rows = $1", rows)
}
