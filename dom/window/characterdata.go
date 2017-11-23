package window

// js:"CharacterData,omit"
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

	// Data prop
	SetData(data string)

	// Length prop
	// js:"length"
	Length() (length uint)
}
