package msgesture

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/element"
	"github.com/matthewmueller/golly/js"
)

type MSGesture struct {
}

func (*MSGesture) AddPointer(pointerId int) {
	js.Rewrite("$<.addPointer($1)", pointerId)
}

func (*MSGesture) Stop() {
	js.Rewrite("$<.stop()")
}

func (*MSGesture) GetTarget() (target *element.Element) {
	js.Rewrite("$<.target")
	return target
}

func (*MSGesture) SetTarget(target *element.Element) {
	js.Rewrite("$<.target = $1", target)
}
