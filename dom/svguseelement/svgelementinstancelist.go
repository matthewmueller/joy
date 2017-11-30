package svguseelement

import "github.com/matthewmueller/golly/js"

// SVGElementInstanceList struct
// js:"SVGElementInstanceList,omit"
type SVGElementInstanceList struct {
}

// Item fn
// js:"item"
func (*SVGElementInstanceList) Item(index uint) (s *SVGElementInstance) {
	js.Rewrite("$_.item($1)", index)
	return s
}

// Length prop
// js:"length"
func (*SVGElementInstanceList) Length() (length uint) {
	js.Rewrite("$_.length")
	return length
}
