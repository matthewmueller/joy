package performancetiming

import "github.com/matthewmueller/golly/js"

// PerformanceTiming struct
// js:"PerformanceTiming,omit"
type PerformanceTiming struct {
}

// ToJSON
func (*PerformanceTiming) ToJSON() (i interface{}) {
	js.Rewrite("$<.ToJSON()")
	return i
}

// ConnectEnd
func (*PerformanceTiming) ConnectEnd() (connectEnd uint64) {
	js.Rewrite("$<.ConnectEnd")
	return connectEnd
}

// ConnectStart
func (*PerformanceTiming) ConnectStart() (connectStart uint64) {
	js.Rewrite("$<.ConnectStart")
	return connectStart
}

// DomainLookupEnd
func (*PerformanceTiming) DomainLookupEnd() (domainLookupEnd uint64) {
	js.Rewrite("$<.DomainLookupEnd")
	return domainLookupEnd
}

// DomainLookupStart
func (*PerformanceTiming) DomainLookupStart() (domainLookupStart uint64) {
	js.Rewrite("$<.DomainLookupStart")
	return domainLookupStart
}

// DomComplete
func (*PerformanceTiming) DomComplete() (domComplete uint64) {
	js.Rewrite("$<.DomComplete")
	return domComplete
}

// DomContentLoadedEventEnd
func (*PerformanceTiming) DomContentLoadedEventEnd() (domContentLoadedEventEnd uint64) {
	js.Rewrite("$<.DomContentLoadedEventEnd")
	return domContentLoadedEventEnd
}

// DomContentLoadedEventStart
func (*PerformanceTiming) DomContentLoadedEventStart() (domContentLoadedEventStart uint64) {
	js.Rewrite("$<.DomContentLoadedEventStart")
	return domContentLoadedEventStart
}

// DomInteractive
func (*PerformanceTiming) DomInteractive() (domInteractive uint64) {
	js.Rewrite("$<.DomInteractive")
	return domInteractive
}

// DomLoading
func (*PerformanceTiming) DomLoading() (domLoading uint64) {
	js.Rewrite("$<.DomLoading")
	return domLoading
}

// FetchStart
func (*PerformanceTiming) FetchStart() (fetchStart uint64) {
	js.Rewrite("$<.FetchStart")
	return fetchStart
}

// LoadEventEnd
func (*PerformanceTiming) LoadEventEnd() (loadEventEnd uint64) {
	js.Rewrite("$<.LoadEventEnd")
	return loadEventEnd
}

// LoadEventStart
func (*PerformanceTiming) LoadEventStart() (loadEventStart uint64) {
	js.Rewrite("$<.LoadEventStart")
	return loadEventStart
}

// MsFirstPaint
func (*PerformanceTiming) MsFirstPaint() (msFirstPaint uint64) {
	js.Rewrite("$<.MsFirstPaint")
	return msFirstPaint
}

// NavigationStart
func (*PerformanceTiming) NavigationStart() (navigationStart uint64) {
	js.Rewrite("$<.NavigationStart")
	return navigationStart
}

// RedirectEnd
func (*PerformanceTiming) RedirectEnd() (redirectEnd uint64) {
	js.Rewrite("$<.RedirectEnd")
	return redirectEnd
}

// RedirectStart
func (*PerformanceTiming) RedirectStart() (redirectStart uint64) {
	js.Rewrite("$<.RedirectStart")
	return redirectStart
}

// RequestStart
func (*PerformanceTiming) RequestStart() (requestStart uint64) {
	js.Rewrite("$<.RequestStart")
	return requestStart
}

// ResponseEnd
func (*PerformanceTiming) ResponseEnd() (responseEnd uint64) {
	js.Rewrite("$<.ResponseEnd")
	return responseEnd
}

// ResponseStart
func (*PerformanceTiming) ResponseStart() (responseStart uint64) {
	js.Rewrite("$<.ResponseStart")
	return responseStart
}

// UnloadEventEnd
func (*PerformanceTiming) UnloadEventEnd() (unloadEventEnd uint64) {
	js.Rewrite("$<.UnloadEventEnd")
	return unloadEventEnd
}

// UnloadEventStart
func (*PerformanceTiming) UnloadEventStart() (unloadEventStart uint64) {
	js.Rewrite("$<.UnloadEventStart")
	return unloadEventStart
}
