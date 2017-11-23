package htmlhrelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLHRElement struct
// js:"HTMLHRElement,omit"
type HTMLHRElement struct {
	window.HTMLElement
}

// Color prop
func (*HTMLHRElement) Color() (color string) {
	js.Rewrite("$<.color")
	return color
}

// Color prop
func (*HTMLHRElement) SetColor(color string) {
	js.Rewrite("$<.color = $1", color)
}

// Size prop
func (*HTMLHRElement) Size() (size int) {
	js.Rewrite("$<.size")
	return size
}

// Size prop
func (*HTMLHRElement) SetSize(size int) {
	js.Rewrite("$<.size = $1", size)
}

// Align prop Sets or retrieves how the object is aligned with adjacent text.
func (*HTMLHRElement) Align() (align string) {
	js.Rewrite("$<.align")
	return align
}

// Align prop Sets or retrieves how the object is aligned with adjacent text.
func (*HTMLHRElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

// NoShade prop Sets or retrieves whether the horizontal rule is drawn with 3-D shading.
func (*HTMLHRElement) NoShade() (noShade bool) {
	js.Rewrite("$<.noShade")
	return noShade
}

// NoShade prop Sets or retrieves whether the horizontal rule is drawn with 3-D shading.
func (*HTMLHRElement) SetNoShade(noShade bool) {
	js.Rewrite("$<.noShade = $1", noShade)
}

// Width prop Sets or retrieves the width of the object.
func (*HTMLHRElement) Width() (width int) {
	js.Rewrite("$<.width")
	return width
}

// Width prop Sets or retrieves the width of the object.
func (*HTMLHRElement) SetWidth(width int) {
	js.Rewrite("$<.width = $1", width)
}
