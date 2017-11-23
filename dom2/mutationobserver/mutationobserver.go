package mutationobserver

import (
	"github.com/matthewmueller/golly/dom2/mutationobserverinit"
	"github.com/matthewmueller/golly/dom2/mutationrecord"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"MutationObserver,omit"
type MutationObserver struct {
}

// Disconnect
func (*MutationObserver) Disconnect() {
	js.Rewrite("$<.Disconnect()")
}

// Observe
func (*MutationObserver) Observe(target window.Node, options *mutationobserverinit.MutationObserverInit) {
	js.Rewrite("$<.Observe($1, $2)", target, options)
}

// TakeRecords
func (*MutationObserver) TakeRecords() (m []*mutationrecord.MutationRecord) {
	js.Rewrite("$<.TakeRecords()")
	return m
}
