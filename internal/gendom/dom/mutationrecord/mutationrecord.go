package mutationrecord

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/node"
	"github.com/matthewmueller/golly/internal/gendom/dom/nodelist"
	"github.com/matthewmueller/golly/js"
)

type MutationRecord struct {
}

func (*MutationRecord) GetAddedNodes() (addedNodes *nodelist.NodeList) {
	js.Rewrite("$<.addedNodes")
	return addedNodes
}

func (*MutationRecord) GetAttributeName() (attributeName string) {
	js.Rewrite("$<.attributeName")
	return attributeName
}

func (*MutationRecord) GetAttributeNamespace() (attributeNamespace string) {
	js.Rewrite("$<.attributeNamespace")
	return attributeNamespace
}

func (*MutationRecord) GetNextSibling() (nextSibling *node.Node) {
	js.Rewrite("$<.nextSibling")
	return nextSibling
}

func (*MutationRecord) GetOldValue() (oldValue string) {
	js.Rewrite("$<.oldValue")
	return oldValue
}

func (*MutationRecord) GetPreviousSibling() (previousSibling *node.Node) {
	js.Rewrite("$<.previousSibling")
	return previousSibling
}

func (*MutationRecord) GetRemovedNodes() (removedNodes *nodelist.NodeList) {
	js.Rewrite("$<.removedNodes")
	return removedNodes
}

func (*MutationRecord) GetTarget() (target *node.Node) {
	js.Rewrite("$<.target")
	return target
}

func (*MutationRecord) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}
