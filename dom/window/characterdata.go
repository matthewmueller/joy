package window

// js:"CharacterData,omit"
type CharacterData interface {
	Node

	// 	// Remove
	// 	// js:"remove,rewrite=remove"
	// 	Remove()

	// 	// AppendData
	// 	// js:"appendData,rewrite=appenddata"
	// 	AppendData(arg string)

	// 	// DeleteData
	// 	// js:"deleteData,rewrite=deletedata"
	// 	DeleteData(offset uint, count uint)

	// 	// InsertData
	// 	// js:"insertData,rewrite=insertdata"
	// 	InsertData(offset uint, arg string)

	// 	// ReplaceData
	// 	// js:"replaceData,rewrite=replacedata"
	// 	ReplaceData(offset uint, count uint, arg string)

	// 	// SubstringData
	// 	// js:"substringData,rewrite=substringdata"
	// 	SubstringData(offset uint, count uint) (s string)

	// 	// Data prop
	// 	// js:"data,rewrite=data"
	// 	Data() (data string)

	// 	// Data prop
	// 	// js:"setdata,rewrite=setdata"
	// 	SetData(data string)

	// 	// Length prop
	// 	// js:"length,rewrite=length"
	// 	Length() (length uint)
}

// // appenddata fn
// func appenddata(arg string) {
// 	js.Rewrite("$<.appendData($1)", arg)
// }

// // deletedata fn
// func deletedata(offset uint, count uint) {
// 	js.Rewrite("$<.deleteData($1, $2)", offset, count)
// }

// // insertdata fn
// func insertdata(offset uint, arg string) {
// 	js.Rewrite("$<.insertData($1, $2)", offset, arg)
// }

// // replacedata fn
// func replacedata(offset uint, count uint, arg string) {
// 	js.Rewrite("$<.replaceData($1, $2, $3)", offset, count, arg)
// }

// // substringdata fn
// func substringdata(offset uint, count uint) (s string) {
// 	js.Rewrite("$<.substringData($1, $2)", offset, count)
// 	return s
// }

// // data prop
// func data() (data string) {
// 	js.Rewrite("$<.data")
// 	return data
// }

// // setdata prop
// func setdata(data string) {
// 	js.Rewrite("$<.data = data")
// }

// // length prop
// func length() (length uint) {
// 	js.Rewrite("$<.length")
// 	return length
// }
