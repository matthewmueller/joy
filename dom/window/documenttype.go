package window

import "github.com/matthewmueller/golly/js"

// DocumentType struct
// js:"DocumentType,omit"
type DocumentType struct {
	Node
}

// Remove fn
func (*DocumentType) Remove() {
	js.Rewrite("$<.remove()")
}

// Entities prop
func (*DocumentType) Entities() (entities *NamedNodeMap) {
	js.Rewrite("$<.entities")
	return entities
}

// InternalSubset prop
func (*DocumentType) InternalSubset() (internalSubset string) {
	js.Rewrite("$<.internalSubset")
	return internalSubset
}

// Name prop
func (*DocumentType) Name() (name string) {
	js.Rewrite("$<.name")
	return name
}

// Notations prop
func (*DocumentType) Notations() (notations *NamedNodeMap) {
	js.Rewrite("$<.notations")
	return notations
}

// PublicID prop
func (*DocumentType) PublicID() (publicId string) {
	js.Rewrite("$<.publicId")
	return publicId
}

// SystemID prop
func (*DocumentType) SystemID() (systemId string) {
	js.Rewrite("$<.systemId")
	return systemId
}
