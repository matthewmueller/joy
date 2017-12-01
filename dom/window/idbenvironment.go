package window

// IDBEnvironment interface
// js:"IDBEnvironment"
type IDBEnvironment interface {

	// IndexedDB prop
	// js:"indexedDB"
	// jsrewrite:"$_.indexedDB"
	IndexedDB() (indexedDB *IDBFactory)
}
