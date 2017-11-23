package htmlframeelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLFrameElement struct
// js:"HTMLFrameElement,omit"
type HTMLFrameElement struct {
	window.HTMLElement
}

// GetSVGDocument fn
func (*HTMLFrameElement) GetSVGDocument() (w window.Document) {
	js.Rewrite("$<.getSVGDocument()")
	return w
}

// Border prop Specifies the properties of a border drawn around an object.
func (*HTMLFrameElement) Border() (border string) {
	js.Rewrite("$<.border")
	return border
}

// Border prop Specifies the properties of a border drawn around an object.
func (*HTMLFrameElement) SetBorder(border string) {
	js.Rewrite("$<.border = $1", border)
}

// BorderColor prop Sets or retrieves the border color of the object.
func (*HTMLFrameElement) BorderColor() (borderColor interface{}) {
	js.Rewrite("$<.borderColor")
	return borderColor
}

// BorderColor prop Sets or retrieves the border color of the object.
func (*HTMLFrameElement) SetBorderColor(borderColor interface{}) {
	js.Rewrite("$<.borderColor = $1", borderColor)
}

// ContentDocument prop Retrieves the document object of the page or frame.
func (*HTMLFrameElement) ContentDocument() (contentDocument window.Document) {
	js.Rewrite("$<.contentDocument")
	return contentDocument
}

// ContentWindow prop Retrieves the object of the specified.
func (*HTMLFrameElement) ContentWindow() (contentWindow *window.Window) {
	js.Rewrite("$<.contentWindow")
	return contentWindow
}

// FrameBorder prop Sets or retrieves whether to display a border for the frame.
func (*HTMLFrameElement) FrameBorder() (frameBorder string) {
	js.Rewrite("$<.frameBorder")
	return frameBorder
}

// FrameBorder prop Sets or retrieves whether to display a border for the frame.
func (*HTMLFrameElement) SetFrameBorder(frameBorder string) {
	js.Rewrite("$<.frameBorder = $1", frameBorder)
}

// FrameSpacing prop Sets or retrieves the amount of additional space between the frames.
func (*HTMLFrameElement) FrameSpacing() (frameSpacing interface{}) {
	js.Rewrite("$<.frameSpacing")
	return frameSpacing
}

// FrameSpacing prop Sets or retrieves the amount of additional space between the frames.
func (*HTMLFrameElement) SetFrameSpacing(frameSpacing interface{}) {
	js.Rewrite("$<.frameSpacing = $1", frameSpacing)
}

// Height prop Sets or retrieves the height of the object.
func (*HTMLFrameElement) Height() (height interface{}) {
	js.Rewrite("$<.height")
	return height
}

// Height prop Sets or retrieves the height of the object.
func (*HTMLFrameElement) SetHeight(height interface{}) {
	js.Rewrite("$<.height = $1", height)
}

// LongDesc prop Sets or retrieves a URI to a long description of the object.
func (*HTMLFrameElement) LongDesc() (longDesc string) {
	js.Rewrite("$<.longDesc")
	return longDesc
}

// LongDesc prop Sets or retrieves a URI to a long description of the object.
func (*HTMLFrameElement) SetLongDesc(longDesc string) {
	js.Rewrite("$<.longDesc = $1", longDesc)
}

// MarginHeight prop Sets or retrieves the top and bottom margin heights before displaying the text in a frame.
func (*HTMLFrameElement) MarginHeight() (marginHeight string) {
	js.Rewrite("$<.marginHeight")
	return marginHeight
}

// MarginHeight prop Sets or retrieves the top and bottom margin heights before displaying the text in a frame.
func (*HTMLFrameElement) SetMarginHeight(marginHeight string) {
	js.Rewrite("$<.marginHeight = $1", marginHeight)
}

// MarginWidth prop Sets or retrieves the left and right margin widths before displaying the text in a frame.
func (*HTMLFrameElement) MarginWidth() (marginWidth string) {
	js.Rewrite("$<.marginWidth")
	return marginWidth
}

// MarginWidth prop Sets or retrieves the left and right margin widths before displaying the text in a frame.
func (*HTMLFrameElement) SetMarginWidth(marginWidth string) {
	js.Rewrite("$<.marginWidth = $1", marginWidth)
}

// Name prop Sets or retrieves the frame name.
func (*HTMLFrameElement) Name() (name string) {
	js.Rewrite("$<.name")
	return name
}

// Name prop Sets or retrieves the frame name.
func (*HTMLFrameElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

// NoResize prop Sets or retrieves whether the user can resize the frame.
func (*HTMLFrameElement) NoResize() (noResize bool) {
	js.Rewrite("$<.noResize")
	return noResize
}

// NoResize prop Sets or retrieves whether the user can resize the frame.
func (*HTMLFrameElement) SetNoResize(noResize bool) {
	js.Rewrite("$<.noResize = $1", noResize)
}

// Onload prop Raised when the object has been completely received from the server.
func (*HTMLFrameElement) Onload() (onload func(window.Event)) {
	js.Rewrite("$<.onload")
	return onload
}

// Onload prop Raised when the object has been completely received from the server.
func (*HTMLFrameElement) SetOnload(onload func(window.Event)) {
	js.Rewrite("$<.onload = $1", onload)
}

// Scrolling prop Sets or retrieves whether the frame can be scrolled.
func (*HTMLFrameElement) Scrolling() (scrolling string) {
	js.Rewrite("$<.scrolling")
	return scrolling
}

// Scrolling prop Sets or retrieves whether the frame can be scrolled.
func (*HTMLFrameElement) SetScrolling(scrolling string) {
	js.Rewrite("$<.scrolling = $1", scrolling)
}

// Src prop Sets or retrieves a URL to be loaded by the object.
func (*HTMLFrameElement) Src() (src string) {
	js.Rewrite("$<.src")
	return src
}

// Src prop Sets or retrieves a URL to be loaded by the object.
func (*HTMLFrameElement) SetSrc(src string) {
	js.Rewrite("$<.src = $1", src)
}

// Width prop Sets or retrieves the width of the object.
func (*HTMLFrameElement) Width() (width interface{}) {
	js.Rewrite("$<.width")
	return width
}

// Width prop Sets or retrieves the width of the object.
func (*HTMLFrameElement) SetWidth(width interface{}) {
	js.Rewrite("$<.width = $1", width)
}
