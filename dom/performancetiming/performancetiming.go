package performancetiming

import "github.com/matthewmueller/joy/js"

// PerformanceTiming struct
// js:"PerformanceTiming,omit"
type PerformanceTiming struct {
}

// ToJSON fn
// js:"toJSON"
func (*PerformanceTiming) ToJSON() (i interface{}) {
	js.Rewrite("$_.toJSON()")
	return i
}

// ConnectEnd prop
// js:"connectEnd"
func (*PerformanceTiming) ConnectEnd() (connectEnd uint64) {
	js.Rewrite("$_.connectEnd")
	return connectEnd
}

// ConnectStart prop
// js:"connectStart"
func (*PerformanceTiming) ConnectStart() (connectStart uint64) {
	js.Rewrite("$_.connectStart")
	return connectStart
}

// DomainLookupEnd prop
// js:"domainLookupEnd"
func (*PerformanceTiming) DomainLookupEnd() (domainLookupEnd uint64) {
	js.Rewrite("$_.domainLookupEnd")
	return domainLookupEnd
}

// DomainLookupStart prop
// js:"domainLookupStart"
func (*PerformanceTiming) DomainLookupStart() (domainLookupStart uint64) {
	js.Rewrite("$_.domainLookupStart")
	return domainLookupStart
}

// DomComplete prop
// js:"domComplete"
func (*PerformanceTiming) DomComplete() (domComplete uint64) {
	js.Rewrite("$_.domComplete")
	return domComplete
}

// DomContentLoadedEventEnd prop
// js:"domContentLoadedEventEnd"
func (*PerformanceTiming) DomContentLoadedEventEnd() (domContentLoadedEventEnd uint64) {
	js.Rewrite("$_.domContentLoadedEventEnd")
	return domContentLoadedEventEnd
}

// DomContentLoadedEventStart prop
// js:"domContentLoadedEventStart"
func (*PerformanceTiming) DomContentLoadedEventStart() (domContentLoadedEventStart uint64) {
	js.Rewrite("$_.domContentLoadedEventStart")
	return domContentLoadedEventStart
}

// DomInteractive prop
// js:"domInteractive"
func (*PerformanceTiming) DomInteractive() (domInteractive uint64) {
	js.Rewrite("$_.domInteractive")
	return domInteractive
}

// DomLoading prop
// js:"domLoading"
func (*PerformanceTiming) DomLoading() (domLoading uint64) {
	js.Rewrite("$_.domLoading")
	return domLoading
}

// FetchStart prop
// js:"fetchStart"
func (*PerformanceTiming) FetchStart() (fetchStart uint64) {
	js.Rewrite("$_.fetchStart")
	return fetchStart
}

// LoadEventEnd prop
// js:"loadEventEnd"
func (*PerformanceTiming) LoadEventEnd() (loadEventEnd uint64) {
	js.Rewrite("$_.loadEventEnd")
	return loadEventEnd
}

// LoadEventStart prop
// js:"loadEventStart"
func (*PerformanceTiming) LoadEventStart() (loadEventStart uint64) {
	js.Rewrite("$_.loadEventStart")
	return loadEventStart
}

// MsFirstPaint prop
// js:"msFirstPaint"
func (*PerformanceTiming) MsFirstPaint() (msFirstPaint uint64) {
	js.Rewrite("$_.msFirstPaint")
	return msFirstPaint
}

// NavigationStart prop
// js:"navigationStart"
func (*PerformanceTiming) NavigationStart() (navigationStart uint64) {
	js.Rewrite("$_.navigationStart")
	return navigationStart
}

// RedirectEnd prop
// js:"redirectEnd"
func (*PerformanceTiming) RedirectEnd() (redirectEnd uint64) {
	js.Rewrite("$_.redirectEnd")
	return redirectEnd
}

// RedirectStart prop
// js:"redirectStart"
func (*PerformanceTiming) RedirectStart() (redirectStart uint64) {
	js.Rewrite("$_.redirectStart")
	return redirectStart
}

// RequestStart prop
// js:"requestStart"
func (*PerformanceTiming) RequestStart() (requestStart uint64) {
	js.Rewrite("$_.requestStart")
	return requestStart
}

// ResponseEnd prop
// js:"responseEnd"
func (*PerformanceTiming) ResponseEnd() (responseEnd uint64) {
	js.Rewrite("$_.responseEnd")
	return responseEnd
}

// ResponseStart prop
// js:"responseStart"
func (*PerformanceTiming) ResponseStart() (responseStart uint64) {
	js.Rewrite("$_.responseStart")
	return responseStart
}

// UnloadEventEnd prop
// js:"unloadEventEnd"
func (*PerformanceTiming) UnloadEventEnd() (unloadEventEnd uint64) {
	js.Rewrite("$_.unloadEventEnd")
	return unloadEventEnd
}

// UnloadEventStart prop
// js:"unloadEventStart"
func (*PerformanceTiming) UnloadEventStart() (unloadEventStart uint64) {
	js.Rewrite("$_.unloadEventStart")
	return unloadEventStart
}
