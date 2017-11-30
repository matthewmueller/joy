package svgpathsegclosepath

import (
	"github.com/matthewmueller/golly/dom/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

var _ svgpathseg.SVGPathSeg = (*SVGPathSegClosePath)(nil)

// SVGPathSegClosePath struct
// js:"SVGPathSegClosePath,omit"
type SVGPathSegClosePath struct {
}

// PathSegType prop
// js:"pathSegType"
func (*SVGPathSegClosePath) PathSegType() (pathSegType uint8) {
	js.Rewrite("$_.pathSegType")
	return pathSegType
}

// PathSegTypeAsLetter prop
// js:"pathSegTypeAsLetter"
func (*SVGPathSegClosePath) PathSegTypeAsLetter() (pathSegTypeAsLetter string) {
	js.Rewrite("$_.pathSegTypeAsLetter")
	return pathSegTypeAsLetter
}
