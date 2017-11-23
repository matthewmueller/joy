package window

import "github.com/matthewmueller/golly/js"

// HTMLHeadElement struct
// js:"HTMLHeadElement,omit"
type HTMLHeadElement struct {
	HTMLElement
}

// Profile prop
func (*HTMLHeadElement) Profile() (profile string) {
	js.Rewrite("$<.profile")
	return profile
}

// Profile prop
func (*HTMLHeadElement) SetProfile(profile string) {
	js.Rewrite("$<.profile = $1", profile)
}
