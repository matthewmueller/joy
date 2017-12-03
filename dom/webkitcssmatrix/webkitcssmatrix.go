package webkitcssmatrix

import "github.com/matthewmueller/joy/macro"

// New fn
func New(text *string) *WebKitCSSMatrix {
	macro.Rewrite("WebKitCSSMatrix")
	return &WebKitCSSMatrix{}
}

// WebKitCSSMatrix struct
// js:"WebKitCSSMatrix,omit"
type WebKitCSSMatrix struct {
}

// Inverse fn
// js:"inverse"
func (*WebKitCSSMatrix) Inverse() (w *WebKitCSSMatrix) {
	macro.Rewrite("$_.inverse()")
	return w
}

// Multiply fn
// js:"multiply"
func (*WebKitCSSMatrix) Multiply(secondMatrix *WebKitCSSMatrix) (w *WebKitCSSMatrix) {
	macro.Rewrite("$_.multiply($1)", secondMatrix)
	return w
}

// Rotate fn
// js:"rotate"
func (*WebKitCSSMatrix) Rotate(angleX float32, angleY *float32, angleZ *float32) (w *WebKitCSSMatrix) {
	macro.Rewrite("$_.rotate($1, $2, $3)", angleX, angleY, angleZ)
	return w
}

// RotateAxisAngle fn
// js:"rotateAxisAngle"
func (*WebKitCSSMatrix) RotateAxisAngle(x float32, y float32, z float32, angle float32) (w *WebKitCSSMatrix) {
	macro.Rewrite("$_.rotateAxisAngle($1, $2, $3, $4)", x, y, z, angle)
	return w
}

// Scale fn
// js:"scale"
func (*WebKitCSSMatrix) Scale(scaleX float32, scaleY *float32, scaleZ *float32) (w *WebKitCSSMatrix) {
	macro.Rewrite("$_.scale($1, $2, $3)", scaleX, scaleY, scaleZ)
	return w
}

// SetMatrixValue fn
// js:"setMatrixValue"
func (*WebKitCSSMatrix) SetMatrixValue(value string) {
	macro.Rewrite("$_.setMatrixValue($1)", value)
}

// SkewX fn
// js:"skewX"
func (*WebKitCSSMatrix) SkewX(angle float32) (w *WebKitCSSMatrix) {
	macro.Rewrite("$_.skewX($1)", angle)
	return w
}

// SkewY fn
// js:"skewY"
func (*WebKitCSSMatrix) SkewY(angle float32) (w *WebKitCSSMatrix) {
	macro.Rewrite("$_.skewY($1)", angle)
	return w
}

// ToString fn
// js:"toString"
func (*WebKitCSSMatrix) ToString() (s string) {
	macro.Rewrite("$_.toString()")
	return s
}

// Translate fn
// js:"translate"
func (*WebKitCSSMatrix) Translate(x float32, y float32, z *float32) (w *WebKitCSSMatrix) {
	macro.Rewrite("$_.translate($1, $2, $3)", x, y, z)
	return w
}

// A prop
// js:"a"
func (*WebKitCSSMatrix) A() (a float32) {
	macro.Rewrite("$_.a")
	return a
}

// SetA prop
// js:"a"
func (*WebKitCSSMatrix) SetA(a float32) {
	macro.Rewrite("$_.a = $1", a)
}

// B prop
// js:"b"
func (*WebKitCSSMatrix) B() (b float32) {
	macro.Rewrite("$_.b")
	return b
}

// SetB prop
// js:"b"
func (*WebKitCSSMatrix) SetB(b float32) {
	macro.Rewrite("$_.b = $1", b)
}

// C prop
// js:"c"
func (*WebKitCSSMatrix) C() (c float32) {
	macro.Rewrite("$_.c")
	return c
}

// SetC prop
// js:"c"
func (*WebKitCSSMatrix) SetC(c float32) {
	macro.Rewrite("$_.c = $1", c)
}

// D prop
// js:"d"
func (*WebKitCSSMatrix) D() (d float32) {
	macro.Rewrite("$_.d")
	return d
}

// SetD prop
// js:"d"
func (*WebKitCSSMatrix) SetD(d float32) {
	macro.Rewrite("$_.d = $1", d)
}

// E prop
// js:"e"
func (*WebKitCSSMatrix) E() (e float32) {
	macro.Rewrite("$_.e")
	return e
}

// SetE prop
// js:"e"
func (*WebKitCSSMatrix) SetE(e float32) {
	macro.Rewrite("$_.e = $1", e)
}

// F prop
// js:"f"
func (*WebKitCSSMatrix) F() (f float32) {
	macro.Rewrite("$_.f")
	return f
}

// SetF prop
// js:"f"
func (*WebKitCSSMatrix) SetF(f float32) {
	macro.Rewrite("$_.f = $1", f)
}

// M11 prop
// js:"m11"
func (*WebKitCSSMatrix) M11() (m11 float32) {
	macro.Rewrite("$_.m11")
	return m11
}

// SetM11 prop
// js:"m11"
func (*WebKitCSSMatrix) SetM11(m11 float32) {
	macro.Rewrite("$_.m11 = $1", m11)
}

// M12 prop
// js:"m12"
func (*WebKitCSSMatrix) M12() (m12 float32) {
	macro.Rewrite("$_.m12")
	return m12
}

// SetM12 prop
// js:"m12"
func (*WebKitCSSMatrix) SetM12(m12 float32) {
	macro.Rewrite("$_.m12 = $1", m12)
}

// M13 prop
// js:"m13"
func (*WebKitCSSMatrix) M13() (m13 float32) {
	macro.Rewrite("$_.m13")
	return m13
}

// SetM13 prop
// js:"m13"
func (*WebKitCSSMatrix) SetM13(m13 float32) {
	macro.Rewrite("$_.m13 = $1", m13)
}

// M14 prop
// js:"m14"
func (*WebKitCSSMatrix) M14() (m14 float32) {
	macro.Rewrite("$_.m14")
	return m14
}

// SetM14 prop
// js:"m14"
func (*WebKitCSSMatrix) SetM14(m14 float32) {
	macro.Rewrite("$_.m14 = $1", m14)
}

// M21 prop
// js:"m21"
func (*WebKitCSSMatrix) M21() (m21 float32) {
	macro.Rewrite("$_.m21")
	return m21
}

// SetM21 prop
// js:"m21"
func (*WebKitCSSMatrix) SetM21(m21 float32) {
	macro.Rewrite("$_.m21 = $1", m21)
}

// M22 prop
// js:"m22"
func (*WebKitCSSMatrix) M22() (m22 float32) {
	macro.Rewrite("$_.m22")
	return m22
}

// SetM22 prop
// js:"m22"
func (*WebKitCSSMatrix) SetM22(m22 float32) {
	macro.Rewrite("$_.m22 = $1", m22)
}

// M23 prop
// js:"m23"
func (*WebKitCSSMatrix) M23() (m23 float32) {
	macro.Rewrite("$_.m23")
	return m23
}

// SetM23 prop
// js:"m23"
func (*WebKitCSSMatrix) SetM23(m23 float32) {
	macro.Rewrite("$_.m23 = $1", m23)
}

// M24 prop
// js:"m24"
func (*WebKitCSSMatrix) M24() (m24 float32) {
	macro.Rewrite("$_.m24")
	return m24
}

// SetM24 prop
// js:"m24"
func (*WebKitCSSMatrix) SetM24(m24 float32) {
	macro.Rewrite("$_.m24 = $1", m24)
}

// M31 prop
// js:"m31"
func (*WebKitCSSMatrix) M31() (m31 float32) {
	macro.Rewrite("$_.m31")
	return m31
}

// SetM31 prop
// js:"m31"
func (*WebKitCSSMatrix) SetM31(m31 float32) {
	macro.Rewrite("$_.m31 = $1", m31)
}

// M32 prop
// js:"m32"
func (*WebKitCSSMatrix) M32() (m32 float32) {
	macro.Rewrite("$_.m32")
	return m32
}

// SetM32 prop
// js:"m32"
func (*WebKitCSSMatrix) SetM32(m32 float32) {
	macro.Rewrite("$_.m32 = $1", m32)
}

// M33 prop
// js:"m33"
func (*WebKitCSSMatrix) M33() (m33 float32) {
	macro.Rewrite("$_.m33")
	return m33
}

// SetM33 prop
// js:"m33"
func (*WebKitCSSMatrix) SetM33(m33 float32) {
	macro.Rewrite("$_.m33 = $1", m33)
}

// M34 prop
// js:"m34"
func (*WebKitCSSMatrix) M34() (m34 float32) {
	macro.Rewrite("$_.m34")
	return m34
}

// SetM34 prop
// js:"m34"
func (*WebKitCSSMatrix) SetM34(m34 float32) {
	macro.Rewrite("$_.m34 = $1", m34)
}

// M41 prop
// js:"m41"
func (*WebKitCSSMatrix) M41() (m41 float32) {
	macro.Rewrite("$_.m41")
	return m41
}

// SetM41 prop
// js:"m41"
func (*WebKitCSSMatrix) SetM41(m41 float32) {
	macro.Rewrite("$_.m41 = $1", m41)
}

// M42 prop
// js:"m42"
func (*WebKitCSSMatrix) M42() (m42 float32) {
	macro.Rewrite("$_.m42")
	return m42
}

// SetM42 prop
// js:"m42"
func (*WebKitCSSMatrix) SetM42(m42 float32) {
	macro.Rewrite("$_.m42 = $1", m42)
}

// M43 prop
// js:"m43"
func (*WebKitCSSMatrix) M43() (m43 float32) {
	macro.Rewrite("$_.m43")
	return m43
}

// SetM43 prop
// js:"m43"
func (*WebKitCSSMatrix) SetM43(m43 float32) {
	macro.Rewrite("$_.m43 = $1", m43)
}

// M44 prop
// js:"m44"
func (*WebKitCSSMatrix) M44() (m44 float32) {
	macro.Rewrite("$_.m44")
	return m44
}

// SetM44 prop
// js:"m44"
func (*WebKitCSSMatrix) SetM44(m44 float32) {
	macro.Rewrite("$_.m44 = $1", m44)
}
