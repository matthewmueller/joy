package svgaelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedstring"
	"github.com/matthewmueller/golly/internal/gendom/dom/svggraphicselement"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgurireference"
	"github.com/matthewmueller/golly/js"
)

type SVGAElement struct {
	*svggraphicselement.SVGGraphicsElement
	*svgurireference.SVGURIReference
}

func (*SVGAElement) GetTarget() (target *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.target")
	return target
}
