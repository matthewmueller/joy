package storage

import "github.com/matthewmueller/joy/macro"

// Storage struct
// js:"Storage,omit"
type Storage struct {
}

// Clear fn
// js:"clear"
func (*Storage) Clear() {
	macro.Rewrite("$_.clear()")
}

// GetItem fn
// js:"getItem"
func (*Storage) GetItem(key string) (i interface{}) {
	macro.Rewrite("$_.getItem($1)", key)
	return i
}

// Key fn
// js:"key"
func (*Storage) Key(index uint) (s string) {
	macro.Rewrite("$_.key($1)", index)
	return s
}

// RemoveItem fn
// js:"removeItem"
func (*Storage) RemoveItem(key string) {
	macro.Rewrite("$_.removeItem($1)", key)
}

// SetItem fn
// js:"setItem"
func (*Storage) SetItem(key string, data string) {
	macro.Rewrite("$_.setItem($1, $2)", key, data)
}

// Length prop
// js:"length"
func (*Storage) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}
