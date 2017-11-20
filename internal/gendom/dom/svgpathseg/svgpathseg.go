package svgpathseg

import "github.com/matthewmueller/golly/js"

type SVGPathSeg struct {
}

func (*SVGPathSeg) GetPathSegType() (pathSegType uint8) {
	js.Rewrite("$<.pathSegType")
	return pathSegType
}

func (*SVGPathSeg) GetPathSegTypeAsLetter() (pathSegTypeAsLetter string) {
	js.Rewrite("$<.pathSegTypeAsLetter")
	return pathSegTypeAsLetter
}
