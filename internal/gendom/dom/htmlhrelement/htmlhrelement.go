package htmlhrelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/doml2deprecatedcolorproperty"
	"github.com/matthewmueller/golly/internal/gendom/dom/doml2deprecatedsizeproperty"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/js"
)

type HTMLHRElement struct {
	*htmlelement.HTMLElement
	*doml2deprecatedcolorproperty.DOML2DeprecatedColorProperty
	*doml2deprecatedsizeproperty.DOML2DeprecatedSizeProperty
}

func (*HTMLHRElement) GetAlign() (align string) {
	js.Rewrite("$<.align")
	return align
}

func (*HTMLHRElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

func (*HTMLHRElement) GetNoShade() (noShade bool) {
	js.Rewrite("$<.noShade")
	return noShade
}

func (*HTMLHRElement) SetNoShade(noShade bool) {
	js.Rewrite("$<.noShade = $1", noShade)
}

func (*HTMLHRElement) GetWidth() (width int) {
	js.Rewrite("$<.width")
	return width
}

func (*HTMLHRElement) SetWidth(width int) {
	js.Rewrite("$<.width = $1", width)
}
