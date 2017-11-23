package mutationobserver

import (
	"github.com/matthewmueller/golly/dom2/mutationobserverinit"
	"github.com/matthewmueller/golly/dom2/mutationrecord"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New(callback func(mutations []*mutationrecord.MutationRecord, observer *MutationObserver)) *MutationObserver {
	js.Rewrite("MutationObserver")
	return &MutationObserver{}
}

// MutationObserver struct
// js:"MutationObserver,omit"
type MutationObserver struct {
}

// Disconnect fn
func (*MutationObserver) Disconnect() {
	js.Rewrite("$<.disconnect()")
}

// Observe fn
func (*MutationObserver) Observe(target window.Node, options *mutationobserverinit.MutationObserverInit) {
	js.Rewrite("$<.observe($1, $2)", target, options)
}

// TakeRecords fn
func (*MutationObserver) TakeRecords() (m []*mutationrecord.MutationRecord) {
	js.Rewrite("$<.takeRecords()")
	return m
}
