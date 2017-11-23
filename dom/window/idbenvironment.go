package window

import "github.com/matthewmueller/golly/js"

// IDBEnvironment struct
// js:"IDBEnvironment,omit"
type IDBEnvironment struct {
}

// IndexedDB prop
func (*IDBEnvironment) IndexedDB() (indexedDB *IDBFactory) {
	js.Rewrite("$<.indexedDB")
	return indexedDB
}
