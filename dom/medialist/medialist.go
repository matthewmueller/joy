package medialist

import "github.com/matthewmueller/joy/macro"

// MediaList struct
// js:"MediaList,omit"
type MediaList struct {
}

// AppendMedium fn
// js:"appendMedium"
func (*MediaList) AppendMedium(newMedium string) {
	macro.Rewrite("$_.appendMedium($1)", newMedium)
}

// DeleteMedium fn
// js:"deleteMedium"
func (*MediaList) DeleteMedium(oldMedium string) {
	macro.Rewrite("$_.deleteMedium($1)", oldMedium)
}

// Item fn
// js:"item"
func (*MediaList) Item(index uint) (s string) {
	macro.Rewrite("$_.item($1)", index)
	return s
}

// ToString fn
// js:"toString"
func (*MediaList) ToString() (s string) {
	macro.Rewrite("$_.toString()")
	return s
}

// Length prop
// js:"length"
func (*MediaList) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}

// MediaText prop
// js:"mediaText"
func (*MediaList) MediaText() (mediaText string) {
	macro.Rewrite("$_.mediaText")
	return mediaText
}

// SetMediaText prop
// js:"mediaText"
func (*MediaList) SetMediaText(mediaText string) {
	macro.Rewrite("$_.mediaText = $1", mediaText)
}
