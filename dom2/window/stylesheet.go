package window

import "github.com/matthewmueller/golly/dom2/medialist"

// js:"StyleSheet,omit"
type StyleSheet interface {

	// Disabled
	Disabled() (disabled bool)

	// Disabled
	SetDisabled(disabled bool)

	// Href
	Href() (href string)

	// Media
	Media() (media *medialist.MediaList)

	// OwnerNode
	OwnerNode() (ownerNode Node)

	// ParentStyleSheet
	ParentStyleSheet() (parentStyleSheet StyleSheet)

	// Title
	Title() (title string)

	// Type
	Type() (kind string)
}
