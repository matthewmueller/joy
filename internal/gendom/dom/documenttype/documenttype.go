package documenttype

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/childnode"
	"github.com/matthewmueller/golly/internal/gendom/dom/namednodemap"
	"github.com/matthewmueller/golly/internal/gendom/dom/node"
	"github.com/matthewmueller/golly/js"
)

type DocumentType struct {
	*node.Node
	*childnode.ChildNode
}

func (*DocumentType) GetEntities() (entities *namednodemap.NamedNodeMap) {
	js.Rewrite("$<.entities")
	return entities
}

func (*DocumentType) GetInternalSubset() (internalSubset string) {
	js.Rewrite("$<.internalSubset")
	return internalSubset
}

func (*DocumentType) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*DocumentType) GetNotations() (notations *namednodemap.NamedNodeMap) {
	js.Rewrite("$<.notations")
	return notations
}

func (*DocumentType) GetPublicID() (publicId string) {
	js.Rewrite("$<.publicId")
	return publicId
}

func (*DocumentType) GetSystemID() (systemId string) {
	js.Rewrite("$<.systemId")
	return systemId
}
