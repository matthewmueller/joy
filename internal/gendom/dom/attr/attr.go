package attr

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/element"
	"github.com/matthewmueller/golly/internal/gendom/dom/node"
	"github.com/matthewmueller/golly/js"
)

type Attr struct {
	*node.Node
}

func (*Attr) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*Attr) GetOwnerElement() (ownerElement *element.Element) {
	js.Rewrite("$<.ownerElement")
	return ownerElement
}

func (*Attr) GetPrefix() (prefix string) {
	js.Rewrite("$<.prefix")
	return prefix
}

func (*Attr) GetSpecified() (specified bool) {
	js.Rewrite("$<.specified")
	return specified
}

func (*Attr) GetValue() (value string) {
	js.Rewrite("$<.value")
	return value
}

func (*Attr) SetValue(value string) {
	js.Rewrite("$<.value = $1", value)
}
