package svgpathseg

// SVGPathSeg interface
// js:"SVGPathSeg"
type SVGPathSeg interface {

	// PathSegType prop
	// js:"pathSegType"
	PathSegType() (pathSegType uint8)

	// PathSegTypeAsLetter prop
	// js:"pathSegTypeAsLetter"
	PathSegTypeAsLetter() (pathSegTypeAsLetter string)
}
