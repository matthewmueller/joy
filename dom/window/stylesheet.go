package window

import "github.com/matthewmueller/golly/dom/medialist"

// js:"StyleSheet,omit"
type StyleSheet interface {

	// Disabled prop
	// js:"disabled"
	Disabled() (disabled bool)

	// Disabled prop
	// js:"setdisabled"
	SetDisabled(disabled bool)

	// Href prop
	// js:"href"
	Href() (href string)

	// Media prop
	// js:"media"
	Media() (media *medialist.MediaList)

	// OwnerNode prop
	// js:"ownerNode"
	OwnerNode() (ownerNode Node)

	// ParentStyleSheet prop
	// js:"parentStyleSheet"
	ParentStyleSheet() (parentStyleSheet StyleSheet)

	// Title prop
	// js:"title"
	Title() (title string)

	// Type prop
	// js:"type"
	Type() (kind string)
}
