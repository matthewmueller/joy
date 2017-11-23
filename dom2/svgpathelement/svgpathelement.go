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

// js:"SVGPathElement,omit"
type SVGPathElement struct {
	window.SVGGraphicsElement
}

// CreateSVGPathSegArcAbs
func (*SVGPathElement) CreateSVGPathSegArcAbs(x float32, y float32, r1 float32, r2 float32, angle float32, largeArcFlag bool, sweepFlag bool) (s *svgpathsegarcabs.SVGPathSegArcAbs) {
	js.Rewrite("$<.CreateSVGPathSegArcAbs($1, $2, $3, $4, $5, $6, $7)", x, y, r1, r2, angle, largeArcFlag, sweepFlag)
	return s
}

// CreateSVGPathSegArcRel
func (*SVGPathElement) CreateSVGPathSegArcRel(x float32, y float32, r1 float32, r2 float32, angle float32, largeArcFlag bool, sweepFlag bool) (s *svgpathsegarcrel.SVGPathSegArcRel) {
	js.Rewrite("$<.CreateSVGPathSegArcRel($1, $2, $3, $4, $5, $6, $7)", x, y, r1, r2, angle, largeArcFlag, sweepFlag)
	return s
}

// CreateSVGPathSegClosePath
func (*SVGPathElement) CreateSVGPathSegClosePath() (s *svgpathsegclosepath.SVGPathSegClosePath) {
	js.Rewrite("$<.CreateSVGPathSegClosePath()")
	return s
}

// CreateSVGPathSegCurvetoCubicAbs
func (*SVGPathElement) CreateSVGPathSegCurvetoCubicAbs(x float32, y float32, x1 float32, y1 float32, x2 float32, y2 float32) (s *svgpathsegcurvetocubicabs.SVGPathSegCurvetoCubicAbs) {
	js.Rewrite("$<.CreateSVGPathSegCurvetoCubicAbs($1, $2, $3, $4, $5, $6)", x, y, x1, y1, x2, y2)
	return s
}

// CreateSVGPathSegCurvetoCubicRel
func (*SVGPathElement) CreateSVGPathSegCurvetoCubicRel(x float32, y float32, x1 float32, y1 float32, x2 float32, y2 float32) (s *svgpathsegcurvetocubicrel.SVGPathSegCurvetoCubicRel) {
	js.Rewrite("$<.CreateSVGPathSegCurvetoCubicRel($1, $2, $3, $4, $5, $6)", x, y, x1, y1, x2, y2)
	return s
}

// CreateSVGPathSegCurvetoCubicSmoothAbs
func (*SVGPathElement) CreateSVGPathSegCurvetoCubicSmoothAbs(x float32, y float32, x2 float32, y2 float32) (s *svgpathsegcurvetocubicsmoothabs.SVGPathSegCurvetoCubicSmoothAbs) {
	js.Rewrite("$<.CreateSVGPathSegCurvetoCubicSmoothAbs($1, $2, $3, $4)", x, y, x2, y2)
	return s
}

// CreateSVGPathSegCurvetoCubicSmoothRel
func (*SVGPathElement) CreateSVGPathSegCurvetoCubicSmoothRel(x float32, y float32, x2 float32, y2 float32) (s *svgpathsegcurvetocubicsmoothrel.SVGPathSegCurvetoCubicSmoothRel) {
	js.Rewrite("$<.CreateSVGPathSegCurvetoCubicSmoothRel($1, $2, $3, $4)", x, y, x2, y2)
	return s
}

// CreateSVGPathSegCurvetoQuadraticAbs
func (*SVGPathElement) CreateSVGPathSegCurvetoQuadraticAbs(x float32, y float32, x1 float32, y1 float32) (s *svgpathsegcurvetoquadraticabs.SVGPathSegCurvetoQuadraticAbs) {
	js.Rewrite("$<.CreateSVGPathSegCurvetoQuadraticAbs($1, $2, $3, $4)", x, y, x1, y1)
	return s
}

// CreateSVGPathSegCurvetoQuadraticRel
func (*SVGPathElement) CreateSVGPathSegCurvetoQuadraticRel(x float32, y float32, x1 float32, y1 float32) (s *svgpathsegcurvetoquadraticrel.SVGPathSegCurvetoQuadraticRel) {
	js.Rewrite("$<.CreateSVGPathSegCurvetoQuadraticRel($1, $2, $3, $4)", x, y, x1, y1)
	return s
}

// CreateSVGPathSegCurvetoQuadraticSmoothAbs
func (*SVGPathElement) CreateSVGPathSegCurvetoQuadraticSmoothAbs(x float32, y float32) (s *svgpathsegcurvetoquadraticsmoothabs.SVGPathSegCurvetoQuadraticSmoothAbs) {
	js.Rewrite("$<.CreateSVGPathSegCurvetoQuadraticSmoothAbs($1, $2)", x, y)
	return s
}

// CreateSVGPathSegCurvetoQuadraticSmoothRel
func (*SVGPathElement) CreateSVGPathSegCurvetoQuadraticSmoothRel(x float32, y float32) (s *svgpathsegcurvetoquadraticsmoothrel.SVGPathSegCurvetoQuadraticSmoothRel) {
	js.Rewrite("$<.CreateSVGPathSegCurvetoQuadraticSmoothRel($1, $2)", x, y)
	return s
}

// CreateSVGPathSegLinetoAbs
func (*SVGPathElement) CreateSVGPathSegLinetoAbs(x float32, y float32) (s *svgpathseglinetoabs.SVGPathSegLinetoAbs) {
	js.Rewrite("$<.CreateSVGPathSegLinetoAbs($1, $2)", x, y)
	return s
}

// CreateSVGPathSegLinetoHorizontalAbs
func (*SVGPathElement) CreateSVGPathSegLinetoHorizontalAbs(x float32) (s *svgpathseglinetohorizontalabs.SVGPathSegLinetoHorizontalAbs) {
	js.Rewrite("$<.CreateSVGPathSegLinetoHorizontalAbs($1)", x)
	return s
}

// CreateSVGPathSegLinetoHorizontalRel
func (*SVGPathElement) CreateSVGPathSegLinetoHorizontalRel(x float32) (s *svgpathseglinetohorizontalrel.SVGPathSegLinetoHorizontalRel) {
	js.Rewrite("$<.CreateSVGPathSegLinetoHorizontalRel($1)", x)
	return s
}

// CreateSVGPathSegLinetoRel
func (*SVGPathElement) CreateSVGPathSegLinetoRel(x float32, y float32) (s *svgpathseglinetorel.SVGPathSegLinetoRel) {
	js.Rewrite("$<.CreateSVGPathSegLinetoRel($1, $2)", x, y)
	return s
}

// CreateSVGPathSegLinetoVerticalAbs
func (*SVGPathElement) CreateSVGPathSegLinetoVerticalAbs(y float32) (s *svgpathseglinetoverticalabs.SVGPathSegLinetoVerticalAbs) {
	js.Rewrite("$<.CreateSVGPathSegLinetoVerticalAbs($1)", y)
	return s
}

// CreateSVGPathSegLinetoVerticalRel
func (*SVGPathElement) CreateSVGPathSegLinetoVerticalRel(y float32) (s *svgpathseglinetoverticalrel.SVGPathSegLinetoVerticalRel) {
	js.Rewrite("$<.CreateSVGPathSegLinetoVerticalRel($1)", y)
	return s
}

// CreateSVGPathSegMovetoAbs
func (*SVGPathElement) CreateSVGPathSegMovetoAbs(x float32, y float32) (s *svgpathsegmovetoabs.SVGPathSegMovetoAbs) {
	js.Rewrite("$<.CreateSVGPathSegMovetoAbs($1, $2)", x, y)
	return s
}

// CreateSVGPathSegMovetoRel
func (*SVGPathElement) CreateSVGPathSegMovetoRel(x float32, y float32) (s *svgpathsegmovetorel.SVGPathSegMovetoRel) {
	js.Rewrite("$<.CreateSVGPathSegMovetoRel($1, $2)", x, y)
	return s
}

// GetPathSegAtLength
func (*SVGPathElement) GetPathSegAtLength(distance float32) (u uint) {
	js.Rewrite("$<.GetPathSegAtLength($1)", distance)
	return u
}

// GetPointAtLength
func (*SVGPathElement) GetPointAtLength(distance float32) (s *svgpoint.SVGPoint) {
	js.Rewrite("$<.GetPointAtLength($1)", distance)
	return s
}

// GetTotalLength
func (*SVGPathElement) GetTotalLength() (f float32) {
	js.Rewrite("$<.GetTotalLength()")
	return f
}

// PathSegList
func (*SVGPathElement) PathSegList() (pathSegList *svgpathseglist.SVGPathSegList) {
	js.Rewrite("$<.PathSegList")
	return pathSegList
}
