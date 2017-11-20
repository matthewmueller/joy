package idbopendbrequest

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/idbrequest"
	"github.com/matthewmueller/golly/internal/gendom/dom/idbversionchangeevent"
	"github.com/matthewmueller/golly/js"
)

type IDBOpenDBRequest struct {
	*idbrequest.IDBRequest
}

func (*IDBOpenDBRequest) GetOnblocked() (blocked *event.Event) {
	js.Rewrite("$<.onblocked")
	return blocked
}

func (*IDBOpenDBRequest) SetOnblocked(blocked *event.Event) {
	js.Rewrite("$<.onblocked = $1", blocked)
}

func (*IDBOpenDBRequest) GetOnupgradeneeded() (upgradeneeded *idbversionchangeevent.IDBVersionChangeEvent) {
	js.Rewrite("$<.onupgradeneeded")
	return upgradeneeded
}

func (*IDBOpenDBRequest) SetOnupgradeneeded(upgradeneeded *idbversionchangeevent.IDBVersionChangeEvent) {
	js.Rewrite("$<.onupgradeneeded = $1", upgradeneeded)
}
