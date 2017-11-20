package imagedata

import "github.com/matthewmueller/golly/js"

type ImageData struct {
}

func (*ImageData) GetData() (data []uint8) {
	js.Rewrite("$<.data")
	return data
}

func (*ImageData) GetHeight() (height uint) {
	js.Rewrite("$<.height")
	return height
}

func (*ImageData) GetWidth() (width uint) {
	js.Rewrite("$<.width")
	return width
}
