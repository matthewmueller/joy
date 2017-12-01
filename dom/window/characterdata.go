package window

// CharacterData interface
// js:"CharacterData"
type CharacterData interface {
	Node

	// AppendData
	// js:"appendData"
	// jsrewrite:"$_.appendData($1)"
	AppendData(arg string)

	// DeleteData
	// js:"deleteData"
	// jsrewrite:"$_.deleteData($1, $2)"
	DeleteData(offset uint, count uint)

	// InsertData
	// js:"insertData"
	// jsrewrite:"$_.insertData($1, $2)"
	InsertData(offset uint, arg string)

	// ReplaceData
	// js:"replaceData"
	// jsrewrite:"$_.replaceData($1, $2, $3)"
	ReplaceData(offset uint, count uint, arg string)

	// SubstringData
	// js:"substringData"
	// jsrewrite:"$_.substringData($1, $2)"
	SubstringData(offset uint, count uint) (s string)

	// Data prop
	// js:"data"
	// jsrewrite:"$_.data"
	Data() (data string)

	// SetData prop
	// js:"data"
	// jsrewrite:"$_.data = $1"
	SetData(data string)

	// Length prop
	// js:"length"
	// jsrewrite:"$_.length"
	Length() (length uint)
}
