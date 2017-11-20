package medialist

import "github.com/matthewmueller/golly/js"

type MediaList struct {
}

func (*MediaList) AppendMedium(newMedium string) {
	js.Rewrite("$<.appendMedium($1)", newMedium)
}

func (*MediaList) DeleteMedium(oldMedium string) {
	js.Rewrite("$<.deleteMedium($1)", oldMedium)
}

func (*MediaList) Item(index uint) (s string) {
	js.Rewrite("$<.item($1)", index)
	return s
}

func (*MediaList) ToString() (s string) {
	js.Rewrite("$<.toString()")
	return s
}

func (*MediaList) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}

func (*MediaList) GetMediaText() (mediaText string) {
	js.Rewrite("$<.mediaText")
	return mediaText
}

func (*MediaList) SetMediaText(mediaText string) {
	js.Rewrite("$<.mediaText = $1", mediaText)
}
