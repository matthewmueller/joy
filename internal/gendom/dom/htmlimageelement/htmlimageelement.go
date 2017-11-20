package htmlimageelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/js"
)

type HTMLImageElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLImageElement) MsGetAsCastingSource() (i interface{}) {
	js.Rewrite("$<.msGetAsCastingSource()")
	return i
}

func (*HTMLImageElement) GetAlign() (align string) {
	js.Rewrite("$<.align")
	return align
}

func (*HTMLImageElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

func (*HTMLImageElement) GetAlt() (alt string) {
	js.Rewrite("$<.alt")
	return alt
}

func (*HTMLImageElement) SetAlt(alt string) {
	js.Rewrite("$<.alt = $1", alt)
}

func (*HTMLImageElement) GetBorder() (border string) {
	js.Rewrite("$<.border")
	return border
}

func (*HTMLImageElement) SetBorder(border string) {
	js.Rewrite("$<.border = $1", border)
}

func (*HTMLImageElement) GetComplete() (complete bool) {
	js.Rewrite("$<.complete")
	return complete
}

func (*HTMLImageElement) GetCrossOrigin() (crossOrigin string) {
	js.Rewrite("$<.crossOrigin")
	return crossOrigin
}

func (*HTMLImageElement) SetCrossOrigin(crossOrigin string) {
	js.Rewrite("$<.crossOrigin = $1", crossOrigin)
}

func (*HTMLImageElement) GetCurrentSrc() (currentSrc string) {
	js.Rewrite("$<.currentSrc")
	return currentSrc
}

func (*HTMLImageElement) GetHeight() (height uint) {
	js.Rewrite("$<.height")
	return height
}

func (*HTMLImageElement) SetHeight(height uint) {
	js.Rewrite("$<.height = $1", height)
}

func (*HTMLImageElement) GetHspace() (hspace int) {
	js.Rewrite("$<.hspace")
	return hspace
}

func (*HTMLImageElement) SetHspace(hspace int) {
	js.Rewrite("$<.hspace = $1", hspace)
}

func (*HTMLImageElement) GetIsMap() (isMap bool) {
	js.Rewrite("$<.isMap")
	return isMap
}

func (*HTMLImageElement) SetIsMap(isMap bool) {
	js.Rewrite("$<.isMap = $1", isMap)
}

func (*HTMLImageElement) GetLongDesc() (longDesc string) {
	js.Rewrite("$<.longDesc")
	return longDesc
}

func (*HTMLImageElement) SetLongDesc(longDesc string) {
	js.Rewrite("$<.longDesc = $1", longDesc)
}

func (*HTMLImageElement) GetLowsrc() (lowsrc string) {
	js.Rewrite("$<.lowsrc")
	return lowsrc
}

func (*HTMLImageElement) SetLowsrc(lowsrc string) {
	js.Rewrite("$<.lowsrc = $1", lowsrc)
}

func (*HTMLImageElement) GetMsPlayToDisabled() (msPlayToDisabled bool) {
	js.Rewrite("$<.msPlayToDisabled")
	return msPlayToDisabled
}

func (*HTMLImageElement) SetMsPlayToDisabled(msPlayToDisabled bool) {
	js.Rewrite("$<.msPlayToDisabled = $1", msPlayToDisabled)
}

func (*HTMLImageElement) GetMsPlayToPreferredSourceURI() (msPlayToPreferredSourceUri string) {
	js.Rewrite("$<.msPlayToPreferredSourceUri")
	return msPlayToPreferredSourceUri
}

func (*HTMLImageElement) SetMsPlayToPreferredSourceURI(msPlayToPreferredSourceUri string) {
	js.Rewrite("$<.msPlayToPreferredSourceUri = $1", msPlayToPreferredSourceUri)
}

func (*HTMLImageElement) GetMsPlayToPrimary() (msPlayToPrimary bool) {
	js.Rewrite("$<.msPlayToPrimary")
	return msPlayToPrimary
}

func (*HTMLImageElement) SetMsPlayToPrimary(msPlayToPrimary bool) {
	js.Rewrite("$<.msPlayToPrimary = $1", msPlayToPrimary)
}

func (*HTMLImageElement) GetMsPlayToSource() (msPlayToSource interface{}) {
	js.Rewrite("$<.msPlayToSource")
	return msPlayToSource
}

func (*HTMLImageElement) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*HTMLImageElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

func (*HTMLImageElement) GetNaturalHeight() (naturalHeight uint) {
	js.Rewrite("$<.naturalHeight")
	return naturalHeight
}

func (*HTMLImageElement) GetNaturalWidth() (naturalWidth uint) {
	js.Rewrite("$<.naturalWidth")
	return naturalWidth
}

func (*HTMLImageElement) GetSizes() (sizes string) {
	js.Rewrite("$<.sizes")
	return sizes
}

func (*HTMLImageElement) SetSizes(sizes string) {
	js.Rewrite("$<.sizes = $1", sizes)
}

func (*HTMLImageElement) GetSrc() (src string) {
	js.Rewrite("$<.src")
	return src
}

func (*HTMLImageElement) SetSrc(src string) {
	js.Rewrite("$<.src = $1", src)
}

func (*HTMLImageElement) GetSrcset() (srcset string) {
	js.Rewrite("$<.srcset")
	return srcset
}

func (*HTMLImageElement) SetSrcset(srcset string) {
	js.Rewrite("$<.srcset = $1", srcset)
}

func (*HTMLImageElement) GetUseMap() (useMap string) {
	js.Rewrite("$<.useMap")
	return useMap
}

func (*HTMLImageElement) SetUseMap(useMap string) {
	js.Rewrite("$<.useMap = $1", useMap)
}

func (*HTMLImageElement) GetVspace() (vspace int) {
	js.Rewrite("$<.vspace")
	return vspace
}

func (*HTMLImageElement) SetVspace(vspace int) {
	js.Rewrite("$<.vspace = $1", vspace)
}

func (*HTMLImageElement) GetWidth() (width uint) {
	js.Rewrite("$<.width")
	return width
}

func (*HTMLImageElement) SetWidth(width uint) {
	js.Rewrite("$<.width = $1", width)
}

func (*HTMLImageElement) GetX() (x int) {
	js.Rewrite("$<.x")
	return x
}

func (*HTMLImageElement) GetY() (y int) {
	js.Rewrite("$<.y")
	return y
}
