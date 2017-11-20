package performanceresourcetiming

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/performanceentry"
	"github.com/matthewmueller/golly/js"
)

type PerformanceResourceTiming struct {
	*performanceentry.PerformanceEntry
}

func (*PerformanceResourceTiming) GetConnectEnd() (connectEnd int) {
	js.Rewrite("$<.connectEnd")
	return connectEnd
}

func (*PerformanceResourceTiming) GetConnectStart() (connectStart int) {
	js.Rewrite("$<.connectStart")
	return connectStart
}

func (*PerformanceResourceTiming) GetDomainLookupEnd() (domainLookupEnd int) {
	js.Rewrite("$<.domainLookupEnd")
	return domainLookupEnd
}

func (*PerformanceResourceTiming) GetDomainLookupStart() (domainLookupStart int) {
	js.Rewrite("$<.domainLookupStart")
	return domainLookupStart
}

func (*PerformanceResourceTiming) GetFetchStart() (fetchStart int) {
	js.Rewrite("$<.fetchStart")
	return fetchStart
}

func (*PerformanceResourceTiming) GetInitiatorType() (initiatorType string) {
	js.Rewrite("$<.initiatorType")
	return initiatorType
}

func (*PerformanceResourceTiming) GetRedirectEnd() (redirectEnd int) {
	js.Rewrite("$<.redirectEnd")
	return redirectEnd
}

func (*PerformanceResourceTiming) GetRedirectStart() (redirectStart int) {
	js.Rewrite("$<.redirectStart")
	return redirectStart
}

func (*PerformanceResourceTiming) GetRequestStart() (requestStart int) {
	js.Rewrite("$<.requestStart")
	return requestStart
}

func (*PerformanceResourceTiming) GetResponseEnd() (responseEnd int) {
	js.Rewrite("$<.responseEnd")
	return responseEnd
}

func (*PerformanceResourceTiming) GetResponseStart() (responseStart int) {
	js.Rewrite("$<.responseStart")
	return responseStart
}
