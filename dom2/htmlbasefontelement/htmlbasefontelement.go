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

// Color
func (*HTMLBaseFontElement) Color() (color string) {
	js.Rewrite("$<.Color")
	return color
}

// Color
func (*HTMLBaseFontElement) SetColor(color string) {
	js.Rewrite("$<.Color = $1", color)
}

// Face Sets or retrieves the current typeface family.
func (*HTMLBaseFontElement) Face() (face string) {
	js.Rewrite("$<.Face")
	return face
}

// Face Sets or retrieves the current typeface family.
func (*HTMLBaseFontElement) SetFace(face string) {
	js.Rewrite("$<.Face = $1", face)
}

// Size Sets or retrieves the font size of the object.
func (*HTMLBaseFontElement) Size() (size int) {
	js.Rewrite("$<.Size")
	return size
}

// Size Sets or retrieves the font size of the object.
func (*HTMLBaseFontElement) SetSize(size int) {
	js.Rewrite("$<.Size = $1", size)
}
