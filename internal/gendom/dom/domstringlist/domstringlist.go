package domstringlist

import "github.com/matthewmueller/golly/js"

type DOMStringList struct {
}

func (*DOMStringList) Contains(str string) (b bool) {
	js.Rewrite("$<.contains($1)", str)
	return b
}

func (*DOMStringList) Item(index uint) (s string) {
	js.Rewrite("$<.item($1)", index)
	return s
}

func (*DOMStringList) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}
