package mutationobserver

import (
	"github.com/matthewmueller/joy/dom/mutationobserverinit"
	"github.com/matthewmueller/joy/dom/mutationrecord"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

// New fn
func New(callback func(mutations []*mutationrecord.MutationRecord, observer *MutationObserver)) *MutationObserver {
	macro.Rewrite("MutationObserver")
	return &MutationObserver{}
}

// MutationObserver struct
// js:"MutationObserver,omit"
type MutationObserver struct {
}

// Disconnect fn
// js:"disconnect"
func (*MutationObserver) Disconnect() {
	macro.Rewrite("$_.disconnect()")
}

// Observe fn
// js:"observe"
func (*MutationObserver) Observe(target window.Node, options *mutationobserverinit.MutationObserverInit) {
	macro.Rewrite("$_.observe($1, $2)", target, options)
}

// TakeRecords fn
// js:"takeRecords"
func (*MutationObserver) TakeRecords() (m []*mutationrecord.MutationRecord) {
	macro.Rewrite("$_.takeRecords()")
	return m
}
