package htmlcanvaselement

import (
	"github.com/matthewmueller/golly/dom/blob"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLCanvasElement struct
// js:"HTMLCanvasElement,omit"
type HTMLCanvasElement struct {
	window.HTMLElement
}

// GetContext fn Returns an object that provides methods and properties for drawing and manipulating images and graphics on a canvas element in a document. A context object includes information about colors, line widths, fonts, and other graphic parameters that can be drawn on a canvas.
//     * @param contextId The identifier (ID) of the type of canvas to create. Internet Explorer 9 and Internet Explorer 10 support only a 2-D context using canvas.getContext("2d"); IE11 Preview also supports 3-D or WebGL context using canvas.getContext("experimental-webgl");
func (*HTMLCanvasElement) GetContext(contextId string, args interface{}) (i interface{}) {
	js.Rewrite("$<.getContext($1, $2)", contextId, args)
	return i
}

// MsToBlob fn Returns a blob object encoded as a Portable Network Graphics (PNG) format from a canvas image or drawing.
func (*HTMLCanvasElement) MsToBlob() (b blob.Blob) {
	js.Rewrite("$<.msToBlob()")
	return b
}

// ToDataURL fn Returns the content of the current canvas as an image that you can use as a source for another canvas or an HTML element.
//     * @param type The standard MIME type for the image format to return. If you do not specify this parameter, the default value is a PNG format image.
func (*HTMLCanvasElement) ToDataURL(kind *string, args interface{}) (s string) {
	js.Rewrite("$<.toDataURL($1, $2)", kind, args)
	return s
}

// Height prop Gets or sets the height of a canvas element on a document.
func (*HTMLCanvasElement) Height() (height uint) {
	js.Rewrite("$<.height")
	return height
}

// Height prop Gets or sets the height of a canvas element on a document.
func (*HTMLCanvasElement) SetHeight(height uint) {
	js.Rewrite("$<.height = $1", height)
}

// Width prop Gets or sets the width of a canvas element on a document.
func (*HTMLCanvasElement) Width() (width uint) {
	js.Rewrite("$<.width")
	return width
}

// Width prop Gets or sets the width of a canvas element on a document.
func (*HTMLCanvasElement) SetWidth(width uint) {
	js.Rewrite("$<.width = $1", width)
}
