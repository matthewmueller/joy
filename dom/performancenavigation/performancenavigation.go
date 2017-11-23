package performancenavigation

import "github.com/matthewmueller/golly/js"

// PerformanceNavigation struct
// js:"PerformanceNavigation,omit"
type PerformanceNavigation struct {
}

// ToJSON fn
func (*PerformanceNavigation) ToJSON() (i interface{}) {
	js.Rewrite("$<.toJSON()")
	return i
}

// RedirectCount prop
func (*PerformanceNavigation) RedirectCount() (redirectCount uint) {
	js.Rewrite("$<.redirectCount")
	return redirectCount
}

// Type prop
func (*PerformanceNavigation) Type() (kind uint) {
	js.Rewrite("$<.type")
	return kind
}
