package touchevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/touchlist"
	"github.com/matthewmueller/golly/internal/gendom/dom/uievent"
	"github.com/matthewmueller/golly/js"
)

type TouchEvent struct {
	*uievent.UIEvent
}

func (*TouchEvent) GetAltKey() (altKey bool) {
	js.Rewrite("$<.altKey")
	return altKey
}

func (*TouchEvent) GetChangedTouches() (changedTouches *touchlist.TouchList) {
	js.Rewrite("$<.changedTouches")
	return changedTouches
}

func (*TouchEvent) GetCharCode() (charCode int8) {
	js.Rewrite("$<.charCode")
	return charCode
}

func (*TouchEvent) GetCtrlKey() (ctrlKey bool) {
	js.Rewrite("$<.ctrlKey")
	return ctrlKey
}

func (*TouchEvent) GetKeyCode() (keyCode int8) {
	js.Rewrite("$<.keyCode")
	return keyCode
}

func (*TouchEvent) GetMetaKey() (metaKey bool) {
	js.Rewrite("$<.metaKey")
	return metaKey
}

func (*TouchEvent) GetShiftKey() (shiftKey bool) {
	js.Rewrite("$<.shiftKey")
	return shiftKey
}

func (*TouchEvent) GetTargetTouches() (targetTouches *touchlist.TouchList) {
	js.Rewrite("$<.targetTouches")
	return targetTouches
}

func (*TouchEvent) GetTouches() (touches *touchlist.TouchList) {
	js.Rewrite("$<.touches")
	return touches
}

func (*TouchEvent) GetWhich() (which uint8) {
	js.Rewrite("$<.which")
	return which
}
