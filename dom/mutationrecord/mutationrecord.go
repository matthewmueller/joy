package mutationrecord

import (
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/js"
)

// MutationRecord struct
// js:"MutationRecord,omit"
type MutationRecord struct {
}

// AddedNodes prop
// js:"addedNodes"
func (*MutationRecord) AddedNodes() (addedNodes *window.NodeList) {
	js.Rewrite("$_.addedNodes")
	return addedNodes
}

// AttributeName prop
// js:"attributeName"
func (*MutationRecord) AttributeName() (attributeName string) {
	js.Rewrite("$_.attributeName")
	return attributeName
}

// AttributeNamespace prop
// js:"attributeNamespace"
func (*MutationRecord) AttributeNamespace() (attributeNamespace string) {
	js.Rewrite("$_.attributeNamespace")
	return attributeNamespace
}

// NextSibling prop
// js:"nextSibling"
func (*MutationRecord) NextSibling() (nextSibling window.Node) {
	js.Rewrite("$_.nextSibling")
	return nextSibling
}

// OldValue prop
// js:"oldValue"
func (*MutationRecord) OldValue() (oldValue string) {
	js.Rewrite("$_.oldValue")
	return oldValue
}

// PreviousSibling prop
// js:"previousSibling"
func (*MutationRecord) PreviousSibling() (previousSibling window.Node) {
	js.Rewrite("$_.previousSibling")
	return previousSibling
}

// RemovedNodes prop
// js:"removedNodes"
func (*MutationRecord) RemovedNodes() (removedNodes *window.NodeList) {
	js.Rewrite("$_.removedNodes")
	return removedNodes
}

// Target prop
// js:"target"
func (*MutationRecord) Target() (target window.Node) {
	js.Rewrite("$_.target")
	return target
}

// Type prop
// js:"type"
func (*MutationRecord) Type() (kind string) {
	js.Rewrite("$_.type")
	return kind
}
