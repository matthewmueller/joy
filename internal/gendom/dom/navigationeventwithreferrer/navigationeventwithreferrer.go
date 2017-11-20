package navigationeventwithreferrer

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/navigationevent"
	"github.com/matthewmueller/golly/js"
)

type NavigationEventWithReferrer struct {
	*navigationevent.NavigationEvent
}

func (*NavigationEventWithReferrer) GetReferer() (referer string) {
	js.Rewrite("$<.referer")
	return referer
}
