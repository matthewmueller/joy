package mutationrecord

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// MutationRecord struct
// js:"MutationRecord,omit"
type MutationRecord struct {
}

// AddedNodes
func (*MutationRecord) AddedNodes() (addedNodes *window.NodeList) {
	js.Rewrite("$<.AddedNodes")
	return addedNodes
}

// AttributeName
func (*MutationRecord) AttributeName() (attributeName string) {
	js.Rewrite("$<.AttributeName")
	return attributeName
}

// AttributeNamespace
func (*MutationRecord) AttributeNamespace() (attributeNamespace string) {
	js.Rewrite("$<.AttributeNamespace")
	return attributeNamespace
}

// NextSibling
func (*MutationRecord) NextSibling() (nextSibling window.Node) {
	js.Rewrite("$<.NextSibling")
	return nextSibling
}

// OldValue
func (*MutationRecord) OldValue() (oldValue string) {
	js.Rewrite("$<.OldValue")
	return oldValue
}

// PreviousSibling
func (*MutationRecord) PreviousSibling() (previousSibling window.Node) {
	js.Rewrite("$<.PreviousSibling")
	return previousSibling
}

// RemovedNodes
func (*MutationRecord) RemovedNodes() (removedNodes *window.NodeList) {
	js.Rewrite("$<.RemovedNodes")
	return removedNodes
}

// Target
func (*MutationRecord) Target() (target window.Node) {
	js.Rewrite("$<.Target")
	return target
}

// Type
func (*MutationRecord) Type() (kind string) {
	js.Rewrite("$<.Type")
	return kind
}
