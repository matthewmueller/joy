package applicationcache

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/progressevent"
	"github.com/matthewmueller/golly/js"
)

type ApplicationCache struct {
	*eventtarget.EventTarget
}

func (*ApplicationCache) Abort() {
	js.Rewrite("$<.abort()")
}

func (*ApplicationCache) SwapCache() {
	js.Rewrite("$<.swapCache()")
}

func (*ApplicationCache) Update() {
	js.Rewrite("$<.update()")
}

func (*ApplicationCache) GetOncached() (cached *event.Event) {
	js.Rewrite("$<.oncached")
	return cached
}

func (*ApplicationCache) SetOncached(cached *event.Event) {
	js.Rewrite("$<.oncached = $1", cached)
}

func (*ApplicationCache) GetOnchecking() (checking *event.Event) {
	js.Rewrite("$<.onchecking")
	return checking
}

func (*ApplicationCache) SetOnchecking(checking *event.Event) {
	js.Rewrite("$<.onchecking = $1", checking)
}

func (*ApplicationCache) GetOndownloading() (downloading *event.Event) {
	js.Rewrite("$<.ondownloading")
	return downloading
}

func (*ApplicationCache) SetOndownloading(downloading *event.Event) {
	js.Rewrite("$<.ondownloading = $1", downloading)
}

func (*ApplicationCache) GetOnerror() (err *event.Event) {
	js.Rewrite("$<.onerror")
	return err
}

func (*ApplicationCache) SetOnerror(err *event.Event) {
	js.Rewrite("$<.onerror = $1", err)
}

func (*ApplicationCache) GetOnnoupdate() (noupdate *event.Event) {
	js.Rewrite("$<.onnoupdate")
	return noupdate
}

func (*ApplicationCache) SetOnnoupdate(noupdate *event.Event) {
	js.Rewrite("$<.onnoupdate = $1", noupdate)
}

func (*ApplicationCache) GetOnobsolete() (obsolete *event.Event) {
	js.Rewrite("$<.onobsolete")
	return obsolete
}

func (*ApplicationCache) SetOnobsolete(obsolete *event.Event) {
	js.Rewrite("$<.onobsolete = $1", obsolete)
}

func (*ApplicationCache) GetOnprogress() (progress *progressevent.ProgressEvent) {
	js.Rewrite("$<.onprogress")
	return progress
}

func (*ApplicationCache) SetOnprogress(progress *progressevent.ProgressEvent) {
	js.Rewrite("$<.onprogress = $1", progress)
}

func (*ApplicationCache) GetOnupdateready() (updateready *event.Event) {
	js.Rewrite("$<.onupdateready")
	return updateready
}

func (*ApplicationCache) SetOnupdateready(updateready *event.Event) {
	js.Rewrite("$<.onupdateready = $1", updateready)
}

func (*ApplicationCache) GetStatus() (status uint8) {
	js.Rewrite("$<.status")
	return status
}
