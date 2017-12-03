package storage

import "github.com/matthewmueller/joy/js"

// Storage struct
// js:"Storage,omit"
type Storage struct {
}

// Clear fn
// js:"clear"
func (*Storage) Clear() {
	js.Rewrite("$_.clear()")
}

// GetItem fn
// js:"getItem"
func (*Storage) GetItem(key string) (i interface{}) {
	js.Rewrite("$_.getItem($1)", key)
	return i
}

// Key fn
// js:"key"
func (*Storage) Key(index uint) (s string) {
	js.Rewrite("$_.key($1)", index)
	return s
}

// RemoveItem fn
// js:"removeItem"
func (*Storage) RemoveItem(key string) {
	js.Rewrite("$_.removeItem($1)", key)
}

// SetItem fn
// js:"setItem"
func (*Storage) SetItem(key string, data string) {
	js.Rewrite("$_.setItem($1, $2)", key, data)
}

// Length prop
// js:"length"
func (*Storage) Length() (length uint) {
	js.Rewrite("$_.length")
	return length
}
