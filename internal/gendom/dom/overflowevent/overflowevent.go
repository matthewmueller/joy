package overflowevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/uievent"
	"github.com/matthewmueller/golly/js"
)

type OverflowEvent struct {
	*uievent.UIEvent
}

func (*OverflowEvent) GetHorizontalOverflow() (horizontalOverflow bool) {
	js.Rewrite("$<.horizontalOverflow")
	return horizontalOverflow
}

func (*OverflowEvent) GetOrient() (orient uint) {
	js.Rewrite("$<.orient")
	return orient
}

func (*OverflowEvent) GetVerticalOverflow() (verticalOverflow bool) {
	js.Rewrite("$<.verticalOverflow")
	return verticalOverflow
}
