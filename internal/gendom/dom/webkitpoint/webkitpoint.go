package webkitpoint

import "github.com/matthewmueller/golly/js"

type WebKitPoint struct {
}

func (*WebKitPoint) GetX() (x float32) {
	js.Rewrite("$<.x")
	return x
}

func (*WebKitPoint) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

func (*WebKitPoint) GetY() (y float32) {
	js.Rewrite("$<.y")
	return y
}

func (*WebKitPoint) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}
