package htmlcanvaselement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/blob"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/js"
)

type HTMLCanvasElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLCanvasElement) GetContext(contextId string, args interface{}) (i interface{}) {
	js.Rewrite("$<.getContext($1, $2)", contextId, args)
	return i
}

func (*HTMLCanvasElement) MsToBlob() (b *blob.Blob) {
	js.Rewrite("$<.msToBlob()")
	return b
}

func (*HTMLCanvasElement) ToDataURL(kind string, args interface{}) (s string) {
	js.Rewrite("$<.toDataURL($1, $2)", kind, args)
	return s
}

func (*HTMLCanvasElement) GetHeight() (height uint) {
	js.Rewrite("$<.height")
	return height
}

func (*HTMLCanvasElement) SetHeight(height uint) {
	js.Rewrite("$<.height = $1", height)
}

func (*HTMLCanvasElement) GetWidth() (width uint) {
	js.Rewrite("$<.width")
	return width
}

func (*HTMLCanvasElement) SetWidth(width uint) {
	js.Rewrite("$<.width = $1", width)
}
