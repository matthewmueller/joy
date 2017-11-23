package idb

// js:"CharacterData,omit"
type CharacterData interface {
	Node

	// AppendData
	AppendData(arg string)

	// DeleteData
	DeleteData(offset uint, count uint)

	// InsertData
	InsertData(offset uint, arg string)

	// ReplaceData
	ReplaceData(offset uint, count uint, arg string)

	// SubstringData
	SubstringData(offset uint, count uint) (s string)

	// Data
	Data() (data string)

	// Data
	SetData(data string)

	// Length
	Length() (length uint)
}
