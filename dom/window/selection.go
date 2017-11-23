package window

import "github.com/matthewmueller/golly/js"

// Selection struct
// js:"Selection,omit"
type Selection struct {
}

// AddRange fn
func (*Selection) AddRange(rng *Range) {
	js.Rewrite("$<.addRange($1)", rng)
}

// Collapse fn
func (*Selection) Collapse(parentNode Node, offset int) {
	js.Rewrite("$<.collapse($1, $2)", parentNode, offset)
}

// CollapseToEnd fn
func (*Selection) CollapseToEnd() {
	js.Rewrite("$<.collapseToEnd()")
}

// CollapseToStart fn
func (*Selection) CollapseToStart() {
	js.Rewrite("$<.collapseToStart()")
}

// ContainsNode fn
func (*Selection) ContainsNode(node Node, partlyContained bool) (b bool) {
	js.Rewrite("$<.containsNode($1, $2)", node, partlyContained)
	return b
}

// DeleteFromDocument fn
func (*Selection) DeleteFromDocument() {
	js.Rewrite("$<.deleteFromDocument()")
}

// Empty fn
func (*Selection) Empty() {
	js.Rewrite("$<.empty()")
}

// Extend fn
func (*Selection) Extend(newNode Node, offset int) {
	js.Rewrite("$<.extend($1, $2)", newNode, offset)
}

// GetRangeAt fn
func (*Selection) GetRangeAt(index int) (r *Range) {
	js.Rewrite("$<.getRangeAt($1)", index)
	return r
}

// RemoveAllRanges fn
func (*Selection) RemoveAllRanges() {
	js.Rewrite("$<.removeAllRanges()")
}

// RemoveRange fn
func (*Selection) RemoveRange(rng *Range) {
	js.Rewrite("$<.removeRange($1)", rng)
}

// SelectAllChildren fn
func (*Selection) SelectAllChildren(parentNode Node) {
	js.Rewrite("$<.selectAllChildren($1)", parentNode)
}

// SetBaseAndExtent fn
func (*Selection) SetBaseAndExtent(baseNode Node, baseOffset int, extentNode Node, extentOffset int) {
	js.Rewrite("$<.setBaseAndExtent($1, $2, $3, $4)", baseNode, baseOffset, extentNode, extentOffset)
}

// SetPosition fn
func (*Selection) SetPosition(parentNode Node, offset int) {
	js.Rewrite("$<.setPosition($1, $2)", parentNode, offset)
}

// ToString fn
func (*Selection) ToString() (s string) {
	js.Rewrite("$<.toString()")
	return s
}

// AnchorNode prop
func (*Selection) AnchorNode() (anchorNode Node) {
	js.Rewrite("$<.anchorNode")
	return anchorNode
}

// AnchorOffset prop
func (*Selection) AnchorOffset() (anchorOffset int) {
	js.Rewrite("$<.anchorOffset")
	return anchorOffset
}

// BaseNode prop
func (*Selection) BaseNode() (baseNode Node) {
	js.Rewrite("$<.baseNode")
	return baseNode
}

// BaseOffset prop
func (*Selection) BaseOffset() (baseOffset int) {
	js.Rewrite("$<.baseOffset")
	return baseOffset
}

// ExtentNode prop
func (*Selection) ExtentNode() (extentNode Node) {
	js.Rewrite("$<.extentNode")
	return extentNode
}

// ExtentOffset prop
func (*Selection) ExtentOffset() (extentOffset int) {
	js.Rewrite("$<.extentOffset")
	return extentOffset
}

// FocusNode prop
func (*Selection) FocusNode() (focusNode Node) {
	js.Rewrite("$<.focusNode")
	return focusNode
}

// FocusOffset prop
func (*Selection) FocusOffset() (focusOffset int) {
	js.Rewrite("$<.focusOffset")
	return focusOffset
}

// IsCollapsed prop
func (*Selection) IsCollapsed() (isCollapsed bool) {
	js.Rewrite("$<.isCollapsed")
	return isCollapsed
}

// RangeCount prop
func (*Selection) RangeCount() (rangeCount int) {
	js.Rewrite("$<.rangeCount")
	return rangeCount
}

// Type prop
func (*Selection) Type() (kind string) {
	js.Rewrite("$<.type")
	return kind
}
