package svgpathseglinetoverticalabs

import (
	"github.com/matthewmueller/golly/dom/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

var _ svgpathseg.SVGPathSeg = (*SVGPathSegLinetoVerticalAbs)(nil)

// SVGPathSegLinetoVerticalAbs struct
// js:"SVGPathSegLinetoVerticalAbs,omit"
type SVGPathSegLinetoVerticalAbs struct {
}

// Y prop
// js:"y"
func (*SVGPathSegLinetoVerticalAbs) Y() (y float32) {
	js.Rewrite("$_.y")
	return y
}

// SetY prop
// js:"y"
func (*SVGPathSegLinetoVerticalAbs) SetY(y float32) {
	js.Rewrite("$_.y = $1", y)
}

// PathSegType prop
// js:"pathSegType"
func (*SVGPathSegLinetoVerticalAbs) PathSegType() (pathSegType uint8) {
	js.Rewrite("$_.pathSegType")
	return pathSegType
}

// PathSegTypeAsLetter prop
// js:"pathSegTypeAsLetter"
func (*SVGPathSegLinetoVerticalAbs) PathSegTypeAsLetter() (pathSegTypeAsLetter string) {
	js.Rewrite("$_.pathSegTypeAsLetter")
	return pathSegTypeAsLetter
}
