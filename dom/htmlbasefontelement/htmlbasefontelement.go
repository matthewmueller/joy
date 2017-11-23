package htmlbasefontelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLBaseFontElement struct
// js:"HTMLBaseFontElement,omit"
type HTMLBaseFontElement struct {
	window.HTMLElement
}

// Color prop
func (*HTMLBaseFontElement) Color() (color string) {
	js.Rewrite("$<.color")
	return color
}

// Color prop
func (*HTMLBaseFontElement) SetColor(color string) {
	js.Rewrite("$<.color = $1", color)
}

// Face prop Sets or retrieves the current typeface family.
func (*HTMLBaseFontElement) Face() (face string) {
	js.Rewrite("$<.face")
	return face
}

// Face prop Sets or retrieves the current typeface family.
func (*HTMLBaseFontElement) SetFace(face string) {
	js.Rewrite("$<.face = $1", face)
}

// Size prop Sets or retrieves the font size of the object.
func (*HTMLBaseFontElement) Size() (size int) {
	js.Rewrite("$<.size")
	return size
}

// Size prop Sets or retrieves the font size of the object.
func (*HTMLBaseFontElement) SetSize(size int) {
	js.Rewrite("$<.size = $1", size)
}
