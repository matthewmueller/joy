package svgpathelement

import (
	"github.com/matthewmueller/golly/dom/childnode"
	"github.com/matthewmueller/golly/dom/clientrect"
	"github.com/matthewmueller/golly/dom/clientrectlist"
	"github.com/matthewmueller/golly/dom/domstringmap"
	"github.com/matthewmueller/golly/dom/domtokenlist"
	"github.com/matthewmueller/golly/dom/mszoomtooptions"
	"github.com/matthewmueller/golly/dom/svganimatedtransformlist"
	"github.com/matthewmueller/golly/dom/svgmatrix"
	"github.com/matthewmueller/golly/dom/svgpathsegarcabs"
	"github.com/matthewmueller/golly/dom/svgpathsegarcrel"
	"github.com/matthewmueller/golly/dom/svgpathsegclosepath"
	"github.com/matthewmueller/golly/dom/svgpathsegcurvetocubicabs"
	"github.com/matthewmueller/golly/dom/svgpathsegcurvetocubicrel"
	"github.com/matthewmueller/golly/dom/svgpathsegcurvetocubicsmoothabs"
	"github.com/matthewmueller/golly/dom/svgpathsegcurvetocubicsmoothrel"
	"github.com/matthewmueller/golly/dom/svgpathsegcurvetoquadraticabs"
	"github.com/matthewmueller/golly/dom/svgpathsegcurvetoquadraticrel"
	"github.com/matthewmueller/golly/dom/svgpathsegcurvetoquadraticsmoothabs"
	"github.com/matthewmueller/golly/dom/svgpathsegcurvetoquadraticsmoothrel"
	"github.com/matthewmueller/golly/dom/svgpathseglinetoabs"
	"github.com/matthewmueller/golly/dom/svgpathseglinetohorizontalabs"
	"github.com/matthewmueller/golly/dom/svgpathseglinetohorizontalrel"
	"github.com/matthewmueller/golly/dom/svgpathseglinetorel"
	"github.com/matthewmueller/golly/dom/svgpathseglinetoverticalabs"
	"github.com/matthewmueller/golly/dom/svgpathseglinetoverticalrel"
	"github.com/matthewmueller/golly/dom/svgpathseglist"
	"github.com/matthewmueller/golly/dom/svgpathsegmovetoabs"
	"github.com/matthewmueller/golly/dom/svgpathsegmovetorel"
	"github.com/matthewmueller/golly/dom/svgpoint"
	"github.com/matthewmueller/golly/dom/svgrect"
	"github.com/matthewmueller/golly/dom/svgstringlist"
	"github.com/matthewmueller/golly/dom/svgtests"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

var _ window.SVGGraphicsElement = (*SVGPathElement)(nil)
var _ svgtests.SVGTests = (*SVGPathElement)(nil)
var _ window.SVGElement = (*SVGPathElement)(nil)
var _ window.Element = (*SVGPathElement)(nil)
var _ window.GlobalEventHandlers = (*SVGPathElement)(nil)
var _ window.ElementTraversal = (*SVGPathElement)(nil)
var _ window.NodeSelector = (*SVGPathElement)(nil)
var _ childnode.ChildNode = (*SVGPathElement)(nil)
var _ window.Node = (*SVGPathElement)(nil)
var _ window.EventTarget = (*SVGPathElement)(nil)

// SVGPathElement struct
// js:"SVGPathElement,omit"
type SVGPathElement struct {
}

// CreateSVGPathSegArcAbs fn
// js:"createSVGPathSegArcAbs"
func (*SVGPathElement) CreateSVGPathSegArcAbs(x float32, y float32, r1 float32, r2 float32, angle float32, largeArcFlag bool, sweepFlag bool) (s *svgpathsegarcabs.SVGPathSegArcAbs) {
	js.Rewrite("$_.createSVGPathSegArcAbs($1, $2, $3, $4, $5, $6, $7)", x, y, r1, r2, angle, largeArcFlag, sweepFlag)
	return s
}

// CreateSVGPathSegArcRel fn
// js:"createSVGPathSegArcRel"
func (*SVGPathElement) CreateSVGPathSegArcRel(x float32, y float32, r1 float32, r2 float32, angle float32, largeArcFlag bool, sweepFlag bool) (s *svgpathsegarcrel.SVGPathSegArcRel) {
	js.Rewrite("$_.createSVGPathSegArcRel($1, $2, $3, $4, $5, $6, $7)", x, y, r1, r2, angle, largeArcFlag, sweepFlag)
	return s
}

// CreateSVGPathSegClosePath fn
// js:"createSVGPathSegClosePath"
func (*SVGPathElement) CreateSVGPathSegClosePath() (s *svgpathsegclosepath.SVGPathSegClosePath) {
	js.Rewrite("$_.createSVGPathSegClosePath()")
	return s
}

// CreateSVGPathSegCurvetoCubicAbs fn
// js:"createSVGPathSegCurvetoCubicAbs"
func (*SVGPathElement) CreateSVGPathSegCurvetoCubicAbs(x float32, y float32, x1 float32, y1 float32, x2 float32, y2 float32) (s *svgpathsegcurvetocubicabs.SVGPathSegCurvetoCubicAbs) {
	js.Rewrite("$_.createSVGPathSegCurvetoCubicAbs($1, $2, $3, $4, $5, $6)", x, y, x1, y1, x2, y2)
	return s
}

// CreateSVGPathSegCurvetoCubicRel fn
// js:"createSVGPathSegCurvetoCubicRel"
func (*SVGPathElement) CreateSVGPathSegCurvetoCubicRel(x float32, y float32, x1 float32, y1 float32, x2 float32, y2 float32) (s *svgpathsegcurvetocubicrel.SVGPathSegCurvetoCubicRel) {
	js.Rewrite("$_.createSVGPathSegCurvetoCubicRel($1, $2, $3, $4, $5, $6)", x, y, x1, y1, x2, y2)
	return s
}

// CreateSVGPathSegCurvetoCubicSmoothAbs fn
// js:"createSVGPathSegCurvetoCubicSmoothAbs"
func (*SVGPathElement) CreateSVGPathSegCurvetoCubicSmoothAbs(x float32, y float32, x2 float32, y2 float32) (s *svgpathsegcurvetocubicsmoothabs.SVGPathSegCurvetoCubicSmoothAbs) {
	js.Rewrite("$_.createSVGPathSegCurvetoCubicSmoothAbs($1, $2, $3, $4)", x, y, x2, y2)
	return s
}

// CreateSVGPathSegCurvetoCubicSmoothRel fn
// js:"createSVGPathSegCurvetoCubicSmoothRel"
func (*SVGPathElement) CreateSVGPathSegCurvetoCubicSmoothRel(x float32, y float32, x2 float32, y2 float32) (s *svgpathsegcurvetocubicsmoothrel.SVGPathSegCurvetoCubicSmoothRel) {
	js.Rewrite("$_.createSVGPathSegCurvetoCubicSmoothRel($1, $2, $3, $4)", x, y, x2, y2)
	return s
}

// CreateSVGPathSegCurvetoQuadraticAbs fn
// js:"createSVGPathSegCurvetoQuadraticAbs"
func (*SVGPathElement) CreateSVGPathSegCurvetoQuadraticAbs(x float32, y float32, x1 float32, y1 float32) (s *svgpathsegcurvetoquadraticabs.SVGPathSegCurvetoQuadraticAbs) {
	js.Rewrite("$_.createSVGPathSegCurvetoQuadraticAbs($1, $2, $3, $4)", x, y, x1, y1)
	return s
}

// CreateSVGPathSegCurvetoQuadraticRel fn
// js:"createSVGPathSegCurvetoQuadraticRel"
func (*SVGPathElement) CreateSVGPathSegCurvetoQuadraticRel(x float32, y float32, x1 float32, y1 float32) (s *svgpathsegcurvetoquadraticrel.SVGPathSegCurvetoQuadraticRel) {
	js.Rewrite("$_.createSVGPathSegCurvetoQuadraticRel($1, $2, $3, $4)", x, y, x1, y1)
	return s
}

// CreateSVGPathSegCurvetoQuadraticSmoothAbs fn
// js:"createSVGPathSegCurvetoQuadraticSmoothAbs"
func (*SVGPathElement) CreateSVGPathSegCurvetoQuadraticSmoothAbs(x float32, y float32) (s *svgpathsegcurvetoquadraticsmoothabs.SVGPathSegCurvetoQuadraticSmoothAbs) {
	js.Rewrite("$_.createSVGPathSegCurvetoQuadraticSmoothAbs($1, $2)", x, y)
	return s
}

// CreateSVGPathSegCurvetoQuadraticSmoothRel fn
// js:"createSVGPathSegCurvetoQuadraticSmoothRel"
func (*SVGPathElement) CreateSVGPathSegCurvetoQuadraticSmoothRel(x float32, y float32) (s *svgpathsegcurvetoquadraticsmoothrel.SVGPathSegCurvetoQuadraticSmoothRel) {
	js.Rewrite("$_.createSVGPathSegCurvetoQuadraticSmoothRel($1, $2)", x, y)
	return s
}

// CreateSVGPathSegLinetoAbs fn
// js:"createSVGPathSegLinetoAbs"
func (*SVGPathElement) CreateSVGPathSegLinetoAbs(x float32, y float32) (s *svgpathseglinetoabs.SVGPathSegLinetoAbs) {
	js.Rewrite("$_.createSVGPathSegLinetoAbs($1, $2)", x, y)
	return s
}

// CreateSVGPathSegLinetoHorizontalAbs fn
// js:"createSVGPathSegLinetoHorizontalAbs"
func (*SVGPathElement) CreateSVGPathSegLinetoHorizontalAbs(x float32) (s *svgpathseglinetohorizontalabs.SVGPathSegLinetoHorizontalAbs) {
	js.Rewrite("$_.createSVGPathSegLinetoHorizontalAbs($1)", x)
	return s
}

// CreateSVGPathSegLinetoHorizontalRel fn
// js:"createSVGPathSegLinetoHorizontalRel"
func (*SVGPathElement) CreateSVGPathSegLinetoHorizontalRel(x float32) (s *svgpathseglinetohorizontalrel.SVGPathSegLinetoHorizontalRel) {
	js.Rewrite("$_.createSVGPathSegLinetoHorizontalRel($1)", x)
	return s
}

// CreateSVGPathSegLinetoRel fn
// js:"createSVGPathSegLinetoRel"
func (*SVGPathElement) CreateSVGPathSegLinetoRel(x float32, y float32) (s *svgpathseglinetorel.SVGPathSegLinetoRel) {
	js.Rewrite("$_.createSVGPathSegLinetoRel($1, $2)", x, y)
	return s
}

// CreateSVGPathSegLinetoVerticalAbs fn
// js:"createSVGPathSegLinetoVerticalAbs"
func (*SVGPathElement) CreateSVGPathSegLinetoVerticalAbs(y float32) (s *svgpathseglinetoverticalabs.SVGPathSegLinetoVerticalAbs) {
	js.Rewrite("$_.createSVGPathSegLinetoVerticalAbs($1)", y)
	return s
}

// CreateSVGPathSegLinetoVerticalRel fn
// js:"createSVGPathSegLinetoVerticalRel"
func (*SVGPathElement) CreateSVGPathSegLinetoVerticalRel(y float32) (s *svgpathseglinetoverticalrel.SVGPathSegLinetoVerticalRel) {
	js.Rewrite("$_.createSVGPathSegLinetoVerticalRel($1)", y)
	return s
}

// CreateSVGPathSegMovetoAbs fn
// js:"createSVGPathSegMovetoAbs"
func (*SVGPathElement) CreateSVGPathSegMovetoAbs(x float32, y float32) (s *svgpathsegmovetoabs.SVGPathSegMovetoAbs) {
	js.Rewrite("$_.createSVGPathSegMovetoAbs($1, $2)", x, y)
	return s
}

// CreateSVGPathSegMovetoRel fn
// js:"createSVGPathSegMovetoRel"
func (*SVGPathElement) CreateSVGPathSegMovetoRel(x float32, y float32) (s *svgpathsegmovetorel.SVGPathSegMovetoRel) {
	js.Rewrite("$_.createSVGPathSegMovetoRel($1, $2)", x, y)
	return s
}

// GetPathSegAtLength fn
// js:"getPathSegAtLength"
func (*SVGPathElement) GetPathSegAtLength(distance float32) (u uint) {
	js.Rewrite("$_.getPathSegAtLength($1)", distance)
	return u
}

// GetPointAtLength fn
// js:"getPointAtLength"
func (*SVGPathElement) GetPointAtLength(distance float32) (s *svgpoint.SVGPoint) {
	js.Rewrite("$_.getPointAtLength($1)", distance)
	return s
}

// GetTotalLength fn
// js:"getTotalLength"
func (*SVGPathElement) GetTotalLength() (f float32) {
	js.Rewrite("$_.getTotalLength()")
	return f
}

// GetBBox fn
// js:"getBBox"
func (*SVGPathElement) GetBBox() (s *svgrect.SVGRect) {
	js.Rewrite("$_.getBBox()")
	return s
}

// GetCTM fn
// js:"getCTM"
func (*SVGPathElement) GetCTM() (s *svgmatrix.SVGMatrix) {
	js.Rewrite("$_.getCTM()")
	return s
}

// GetScreenCTM fn
// js:"getScreenCTM"
func (*SVGPathElement) GetScreenCTM() (s *svgmatrix.SVGMatrix) {
	js.Rewrite("$_.getScreenCTM()")
	return s
}

// GetTransformToElement fn
// js:"getTransformToElement"
func (*SVGPathElement) GetTransformToElement(element window.SVGElement) (s *svgmatrix.SVGMatrix) {
	js.Rewrite("$_.getTransformToElement($1)", element)
	return s
}

// HasExtension fn
// js:"hasExtension"
func (*SVGPathElement) HasExtension(extension string) (b bool) {
	js.Rewrite("$_.hasExtension($1)", extension)
	return b
}

// GetAttribute fn
// js:"getAttribute"
func (*SVGPathElement) GetAttribute(qualifiedName string) (s string) {
	js.Rewrite("$_.getAttribute($1)", qualifiedName)
	return s
}

// GetAttributeNode fn
// js:"getAttributeNode"
func (*SVGPathElement) GetAttributeNode(name string) (w *window.Attr) {
	js.Rewrite("$_.getAttributeNode($1)", name)
	return w
}

// GetAttributeNodeNS fn
// js:"getAttributeNodeNS"
func (*SVGPathElement) GetAttributeNodeNS(namespaceURI string, localName string) (w *window.Attr) {
	js.Rewrite("$_.getAttributeNodeNS($1, $2)", namespaceURI, localName)
	return w
}

// GetAttributeNS fn
// js:"getAttributeNS"
func (*SVGPathElement) GetAttributeNS(namespaceURI string, localName string) (s string) {
	js.Rewrite("$_.getAttributeNS($1, $2)", namespaceURI, localName)
	return s
}

// GetBoundingClientRect fn
// js:"getBoundingClientRect"
func (*SVGPathElement) GetBoundingClientRect() (c *clientrect.ClientRect) {
	js.Rewrite("$_.getBoundingClientRect()")
	return c
}

// GetClientRects fn
// js:"getClientRects"
func (*SVGPathElement) GetClientRects() (c *clientrectlist.ClientRectList) {
	js.Rewrite("$_.getClientRects()")
	return c
}

// GetElementsByTagName fn
// js:"getElementsByTagName"
func (*SVGPathElement) GetElementsByTagName(name string) (w *window.NodeList) {
	js.Rewrite("$_.getElementsByTagName($1)", name)
	return w
}

// GetElementsByTagNameNS fn
// js:"getElementsByTagNameNS"
func (*SVGPathElement) GetElementsByTagNameNS(namespaceURI string, localName string) (w *window.NodeList) {
	js.Rewrite("$_.getElementsByTagNameNS($1, $2)", namespaceURI, localName)
	return w
}

// HasAttribute fn
// js:"hasAttribute"
func (*SVGPathElement) HasAttribute(name string) (b bool) {
	js.Rewrite("$_.hasAttribute($1)", name)
	return b
}

// HasAttributeNS fn
// js:"hasAttributeNS"
func (*SVGPathElement) HasAttributeNS(namespaceURI string, localName string) (b bool) {
	js.Rewrite("$_.hasAttributeNS($1, $2)", namespaceURI, localName)
	return b
}

// MsGetRegionContent fn
// js:"msGetRegionContent"
func (*SVGPathElement) MsGetRegionContent() (w *window.MSRangeCollection) {
	js.Rewrite("$_.msGetRegionContent()")
	return w
}

// MsGetUntransformedBounds fn
// js:"msGetUntransformedBounds"
func (*SVGPathElement) MsGetUntransformedBounds() (c *clientrect.ClientRect) {
	js.Rewrite("$_.msGetUntransformedBounds()")
	return c
}

// MsMatchesSelector fn
// js:"msMatchesSelector"
func (*SVGPathElement) MsMatchesSelector(selectors string) (b bool) {
	js.Rewrite("$_.msMatchesSelector($1)", selectors)
	return b
}

// MsReleasePointerCapture fn
// js:"msReleasePointerCapture"
func (*SVGPathElement) MsReleasePointerCapture(pointerId int) {
	js.Rewrite("$_.msReleasePointerCapture($1)", pointerId)
}

// MsSetPointerCapture fn
// js:"msSetPointerCapture"
func (*SVGPathElement) MsSetPointerCapture(pointerId int) {
	js.Rewrite("$_.msSetPointerCapture($1)", pointerId)
}

// MsZoomTo fn
// js:"msZoomTo"
func (*SVGPathElement) MsZoomTo(args *mszoomtooptions.MsZoomToOptions) {
	js.Rewrite("$_.msZoomTo($1)", args)
}

// ReleasePointerCapture fn
// js:"releasePointerCapture"
func (*SVGPathElement) ReleasePointerCapture(pointerId int) {
	js.Rewrite("$_.releasePointerCapture($1)", pointerId)
}

// RemoveAttribute fn
// js:"removeAttribute"
func (*SVGPathElement) RemoveAttribute(qualifiedName string) {
	js.Rewrite("$_.removeAttribute($1)", qualifiedName)
}

// RemoveAttributeNode fn
// js:"removeAttributeNode"
func (*SVGPathElement) RemoveAttributeNode(oldAttr *window.Attr) (w *window.Attr) {
	js.Rewrite("$_.removeAttributeNode($1)", oldAttr)
	return w
}

// RemoveAttributeNS fn
// js:"removeAttributeNS"
func (*SVGPathElement) RemoveAttributeNS(namespaceURI string, localName string) {
	js.Rewrite("$_.removeAttributeNS($1, $2)", namespaceURI, localName)
}

// RequestFullscreen fn
// js:"requestFullscreen"
func (*SVGPathElement) RequestFullscreen() {
	js.Rewrite("$_.requestFullscreen()")
}

// RequestPointerLock fn
// js:"requestPointerLock"
func (*SVGPathElement) RequestPointerLock() {
	js.Rewrite("$_.requestPointerLock()")
}

// SetAttribute fn
// js:"setAttribute"
func (*SVGPathElement) SetAttribute(qualifiedName string, value string) {
	js.Rewrite("$_.setAttribute($1, $2)", qualifiedName, value)
}

// SetAttributeNode fn
// js:"setAttributeNode"
func (*SVGPathElement) SetAttributeNode(newAttr *window.Attr) (w *window.Attr) {
	js.Rewrite("$_.setAttributeNode($1)", newAttr)
	return w
}

// SetAttributeNodeNS fn
// js:"setAttributeNodeNS"
func (*SVGPathElement) SetAttributeNodeNS(newAttr *window.Attr) (w *window.Attr) {
	js.Rewrite("$_.setAttributeNodeNS($1)", newAttr)
	return w
}

// SetAttributeNS fn
// js:"setAttributeNS"
func (*SVGPathElement) SetAttributeNS(namespaceURI string, qualifiedName string, value string) {
	js.Rewrite("$_.setAttributeNS($1, $2, $3)", namespaceURI, qualifiedName, value)
}

// SetPointerCapture fn
// js:"setPointerCapture"
func (*SVGPathElement) SetPointerCapture(pointerId int) {
	js.Rewrite("$_.setPointerCapture($1)", pointerId)
}

// WebkitMatchesSelector fn
// js:"webkitMatchesSelector"
func (*SVGPathElement) WebkitMatchesSelector(selectors string) (b bool) {
	js.Rewrite("$_.webkitMatchesSelector($1)", selectors)
	return b
}

// WebkitRequestFullscreen fn
// js:"webkitRequestFullscreen"
func (*SVGPathElement) WebkitRequestFullscreen() {
	js.Rewrite("$_.webkitRequestFullscreen()")
}

// WebkitRequestFullScreen fn
// js:"webkitRequestFullScreen"
func (*SVGPathElement) WebkitRequestFullScreen() {
	js.Rewrite("$_.webkitRequestFullScreen()")
}

// QuerySelector fn
// js:"querySelector"
func (*SVGPathElement) QuerySelector(selectors string) (w window.Element) {
	js.Rewrite("$_.querySelector($1)", selectors)
	return w
}

// QuerySelectorAll fn
// js:"querySelectorAll"
func (*SVGPathElement) QuerySelectorAll(selectors string) (w *window.NodeList) {
	js.Rewrite("$_.querySelectorAll($1)", selectors)
	return w
}

// AppendChild fn
// js:"appendChild"
func (*SVGPathElement) AppendChild(newChild window.Node) (w window.Node) {
	js.Rewrite("$_.appendChild($1)", newChild)
	return w
}

// CloneNode fn
// js:"cloneNode"
func (*SVGPathElement) CloneNode(deep *bool) (w window.Node) {
	js.Rewrite("$_.cloneNode($1)", deep)
	return w
}

// CompareDocumentPosition fn
// js:"compareDocumentPosition"
func (*SVGPathElement) CompareDocumentPosition(other window.Node) (u uint8) {
	js.Rewrite("$_.compareDocumentPosition($1)", other)
	return u
}

// Contains fn
// js:"contains"
func (*SVGPathElement) Contains(child window.Node) (b bool) {
	js.Rewrite("$_.contains($1)", child)
	return b
}

// HasAttributes fn
// js:"hasAttributes"
func (*SVGPathElement) HasAttributes() (b bool) {
	js.Rewrite("$_.hasAttributes()")
	return b
}

// HasChildNodes fn
// js:"hasChildNodes"
func (*SVGPathElement) HasChildNodes() (b bool) {
	js.Rewrite("$_.hasChildNodes()")
	return b
}

// InsertBefore fn
// js:"insertBefore"
func (*SVGPathElement) InsertBefore(newChild window.Node, refChild *window.Node) (w window.Node) {
	js.Rewrite("$_.insertBefore($1, $2)", newChild, refChild)
	return w
}

// IsDefaultNamespace fn
// js:"isDefaultNamespace"
func (*SVGPathElement) IsDefaultNamespace(namespaceURI string) (b bool) {
	js.Rewrite("$_.isDefaultNamespace($1)", namespaceURI)
	return b
}

// IsEqualNode fn
// js:"isEqualNode"
func (*SVGPathElement) IsEqualNode(arg window.Node) (b bool) {
	js.Rewrite("$_.isEqualNode($1)", arg)
	return b
}

// IsSameNode fn
// js:"isSameNode"
func (*SVGPathElement) IsSameNode(other window.Node) (b bool) {
	js.Rewrite("$_.isSameNode($1)", other)
	return b
}

// LookupNamespaceURI fn
// js:"lookupNamespaceURI"
func (*SVGPathElement) LookupNamespaceURI(prefix string) (s string) {
	js.Rewrite("$_.lookupNamespaceURI($1)", prefix)
	return s
}

// LookupPrefix fn
// js:"lookupPrefix"
func (*SVGPathElement) LookupPrefix(namespaceURI string) (s string) {
	js.Rewrite("$_.lookupPrefix($1)", namespaceURI)
	return s
}

// Normalize fn
// js:"normalize"
func (*SVGPathElement) Normalize() {
	js.Rewrite("$_.normalize()")
}

// RemoveChild fn
// js:"removeChild"
func (*SVGPathElement) RemoveChild(oldChild window.Node) (w window.Node) {
	js.Rewrite("$_.removeChild($1)", oldChild)
	return w
}

// ReplaceChild fn
// js:"replaceChild"
func (*SVGPathElement) ReplaceChild(newChild window.Node, oldChild window.Node) (w window.Node) {
	js.Rewrite("$_.replaceChild($1, $2)", newChild, oldChild)
	return w
}

// AddEventListener fn
// js:"addEventListener"
func (*SVGPathElement) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*SVGPathElement) DispatchEvent(evt window.Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*SVGPathElement) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// PathSegList prop
// js:"pathSegList"
func (*SVGPathElement) PathSegList() (pathSegList *svgpathseglist.SVGPathSegList) {
	js.Rewrite("$_.pathSegList")
	return pathSegList
}

// FarthestViewportElement prop
// js:"farthestViewportElement"
func (*SVGPathElement) FarthestViewportElement() (farthestViewportElement window.SVGElement) {
	js.Rewrite("$_.farthestViewportElement")
	return farthestViewportElement
}

// NearestViewportElement prop
// js:"nearestViewportElement"
func (*SVGPathElement) NearestViewportElement() (nearestViewportElement window.SVGElement) {
	js.Rewrite("$_.nearestViewportElement")
	return nearestViewportElement
}

// Transform prop
// js:"transform"
func (*SVGPathElement) Transform() (transform *svganimatedtransformlist.SVGAnimatedTransformList) {
	js.Rewrite("$_.transform")
	return transform
}

// RequiredExtensions prop
// js:"requiredExtensions"
func (*SVGPathElement) RequiredExtensions() (requiredExtensions *svgstringlist.SVGStringList) {
	js.Rewrite("$_.requiredExtensions")
	return requiredExtensions
}

// RequiredFeatures prop
// js:"requiredFeatures"
func (*SVGPathElement) RequiredFeatures() (requiredFeatures *svgstringlist.SVGStringList) {
	js.Rewrite("$_.requiredFeatures")
	return requiredFeatures
}

// SystemLanguage prop
// js:"systemLanguage"
func (*SVGPathElement) SystemLanguage() (systemLanguage *svgstringlist.SVGStringList) {
	js.Rewrite("$_.systemLanguage")
	return systemLanguage
}

// Dataset prop
// js:"dataset"
func (*SVGPathElement) Dataset() (dataset *domstringmap.DOMStringMap) {
	js.Rewrite("$_.dataset")
	return dataset
}

// OwnerSVGElement prop
// js:"ownerSVGElement"
func (*SVGPathElement) OwnerSVGElement() (ownerSVGElement *window.SVGSVGElement) {
	js.Rewrite("$_.ownerSVGElement")
	return ownerSVGElement
}

// ViewportElement prop
// js:"viewportElement"
func (*SVGPathElement) ViewportElement() (viewportElement window.SVGElement) {
	js.Rewrite("$_.viewportElement")
	return viewportElement
}

// Xmlbase prop
// js:"xmlbase"
func (*SVGPathElement) Xmlbase() (xmlbase string) {
	js.Rewrite("$_.xmlbase")
	return xmlbase
}

// SetXmlbase prop
// js:"xmlbase"
func (*SVGPathElement) SetXmlbase(xmlbase string) {
	js.Rewrite("$_.xmlbase = $1", xmlbase)
}

// ClassList prop
// js:"classList"
func (*SVGPathElement) ClassList() (classList domtokenlist.DOMTokenList) {
	js.Rewrite("$_.classList")
	return classList
}

// ClassName prop
// js:"className"
func (*SVGPathElement) ClassName() (className string) {
	js.Rewrite("$_.className")
	return className
}

// SetClassName prop
// js:"className"
func (*SVGPathElement) SetClassName(className string) {
	js.Rewrite("$_.className = $1", className)
}

// ClientHeight prop
// js:"clientHeight"
func (*SVGPathElement) ClientHeight() (clientHeight int) {
	js.Rewrite("$_.clientHeight")
	return clientHeight
}

// ClientLeft prop
// js:"clientLeft"
func (*SVGPathElement) ClientLeft() (clientLeft int) {
	js.Rewrite("$_.clientLeft")
	return clientLeft
}

// ClientTop prop
// js:"clientTop"
func (*SVGPathElement) ClientTop() (clientTop int) {
	js.Rewrite("$_.clientTop")
	return clientTop
}

// ClientWidth prop
// js:"clientWidth"
func (*SVGPathElement) ClientWidth() (clientWidth int) {
	js.Rewrite("$_.clientWidth")
	return clientWidth
}

// ID prop
// js:"id"
func (*SVGPathElement) ID() (id string) {
	js.Rewrite("$_.id")
	return id
}

// SetID prop
// js:"id"
func (*SVGPathElement) SetID(id string) {
	js.Rewrite("$_.id = $1", id)
}

// InnerHTML prop
// js:"innerHTML"
func (*SVGPathElement) InnerHTML() (innerHTML string) {
	js.Rewrite("$_.innerHTML")
	return innerHTML
}

// SetInnerHTML prop
// js:"innerHTML"
func (*SVGPathElement) SetInnerHTML(innerHTML string) {
	js.Rewrite("$_.innerHTML = $1", innerHTML)
}

// MsContentZoomFactor prop
// js:"msContentZoomFactor"
func (*SVGPathElement) MsContentZoomFactor() (msContentZoomFactor float32) {
	js.Rewrite("$_.msContentZoomFactor")
	return msContentZoomFactor
}

// SetMsContentZoomFactor prop
// js:"msContentZoomFactor"
func (*SVGPathElement) SetMsContentZoomFactor(msContentZoomFactor float32) {
	js.Rewrite("$_.msContentZoomFactor = $1", msContentZoomFactor)
}

// MsRegionOverflow prop
// js:"msRegionOverflow"
func (*SVGPathElement) MsRegionOverflow() (msRegionOverflow string) {
	js.Rewrite("$_.msRegionOverflow")
	return msRegionOverflow
}

// Onariarequest prop
// js:"onariarequest"
func (*SVGPathElement) Onariarequest() (onariarequest func(window.Event)) {
	js.Rewrite("$_.onariarequest")
	return onariarequest
}

// SetOnariarequest prop
// js:"onariarequest"
func (*SVGPathElement) SetOnariarequest(onariarequest func(window.Event)) {
	js.Rewrite("$_.onariarequest = $1", onariarequest)
}

// Oncommand prop
// js:"oncommand"
func (*SVGPathElement) Oncommand() (oncommand func(window.Event)) {
	js.Rewrite("$_.oncommand")
	return oncommand
}

// SetOncommand prop
// js:"oncommand"
func (*SVGPathElement) SetOncommand(oncommand func(window.Event)) {
	js.Rewrite("$_.oncommand = $1", oncommand)
}

// Ongotpointercapture prop
// js:"ongotpointercapture"
func (*SVGPathElement) Ongotpointercapture() (ongotpointercapture func(*window.PointerEvent)) {
	js.Rewrite("$_.ongotpointercapture")
	return ongotpointercapture
}

// SetOngotpointercapture prop
// js:"ongotpointercapture"
func (*SVGPathElement) SetOngotpointercapture(ongotpointercapture func(*window.PointerEvent)) {
	js.Rewrite("$_.ongotpointercapture = $1", ongotpointercapture)
}

// Onlostpointercapture prop
// js:"onlostpointercapture"
func (*SVGPathElement) Onlostpointercapture() (onlostpointercapture func(*window.PointerEvent)) {
	js.Rewrite("$_.onlostpointercapture")
	return onlostpointercapture
}

// SetOnlostpointercapture prop
// js:"onlostpointercapture"
func (*SVGPathElement) SetOnlostpointercapture(onlostpointercapture func(*window.PointerEvent)) {
	js.Rewrite("$_.onlostpointercapture = $1", onlostpointercapture)
}

// Onmsgesturechange prop
// js:"onmsgesturechange"
func (*SVGPathElement) Onmsgesturechange() (onmsgesturechange func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturechange")
	return onmsgesturechange
}

// SetOnmsgesturechange prop
// js:"onmsgesturechange"
func (*SVGPathElement) SetOnmsgesturechange(onmsgesturechange func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturechange = $1", onmsgesturechange)
}

// Onmsgesturedoubletap prop
// js:"onmsgesturedoubletap"
func (*SVGPathElement) Onmsgesturedoubletap() (onmsgesturedoubletap func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturedoubletap")
	return onmsgesturedoubletap
}

// SetOnmsgesturedoubletap prop
// js:"onmsgesturedoubletap"
func (*SVGPathElement) SetOnmsgesturedoubletap(onmsgesturedoubletap func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturedoubletap = $1", onmsgesturedoubletap)
}

// Onmsgestureend prop
// js:"onmsgestureend"
func (*SVGPathElement) Onmsgestureend() (onmsgestureend func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgestureend")
	return onmsgestureend
}

// SetOnmsgestureend prop
// js:"onmsgestureend"
func (*SVGPathElement) SetOnmsgestureend(onmsgestureend func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgestureend = $1", onmsgestureend)
}

// Onmsgesturehold prop
// js:"onmsgesturehold"
func (*SVGPathElement) Onmsgesturehold() (onmsgesturehold func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturehold")
	return onmsgesturehold
}

// SetOnmsgesturehold prop
// js:"onmsgesturehold"
func (*SVGPathElement) SetOnmsgesturehold(onmsgesturehold func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturehold = $1", onmsgesturehold)
}

// Onmsgesturestart prop
// js:"onmsgesturestart"
func (*SVGPathElement) Onmsgesturestart() (onmsgesturestart func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturestart")
	return onmsgesturestart
}

// SetOnmsgesturestart prop
// js:"onmsgesturestart"
func (*SVGPathElement) SetOnmsgesturestart(onmsgesturestart func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturestart = $1", onmsgesturestart)
}

// Onmsgesturetap prop
// js:"onmsgesturetap"
func (*SVGPathElement) Onmsgesturetap() (onmsgesturetap func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturetap")
	return onmsgesturetap
}

// SetOnmsgesturetap prop
// js:"onmsgesturetap"
func (*SVGPathElement) SetOnmsgesturetap(onmsgesturetap func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturetap = $1", onmsgesturetap)
}

// Onmsgotpointercapture prop
// js:"onmsgotpointercapture"
func (*SVGPathElement) Onmsgotpointercapture() (onmsgotpointercapture func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmsgotpointercapture")
	return onmsgotpointercapture
}

// SetOnmsgotpointercapture prop
// js:"onmsgotpointercapture"
func (*SVGPathElement) SetOnmsgotpointercapture(onmsgotpointercapture func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmsgotpointercapture = $1", onmsgotpointercapture)
}

// Onmsinertiastart prop
// js:"onmsinertiastart"
func (*SVGPathElement) Onmsinertiastart() (onmsinertiastart func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsinertiastart")
	return onmsinertiastart
}

// SetOnmsinertiastart prop
// js:"onmsinertiastart"
func (*SVGPathElement) SetOnmsinertiastart(onmsinertiastart func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsinertiastart = $1", onmsinertiastart)
}

// Onmslostpointercapture prop
// js:"onmslostpointercapture"
func (*SVGPathElement) Onmslostpointercapture() (onmslostpointercapture func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmslostpointercapture")
	return onmslostpointercapture
}

// SetOnmslostpointercapture prop
// js:"onmslostpointercapture"
func (*SVGPathElement) SetOnmslostpointercapture(onmslostpointercapture func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmslostpointercapture = $1", onmslostpointercapture)
}

// Onmspointercancel prop
// js:"onmspointercancel"
func (*SVGPathElement) Onmspointercancel() (onmspointercancel func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointercancel")
	return onmspointercancel
}

// SetOnmspointercancel prop
// js:"onmspointercancel"
func (*SVGPathElement) SetOnmspointercancel(onmspointercancel func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointercancel = $1", onmspointercancel)
}

// Onmspointerdown prop
// js:"onmspointerdown"
func (*SVGPathElement) Onmspointerdown() (onmspointerdown func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerdown")
	return onmspointerdown
}

// SetOnmspointerdown prop
// js:"onmspointerdown"
func (*SVGPathElement) SetOnmspointerdown(onmspointerdown func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerdown = $1", onmspointerdown)
}

// Onmspointerenter prop
// js:"onmspointerenter"
func (*SVGPathElement) Onmspointerenter() (onmspointerenter func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerenter")
	return onmspointerenter
}

// SetOnmspointerenter prop
// js:"onmspointerenter"
func (*SVGPathElement) SetOnmspointerenter(onmspointerenter func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerenter = $1", onmspointerenter)
}

// Onmspointerleave prop
// js:"onmspointerleave"
func (*SVGPathElement) Onmspointerleave() (onmspointerleave func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerleave")
	return onmspointerleave
}

// SetOnmspointerleave prop
// js:"onmspointerleave"
func (*SVGPathElement) SetOnmspointerleave(onmspointerleave func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerleave = $1", onmspointerleave)
}

// Onmspointermove prop
// js:"onmspointermove"
func (*SVGPathElement) Onmspointermove() (onmspointermove func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointermove")
	return onmspointermove
}

// SetOnmspointermove prop
// js:"onmspointermove"
func (*SVGPathElement) SetOnmspointermove(onmspointermove func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointermove = $1", onmspointermove)
}

// Onmspointerout prop
// js:"onmspointerout"
func (*SVGPathElement) Onmspointerout() (onmspointerout func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerout")
	return onmspointerout
}

// SetOnmspointerout prop
// js:"onmspointerout"
func (*SVGPathElement) SetOnmspointerout(onmspointerout func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerout = $1", onmspointerout)
}

// Onmspointerover prop
// js:"onmspointerover"
func (*SVGPathElement) Onmspointerover() (onmspointerover func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerover")
	return onmspointerover
}

// SetOnmspointerover prop
// js:"onmspointerover"
func (*SVGPathElement) SetOnmspointerover(onmspointerover func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerover = $1", onmspointerover)
}

// Onmspointerup prop
// js:"onmspointerup"
func (*SVGPathElement) Onmspointerup() (onmspointerup func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerup")
	return onmspointerup
}

// SetOnmspointerup prop
// js:"onmspointerup"
func (*SVGPathElement) SetOnmspointerup(onmspointerup func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerup = $1", onmspointerup)
}

// Ontouchcancel prop
// js:"ontouchcancel"
func (*SVGPathElement) Ontouchcancel() (ontouchcancel func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchcancel")
	return ontouchcancel
}

// SetOntouchcancel prop
// js:"ontouchcancel"
func (*SVGPathElement) SetOntouchcancel(ontouchcancel func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchcancel = $1", ontouchcancel)
}

// Ontouchend prop
// js:"ontouchend"
func (*SVGPathElement) Ontouchend() (ontouchend func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchend")
	return ontouchend
}

// SetOntouchend prop
// js:"ontouchend"
func (*SVGPathElement) SetOntouchend(ontouchend func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchend = $1", ontouchend)
}

// Ontouchmove prop
// js:"ontouchmove"
func (*SVGPathElement) Ontouchmove() (ontouchmove func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchmove")
	return ontouchmove
}

// SetOntouchmove prop
// js:"ontouchmove"
func (*SVGPathElement) SetOntouchmove(ontouchmove func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchmove = $1", ontouchmove)
}

// Ontouchstart prop
// js:"ontouchstart"
func (*SVGPathElement) Ontouchstart() (ontouchstart func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchstart")
	return ontouchstart
}

// SetOntouchstart prop
// js:"ontouchstart"
func (*SVGPathElement) SetOntouchstart(ontouchstart func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchstart = $1", ontouchstart)
}

// Onwebkitfullscreenchange prop
// js:"onwebkitfullscreenchange"
func (*SVGPathElement) Onwebkitfullscreenchange() (onwebkitfullscreenchange func(window.Event)) {
	js.Rewrite("$_.onwebkitfullscreenchange")
	return onwebkitfullscreenchange
}

// SetOnwebkitfullscreenchange prop
// js:"onwebkitfullscreenchange"
func (*SVGPathElement) SetOnwebkitfullscreenchange(onwebkitfullscreenchange func(window.Event)) {
	js.Rewrite("$_.onwebkitfullscreenchange = $1", onwebkitfullscreenchange)
}

// Onwebkitfullscreenerror prop
// js:"onwebkitfullscreenerror"
func (*SVGPathElement) Onwebkitfullscreenerror() (onwebkitfullscreenerror func(window.Event)) {
	js.Rewrite("$_.onwebkitfullscreenerror")
	return onwebkitfullscreenerror
}

// SetOnwebkitfullscreenerror prop
// js:"onwebkitfullscreenerror"
func (*SVGPathElement) SetOnwebkitfullscreenerror(onwebkitfullscreenerror func(window.Event)) {
	js.Rewrite("$_.onwebkitfullscreenerror = $1", onwebkitfullscreenerror)
}

// OuterHTML prop
// js:"outerHTML"
func (*SVGPathElement) OuterHTML() (outerHTML string) {
	js.Rewrite("$_.outerHTML")
	return outerHTML
}

// SetOuterHTML prop
// js:"outerHTML"
func (*SVGPathElement) SetOuterHTML(outerHTML string) {
	js.Rewrite("$_.outerHTML = $1", outerHTML)
}

// Prefix prop
// js:"prefix"
func (*SVGPathElement) Prefix() (prefix string) {
	js.Rewrite("$_.prefix")
	return prefix
}

// ScrollHeight prop
// js:"scrollHeight"
func (*SVGPathElement) ScrollHeight() (scrollHeight int) {
	js.Rewrite("$_.scrollHeight")
	return scrollHeight
}

// ScrollLeft prop
// js:"scrollLeft"
func (*SVGPathElement) ScrollLeft() (scrollLeft int) {
	js.Rewrite("$_.scrollLeft")
	return scrollLeft
}

// SetScrollLeft prop
// js:"scrollLeft"
func (*SVGPathElement) SetScrollLeft(scrollLeft int) {
	js.Rewrite("$_.scrollLeft = $1", scrollLeft)
}

// ScrollTop prop
// js:"scrollTop"
func (*SVGPathElement) ScrollTop() (scrollTop int) {
	js.Rewrite("$_.scrollTop")
	return scrollTop
}

// SetScrollTop prop
// js:"scrollTop"
func (*SVGPathElement) SetScrollTop(scrollTop int) {
	js.Rewrite("$_.scrollTop = $1", scrollTop)
}

// ScrollWidth prop
// js:"scrollWidth"
func (*SVGPathElement) ScrollWidth() (scrollWidth int) {
	js.Rewrite("$_.scrollWidth")
	return scrollWidth
}

// TagName prop
// js:"tagName"
func (*SVGPathElement) TagName() (tagName string) {
	js.Rewrite("$_.tagName")
	return tagName
}

// Onpointercancel prop
// js:"onpointercancel"
func (*SVGPathElement) Onpointercancel() (onpointercancel func(window.Event)) {
	js.Rewrite("$_.onpointercancel")
	return onpointercancel
}

// SetOnpointercancel prop
// js:"onpointercancel"
func (*SVGPathElement) SetOnpointercancel(onpointercancel func(window.Event)) {
	js.Rewrite("$_.onpointercancel = $1", onpointercancel)
}

// Onpointerdown prop
// js:"onpointerdown"
func (*SVGPathElement) Onpointerdown() (onpointerdown func(window.Event)) {
	js.Rewrite("$_.onpointerdown")
	return onpointerdown
}

// SetOnpointerdown prop
// js:"onpointerdown"
func (*SVGPathElement) SetOnpointerdown(onpointerdown func(window.Event)) {
	js.Rewrite("$_.onpointerdown = $1", onpointerdown)
}

// Onpointerenter prop
// js:"onpointerenter"
func (*SVGPathElement) Onpointerenter() (onpointerenter func(window.Event)) {
	js.Rewrite("$_.onpointerenter")
	return onpointerenter
}

// SetOnpointerenter prop
// js:"onpointerenter"
func (*SVGPathElement) SetOnpointerenter(onpointerenter func(window.Event)) {
	js.Rewrite("$_.onpointerenter = $1", onpointerenter)
}

// Onpointerleave prop
// js:"onpointerleave"
func (*SVGPathElement) Onpointerleave() (onpointerleave func(window.Event)) {
	js.Rewrite("$_.onpointerleave")
	return onpointerleave
}

// SetOnpointerleave prop
// js:"onpointerleave"
func (*SVGPathElement) SetOnpointerleave(onpointerleave func(window.Event)) {
	js.Rewrite("$_.onpointerleave = $1", onpointerleave)
}

// Onpointermove prop
// js:"onpointermove"
func (*SVGPathElement) Onpointermove() (onpointermove func(window.Event)) {
	js.Rewrite("$_.onpointermove")
	return onpointermove
}

// SetOnpointermove prop
// js:"onpointermove"
func (*SVGPathElement) SetOnpointermove(onpointermove func(window.Event)) {
	js.Rewrite("$_.onpointermove = $1", onpointermove)
}

// Onpointerout prop
// js:"onpointerout"
func (*SVGPathElement) Onpointerout() (onpointerout func(window.Event)) {
	js.Rewrite("$_.onpointerout")
	return onpointerout
}

// SetOnpointerout prop
// js:"onpointerout"
func (*SVGPathElement) SetOnpointerout(onpointerout func(window.Event)) {
	js.Rewrite("$_.onpointerout = $1", onpointerout)
}

// Onpointerover prop
// js:"onpointerover"
func (*SVGPathElement) Onpointerover() (onpointerover func(window.Event)) {
	js.Rewrite("$_.onpointerover")
	return onpointerover
}

// SetOnpointerover prop
// js:"onpointerover"
func (*SVGPathElement) SetOnpointerover(onpointerover func(window.Event)) {
	js.Rewrite("$_.onpointerover = $1", onpointerover)
}

// Onpointerup prop
// js:"onpointerup"
func (*SVGPathElement) Onpointerup() (onpointerup func(window.Event)) {
	js.Rewrite("$_.onpointerup")
	return onpointerup
}

// SetOnpointerup prop
// js:"onpointerup"
func (*SVGPathElement) SetOnpointerup(onpointerup func(window.Event)) {
	js.Rewrite("$_.onpointerup = $1", onpointerup)
}

// Onwheel prop
// js:"onwheel"
func (*SVGPathElement) Onwheel() (onwheel func(window.Event)) {
	js.Rewrite("$_.onwheel")
	return onwheel
}

// SetOnwheel prop
// js:"onwheel"
func (*SVGPathElement) SetOnwheel(onwheel func(window.Event)) {
	js.Rewrite("$_.onwheel = $1", onwheel)
}

// ChildElementCount prop
// js:"childElementCount"
func (*SVGPathElement) ChildElementCount() (childElementCount uint) {
	js.Rewrite("$_.childElementCount")
	return childElementCount
}

// FirstElementChild prop
// js:"firstElementChild"
func (*SVGPathElement) FirstElementChild() (firstElementChild window.Element) {
	js.Rewrite("$_.firstElementChild")
	return firstElementChild
}

// LastElementChild prop
// js:"lastElementChild"
func (*SVGPathElement) LastElementChild() (lastElementChild window.Element) {
	js.Rewrite("$_.lastElementChild")
	return lastElementChild
}

// NextElementSibling prop
// js:"nextElementSibling"
func (*SVGPathElement) NextElementSibling() (nextElementSibling window.Element) {
	js.Rewrite("$_.nextElementSibling")
	return nextElementSibling
}

// PreviousElementSibling prop
// js:"previousElementSibling"
func (*SVGPathElement) PreviousElementSibling() (previousElementSibling window.Element) {
	js.Rewrite("$_.previousElementSibling")
	return previousElementSibling
}

// Attributes prop
// js:"attributes"
func (*SVGPathElement) Attributes() (attributes *window.NamedNodeMap) {
	js.Rewrite("$_.attributes")
	return attributes
}

// BaseURI prop
// js:"baseURI"
func (*SVGPathElement) BaseURI() (baseURI string) {
	js.Rewrite("$_.baseURI")
	return baseURI
}

// ChildNodes prop
// js:"childNodes"
func (*SVGPathElement) ChildNodes() (childNodes *window.NodeList) {
	js.Rewrite("$_.childNodes")
	return childNodes
}

// FirstChild prop
// js:"firstChild"
func (*SVGPathElement) FirstChild() (firstChild window.Node) {
	js.Rewrite("$_.firstChild")
	return firstChild
}

// LastChild prop
// js:"lastChild"
func (*SVGPathElement) LastChild() (lastChild window.Node) {
	js.Rewrite("$_.lastChild")
	return lastChild
}

// LocalName prop
// js:"localName"
func (*SVGPathElement) LocalName() (localName string) {
	js.Rewrite("$_.localName")
	return localName
}

// NamespaceURI prop
// js:"namespaceURI"
func (*SVGPathElement) NamespaceURI() (namespaceURI string) {
	js.Rewrite("$_.namespaceURI")
	return namespaceURI
}

// NextSibling prop
// js:"nextSibling"
func (*SVGPathElement) NextSibling() (nextSibling window.Node) {
	js.Rewrite("$_.nextSibling")
	return nextSibling
}

// NodeName prop
// js:"nodeName"
func (*SVGPathElement) NodeName() (nodeName string) {
	js.Rewrite("$_.nodeName")
	return nodeName
}

// NodeType prop
// js:"nodeType"
func (*SVGPathElement) NodeType() (nodeType uint8) {
	js.Rewrite("$_.nodeType")
	return nodeType
}

// NodeValue prop
// js:"nodeValue"
func (*SVGPathElement) NodeValue() (nodeValue string) {
	js.Rewrite("$_.nodeValue")
	return nodeValue
}

// SetNodeValue prop
// js:"nodeValue"
func (*SVGPathElement) SetNodeValue(nodeValue string) {
	js.Rewrite("$_.nodeValue = $1", nodeValue)
}

// OwnerDocument prop
// js:"ownerDocument"
func (*SVGPathElement) OwnerDocument() (ownerDocument window.Document) {
	js.Rewrite("$_.ownerDocument")
	return ownerDocument
}

// ParentElement prop
// js:"parentElement"
func (*SVGPathElement) ParentElement() (parentElement window.HTMLElement) {
	js.Rewrite("$_.parentElement")
	return parentElement
}

// ParentNode prop
// js:"parentNode"
func (*SVGPathElement) ParentNode() (parentNode window.Node) {
	js.Rewrite("$_.parentNode")
	return parentNode
}

// PreviousSibling prop
// js:"previousSibling"
func (*SVGPathElement) PreviousSibling() (previousSibling window.Node) {
	js.Rewrite("$_.previousSibling")
	return previousSibling
}

// TextContent prop
// js:"textContent"
func (*SVGPathElement) TextContent() (textContent string) {
	js.Rewrite("$_.textContent")
	return textContent
}

// SetTextContent prop
// js:"textContent"
func (*SVGPathElement) SetTextContent(textContent string) {
	js.Rewrite("$_.textContent = $1", textContent)
}
