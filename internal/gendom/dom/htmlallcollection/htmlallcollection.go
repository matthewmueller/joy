package htmlallcollection

import "github.com/matthewmueller/golly/js"

type HTMLAllCollection struct {
}

func (*HTMLAllCollection) Item(nameOrIndex string) (i interface{}) {
	js.Rewrite("$<.item($1)", nameOrIndex)
	return i
}

func (*HTMLAllCollection) NamedItem(name string) (i interface{}) {
	js.Rewrite("$<.namedItem($1)", name)
	return i
}

func (*HTMLAllCollection) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}
