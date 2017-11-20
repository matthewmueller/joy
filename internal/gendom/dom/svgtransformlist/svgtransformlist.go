package svgtransformlist

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svgmatrix"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgtransform"
	"github.com/matthewmueller/golly/js"
)

type SVGTransformList struct {
}

func (*SVGTransformList) AppendItem(newItem *svgtransform.SVGTransform) (s *svgtransform.SVGTransform) {
	js.Rewrite("$<.appendItem($1)", newItem)
	return s
}

func (*SVGTransformList) Clear() {
	js.Rewrite("$<.clear()")
}

func (*SVGTransformList) Consolidate() (s *svgtransform.SVGTransform) {
	js.Rewrite("$<.consolidate()")
	return s
}

func (*SVGTransformList) CreateSVGTransformFromMatrix(matrix *svgmatrix.SVGMatrix) (s *svgtransform.SVGTransform) {
	js.Rewrite("$<.createSVGTransformFromMatrix($1)", matrix)
	return s
}

func (*SVGTransformList) GetItem(index uint) (s *svgtransform.SVGTransform) {
	js.Rewrite("$<.getItem($1)", index)
	return s
}

func (*SVGTransformList) Initialize(newItem *svgtransform.SVGTransform) (s *svgtransform.SVGTransform) {
	js.Rewrite("$<.initialize($1)", newItem)
	return s
}

func (*SVGTransformList) InsertItemBefore(newItem *svgtransform.SVGTransform, index uint) (s *svgtransform.SVGTransform) {
	js.Rewrite("$<.insertItemBefore($1, $2)", newItem, index)
	return s
}

func (*SVGTransformList) RemoveItem(index uint) (s *svgtransform.SVGTransform) {
	js.Rewrite("$<.removeItem($1)", index)
	return s
}

func (*SVGTransformList) ReplaceItem(newItem *svgtransform.SVGTransform, index uint) (s *svgtransform.SVGTransform) {
	js.Rewrite("$<.replaceItem($1, $2)", newItem, index)
	return s
}

func (*SVGTransformList) GetNumberOfItems() (numberOfItems uint) {
	js.Rewrite("$<.numberOfItems")
	return numberOfItems
}
