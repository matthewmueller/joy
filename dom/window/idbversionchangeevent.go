package window

import "github.com/matthewmueller/golly/js"

// IDBVersionChangeEvent struct
// js:"IDBVersionChangeEvent,omit"
type IDBVersionChangeEvent struct {
	Event
}

// NewVersion prop
func (*IDBVersionChangeEvent) NewVersion() (newVersion uint64) {
	js.Rewrite("$<.newVersion")
	return newVersion
}

// OldVersion prop
func (*IDBVersionChangeEvent) OldVersion() (oldVersion uint64) {
	js.Rewrite("$<.oldVersion")
	return oldVersion
}
