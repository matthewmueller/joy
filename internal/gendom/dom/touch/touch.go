package touch

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/js"
)

type Touch struct {
}

func (*Touch) GetClientX() (clientX int) {
	js.Rewrite("$<.clientX")
	return clientX
}

func (*Touch) GetClientY() (clientY int) {
	js.Rewrite("$<.clientY")
	return clientY
}

func (*Touch) GetIdentifier() (identifier int) {
	js.Rewrite("$<.identifier")
	return identifier
}

func (*Touch) GetPageX() (pageX int) {
	js.Rewrite("$<.pageX")
	return pageX
}

func (*Touch) GetPageY() (pageY int) {
	js.Rewrite("$<.pageY")
	return pageY
}

func (*Touch) GetScreenX() (screenX int) {
	js.Rewrite("$<.screenX")
	return screenX
}

func (*Touch) GetScreenY() (screenY int) {
	js.Rewrite("$<.screenY")
	return screenY
}

func (*Touch) GetTarget() (target *eventtarget.EventTarget) {
	js.Rewrite("$<.target")
	return target
}
