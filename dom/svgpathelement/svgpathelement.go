package svgpathelement

import (
	"github.com/matthewmueller/golly/dom2/svgpathsegarcabs"
	"github.com/matthewmueller/golly/dom2/svgpathsegarcrel"
	"github.com/matthewmueller/golly/dom2/svgpathsegclosepath"
	"github.com/matthewmueller/golly/dom2/svgpathsegcurvetocubicabs"
	"github.com/matthewmueller/golly/dom2/svgpathsegcurvetocubicrel"
	"github.com/matthewmueller/golly/dom2/svgpathsegcurvetocubicsmoothabs"
	"github.com/matthewmueller/golly/dom2/svgpathsegcurvetocubicsmoothrel"
	"github.com/matthewmueller/golly/dom2/svgpathsegcurvetoquadraticabs"
	"github.com/matthewmueller/golly/dom2/svgpathsegcurvetoquadraticrel"
	"github.com/matthewmueller/golly/dom2/svgpathsegcurvetoquadraticsmoothabs"
	"github.com/matthewmueller/golly/dom2/svgpathsegcurvetoquadraticsmoothrel"
	"github.com/matthewmueller/golly/dom2/svgpathseglinetoabs"
	"github.com/matthewmueller/golly/dom2/svgpathseglinetohorizontalabs"
	"github.com/matthewmueller/golly/dom2/svgpathseglinetohorizontalrel"
	"github.com/matthewmueller/golly/dom2/svgpathseglinetorel"
	"github.com/matthewmueller/golly/dom2/svgpathseglinetoverticalabs"
	"github.com/matthewmueller/golly/dom2/svgpathseglinetoverticalrel"
	"github.com/matthewmueller/golly/dom2/svgpathseglist"
	"github.com/matthewmueller/golly/dom2/svgpathsegmovetoabs"
	"github.com/matthewmueller/golly/dom2/svgpathsegmovetorel"
	"github.com/matthewmueller/golly/dom2/svgpoint"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// SVGPathElement struct
// js:"SVGPathElement,omit"
type SVGPathElement struct {
	window.SVGGraphicsElement
}

// CreateSVGPathSegArcAbs fn
func (*SVGPathElement) CreateSVGPathSegArcAbs(x float32, y float32, r1 float32, r2 float32, angle float32, largeArcFlag bool, sweepFlag bool) (s *svgpathsegarcabs.SVGPathSegArcAbs) {
	js.Rewrite("$<.createSVGPathSegArcAbs($1, $2, $3, $4, $5, $6, $7)", x, y, r1, r2, angle, largeArcFlag, sweepFlag)
	return s
}

// CreateSVGPathSegArcRel fn
func (*SVGPathElement) CreateSVGPathSegArcRel(x float32, y float32, r1 float32, r2 float32, angle float32, largeArcFlag bool, sweepFlag bool) (s *svgpathsegarcrel.SVGPathSegArcRel) {
	js.Rewrite("$<.createSVGPathSegArcRel($1, $2, $3, $4, $5, $6, $7)", x, y, r1, r2, angle, largeArcFlag, sweepFlag)
	return s
}

// CreateSVGPathSegClosePath fn
func (*SVGPathElement) CreateSVGPathSegClosePath() (s *svgpathsegclosepath.SVGPathSegClosePath) {
	js.Rewrite("$<.createSVGPathSegClosePath()")
	return s
}

// CreateSVGPathSegCurvetoCubicAbs fn
func (*SVGPathElement) CreateSVGPathSegCurvetoCubicAbs(x float32, y float32, x1 float32, y1 float32, x2 float32, y2 float32) (s *svgpathsegcurvetocubicabs.SVGPathSegCurvetoCubicAbs) {
	js.Rewrite("$<.createSVGPathSegCurvetoCubicAbs($1, $2, $3, $4, $5, $6)", x, y, x1, y1, x2, y2)
	return s
}

// CreateSVGPathSegCurvetoCubicRel fn
func (*SVGPathElement) CreateSVGPathSegCurvetoCubicRel(x float32, y float32, x1 float32, y1 float32, x2 float32, y2 float32) (s *svgpathsegcurvetocubicrel.SVGPathSegCurvetoCubicRel) {
	js.Rewrite("$<.createSVGPathSegCurvetoCubicRel($1, $2, $3, $4, $5, $6)", x, y, x1, y1, x2, y2)
	return s
}

// CreateSVGPathSegCurvetoCubicSmoothAbs fn
func (*SVGPathElement) CreateSVGPathSegCurvetoCubicSmoothAbs(x float32, y float32, x2 float32, y2 float32) (s *svgpathsegcurvetocubicsmoothabs.SVGPathSegCurvetoCubicSmoothAbs) {
	js.Rewrite("$<.createSVGPathSegCurvetoCubicSmoothAbs($1, $2, $3, $4)", x, y, x2, y2)
	return s
}

// CreateSVGPathSegCurvetoCubicSmoothRel fn
func (*SVGPathElement) CreateSVGPathSegCurvetoCubicSmoothRel(x float32, y float32, x2 float32, y2 float32) (s *svgpathsegcurvetocubicsmoothrel.SVGPathSegCurvetoCubicSmoothRel) {
	js.Rewrite("$<.createSVGPathSegCurvetoCubicSmoothRel($1, $2, $3, $4)", x, y, x2, y2)
	return s
}

// CreateSVGPathSegCurvetoQuadraticAbs fn
func (*SVGPathElement) CreateSVGPathSegCurvetoQuadraticAbs(x float32, y float32, x1 float32, y1 float32) (s *svgpathsegcurvetoquadraticabs.SVGPathSegCurvetoQuadraticAbs) {
	js.Rewrite("$<.createSVGPathSegCurvetoQuadraticAbs($1, $2, $3, $4)", x, y, x1, y1)
	return s
}

// CreateSVGPathSegCurvetoQuadraticRel fn
func (*SVGPathElement) CreateSVGPathSegCurvetoQuadraticRel(x float32, y float32, x1 float32, y1 float32) (s *svgpathsegcurvetoquadraticrel.SVGPathSegCurvetoQuadraticRel) {
	js.Rewrite("$<.createSVGPathSegCurvetoQuadraticRel($1, $2, $3, $4)", x, y, x1, y1)
	return s
}

// CreateSVGPathSegCurvetoQuadraticSmoothAbs fn
func (*SVGPathElement) CreateSVGPathSegCurvetoQuadraticSmoothAbs(x float32, y float32) (s *svgpathsegcurvetoquadraticsmoothabs.SVGPathSegCurvetoQuadraticSmoothAbs) {
	js.Rewrite("$<.createSVGPathSegCurvetoQuadraticSmoothAbs($1, $2)", x, y)
	return s
}

// CreateSVGPathSegCurvetoQuadraticSmoothRel fn
func (*SVGPathElement) CreateSVGPathSegCurvetoQuadraticSmoothRel(x float32, y float32) (s *svgpathsegcurvetoquadraticsmoothrel.SVGPathSegCurvetoQuadraticSmoothRel) {
	js.Rewrite("$<.createSVGPathSegCurvetoQuadraticSmoothRel($1, $2)", x, y)
	return s
}

// CreateSVGPathSegLinetoAbs fn
func (*SVGPathElement) CreateSVGPathSegLinetoAbs(x float32, y float32) (s *svgpathseglinetoabs.SVGPathSegLinetoAbs) {
	js.Rewrite("$<.createSVGPathSegLinetoAbs($1, $2)", x, y)
	return s
}

// CreateSVGPathSegLinetoHorizontalAbs fn
func (*SVGPathElement) CreateSVGPathSegLinetoHorizontalAbs(x float32) (s *svgpathseglinetohorizontalabs.SVGPathSegLinetoHorizontalAbs) {
	js.Rewrite("$<.createSVGPathSegLinetoHorizontalAbs($1)", x)
	return s
}

// CreateSVGPathSegLinetoHorizontalRel fn
func (*SVGPathElement) CreateSVGPathSegLinetoHorizontalRel(x float32) (s *svgpathseglinetohorizontalrel.SVGPathSegLinetoHorizontalRel) {
	js.Rewrite("$<.createSVGPathSegLinetoHorizontalRel($1)", x)
	return s
}

// CreateSVGPathSegLinetoRel fn
func (*SVGPathElement) CreateSVGPathSegLinetoRel(x float32, y float32) (s *svgpathseglinetorel.SVGPathSegLinetoRel) {
	js.Rewrite("$<.createSVGPathSegLinetoRel($1, $2)", x, y)
	return s
}

// CreateSVGPathSegLinetoVerticalAbs fn
func (*SVGPathElement) CreateSVGPathSegLinetoVerticalAbs(y float32) (s *svgpathseglinetoverticalabs.SVGPathSegLinetoVerticalAbs) {
	js.Rewrite("$<.createSVGPathSegLinetoVerticalAbs($1)", y)
	return s
}

// CreateSVGPathSegLinetoVerticalRel fn
func (*SVGPathElement) CreateSVGPathSegLinetoVerticalRel(y float32) (s *svgpathseglinetoverticalrel.SVGPathSegLinetoVerticalRel) {
	js.Rewrite("$<.createSVGPathSegLinetoVerticalRel($1)", y)
	return s
}

// CreateSVGPathSegMovetoAbs fn
func (*SVGPathElement) CreateSVGPathSegMovetoAbs(x float32, y float32) (s *svgpathsegmovetoabs.SVGPathSegMovetoAbs) {
	js.Rewrite("$<.createSVGPathSegMovetoAbs($1, $2)", x, y)
	return s
}

// CreateSVGPathSegMovetoRel fn
func (*SVGPathElement) CreateSVGPathSegMovetoRel(x float32, y float32) (s *svgpathsegmovetorel.SVGPathSegMovetoRel) {
	js.Rewrite("$<.createSVGPathSegMovetoRel($1, $2)", x, y)
	return s
}

// GetPathSegAtLength fn
func (*SVGPathElement) GetPathSegAtLength(distance float32) (u uint) {
	js.Rewrite("$<.getPathSegAtLength($1)", distance)
	return u
}

// GetPointAtLength fn
func (*SVGPathElement) GetPointAtLength(distance float32) (s *svgpoint.SVGPoint) {
	js.Rewrite("$<.getPointAtLength($1)", distance)
	return s
}

// GetTotalLength fn
func (*SVGPathElement) GetTotalLength() (f float32) {
	js.Rewrite("$<.getTotalLength()")
	return f
}

// PathSegList prop
func (*SVGPathElement) PathSegList() (pathSegList *svgpathseglist.SVGPathSegList) {
	js.Rewrite("$<.pathSegList")
	return pathSegList
}
