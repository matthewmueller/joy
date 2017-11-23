package htmlhrelement

import (
	"github.com/matthewmueller/golly/dom2/doml2deprecatedcolorproperty"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"HTMLHRElement,omit"
type HTMLHRElement struct {
	window.HTMLElement
	doml2deprecatedcolorproperty.DOML2deprecatedColorProperty
	doml2deprecatedsizeproperty.DOML2deprecatedSizeProperty
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
