package htmlfontelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLFontElement struct
// js:"HTMLFontElement,omit"
type HTMLFontElement struct {
	window.HTMLElement
}

// Color prop
func (*HTMLFontElement) Color() (color string) {
	js.Rewrite("$<.color")
	return color
}

// Color prop
func (*HTMLFontElement) SetColor(color string) {
	js.Rewrite("$<.color = $1", color)
}

// Size prop
func (*HTMLFontElement) Size() (size int) {
	js.Rewrite("$<.size")
	return size
}

// Size prop
func (*HTMLFontElement) SetSize(size int) {
	js.Rewrite("$<.size = $1", size)
}

// Face prop Sets or retrieves the current typeface family.
func (*HTMLFontElement) Face() (face string) {
	js.Rewrite("$<.face")
	return face
}

// Face prop Sets or retrieves the current typeface family.
func (*HTMLFontElement) SetFace(face string) {
	js.Rewrite("$<.face = $1", face)
}
