package svgrect

import "github.com/matthewmueller/golly/js"

type SVGRect struct {
}

func (*SVGRect) GetHeight() (height float32) {
	js.Rewrite("$<.height")
	return height
}

func (*SVGRect) SetHeight(height float32) {
	js.Rewrite("$<.height = $1", height)
}

func (*SVGRect) GetWidth() (width float32) {
	js.Rewrite("$<.width")
	return width
}

func (*SVGRect) SetWidth(width float32) {
	js.Rewrite("$<.width = $1", width)
}

func (*SVGRect) GetX() (x float32) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGRect) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

func (*SVGRect) GetY() (y float32) {
	js.Rewrite("$<.y")
	return y
}

func (*SVGRect) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}
