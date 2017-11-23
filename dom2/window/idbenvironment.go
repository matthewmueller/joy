package window

// js:"IDBEnvironment,omit"
type IDBEnvironment interface {

	// IndexedDB
	IndexedDB() (indexedDB *IDBFactory)
}
