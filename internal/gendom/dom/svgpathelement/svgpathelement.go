package svgpathelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svggraphicselement"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgpathsegarcabs"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgpathsegarcrel"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgpathsegclosepath"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgpathsegcurvetocubicabs"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgpathsegcurvetocubicrel"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgpathsegcurvetocubicsmoothabs"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgpathsegcurvetocubicsmoothrel"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgpathsegcurvetoquadraticabs"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgpathsegcurvetoquadraticrel"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgpathsegcurvetoquadraticsmoothabs"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgpathsegcurvetoquadraticsmoothrel"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgpathseglinetoabs"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgpathseglinetohorizontalabs"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgpathseglinetohorizontalrel"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgpathseglinetorel"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgpathseglinetoverticalabs"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgpathseglinetoverticalrel"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgpathseglist"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgpathsegmovetoabs"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgpathsegmovetorel"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgpoint"
	"github.com/matthewmueller/golly/js"
)

type SVGPathElement struct {
	*svggraphicselement.SVGGraphicsElement
}

func (*SVGPathElement) CreateSVGPathSegArcAbs(x float32, y float32, r1 float32, r2 float32, angle float32, largeArcFlag bool, sweepFlag bool) (s *svgpathsegarcabs.SVGPathSegArcAbs) {
	js.Rewrite("$<.createSVGPathSegArcAbs($1, $2, $3, $4, $5, $6, $7)", x, y, r1, r2, angle, largeArcFlag, sweepFlag)
	return s
}

func (*SVGPathElement) CreateSVGPathSegArcRel(x float32, y float32, r1 float32, r2 float32, angle float32, largeArcFlag bool, sweepFlag bool) (s *svgpathsegarcrel.SVGPathSegArcRel) {
	js.Rewrite("$<.createSVGPathSegArcRel($1, $2, $3, $4, $5, $6, $7)", x, y, r1, r2, angle, largeArcFlag, sweepFlag)
	return s
}

func (*SVGPathElement) CreateSVGPathSegClosePath() (s *svgpathsegclosepath.SVGPathSegClosePath) {
	js.Rewrite("$<.createSVGPathSegClosePath()")
	return s
}

func (*SVGPathElement) CreateSVGPathSegCurvetoCubicAbs(x float32, y float32, x1 float32, y1 float32, x2 float32, y2 float32) (s *svgpathsegcurvetocubicabs.SVGPathSegCurvetoCubicAbs) {
	js.Rewrite("$<.createSVGPathSegCurvetoCubicAbs($1, $2, $3, $4, $5, $6)", x, y, x1, y1, x2, y2)
	return s
}

func (*SVGPathElement) CreateSVGPathSegCurvetoCubicRel(x float32, y float32, x1 float32, y1 float32, x2 float32, y2 float32) (s *svgpathsegcurvetocubicrel.SVGPathSegCurvetoCubicRel) {
	js.Rewrite("$<.createSVGPathSegCurvetoCubicRel($1, $2, $3, $4, $5, $6)", x, y, x1, y1, x2, y2)
	return s
}

func (*SVGPathElement) CreateSVGPathSegCurvetoCubicSmoothAbs(x float32, y float32, x2 float32, y2 float32) (s *svgpathsegcurvetocubicsmoothabs.SVGPathSegCurvetoCubicSmoothAbs) {
	js.Rewrite("$<.createSVGPathSegCurvetoCubicSmoothAbs($1, $2, $3, $4)", x, y, x2, y2)
	return s
}

func (*SVGPathElement) CreateSVGPathSegCurvetoCubicSmoothRel(x float32, y float32, x2 float32, y2 float32) (s *svgpathsegcurvetocubicsmoothrel.SVGPathSegCurvetoCubicSmoothRel) {
	js.Rewrite("$<.createSVGPathSegCurvetoCubicSmoothRel($1, $2, $3, $4)", x, y, x2, y2)
	return s
}

func (*SVGPathElement) CreateSVGPathSegCurvetoQuadraticAbs(x float32, y float32, x1 float32, y1 float32) (s *svgpathsegcurvetoquadraticabs.SVGPathSegCurvetoQuadraticAbs) {
	js.Rewrite("$<.createSVGPathSegCurvetoQuadraticAbs($1, $2, $3, $4)", x, y, x1, y1)
	return s
}

func (*SVGPathElement) CreateSVGPathSegCurvetoQuadraticRel(x float32, y float32, x1 float32, y1 float32) (s *svgpathsegcurvetoquadraticrel.SVGPathSegCurvetoQuadraticRel) {
	js.Rewrite("$<.createSVGPathSegCurvetoQuadraticRel($1, $2, $3, $4)", x, y, x1, y1)
	return s
}

func (*SVGPathElement) CreateSVGPathSegCurvetoQuadraticSmoothAbs(x float32, y float32) (s *svgpathsegcurvetoquadraticsmoothabs.SVGPathSegCurvetoQuadraticSmoothAbs) {
	js.Rewrite("$<.createSVGPathSegCurvetoQuadraticSmoothAbs($1, $2)", x, y)
	return s
}

func (*SVGPathElement) CreateSVGPathSegCurvetoQuadraticSmoothRel(x float32, y float32) (s *svgpathsegcurvetoquadraticsmoothrel.SVGPathSegCurvetoQuadraticSmoothRel) {
	js.Rewrite("$<.createSVGPathSegCurvetoQuadraticSmoothRel($1, $2)", x, y)
	return s
}

func (*SVGPathElement) CreateSVGPathSegLinetoAbs(x float32, y float32) (s *svgpathseglinetoabs.SVGPathSegLinetoAbs) {
	js.Rewrite("$<.createSVGPathSegLinetoAbs($1, $2)", x, y)
	return s
}

func (*SVGPathElement) CreateSVGPathSegLinetoHorizontalAbs(x float32) (s *svgpathseglinetohorizontalabs.SVGPathSegLinetoHorizontalAbs) {
	js.Rewrite("$<.createSVGPathSegLinetoHorizontalAbs($1)", x)
	return s
}

func (*SVGPathElement) CreateSVGPathSegLinetoHorizontalRel(x float32) (s *svgpathseglinetohorizontalrel.SVGPathSegLinetoHorizontalRel) {
	js.Rewrite("$<.createSVGPathSegLinetoHorizontalRel($1)", x)
	return s
}

func (*SVGPathElement) CreateSVGPathSegLinetoRel(x float32, y float32) (s *svgpathseglinetorel.SVGPathSegLinetoRel) {
	js.Rewrite("$<.createSVGPathSegLinetoRel($1, $2)", x, y)
	return s
}

func (*SVGPathElement) CreateSVGPathSegLinetoVerticalAbs(y float32) (s *svgpathseglinetoverticalabs.SVGPathSegLinetoVerticalAbs) {
	js.Rewrite("$<.createSVGPathSegLinetoVerticalAbs($1)", y)
	return s
}

func (*SVGPathElement) CreateSVGPathSegLinetoVerticalRel(y float32) (s *svgpathseglinetoverticalrel.SVGPathSegLinetoVerticalRel) {
	js.Rewrite("$<.createSVGPathSegLinetoVerticalRel($1)", y)
	return s
}

func (*SVGPathElement) CreateSVGPathSegMovetoAbs(x float32, y float32) (s *svgpathsegmovetoabs.SVGPathSegMovetoAbs) {
	js.Rewrite("$<.createSVGPathSegMovetoAbs($1, $2)", x, y)
	return s
}

func (*SVGPathElement) CreateSVGPathSegMovetoRel(x float32, y float32) (s *svgpathsegmovetorel.SVGPathSegMovetoRel) {
	js.Rewrite("$<.createSVGPathSegMovetoRel($1, $2)", x, y)
	return s
}

func (*SVGPathElement) GetPathSegAtLength(distance float32) (u uint) {
	js.Rewrite("$<.getPathSegAtLength($1)", distance)
	return u
}

func (*SVGPathElement) GetPointAtLength(distance float32) (s *svgpoint.SVGPoint) {
	js.Rewrite("$<.getPointAtLength($1)", distance)
	return s
}

func (*SVGPathElement) GetTotalLength() (f float32) {
	js.Rewrite("$<.getTotalLength()")
	return f
}

func (*SVGPathElement) GetPathSegList() (pathSegList *svgpathseglist.SVGPathSegList) {
	js.Rewrite("$<.pathSegList")
	return pathSegList
}
