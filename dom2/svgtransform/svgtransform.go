package svgtransform

import (
	"github.com/matthewmueller/golly/dom2/svgmatrix"
	"github.com/matthewmueller/golly/js"
)

// js:"SVGTransform,omit"
type SVGTransform struct {
}

// SetMatrix
func (*SVGTransform) SetMatrix(matrix *svgmatrix.SVGMatrix) {
	js.Rewrite("$<.SetMatrix($1)", matrix)
}

// SetRotate
func (*SVGTransform) SetRotate(angle float32, cx float32, cy float32) {
	js.Rewrite("$<.SetRotate($1, $2, $3)", angle, cx, cy)
}

// SetScale
func (*SVGTransform) SetScale(sx float32, sy float32) {
	js.Rewrite("$<.SetScale($1, $2)", sx, sy)
}

// SetSkewX
func (*SVGTransform) SetSkewX(angle float32) {
	js.Rewrite("$<.SetSkewX($1)", angle)
}

// SetSkewY
func (*SVGTransform) SetSkewY(angle float32) {
	js.Rewrite("$<.SetSkewY($1)", angle)
}

// SetTranslate
func (*SVGTransform) SetTranslate(tx float32, ty float32) {
	js.Rewrite("$<.SetTranslate($1, $2)", tx, ty)
}

// Angle
func (*SVGTransform) Angle() (angle float32) {
	js.Rewrite("$<.Angle")
	return angle
}

// Matrix
func (*SVGTransform) Matrix() (matrix *svgmatrix.SVGMatrix) {
	js.Rewrite("$<.Matrix")
	return matrix
}

// Type
func (*SVGTransform) Type() (kind uint8) {
	js.Rewrite("$<.Type")
	return kind
}
