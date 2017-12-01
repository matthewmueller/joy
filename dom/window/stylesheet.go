package window

import "github.com/matthewmueller/golly/dom/medialist"

// StyleSheet interface
// js:"StyleSheet"
type StyleSheet interface {

	// Disabled prop
	// js:"disabled"
	// jsrewrite:"$_.disabled"
	Disabled() (disabled bool)

	// SetDisabled prop
	// js:"disabled"
	// jsrewrite:"$_.disabled = $1"
	SetDisabled(disabled bool)

	// Href prop
	// js:"href"
	// jsrewrite:"$_.href"
	Href() (href string)

	// Media prop
	// js:"media"
	// jsrewrite:"$_.media"
	Media() (media *medialist.MediaList)

	// OwnerNode prop
	// js:"ownerNode"
	// jsrewrite:"$_.ownerNode"
	OwnerNode() (ownerNode Node)

	// ParentStyleSheet prop
	// js:"parentStyleSheet"
	// jsrewrite:"$_.parentStyleSheet"
	ParentStyleSheet() (parentStyleSheet StyleSheet)

	// Title prop
	// js:"title"
	// jsrewrite:"$_.title"
	Title() (title string)

	// Type prop
	// js:"type"
	// jsrewrite:"$_.type"
	Type() (kind string)
}
