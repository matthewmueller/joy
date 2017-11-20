package xpathresult

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/node"
	"github.com/matthewmueller/golly/js"
)

type XPathResult struct {
}

func (*XPathResult) IterateNext() (n *node.Node) {
	js.Rewrite("$<.iterateNext()")
	return n
}

func (*XPathResult) SnapshotItem(index uint) (n *node.Node) {
	js.Rewrite("$<.snapshotItem($1)", index)
	return n
}

func (*XPathResult) GetBooleanValue() (booleanValue bool) {
	js.Rewrite("$<.booleanValue")
	return booleanValue
}

func (*XPathResult) GetInvalidIteratorState() (invalidIteratorState bool) {
	js.Rewrite("$<.invalidIteratorState")
	return invalidIteratorState
}

func (*XPathResult) GetNumberValue() (numberValue float32) {
	js.Rewrite("$<.numberValue")
	return numberValue
}

func (*XPathResult) GetResultType() (resultType uint8) {
	js.Rewrite("$<.resultType")
	return resultType
}

func (*XPathResult) GetSingleNodeValue() (singleNodeValue *node.Node) {
	js.Rewrite("$<.singleNodeValue")
	return singleNodeValue
}

func (*XPathResult) GetSnapshotLength() (snapshotLength uint) {
	js.Rewrite("$<.snapshotLength")
	return snapshotLength
}

func (*XPathResult) GetStringValue() (stringValue string) {
	js.Rewrite("$<.stringValue")
	return stringValue
}
