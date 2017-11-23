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

// Color
func (*HTMLFontElement) Color() (color string) {
	js.Rewrite("$<.Color")
	return color
}

// Color
func (*HTMLFontElement) SetColor(color string) {
	js.Rewrite("$<.Color = $1", color)
}

// Size
func (*HTMLFontElement) Size() (size int) {
	js.Rewrite("$<.Size")
	return size
}

// Size
func (*HTMLFontElement) SetSize(size int) {
	js.Rewrite("$<.Size = $1", size)
}

// Face Sets or retrieves the current typeface family.
func (*HTMLFontElement) Face() (face string) {
	js.Rewrite("$<.Face")
	return face
}

// Face Sets or retrieves the current typeface family.
func (*HTMLFontElement) SetFace(face string) {
	js.Rewrite("$<.Face = $1", face)
}
