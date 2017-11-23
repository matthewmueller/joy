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

// Inverse
func (*WebKitCSSMatrix) Inverse() (w *WebKitCSSMatrix) {
	js.Rewrite("$<.Inverse()")
	return w
}

// Multiply
func (*WebKitCSSMatrix) Multiply(secondMatrix *WebKitCSSMatrix) (w *WebKitCSSMatrix) {
	js.Rewrite("$<.Multiply($1)", secondMatrix)
	return w
}

// Rotate
func (*WebKitCSSMatrix) Rotate(angleX float32, angleY *float32, angleZ *float32) (w *WebKitCSSMatrix) {
	js.Rewrite("$<.Rotate($1, $2, $3)", angleX, angleY, angleZ)
	return w
}

// RotateAxisAngle
func (*WebKitCSSMatrix) RotateAxisAngle(x float32, y float32, z float32, angle float32) (w *WebKitCSSMatrix) {
	js.Rewrite("$<.RotateAxisAngle($1, $2, $3, $4)", x, y, z, angle)
	return w
}

// Scale
func (*WebKitCSSMatrix) Scale(scaleX float32, scaleY *float32, scaleZ *float32) (w *WebKitCSSMatrix) {
	js.Rewrite("$<.Scale($1, $2, $3)", scaleX, scaleY, scaleZ)
	return w
}

// SetMatrixValue
func (*WebKitCSSMatrix) SetMatrixValue(value string) {
	js.Rewrite("$<.SetMatrixValue($1)", value)
}

// SkewX
func (*WebKitCSSMatrix) SkewX(angle float32) (w *WebKitCSSMatrix) {
	js.Rewrite("$<.SkewX($1)", angle)
	return w
}

// SkewY
func (*WebKitCSSMatrix) SkewY(angle float32) (w *WebKitCSSMatrix) {
	js.Rewrite("$<.SkewY($1)", angle)
	return w
}

// ToString
func (*WebKitCSSMatrix) ToString() (s string) {
	js.Rewrite("$<.ToString()")
	return s
}

// Translate
func (*WebKitCSSMatrix) Translate(x float32, y float32, z *float32) (w *WebKitCSSMatrix) {
	js.Rewrite("$<.Translate($1, $2, $3)", x, y, z)
	return w
}

// A
func (*WebKitCSSMatrix) A() (a float32) {
	js.Rewrite("$<.A")
	return a
}

// A
func (*WebKitCSSMatrix) SetA(a float32) {
	js.Rewrite("$<.A = $1", a)
}

// B
func (*WebKitCSSMatrix) B() (b float32) {
	js.Rewrite("$<.B")
	return b
}

// B
func (*WebKitCSSMatrix) SetB(b float32) {
	js.Rewrite("$<.B = $1", b)
}

// C
func (*WebKitCSSMatrix) C() (c float32) {
	js.Rewrite("$<.C")
	return c
}

// C
func (*WebKitCSSMatrix) SetC(c float32) {
	js.Rewrite("$<.C = $1", c)
}

// D
func (*WebKitCSSMatrix) D() (d float32) {
	js.Rewrite("$<.D")
	return d
}

// D
func (*WebKitCSSMatrix) SetD(d float32) {
	js.Rewrite("$<.D = $1", d)
}

// E
func (*WebKitCSSMatrix) E() (e float32) {
	js.Rewrite("$<.E")
	return e
}

// E
func (*WebKitCSSMatrix) SetE(e float32) {
	js.Rewrite("$<.E = $1", e)
}

// F
func (*WebKitCSSMatrix) F() (f float32) {
	js.Rewrite("$<.F")
	return f
}

// F
func (*WebKitCSSMatrix) SetF(f float32) {
	js.Rewrite("$<.F = $1", f)
}

// M11
func (*WebKitCSSMatrix) M11() (m11 float32) {
	js.Rewrite("$<.M11")
	return m11
}

// M11
func (*WebKitCSSMatrix) SetM11(m11 float32) {
	js.Rewrite("$<.M11 = $1", m11)
}

// M12
func (*WebKitCSSMatrix) M12() (m12 float32) {
	js.Rewrite("$<.M12")
	return m12
}

// M12
func (*WebKitCSSMatrix) SetM12(m12 float32) {
	js.Rewrite("$<.M12 = $1", m12)
}

// M13
func (*WebKitCSSMatrix) M13() (m13 float32) {
	js.Rewrite("$<.M13")
	return m13
}

// M13
func (*WebKitCSSMatrix) SetM13(m13 float32) {
	js.Rewrite("$<.M13 = $1", m13)
}

// M14
func (*WebKitCSSMatrix) M14() (m14 float32) {
	js.Rewrite("$<.M14")
	return m14
}

// M14
func (*WebKitCSSMatrix) SetM14(m14 float32) {
	js.Rewrite("$<.M14 = $1", m14)
}

// M21
func (*WebKitCSSMatrix) M21() (m21 float32) {
	js.Rewrite("$<.M21")
	return m21
}

// M21
func (*WebKitCSSMatrix) SetM21(m21 float32) {
	js.Rewrite("$<.M21 = $1", m21)
}

// M22
func (*WebKitCSSMatrix) M22() (m22 float32) {
	js.Rewrite("$<.M22")
	return m22
}

// M22
func (*WebKitCSSMatrix) SetM22(m22 float32) {
	js.Rewrite("$<.M22 = $1", m22)
}

// M23
func (*WebKitCSSMatrix) M23() (m23 float32) {
	js.Rewrite("$<.M23")
	return m23
}

// M23
func (*WebKitCSSMatrix) SetM23(m23 float32) {
	js.Rewrite("$<.M23 = $1", m23)
}

// M24
func (*WebKitCSSMatrix) M24() (m24 float32) {
	js.Rewrite("$<.M24")
	return m24
}

// M24
func (*WebKitCSSMatrix) SetM24(m24 float32) {
	js.Rewrite("$<.M24 = $1", m24)
}

// M31
func (*WebKitCSSMatrix) M31() (m31 float32) {
	js.Rewrite("$<.M31")
	return m31
}

// M31
func (*WebKitCSSMatrix) SetM31(m31 float32) {
	js.Rewrite("$<.M31 = $1", m31)
}

// M32
func (*WebKitCSSMatrix) M32() (m32 float32) {
	js.Rewrite("$<.M32")
	return m32
}

// M32
func (*WebKitCSSMatrix) SetM32(m32 float32) {
	js.Rewrite("$<.M32 = $1", m32)
}

// M33
func (*WebKitCSSMatrix) M33() (m33 float32) {
	js.Rewrite("$<.M33")
	return m33
}

// M33
func (*WebKitCSSMatrix) SetM33(m33 float32) {
	js.Rewrite("$<.M33 = $1", m33)
}

// M34
func (*WebKitCSSMatrix) M34() (m34 float32) {
	js.Rewrite("$<.M34")
	return m34
}

// M34
func (*WebKitCSSMatrix) SetM34(m34 float32) {
	js.Rewrite("$<.M34 = $1", m34)
}

// M41
func (*WebKitCSSMatrix) M41() (m41 float32) {
	js.Rewrite("$<.M41")
	return m41
}

// M41
func (*WebKitCSSMatrix) SetM41(m41 float32) {
	js.Rewrite("$<.M41 = $1", m41)
}

// M42
func (*WebKitCSSMatrix) M42() (m42 float32) {
	js.Rewrite("$<.M42")
	return m42
}

// M42
func (*WebKitCSSMatrix) SetM42(m42 float32) {
	js.Rewrite("$<.M42 = $1", m42)
}

// M43
func (*WebKitCSSMatrix) M43() (m43 float32) {
	js.Rewrite("$<.M43")
	return m43
}

// M43
func (*WebKitCSSMatrix) SetM43(m43 float32) {
	js.Rewrite("$<.M43 = $1", m43)
}

// M44
func (*WebKitCSSMatrix) M44() (m44 float32) {
	js.Rewrite("$<.M44")
	return m44
}

// M44
func (*WebKitCSSMatrix) SetM44(m44 float32) {
	js.Rewrite("$<.M44 = $1", m44)
}
