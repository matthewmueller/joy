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

// Color
func (*HTMLHRElement) Color() (color string) {
	js.Rewrite("$<.Color")
	return color
}

// Color
func (*HTMLHRElement) SetColor(color string) {
	js.Rewrite("$<.Color = $1", color)
}

// Size
func (*HTMLHRElement) Size() (size int) {
	js.Rewrite("$<.Size")
	return size
}

// Size
func (*HTMLHRElement) SetSize(size int) {
	js.Rewrite("$<.Size = $1", size)
}

// Align Sets or retrieves how the object is aligned with adjacent text.
func (*HTMLHRElement) Align() (align string) {
	js.Rewrite("$<.Align")
	return align
}

// Align Sets or retrieves how the object is aligned with adjacent text.
func (*HTMLHRElement) SetAlign(align string) {
	js.Rewrite("$<.Align = $1", align)
}

// NoShade Sets or retrieves whether the horizontal rule is drawn with 3-D shading.
func (*HTMLHRElement) NoShade() (noShade bool) {
	js.Rewrite("$<.NoShade")
	return noShade
}

// NoShade Sets or retrieves whether the horizontal rule is drawn with 3-D shading.
func (*HTMLHRElement) SetNoShade(noShade bool) {
	js.Rewrite("$<.NoShade = $1", noShade)
}

// Width Sets or retrieves the width of the object.
func (*HTMLHRElement) Width() (width int) {
	js.Rewrite("$<.Width")
	return width
}

// Width Sets or retrieves the width of the object.
func (*HTMLHRElement) SetWidth(width int) {
	js.Rewrite("$<.Width = $1", width)
}
