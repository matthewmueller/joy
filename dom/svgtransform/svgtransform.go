package svgtransform

import (
	"github.com/matthewmueller/golly/dom/svgmatrix"
	"github.com/matthewmueller/golly/js"
)

// SVGTransform struct
// js:"SVGTransform,omit"
type SVGTransform struct {
}

// SetMatrix fn
func (*SVGTransform) SetMatrix(matrix *svgmatrix.SVGMatrix) {
	js.Rewrite("$<.setMatrix($1)", matrix)
}

// SetRotate fn
func (*SVGTransform) SetRotate(angle float32, cx float32, cy float32) {
	js.Rewrite("$<.setRotate($1, $2, $3)", angle, cx, cy)
}

// SetScale fn
func (*SVGTransform) SetScale(sx float32, sy float32) {
	js.Rewrite("$<.setScale($1, $2)", sx, sy)
}

// SetSkewX fn
func (*SVGTransform) SetSkewX(angle float32) {
	js.Rewrite("$<.setSkewX($1)", angle)
}

// SetSkewY fn
func (*SVGTransform) SetSkewY(angle float32) {
	js.Rewrite("$<.setSkewY($1)", angle)
}

// SetTranslate fn
func (*SVGTransform) SetTranslate(tx float32, ty float32) {
	js.Rewrite("$<.setTranslate($1, $2)", tx, ty)
}

// Angle prop
func (*SVGTransform) Angle() (angle float32) {
	js.Rewrite("$<.angle")
	return angle
}

// Matrix prop
func (*SVGTransform) Matrix() (matrix *svgmatrix.SVGMatrix) {
	js.Rewrite("$<.matrix")
	return matrix
}

// Type prop
func (*SVGTransform) Type() (kind uint8) {
	js.Rewrite("$<.type")
	return kind
}
