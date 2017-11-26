package window

import (
	"github.com/matthewmueller/golly/dom/medialist"
	"github.com/matthewmueller/golly/js"
)

// js:"StyleSheet,omit"
type StyleSheet interface {

	// Disabled prop
	// js:"disabled,rewrite=disabled"
	Disabled() (disabled bool)

	// Disabled prop
	// js:"setdisabled,rewrite=setdisabled"
	SetDisabled(disabled bool)

	// Href prop
	// js:"href,rewrite=href"
	Href() (href string)

	// Media prop
	// js:"media,rewrite=media"
	Media() (media *medialist.MediaList)

	// OwnerNode prop
	// js:"ownerNode,rewrite=ownernode"
	OwnerNode() (ownerNode Node)

	// ParentStyleSheet prop
	// js:"parentStyleSheet,rewrite=parentstylesheet"
	ParentStyleSheet() (parentStyleSheet StyleSheet)

	// Title prop
	// js:"title,rewrite=title"
	Title() (title string)

	// Type prop
	// js:"type,rewrite=kind"
	Type() (kind string)
}

// disabled prop
func disabled() (disabled bool) {
	js.Rewrite("$<.disabled")
	return disabled
}

// setdisabled prop
func setdisabled(disabled bool) {
	js.Rewrite("$<.disabled = disabled")
}

// href prop
func href() (href string) {
	js.Rewrite("$<.href")
	return href
}

// media prop
func media() (media *medialist.MediaList) {
	js.Rewrite("$<.media")
	return media
}

// ownernode prop
func ownernode() (ownerNode Node) {
	js.Rewrite("$<.ownerNode")
	return ownerNode
}

// // parentstylesheet prop
// func parentstylesheet() (parentStyleSheet StyleSheet) {
// 	js.Rewrite("$<.parentStyleSheet")
// 	return parentStyleSheet
// }

// // title prop
// func title() (title string) {
// 	js.Rewrite("$<.title")
// 	return title
// }

// // kind prop
// func kind() (kind string) {
// 	js.Rewrite("$<.type")
// 	return kind
// }
