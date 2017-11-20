package performancenavigationtiming

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/navigationtype"
	"github.com/matthewmueller/golly/internal/gendom/dom/performanceentry"
	"github.com/matthewmueller/golly/js"
)

type PerformanceNavigationTiming struct {
	*performanceentry.PerformanceEntry
}

func (*PerformanceNavigationTiming) GetConnectEnd() (connectEnd int) {
	js.Rewrite("$<.connectEnd")
	return connectEnd
}

func (*PerformanceNavigationTiming) GetConnectStart() (connectStart int) {
	js.Rewrite("$<.connectStart")
	return connectStart
}

func (*PerformanceNavigationTiming) GetDomainLookupEnd() (domainLookupEnd int) {
	js.Rewrite("$<.domainLookupEnd")
	return domainLookupEnd
}

func (*PerformanceNavigationTiming) GetDomainLookupStart() (domainLookupStart int) {
	js.Rewrite("$<.domainLookupStart")
	return domainLookupStart
}

func (*PerformanceNavigationTiming) GetDomComplete() (domComplete int) {
	js.Rewrite("$<.domComplete")
	return domComplete
}

func (*PerformanceNavigationTiming) GetDomContentLoadedEventEnd() (domContentLoadedEventEnd int) {
	js.Rewrite("$<.domContentLoadedEventEnd")
	return domContentLoadedEventEnd
}

func (*PerformanceNavigationTiming) GetDomContentLoadedEventStart() (domContentLoadedEventStart int) {
	js.Rewrite("$<.domContentLoadedEventStart")
	return domContentLoadedEventStart
}

func (*PerformanceNavigationTiming) GetDomInteractive() (domInteractive int) {
	js.Rewrite("$<.domInteractive")
	return domInteractive
}

func (*PerformanceNavigationTiming) GetDomLoading() (domLoading int) {
	js.Rewrite("$<.domLoading")
	return domLoading
}

func (*PerformanceNavigationTiming) GetFetchStart() (fetchStart int) {
	js.Rewrite("$<.fetchStart")
	return fetchStart
}

func (*PerformanceNavigationTiming) GetLoadEventEnd() (loadEventEnd int) {
	js.Rewrite("$<.loadEventEnd")
	return loadEventEnd
}

func (*PerformanceNavigationTiming) GetLoadEventStart() (loadEventStart int) {
	js.Rewrite("$<.loadEventStart")
	return loadEventStart
}

func (*PerformanceNavigationTiming) GetNavigationStart() (navigationStart int) {
	js.Rewrite("$<.navigationStart")
	return navigationStart
}

func (*PerformanceNavigationTiming) GetRedirectCount() (redirectCount uint8) {
	js.Rewrite("$<.redirectCount")
	return redirectCount
}

func (*PerformanceNavigationTiming) GetRedirectEnd() (redirectEnd int) {
	js.Rewrite("$<.redirectEnd")
	return redirectEnd
}

func (*PerformanceNavigationTiming) GetRedirectStart() (redirectStart int) {
	js.Rewrite("$<.redirectStart")
	return redirectStart
}

func (*PerformanceNavigationTiming) GetRequestStart() (requestStart int) {
	js.Rewrite("$<.requestStart")
	return requestStart
}

func (*PerformanceNavigationTiming) GetResponseEnd() (responseEnd int) {
	js.Rewrite("$<.responseEnd")
	return responseEnd
}

func (*PerformanceNavigationTiming) GetResponseStart() (responseStart int) {
	js.Rewrite("$<.responseStart")
	return responseStart
}

func (*PerformanceNavigationTiming) GetType() (kind *navigationtype.NavigationType) {
	js.Rewrite("$<.type")
	return kind
}

func (*PerformanceNavigationTiming) GetUnloadEventEnd() (unloadEventEnd int) {
	js.Rewrite("$<.unloadEventEnd")
	return unloadEventEnd
}

func (*PerformanceNavigationTiming) GetUnloadEventStart() (unloadEventStart int) {
	js.Rewrite("$<.unloadEventStart")
	return unloadEventStart
}
