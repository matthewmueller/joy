package svgmatrix

import "github.com/matthewmueller/golly/js"

// SVGMatrix struct
// js:"SVGMatrix,omit"
type SVGMatrix struct {
}

// FlipX fn
func (*SVGMatrix) FlipX() (s *SVGMatrix) {
	js.Rewrite("$<.flipX()")
	return s
}

// FlipY fn
func (*SVGMatrix) FlipY() (s *SVGMatrix) {
	js.Rewrite("$<.flipY()")
	return s
}

// Inverse fn
func (*SVGMatrix) Inverse() (s *SVGMatrix) {
	js.Rewrite("$<.inverse()")
	return s
}

// Multiply fn
func (*SVGMatrix) Multiply(secondMatrix *SVGMatrix) (s *SVGMatrix) {
	js.Rewrite("$<.multiply($1)", secondMatrix)
	return s
}

// Rotate fn
func (*SVGMatrix) Rotate(angle float32) (s *SVGMatrix) {
	js.Rewrite("$<.rotate($1)", angle)
	return s
}

// RotateFromVector fn
func (*SVGMatrix) RotateFromVector(x float32, y float32) (s *SVGMatrix) {
	js.Rewrite("$<.rotateFromVector($1, $2)", x, y)
	return s
}

// Scale fn
func (*SVGMatrix) Scale(scaleFactor float32) (s *SVGMatrix) {
	js.Rewrite("$<.scale($1)", scaleFactor)
	return s
}

// ScaleNonUniform fn
func (*SVGMatrix) ScaleNonUniform(scaleFactorX float32, scaleFactorY float32) (s *SVGMatrix) {
	js.Rewrite("$<.scaleNonUniform($1, $2)", scaleFactorX, scaleFactorY)
	return s
}

// SkewX fn
func (*SVGMatrix) SkewX(angle float32) (s *SVGMatrix) {
	js.Rewrite("$<.skewX($1)", angle)
	return s
}

// SkewY fn
func (*SVGMatrix) SkewY(angle float32) (s *SVGMatrix) {
	js.Rewrite("$<.skewY($1)", angle)
	return s
}

// Translate fn
func (*SVGMatrix) Translate(x float32, y float32) (s *SVGMatrix) {
	js.Rewrite("$<.translate($1, $2)", x, y)
	return s
}

// A prop
func (*SVGMatrix) A() (a float32) {
	js.Rewrite("$<.a")
	return a
}

// A prop
func (*SVGMatrix) SetA(a float32) {
	js.Rewrite("$<.a = $1", a)
}

// B prop
func (*SVGMatrix) B() (b float32) {
	js.Rewrite("$<.b")
	return b
}

// B prop
func (*SVGMatrix) SetB(b float32) {
	js.Rewrite("$<.b = $1", b)
}

// C prop
func (*SVGMatrix) C() (c float32) {
	js.Rewrite("$<.c")
	return c
}

// C prop
func (*SVGMatrix) SetC(c float32) {
	js.Rewrite("$<.c = $1", c)
}

// D prop
func (*SVGMatrix) D() (d float32) {
	js.Rewrite("$<.d")
	return d
}

// D prop
func (*SVGMatrix) SetD(d float32) {
	js.Rewrite("$<.d = $1", d)
}

// E prop
func (*SVGMatrix) E() (e float32) {
	js.Rewrite("$<.e")
	return e
}

// E prop
func (*SVGMatrix) SetE(e float32) {
	js.Rewrite("$<.e = $1", e)
}

// F prop
func (*SVGMatrix) F() (f float32) {
	js.Rewrite("$<.f")
	return f
}

// F prop
func (*SVGMatrix) SetF(f float32) {
	js.Rewrite("$<.f = $1", f)
}
