package window

import "github.com/matthewmueller/golly/js"

// Attr struct
// js:"Attr,omit"
type Attr struct {
	Node
}

// Name prop
func (*Attr) Name() (name string) {
	js.Rewrite("$<.name")
	return name
}

// OwnerElement prop
func (*Attr) OwnerElement() (ownerElement Element) {
	js.Rewrite("$<.ownerElement")
	return ownerElement
}

// Prefix prop
func (*Attr) Prefix() (prefix string) {
	js.Rewrite("$<.prefix")
	return prefix
}

// Specified prop
func (*Attr) Specified() (specified bool) {
	js.Rewrite("$<.specified")
	return specified
}

// Value prop
func (*Attr) Value() (value string) {
	js.Rewrite("$<.value")
	return value
}

// Value prop
func (*Attr) SetValue(value string) {
	js.Rewrite("$<.value = $1", value)
}
