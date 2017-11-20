package mediastreamerrorevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/mediastreamerror"
	"github.com/matthewmueller/golly/js"
)

type MediaStreamErrorEvent struct {
	*event.Event
}

func (*MediaStreamErrorEvent) GetError() (err *mediastreamerror.MediaStreamError) {
	js.Rewrite("$<.error")
	return err
}
