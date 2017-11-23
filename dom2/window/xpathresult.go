package window

import "github.com/matthewmueller/golly/js"

// js:"XPathResult,omit"
type XPathResult struct {
}

// IterateNext
func (*XPathResult) IterateNext() (n Node) {
	js.Rewrite("$<.IterateNext()")
	return n
}

// SnapshotItem
func (*XPathResult) SnapshotItem(index uint) (n Node) {
	js.Rewrite("$<.SnapshotItem($1)", index)
	return n
}

// BooleanValue
func (*XPathResult) BooleanValue() (booleanValue bool) {
	js.Rewrite("$<.BooleanValue")
	return booleanValue
}

// InvalidIteratorState
func (*XPathResult) InvalidIteratorState() (invalidIteratorState bool) {
	js.Rewrite("$<.InvalidIteratorState")
	return invalidIteratorState
}

// NumberValue
func (*XPathResult) NumberValue() (numberValue float32) {
	js.Rewrite("$<.NumberValue")
	return numberValue
}

// ResultType
func (*XPathResult) ResultType() (resultType uint8) {
	js.Rewrite("$<.ResultType")
	return resultType
}

// SingleNodeValue
func (*XPathResult) SingleNodeValue() (singleNodeValue Node) {
	js.Rewrite("$<.SingleNodeValue")
	return singleNodeValue
}

// SnapshotLength
func (*XPathResult) SnapshotLength() (snapshotLength uint) {
	js.Rewrite("$<.SnapshotLength")
	return snapshotLength
}

// StringValue
func (*XPathResult) StringValue() (stringValue string) {
	js.Rewrite("$<.StringValue")
	return stringValue
}
