package webkitcssmatrix

import "github.com/matthewmueller/golly/js"

type WebKitCSSMatrix struct {
}

func (*WebKitCSSMatrix) Inverse() (w *WebKitCSSMatrix) {
	js.Rewrite("$<.inverse()")
	return w
}

func (*WebKitCSSMatrix) Multiply(secondMatrix *WebKitCSSMatrix) (w *WebKitCSSMatrix) {
	js.Rewrite("$<.multiply($1)", secondMatrix)
	return w
}

func (*WebKitCSSMatrix) Rotate(angleX float32, angleY float32, angleZ float32) (w *WebKitCSSMatrix) {
	js.Rewrite("$<.rotate($1, $2, $3)", angleX, angleY, angleZ)
	return w
}

func (*WebKitCSSMatrix) RotateAxisAngle(x float32, y float32, z float32, angle float32) (w *WebKitCSSMatrix) {
	js.Rewrite("$<.rotateAxisAngle($1, $2, $3, $4)", x, y, z, angle)
	return w
}

func (*WebKitCSSMatrix) Scale(scaleX float32, scaleY float32, scaleZ float32) (w *WebKitCSSMatrix) {
	js.Rewrite("$<.scale($1, $2, $3)", scaleX, scaleY, scaleZ)
	return w
}

func (*WebKitCSSMatrix) SetMatrixValue(value string) {
	js.Rewrite("$<.setMatrixValue($1)", value)
}

func (*WebKitCSSMatrix) SkewX(angle float32) (w *WebKitCSSMatrix) {
	js.Rewrite("$<.skewX($1)", angle)
	return w
}

func (*WebKitCSSMatrix) SkewY(angle float32) (w *WebKitCSSMatrix) {
	js.Rewrite("$<.skewY($1)", angle)
	return w
}

func (*WebKitCSSMatrix) ToString() (s string) {
	js.Rewrite("$<.toString()")
	return s
}

func (*WebKitCSSMatrix) Translate(x float32, y float32, z float32) (w *WebKitCSSMatrix) {
	js.Rewrite("$<.translate($1, $2, $3)", x, y, z)
	return w
}

func (*WebKitCSSMatrix) GetA() (a float32) {
	js.Rewrite("$<.a")
	return a
}

func (*WebKitCSSMatrix) SetA(a float32) {
	js.Rewrite("$<.a = $1", a)
}

func (*WebKitCSSMatrix) GetB() (b float32) {
	js.Rewrite("$<.b")
	return b
}

func (*WebKitCSSMatrix) SetB(b float32) {
	js.Rewrite("$<.b = $1", b)
}

func (*WebKitCSSMatrix) GetC() (c float32) {
	js.Rewrite("$<.c")
	return c
}

func (*WebKitCSSMatrix) SetC(c float32) {
	js.Rewrite("$<.c = $1", c)
}

func (*WebKitCSSMatrix) GetD() (d float32) {
	js.Rewrite("$<.d")
	return d
}

func (*WebKitCSSMatrix) SetD(d float32) {
	js.Rewrite("$<.d = $1", d)
}

func (*WebKitCSSMatrix) GetE() (e float32) {
	js.Rewrite("$<.e")
	return e
}

func (*WebKitCSSMatrix) SetE(e float32) {
	js.Rewrite("$<.e = $1", e)
}

func (*WebKitCSSMatrix) GetF() (f float32) {
	js.Rewrite("$<.f")
	return f
}

func (*WebKitCSSMatrix) SetF(f float32) {
	js.Rewrite("$<.f = $1", f)
}

func (*WebKitCSSMatrix) GetM11() (m11 float32) {
	js.Rewrite("$<.m11")
	return m11
}

func (*WebKitCSSMatrix) SetM11(m11 float32) {
	js.Rewrite("$<.m11 = $1", m11)
}

func (*WebKitCSSMatrix) GetM12() (m12 float32) {
	js.Rewrite("$<.m12")
	return m12
}

func (*WebKitCSSMatrix) SetM12(m12 float32) {
	js.Rewrite("$<.m12 = $1", m12)
}

func (*WebKitCSSMatrix) GetM13() (m13 float32) {
	js.Rewrite("$<.m13")
	return m13
}

func (*WebKitCSSMatrix) SetM13(m13 float32) {
	js.Rewrite("$<.m13 = $1", m13)
}

func (*WebKitCSSMatrix) GetM14() (m14 float32) {
	js.Rewrite("$<.m14")
	return m14
}

func (*WebKitCSSMatrix) SetM14(m14 float32) {
	js.Rewrite("$<.m14 = $1", m14)
}

func (*WebKitCSSMatrix) GetM21() (m21 float32) {
	js.Rewrite("$<.m21")
	return m21
}

func (*WebKitCSSMatrix) SetM21(m21 float32) {
	js.Rewrite("$<.m21 = $1", m21)
}

func (*WebKitCSSMatrix) GetM22() (m22 float32) {
	js.Rewrite("$<.m22")
	return m22
}

func (*WebKitCSSMatrix) SetM22(m22 float32) {
	js.Rewrite("$<.m22 = $1", m22)
}

func (*WebKitCSSMatrix) GetM23() (m23 float32) {
	js.Rewrite("$<.m23")
	return m23
}

func (*WebKitCSSMatrix) SetM23(m23 float32) {
	js.Rewrite("$<.m23 = $1", m23)
}

func (*WebKitCSSMatrix) GetM24() (m24 float32) {
	js.Rewrite("$<.m24")
	return m24
}

func (*WebKitCSSMatrix) SetM24(m24 float32) {
	js.Rewrite("$<.m24 = $1", m24)
}

func (*WebKitCSSMatrix) GetM31() (m31 float32) {
	js.Rewrite("$<.m31")
	return m31
}

func (*WebKitCSSMatrix) SetM31(m31 float32) {
	js.Rewrite("$<.m31 = $1", m31)
}

func (*WebKitCSSMatrix) GetM32() (m32 float32) {
	js.Rewrite("$<.m32")
	return m32
}

func (*WebKitCSSMatrix) SetM32(m32 float32) {
	js.Rewrite("$<.m32 = $1", m32)
}

func (*WebKitCSSMatrix) GetM33() (m33 float32) {
	js.Rewrite("$<.m33")
	return m33
}

func (*WebKitCSSMatrix) SetM33(m33 float32) {
	js.Rewrite("$<.m33 = $1", m33)
}

func (*WebKitCSSMatrix) GetM34() (m34 float32) {
	js.Rewrite("$<.m34")
	return m34
}

func (*WebKitCSSMatrix) SetM34(m34 float32) {
	js.Rewrite("$<.m34 = $1", m34)
}

func (*WebKitCSSMatrix) GetM41() (m41 float32) {
	js.Rewrite("$<.m41")
	return m41
}

func (*WebKitCSSMatrix) SetM41(m41 float32) {
	js.Rewrite("$<.m41 = $1", m41)
}

func (*WebKitCSSMatrix) GetM42() (m42 float32) {
	js.Rewrite("$<.m42")
	return m42
}

func (*WebKitCSSMatrix) SetM42(m42 float32) {
	js.Rewrite("$<.m42 = $1", m42)
}

func (*WebKitCSSMatrix) GetM43() (m43 float32) {
	js.Rewrite("$<.m43")
	return m43
}

func (*WebKitCSSMatrix) SetM43(m43 float32) {
	js.Rewrite("$<.m43 = $1", m43)
}

func (*WebKitCSSMatrix) GetM44() (m44 float32) {
	js.Rewrite("$<.m44")
	return m44
}

func (*WebKitCSSMatrix) SetM44(m44 float32) {
	js.Rewrite("$<.m44 = $1", m44)
}
