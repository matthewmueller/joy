package imagedata

import "github.com/matthewmueller/golly/js"

// ImageData struct
// js:"ImageData,omit"
type ImageData struct {
}

// Data prop
// js:"data"
func (*ImageData) Data() (data []uint8) {
	js.Rewrite("$_.data")
	return data
}

// Height prop
// js:"height"
func (*ImageData) Height() (height uint) {
	js.Rewrite("$_.height")
	return height
}

// Width prop
// js:"width"
func (*ImageData) Width() (width uint) {
	js.Rewrite("$_.width")
	return width
}
