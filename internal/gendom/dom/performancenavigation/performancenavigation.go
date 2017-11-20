package performancenavigation

import "github.com/matthewmueller/golly/js"

type PerformanceNavigation struct {
}

func (*PerformanceNavigation) ToJSON() (i interface{}) {
	js.Rewrite("$<.toJSON()")
	return i
}

func (*PerformanceNavigation) GetRedirectCount() (redirectCount uint) {
	js.Rewrite("$<.redirectCount")
	return redirectCount
}

func (*PerformanceNavigation) GetType() (kind uint) {
	js.Rewrite("$<.type")
	return kind
}
