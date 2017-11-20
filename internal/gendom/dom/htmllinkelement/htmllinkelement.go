package htmllinkelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/linkstyle"
	"github.com/matthewmueller/golly/js"
)

type HTMLLinkElement struct {
	*htmlelement.HTMLElement
	*linkstyle.LinkStyle
}

func (*HTMLLinkElement) GetCharset() (charset string) {
	js.Rewrite("$<.charset")
	return charset
}

func (*HTMLLinkElement) SetCharset(charset string) {
	js.Rewrite("$<.charset = $1", charset)
}

func (*HTMLLinkElement) GetDisabled() (disabled bool) {
	js.Rewrite("$<.disabled")
	return disabled
}

func (*HTMLLinkElement) SetDisabled(disabled bool) {
	js.Rewrite("$<.disabled = $1", disabled)
}

func (*HTMLLinkElement) GetHref() (href string) {
	js.Rewrite("$<.href")
	return href
}

func (*HTMLLinkElement) SetHref(href string) {
	js.Rewrite("$<.href = $1", href)
}

func (*HTMLLinkElement) GetHreflang() (hreflang string) {
	js.Rewrite("$<.hreflang")
	return hreflang
}

func (*HTMLLinkElement) SetHreflang(hreflang string) {
	js.Rewrite("$<.hreflang = $1", hreflang)
}

func (*HTMLLinkElement) GetMedia() (media string) {
	js.Rewrite("$<.media")
	return media
}

func (*HTMLLinkElement) SetMedia(media string) {
	js.Rewrite("$<.media = $1", media)
}

func (*HTMLLinkElement) GetRel() (rel string) {
	js.Rewrite("$<.rel")
	return rel
}

func (*HTMLLinkElement) SetRel(rel string) {
	js.Rewrite("$<.rel = $1", rel)
}

func (*HTMLLinkElement) GetRev() (rev string) {
	js.Rewrite("$<.rev")
	return rev
}

func (*HTMLLinkElement) SetRev(rev string) {
	js.Rewrite("$<.rev = $1", rev)
}

func (*HTMLLinkElement) GetTarget() (target string) {
	js.Rewrite("$<.target")
	return target
}

func (*HTMLLinkElement) SetTarget(target string) {
	js.Rewrite("$<.target = $1", target)
}

func (*HTMLLinkElement) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

func (*HTMLLinkElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}
