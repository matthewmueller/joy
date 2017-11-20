package domtokenlist

import "github.com/matthewmueller/golly/js"

type DOMTokenList struct {
}

func (*DOMTokenList) Add(token string) {
	js.Rewrite("$<.add($1)", token)
}

func (*DOMTokenList) Contains(token string) (b bool) {
	js.Rewrite("$<.contains($1)", token)
	return b
}

func (*DOMTokenList) Item(index uint) (s string) {
	js.Rewrite("$<.item($1)", index)
	return s
}

func (*DOMTokenList) Remove(token string) {
	js.Rewrite("$<.remove($1)", token)
}

func (*DOMTokenList) Toggle(token string, force bool) (b bool) {
	js.Rewrite("$<.toggle($1, $2)", token, force)
	return b
}

func (*DOMTokenList) ToString() (s string) {
	js.Rewrite("$<.toString()")
	return s
}

func (*DOMTokenList) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}
