package performancenavigationtiming

import (
	"github.com/matthewmueller/golly/dom/navigationtype"
	"github.com/matthewmueller/golly/dom/performanceentry"
	"github.com/matthewmueller/golly/js"
)

var _ performanceentry.PerformanceEntry = (*PerformanceNavigationTiming)(nil)

// PerformanceNavigationTiming struct
// js:"PerformanceNavigationTiming,omit"
type PerformanceNavigationTiming struct {
}

// ConnectEnd prop
// js:"connectEnd"
func (*PerformanceNavigationTiming) ConnectEnd() (connectEnd int) {
	js.Rewrite("$_.connectEnd")
	return connectEnd
}

// ConnectStart prop
// js:"connectStart"
func (*PerformanceNavigationTiming) ConnectStart() (connectStart int) {
	js.Rewrite("$_.connectStart")
	return connectStart
}

// DomainLookupEnd prop
// js:"domainLookupEnd"
func (*PerformanceNavigationTiming) DomainLookupEnd() (domainLookupEnd int) {
	js.Rewrite("$_.domainLookupEnd")
	return domainLookupEnd
}

// DomainLookupStart prop
// js:"domainLookupStart"
func (*PerformanceNavigationTiming) DomainLookupStart() (domainLookupStart int) {
	js.Rewrite("$_.domainLookupStart")
	return domainLookupStart
}

// DomComplete prop
// js:"domComplete"
func (*PerformanceNavigationTiming) DomComplete() (domComplete int) {
	js.Rewrite("$_.domComplete")
	return domComplete
}

// DomContentLoadedEventEnd prop
// js:"domContentLoadedEventEnd"
func (*PerformanceNavigationTiming) DomContentLoadedEventEnd() (domContentLoadedEventEnd int) {
	js.Rewrite("$_.domContentLoadedEventEnd")
	return domContentLoadedEventEnd
}

// DomContentLoadedEventStart prop
// js:"domContentLoadedEventStart"
func (*PerformanceNavigationTiming) DomContentLoadedEventStart() (domContentLoadedEventStart int) {
	js.Rewrite("$_.domContentLoadedEventStart")
	return domContentLoadedEventStart
}

// DomInteractive prop
// js:"domInteractive"
func (*PerformanceNavigationTiming) DomInteractive() (domInteractive int) {
	js.Rewrite("$_.domInteractive")
	return domInteractive
}

// DomLoading prop
// js:"domLoading"
func (*PerformanceNavigationTiming) DomLoading() (domLoading int) {
	js.Rewrite("$_.domLoading")
	return domLoading
}

// FetchStart prop
// js:"fetchStart"
func (*PerformanceNavigationTiming) FetchStart() (fetchStart int) {
	js.Rewrite("$_.fetchStart")
	return fetchStart
}

// LoadEventEnd prop
// js:"loadEventEnd"
func (*PerformanceNavigationTiming) LoadEventEnd() (loadEventEnd int) {
	js.Rewrite("$_.loadEventEnd")
	return loadEventEnd
}

// LoadEventStart prop
// js:"loadEventStart"
func (*PerformanceNavigationTiming) LoadEventStart() (loadEventStart int) {
	js.Rewrite("$_.loadEventStart")
	return loadEventStart
}

// NavigationStart prop
// js:"navigationStart"
func (*PerformanceNavigationTiming) NavigationStart() (navigationStart int) {
	js.Rewrite("$_.navigationStart")
	return navigationStart
}

// RedirectCount prop
// js:"redirectCount"
func (*PerformanceNavigationTiming) RedirectCount() (redirectCount uint8) {
	js.Rewrite("$_.redirectCount")
	return redirectCount
}

// RedirectEnd prop
// js:"redirectEnd"
func (*PerformanceNavigationTiming) RedirectEnd() (redirectEnd int) {
	js.Rewrite("$_.redirectEnd")
	return redirectEnd
}

// RedirectStart prop
// js:"redirectStart"
func (*PerformanceNavigationTiming) RedirectStart() (redirectStart int) {
	js.Rewrite("$_.redirectStart")
	return redirectStart
}

// RequestStart prop
// js:"requestStart"
func (*PerformanceNavigationTiming) RequestStart() (requestStart int) {
	js.Rewrite("$_.requestStart")
	return requestStart
}

// ResponseEnd prop
// js:"responseEnd"
func (*PerformanceNavigationTiming) ResponseEnd() (responseEnd int) {
	js.Rewrite("$_.responseEnd")
	return responseEnd
}

// ResponseStart prop
// js:"responseStart"
func (*PerformanceNavigationTiming) ResponseStart() (responseStart int) {
	js.Rewrite("$_.responseStart")
	return responseStart
}

// Type prop
// js:"type"
func (*PerformanceNavigationTiming) Type() (kind *navigationtype.NavigationType) {
	js.Rewrite("$_.type")
	return kind
}

// UnloadEventEnd prop
// js:"unloadEventEnd"
func (*PerformanceNavigationTiming) UnloadEventEnd() (unloadEventEnd int) {
	js.Rewrite("$_.unloadEventEnd")
	return unloadEventEnd
}

// UnloadEventStart prop
// js:"unloadEventStart"
func (*PerformanceNavigationTiming) UnloadEventStart() (unloadEventStart int) {
	js.Rewrite("$_.unloadEventStart")
	return unloadEventStart
}

// Duration prop
// js:"duration"
func (*PerformanceNavigationTiming) Duration() (duration int) {
	js.Rewrite("$_.duration")
	return duration
}

// EntryType prop
// js:"entryType"
func (*PerformanceNavigationTiming) EntryType() (entryType string) {
	js.Rewrite("$_.entryType")
	return entryType
}

// Name prop
// js:"name"
func (*PerformanceNavigationTiming) Name() (name string) {
	js.Rewrite("$_.name")
	return name
}

// StartTime prop
// js:"startTime"
func (*PerformanceNavigationTiming) StartTime() (startTime int) {
	js.Rewrite("$_.startTime")
	return startTime
}
