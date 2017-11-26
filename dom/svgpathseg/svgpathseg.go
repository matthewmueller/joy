package svgpathseg

import "github.com/matthewmueller/golly/js"

// js:"SVGPathSeg,omit"
type SVGPathSeg interface {

	// PathSegType prop
	// js:"pathSegType,rewrite=pathsegtype"
	PathSegType() (pathSegType uint8)

	// PathSegTypeAsLetter prop
	// js:"pathSegTypeAsLetter,rewrite=pathsegtypeasletter"
	PathSegTypeAsLetter() (pathSegTypeAsLetter string)
}

// pathsegtype prop
func pathsegtype() (pathSegType uint8) {
	js.Rewrite("$<.pathSegType")
	return pathSegType
}

// pathsegtypeasletter prop
func pathsegtypeasletter() (pathSegTypeAsLetter string) {
	js.Rewrite("$<.pathSegTypeAsLetter")
	return pathSegTypeAsLetter
}
