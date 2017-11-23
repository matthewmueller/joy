package window

import "github.com/matthewmueller/golly/js"

// js:"Attr,omit"
type Attr struct {
	Node
}

// Name
func (*Attr) Name() (name string) {
	js.Rewrite("$<.Name")
	return name
}

// OwnerElement
func (*Attr) OwnerElement() (ownerElement Element) {
	js.Rewrite("$<.OwnerElement")
	return ownerElement
}

// Prefix
func (*Attr) Prefix() (prefix string) {
	js.Rewrite("$<.Prefix")
	return prefix
}

// Specified
func (*Attr) Specified() (specified bool) {
	js.Rewrite("$<.Specified")
	return specified
}

// Value
func (*Attr) Value() (value string) {
	js.Rewrite("$<.Value")
	return value
}

// Value
func (*Attr) SetValue(value string) {
	js.Rewrite("$<.Value = $1", value)
}
