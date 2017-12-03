package performancenavigationtiming

import (
	"github.com/matthewmueller/joy/dom/navigationtype"
	"github.com/matthewmueller/joy/dom/performanceentry"
	"github.com/matthewmueller/joy/macro"
)

var _ performanceentry.PerformanceEntry = (*PerformanceNavigationTiming)(nil)

// PerformanceNavigationTiming struct
// js:"PerformanceNavigationTiming,omit"
type PerformanceNavigationTiming struct {
}

// ConnectEnd prop
// js:"connectEnd"
func (*PerformanceNavigationTiming) ConnectEnd() (connectEnd int) {
	macro.Rewrite("$_.connectEnd")
	return connectEnd
}

// ConnectStart prop
// js:"connectStart"
func (*PerformanceNavigationTiming) ConnectStart() (connectStart int) {
	macro.Rewrite("$_.connectStart")
	return connectStart
}

// DomainLookupEnd prop
// js:"domainLookupEnd"
func (*PerformanceNavigationTiming) DomainLookupEnd() (domainLookupEnd int) {
	macro.Rewrite("$_.domainLookupEnd")
	return domainLookupEnd
}

// DomainLookupStart prop
// js:"domainLookupStart"
func (*PerformanceNavigationTiming) DomainLookupStart() (domainLookupStart int) {
	macro.Rewrite("$_.domainLookupStart")
	return domainLookupStart
}

// DomComplete prop
// js:"domComplete"
func (*PerformanceNavigationTiming) DomComplete() (domComplete int) {
	macro.Rewrite("$_.domComplete")
	return domComplete
}

// DomContentLoadedEventEnd prop
// js:"domContentLoadedEventEnd"
func (*PerformanceNavigationTiming) DomContentLoadedEventEnd() (domContentLoadedEventEnd int) {
	macro.Rewrite("$_.domContentLoadedEventEnd")
	return domContentLoadedEventEnd
}

// DomContentLoadedEventStart prop
// js:"domContentLoadedEventStart"
func (*PerformanceNavigationTiming) DomContentLoadedEventStart() (domContentLoadedEventStart int) {
	macro.Rewrite("$_.domContentLoadedEventStart")
	return domContentLoadedEventStart
}

// DomInteractive prop
// js:"domInteractive"
func (*PerformanceNavigationTiming) DomInteractive() (domInteractive int) {
	macro.Rewrite("$_.domInteractive")
	return domInteractive
}

// DomLoading prop
// js:"domLoading"
func (*PerformanceNavigationTiming) DomLoading() (domLoading int) {
	macro.Rewrite("$_.domLoading")
	return domLoading
}

// FetchStart prop
// js:"fetchStart"
func (*PerformanceNavigationTiming) FetchStart() (fetchStart int) {
	macro.Rewrite("$_.fetchStart")
	return fetchStart
}

// LoadEventEnd prop
// js:"loadEventEnd"
func (*PerformanceNavigationTiming) LoadEventEnd() (loadEventEnd int) {
	macro.Rewrite("$_.loadEventEnd")
	return loadEventEnd
}

// LoadEventStart prop
// js:"loadEventStart"
func (*PerformanceNavigationTiming) LoadEventStart() (loadEventStart int) {
	macro.Rewrite("$_.loadEventStart")
	return loadEventStart
}

// NavigationStart prop
// js:"navigationStart"
func (*PerformanceNavigationTiming) NavigationStart() (navigationStart int) {
	macro.Rewrite("$_.navigationStart")
	return navigationStart
}

// RedirectCount prop
// js:"redirectCount"
func (*PerformanceNavigationTiming) RedirectCount() (redirectCount uint8) {
	macro.Rewrite("$_.redirectCount")
	return redirectCount
}

// RedirectEnd prop
// js:"redirectEnd"
func (*PerformanceNavigationTiming) RedirectEnd() (redirectEnd int) {
	macro.Rewrite("$_.redirectEnd")
	return redirectEnd
}

// RedirectStart prop
// js:"redirectStart"
func (*PerformanceNavigationTiming) RedirectStart() (redirectStart int) {
	macro.Rewrite("$_.redirectStart")
	return redirectStart
}

// RequestStart prop
// js:"requestStart"
func (*PerformanceNavigationTiming) RequestStart() (requestStart int) {
	macro.Rewrite("$_.requestStart")
	return requestStart
}

// ResponseEnd prop
// js:"responseEnd"
func (*PerformanceNavigationTiming) ResponseEnd() (responseEnd int) {
	macro.Rewrite("$_.responseEnd")
	return responseEnd
}

// ResponseStart prop
// js:"responseStart"
func (*PerformanceNavigationTiming) ResponseStart() (responseStart int) {
	macro.Rewrite("$_.responseStart")
	return responseStart
}

// Type prop
// js:"type"
func (*PerformanceNavigationTiming) Type() (kind *navigationtype.NavigationType) {
	macro.Rewrite("$_.type")
	return kind
}

// UnloadEventEnd prop
// js:"unloadEventEnd"
func (*PerformanceNavigationTiming) UnloadEventEnd() (unloadEventEnd int) {
	macro.Rewrite("$_.unloadEventEnd")
	return unloadEventEnd
}

// UnloadEventStart prop
// js:"unloadEventStart"
func (*PerformanceNavigationTiming) UnloadEventStart() (unloadEventStart int) {
	macro.Rewrite("$_.unloadEventStart")
	return unloadEventStart
}

// Duration prop
// js:"duration"
func (*PerformanceNavigationTiming) Duration() (duration int) {
	macro.Rewrite("$_.duration")
	return duration
}

// EntryType prop
// js:"entryType"
func (*PerformanceNavigationTiming) EntryType() (entryType string) {
	macro.Rewrite("$_.entryType")
	return entryType
}

// Name prop
// js:"name"
func (*PerformanceNavigationTiming) Name() (name string) {
	macro.Rewrite("$_.name")
	return name
}

// StartTime prop
// js:"startTime"
func (*PerformanceNavigationTiming) StartTime() (startTime int) {
	macro.Rewrite("$_.startTime")
	return startTime
}
