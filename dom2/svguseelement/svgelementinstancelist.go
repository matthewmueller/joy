package svguseelement

import "github.com/matthewmueller/golly/js"

// SVGElementInstanceList struct
// js:"SVGElementInstanceList,omit"
type SVGElementInstanceList struct {
}

// Item
func (*SVGElementInstanceList) Item(index uint) (s *SVGElementInstance) {
	js.Rewrite("$<.Item($1)", index)
	return s
}

// Length
func (*SVGElementInstanceList) Length() (length uint) {
	js.Rewrite("$<.Length")
	return length
}
