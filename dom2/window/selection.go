package window

import "github.com/matthewmueller/golly/js"

// Selection struct
// js:"Selection,omit"
type Selection struct {
}

// AddRange
func (*Selection) AddRange(rng *Range) {
	js.Rewrite("$<.AddRange($1)", rng)
}

// Collapse
func (*Selection) Collapse(parentNode Node, offset int) {
	js.Rewrite("$<.Collapse($1, $2)", parentNode, offset)
}

// CollapseToEnd
func (*Selection) CollapseToEnd() {
	js.Rewrite("$<.CollapseToEnd()")
}

// CollapseToStart
func (*Selection) CollapseToStart() {
	js.Rewrite("$<.CollapseToStart()")
}

// ContainsNode
func (*Selection) ContainsNode(node Node, partlyContained bool) (b bool) {
	js.Rewrite("$<.ContainsNode($1, $2)", node, partlyContained)
	return b
}

// DeleteFromDocument
func (*Selection) DeleteFromDocument() {
	js.Rewrite("$<.DeleteFromDocument()")
}

// Empty
func (*Selection) Empty() {
	js.Rewrite("$<.Empty()")
}

// Extend
func (*Selection) Extend(newNode Node, offset int) {
	js.Rewrite("$<.Extend($1, $2)", newNode, offset)
}

// GetRangeAt
func (*Selection) GetRangeAt(index int) (r *Range) {
	js.Rewrite("$<.GetRangeAt($1)", index)
	return r
}

// RemoveAllRanges
func (*Selection) RemoveAllRanges() {
	js.Rewrite("$<.RemoveAllRanges()")
}

// RemoveRange
func (*Selection) RemoveRange(rng *Range) {
	js.Rewrite("$<.RemoveRange($1)", rng)
}

// SelectAllChildren
func (*Selection) SelectAllChildren(parentNode Node) {
	js.Rewrite("$<.SelectAllChildren($1)", parentNode)
}

// SetBaseAndExtent
func (*Selection) SetBaseAndExtent(baseNode Node, baseOffset int, extentNode Node, extentOffset int) {
	js.Rewrite("$<.SetBaseAndExtent($1, $2, $3, $4)", baseNode, baseOffset, extentNode, extentOffset)
}

// SetPosition
func (*Selection) SetPosition(parentNode Node, offset int) {
	js.Rewrite("$<.SetPosition($1, $2)", parentNode, offset)
}

// ToString
func (*Selection) ToString() (s string) {
	js.Rewrite("$<.ToString()")
	return s
}

// AnchorNode
func (*Selection) AnchorNode() (anchorNode Node) {
	js.Rewrite("$<.AnchorNode")
	return anchorNode
}

// AnchorOffset
func (*Selection) AnchorOffset() (anchorOffset int) {
	js.Rewrite("$<.AnchorOffset")
	return anchorOffset
}

// BaseNode
func (*Selection) BaseNode() (baseNode Node) {
	js.Rewrite("$<.BaseNode")
	return baseNode
}

// BaseOffset
func (*Selection) BaseOffset() (baseOffset int) {
	js.Rewrite("$<.BaseOffset")
	return baseOffset
}

// ExtentNode
func (*Selection) ExtentNode() (extentNode Node) {
	js.Rewrite("$<.ExtentNode")
	return extentNode
}

// ExtentOffset
func (*Selection) ExtentOffset() (extentOffset int) {
	js.Rewrite("$<.ExtentOffset")
	return extentOffset
}

// FocusNode
func (*Selection) FocusNode() (focusNode Node) {
	js.Rewrite("$<.FocusNode")
	return focusNode
}

// FocusOffset
func (*Selection) FocusOffset() (focusOffset int) {
	js.Rewrite("$<.FocusOffset")
	return focusOffset
}

// IsCollapsed
func (*Selection) IsCollapsed() (isCollapsed bool) {
	js.Rewrite("$<.IsCollapsed")
	return isCollapsed
}

// RangeCount
func (*Selection) RangeCount() (rangeCount int) {
	js.Rewrite("$<.RangeCount")
	return rangeCount
}

// Type
func (*Selection) Type() (kind string) {
	js.Rewrite("$<.Type")
	return kind
}
