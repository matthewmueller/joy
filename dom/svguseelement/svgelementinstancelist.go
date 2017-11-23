package svguseelement

import "github.com/matthewmueller/golly/js"

// SVGElementInstanceList struct
// js:"SVGElementInstanceList,omit"
type SVGElementInstanceList struct {
}

// Item fn
func (*SVGElementInstanceList) Item(index uint) (s *SVGElementInstance) {
	js.Rewrite("$<.item($1)", index)
	return s
}

// Length prop
func (*SVGElementInstanceList) Length() (length uint) {
	js.Rewrite("$<.length")
	return length
}
