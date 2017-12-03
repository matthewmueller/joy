package svguseelement

import "github.com/matthewmueller/joy/macro"

// SVGElementInstanceList struct
// js:"SVGElementInstanceList,omit"
type SVGElementInstanceList struct {
}

// Item fn
// js:"item"
func (*SVGElementInstanceList) Item(index uint) (s *SVGElementInstance) {
	macro.Rewrite("$_.item($1)", index)
	return s
}

// Length prop
// js:"length"
func (*SVGElementInstanceList) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}
