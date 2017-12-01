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
// js:"setMatrix"
func (*SVGTransform) SetMatrix(matrix *svgmatrix.SVGMatrix) {
	js.Rewrite("$_.setMatrix($1)", matrix)
}

// SetRotate fn
// js:"setRotate"
func (*SVGTransform) SetRotate(angle float32, cx float32, cy float32) {
	js.Rewrite("$_.setRotate($1, $2, $3)", angle, cx, cy)
}

// SetScale fn
// js:"setScale"
func (*SVGTransform) SetScale(sx float32, sy float32) {
	js.Rewrite("$_.setScale($1, $2)", sx, sy)
}

// SetSkewX fn
// js:"setSkewX"
func (*SVGTransform) SetSkewX(angle float32) {
	js.Rewrite("$_.setSkewX($1)", angle)
}

// SetSkewY fn
// js:"setSkewY"
func (*SVGTransform) SetSkewY(angle float32) {
	js.Rewrite("$_.setSkewY($1)", angle)
}

// SetTranslate fn
// js:"setTranslate"
func (*SVGTransform) SetTranslate(tx float32, ty float32) {
	js.Rewrite("$_.setTranslate($1, $2)", tx, ty)
}

// Angle prop
// js:"angle"
func (*SVGTransform) Angle() (angle float32) {
	js.Rewrite("$_.angle")
	return angle
}

// Matrix prop
// js:"matrix"
func (*SVGTransform) Matrix() (matrix *svgmatrix.SVGMatrix) {
	js.Rewrite("$_.matrix")
	return matrix
}

// Type prop
// js:"type"
func (*SVGTransform) Type() (kind uint8) {
	js.Rewrite("$_.type")
	return kind
}
