package webkitcssmatrix

import "github.com/matthewmueller/golly/js"

// New fn
func New(text *string) *WebKitCSSMatrix {
	js.Rewrite("WebKitCSSMatrix")
	return &WebKitCSSMatrix{}
}

// WebKitCSSMatrix struct
// js:"WebKitCSSMatrix,omit"
type WebKitCSSMatrix struct {
}

// Inverse fn
func (*WebKitCSSMatrix) Inverse() (w *WebKitCSSMatrix) {
	js.Rewrite("$<.inverse()")
	return w
}

// Multiply fn
func (*WebKitCSSMatrix) Multiply(secondMatrix *WebKitCSSMatrix) (w *WebKitCSSMatrix) {
	js.Rewrite("$<.multiply($1)", secondMatrix)
	return w
}

// Rotate fn
func (*WebKitCSSMatrix) Rotate(angleX float32, angleY *float32, angleZ *float32) (w *WebKitCSSMatrix) {
	js.Rewrite("$<.rotate($1, $2, $3)", angleX, angleY, angleZ)
	return w
}

// RotateAxisAngle fn
func (*WebKitCSSMatrix) RotateAxisAngle(x float32, y float32, z float32, angle float32) (w *WebKitCSSMatrix) {
	js.Rewrite("$<.rotateAxisAngle($1, $2, $3, $4)", x, y, z, angle)
	return w
}

// Scale fn
func (*WebKitCSSMatrix) Scale(scaleX float32, scaleY *float32, scaleZ *float32) (w *WebKitCSSMatrix) {
	js.Rewrite("$<.scale($1, $2, $3)", scaleX, scaleY, scaleZ)
	return w
}

// SetMatrixValue fn
func (*WebKitCSSMatrix) SetMatrixValue(value string) {
	js.Rewrite("$<.setMatrixValue($1)", value)
}

// SkewX fn
func (*WebKitCSSMatrix) SkewX(angle float32) (w *WebKitCSSMatrix) {
	js.Rewrite("$<.skewX($1)", angle)
	return w
}

// SkewY fn
func (*WebKitCSSMatrix) SkewY(angle float32) (w *WebKitCSSMatrix) {
	js.Rewrite("$<.skewY($1)", angle)
	return w
}

// ToString fn
func (*WebKitCSSMatrix) ToString() (s string) {
	js.Rewrite("$<.toString()")
	return s
}

// Translate fn
func (*WebKitCSSMatrix) Translate(x float32, y float32, z *float32) (w *WebKitCSSMatrix) {
	js.Rewrite("$<.translate($1, $2, $3)", x, y, z)
	return w
}

// A prop
func (*WebKitCSSMatrix) A() (a float32) {
	js.Rewrite("$<.a")
	return a
}

// A prop
func (*WebKitCSSMatrix) SetA(a float32) {
	js.Rewrite("$<.a = $1", a)
}

// B prop
func (*WebKitCSSMatrix) B() (b float32) {
	js.Rewrite("$<.b")
	return b
}

// B prop
func (*WebKitCSSMatrix) SetB(b float32) {
	js.Rewrite("$<.b = $1", b)
}

// C prop
func (*WebKitCSSMatrix) C() (c float32) {
	js.Rewrite("$<.c")
	return c
}

// C prop
func (*WebKitCSSMatrix) SetC(c float32) {
	js.Rewrite("$<.c = $1", c)
}

// D prop
func (*WebKitCSSMatrix) D() (d float32) {
	js.Rewrite("$<.d")
	return d
}

// D prop
func (*WebKitCSSMatrix) SetD(d float32) {
	js.Rewrite("$<.d = $1", d)
}

// E prop
func (*WebKitCSSMatrix) E() (e float32) {
	js.Rewrite("$<.e")
	return e
}

// E prop
func (*WebKitCSSMatrix) SetE(e float32) {
	js.Rewrite("$<.e = $1", e)
}

// F prop
func (*WebKitCSSMatrix) F() (f float32) {
	js.Rewrite("$<.f")
	return f
}

// F prop
func (*WebKitCSSMatrix) SetF(f float32) {
	js.Rewrite("$<.f = $1", f)
}

// M11 prop
func (*WebKitCSSMatrix) M11() (m11 float32) {
	js.Rewrite("$<.m11")
	return m11
}

// M11 prop
func (*WebKitCSSMatrix) SetM11(m11 float32) {
	js.Rewrite("$<.m11 = $1", m11)
}

// M12 prop
func (*WebKitCSSMatrix) M12() (m12 float32) {
	js.Rewrite("$<.m12")
	return m12
}

// M12 prop
func (*WebKitCSSMatrix) SetM12(m12 float32) {
	js.Rewrite("$<.m12 = $1", m12)
}

// M13 prop
func (*WebKitCSSMatrix) M13() (m13 float32) {
	js.Rewrite("$<.m13")
	return m13
}

// M13 prop
func (*WebKitCSSMatrix) SetM13(m13 float32) {
	js.Rewrite("$<.m13 = $1", m13)
}

// M14 prop
func (*WebKitCSSMatrix) M14() (m14 float32) {
	js.Rewrite("$<.m14")
	return m14
}

// M14 prop
func (*WebKitCSSMatrix) SetM14(m14 float32) {
	js.Rewrite("$<.m14 = $1", m14)
}

// M21 prop
func (*WebKitCSSMatrix) M21() (m21 float32) {
	js.Rewrite("$<.m21")
	return m21
}

// M21 prop
func (*WebKitCSSMatrix) SetM21(m21 float32) {
	js.Rewrite("$<.m21 = $1", m21)
}

// M22 prop
func (*WebKitCSSMatrix) M22() (m22 float32) {
	js.Rewrite("$<.m22")
	return m22
}

// M22 prop
func (*WebKitCSSMatrix) SetM22(m22 float32) {
	js.Rewrite("$<.m22 = $1", m22)
}

// M23 prop
func (*WebKitCSSMatrix) M23() (m23 float32) {
	js.Rewrite("$<.m23")
	return m23
}

// M23 prop
func (*WebKitCSSMatrix) SetM23(m23 float32) {
	js.Rewrite("$<.m23 = $1", m23)
}

// M24 prop
func (*WebKitCSSMatrix) M24() (m24 float32) {
	js.Rewrite("$<.m24")
	return m24
}

// M24 prop
func (*WebKitCSSMatrix) SetM24(m24 float32) {
	js.Rewrite("$<.m24 = $1", m24)
}

// M31 prop
func (*WebKitCSSMatrix) M31() (m31 float32) {
	js.Rewrite("$<.m31")
	return m31
}

// M31 prop
func (*WebKitCSSMatrix) SetM31(m31 float32) {
	js.Rewrite("$<.m31 = $1", m31)
}

// M32 prop
func (*WebKitCSSMatrix) M32() (m32 float32) {
	js.Rewrite("$<.m32")
	return m32
}

// M32 prop
func (*WebKitCSSMatrix) SetM32(m32 float32) {
	js.Rewrite("$<.m32 = $1", m32)
}

// M33 prop
func (*WebKitCSSMatrix) M33() (m33 float32) {
	js.Rewrite("$<.m33")
	return m33
}

// M33 prop
func (*WebKitCSSMatrix) SetM33(m33 float32) {
	js.Rewrite("$<.m33 = $1", m33)
}

// M34 prop
func (*WebKitCSSMatrix) M34() (m34 float32) {
	js.Rewrite("$<.m34")
	return m34
}

// M34 prop
func (*WebKitCSSMatrix) SetM34(m34 float32) {
	js.Rewrite("$<.m34 = $1", m34)
}

// M41 prop
func (*WebKitCSSMatrix) M41() (m41 float32) {
	js.Rewrite("$<.m41")
	return m41
}

// M41 prop
func (*WebKitCSSMatrix) SetM41(m41 float32) {
	js.Rewrite("$<.m41 = $1", m41)
}

// M42 prop
func (*WebKitCSSMatrix) M42() (m42 float32) {
	js.Rewrite("$<.m42")
	return m42
}

// M42 prop
func (*WebKitCSSMatrix) SetM42(m42 float32) {
	js.Rewrite("$<.m42 = $1", m42)
}

// M43 prop
func (*WebKitCSSMatrix) M43() (m43 float32) {
	js.Rewrite("$<.m43")
	return m43
}

// M43 prop
func (*WebKitCSSMatrix) SetM43(m43 float32) {
	js.Rewrite("$<.m43 = $1", m43)
}

// M44 prop
func (*WebKitCSSMatrix) M44() (m44 float32) {
	js.Rewrite("$<.m44")
	return m44
}

// M44 prop
func (*WebKitCSSMatrix) SetM44(m44 float32) {
	js.Rewrite("$<.m44 = $1", m44)
}
