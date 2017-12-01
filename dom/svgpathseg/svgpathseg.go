package svgpathseg

// SVGPathSeg interface
// js:"SVGPathSeg"
type SVGPathSeg interface {

	// PathSegType prop
	// js:"pathSegType"
	// jsrewrite:"$_.pathSegType"
	PathSegType() (pathSegType uint8)

	// PathSegTypeAsLetter prop
	// js:"pathSegTypeAsLetter"
	// jsrewrite:"$_.pathSegTypeAsLetter"
	PathSegTypeAsLetter() (pathSegTypeAsLetter string)
}
