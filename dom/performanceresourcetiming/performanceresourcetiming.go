package performanceresourcetiming

import (
	"github.com/matthewmueller/golly/dom/performanceentry"
	"github.com/matthewmueller/golly/js"
)

// PerformanceResourceTiming struct
// js:"PerformanceResourceTiming,omit"
type PerformanceResourceTiming struct {
	performanceentry.PerformanceEntry
}

// ConnectEnd prop
func (*PerformanceResourceTiming) ConnectEnd() (connectEnd int) {
	js.Rewrite("$<.connectEnd")
	return connectEnd
}

// ConnectStart prop
func (*PerformanceResourceTiming) ConnectStart() (connectStart int) {
	js.Rewrite("$<.connectStart")
	return connectStart
}

// DomainLookupEnd prop
func (*PerformanceResourceTiming) DomainLookupEnd() (domainLookupEnd int) {
	js.Rewrite("$<.domainLookupEnd")
	return domainLookupEnd
}

// DomainLookupStart prop
func (*PerformanceResourceTiming) DomainLookupStart() (domainLookupStart int) {
	js.Rewrite("$<.domainLookupStart")
	return domainLookupStart
}

// FetchStart prop
func (*PerformanceResourceTiming) FetchStart() (fetchStart int) {
	js.Rewrite("$<.fetchStart")
	return fetchStart
}

// InitiatorType prop
func (*PerformanceResourceTiming) InitiatorType() (initiatorType string) {
	js.Rewrite("$<.initiatorType")
	return initiatorType
}

// RedirectEnd prop
func (*PerformanceResourceTiming) RedirectEnd() (redirectEnd int) {
	js.Rewrite("$<.redirectEnd")
	return redirectEnd
}

// RedirectStart prop
func (*PerformanceResourceTiming) RedirectStart() (redirectStart int) {
	js.Rewrite("$<.redirectStart")
	return redirectStart
}

// RequestStart prop
func (*PerformanceResourceTiming) RequestStart() (requestStart int) {
	js.Rewrite("$<.requestStart")
	return requestStart
}

// ResponseEnd prop
func (*PerformanceResourceTiming) ResponseEnd() (responseEnd int) {
	js.Rewrite("$<.responseEnd")
	return responseEnd
}

// ResponseStart prop
func (*PerformanceResourceTiming) ResponseStart() (responseStart int) {
	js.Rewrite("$<.responseStart")
	return responseStart
}
