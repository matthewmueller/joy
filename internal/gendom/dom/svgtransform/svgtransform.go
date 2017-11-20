package svgtransform

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svgmatrix"
	"github.com/matthewmueller/golly/js"
)

type SVGTransform struct {
}

func (*SVGTransform) SetMatrix(matrix *svgmatrix.SVGMatrix) {
	js.Rewrite("$<.setMatrix($1)", matrix)
}

func (*SVGTransform) SetRotate(angle float32, cx float32, cy float32) {
	js.Rewrite("$<.setRotate($1, $2, $3)", angle, cx, cy)
}

func (*SVGTransform) SetScale(sx float32, sy float32) {
	js.Rewrite("$<.setScale($1, $2)", sx, sy)
}

func (*SVGTransform) SetSkewX(angle float32) {
	js.Rewrite("$<.setSkewX($1)", angle)
}

func (*SVGTransform) SetSkewY(angle float32) {
	js.Rewrite("$<.setSkewY($1)", angle)
}

func (*SVGTransform) SetTranslate(tx float32, ty float32) {
	js.Rewrite("$<.setTranslate($1, $2)", tx, ty)
}

func (*SVGTransform) GetAngle() (angle float32) {
	js.Rewrite("$<.angle")
	return angle
}

func (*SVGTransform) GetMatrix() (matrix *svgmatrix.SVGMatrix) {
	js.Rewrite("$<.matrix")
	return matrix
}

func (*SVGTransform) GetType() (kind uint8) {
	js.Rewrite("$<.type")
	return kind
}
