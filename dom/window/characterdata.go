package window

// CharacterData interface
// js:"CharacterData"
type CharacterData interface {
	Node

	// AppendData
	// js:"appendData"
	AppendData(arg string)

	// DeleteData
	// js:"deleteData"
	DeleteData(offset uint, count uint)

	// InsertData
	// js:"insertData"
	InsertData(offset uint, arg string)

	// ReplaceData
	// js:"replaceData"
	ReplaceData(offset uint, count uint, arg string)

	// SubstringData
	// js:"substringData"
	SubstringData(offset uint, count uint) (s string)

	// Data prop
	// js:"data"
	Data() (data string)

	// SetData prop
	// js:"data"
	SetData(data string)

	// Length prop
	// js:"length"
	Length() (length uint)
}
