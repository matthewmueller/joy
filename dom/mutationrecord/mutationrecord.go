package mutationrecord

import (
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// MutationRecord struct
// js:"MutationRecord,omit"
type MutationRecord struct {
}

// AddedNodes prop
func (*MutationRecord) AddedNodes() (addedNodes *window.NodeList) {
	js.Rewrite("$<.addedNodes")
	return addedNodes
}

// AttributeName prop
func (*MutationRecord) AttributeName() (attributeName string) {
	js.Rewrite("$<.attributeName")
	return attributeName
}

// AttributeNamespace prop
func (*MutationRecord) AttributeNamespace() (attributeNamespace string) {
	js.Rewrite("$<.attributeNamespace")
	return attributeNamespace
}

// NextSibling prop
func (*MutationRecord) NextSibling() (nextSibling window.Node) {
	js.Rewrite("$<.nextSibling")
	return nextSibling
}

// OldValue prop
func (*MutationRecord) OldValue() (oldValue string) {
	js.Rewrite("$<.oldValue")
	return oldValue
}

// PreviousSibling prop
func (*MutationRecord) PreviousSibling() (previousSibling window.Node) {
	js.Rewrite("$<.previousSibling")
	return previousSibling
}

// RemovedNodes prop
func (*MutationRecord) RemovedNodes() (removedNodes *window.NodeList) {
	js.Rewrite("$<.removedNodes")
	return removedNodes
}

// Target prop
func (*MutationRecord) Target() (target window.Node) {
	js.Rewrite("$<.target")
	return target
}

// Type prop
func (*MutationRecord) Type() (kind string) {
	js.Rewrite("$<.type")
	return kind
}
