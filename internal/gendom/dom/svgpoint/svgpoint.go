package svgpoint

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svgmatrix"
	"github.com/matthewmueller/golly/js"
)

type SVGPoint struct {
}

func (*SVGPoint) MatrixTransform(matrix *svgmatrix.SVGMatrix) (s *SVGPoint) {
	js.Rewrite("$<.matrixTransform($1)", matrix)
	return s
}

func (*SVGPoint) GetX() (x float32) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGPoint) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

func (*SVGPoint) GetY() (y float32) {
	js.Rewrite("$<.y")
	return y
}

func (*SVGPoint) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}
