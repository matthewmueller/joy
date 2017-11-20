package htmlbasefontelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/doml2deprecatedcolorproperty"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/js"
)

type HTMLBaseFontElement struct {
	*htmlelement.HTMLElement
	*doml2deprecatedcolorproperty.DOML2DeprecatedColorProperty
}

func (*HTMLBaseFontElement) GetFace() (face string) {
	js.Rewrite("$<.face")
	return face
}

func (*HTMLBaseFontElement) SetFace(face string) {
	js.Rewrite("$<.face = $1", face)
}

func (*HTMLBaseFontElement) GetSize() (size int) {
	js.Rewrite("$<.size")
	return size
}

func (*HTMLBaseFontElement) SetSize(size int) {
	js.Rewrite("$<.size = $1", size)
}
