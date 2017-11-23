package performancenavigationtiming

import (
	"github.com/matthewmueller/golly/dom/navigationtype"
	"github.com/matthewmueller/golly/dom/performanceentry"
	"github.com/matthewmueller/golly/js"
)

// PerformanceNavigationTiming struct
// js:"PerformanceNavigationTiming,omit"
type PerformanceNavigationTiming struct {
	performanceentry.PerformanceEntry
}

// ConnectEnd prop
func (*PerformanceNavigationTiming) ConnectEnd() (connectEnd int) {
	js.Rewrite("$<.connectEnd")
	return connectEnd
}

// ConnectStart prop
func (*PerformanceNavigationTiming) ConnectStart() (connectStart int) {
	js.Rewrite("$<.connectStart")
	return connectStart
}

// DomainLookupEnd prop
func (*PerformanceNavigationTiming) DomainLookupEnd() (domainLookupEnd int) {
	js.Rewrite("$<.domainLookupEnd")
	return domainLookupEnd
}

// DomainLookupStart prop
func (*PerformanceNavigationTiming) DomainLookupStart() (domainLookupStart int) {
	js.Rewrite("$<.domainLookupStart")
	return domainLookupStart
}

// DomComplete prop
func (*PerformanceNavigationTiming) DomComplete() (domComplete int) {
	js.Rewrite("$<.domComplete")
	return domComplete
}

// DomContentLoadedEventEnd prop
func (*PerformanceNavigationTiming) DomContentLoadedEventEnd() (domContentLoadedEventEnd int) {
	js.Rewrite("$<.domContentLoadedEventEnd")
	return domContentLoadedEventEnd
}

// DomContentLoadedEventStart prop
func (*PerformanceNavigationTiming) DomContentLoadedEventStart() (domContentLoadedEventStart int) {
	js.Rewrite("$<.domContentLoadedEventStart")
	return domContentLoadedEventStart
}

// DomInteractive prop
func (*PerformanceNavigationTiming) DomInteractive() (domInteractive int) {
	js.Rewrite("$<.domInteractive")
	return domInteractive
}

// DomLoading prop
func (*PerformanceNavigationTiming) DomLoading() (domLoading int) {
	js.Rewrite("$<.domLoading")
	return domLoading
}

// FetchStart prop
func (*PerformanceNavigationTiming) FetchStart() (fetchStart int) {
	js.Rewrite("$<.fetchStart")
	return fetchStart
}

// LoadEventEnd prop
func (*PerformanceNavigationTiming) LoadEventEnd() (loadEventEnd int) {
	js.Rewrite("$<.loadEventEnd")
	return loadEventEnd
}

// LoadEventStart prop
func (*PerformanceNavigationTiming) LoadEventStart() (loadEventStart int) {
	js.Rewrite("$<.loadEventStart")
	return loadEventStart
}

// NavigationStart prop
func (*PerformanceNavigationTiming) NavigationStart() (navigationStart int) {
	js.Rewrite("$<.navigationStart")
	return navigationStart
}

// RedirectCount prop
func (*PerformanceNavigationTiming) RedirectCount() (redirectCount uint8) {
	js.Rewrite("$<.redirectCount")
	return redirectCount
}

// RedirectEnd prop
func (*PerformanceNavigationTiming) RedirectEnd() (redirectEnd int) {
	js.Rewrite("$<.redirectEnd")
	return redirectEnd
}

// RedirectStart prop
func (*PerformanceNavigationTiming) RedirectStart() (redirectStart int) {
	js.Rewrite("$<.redirectStart")
	return redirectStart
}

// RequestStart prop
func (*PerformanceNavigationTiming) RequestStart() (requestStart int) {
	js.Rewrite("$<.requestStart")
	return requestStart
}

// ResponseEnd prop
func (*PerformanceNavigationTiming) ResponseEnd() (responseEnd int) {
	js.Rewrite("$<.responseEnd")
	return responseEnd
}

// ResponseStart prop
func (*PerformanceNavigationTiming) ResponseStart() (responseStart int) {
	js.Rewrite("$<.responseStart")
	return responseStart
}

// Type prop
func (*PerformanceNavigationTiming) Type() (kind *navigationtype.NavigationType) {
	js.Rewrite("$<.type")
	return kind
}

// UnloadEventEnd prop
func (*PerformanceNavigationTiming) UnloadEventEnd() (unloadEventEnd int) {
	js.Rewrite("$<.unloadEventEnd")
	return unloadEventEnd
}

// UnloadEventStart prop
func (*PerformanceNavigationTiming) UnloadEventStart() (unloadEventStart int) {
	js.Rewrite("$<.unloadEventStart")
	return unloadEventStart
}
