package focusevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/uievent"
	"github.com/matthewmueller/golly/internal/gendom/dom/window"
	"github.com/matthewmueller/golly/js"
)

type FocusEvent struct {
	*uievent.UIEvent
}

func (*FocusEvent) InitFocusEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *window.Window, detailArg int, relatedTargetArg *eventtarget.EventTarget) {
	js.Rewrite("$<.initFocusEvent($1, $2, $3, $4, $5, $6)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, relatedTargetArg)
}

func (*FocusEvent) GetRelatedTarget() (relatedTarget *eventtarget.EventTarget) {
	js.Rewrite("$<.relatedTarget")
	return relatedTarget
}
