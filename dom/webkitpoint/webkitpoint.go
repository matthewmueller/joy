package webkitpoint

import "github.com/matthewmueller/golly/js"

// New fn
func New(x *float32, y *float32) *WebKitPoint {
	js.Rewrite("WebKitPoint")
	return &WebKitPoint{}
}

// WebKitPoint struct
// js:"WebKitPoint,omit"
type WebKitPoint struct {
}

// X prop
func (*WebKitPoint) X() (x float32) {
	js.Rewrite("$<.x")
	return x
}

// X prop
func (*WebKitPoint) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

// Y prop
func (*WebKitPoint) Y() (y float32) {
	js.Rewrite("$<.y")
	return y
}

// Y prop
func (*WebKitPoint) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}
