package svgpathsegclosepath

import (
	"github.com/matthewmueller/joy/dom/svgpathseg"
	"github.com/matthewmueller/joy/macro"
)

var _ svgpathseg.SVGPathSeg = (*SVGPathSegClosePath)(nil)

// SVGPathSegClosePath struct
// js:"SVGPathSegClosePath,omit"
type SVGPathSegClosePath struct {
}

// PathSegType prop
// js:"pathSegType"
func (*SVGPathSegClosePath) PathSegType() (pathSegType uint8) {
	macro.Rewrite("$_.pathSegType")
	return pathSegType
}

// PathSegTypeAsLetter prop
// js:"pathSegTypeAsLetter"
func (*SVGPathSegClosePath) PathSegTypeAsLetter() (pathSegTypeAsLetter string) {
	macro.Rewrite("$_.pathSegTypeAsLetter")
	return pathSegTypeAsLetter
}
