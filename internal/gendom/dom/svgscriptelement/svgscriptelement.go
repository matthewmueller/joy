package svgscriptelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svgelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgurireference"
	"github.com/matthewmueller/golly/js"
)

type SVGScriptElement struct {
	*svgelement.SVGElement
	*svgurireference.SVGURIReference
}

func (*SVGScriptElement) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

func (*SVGScriptElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}
