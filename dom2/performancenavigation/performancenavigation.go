package performancenavigation

import "github.com/matthewmueller/golly/js"

// js:"PerformanceNavigation,omit"
type PerformanceNavigation struct {
}

// ToJSON
func (*PerformanceNavigation) ToJSON() (i interface{}) {
	js.Rewrite("$<.ToJSON()")
	return i
}

// RedirectCount
func (*PerformanceNavigation) RedirectCount() (redirectCount uint) {
	js.Rewrite("$<.RedirectCount")
	return redirectCount
}

// Type
func (*PerformanceNavigation) Type() (kind uint) {
	js.Rewrite("$<.Type")
	return kind
}
