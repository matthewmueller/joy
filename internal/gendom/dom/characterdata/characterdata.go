package characterdata

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/childnode"
	"github.com/matthewmueller/golly/internal/gendom/dom/node"
	"github.com/matthewmueller/golly/js"
)

type CharacterData struct {
	*node.Node
	*childnode.ChildNode
}

func (*CharacterData) AppendData(arg string) {
	js.Rewrite("$<.appendData($1)", arg)
}

func (*CharacterData) DeleteData(offset uint, count uint) {
	js.Rewrite("$<.deleteData($1, $2)", offset, count)
}

func (*CharacterData) InsertData(offset uint, arg string) {
	js.Rewrite("$<.insertData($1, $2)", offset, arg)
}

func (*CharacterData) ReplaceData(offset uint, count uint, arg string) {
	js.Rewrite("$<.replaceData($1, $2, $3)", offset, count, arg)
}

func (*CharacterData) SubstringData(offset uint, count uint) (s string) {
	js.Rewrite("$<.substringData($1, $2)", offset, count)
	return s
}

func (*CharacterData) GetData() (data string) {
	js.Rewrite("$<.data")
	return data
}

func (*CharacterData) SetData(data string) {
	js.Rewrite("$<.data = $1", data)
}

func (*CharacterData) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}
