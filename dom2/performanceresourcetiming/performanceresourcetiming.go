package performanceresourcetiming

import (
	"github.com/matthewmueller/golly/dom2/performanceentry"
	"github.com/matthewmueller/golly/js"
)

// js:"PerformanceResourceTiming,omit"
type PerformanceResourceTiming struct {
	performanceentry.PerformanceEntry
}

// ConnectEnd
func (*PerformanceResourceTiming) ConnectEnd() (connectEnd int) {
	js.Rewrite("$<.ConnectEnd")
	return connectEnd
}

// ConnectStart
func (*PerformanceResourceTiming) ConnectStart() (connectStart int) {
	js.Rewrite("$<.ConnectStart")
	return connectStart
}

// DomainLookupEnd
func (*PerformanceResourceTiming) DomainLookupEnd() (domainLookupEnd int) {
	js.Rewrite("$<.DomainLookupEnd")
	return domainLookupEnd
}

// DomainLookupStart
func (*PerformanceResourceTiming) DomainLookupStart() (domainLookupStart int) {
	js.Rewrite("$<.DomainLookupStart")
	return domainLookupStart
}

// FetchStart
func (*PerformanceResourceTiming) FetchStart() (fetchStart int) {
	js.Rewrite("$<.FetchStart")
	return fetchStart
}

// InitiatorType
func (*PerformanceResourceTiming) InitiatorType() (initiatorType string) {
	js.Rewrite("$<.InitiatorType")
	return initiatorType
}

// RedirectEnd
func (*PerformanceResourceTiming) RedirectEnd() (redirectEnd int) {
	js.Rewrite("$<.RedirectEnd")
	return redirectEnd
}

// RedirectStart
func (*PerformanceResourceTiming) RedirectStart() (redirectStart int) {
	js.Rewrite("$<.RedirectStart")
	return redirectStart
}

// RequestStart
func (*PerformanceResourceTiming) RequestStart() (requestStart int) {
	js.Rewrite("$<.RequestStart")
	return requestStart
}

// ResponseEnd
func (*PerformanceResourceTiming) ResponseEnd() (responseEnd int) {
	js.Rewrite("$<.ResponseEnd")
	return responseEnd
}

// ResponseStart
func (*PerformanceResourceTiming) ResponseStart() (responseStart int) {
	js.Rewrite("$<.ResponseStart")
	return responseStart
}
