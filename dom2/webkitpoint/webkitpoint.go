package webkitpoint

import "github.com/matthewmueller/golly/js"

// js:"WebKitPoint,omit"
type WebKitPoint struct {
}

// X
func (*WebKitPoint) X() (x float32) {
	js.Rewrite("$<.X")
	return x
}

// X
func (*WebKitPoint) SetX(x float32) {
	js.Rewrite("$<.X = $1", x)
}

// Y
func (*WebKitPoint) Y() (y float32) {
	js.Rewrite("$<.Y")
	return y
}

// Y
func (*WebKitPoint) SetY(y float32) {
	js.Rewrite("$<.Y = $1", y)
}
