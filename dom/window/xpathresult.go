package window

import "github.com/matthewmueller/golly/js"

// XPathResult struct
// js:"XPathResult,omit"
type XPathResult struct {
}

// IterateNext fn
func (*XPathResult) IterateNext() (n Node) {
	js.Rewrite("$<.iterateNext()")
	return n
}

// SnapshotItem fn
func (*XPathResult) SnapshotItem(index uint) (n Node) {
	js.Rewrite("$<.snapshotItem($1)", index)
	return n
}

// BooleanValue prop
func (*XPathResult) BooleanValue() (booleanValue bool) {
	js.Rewrite("$<.booleanValue")
	return booleanValue
}

// InvalidIteratorState prop
func (*XPathResult) InvalidIteratorState() (invalidIteratorState bool) {
	js.Rewrite("$<.invalidIteratorState")
	return invalidIteratorState
}

// NumberValue prop
func (*XPathResult) NumberValue() (numberValue float32) {
	js.Rewrite("$<.numberValue")
	return numberValue
}

// ResultType prop
func (*XPathResult) ResultType() (resultType uint8) {
	js.Rewrite("$<.resultType")
	return resultType
}

// SingleNodeValue prop
func (*XPathResult) SingleNodeValue() (singleNodeValue Node) {
	js.Rewrite("$<.singleNodeValue")
	return singleNodeValue
}

// SnapshotLength prop
func (*XPathResult) SnapshotLength() (snapshotLength uint) {
	js.Rewrite("$<.snapshotLength")
	return snapshotLength
}

// StringValue prop
func (*XPathResult) StringValue() (stringValue string) {
	js.Rewrite("$<.stringValue")
	return stringValue
}
