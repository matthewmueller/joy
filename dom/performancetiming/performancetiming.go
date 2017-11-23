package performancetiming

import "github.com/matthewmueller/golly/js"

// PerformanceTiming struct
// js:"PerformanceTiming,omit"
type PerformanceTiming struct {
}

// ToJSON fn
func (*PerformanceTiming) ToJSON() (i interface{}) {
	js.Rewrite("$<.toJSON()")
	return i
}

// ConnectEnd prop
func (*PerformanceTiming) ConnectEnd() (connectEnd uint64) {
	js.Rewrite("$<.connectEnd")
	return connectEnd
}

// ConnectStart prop
func (*PerformanceTiming) ConnectStart() (connectStart uint64) {
	js.Rewrite("$<.connectStart")
	return connectStart
}

// DomainLookupEnd prop
func (*PerformanceTiming) DomainLookupEnd() (domainLookupEnd uint64) {
	js.Rewrite("$<.domainLookupEnd")
	return domainLookupEnd
}

// DomainLookupStart prop
func (*PerformanceTiming) DomainLookupStart() (domainLookupStart uint64) {
	js.Rewrite("$<.domainLookupStart")
	return domainLookupStart
}

// DomComplete prop
func (*PerformanceTiming) DomComplete() (domComplete uint64) {
	js.Rewrite("$<.domComplete")
	return domComplete
}

// DomContentLoadedEventEnd prop
func (*PerformanceTiming) DomContentLoadedEventEnd() (domContentLoadedEventEnd uint64) {
	js.Rewrite("$<.domContentLoadedEventEnd")
	return domContentLoadedEventEnd
}

// DomContentLoadedEventStart prop
func (*PerformanceTiming) DomContentLoadedEventStart() (domContentLoadedEventStart uint64) {
	js.Rewrite("$<.domContentLoadedEventStart")
	return domContentLoadedEventStart
}

// DomInteractive prop
func (*PerformanceTiming) DomInteractive() (domInteractive uint64) {
	js.Rewrite("$<.domInteractive")
	return domInteractive
}

// DomLoading prop
func (*PerformanceTiming) DomLoading() (domLoading uint64) {
	js.Rewrite("$<.domLoading")
	return domLoading
}

// FetchStart prop
func (*PerformanceTiming) FetchStart() (fetchStart uint64) {
	js.Rewrite("$<.fetchStart")
	return fetchStart
}

// LoadEventEnd prop
func (*PerformanceTiming) LoadEventEnd() (loadEventEnd uint64) {
	js.Rewrite("$<.loadEventEnd")
	return loadEventEnd
}

// LoadEventStart prop
func (*PerformanceTiming) LoadEventStart() (loadEventStart uint64) {
	js.Rewrite("$<.loadEventStart")
	return loadEventStart
}

// MsFirstPaint prop
func (*PerformanceTiming) MsFirstPaint() (msFirstPaint uint64) {
	js.Rewrite("$<.msFirstPaint")
	return msFirstPaint
}

// NavigationStart prop
func (*PerformanceTiming) NavigationStart() (navigationStart uint64) {
	js.Rewrite("$<.navigationStart")
	return navigationStart
}

// RedirectEnd prop
func (*PerformanceTiming) RedirectEnd() (redirectEnd uint64) {
	js.Rewrite("$<.redirectEnd")
	return redirectEnd
}

// RedirectStart prop
func (*PerformanceTiming) RedirectStart() (redirectStart uint64) {
	js.Rewrite("$<.redirectStart")
	return redirectStart
}

// RequestStart prop
func (*PerformanceTiming) RequestStart() (requestStart uint64) {
	js.Rewrite("$<.requestStart")
	return requestStart
}

// ResponseEnd prop
func (*PerformanceTiming) ResponseEnd() (responseEnd uint64) {
	js.Rewrite("$<.responseEnd")
	return responseEnd
}

// ResponseStart prop
func (*PerformanceTiming) ResponseStart() (responseStart uint64) {
	js.Rewrite("$<.responseStart")
	return responseStart
}

// UnloadEventEnd prop
func (*PerformanceTiming) UnloadEventEnd() (unloadEventEnd uint64) {
	js.Rewrite("$<.unloadEventEnd")
	return unloadEventEnd
}

// UnloadEventStart prop
func (*PerformanceTiming) UnloadEventStart() (unloadEventStart uint64) {
	js.Rewrite("$<.unloadEventStart")
	return unloadEventStart
}
