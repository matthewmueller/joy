package svgpathseglinetoverticalrel

import (
	"github.com/matthewmueller/joy/dom/svgpathseg"
	"github.com/matthewmueller/joy/macro"
)

var _ svgpathseg.SVGPathSeg = (*SVGPathSegLinetoVerticalRel)(nil)

// SVGPathSegLinetoVerticalRel struct
// js:"SVGPathSegLinetoVerticalRel,omit"
type SVGPathSegLinetoVerticalRel struct {
}

// Y prop
// js:"y"
func (*SVGPathSegLinetoVerticalRel) Y() (y float32) {
	macro.Rewrite("$_.y")
	return y
}

// SetY prop
// js:"y"
func (*SVGPathSegLinetoVerticalRel) SetY(y float32) {
	macro.Rewrite("$_.y = $1", y)
}

// PathSegType prop
// js:"pathSegType"
func (*SVGPathSegLinetoVerticalRel) PathSegType() (pathSegType uint8) {
	macro.Rewrite("$_.pathSegType")
	return pathSegType
}

// PathSegTypeAsLetter prop
// js:"pathSegTypeAsLetter"
func (*SVGPathSegLinetoVerticalRel) PathSegTypeAsLetter() (pathSegTypeAsLetter string) {
	macro.Rewrite("$_.pathSegTypeAsLetter")
	return pathSegTypeAsLetter
}
