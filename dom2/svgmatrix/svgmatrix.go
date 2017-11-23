package svgmatrix

import "github.com/matthewmueller/golly/js"

// SVGMatrix struct
// js:"SVGMatrix,omit"
type SVGMatrix struct {
}

// FlipX
func (*SVGMatrix) FlipX() (s *SVGMatrix) {
	js.Rewrite("$<.FlipX()")
	return s
}

// FlipY
func (*SVGMatrix) FlipY() (s *SVGMatrix) {
	js.Rewrite("$<.FlipY()")
	return s
}

// Inverse
func (*SVGMatrix) Inverse() (s *SVGMatrix) {
	js.Rewrite("$<.Inverse()")
	return s
}

// Multiply
func (*SVGMatrix) Multiply(secondMatrix *SVGMatrix) (s *SVGMatrix) {
	js.Rewrite("$<.Multiply($1)", secondMatrix)
	return s
}

// Rotate
func (*SVGMatrix) Rotate(angle float32) (s *SVGMatrix) {
	js.Rewrite("$<.Rotate($1)", angle)
	return s
}

// RotateFromVector
func (*SVGMatrix) RotateFromVector(x float32, y float32) (s *SVGMatrix) {
	js.Rewrite("$<.RotateFromVector($1, $2)", x, y)
	return s
}

// Scale
func (*SVGMatrix) Scale(scaleFactor float32) (s *SVGMatrix) {
	js.Rewrite("$<.Scale($1)", scaleFactor)
	return s
}

// ScaleNonUniform
func (*SVGMatrix) ScaleNonUniform(scaleFactorX float32, scaleFactorY float32) (s *SVGMatrix) {
	js.Rewrite("$<.ScaleNonUniform($1, $2)", scaleFactorX, scaleFactorY)
	return s
}

// SkewX
func (*SVGMatrix) SkewX(angle float32) (s *SVGMatrix) {
	js.Rewrite("$<.SkewX($1)", angle)
	return s
}

// SkewY
func (*SVGMatrix) SkewY(angle float32) (s *SVGMatrix) {
	js.Rewrite("$<.SkewY($1)", angle)
	return s
}

// Translate
func (*SVGMatrix) Translate(x float32, y float32) (s *SVGMatrix) {
	js.Rewrite("$<.Translate($1, $2)", x, y)
	return s
}

// A
func (*SVGMatrix) A() (a float32) {
	js.Rewrite("$<.A")
	return a
}

// A
func (*SVGMatrix) SetA(a float32) {
	js.Rewrite("$<.A = $1", a)
}

// B
func (*SVGMatrix) B() (b float32) {
	js.Rewrite("$<.B")
	return b
}

// B
func (*SVGMatrix) SetB(b float32) {
	js.Rewrite("$<.B = $1", b)
}

// C
func (*SVGMatrix) C() (c float32) {
	js.Rewrite("$<.C")
	return c
}

// C
func (*SVGMatrix) SetC(c float32) {
	js.Rewrite("$<.C = $1", c)
}

// D
func (*SVGMatrix) D() (d float32) {
	js.Rewrite("$<.D")
	return d
}

// D
func (*SVGMatrix) SetD(d float32) {
	js.Rewrite("$<.D = $1", d)
}

// E
func (*SVGMatrix) E() (e float32) {
	js.Rewrite("$<.E")
	return e
}

// E
func (*SVGMatrix) SetE(e float32) {
	js.Rewrite("$<.E = $1", e)
}

// F
func (*SVGMatrix) F() (f float32) {
	js.Rewrite("$<.F")
	return f
}

// F
func (*SVGMatrix) SetF(f float32) {
	js.Rewrite("$<.F = $1", f)
}
