package htmlfontelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/doml2deprecatedcolorproperty"
	"github.com/matthewmueller/golly/internal/gendom/dom/doml2deprecatedsizeproperty"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/js"
)

type HTMLFontElement struct {
	*htmlelement.HTMLElement
	*doml2deprecatedcolorproperty.DOML2DeprecatedColorProperty
	*doml2deprecatedsizeproperty.DOML2DeprecatedSizeProperty
}

func (*HTMLFontElement) GetFace() (face string) {
	js.Rewrite("$<.face")
	return face
}

func (*HTMLFontElement) SetFace(face string) {
	js.Rewrite("$<.face = $1", face)
}
