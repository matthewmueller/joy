package mutationevent

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"MutationEvent,omit"
type MutationEvent struct {
	window.Event
}

// InitMutationEvent
func (*MutationEvent) InitMutationEvent(typeArg string, canBubbleArg bool, cancelableArg bool, relatedNodeArg window.Node, prevValueArg string, newValueArg string, attrNameArg string, attrChangeArg uint8) {
	js.Rewrite("$<.InitMutationEvent($1, $2, $3, $4, $5, $6, $7, $8)", typeArg, canBubbleArg, cancelableArg, relatedNodeArg, prevValueArg, newValueArg, attrNameArg, attrChangeArg)
}

// AttrChange
func (*MutationEvent) AttrChange() (attrChange uint8) {
	js.Rewrite("$<.AttrChange")
	return attrChange
}

// AttrName
func (*MutationEvent) AttrName() (attrName string) {
	js.Rewrite("$<.AttrName")
	return attrName
}

// NewValue
func (*MutationEvent) NewValue() (newValue string) {
	js.Rewrite("$<.NewValue")
	return newValue
}

// PrevValue
func (*MutationEvent) PrevValue() (prevValue string) {
	js.Rewrite("$<.PrevValue")
	return prevValue
}

// RelatedNode
func (*MutationEvent) RelatedNode() (relatedNode window.Node) {
	js.Rewrite("$<.RelatedNode")
	return relatedNode
}
