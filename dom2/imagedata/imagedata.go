package imagedata

import "github.com/matthewmueller/golly/js"

// js:"ImageData,omit"
type ImageData struct {
}

// Data
func (*ImageData) Data() (data []uint8) {
	js.Rewrite("$<.Data")
	return data
}

// Height
func (*ImageData) Height() (height uint) {
	js.Rewrite("$<.Height")
	return height
}

// Width
func (*ImageData) Width() (width uint) {
	js.Rewrite("$<.Width")
	return width
}
