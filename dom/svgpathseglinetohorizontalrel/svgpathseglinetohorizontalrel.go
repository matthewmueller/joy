package svgpathseglinetohorizontalrel

import (
	"github.com/matthewmueller/joy/dom/svgpathseg"
	"github.com/matthewmueller/joy/js"
)

var _ svgpathseg.SVGPathSeg = (*SVGPathSegLinetoHorizontalRel)(nil)

// SVGPathSegLinetoHorizontalRel struct
// js:"SVGPathSegLinetoHorizontalRel,omit"
type SVGPathSegLinetoHorizontalRel struct {
}

// X prop
// js:"x"
func (*SVGPathSegLinetoHorizontalRel) X() (x float32) {
	js.Rewrite("$_.x")
	return x
}

// SetX prop
// js:"x"
func (*SVGPathSegLinetoHorizontalRel) SetX(x float32) {
	js.Rewrite("$_.x = $1", x)
}

// PathSegType prop
// js:"pathSegType"
func (*SVGPathSegLinetoHorizontalRel) PathSegType() (pathSegType uint8) {
	js.Rewrite("$_.pathSegType")
	return pathSegType
}

// PathSegTypeAsLetter prop
// js:"pathSegTypeAsLetter"
func (*SVGPathSegLinetoHorizontalRel) PathSegTypeAsLetter() (pathSegTypeAsLetter string) {
	js.Rewrite("$_.pathSegTypeAsLetter")
	return pathSegTypeAsLetter
}
