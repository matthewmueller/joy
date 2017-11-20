package clientrect

import "github.com/matthewmueller/golly/js"

type ClientRect struct {
}

func (*ClientRect) GetBottom() (bottom int) {
	js.Rewrite("$<.bottom")
	return bottom
}

func (*ClientRect) SetBottom(bottom int) {
	js.Rewrite("$<.bottom = $1", bottom)
}

func (*ClientRect) GetHeight() (height float32) {
	js.Rewrite("$<.height")
	return height
}

func (*ClientRect) GetLeft() (left int) {
	js.Rewrite("$<.left")
	return left
}

func (*ClientRect) SetLeft(left int) {
	js.Rewrite("$<.left = $1", left)
}

func (*ClientRect) GetRight() (right int) {
	js.Rewrite("$<.right")
	return right
}

func (*ClientRect) SetRight(right int) {
	js.Rewrite("$<.right = $1", right)
}

func (*ClientRect) GetTop() (top int) {
	js.Rewrite("$<.top")
	return top
}

func (*ClientRect) SetTop(top int) {
	js.Rewrite("$<.top = $1", top)
}

func (*ClientRect) GetWidth() (width float32) {
	js.Rewrite("$<.width")
	return width
}
