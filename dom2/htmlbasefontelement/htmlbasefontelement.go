package htmlbasefontelement

import (
	"github.com/matthewmueller/golly/dom2/doml2deprecatedcolorproperty"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"HTMLBaseFontElement,omit"
type HTMLBaseFontElement struct {
	window.HTMLElement
	doml2deprecatedcolorproperty.DOML2deprecatedColorProperty
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
