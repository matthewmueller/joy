package storageevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/storage"
	"github.com/matthewmueller/golly/js"
)

type StorageEvent struct {
	*event.Event
}

func (*StorageEvent) InitStorageEvent(typeArg string, canBubbleArg bool, cancelableArg bool, keyArg string, oldValueArg interface{}, newValueArg interface{}, urlArg string, storageAreaArg *storage.Storage) {
	js.Rewrite("$<.initStorageEvent($1, $2, $3, $4, $5, $6, $7, $8)", typeArg, canBubbleArg, cancelableArg, keyArg, oldValueArg, newValueArg, urlArg, storageAreaArg)
}

func (*StorageEvent) GetKey() (key string) {
	js.Rewrite("$<.key")
	return key
}

func (*StorageEvent) GetNewValue() (newValue interface{}) {
	js.Rewrite("$<.newValue")
	return newValue
}

func (*StorageEvent) GetOldValue() (oldValue interface{}) {
	js.Rewrite("$<.oldValue")
	return oldValue
}

func (*StorageEvent) GetStorageArea() (storageArea *storage.Storage) {
	js.Rewrite("$<.storageArea")
	return storageArea
}

func (*StorageEvent) GetURL() (url string) {
	js.Rewrite("$<.url")
	return url
}
