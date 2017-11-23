package svgpoint

import (
	"github.com/matthewmueller/golly/dom2/svgmatrix"
	"github.com/matthewmueller/golly/js"
)

// js:"SVGPoint,omit"
type SVGPoint struct {
}

// MatrixTransform
func (*SVGPoint) MatrixTransform(matrix *svgmatrix.SVGMatrix) (s *SVGPoint) {
	js.Rewrite("$<.MatrixTransform($1)", matrix)
	return s
}

// X
func (*SVGPoint) X() (x float32) {
	js.Rewrite("$<.X")
	return x
}

// X
func (*SVGPoint) SetX(x float32) {
	js.Rewrite("$<.X = $1", x)
}

// Y
func (*SVGPoint) Y() (y float32) {
	js.Rewrite("$<.Y")
	return y
}

// Y
func (*SVGPoint) SetY(y float32) {
	js.Rewrite("$<.Y = $1", y)
}
