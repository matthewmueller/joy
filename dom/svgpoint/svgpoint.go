package svgpoint

import (
	"github.com/matthewmueller/golly/dom2/svgmatrix"
	"github.com/matthewmueller/golly/js"
)

// SVGPoint struct
// js:"SVGPoint,omit"
type SVGPoint struct {
}

// MatrixTransform fn
func (*SVGPoint) MatrixTransform(matrix *svgmatrix.SVGMatrix) (s *SVGPoint) {
	js.Rewrite("$<.matrixTransform($1)", matrix)
	return s
}

// X prop
func (*SVGPoint) X() (x float32) {
	js.Rewrite("$<.x")
	return x
}

// X prop
func (*SVGPoint) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

// Y prop
func (*SVGPoint) Y() (y float32) {
	js.Rewrite("$<.y")
	return y
}

// Y prop
func (*SVGPoint) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}
