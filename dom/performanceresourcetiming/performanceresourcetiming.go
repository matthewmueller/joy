package performanceresourcetiming

import (
	"github.com/matthewmueller/joy/dom/performanceentry"
	"github.com/matthewmueller/joy/js"
)

var _ performanceentry.PerformanceEntry = (*PerformanceResourceTiming)(nil)

// PerformanceResourceTiming struct
// js:"PerformanceResourceTiming,omit"
type PerformanceResourceTiming struct {
}

// ConnectEnd prop
// js:"connectEnd"
func (*PerformanceResourceTiming) ConnectEnd() (connectEnd int) {
	js.Rewrite("$_.connectEnd")
	return connectEnd
}

// ConnectStart prop
// js:"connectStart"
func (*PerformanceResourceTiming) ConnectStart() (connectStart int) {
	js.Rewrite("$_.connectStart")
	return connectStart
}

// DomainLookupEnd prop
// js:"domainLookupEnd"
func (*PerformanceResourceTiming) DomainLookupEnd() (domainLookupEnd int) {
	js.Rewrite("$_.domainLookupEnd")
	return domainLookupEnd
}

// DomainLookupStart prop
// js:"domainLookupStart"
func (*PerformanceResourceTiming) DomainLookupStart() (domainLookupStart int) {
	js.Rewrite("$_.domainLookupStart")
	return domainLookupStart
}

// FetchStart prop
// js:"fetchStart"
func (*PerformanceResourceTiming) FetchStart() (fetchStart int) {
	js.Rewrite("$_.fetchStart")
	return fetchStart
}

// InitiatorType prop
// js:"initiatorType"
func (*PerformanceResourceTiming) InitiatorType() (initiatorType string) {
	js.Rewrite("$_.initiatorType")
	return initiatorType
}

// RedirectEnd prop
// js:"redirectEnd"
func (*PerformanceResourceTiming) RedirectEnd() (redirectEnd int) {
	js.Rewrite("$_.redirectEnd")
	return redirectEnd
}

// RedirectStart prop
// js:"redirectStart"
func (*PerformanceResourceTiming) RedirectStart() (redirectStart int) {
	js.Rewrite("$_.redirectStart")
	return redirectStart
}

// RequestStart prop
// js:"requestStart"
func (*PerformanceResourceTiming) RequestStart() (requestStart int) {
	js.Rewrite("$_.requestStart")
	return requestStart
}

// ResponseEnd prop
// js:"responseEnd"
func (*PerformanceResourceTiming) ResponseEnd() (responseEnd int) {
	js.Rewrite("$_.responseEnd")
	return responseEnd
}

// ResponseStart prop
// js:"responseStart"
func (*PerformanceResourceTiming) ResponseStart() (responseStart int) {
	js.Rewrite("$_.responseStart")
	return responseStart
}

// Duration prop
// js:"duration"
func (*PerformanceResourceTiming) Duration() (duration int) {
	js.Rewrite("$_.duration")
	return duration
}

// EntryType prop
// js:"entryType"
func (*PerformanceResourceTiming) EntryType() (entryType string) {
	js.Rewrite("$_.entryType")
	return entryType
}

// Name prop
// js:"name"
func (*PerformanceResourceTiming) Name() (name string) {
	js.Rewrite("$_.name")
	return name
}

// StartTime prop
// js:"startTime"
func (*PerformanceResourceTiming) StartTime() (startTime int) {
	js.Rewrite("$_.startTime")
	return startTime
}
