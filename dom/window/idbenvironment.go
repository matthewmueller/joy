package window

// IDBEnvironment interface
// js:"IDBEnvironment"
type IDBEnvironment interface {

	// IndexedDB prop
	// js:"indexedDB"
	IndexedDB() (indexedDB *IDBFactory)
}
