package storage

import "github.com/matthewmueller/golly/js"

type Storage struct {
}

func (*Storage) Clear() {
	js.Rewrite("$<.clear()")
}

func (*Storage) GetItem(key string) (i interface{}) {
	js.Rewrite("$<.getItem($1)", key)
	return i
}

func (*Storage) Key(index uint) (s string) {
	js.Rewrite("$<.key($1)", index)
	return s
}

func (*Storage) RemoveItem(key string) {
	js.Rewrite("$<.removeItem($1)", key)
}

func (*Storage) SetItem(key string, data string) {
	js.Rewrite("$<.setItem($1, $2)", key, data)
}

func (*Storage) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}
