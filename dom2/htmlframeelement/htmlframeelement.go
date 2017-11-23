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

// GetSVGDocument
func (*HTMLFrameElement) GetSVGDocument() (w window.Document) {
	js.Rewrite("$<.GetSVGDocument()")
	return w
}

// Border Specifies the properties of a border drawn around an object.
func (*HTMLFrameElement) Border() (border string) {
	js.Rewrite("$<.Border")
	return border
}

// Border Specifies the properties of a border drawn around an object.
func (*HTMLFrameElement) SetBorder(border string) {
	js.Rewrite("$<.Border = $1", border)
}

// BorderColor Sets or retrieves the border color of the object.
func (*HTMLFrameElement) BorderColor() (borderColor interface{}) {
	js.Rewrite("$<.BorderColor")
	return borderColor
}

// BorderColor Sets or retrieves the border color of the object.
func (*HTMLFrameElement) SetBorderColor(borderColor interface{}) {
	js.Rewrite("$<.BorderColor = $1", borderColor)
}

// ContentDocument Retrieves the document object of the page or frame.
func (*HTMLFrameElement) ContentDocument() (contentDocument window.Document) {
	js.Rewrite("$<.ContentDocument")
	return contentDocument
}

// ContentWindow Retrieves the object of the specified.
func (*HTMLFrameElement) ContentWindow() (contentWindow *window.Window) {
	js.Rewrite("$<.ContentWindow")
	return contentWindow
}

// FrameBorder Sets or retrieves whether to display a border for the frame.
func (*HTMLFrameElement) FrameBorder() (frameBorder string) {
	js.Rewrite("$<.FrameBorder")
	return frameBorder
}

// FrameBorder Sets or retrieves whether to display a border for the frame.
func (*HTMLFrameElement) SetFrameBorder(frameBorder string) {
	js.Rewrite("$<.FrameBorder = $1", frameBorder)
}

// FrameSpacing Sets or retrieves the amount of additional space between the frames.
func (*HTMLFrameElement) FrameSpacing() (frameSpacing interface{}) {
	js.Rewrite("$<.FrameSpacing")
	return frameSpacing
}

// FrameSpacing Sets or retrieves the amount of additional space between the frames.
func (*HTMLFrameElement) SetFrameSpacing(frameSpacing interface{}) {
	js.Rewrite("$<.FrameSpacing = $1", frameSpacing)
}

// Height Sets or retrieves the height of the object.
func (*HTMLFrameElement) Height() (height interface{}) {
	js.Rewrite("$<.Height")
	return height
}

// Height Sets or retrieves the height of the object.
func (*HTMLFrameElement) SetHeight(height interface{}) {
	js.Rewrite("$<.Height = $1", height)
}

// LongDesc Sets or retrieves a URI to a long description of the object.
func (*HTMLFrameElement) LongDesc() (longDesc string) {
	js.Rewrite("$<.LongDesc")
	return longDesc
}

// LongDesc Sets or retrieves a URI to a long description of the object.
func (*HTMLFrameElement) SetLongDesc(longDesc string) {
	js.Rewrite("$<.LongDesc = $1", longDesc)
}

// MarginHeight Sets or retrieves the top and bottom margin heights before displaying the text in a frame.
func (*HTMLFrameElement) MarginHeight() (marginHeight string) {
	js.Rewrite("$<.MarginHeight")
	return marginHeight
}

// MarginHeight Sets or retrieves the top and bottom margin heights before displaying the text in a frame.
func (*HTMLFrameElement) SetMarginHeight(marginHeight string) {
	js.Rewrite("$<.MarginHeight = $1", marginHeight)
}

// MarginWidth Sets or retrieves the left and right margin widths before displaying the text in a frame.
func (*HTMLFrameElement) MarginWidth() (marginWidth string) {
	js.Rewrite("$<.MarginWidth")
	return marginWidth
}

// MarginWidth Sets or retrieves the left and right margin widths before displaying the text in a frame.
func (*HTMLFrameElement) SetMarginWidth(marginWidth string) {
	js.Rewrite("$<.MarginWidth = $1", marginWidth)
}

// Name Sets or retrieves the frame name.
func (*HTMLFrameElement) Name() (name string) {
	js.Rewrite("$<.Name")
	return name
}

// Name Sets or retrieves the frame name.
func (*HTMLFrameElement) SetName(name string) {
	js.Rewrite("$<.Name = $1", name)
}

// NoResize Sets or retrieves whether the user can resize the frame.
func (*HTMLFrameElement) NoResize() (noResize bool) {
	js.Rewrite("$<.NoResize")
	return noResize
}

// NoResize Sets or retrieves whether the user can resize the frame.
func (*HTMLFrameElement) SetNoResize(noResize bool) {
	js.Rewrite("$<.NoResize = $1", noResize)
}

// Onload Raised when the object has been completely received from the server.
func (*HTMLFrameElement) Onload() (onload func(window.Event)) {
	js.Rewrite("$<.Onload")
	return onload
}

// Onload Raised when the object has been completely received from the server.
func (*HTMLFrameElement) SetOnload(onload func(window.Event)) {
	js.Rewrite("$<.Onload = $1", onload)
}

// Scrolling Sets or retrieves whether the frame can be scrolled.
func (*HTMLFrameElement) Scrolling() (scrolling string) {
	js.Rewrite("$<.Scrolling")
	return scrolling
}

// Scrolling Sets or retrieves whether the frame can be scrolled.
func (*HTMLFrameElement) SetScrolling(scrolling string) {
	js.Rewrite("$<.Scrolling = $1", scrolling)
}

// Src Sets or retrieves a URL to be loaded by the object.
func (*HTMLFrameElement) Src() (src string) {
	js.Rewrite("$<.Src")
	return src
}

// Src Sets or retrieves a URL to be loaded by the object.
func (*HTMLFrameElement) SetSrc(src string) {
	js.Rewrite("$<.Src = $1", src)
}

// Width Sets or retrieves the width of the object.
func (*HTMLFrameElement) Width() (width interface{}) {
	js.Rewrite("$<.Width")
	return width
}

// Width Sets or retrieves the width of the object.
func (*HTMLFrameElement) SetWidth(width interface{}) {
	js.Rewrite("$<.Width = $1", width)
}
