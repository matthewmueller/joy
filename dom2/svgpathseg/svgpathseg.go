package svgpathseg

// js:"SVGPathSeg,omit"
type SVGPathSeg interface {

	// PathSegType
	PathSegType() (pathSegType uint8)

	// PathSegTypeAsLetter
	PathSegTypeAsLetter() (pathSegTypeAsLetter string)
}
