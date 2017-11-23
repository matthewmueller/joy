package mutationevent

import (
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// MutationEvent struct
// js:"MutationEvent,omit"
type MutationEvent struct {
	window.Event
}

// InitMutationEvent fn
func (*MutationEvent) InitMutationEvent(typeArg string, canBubbleArg bool, cancelableArg bool, relatedNodeArg window.Node, prevValueArg string, newValueArg string, attrNameArg string, attrChangeArg uint8) {
	js.Rewrite("$<.initMutationEvent($1, $2, $3, $4, $5, $6, $7, $8)", typeArg, canBubbleArg, cancelableArg, relatedNodeArg, prevValueArg, newValueArg, attrNameArg, attrChangeArg)
}

// AttrChange prop
func (*MutationEvent) AttrChange() (attrChange uint8) {
	js.Rewrite("$<.attrChange")
	return attrChange
}

// AttrName prop
func (*MutationEvent) AttrName() (attrName string) {
	js.Rewrite("$<.attrName")
	return attrName
}

// NewValue prop
func (*MutationEvent) NewValue() (newValue string) {
	js.Rewrite("$<.newValue")
	return newValue
}

// PrevValue prop
func (*MutationEvent) PrevValue() (prevValue string) {
	js.Rewrite("$<.prevValue")
	return prevValue
}

// RelatedNode prop
func (*MutationEvent) RelatedNode() (relatedNode window.Node) {
	js.Rewrite("$<.relatedNode")
	return relatedNode
}
