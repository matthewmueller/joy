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

// InitStorageEvent
func (*StorageEvent) InitStorageEvent(typeArg string, canBubbleArg bool, cancelableArg bool, keyArg string, oldValueArg interface{}, newValueArg interface{}, urlArg string, storageAreaArg *storage.Storage) {
	js.Rewrite("$<.InitStorageEvent($1, $2, $3, $4, $5, $6, $7, $8)", typeArg, canBubbleArg, cancelableArg, keyArg, oldValueArg, newValueArg, urlArg, storageAreaArg)
}

// Key
func (*StorageEvent) Key() (key string) {
	js.Rewrite("$<.Key")
	return key
}

// NewValue
func (*StorageEvent) NewValue() (newValue interface{}) {
	js.Rewrite("$<.NewValue")
	return newValue
}

// OldValue
func (*StorageEvent) OldValue() (oldValue interface{}) {
	js.Rewrite("$<.OldValue")
	return oldValue
}

// StorageArea
func (*StorageEvent) StorageArea() (storageArea *storage.Storage) {
	js.Rewrite("$<.StorageArea")
	return storageArea
}

// URL
func (*StorageEvent) URL() (url string) {
	js.Rewrite("$<.URL")
	return url
}
