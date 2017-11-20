package idbversionchangeevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/js"
)

type IDBVersionChangeEvent struct {
	*event.Event
}

func (*IDBVersionChangeEvent) GetNewVersion() (newVersion uint64) {
	js.Rewrite("$<.newVersion")
	return newVersion
}

func (*IDBVersionChangeEvent) GetOldVersion() (oldVersion uint64) {
	js.Rewrite("$<.oldVersion")
	return oldVersion
}
