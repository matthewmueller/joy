package htmliframeelement

import (
	"github.com/matthewmueller/golly/dom2/domsettabletokenlist"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLIFrameElement struct
// js:"HTMLIFrameElement,omit"
type HTMLIFrameElement struct {
	window.HTMLElement
}

// GetSVGDocument fn
func (*HTMLIFrameElement) GetSVGDocument() (w window.Document) {
	js.Rewrite("$<.getSVGDocument()")
	return w
}

// Align prop Sets or retrieves how the object is aligned with adjacent text.
func (*HTMLIFrameElement) Align() (align string) {
	js.Rewrite("$<.align")
	return align
}

// Align prop Sets or retrieves how the object is aligned with adjacent text.
func (*HTMLIFrameElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

// AllowFullscreen prop
func (*HTMLIFrameElement) AllowFullscreen() (allowFullscreen bool) {
	js.Rewrite("$<.allowFullscreen")
	return allowFullscreen
}

// AllowFullscreen prop
func (*HTMLIFrameElement) SetAllowFullscreen(allowFullscreen bool) {
	js.Rewrite("$<.allowFullscreen = $1", allowFullscreen)
}

// AllowPaymentRequest prop
func (*HTMLIFrameElement) AllowPaymentRequest() (allowPaymentRequest bool) {
	js.Rewrite("$<.allowPaymentRequest")
	return allowPaymentRequest
}

// AllowPaymentRequest prop
func (*HTMLIFrameElement) SetAllowPaymentRequest(allowPaymentRequest bool) {
	js.Rewrite("$<.allowPaymentRequest = $1", allowPaymentRequest)
}

// Border prop Specifies the properties of a border drawn around an object.
func (*HTMLIFrameElement) Border() (border string) {
	js.Rewrite("$<.border")
	return border
}

// Border prop Specifies the properties of a border drawn around an object.
func (*HTMLIFrameElement) SetBorder(border string) {
	js.Rewrite("$<.border = $1", border)
}

// ContentDocument prop Retrieves the document object of the page or frame.
func (*HTMLIFrameElement) ContentDocument() (contentDocument window.Document) {
	js.Rewrite("$<.contentDocument")
	return contentDocument
}

// ContentWindow prop Retrieves the object of the specified.
func (*HTMLIFrameElement) ContentWindow() (contentWindow *window.Window) {
	js.Rewrite("$<.contentWindow")
	return contentWindow
}

// FrameBorder prop Sets or retrieves whether to display a border for the frame.
func (*HTMLIFrameElement) FrameBorder() (frameBorder string) {
	js.Rewrite("$<.frameBorder")
	return frameBorder
}

// FrameBorder prop Sets or retrieves whether to display a border for the frame.
func (*HTMLIFrameElement) SetFrameBorder(frameBorder string) {
	js.Rewrite("$<.frameBorder = $1", frameBorder)
}

// FrameSpacing prop Sets or retrieves the amount of additional space between the frames.
func (*HTMLIFrameElement) FrameSpacing() (frameSpacing interface{}) {
	js.Rewrite("$<.frameSpacing")
	return frameSpacing
}

// FrameSpacing prop Sets or retrieves the amount of additional space between the frames.
func (*HTMLIFrameElement) SetFrameSpacing(frameSpacing interface{}) {
	js.Rewrite("$<.frameSpacing = $1", frameSpacing)
}

// Height prop Sets or retrieves the height of the object.
func (*HTMLIFrameElement) Height() (height string) {
	js.Rewrite("$<.height")
	return height
}

// Height prop Sets or retrieves the height of the object.
func (*HTMLIFrameElement) SetHeight(height string) {
	js.Rewrite("$<.height = $1", height)
}

// Hspace prop Sets or retrieves the horizontal margin for the object.
func (*HTMLIFrameElement) Hspace() (hspace int) {
	js.Rewrite("$<.hspace")
	return hspace
}

// Hspace prop Sets or retrieves the horizontal margin for the object.
func (*HTMLIFrameElement) SetHspace(hspace int) {
	js.Rewrite("$<.hspace = $1", hspace)
}

// LongDesc prop Sets or retrieves a URI to a long description of the object.
func (*HTMLIFrameElement) LongDesc() (longDesc string) {
	js.Rewrite("$<.longDesc")
	return longDesc
}

// LongDesc prop Sets or retrieves a URI to a long description of the object.
func (*HTMLIFrameElement) SetLongDesc(longDesc string) {
	js.Rewrite("$<.longDesc = $1", longDesc)
}

// MarginHeight prop Sets or retrieves the top and bottom margin heights before displaying the text in a frame.
func (*HTMLIFrameElement) MarginHeight() (marginHeight string) {
	js.Rewrite("$<.marginHeight")
	return marginHeight
}

// MarginHeight prop Sets or retrieves the top and bottom margin heights before displaying the text in a frame.
func (*HTMLIFrameElement) SetMarginHeight(marginHeight string) {
	js.Rewrite("$<.marginHeight = $1", marginHeight)
}

// MarginWidth prop Sets or retrieves the left and right margin widths before displaying the text in a frame.
func (*HTMLIFrameElement) MarginWidth() (marginWidth string) {
	js.Rewrite("$<.marginWidth")
	return marginWidth
}

// MarginWidth prop Sets or retrieves the left and right margin widths before displaying the text in a frame.
func (*HTMLIFrameElement) SetMarginWidth(marginWidth string) {
	js.Rewrite("$<.marginWidth = $1", marginWidth)
}

// Name prop Sets or retrieves the frame name.
func (*HTMLIFrameElement) Name() (name string) {
	js.Rewrite("$<.name")
	return name
}

// Name prop Sets or retrieves the frame name.
func (*HTMLIFrameElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

// NoResize prop Sets or retrieves whether the user can resize the frame.
func (*HTMLIFrameElement) NoResize() (noResize bool) {
	js.Rewrite("$<.noResize")
	return noResize
}

// NoResize prop Sets or retrieves whether the user can resize the frame.
func (*HTMLIFrameElement) SetNoResize(noResize bool) {
	js.Rewrite("$<.noResize = $1", noResize)
}

// Onload prop Raised when the object has been completely received from the server.
func (*HTMLIFrameElement) Onload() (onload func(window.Event)) {
	js.Rewrite("$<.onload")
	return onload
}

// Onload prop Raised when the object has been completely received from the server.
func (*HTMLIFrameElement) SetOnload(onload func(window.Event)) {
	js.Rewrite("$<.onload = $1", onload)
}

// Sandbox prop
func (*HTMLIFrameElement) Sandbox() (sandbox *domsettabletokenlist.DOMSettableTokenList) {
	js.Rewrite("$<.sandbox")
	return sandbox
}

// Scrolling prop Sets or retrieves whether the frame can be scrolled.
func (*HTMLIFrameElement) Scrolling() (scrolling string) {
	js.Rewrite("$<.scrolling")
	return scrolling
}

// Scrolling prop Sets or retrieves whether the frame can be scrolled.
func (*HTMLIFrameElement) SetScrolling(scrolling string) {
	js.Rewrite("$<.scrolling = $1", scrolling)
}

// Src prop Sets or retrieves a URL to be loaded by the object.
func (*HTMLIFrameElement) Src() (src string) {
	js.Rewrite("$<.src")
	return src
}

// Src prop Sets or retrieves a URL to be loaded by the object.
func (*HTMLIFrameElement) SetSrc(src string) {
	js.Rewrite("$<.src = $1", src)
}

// Vspace prop Sets or retrieves the vertical margin for the object.
func (*HTMLIFrameElement) Vspace() (vspace int) {
	js.Rewrite("$<.vspace")
	return vspace
}

// Vspace prop Sets or retrieves the vertical margin for the object.
func (*HTMLIFrameElement) SetVspace(vspace int) {
	js.Rewrite("$<.vspace = $1", vspace)
}

// Width prop Sets or retrieves the width of the object.
func (*HTMLIFrameElement) Width() (width string) {
	js.Rewrite("$<.width")
	return width
}

// Width prop Sets or retrieves the width of the object.
func (*HTMLIFrameElement) SetWidth(width string) {
	js.Rewrite("$<.width = $1", width)
}
