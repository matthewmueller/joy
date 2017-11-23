package window

import "github.com/matthewmueller/golly/js"

// js:"HTMLHeadElement,omit"
type HTMLHeadElement struct {
	HTMLElement
}

// Profile
func (*HTMLHeadElement) Profile() (profile string) {
	js.Rewrite("$<.Profile")
	return profile
}

// Profile
func (*HTMLHeadElement) SetProfile(profile string) {
	js.Rewrite("$<.Profile = $1", profile)
}
