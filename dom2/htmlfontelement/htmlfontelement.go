package htmlfontelement

import (
	"github.com/matthewmueller/golly/dom2/doml2deprecatedcolorproperty"
	"github.com/matthewmueller/golly/dom2/doml2deprecatedsizeproperty"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"HTMLFontElement,omit"
type HTMLFontElement struct {
	window.HTMLElement
	doml2deprecatedcolorproperty.DOML2deprecatedColorProperty
	doml2deprecatedsizeproperty.DOML2deprecatedSizeProperty
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
