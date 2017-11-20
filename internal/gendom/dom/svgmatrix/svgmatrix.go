package svgmatrix

import "github.com/matthewmueller/golly/js"

type SVGMatrix struct {
}

func (*SVGMatrix) FlipX() (s *SVGMatrix) {
	js.Rewrite("$<.flipX()")
	return s
}

func (*SVGMatrix) FlipY() (s *SVGMatrix) {
	js.Rewrite("$<.flipY()")
	return s
}

func (*SVGMatrix) Inverse() (s *SVGMatrix) {
	js.Rewrite("$<.inverse()")
	return s
}

func (*SVGMatrix) Multiply(secondMatrix *SVGMatrix) (s *SVGMatrix) {
	js.Rewrite("$<.multiply($1)", secondMatrix)
	return s
}

func (*SVGMatrix) Rotate(angle float32) (s *SVGMatrix) {
	js.Rewrite("$<.rotate($1)", angle)
	return s
}

func (*SVGMatrix) RotateFromVector(x float32, y float32) (s *SVGMatrix) {
	js.Rewrite("$<.rotateFromVector($1, $2)", x, y)
	return s
}

func (*SVGMatrix) Scale(scaleFactor float32) (s *SVGMatrix) {
	js.Rewrite("$<.scale($1)", scaleFactor)
	return s
}

func (*SVGMatrix) ScaleNonUniform(scaleFactorX float32, scaleFactorY float32) (s *SVGMatrix) {
	js.Rewrite("$<.scaleNonUniform($1, $2)", scaleFactorX, scaleFactorY)
	return s
}

func (*SVGMatrix) SkewX(angle float32) (s *SVGMatrix) {
	js.Rewrite("$<.skewX($1)", angle)
	return s
}

func (*SVGMatrix) SkewY(angle float32) (s *SVGMatrix) {
	js.Rewrite("$<.skewY($1)", angle)
	return s
}

func (*SVGMatrix) Translate(x float32, y float32) (s *SVGMatrix) {
	js.Rewrite("$<.translate($1, $2)", x, y)
	return s
}

func (*SVGMatrix) GetA() (a float32) {
	js.Rewrite("$<.a")
	return a
}

func (*SVGMatrix) SetA(a float32) {
	js.Rewrite("$<.a = $1", a)
}

func (*SVGMatrix) GetB() (b float32) {
	js.Rewrite("$<.b")
	return b
}

func (*SVGMatrix) SetB(b float32) {
	js.Rewrite("$<.b = $1", b)
}

func (*SVGMatrix) GetC() (c float32) {
	js.Rewrite("$<.c")
	return c
}

func (*SVGMatrix) SetC(c float32) {
	js.Rewrite("$<.c = $1", c)
}

func (*SVGMatrix) GetD() (d float32) {
	js.Rewrite("$<.d")
	return d
}

func (*SVGMatrix) SetD(d float32) {
	js.Rewrite("$<.d = $1", d)
}

func (*SVGMatrix) GetE() (e float32) {
	js.Rewrite("$<.e")
	return e
}

func (*SVGMatrix) SetE(e float32) {
	js.Rewrite("$<.e = $1", e)
}

func (*SVGMatrix) GetF() (f float32) {
	js.Rewrite("$<.f")
	return f
}

func (*SVGMatrix) SetF(f float32) {
	js.Rewrite("$<.f = $1", f)
}
