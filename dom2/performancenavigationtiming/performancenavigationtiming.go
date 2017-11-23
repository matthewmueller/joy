package performancenavigationtiming

import (
	"github.com/matthewmueller/golly/dom2/navigationtype"
	"github.com/matthewmueller/golly/dom2/performanceentry"
	"github.com/matthewmueller/golly/js"
)

// js:"PerformanceNavigationTiming,omit"
type PerformanceNavigationTiming struct {
	performanceentry.PerformanceEntry
}

// ConnectEnd
func (*PerformanceNavigationTiming) ConnectEnd() (connectEnd int) {
	js.Rewrite("$<.ConnectEnd")
	return connectEnd
}

// ConnectStart
func (*PerformanceNavigationTiming) ConnectStart() (connectStart int) {
	js.Rewrite("$<.ConnectStart")
	return connectStart
}

// DomainLookupEnd
func (*PerformanceNavigationTiming) DomainLookupEnd() (domainLookupEnd int) {
	js.Rewrite("$<.DomainLookupEnd")
	return domainLookupEnd
}

// DomainLookupStart
func (*PerformanceNavigationTiming) DomainLookupStart() (domainLookupStart int) {
	js.Rewrite("$<.DomainLookupStart")
	return domainLookupStart
}

// DomComplete
func (*PerformanceNavigationTiming) DomComplete() (domComplete int) {
	js.Rewrite("$<.DomComplete")
	return domComplete
}

// DomContentLoadedEventEnd
func (*PerformanceNavigationTiming) DomContentLoadedEventEnd() (domContentLoadedEventEnd int) {
	js.Rewrite("$<.DomContentLoadedEventEnd")
	return domContentLoadedEventEnd
}

// DomContentLoadedEventStart
func (*PerformanceNavigationTiming) DomContentLoadedEventStart() (domContentLoadedEventStart int) {
	js.Rewrite("$<.DomContentLoadedEventStart")
	return domContentLoadedEventStart
}

// DomInteractive
func (*PerformanceNavigationTiming) DomInteractive() (domInteractive int) {
	js.Rewrite("$<.DomInteractive")
	return domInteractive
}

// DomLoading
func (*PerformanceNavigationTiming) DomLoading() (domLoading int) {
	js.Rewrite("$<.DomLoading")
	return domLoading
}

// FetchStart
func (*PerformanceNavigationTiming) FetchStart() (fetchStart int) {
	js.Rewrite("$<.FetchStart")
	return fetchStart
}

// LoadEventEnd
func (*PerformanceNavigationTiming) LoadEventEnd() (loadEventEnd int) {
	js.Rewrite("$<.LoadEventEnd")
	return loadEventEnd
}

// LoadEventStart
func (*PerformanceNavigationTiming) LoadEventStart() (loadEventStart int) {
	js.Rewrite("$<.LoadEventStart")
	return loadEventStart
}

// NavigationStart
func (*PerformanceNavigationTiming) NavigationStart() (navigationStart int) {
	js.Rewrite("$<.NavigationStart")
	return navigationStart
}

// RedirectCount
func (*PerformanceNavigationTiming) RedirectCount() (redirectCount uint8) {
	js.Rewrite("$<.RedirectCount")
	return redirectCount
}

// RedirectEnd
func (*PerformanceNavigationTiming) RedirectEnd() (redirectEnd int) {
	js.Rewrite("$<.RedirectEnd")
	return redirectEnd
}

// RedirectStart
func (*PerformanceNavigationTiming) RedirectStart() (redirectStart int) {
	js.Rewrite("$<.RedirectStart")
	return redirectStart
}

// RequestStart
func (*PerformanceNavigationTiming) RequestStart() (requestStart int) {
	js.Rewrite("$<.RequestStart")
	return requestStart
}

// ResponseEnd
func (*PerformanceNavigationTiming) ResponseEnd() (responseEnd int) {
	js.Rewrite("$<.ResponseEnd")
	return responseEnd
}

// ResponseStart
func (*PerformanceNavigationTiming) ResponseStart() (responseStart int) {
	js.Rewrite("$<.ResponseStart")
	return responseStart
}

// Type
func (*PerformanceNavigationTiming) Type() (kind *navigationtype.NavigationType) {
	js.Rewrite("$<.Type")
	return kind
}

// UnloadEventEnd
func (*PerformanceNavigationTiming) UnloadEventEnd() (unloadEventEnd int) {
	js.Rewrite("$<.UnloadEventEnd")
	return unloadEventEnd
}

// UnloadEventStart
func (*PerformanceNavigationTiming) UnloadEventStart() (unloadEventStart int) {
	js.Rewrite("$<.UnloadEventStart")
	return unloadEventStart
}
