package storage

import "github.com/matthewmueller/golly/js"

// Storage struct
// js:"Storage,omit"
type Storage struct {
}

// Clear fn
func (*Storage) Clear() {
	js.Rewrite("$<.clear()")
}

// GetItem fn
func (*Storage) GetItem(key string) (i interface{}) {
	js.Rewrite("$<.getItem($1)", key)
	return i
}

// Key fn
func (*Storage) Key(index uint) (s string) {
	js.Rewrite("$<.key($1)", index)
	return s
}

// RemoveItem fn
func (*Storage) RemoveItem(key string) {
	js.Rewrite("$<.removeItem($1)", key)
}

// SetItem fn
func (*Storage) SetItem(key string, data string) {
	js.Rewrite("$<.setItem($1, $2)", key, data)
}

// Length prop
func (*Storage) Length() (length uint) {
	js.Rewrite("$<.length")
	return length
}
