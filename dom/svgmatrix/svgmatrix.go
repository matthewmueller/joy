package svgmatrix

import "github.com/matthewmueller/joy/js"

// SVGMatrix struct
// js:"SVGMatrix,omit"
type SVGMatrix struct {
}

// FlipX fn
// js:"flipX"
func (*SVGMatrix) FlipX() (s *SVGMatrix) {
	js.Rewrite("$_.flipX()")
	return s
}

// FlipY fn
// js:"flipY"
func (*SVGMatrix) FlipY() (s *SVGMatrix) {
	js.Rewrite("$_.flipY()")
	return s
}

// Inverse fn
// js:"inverse"
func (*SVGMatrix) Inverse() (s *SVGMatrix) {
	js.Rewrite("$_.inverse()")
	return s
}

// Multiply fn
// js:"multiply"
func (*SVGMatrix) Multiply(secondMatrix *SVGMatrix) (s *SVGMatrix) {
	js.Rewrite("$_.multiply($1)", secondMatrix)
	return s
}

// Rotate fn
// js:"rotate"
func (*SVGMatrix) Rotate(angle float32) (s *SVGMatrix) {
	js.Rewrite("$_.rotate($1)", angle)
	return s
}

// RotateFromVector fn
// js:"rotateFromVector"
func (*SVGMatrix) RotateFromVector(x float32, y float32) (s *SVGMatrix) {
	js.Rewrite("$_.rotateFromVector($1, $2)", x, y)
	return s
}

// Scale fn
// js:"scale"
func (*SVGMatrix) Scale(scaleFactor float32) (s *SVGMatrix) {
	js.Rewrite("$_.scale($1)", scaleFactor)
	return s
}

// ScaleNonUniform fn
// js:"scaleNonUniform"
func (*SVGMatrix) ScaleNonUniform(scaleFactorX float32, scaleFactorY float32) (s *SVGMatrix) {
	js.Rewrite("$_.scaleNonUniform($1, $2)", scaleFactorX, scaleFactorY)
	return s
}

// SkewX fn
// js:"skewX"
func (*SVGMatrix) SkewX(angle float32) (s *SVGMatrix) {
	js.Rewrite("$_.skewX($1)", angle)
	return s
}

// SkewY fn
// js:"skewY"
func (*SVGMatrix) SkewY(angle float32) (s *SVGMatrix) {
	js.Rewrite("$_.skewY($1)", angle)
	return s
}

// Translate fn
// js:"translate"
func (*SVGMatrix) Translate(x float32, y float32) (s *SVGMatrix) {
	js.Rewrite("$_.translate($1, $2)", x, y)
	return s
}

// A prop
// js:"a"
func (*SVGMatrix) A() (a float32) {
	js.Rewrite("$_.a")
	return a
}

// SetA prop
// js:"a"
func (*SVGMatrix) SetA(a float32) {
	js.Rewrite("$_.a = $1", a)
}

// B prop
// js:"b"
func (*SVGMatrix) B() (b float32) {
	js.Rewrite("$_.b")
	return b
}

// SetB prop
// js:"b"
func (*SVGMatrix) SetB(b float32) {
	js.Rewrite("$_.b = $1", b)
}

// C prop
// js:"c"
func (*SVGMatrix) C() (c float32) {
	js.Rewrite("$_.c")
	return c
}

// SetC prop
// js:"c"
func (*SVGMatrix) SetC(c float32) {
	js.Rewrite("$_.c = $1", c)
}

// D prop
// js:"d"
func (*SVGMatrix) D() (d float32) {
	js.Rewrite("$_.d")
	return d
}

// SetD prop
// js:"d"
func (*SVGMatrix) SetD(d float32) {
	js.Rewrite("$_.d = $1", d)
}

// E prop
// js:"e"
func (*SVGMatrix) E() (e float32) {
	js.Rewrite("$_.e")
	return e
}

// SetE prop
// js:"e"
func (*SVGMatrix) SetE(e float32) {
	js.Rewrite("$_.e = $1", e)
}

// F prop
// js:"f"
func (*SVGMatrix) F() (f float32) {
	js.Rewrite("$_.f")
	return f
}

// SetF prop
// js:"f"
func (*SVGMatrix) SetF(f float32) {
	js.Rewrite("$_.f = $1", f)
}
