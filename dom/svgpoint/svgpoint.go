package svgpoint

import (
	"github.com/matthewmueller/joy/dom/svgmatrix"
	"github.com/matthewmueller/joy/js"
)

// SVGPoint struct
// js:"SVGPoint,omit"
type SVGPoint struct {
}

// MatrixTransform fn
// js:"matrixTransform"
func (*SVGPoint) MatrixTransform(matrix *svgmatrix.SVGMatrix) (s *SVGPoint) {
	js.Rewrite("$_.matrixTransform($1)", matrix)
	return s
}

// X prop
// js:"x"
func (*SVGPoint) X() (x float32) {
	js.Rewrite("$_.x")
	return x
}

// SetX prop
// js:"x"
func (*SVGPoint) SetX(x float32) {
	js.Rewrite("$_.x = $1", x)
}

// Y prop
// js:"y"
func (*SVGPoint) Y() (y float32) {
	js.Rewrite("$_.y")
	return y
}

// SetY prop
// js:"y"
func (*SVGPoint) SetY(y float32) {
	js.Rewrite("$_.y = $1", y)
}
