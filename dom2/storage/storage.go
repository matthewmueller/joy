package storage

import "github.com/matthewmueller/golly/js"

// Storage struct
// js:"Storage,omit"
type Storage struct {
}

// Clear
func (*Storage) Clear() {
	js.Rewrite("$<.Clear()")
}

// GetItem
func (*Storage) GetItem(key string) (i interface{}) {
	js.Rewrite("$<.GetItem($1)", key)
	return i
}

// Key
func (*Storage) Key(index uint) (s string) {
	js.Rewrite("$<.Key($1)", index)
	return s
}

// RemoveItem
func (*Storage) RemoveItem(key string) {
	js.Rewrite("$<.RemoveItem($1)", key)
}

// SetItem
func (*Storage) SetItem(key string, data string) {
	js.Rewrite("$<.SetItem($1, $2)", key, data)
}

// Length
func (*Storage) Length() (length uint) {
	js.Rewrite("$<.Length")
	return length
}
