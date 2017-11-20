package mutationevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/node"
	"github.com/matthewmueller/golly/js"
)

type MutationEvent struct {
	*event.Event
}

func (*MutationEvent) InitMutationEvent(typeArg string, canBubbleArg bool, cancelableArg bool, relatedNodeArg *node.Node, prevValueArg string, newValueArg string, attrNameArg string, attrChangeArg uint8) {
	js.Rewrite("$<.initMutationEvent($1, $2, $3, $4, $5, $6, $7, $8)", typeArg, canBubbleArg, cancelableArg, relatedNodeArg, prevValueArg, newValueArg, attrNameArg, attrChangeArg)
}

func (*MutationEvent) GetAttrChange() (attrChange uint8) {
	js.Rewrite("$<.attrChange")
	return attrChange
}

func (*MutationEvent) GetAttrName() (attrName string) {
	js.Rewrite("$<.attrName")
	return attrName
}

func (*MutationEvent) GetNewValue() (newValue string) {
	js.Rewrite("$<.newValue")
	return newValue
}

func (*MutationEvent) GetPrevValue() (prevValue string) {
	js.Rewrite("$<.prevValue")
	return prevValue
}

func (*MutationEvent) GetRelatedNode() (relatedNode *node.Node) {
	js.Rewrite("$<.relatedNode")
	return relatedNode
}
