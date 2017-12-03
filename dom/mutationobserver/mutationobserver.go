package mutationobserver

import (
	"github.com/matthewmueller/joy/dom/mutationobserverinit"
	"github.com/matthewmueller/joy/dom/mutationrecord"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/js"
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
// js:"disconnect"
func (*MutationObserver) Disconnect() {
	js.Rewrite("$_.disconnect()")
}

// Observe fn
// js:"observe"
func (*MutationObserver) Observe(target window.Node, options *mutationobserverinit.MutationObserverInit) {
	js.Rewrite("$_.observe($1, $2)", target, options)
}

// TakeRecords fn
// js:"takeRecords"
func (*MutationObserver) TakeRecords() (m []*mutationrecord.MutationRecord) {
	js.Rewrite("$_.takeRecords()")
	return m
}
