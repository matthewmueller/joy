package clientrect

import "github.com/matthewmueller/joy/macro"

// ClientRect struct
// js:"ClientRect,omit"
type ClientRect struct {
}

// Bottom prop
// js:"bottom"
func (*ClientRect) Bottom() (bottom int) {
	macro.Rewrite("$_.bottom")
	return bottom
}

// SetBottom prop
// js:"bottom"
func (*ClientRect) SetBottom(bottom int) {
	macro.Rewrite("$_.bottom = $1", bottom)
}

// Height prop
// js:"height"
func (*ClientRect) Height() (height float32) {
	macro.Rewrite("$_.height")
	return height
}

// Left prop
// js:"left"
func (*ClientRect) Left() (left int) {
	macro.Rewrite("$_.left")
	return left
}

// SetLeft prop
// js:"left"
func (*ClientRect) SetLeft(left int) {
	macro.Rewrite("$_.left = $1", left)
}

// Right prop
// js:"right"
func (*ClientRect) Right() (right int) {
	macro.Rewrite("$_.right")
	return right
}

// SetRight prop
// js:"right"
func (*ClientRect) SetRight(right int) {
	macro.Rewrite("$_.right = $1", right)
}

// Top prop
// js:"top"
func (*ClientRect) Top() (top int) {
	macro.Rewrite("$_.top")
	return top
}

// SetTop prop
// js:"top"
func (*ClientRect) SetTop(top int) {
	macro.Rewrite("$_.top = $1", top)
}

// Width prop
// js:"width"
func (*ClientRect) Width() (width float32) {
	macro.Rewrite("$_.width")
	return width
}
