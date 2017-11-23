package window

import (
	"github.com/matthewmueller/golly/dom2/storage"
	"github.com/matthewmueller/golly/js"
)

// StorageEvent struct
// js:"StorageEvent,omit"
type StorageEvent struct {
	Event
}

// InitStorageEvent fn
func (*StorageEvent) InitStorageEvent(typeArg string, canBubbleArg bool, cancelableArg bool, keyArg string, oldValueArg interface{}, newValueArg interface{}, urlArg string, storageAreaArg *storage.Storage) {
	js.Rewrite("$<.initStorageEvent($1, $2, $3, $4, $5, $6, $7, $8)", typeArg, canBubbleArg, cancelableArg, keyArg, oldValueArg, newValueArg, urlArg, storageAreaArg)
}

// Key prop
func (*StorageEvent) Key() (key string) {
	js.Rewrite("$<.key")
	return key
}

// NewValue prop
func (*StorageEvent) NewValue() (newValue interface{}) {
	js.Rewrite("$<.newValue")
	return newValue
}

// OldValue prop
func (*StorageEvent) OldValue() (oldValue interface{}) {
	js.Rewrite("$<.oldValue")
	return oldValue
}

// StorageArea prop
func (*StorageEvent) StorageArea() (storageArea *storage.Storage) {
	js.Rewrite("$<.storageArea")
	return storageArea
}

// URL prop
func (*StorageEvent) URL() (url string) {
	js.Rewrite("$<.url")
	return url
}
