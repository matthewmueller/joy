package window

import "github.com/matthewmueller/golly/js"

// IDBVersionChangeEvent struct
// js:"IDBVersionChangeEvent,omit"
type IDBVersionChangeEvent struct {
	Event
}

// NewVersion
func (*IDBVersionChangeEvent) NewVersion() (newVersion uint64) {
	js.Rewrite("$<.NewVersion")
	return newVersion
}

// OldVersion
func (*IDBVersionChangeEvent) OldVersion() (oldVersion uint64) {
	js.Rewrite("$<.OldVersion")
	return oldVersion
}
