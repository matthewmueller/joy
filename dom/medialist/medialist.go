package medialist

import "github.com/matthewmueller/golly/js"

// MediaList struct
// js:"MediaList,omit"
type MediaList struct {
}

// AppendMedium fn
// js:"appendMedium"
func (*MediaList) AppendMedium(newMedium string) {
	js.Rewrite("$_.appendMedium($1)", newMedium)
}

// DeleteMedium fn
// js:"deleteMedium"
func (*MediaList) DeleteMedium(oldMedium string) {
	js.Rewrite("$_.deleteMedium($1)", oldMedium)
}

// Item fn
// js:"item"
func (*MediaList) Item(index uint) (s string) {
	js.Rewrite("$_.item($1)", index)
	return s
}

// ToString fn
// js:"toString"
func (*MediaList) ToString() (s string) {
	js.Rewrite("$_.toString()")
	return s
}

// Length prop
// js:"length"
func (*MediaList) Length() (length uint) {
	js.Rewrite("$_.length")
	return length
}

// MediaText prop
// js:"mediaText"
func (*MediaList) MediaText() (mediaText string) {
	js.Rewrite("$_.mediaText")
	return mediaText
}

// SetMediaText prop
// js:"mediaText"
func (*MediaList) SetMediaText(mediaText string) {
	js.Rewrite("$_.mediaText = $1", mediaText)
}
