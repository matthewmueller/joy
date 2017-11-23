package htmliframeelement

import (
	"github.com/matthewmueller/golly/dom2/getsvgdocument"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"HTMLIFrameElement,omit"
type HTMLIFrameElement struct {
	window.HTMLElement
	getsvgdocument.GetSVGDocument
}

// Align Sets or retrieves how the object is aligned with adjacent text.
func (*HTMLIFrameElement) Align() (align string) {
	js.Rewrite("$<.Align")
	return align
}

// Align Sets or retrieves how the object is aligned with adjacent text.
func (*HTMLIFrameElement) SetAlign(align string) {
	js.Rewrite("$<.Align = $1", align)
}

// AllowFullscreen
func (*HTMLIFrameElement) AllowFullscreen() (allowFullscreen bool) {
	js.Rewrite("$<.AllowFullscreen")
	return allowFullscreen
}

// AllowFullscreen
func (*HTMLIFrameElement) SetAllowFullscreen(allowFullscreen bool) {
	js.Rewrite("$<.AllowFullscreen = $1", allowFullscreen)
}

// AllowPaymentRequest
func (*HTMLIFrameElement) AllowPaymentRequest() (allowPaymentRequest bool) {
	js.Rewrite("$<.AllowPaymentRequest")
	return allowPaymentRequest
}

// AllowPaymentRequest
func (*HTMLIFrameElement) SetAllowPaymentRequest(allowPaymentRequest bool) {
	js.Rewrite("$<.AllowPaymentRequest = $1", allowPaymentRequest)
}

// Border Specifies the properties of a border drawn around an object.
func (*HTMLIFrameElement) Border() (border string) {
	js.Rewrite("$<.Border")
	return border
}

// Border Specifies the properties of a border drawn around an object.
func (*HTMLIFrameElement) SetBorder(border string) {
	js.Rewrite("$<.Border = $1", border)
}

// ContentDocument Retrieves the document object of the page or frame.
func (*HTMLIFrameElement) ContentDocument() (contentDocument window.Document) {
	js.Rewrite("$<.ContentDocument")
	return contentDocument
}

// ContentWindow Retrieves the object of the specified.
func (*HTMLIFrameElement) ContentWindow() (contentWindow *window.Window) {
	js.Rewrite("$<.ContentWindow")
	return contentWindow
}

// FrameBorder Sets or retrieves whether to display a border for the frame.
func (*HTMLIFrameElement) FrameBorder() (frameBorder string) {
	js.Rewrite("$<.FrameBorder")
	return frameBorder
}

// FrameBorder Sets or retrieves whether to display a border for the frame.
func (*HTMLIFrameElement) SetFrameBorder(frameBorder string) {
	js.Rewrite("$<.FrameBorder = $1", frameBorder)
}

// FrameSpacing Sets or retrieves the amount of additional space between the frames.
func (*HTMLIFrameElement) FrameSpacing() (frameSpacing interface{}) {
	js.Rewrite("$<.FrameSpacing")
	return frameSpacing
}

// FrameSpacing Sets or retrieves the amount of additional space between the frames.
func (*HTMLIFrameElement) SetFrameSpacing(frameSpacing interface{}) {
	js.Rewrite("$<.FrameSpacing = $1", frameSpacing)
}

// Height Sets or retrieves the height of the object.
func (*HTMLIFrameElement) Height() (height string) {
	js.Rewrite("$<.Height")
	return height
}

// Height Sets or retrieves the height of the object.
func (*HTMLIFrameElement) SetHeight(height string) {
	js.Rewrite("$<.Height = $1", height)
}

// Hspace Sets or retrieves the horizontal margin for the object.
func (*HTMLIFrameElement) Hspace() (hspace int) {
	js.Rewrite("$<.Hspace")
	return hspace
}

// Hspace Sets or retrieves the horizontal margin for the object.
func (*HTMLIFrameElement) SetHspace(hspace int) {
	js.Rewrite("$<.Hspace = $1", hspace)
}

// LongDesc Sets or retrieves a URI to a long description of the object.
func (*HTMLIFrameElement) LongDesc() (longDesc string) {
	js.Rewrite("$<.LongDesc")
	return longDesc
}

// LongDesc Sets or retrieves a URI to a long description of the object.
func (*HTMLIFrameElement) SetLongDesc(longDesc string) {
	js.Rewrite("$<.LongDesc = $1", longDesc)
}

// MarginHeight Sets or retrieves the top and bottom margin heights before displaying the text in a frame.
func (*HTMLIFrameElement) MarginHeight() (marginHeight string) {
	js.Rewrite("$<.MarginHeight")
	return marginHeight
}

// MarginHeight Sets or retrieves the top and bottom margin heights before displaying the text in a frame.
func (*HTMLIFrameElement) SetMarginHeight(marginHeight string) {
	js.Rewrite("$<.MarginHeight = $1", marginHeight)
}

// MarginWidth Sets or retrieves the left and right margin widths before displaying the text in a frame.
func (*HTMLIFrameElement) MarginWidth() (marginWidth string) {
	js.Rewrite("$<.MarginWidth")
	return marginWidth
}

// MarginWidth Sets or retrieves the left and right margin widths before displaying the text in a frame.
func (*HTMLIFrameElement) SetMarginWidth(marginWidth string) {
	js.Rewrite("$<.MarginWidth = $1", marginWidth)
}

// Name Sets or retrieves the frame name.
func (*HTMLIFrameElement) Name() (name string) {
	js.Rewrite("$<.Name")
	return name
}

// Name Sets or retrieves the frame name.
func (*HTMLIFrameElement) SetName(name string) {
	js.Rewrite("$<.Name = $1", name)
}

// NoResize Sets or retrieves whether the user can resize the frame.
func (*HTMLIFrameElement) NoResize() (noResize bool) {
	js.Rewrite("$<.NoResize")
	return noResize
}

// NoResize Sets or retrieves whether the user can resize the frame.
func (*HTMLIFrameElement) SetNoResize(noResize bool) {
	js.Rewrite("$<.NoResize = $1", noResize)
}

// Onload Raised when the object has been completely received from the server.
func (*HTMLIFrameElement) Onload() (onload func(window.Event)) {
	js.Rewrite("$<.Onload")
	return onload
}

// Onload Raised when the object has been completely received from the server.
func (*HTMLIFrameElement) SetOnload(onload func(window.Event)) {
	js.Rewrite("$<.Onload = $1", onload)
}

// Sandbox
func (*HTMLIFrameElement) Sandbox() (sandbox *domsettabletokenlist.DOMSettableTokenList) {
	js.Rewrite("$<.Sandbox")
	return sandbox
}

// Scrolling Sets or retrieves whether the frame can be scrolled.
func (*HTMLIFrameElement) Scrolling() (scrolling string) {
	js.Rewrite("$<.Scrolling")
	return scrolling
}

// Scrolling Sets or retrieves whether the frame can be scrolled.
func (*HTMLIFrameElement) SetScrolling(scrolling string) {
	js.Rewrite("$<.Scrolling = $1", scrolling)
}

// Src Sets or retrieves a URL to be loaded by the object.
func (*HTMLIFrameElement) Src() (src string) {
	js.Rewrite("$<.Src")
	return src
}

// Src Sets or retrieves a URL to be loaded by the object.
func (*HTMLIFrameElement) SetSrc(src string) {
	js.Rewrite("$<.Src = $1", src)
}

// Vspace Sets or retrieves the vertical margin for the object.
func (*HTMLIFrameElement) Vspace() (vspace int) {
	js.Rewrite("$<.Vspace")
	return vspace
}

// Vspace Sets or retrieves the vertical margin for the object.
func (*HTMLIFrameElement) SetVspace(vspace int) {
	js.Rewrite("$<.Vspace = $1", vspace)
}

// Width Sets or retrieves the width of the object.
func (*HTMLIFrameElement) Width() (width string) {
	js.Rewrite("$<.Width")
	return width
}

// Width Sets or retrieves the width of the object.
func (*HTMLIFrameElement) SetWidth(width string) {
	js.Rewrite("$<.Width = $1", width)
}
