package cssstyledeclaration

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/cssrule"
	"github.com/matthewmueller/golly/js"
)

type CSSStyleDeclaration struct {
}

func (*CSSStyleDeclaration) GetPropertyPriority(propertyName string) (s string) {
	js.Rewrite("$<.getPropertyPriority($1)", propertyName)
	return s
}

func (*CSSStyleDeclaration) GetPropertyValue(propertyName string) (s string) {
	js.Rewrite("$<.getPropertyValue($1)", propertyName)
	return s
}

func (*CSSStyleDeclaration) Item(index uint) (s string) {
	js.Rewrite("$<.item($1)", index)
	return s
}

func (*CSSStyleDeclaration) RemoveProperty(propertyName string) (s string) {
	js.Rewrite("$<.removeProperty($1)", propertyName)
	return s
}

func (*CSSStyleDeclaration) SetProperty(propertyName string, value string, priority string) {
	js.Rewrite("$<.setProperty($1, $2, $3)", propertyName, value, priority)
}

func (*CSSStyleDeclaration) GetAlignContent() (alignContent string) {
	js.Rewrite("$<.alignContent")
	return alignContent
}

func (*CSSStyleDeclaration) SetAlignContent(alignContent string) {
	js.Rewrite("$<.alignContent = $1", alignContent)
}

func (*CSSStyleDeclaration) GetAlignItems() (alignItems string) {
	js.Rewrite("$<.alignItems")
	return alignItems
}

func (*CSSStyleDeclaration) SetAlignItems(alignItems string) {
	js.Rewrite("$<.alignItems = $1", alignItems)
}

func (*CSSStyleDeclaration) GetAlignmentBaseline() (alignmentBaseline string) {
	js.Rewrite("$<.alignmentBaseline")
	return alignmentBaseline
}

func (*CSSStyleDeclaration) SetAlignmentBaseline(alignmentBaseline string) {
	js.Rewrite("$<.alignmentBaseline = $1", alignmentBaseline)
}

func (*CSSStyleDeclaration) GetAlignSelf() (alignSelf string) {
	js.Rewrite("$<.alignSelf")
	return alignSelf
}

func (*CSSStyleDeclaration) SetAlignSelf(alignSelf string) {
	js.Rewrite("$<.alignSelf = $1", alignSelf)
}

func (*CSSStyleDeclaration) GetAnimation() (animation string) {
	js.Rewrite("$<.animation")
	return animation
}

func (*CSSStyleDeclaration) SetAnimation(animation string) {
	js.Rewrite("$<.animation = $1", animation)
}

func (*CSSStyleDeclaration) GetAnimationDelay() (animationDelay string) {
	js.Rewrite("$<.animationDelay")
	return animationDelay
}

func (*CSSStyleDeclaration) SetAnimationDelay(animationDelay string) {
	js.Rewrite("$<.animationDelay = $1", animationDelay)
}

func (*CSSStyleDeclaration) GetAnimationDirection() (animationDirection string) {
	js.Rewrite("$<.animationDirection")
	return animationDirection
}

func (*CSSStyleDeclaration) SetAnimationDirection(animationDirection string) {
	js.Rewrite("$<.animationDirection = $1", animationDirection)
}

func (*CSSStyleDeclaration) GetAnimationDuration() (animationDuration string) {
	js.Rewrite("$<.animationDuration")
	return animationDuration
}

func (*CSSStyleDeclaration) SetAnimationDuration(animationDuration string) {
	js.Rewrite("$<.animationDuration = $1", animationDuration)
}

func (*CSSStyleDeclaration) GetAnimationFillMode() (animationFillMode string) {
	js.Rewrite("$<.animationFillMode")
	return animationFillMode
}

func (*CSSStyleDeclaration) SetAnimationFillMode(animationFillMode string) {
	js.Rewrite("$<.animationFillMode = $1", animationFillMode)
}

func (*CSSStyleDeclaration) GetAnimationIterationCount() (animationIterationCount string) {
	js.Rewrite("$<.animationIterationCount")
	return animationIterationCount
}

func (*CSSStyleDeclaration) SetAnimationIterationCount(animationIterationCount string) {
	js.Rewrite("$<.animationIterationCount = $1", animationIterationCount)
}

func (*CSSStyleDeclaration) GetAnimationName() (animationName string) {
	js.Rewrite("$<.animationName")
	return animationName
}

func (*CSSStyleDeclaration) SetAnimationName(animationName string) {
	js.Rewrite("$<.animationName = $1", animationName)
}

func (*CSSStyleDeclaration) GetAnimationPlayState() (animationPlayState string) {
	js.Rewrite("$<.animationPlayState")
	return animationPlayState
}

func (*CSSStyleDeclaration) SetAnimationPlayState(animationPlayState string) {
	js.Rewrite("$<.animationPlayState = $1", animationPlayState)
}

func (*CSSStyleDeclaration) GetAnimationTimingFunction() (animationTimingFunction string) {
	js.Rewrite("$<.animationTimingFunction")
	return animationTimingFunction
}

func (*CSSStyleDeclaration) SetAnimationTimingFunction(animationTimingFunction string) {
	js.Rewrite("$<.animationTimingFunction = $1", animationTimingFunction)
}

func (*CSSStyleDeclaration) GetBackfaceVisibility() (backfaceVisibility string) {
	js.Rewrite("$<.backfaceVisibility")
	return backfaceVisibility
}

func (*CSSStyleDeclaration) SetBackfaceVisibility(backfaceVisibility string) {
	js.Rewrite("$<.backfaceVisibility = $1", backfaceVisibility)
}

func (*CSSStyleDeclaration) GetBackground() (background string) {
	js.Rewrite("$<.background")
	return background
}

func (*CSSStyleDeclaration) SetBackground(background string) {
	js.Rewrite("$<.background = $1", background)
}

func (*CSSStyleDeclaration) GetBackgroundAttachment() (backgroundAttachment string) {
	js.Rewrite("$<.backgroundAttachment")
	return backgroundAttachment
}

func (*CSSStyleDeclaration) SetBackgroundAttachment(backgroundAttachment string) {
	js.Rewrite("$<.backgroundAttachment = $1", backgroundAttachment)
}

func (*CSSStyleDeclaration) GetBackgroundClip() (backgroundClip string) {
	js.Rewrite("$<.backgroundClip")
	return backgroundClip
}

func (*CSSStyleDeclaration) SetBackgroundClip(backgroundClip string) {
	js.Rewrite("$<.backgroundClip = $1", backgroundClip)
}

func (*CSSStyleDeclaration) GetBackgroundColor() (backgroundColor string) {
	js.Rewrite("$<.backgroundColor")
	return backgroundColor
}

func (*CSSStyleDeclaration) SetBackgroundColor(backgroundColor string) {
	js.Rewrite("$<.backgroundColor = $1", backgroundColor)
}

func (*CSSStyleDeclaration) GetBackgroundImage() (backgroundImage string) {
	js.Rewrite("$<.backgroundImage")
	return backgroundImage
}

func (*CSSStyleDeclaration) SetBackgroundImage(backgroundImage string) {
	js.Rewrite("$<.backgroundImage = $1", backgroundImage)
}

func (*CSSStyleDeclaration) GetBackgroundOrigin() (backgroundOrigin string) {
	js.Rewrite("$<.backgroundOrigin")
	return backgroundOrigin
}

func (*CSSStyleDeclaration) SetBackgroundOrigin(backgroundOrigin string) {
	js.Rewrite("$<.backgroundOrigin = $1", backgroundOrigin)
}

func (*CSSStyleDeclaration) GetBackgroundPosition() (backgroundPosition string) {
	js.Rewrite("$<.backgroundPosition")
	return backgroundPosition
}

func (*CSSStyleDeclaration) SetBackgroundPosition(backgroundPosition string) {
	js.Rewrite("$<.backgroundPosition = $1", backgroundPosition)
}

func (*CSSStyleDeclaration) GetBackgroundPositionX() (backgroundPositionX string) {
	js.Rewrite("$<.backgroundPositionX")
	return backgroundPositionX
}

func (*CSSStyleDeclaration) SetBackgroundPositionX(backgroundPositionX string) {
	js.Rewrite("$<.backgroundPositionX = $1", backgroundPositionX)
}

func (*CSSStyleDeclaration) GetBackgroundPositionY() (backgroundPositionY string) {
	js.Rewrite("$<.backgroundPositionY")
	return backgroundPositionY
}

func (*CSSStyleDeclaration) SetBackgroundPositionY(backgroundPositionY string) {
	js.Rewrite("$<.backgroundPositionY = $1", backgroundPositionY)
}

func (*CSSStyleDeclaration) GetBackgroundRepeat() (backgroundRepeat string) {
	js.Rewrite("$<.backgroundRepeat")
	return backgroundRepeat
}

func (*CSSStyleDeclaration) SetBackgroundRepeat(backgroundRepeat string) {
	js.Rewrite("$<.backgroundRepeat = $1", backgroundRepeat)
}

func (*CSSStyleDeclaration) GetBackgroundSize() (backgroundSize string) {
	js.Rewrite("$<.backgroundSize")
	return backgroundSize
}

func (*CSSStyleDeclaration) SetBackgroundSize(backgroundSize string) {
	js.Rewrite("$<.backgroundSize = $1", backgroundSize)
}

func (*CSSStyleDeclaration) GetBaselineShift() (baselineShift string) {
	js.Rewrite("$<.baselineShift")
	return baselineShift
}

func (*CSSStyleDeclaration) SetBaselineShift(baselineShift string) {
	js.Rewrite("$<.baselineShift = $1", baselineShift)
}

func (*CSSStyleDeclaration) GetBorder() (border string) {
	js.Rewrite("$<.border")
	return border
}

func (*CSSStyleDeclaration) SetBorder(border string) {
	js.Rewrite("$<.border = $1", border)
}

func (*CSSStyleDeclaration) GetBorderBottom() (borderBottom string) {
	js.Rewrite("$<.borderBottom")
	return borderBottom
}

func (*CSSStyleDeclaration) SetBorderBottom(borderBottom string) {
	js.Rewrite("$<.borderBottom = $1", borderBottom)
}

func (*CSSStyleDeclaration) GetBorderBottomColor() (borderBottomColor string) {
	js.Rewrite("$<.borderBottomColor")
	return borderBottomColor
}

func (*CSSStyleDeclaration) SetBorderBottomColor(borderBottomColor string) {
	js.Rewrite("$<.borderBottomColor = $1", borderBottomColor)
}

func (*CSSStyleDeclaration) GetBorderBottomLeftRadius() (borderBottomLeftRadius string) {
	js.Rewrite("$<.borderBottomLeftRadius")
	return borderBottomLeftRadius
}

func (*CSSStyleDeclaration) SetBorderBottomLeftRadius(borderBottomLeftRadius string) {
	js.Rewrite("$<.borderBottomLeftRadius = $1", borderBottomLeftRadius)
}

func (*CSSStyleDeclaration) GetBorderBottomRightRadius() (borderBottomRightRadius string) {
	js.Rewrite("$<.borderBottomRightRadius")
	return borderBottomRightRadius
}

func (*CSSStyleDeclaration) SetBorderBottomRightRadius(borderBottomRightRadius string) {
	js.Rewrite("$<.borderBottomRightRadius = $1", borderBottomRightRadius)
}

func (*CSSStyleDeclaration) GetBorderBottomStyle() (borderBottomStyle string) {
	js.Rewrite("$<.borderBottomStyle")
	return borderBottomStyle
}

func (*CSSStyleDeclaration) SetBorderBottomStyle(borderBottomStyle string) {
	js.Rewrite("$<.borderBottomStyle = $1", borderBottomStyle)
}

func (*CSSStyleDeclaration) GetBorderBottomWidth() (borderBottomWidth string) {
	js.Rewrite("$<.borderBottomWidth")
	return borderBottomWidth
}

func (*CSSStyleDeclaration) SetBorderBottomWidth(borderBottomWidth string) {
	js.Rewrite("$<.borderBottomWidth = $1", borderBottomWidth)
}

func (*CSSStyleDeclaration) GetBorderCollapse() (borderCollapse string) {
	js.Rewrite("$<.borderCollapse")
	return borderCollapse
}

func (*CSSStyleDeclaration) SetBorderCollapse(borderCollapse string) {
	js.Rewrite("$<.borderCollapse = $1", borderCollapse)
}

func (*CSSStyleDeclaration) GetBorderColor() (borderColor string) {
	js.Rewrite("$<.borderColor")
	return borderColor
}

func (*CSSStyleDeclaration) SetBorderColor(borderColor string) {
	js.Rewrite("$<.borderColor = $1", borderColor)
}

func (*CSSStyleDeclaration) GetBorderImage() (borderImage string) {
	js.Rewrite("$<.borderImage")
	return borderImage
}

func (*CSSStyleDeclaration) SetBorderImage(borderImage string) {
	js.Rewrite("$<.borderImage = $1", borderImage)
}

func (*CSSStyleDeclaration) GetBorderImageOutset() (borderImageOutset string) {
	js.Rewrite("$<.borderImageOutset")
	return borderImageOutset
}

func (*CSSStyleDeclaration) SetBorderImageOutset(borderImageOutset string) {
	js.Rewrite("$<.borderImageOutset = $1", borderImageOutset)
}

func (*CSSStyleDeclaration) GetBorderImageRepeat() (borderImageRepeat string) {
	js.Rewrite("$<.borderImageRepeat")
	return borderImageRepeat
}

func (*CSSStyleDeclaration) SetBorderImageRepeat(borderImageRepeat string) {
	js.Rewrite("$<.borderImageRepeat = $1", borderImageRepeat)
}

func (*CSSStyleDeclaration) GetBorderImageSlice() (borderImageSlice string) {
	js.Rewrite("$<.borderImageSlice")
	return borderImageSlice
}

func (*CSSStyleDeclaration) SetBorderImageSlice(borderImageSlice string) {
	js.Rewrite("$<.borderImageSlice = $1", borderImageSlice)
}

func (*CSSStyleDeclaration) GetBorderImageSource() (borderImageSource string) {
	js.Rewrite("$<.borderImageSource")
	return borderImageSource
}

func (*CSSStyleDeclaration) SetBorderImageSource(borderImageSource string) {
	js.Rewrite("$<.borderImageSource = $1", borderImageSource)
}

func (*CSSStyleDeclaration) GetBorderImageWidth() (borderImageWidth string) {
	js.Rewrite("$<.borderImageWidth")
	return borderImageWidth
}

func (*CSSStyleDeclaration) SetBorderImageWidth(borderImageWidth string) {
	js.Rewrite("$<.borderImageWidth = $1", borderImageWidth)
}

func (*CSSStyleDeclaration) GetBorderLeft() (borderLeft string) {
	js.Rewrite("$<.borderLeft")
	return borderLeft
}

func (*CSSStyleDeclaration) SetBorderLeft(borderLeft string) {
	js.Rewrite("$<.borderLeft = $1", borderLeft)
}

func (*CSSStyleDeclaration) GetBorderLeftColor() (borderLeftColor string) {
	js.Rewrite("$<.borderLeftColor")
	return borderLeftColor
}

func (*CSSStyleDeclaration) SetBorderLeftColor(borderLeftColor string) {
	js.Rewrite("$<.borderLeftColor = $1", borderLeftColor)
}

func (*CSSStyleDeclaration) GetBorderLeftStyle() (borderLeftStyle string) {
	js.Rewrite("$<.borderLeftStyle")
	return borderLeftStyle
}

func (*CSSStyleDeclaration) SetBorderLeftStyle(borderLeftStyle string) {
	js.Rewrite("$<.borderLeftStyle = $1", borderLeftStyle)
}

func (*CSSStyleDeclaration) GetBorderLeftWidth() (borderLeftWidth string) {
	js.Rewrite("$<.borderLeftWidth")
	return borderLeftWidth
}

func (*CSSStyleDeclaration) SetBorderLeftWidth(borderLeftWidth string) {
	js.Rewrite("$<.borderLeftWidth = $1", borderLeftWidth)
}

func (*CSSStyleDeclaration) GetBorderRadius() (borderRadius string) {
	js.Rewrite("$<.borderRadius")
	return borderRadius
}

func (*CSSStyleDeclaration) SetBorderRadius(borderRadius string) {
	js.Rewrite("$<.borderRadius = $1", borderRadius)
}

func (*CSSStyleDeclaration) GetBorderRight() (borderRight string) {
	js.Rewrite("$<.borderRight")
	return borderRight
}

func (*CSSStyleDeclaration) SetBorderRight(borderRight string) {
	js.Rewrite("$<.borderRight = $1", borderRight)
}

func (*CSSStyleDeclaration) GetBorderRightColor() (borderRightColor string) {
	js.Rewrite("$<.borderRightColor")
	return borderRightColor
}

func (*CSSStyleDeclaration) SetBorderRightColor(borderRightColor string) {
	js.Rewrite("$<.borderRightColor = $1", borderRightColor)
}

func (*CSSStyleDeclaration) GetBorderRightStyle() (borderRightStyle string) {
	js.Rewrite("$<.borderRightStyle")
	return borderRightStyle
}

func (*CSSStyleDeclaration) SetBorderRightStyle(borderRightStyle string) {
	js.Rewrite("$<.borderRightStyle = $1", borderRightStyle)
}

func (*CSSStyleDeclaration) GetBorderRightWidth() (borderRightWidth string) {
	js.Rewrite("$<.borderRightWidth")
	return borderRightWidth
}

func (*CSSStyleDeclaration) SetBorderRightWidth(borderRightWidth string) {
	js.Rewrite("$<.borderRightWidth = $1", borderRightWidth)
}

func (*CSSStyleDeclaration) GetBorderSpacing() (borderSpacing string) {
	js.Rewrite("$<.borderSpacing")
	return borderSpacing
}

func (*CSSStyleDeclaration) SetBorderSpacing(borderSpacing string) {
	js.Rewrite("$<.borderSpacing = $1", borderSpacing)
}

func (*CSSStyleDeclaration) GetBorderStyle() (borderStyle string) {
	js.Rewrite("$<.borderStyle")
	return borderStyle
}

func (*CSSStyleDeclaration) SetBorderStyle(borderStyle string) {
	js.Rewrite("$<.borderStyle = $1", borderStyle)
}

func (*CSSStyleDeclaration) GetBorderTop() (borderTop string) {
	js.Rewrite("$<.borderTop")
	return borderTop
}

func (*CSSStyleDeclaration) SetBorderTop(borderTop string) {
	js.Rewrite("$<.borderTop = $1", borderTop)
}

func (*CSSStyleDeclaration) GetBorderTopColor() (borderTopColor string) {
	js.Rewrite("$<.borderTopColor")
	return borderTopColor
}

func (*CSSStyleDeclaration) SetBorderTopColor(borderTopColor string) {
	js.Rewrite("$<.borderTopColor = $1", borderTopColor)
}

func (*CSSStyleDeclaration) GetBorderTopLeftRadius() (borderTopLeftRadius string) {
	js.Rewrite("$<.borderTopLeftRadius")
	return borderTopLeftRadius
}

func (*CSSStyleDeclaration) SetBorderTopLeftRadius(borderTopLeftRadius string) {
	js.Rewrite("$<.borderTopLeftRadius = $1", borderTopLeftRadius)
}

func (*CSSStyleDeclaration) GetBorderTopRightRadius() (borderTopRightRadius string) {
	js.Rewrite("$<.borderTopRightRadius")
	return borderTopRightRadius
}

func (*CSSStyleDeclaration) SetBorderTopRightRadius(borderTopRightRadius string) {
	js.Rewrite("$<.borderTopRightRadius = $1", borderTopRightRadius)
}

func (*CSSStyleDeclaration) GetBorderTopStyle() (borderTopStyle string) {
	js.Rewrite("$<.borderTopStyle")
	return borderTopStyle
}

func (*CSSStyleDeclaration) SetBorderTopStyle(borderTopStyle string) {
	js.Rewrite("$<.borderTopStyle = $1", borderTopStyle)
}

func (*CSSStyleDeclaration) GetBorderTopWidth() (borderTopWidth string) {
	js.Rewrite("$<.borderTopWidth")
	return borderTopWidth
}

func (*CSSStyleDeclaration) SetBorderTopWidth(borderTopWidth string) {
	js.Rewrite("$<.borderTopWidth = $1", borderTopWidth)
}

func (*CSSStyleDeclaration) GetBorderWidth() (borderWidth string) {
	js.Rewrite("$<.borderWidth")
	return borderWidth
}

func (*CSSStyleDeclaration) SetBorderWidth(borderWidth string) {
	js.Rewrite("$<.borderWidth = $1", borderWidth)
}

func (*CSSStyleDeclaration) GetBottom() (bottom string) {
	js.Rewrite("$<.bottom")
	return bottom
}

func (*CSSStyleDeclaration) SetBottom(bottom string) {
	js.Rewrite("$<.bottom = $1", bottom)
}

func (*CSSStyleDeclaration) GetBoxShadow() (boxShadow string) {
	js.Rewrite("$<.boxShadow")
	return boxShadow
}

func (*CSSStyleDeclaration) SetBoxShadow(boxShadow string) {
	js.Rewrite("$<.boxShadow = $1", boxShadow)
}

func (*CSSStyleDeclaration) GetBoxSizing() (boxSizing string) {
	js.Rewrite("$<.boxSizing")
	return boxSizing
}

func (*CSSStyleDeclaration) SetBoxSizing(boxSizing string) {
	js.Rewrite("$<.boxSizing = $1", boxSizing)
}

func (*CSSStyleDeclaration) GetBreakAfter() (breakAfter string) {
	js.Rewrite("$<.breakAfter")
	return breakAfter
}

func (*CSSStyleDeclaration) SetBreakAfter(breakAfter string) {
	js.Rewrite("$<.breakAfter = $1", breakAfter)
}

func (*CSSStyleDeclaration) GetBreakBefore() (breakBefore string) {
	js.Rewrite("$<.breakBefore")
	return breakBefore
}

func (*CSSStyleDeclaration) SetBreakBefore(breakBefore string) {
	js.Rewrite("$<.breakBefore = $1", breakBefore)
}

func (*CSSStyleDeclaration) GetBreakInside() (breakInside string) {
	js.Rewrite("$<.breakInside")
	return breakInside
}

func (*CSSStyleDeclaration) SetBreakInside(breakInside string) {
	js.Rewrite("$<.breakInside = $1", breakInside)
}

func (*CSSStyleDeclaration) GetCaptionSide() (captionSide string) {
	js.Rewrite("$<.captionSide")
	return captionSide
}

func (*CSSStyleDeclaration) SetCaptionSide(captionSide string) {
	js.Rewrite("$<.captionSide = $1", captionSide)
}

func (*CSSStyleDeclaration) GetClear() (clear string) {
	js.Rewrite("$<.clear")
	return clear
}

func (*CSSStyleDeclaration) SetClear(clear string) {
	js.Rewrite("$<.clear = $1", clear)
}

func (*CSSStyleDeclaration) GetClip() (clip string) {
	js.Rewrite("$<.clip")
	return clip
}

func (*CSSStyleDeclaration) SetClip(clip string) {
	js.Rewrite("$<.clip = $1", clip)
}

func (*CSSStyleDeclaration) GetClipPath() (clipPath string) {
	js.Rewrite("$<.clipPath")
	return clipPath
}

func (*CSSStyleDeclaration) SetClipPath(clipPath string) {
	js.Rewrite("$<.clipPath = $1", clipPath)
}

func (*CSSStyleDeclaration) GetClipRule() (clipRule string) {
	js.Rewrite("$<.clipRule")
	return clipRule
}

func (*CSSStyleDeclaration) SetClipRule(clipRule string) {
	js.Rewrite("$<.clipRule = $1", clipRule)
}

func (*CSSStyleDeclaration) GetColor() (color string) {
	js.Rewrite("$<.color")
	return color
}

func (*CSSStyleDeclaration) SetColor(color string) {
	js.Rewrite("$<.color = $1", color)
}

func (*CSSStyleDeclaration) GetColorInterpolationFilters() (colorInterpolationFilters string) {
	js.Rewrite("$<.colorInterpolationFilters")
	return colorInterpolationFilters
}

func (*CSSStyleDeclaration) SetColorInterpolationFilters(colorInterpolationFilters string) {
	js.Rewrite("$<.colorInterpolationFilters = $1", colorInterpolationFilters)
}

func (*CSSStyleDeclaration) GetColumnCount() (columnCount interface{}) {
	js.Rewrite("$<.columnCount")
	return columnCount
}

func (*CSSStyleDeclaration) SetColumnCount(columnCount interface{}) {
	js.Rewrite("$<.columnCount = $1", columnCount)
}

func (*CSSStyleDeclaration) GetColumnFill() (columnFill string) {
	js.Rewrite("$<.columnFill")
	return columnFill
}

func (*CSSStyleDeclaration) SetColumnFill(columnFill string) {
	js.Rewrite("$<.columnFill = $1", columnFill)
}

func (*CSSStyleDeclaration) GetColumnGap() (columnGap interface{}) {
	js.Rewrite("$<.columnGap")
	return columnGap
}

func (*CSSStyleDeclaration) SetColumnGap(columnGap interface{}) {
	js.Rewrite("$<.columnGap = $1", columnGap)
}

func (*CSSStyleDeclaration) GetColumnRule() (columnRule string) {
	js.Rewrite("$<.columnRule")
	return columnRule
}

func (*CSSStyleDeclaration) SetColumnRule(columnRule string) {
	js.Rewrite("$<.columnRule = $1", columnRule)
}

func (*CSSStyleDeclaration) GetColumnRuleColor() (columnRuleColor interface{}) {
	js.Rewrite("$<.columnRuleColor")
	return columnRuleColor
}

func (*CSSStyleDeclaration) SetColumnRuleColor(columnRuleColor interface{}) {
	js.Rewrite("$<.columnRuleColor = $1", columnRuleColor)
}

func (*CSSStyleDeclaration) GetColumnRuleStyle() (columnRuleStyle string) {
	js.Rewrite("$<.columnRuleStyle")
	return columnRuleStyle
}

func (*CSSStyleDeclaration) SetColumnRuleStyle(columnRuleStyle string) {
	js.Rewrite("$<.columnRuleStyle = $1", columnRuleStyle)
}

func (*CSSStyleDeclaration) GetColumnRuleWidth() (columnRuleWidth interface{}) {
	js.Rewrite("$<.columnRuleWidth")
	return columnRuleWidth
}

func (*CSSStyleDeclaration) SetColumnRuleWidth(columnRuleWidth interface{}) {
	js.Rewrite("$<.columnRuleWidth = $1", columnRuleWidth)
}

func (*CSSStyleDeclaration) GetColumns() (columns string) {
	js.Rewrite("$<.columns")
	return columns
}

func (*CSSStyleDeclaration) SetColumns(columns string) {
	js.Rewrite("$<.columns = $1", columns)
}

func (*CSSStyleDeclaration) GetColumnSpan() (columnSpan string) {
	js.Rewrite("$<.columnSpan")
	return columnSpan
}

func (*CSSStyleDeclaration) SetColumnSpan(columnSpan string) {
	js.Rewrite("$<.columnSpan = $1", columnSpan)
}

func (*CSSStyleDeclaration) GetColumnWidth() (columnWidth interface{}) {
	js.Rewrite("$<.columnWidth")
	return columnWidth
}

func (*CSSStyleDeclaration) SetColumnWidth(columnWidth interface{}) {
	js.Rewrite("$<.columnWidth = $1", columnWidth)
}

func (*CSSStyleDeclaration) GetContent() (content string) {
	js.Rewrite("$<.content")
	return content
}

func (*CSSStyleDeclaration) SetContent(content string) {
	js.Rewrite("$<.content = $1", content)
}

func (*CSSStyleDeclaration) GetCounterIncrement() (counterIncrement string) {
	js.Rewrite("$<.counterIncrement")
	return counterIncrement
}

func (*CSSStyleDeclaration) SetCounterIncrement(counterIncrement string) {
	js.Rewrite("$<.counterIncrement = $1", counterIncrement)
}

func (*CSSStyleDeclaration) GetCounterReset() (counterReset string) {
	js.Rewrite("$<.counterReset")
	return counterReset
}

func (*CSSStyleDeclaration) SetCounterReset(counterReset string) {
	js.Rewrite("$<.counterReset = $1", counterReset)
}

func (*CSSStyleDeclaration) GetCSSFloat() (cssFloat string) {
	js.Rewrite("$<.cssFloat")
	return cssFloat
}

func (*CSSStyleDeclaration) SetCSSFloat(cssFloat string) {
	js.Rewrite("$<.cssFloat = $1", cssFloat)
}

func (*CSSStyleDeclaration) GetCSSText() (cssText string) {
	js.Rewrite("$<.cssText")
	return cssText
}

func (*CSSStyleDeclaration) SetCSSText(cssText string) {
	js.Rewrite("$<.cssText = $1", cssText)
}

func (*CSSStyleDeclaration) GetCursor() (cursor string) {
	js.Rewrite("$<.cursor")
	return cursor
}

func (*CSSStyleDeclaration) SetCursor(cursor string) {
	js.Rewrite("$<.cursor = $1", cursor)
}

func (*CSSStyleDeclaration) GetDirection() (direction string) {
	js.Rewrite("$<.direction")
	return direction
}

func (*CSSStyleDeclaration) SetDirection(direction string) {
	js.Rewrite("$<.direction = $1", direction)
}

func (*CSSStyleDeclaration) GetDisplay() (display string) {
	js.Rewrite("$<.display")
	return display
}

func (*CSSStyleDeclaration) SetDisplay(display string) {
	js.Rewrite("$<.display = $1", display)
}

func (*CSSStyleDeclaration) GetDominantBaseline() (dominantBaseline string) {
	js.Rewrite("$<.dominantBaseline")
	return dominantBaseline
}

func (*CSSStyleDeclaration) SetDominantBaseline(dominantBaseline string) {
	js.Rewrite("$<.dominantBaseline = $1", dominantBaseline)
}

func (*CSSStyleDeclaration) GetEmptyCells() (emptyCells string) {
	js.Rewrite("$<.emptyCells")
	return emptyCells
}

func (*CSSStyleDeclaration) SetEmptyCells(emptyCells string) {
	js.Rewrite("$<.emptyCells = $1", emptyCells)
}

func (*CSSStyleDeclaration) GetEnableBackground() (enableBackground string) {
	js.Rewrite("$<.enableBackground")
	return enableBackground
}

func (*CSSStyleDeclaration) SetEnableBackground(enableBackground string) {
	js.Rewrite("$<.enableBackground = $1", enableBackground)
}

func (*CSSStyleDeclaration) GetFill() (fill string) {
	js.Rewrite("$<.fill")
	return fill
}

func (*CSSStyleDeclaration) SetFill(fill string) {
	js.Rewrite("$<.fill = $1", fill)
}

func (*CSSStyleDeclaration) GetFillOpacity() (fillOpacity string) {
	js.Rewrite("$<.fillOpacity")
	return fillOpacity
}

func (*CSSStyleDeclaration) SetFillOpacity(fillOpacity string) {
	js.Rewrite("$<.fillOpacity = $1", fillOpacity)
}

func (*CSSStyleDeclaration) GetFillRule() (fillRule string) {
	js.Rewrite("$<.fillRule")
	return fillRule
}

func (*CSSStyleDeclaration) SetFillRule(fillRule string) {
	js.Rewrite("$<.fillRule = $1", fillRule)
}

func (*CSSStyleDeclaration) GetFilter() (filter string) {
	js.Rewrite("$<.filter")
	return filter
}

func (*CSSStyleDeclaration) SetFilter(filter string) {
	js.Rewrite("$<.filter = $1", filter)
}

func (*CSSStyleDeclaration) GetFlex() (flex string) {
	js.Rewrite("$<.flex")
	return flex
}

func (*CSSStyleDeclaration) SetFlex(flex string) {
	js.Rewrite("$<.flex = $1", flex)
}

func (*CSSStyleDeclaration) GetFlexBasis() (flexBasis string) {
	js.Rewrite("$<.flexBasis")
	return flexBasis
}

func (*CSSStyleDeclaration) SetFlexBasis(flexBasis string) {
	js.Rewrite("$<.flexBasis = $1", flexBasis)
}

func (*CSSStyleDeclaration) GetFlexDirection() (flexDirection string) {
	js.Rewrite("$<.flexDirection")
	return flexDirection
}

func (*CSSStyleDeclaration) SetFlexDirection(flexDirection string) {
	js.Rewrite("$<.flexDirection = $1", flexDirection)
}

func (*CSSStyleDeclaration) GetFlexFlow() (flexFlow string) {
	js.Rewrite("$<.flexFlow")
	return flexFlow
}

func (*CSSStyleDeclaration) SetFlexFlow(flexFlow string) {
	js.Rewrite("$<.flexFlow = $1", flexFlow)
}

func (*CSSStyleDeclaration) GetFlexGrow() (flexGrow string) {
	js.Rewrite("$<.flexGrow")
	return flexGrow
}

func (*CSSStyleDeclaration) SetFlexGrow(flexGrow string) {
	js.Rewrite("$<.flexGrow = $1", flexGrow)
}

func (*CSSStyleDeclaration) GetFlexShrink() (flexShrink string) {
	js.Rewrite("$<.flexShrink")
	return flexShrink
}

func (*CSSStyleDeclaration) SetFlexShrink(flexShrink string) {
	js.Rewrite("$<.flexShrink = $1", flexShrink)
}

func (*CSSStyleDeclaration) GetFlexWrap() (flexWrap string) {
	js.Rewrite("$<.flexWrap")
	return flexWrap
}

func (*CSSStyleDeclaration) SetFlexWrap(flexWrap string) {
	js.Rewrite("$<.flexWrap = $1", flexWrap)
}

func (*CSSStyleDeclaration) GetFloodColor() (floodColor string) {
	js.Rewrite("$<.floodColor")
	return floodColor
}

func (*CSSStyleDeclaration) SetFloodColor(floodColor string) {
	js.Rewrite("$<.floodColor = $1", floodColor)
}

func (*CSSStyleDeclaration) GetFloodOpacity() (floodOpacity string) {
	js.Rewrite("$<.floodOpacity")
	return floodOpacity
}

func (*CSSStyleDeclaration) SetFloodOpacity(floodOpacity string) {
	js.Rewrite("$<.floodOpacity = $1", floodOpacity)
}

func (*CSSStyleDeclaration) GetFont() (font string) {
	js.Rewrite("$<.font")
	return font
}

func (*CSSStyleDeclaration) SetFont(font string) {
	js.Rewrite("$<.font = $1", font)
}

func (*CSSStyleDeclaration) GetFontFamily() (fontFamily string) {
	js.Rewrite("$<.fontFamily")
	return fontFamily
}

func (*CSSStyleDeclaration) SetFontFamily(fontFamily string) {
	js.Rewrite("$<.fontFamily = $1", fontFamily)
}

func (*CSSStyleDeclaration) GetFontFeatureSettings() (fontFeatureSettings string) {
	js.Rewrite("$<.fontFeatureSettings")
	return fontFeatureSettings
}

func (*CSSStyleDeclaration) SetFontFeatureSettings(fontFeatureSettings string) {
	js.Rewrite("$<.fontFeatureSettings = $1", fontFeatureSettings)
}

func (*CSSStyleDeclaration) GetFontSize() (fontSize string) {
	js.Rewrite("$<.fontSize")
	return fontSize
}

func (*CSSStyleDeclaration) SetFontSize(fontSize string) {
	js.Rewrite("$<.fontSize = $1", fontSize)
}

func (*CSSStyleDeclaration) GetFontSizeAdjust() (fontSizeAdjust string) {
	js.Rewrite("$<.fontSizeAdjust")
	return fontSizeAdjust
}

func (*CSSStyleDeclaration) SetFontSizeAdjust(fontSizeAdjust string) {
	js.Rewrite("$<.fontSizeAdjust = $1", fontSizeAdjust)
}

func (*CSSStyleDeclaration) GetFontStretch() (fontStretch string) {
	js.Rewrite("$<.fontStretch")
	return fontStretch
}

func (*CSSStyleDeclaration) SetFontStretch(fontStretch string) {
	js.Rewrite("$<.fontStretch = $1", fontStretch)
}

func (*CSSStyleDeclaration) GetFontStyle() (fontStyle string) {
	js.Rewrite("$<.fontStyle")
	return fontStyle
}

func (*CSSStyleDeclaration) SetFontStyle(fontStyle string) {
	js.Rewrite("$<.fontStyle = $1", fontStyle)
}

func (*CSSStyleDeclaration) GetFontVariant() (fontVariant string) {
	js.Rewrite("$<.fontVariant")
	return fontVariant
}

func (*CSSStyleDeclaration) SetFontVariant(fontVariant string) {
	js.Rewrite("$<.fontVariant = $1", fontVariant)
}

func (*CSSStyleDeclaration) GetFontWeight() (fontWeight string) {
	js.Rewrite("$<.fontWeight")
	return fontWeight
}

func (*CSSStyleDeclaration) SetFontWeight(fontWeight string) {
	js.Rewrite("$<.fontWeight = $1", fontWeight)
}

func (*CSSStyleDeclaration) GetGlyphOrientationHorizontal() (glyphOrientationHorizontal string) {
	js.Rewrite("$<.glyphOrientationHorizontal")
	return glyphOrientationHorizontal
}

func (*CSSStyleDeclaration) SetGlyphOrientationHorizontal(glyphOrientationHorizontal string) {
	js.Rewrite("$<.glyphOrientationHorizontal = $1", glyphOrientationHorizontal)
}

func (*CSSStyleDeclaration) GetGlyphOrientationVertical() (glyphOrientationVertical string) {
	js.Rewrite("$<.glyphOrientationVertical")
	return glyphOrientationVertical
}

func (*CSSStyleDeclaration) SetGlyphOrientationVertical(glyphOrientationVertical string) {
	js.Rewrite("$<.glyphOrientationVertical = $1", glyphOrientationVertical)
}

func (*CSSStyleDeclaration) GetHeight() (height string) {
	js.Rewrite("$<.height")
	return height
}

func (*CSSStyleDeclaration) SetHeight(height string) {
	js.Rewrite("$<.height = $1", height)
}

func (*CSSStyleDeclaration) GetImeMode() (imeMode string) {
	js.Rewrite("$<.imeMode")
	return imeMode
}

func (*CSSStyleDeclaration) SetImeMode(imeMode string) {
	js.Rewrite("$<.imeMode = $1", imeMode)
}

func (*CSSStyleDeclaration) GetJustifyContent() (justifyContent string) {
	js.Rewrite("$<.justifyContent")
	return justifyContent
}

func (*CSSStyleDeclaration) SetJustifyContent(justifyContent string) {
	js.Rewrite("$<.justifyContent = $1", justifyContent)
}

func (*CSSStyleDeclaration) GetKerning() (kerning string) {
	js.Rewrite("$<.kerning")
	return kerning
}

func (*CSSStyleDeclaration) SetKerning(kerning string) {
	js.Rewrite("$<.kerning = $1", kerning)
}

func (*CSSStyleDeclaration) GetLayoutGrid() (layoutGrid string) {
	js.Rewrite("$<.layoutGrid")
	return layoutGrid
}

func (*CSSStyleDeclaration) SetLayoutGrid(layoutGrid string) {
	js.Rewrite("$<.layoutGrid = $1", layoutGrid)
}

func (*CSSStyleDeclaration) GetLayoutGridChar() (layoutGridChar string) {
	js.Rewrite("$<.layoutGridChar")
	return layoutGridChar
}

func (*CSSStyleDeclaration) SetLayoutGridChar(layoutGridChar string) {
	js.Rewrite("$<.layoutGridChar = $1", layoutGridChar)
}

func (*CSSStyleDeclaration) GetLayoutGridLine() (layoutGridLine string) {
	js.Rewrite("$<.layoutGridLine")
	return layoutGridLine
}

func (*CSSStyleDeclaration) SetLayoutGridLine(layoutGridLine string) {
	js.Rewrite("$<.layoutGridLine = $1", layoutGridLine)
}

func (*CSSStyleDeclaration) GetLayoutGridMode() (layoutGridMode string) {
	js.Rewrite("$<.layoutGridMode")
	return layoutGridMode
}

func (*CSSStyleDeclaration) SetLayoutGridMode(layoutGridMode string) {
	js.Rewrite("$<.layoutGridMode = $1", layoutGridMode)
}

func (*CSSStyleDeclaration) GetLayoutGridType() (layoutGridType string) {
	js.Rewrite("$<.layoutGridType")
	return layoutGridType
}

func (*CSSStyleDeclaration) SetLayoutGridType(layoutGridType string) {
	js.Rewrite("$<.layoutGridType = $1", layoutGridType)
}

func (*CSSStyleDeclaration) GetLeft() (left string) {
	js.Rewrite("$<.left")
	return left
}

func (*CSSStyleDeclaration) SetLeft(left string) {
	js.Rewrite("$<.left = $1", left)
}

func (*CSSStyleDeclaration) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}

func (*CSSStyleDeclaration) GetLetterSpacing() (letterSpacing string) {
	js.Rewrite("$<.letterSpacing")
	return letterSpacing
}

func (*CSSStyleDeclaration) SetLetterSpacing(letterSpacing string) {
	js.Rewrite("$<.letterSpacing = $1", letterSpacing)
}

func (*CSSStyleDeclaration) GetLightingColor() (lightingColor string) {
	js.Rewrite("$<.lightingColor")
	return lightingColor
}

func (*CSSStyleDeclaration) SetLightingColor(lightingColor string) {
	js.Rewrite("$<.lightingColor = $1", lightingColor)
}

func (*CSSStyleDeclaration) GetLineBreak() (lineBreak string) {
	js.Rewrite("$<.lineBreak")
	return lineBreak
}

func (*CSSStyleDeclaration) SetLineBreak(lineBreak string) {
	js.Rewrite("$<.lineBreak = $1", lineBreak)
}

func (*CSSStyleDeclaration) GetLineHeight() (lineHeight string) {
	js.Rewrite("$<.lineHeight")
	return lineHeight
}

func (*CSSStyleDeclaration) SetLineHeight(lineHeight string) {
	js.Rewrite("$<.lineHeight = $1", lineHeight)
}

func (*CSSStyleDeclaration) GetListStyle() (listStyle string) {
	js.Rewrite("$<.listStyle")
	return listStyle
}

func (*CSSStyleDeclaration) SetListStyle(listStyle string) {
	js.Rewrite("$<.listStyle = $1", listStyle)
}

func (*CSSStyleDeclaration) GetListStyleImage() (listStyleImage string) {
	js.Rewrite("$<.listStyleImage")
	return listStyleImage
}

func (*CSSStyleDeclaration) SetListStyleImage(listStyleImage string) {
	js.Rewrite("$<.listStyleImage = $1", listStyleImage)
}

func (*CSSStyleDeclaration) GetListStylePosition() (listStylePosition string) {
	js.Rewrite("$<.listStylePosition")
	return listStylePosition
}

func (*CSSStyleDeclaration) SetListStylePosition(listStylePosition string) {
	js.Rewrite("$<.listStylePosition = $1", listStylePosition)
}

func (*CSSStyleDeclaration) GetListStyleType() (listStyleType string) {
	js.Rewrite("$<.listStyleType")
	return listStyleType
}

func (*CSSStyleDeclaration) SetListStyleType(listStyleType string) {
	js.Rewrite("$<.listStyleType = $1", listStyleType)
}

func (*CSSStyleDeclaration) GetMargin() (margin string) {
	js.Rewrite("$<.margin")
	return margin
}

func (*CSSStyleDeclaration) SetMargin(margin string) {
	js.Rewrite("$<.margin = $1", margin)
}

func (*CSSStyleDeclaration) GetMarginBottom() (marginBottom string) {
	js.Rewrite("$<.marginBottom")
	return marginBottom
}

func (*CSSStyleDeclaration) SetMarginBottom(marginBottom string) {
	js.Rewrite("$<.marginBottom = $1", marginBottom)
}

func (*CSSStyleDeclaration) GetMarginLeft() (marginLeft string) {
	js.Rewrite("$<.marginLeft")
	return marginLeft
}

func (*CSSStyleDeclaration) SetMarginLeft(marginLeft string) {
	js.Rewrite("$<.marginLeft = $1", marginLeft)
}

func (*CSSStyleDeclaration) GetMarginRight() (marginRight string) {
	js.Rewrite("$<.marginRight")
	return marginRight
}

func (*CSSStyleDeclaration) SetMarginRight(marginRight string) {
	js.Rewrite("$<.marginRight = $1", marginRight)
}

func (*CSSStyleDeclaration) GetMarginTop() (marginTop string) {
	js.Rewrite("$<.marginTop")
	return marginTop
}

func (*CSSStyleDeclaration) SetMarginTop(marginTop string) {
	js.Rewrite("$<.marginTop = $1", marginTop)
}

func (*CSSStyleDeclaration) GetMarker() (marker string) {
	js.Rewrite("$<.marker")
	return marker
}

func (*CSSStyleDeclaration) SetMarker(marker string) {
	js.Rewrite("$<.marker = $1", marker)
}

func (*CSSStyleDeclaration) GetMarkerEnd() (markerEnd string) {
	js.Rewrite("$<.markerEnd")
	return markerEnd
}

func (*CSSStyleDeclaration) SetMarkerEnd(markerEnd string) {
	js.Rewrite("$<.markerEnd = $1", markerEnd)
}

func (*CSSStyleDeclaration) GetMarkerMid() (markerMid string) {
	js.Rewrite("$<.markerMid")
	return markerMid
}

func (*CSSStyleDeclaration) SetMarkerMid(markerMid string) {
	js.Rewrite("$<.markerMid = $1", markerMid)
}

func (*CSSStyleDeclaration) GetMarkerStart() (markerStart string) {
	js.Rewrite("$<.markerStart")
	return markerStart
}

func (*CSSStyleDeclaration) SetMarkerStart(markerStart string) {
	js.Rewrite("$<.markerStart = $1", markerStart)
}

func (*CSSStyleDeclaration) GetMask() (mask string) {
	js.Rewrite("$<.mask")
	return mask
}

func (*CSSStyleDeclaration) SetMask(mask string) {
	js.Rewrite("$<.mask = $1", mask)
}

func (*CSSStyleDeclaration) GetMaxHeight() (maxHeight string) {
	js.Rewrite("$<.maxHeight")
	return maxHeight
}

func (*CSSStyleDeclaration) SetMaxHeight(maxHeight string) {
	js.Rewrite("$<.maxHeight = $1", maxHeight)
}

func (*CSSStyleDeclaration) GetMaxWidth() (maxWidth string) {
	js.Rewrite("$<.maxWidth")
	return maxWidth
}

func (*CSSStyleDeclaration) SetMaxWidth(maxWidth string) {
	js.Rewrite("$<.maxWidth = $1", maxWidth)
}

func (*CSSStyleDeclaration) GetMinHeight() (minHeight string) {
	js.Rewrite("$<.minHeight")
	return minHeight
}

func (*CSSStyleDeclaration) SetMinHeight(minHeight string) {
	js.Rewrite("$<.minHeight = $1", minHeight)
}

func (*CSSStyleDeclaration) GetMinWidth() (minWidth string) {
	js.Rewrite("$<.minWidth")
	return minWidth
}

func (*CSSStyleDeclaration) SetMinWidth(minWidth string) {
	js.Rewrite("$<.minWidth = $1", minWidth)
}

func (*CSSStyleDeclaration) GetMsContentZoomChaining() (msContentZoomChaining string) {
	js.Rewrite("$<.msContentZoomChaining")
	return msContentZoomChaining
}

func (*CSSStyleDeclaration) SetMsContentZoomChaining(msContentZoomChaining string) {
	js.Rewrite("$<.msContentZoomChaining = $1", msContentZoomChaining)
}

func (*CSSStyleDeclaration) GetMsContentZooming() (msContentZooming string) {
	js.Rewrite("$<.msContentZooming")
	return msContentZooming
}

func (*CSSStyleDeclaration) SetMsContentZooming(msContentZooming string) {
	js.Rewrite("$<.msContentZooming = $1", msContentZooming)
}

func (*CSSStyleDeclaration) GetMsContentZoomLimit() (msContentZoomLimit string) {
	js.Rewrite("$<.msContentZoomLimit")
	return msContentZoomLimit
}

func (*CSSStyleDeclaration) SetMsContentZoomLimit(msContentZoomLimit string) {
	js.Rewrite("$<.msContentZoomLimit = $1", msContentZoomLimit)
}

func (*CSSStyleDeclaration) GetMsContentZoomLimitMax() (msContentZoomLimitMax interface{}) {
	js.Rewrite("$<.msContentZoomLimitMax")
	return msContentZoomLimitMax
}

func (*CSSStyleDeclaration) SetMsContentZoomLimitMax(msContentZoomLimitMax interface{}) {
	js.Rewrite("$<.msContentZoomLimitMax = $1", msContentZoomLimitMax)
}

func (*CSSStyleDeclaration) GetMsContentZoomLimitMin() (msContentZoomLimitMin interface{}) {
	js.Rewrite("$<.msContentZoomLimitMin")
	return msContentZoomLimitMin
}

func (*CSSStyleDeclaration) SetMsContentZoomLimitMin(msContentZoomLimitMin interface{}) {
	js.Rewrite("$<.msContentZoomLimitMin = $1", msContentZoomLimitMin)
}

func (*CSSStyleDeclaration) GetMsContentZoomSnap() (msContentZoomSnap string) {
	js.Rewrite("$<.msContentZoomSnap")
	return msContentZoomSnap
}

func (*CSSStyleDeclaration) SetMsContentZoomSnap(msContentZoomSnap string) {
	js.Rewrite("$<.msContentZoomSnap = $1", msContentZoomSnap)
}

func (*CSSStyleDeclaration) GetMsContentZoomSnapPoints() (msContentZoomSnapPoints string) {
	js.Rewrite("$<.msContentZoomSnapPoints")
	return msContentZoomSnapPoints
}

func (*CSSStyleDeclaration) SetMsContentZoomSnapPoints(msContentZoomSnapPoints string) {
	js.Rewrite("$<.msContentZoomSnapPoints = $1", msContentZoomSnapPoints)
}

func (*CSSStyleDeclaration) GetMsContentZoomSnapType() (msContentZoomSnapType string) {
	js.Rewrite("$<.msContentZoomSnapType")
	return msContentZoomSnapType
}

func (*CSSStyleDeclaration) SetMsContentZoomSnapType(msContentZoomSnapType string) {
	js.Rewrite("$<.msContentZoomSnapType = $1", msContentZoomSnapType)
}

func (*CSSStyleDeclaration) GetMsFlowFrom() (msFlowFrom string) {
	js.Rewrite("$<.msFlowFrom")
	return msFlowFrom
}

func (*CSSStyleDeclaration) SetMsFlowFrom(msFlowFrom string) {
	js.Rewrite("$<.msFlowFrom = $1", msFlowFrom)
}

func (*CSSStyleDeclaration) GetMsFlowInto() (msFlowInto string) {
	js.Rewrite("$<.msFlowInto")
	return msFlowInto
}

func (*CSSStyleDeclaration) SetMsFlowInto(msFlowInto string) {
	js.Rewrite("$<.msFlowInto = $1", msFlowInto)
}

func (*CSSStyleDeclaration) GetMsFontFeatureSettings() (msFontFeatureSettings string) {
	js.Rewrite("$<.msFontFeatureSettings")
	return msFontFeatureSettings
}

func (*CSSStyleDeclaration) SetMsFontFeatureSettings(msFontFeatureSettings string) {
	js.Rewrite("$<.msFontFeatureSettings = $1", msFontFeatureSettings)
}

func (*CSSStyleDeclaration) GetMsGridColumn() (msGridColumn interface{}) {
	js.Rewrite("$<.msGridColumn")
	return msGridColumn
}

func (*CSSStyleDeclaration) SetMsGridColumn(msGridColumn interface{}) {
	js.Rewrite("$<.msGridColumn = $1", msGridColumn)
}

func (*CSSStyleDeclaration) GetMsGridColumnAlign() (msGridColumnAlign string) {
	js.Rewrite("$<.msGridColumnAlign")
	return msGridColumnAlign
}

func (*CSSStyleDeclaration) SetMsGridColumnAlign(msGridColumnAlign string) {
	js.Rewrite("$<.msGridColumnAlign = $1", msGridColumnAlign)
}

func (*CSSStyleDeclaration) GetMsGridColumns() (msGridColumns string) {
	js.Rewrite("$<.msGridColumns")
	return msGridColumns
}

func (*CSSStyleDeclaration) SetMsGridColumns(msGridColumns string) {
	js.Rewrite("$<.msGridColumns = $1", msGridColumns)
}

func (*CSSStyleDeclaration) GetMsGridColumnSpan() (msGridColumnSpan interface{}) {
	js.Rewrite("$<.msGridColumnSpan")
	return msGridColumnSpan
}

func (*CSSStyleDeclaration) SetMsGridColumnSpan(msGridColumnSpan interface{}) {
	js.Rewrite("$<.msGridColumnSpan = $1", msGridColumnSpan)
}

func (*CSSStyleDeclaration) GetMsGridRow() (msGridRow interface{}) {
	js.Rewrite("$<.msGridRow")
	return msGridRow
}

func (*CSSStyleDeclaration) SetMsGridRow(msGridRow interface{}) {
	js.Rewrite("$<.msGridRow = $1", msGridRow)
}

func (*CSSStyleDeclaration) GetMsGridRowAlign() (msGridRowAlign string) {
	js.Rewrite("$<.msGridRowAlign")
	return msGridRowAlign
}

func (*CSSStyleDeclaration) SetMsGridRowAlign(msGridRowAlign string) {
	js.Rewrite("$<.msGridRowAlign = $1", msGridRowAlign)
}

func (*CSSStyleDeclaration) GetMsGridRows() (msGridRows string) {
	js.Rewrite("$<.msGridRows")
	return msGridRows
}

func (*CSSStyleDeclaration) SetMsGridRows(msGridRows string) {
	js.Rewrite("$<.msGridRows = $1", msGridRows)
}

func (*CSSStyleDeclaration) GetMsGridRowSpan() (msGridRowSpan interface{}) {
	js.Rewrite("$<.msGridRowSpan")
	return msGridRowSpan
}

func (*CSSStyleDeclaration) SetMsGridRowSpan(msGridRowSpan interface{}) {
	js.Rewrite("$<.msGridRowSpan = $1", msGridRowSpan)
}

func (*CSSStyleDeclaration) GetMsHighContrastAdjust() (msHighContrastAdjust string) {
	js.Rewrite("$<.msHighContrastAdjust")
	return msHighContrastAdjust
}

func (*CSSStyleDeclaration) SetMsHighContrastAdjust(msHighContrastAdjust string) {
	js.Rewrite("$<.msHighContrastAdjust = $1", msHighContrastAdjust)
}

func (*CSSStyleDeclaration) GetMsHyphenateLimitChars() (msHyphenateLimitChars string) {
	js.Rewrite("$<.msHyphenateLimitChars")
	return msHyphenateLimitChars
}

func (*CSSStyleDeclaration) SetMsHyphenateLimitChars(msHyphenateLimitChars string) {
	js.Rewrite("$<.msHyphenateLimitChars = $1", msHyphenateLimitChars)
}

func (*CSSStyleDeclaration) GetMsHyphenateLimitLines() (msHyphenateLimitLines interface{}) {
	js.Rewrite("$<.msHyphenateLimitLines")
	return msHyphenateLimitLines
}

func (*CSSStyleDeclaration) SetMsHyphenateLimitLines(msHyphenateLimitLines interface{}) {
	js.Rewrite("$<.msHyphenateLimitLines = $1", msHyphenateLimitLines)
}

func (*CSSStyleDeclaration) GetMsHyphenateLimitZone() (msHyphenateLimitZone interface{}) {
	js.Rewrite("$<.msHyphenateLimitZone")
	return msHyphenateLimitZone
}

func (*CSSStyleDeclaration) SetMsHyphenateLimitZone(msHyphenateLimitZone interface{}) {
	js.Rewrite("$<.msHyphenateLimitZone = $1", msHyphenateLimitZone)
}

func (*CSSStyleDeclaration) GetMsHyphens() (msHyphens string) {
	js.Rewrite("$<.msHyphens")
	return msHyphens
}

func (*CSSStyleDeclaration) SetMsHyphens(msHyphens string) {
	js.Rewrite("$<.msHyphens = $1", msHyphens)
}

func (*CSSStyleDeclaration) GetMsImeAlign() (msImeAlign string) {
	js.Rewrite("$<.msImeAlign")
	return msImeAlign
}

func (*CSSStyleDeclaration) SetMsImeAlign(msImeAlign string) {
	js.Rewrite("$<.msImeAlign = $1", msImeAlign)
}

func (*CSSStyleDeclaration) GetMsOverflowStyle() (msOverflowStyle string) {
	js.Rewrite("$<.msOverflowStyle")
	return msOverflowStyle
}

func (*CSSStyleDeclaration) SetMsOverflowStyle(msOverflowStyle string) {
	js.Rewrite("$<.msOverflowStyle = $1", msOverflowStyle)
}

func (*CSSStyleDeclaration) GetMsScrollChaining() (msScrollChaining string) {
	js.Rewrite("$<.msScrollChaining")
	return msScrollChaining
}

func (*CSSStyleDeclaration) SetMsScrollChaining(msScrollChaining string) {
	js.Rewrite("$<.msScrollChaining = $1", msScrollChaining)
}

func (*CSSStyleDeclaration) GetMsScrollLimit() (msScrollLimit string) {
	js.Rewrite("$<.msScrollLimit")
	return msScrollLimit
}

func (*CSSStyleDeclaration) SetMsScrollLimit(msScrollLimit string) {
	js.Rewrite("$<.msScrollLimit = $1", msScrollLimit)
}

func (*CSSStyleDeclaration) GetMsScrollLimitXMax() (msScrollLimitXMax interface{}) {
	js.Rewrite("$<.msScrollLimitXMax")
	return msScrollLimitXMax
}

func (*CSSStyleDeclaration) SetMsScrollLimitXMax(msScrollLimitXMax interface{}) {
	js.Rewrite("$<.msScrollLimitXMax = $1", msScrollLimitXMax)
}

func (*CSSStyleDeclaration) GetMsScrollLimitXMin() (msScrollLimitXMin interface{}) {
	js.Rewrite("$<.msScrollLimitXMin")
	return msScrollLimitXMin
}

func (*CSSStyleDeclaration) SetMsScrollLimitXMin(msScrollLimitXMin interface{}) {
	js.Rewrite("$<.msScrollLimitXMin = $1", msScrollLimitXMin)
}

func (*CSSStyleDeclaration) GetMsScrollLimitYMax() (msScrollLimitYMax interface{}) {
	js.Rewrite("$<.msScrollLimitYMax")
	return msScrollLimitYMax
}

func (*CSSStyleDeclaration) SetMsScrollLimitYMax(msScrollLimitYMax interface{}) {
	js.Rewrite("$<.msScrollLimitYMax = $1", msScrollLimitYMax)
}

func (*CSSStyleDeclaration) GetMsScrollLimitYMin() (msScrollLimitYMin interface{}) {
	js.Rewrite("$<.msScrollLimitYMin")
	return msScrollLimitYMin
}

func (*CSSStyleDeclaration) SetMsScrollLimitYMin(msScrollLimitYMin interface{}) {
	js.Rewrite("$<.msScrollLimitYMin = $1", msScrollLimitYMin)
}

func (*CSSStyleDeclaration) GetMsScrollRails() (msScrollRails string) {
	js.Rewrite("$<.msScrollRails")
	return msScrollRails
}

func (*CSSStyleDeclaration) SetMsScrollRails(msScrollRails string) {
	js.Rewrite("$<.msScrollRails = $1", msScrollRails)
}

func (*CSSStyleDeclaration) GetMsScrollSnapPointsX() (msScrollSnapPointsX string) {
	js.Rewrite("$<.msScrollSnapPointsX")
	return msScrollSnapPointsX
}

func (*CSSStyleDeclaration) SetMsScrollSnapPointsX(msScrollSnapPointsX string) {
	js.Rewrite("$<.msScrollSnapPointsX = $1", msScrollSnapPointsX)
}

func (*CSSStyleDeclaration) GetMsScrollSnapPointsY() (msScrollSnapPointsY string) {
	js.Rewrite("$<.msScrollSnapPointsY")
	return msScrollSnapPointsY
}

func (*CSSStyleDeclaration) SetMsScrollSnapPointsY(msScrollSnapPointsY string) {
	js.Rewrite("$<.msScrollSnapPointsY = $1", msScrollSnapPointsY)
}

func (*CSSStyleDeclaration) GetMsScrollSnapType() (msScrollSnapType string) {
	js.Rewrite("$<.msScrollSnapType")
	return msScrollSnapType
}

func (*CSSStyleDeclaration) SetMsScrollSnapType(msScrollSnapType string) {
	js.Rewrite("$<.msScrollSnapType = $1", msScrollSnapType)
}

func (*CSSStyleDeclaration) GetMsScrollSnapX() (msScrollSnapX string) {
	js.Rewrite("$<.msScrollSnapX")
	return msScrollSnapX
}

func (*CSSStyleDeclaration) SetMsScrollSnapX(msScrollSnapX string) {
	js.Rewrite("$<.msScrollSnapX = $1", msScrollSnapX)
}

func (*CSSStyleDeclaration) GetMsScrollSnapY() (msScrollSnapY string) {
	js.Rewrite("$<.msScrollSnapY")
	return msScrollSnapY
}

func (*CSSStyleDeclaration) SetMsScrollSnapY(msScrollSnapY string) {
	js.Rewrite("$<.msScrollSnapY = $1", msScrollSnapY)
}

func (*CSSStyleDeclaration) GetMsScrollTranslation() (msScrollTranslation string) {
	js.Rewrite("$<.msScrollTranslation")
	return msScrollTranslation
}

func (*CSSStyleDeclaration) SetMsScrollTranslation(msScrollTranslation string) {
	js.Rewrite("$<.msScrollTranslation = $1", msScrollTranslation)
}

func (*CSSStyleDeclaration) GetMsTextCombineHorizontal() (msTextCombineHorizontal string) {
	js.Rewrite("$<.msTextCombineHorizontal")
	return msTextCombineHorizontal
}

func (*CSSStyleDeclaration) SetMsTextCombineHorizontal(msTextCombineHorizontal string) {
	js.Rewrite("$<.msTextCombineHorizontal = $1", msTextCombineHorizontal)
}

func (*CSSStyleDeclaration) GetMsTextSizeAdjust() (msTextSizeAdjust interface{}) {
	js.Rewrite("$<.msTextSizeAdjust")
	return msTextSizeAdjust
}

func (*CSSStyleDeclaration) SetMsTextSizeAdjust(msTextSizeAdjust interface{}) {
	js.Rewrite("$<.msTextSizeAdjust = $1", msTextSizeAdjust)
}

func (*CSSStyleDeclaration) GetMsTouchAction() (msTouchAction string) {
	js.Rewrite("$<.msTouchAction")
	return msTouchAction
}

func (*CSSStyleDeclaration) SetMsTouchAction(msTouchAction string) {
	js.Rewrite("$<.msTouchAction = $1", msTouchAction)
}

func (*CSSStyleDeclaration) GetMsTouchSelect() (msTouchSelect string) {
	js.Rewrite("$<.msTouchSelect")
	return msTouchSelect
}

func (*CSSStyleDeclaration) SetMsTouchSelect(msTouchSelect string) {
	js.Rewrite("$<.msTouchSelect = $1", msTouchSelect)
}

func (*CSSStyleDeclaration) GetMsUserSelect() (msUserSelect string) {
	js.Rewrite("$<.msUserSelect")
	return msUserSelect
}

func (*CSSStyleDeclaration) SetMsUserSelect(msUserSelect string) {
	js.Rewrite("$<.msUserSelect = $1", msUserSelect)
}

func (*CSSStyleDeclaration) GetMsWrapFlow() (msWrapFlow string) {
	js.Rewrite("$<.msWrapFlow")
	return msWrapFlow
}

func (*CSSStyleDeclaration) SetMsWrapFlow(msWrapFlow string) {
	js.Rewrite("$<.msWrapFlow = $1", msWrapFlow)
}

func (*CSSStyleDeclaration) GetMsWrapMargin() (msWrapMargin interface{}) {
	js.Rewrite("$<.msWrapMargin")
	return msWrapMargin
}

func (*CSSStyleDeclaration) SetMsWrapMargin(msWrapMargin interface{}) {
	js.Rewrite("$<.msWrapMargin = $1", msWrapMargin)
}

func (*CSSStyleDeclaration) GetMsWrapThrough() (msWrapThrough string) {
	js.Rewrite("$<.msWrapThrough")
	return msWrapThrough
}

func (*CSSStyleDeclaration) SetMsWrapThrough(msWrapThrough string) {
	js.Rewrite("$<.msWrapThrough = $1", msWrapThrough)
}

func (*CSSStyleDeclaration) GetOpacity() (opacity string) {
	js.Rewrite("$<.opacity")
	return opacity
}

func (*CSSStyleDeclaration) SetOpacity(opacity string) {
	js.Rewrite("$<.opacity = $1", opacity)
}

func (*CSSStyleDeclaration) GetOrder() (order string) {
	js.Rewrite("$<.order")
	return order
}

func (*CSSStyleDeclaration) SetOrder(order string) {
	js.Rewrite("$<.order = $1", order)
}

func (*CSSStyleDeclaration) GetOrphans() (orphans string) {
	js.Rewrite("$<.orphans")
	return orphans
}

func (*CSSStyleDeclaration) SetOrphans(orphans string) {
	js.Rewrite("$<.orphans = $1", orphans)
}

func (*CSSStyleDeclaration) GetOutline() (outline string) {
	js.Rewrite("$<.outline")
	return outline
}

func (*CSSStyleDeclaration) SetOutline(outline string) {
	js.Rewrite("$<.outline = $1", outline)
}

func (*CSSStyleDeclaration) GetOutlineColor() (outlineColor string) {
	js.Rewrite("$<.outlineColor")
	return outlineColor
}

func (*CSSStyleDeclaration) SetOutlineColor(outlineColor string) {
	js.Rewrite("$<.outlineColor = $1", outlineColor)
}

func (*CSSStyleDeclaration) GetOutlineOffset() (outlineOffset string) {
	js.Rewrite("$<.outlineOffset")
	return outlineOffset
}

func (*CSSStyleDeclaration) SetOutlineOffset(outlineOffset string) {
	js.Rewrite("$<.outlineOffset = $1", outlineOffset)
}

func (*CSSStyleDeclaration) GetOutlineStyle() (outlineStyle string) {
	js.Rewrite("$<.outlineStyle")
	return outlineStyle
}

func (*CSSStyleDeclaration) SetOutlineStyle(outlineStyle string) {
	js.Rewrite("$<.outlineStyle = $1", outlineStyle)
}

func (*CSSStyleDeclaration) GetOutlineWidth() (outlineWidth string) {
	js.Rewrite("$<.outlineWidth")
	return outlineWidth
}

func (*CSSStyleDeclaration) SetOutlineWidth(outlineWidth string) {
	js.Rewrite("$<.outlineWidth = $1", outlineWidth)
}

func (*CSSStyleDeclaration) GetOverflow() (overflow string) {
	js.Rewrite("$<.overflow")
	return overflow
}

func (*CSSStyleDeclaration) SetOverflow(overflow string) {
	js.Rewrite("$<.overflow = $1", overflow)
}

func (*CSSStyleDeclaration) GetOverflowX() (overflowX string) {
	js.Rewrite("$<.overflowX")
	return overflowX
}

func (*CSSStyleDeclaration) SetOverflowX(overflowX string) {
	js.Rewrite("$<.overflowX = $1", overflowX)
}

func (*CSSStyleDeclaration) GetOverflowY() (overflowY string) {
	js.Rewrite("$<.overflowY")
	return overflowY
}

func (*CSSStyleDeclaration) SetOverflowY(overflowY string) {
	js.Rewrite("$<.overflowY = $1", overflowY)
}

func (*CSSStyleDeclaration) GetPadding() (padding string) {
	js.Rewrite("$<.padding")
	return padding
}

func (*CSSStyleDeclaration) SetPadding(padding string) {
	js.Rewrite("$<.padding = $1", padding)
}

func (*CSSStyleDeclaration) GetPaddingBottom() (paddingBottom string) {
	js.Rewrite("$<.paddingBottom")
	return paddingBottom
}

func (*CSSStyleDeclaration) SetPaddingBottom(paddingBottom string) {
	js.Rewrite("$<.paddingBottom = $1", paddingBottom)
}

func (*CSSStyleDeclaration) GetPaddingLeft() (paddingLeft string) {
	js.Rewrite("$<.paddingLeft")
	return paddingLeft
}

func (*CSSStyleDeclaration) SetPaddingLeft(paddingLeft string) {
	js.Rewrite("$<.paddingLeft = $1", paddingLeft)
}

func (*CSSStyleDeclaration) GetPaddingRight() (paddingRight string) {
	js.Rewrite("$<.paddingRight")
	return paddingRight
}

func (*CSSStyleDeclaration) SetPaddingRight(paddingRight string) {
	js.Rewrite("$<.paddingRight = $1", paddingRight)
}

func (*CSSStyleDeclaration) GetPaddingTop() (paddingTop string) {
	js.Rewrite("$<.paddingTop")
	return paddingTop
}

func (*CSSStyleDeclaration) SetPaddingTop(paddingTop string) {
	js.Rewrite("$<.paddingTop = $1", paddingTop)
}

func (*CSSStyleDeclaration) GetPageBreakAfter() (pageBreakAfter string) {
	js.Rewrite("$<.pageBreakAfter")
	return pageBreakAfter
}

func (*CSSStyleDeclaration) SetPageBreakAfter(pageBreakAfter string) {
	js.Rewrite("$<.pageBreakAfter = $1", pageBreakAfter)
}

func (*CSSStyleDeclaration) GetPageBreakBefore() (pageBreakBefore string) {
	js.Rewrite("$<.pageBreakBefore")
	return pageBreakBefore
}

func (*CSSStyleDeclaration) SetPageBreakBefore(pageBreakBefore string) {
	js.Rewrite("$<.pageBreakBefore = $1", pageBreakBefore)
}

func (*CSSStyleDeclaration) GetPageBreakInside() (pageBreakInside string) {
	js.Rewrite("$<.pageBreakInside")
	return pageBreakInside
}

func (*CSSStyleDeclaration) SetPageBreakInside(pageBreakInside string) {
	js.Rewrite("$<.pageBreakInside = $1", pageBreakInside)
}

func (*CSSStyleDeclaration) GetParentRule() (parentRule *cssrule.CSSRule) {
	js.Rewrite("$<.parentRule")
	return parentRule
}

func (*CSSStyleDeclaration) GetPerspective() (perspective string) {
	js.Rewrite("$<.perspective")
	return perspective
}

func (*CSSStyleDeclaration) SetPerspective(perspective string) {
	js.Rewrite("$<.perspective = $1", perspective)
}

func (*CSSStyleDeclaration) GetPerspectiveOrigin() (perspectiveOrigin string) {
	js.Rewrite("$<.perspectiveOrigin")
	return perspectiveOrigin
}

func (*CSSStyleDeclaration) SetPerspectiveOrigin(perspectiveOrigin string) {
	js.Rewrite("$<.perspectiveOrigin = $1", perspectiveOrigin)
}

func (*CSSStyleDeclaration) GetPointerEvents() (pointerEvents string) {
	js.Rewrite("$<.pointerEvents")
	return pointerEvents
}

func (*CSSStyleDeclaration) SetPointerEvents(pointerEvents string) {
	js.Rewrite("$<.pointerEvents = $1", pointerEvents)
}

func (*CSSStyleDeclaration) GetPosition() (position string) {
	js.Rewrite("$<.position")
	return position
}

func (*CSSStyleDeclaration) SetPosition(position string) {
	js.Rewrite("$<.position = $1", position)
}

func (*CSSStyleDeclaration) GetQuotes() (quotes string) {
	js.Rewrite("$<.quotes")
	return quotes
}

func (*CSSStyleDeclaration) SetQuotes(quotes string) {
	js.Rewrite("$<.quotes = $1", quotes)
}

func (*CSSStyleDeclaration) GetRight() (right string) {
	js.Rewrite("$<.right")
	return right
}

func (*CSSStyleDeclaration) SetRight(right string) {
	js.Rewrite("$<.right = $1", right)
}

func (*CSSStyleDeclaration) GetRotate() (rotate string) {
	js.Rewrite("$<.rotate")
	return rotate
}

func (*CSSStyleDeclaration) SetRotate(rotate string) {
	js.Rewrite("$<.rotate = $1", rotate)
}

func (*CSSStyleDeclaration) GetRubyAlign() (rubyAlign string) {
	js.Rewrite("$<.rubyAlign")
	return rubyAlign
}

func (*CSSStyleDeclaration) SetRubyAlign(rubyAlign string) {
	js.Rewrite("$<.rubyAlign = $1", rubyAlign)
}

func (*CSSStyleDeclaration) GetRubyOverhang() (rubyOverhang string) {
	js.Rewrite("$<.rubyOverhang")
	return rubyOverhang
}

func (*CSSStyleDeclaration) SetRubyOverhang(rubyOverhang string) {
	js.Rewrite("$<.rubyOverhang = $1", rubyOverhang)
}

func (*CSSStyleDeclaration) GetRubyPosition() (rubyPosition string) {
	js.Rewrite("$<.rubyPosition")
	return rubyPosition
}

func (*CSSStyleDeclaration) SetRubyPosition(rubyPosition string) {
	js.Rewrite("$<.rubyPosition = $1", rubyPosition)
}

func (*CSSStyleDeclaration) GetScale() (scale string) {
	js.Rewrite("$<.scale")
	return scale
}

func (*CSSStyleDeclaration) SetScale(scale string) {
	js.Rewrite("$<.scale = $1", scale)
}

func (*CSSStyleDeclaration) GetStopColor() (stopColor string) {
	js.Rewrite("$<.stopColor")
	return stopColor
}

func (*CSSStyleDeclaration) SetStopColor(stopColor string) {
	js.Rewrite("$<.stopColor = $1", stopColor)
}

func (*CSSStyleDeclaration) GetStopOpacity() (stopOpacity string) {
	js.Rewrite("$<.stopOpacity")
	return stopOpacity
}

func (*CSSStyleDeclaration) SetStopOpacity(stopOpacity string) {
	js.Rewrite("$<.stopOpacity = $1", stopOpacity)
}

func (*CSSStyleDeclaration) GetStroke() (stroke string) {
	js.Rewrite("$<.stroke")
	return stroke
}

func (*CSSStyleDeclaration) SetStroke(stroke string) {
	js.Rewrite("$<.stroke = $1", stroke)
}

func (*CSSStyleDeclaration) GetStrokeDasharray() (strokeDasharray string) {
	js.Rewrite("$<.strokeDasharray")
	return strokeDasharray
}

func (*CSSStyleDeclaration) SetStrokeDasharray(strokeDasharray string) {
	js.Rewrite("$<.strokeDasharray = $1", strokeDasharray)
}

func (*CSSStyleDeclaration) GetStrokeDashoffset() (strokeDashoffset string) {
	js.Rewrite("$<.strokeDashoffset")
	return strokeDashoffset
}

func (*CSSStyleDeclaration) SetStrokeDashoffset(strokeDashoffset string) {
	js.Rewrite("$<.strokeDashoffset = $1", strokeDashoffset)
}

func (*CSSStyleDeclaration) GetStrokeLinecap() (strokeLinecap string) {
	js.Rewrite("$<.strokeLinecap")
	return strokeLinecap
}

func (*CSSStyleDeclaration) SetStrokeLinecap(strokeLinecap string) {
	js.Rewrite("$<.strokeLinecap = $1", strokeLinecap)
}

func (*CSSStyleDeclaration) GetStrokeLinejoin() (strokeLinejoin string) {
	js.Rewrite("$<.strokeLinejoin")
	return strokeLinejoin
}

func (*CSSStyleDeclaration) SetStrokeLinejoin(strokeLinejoin string) {
	js.Rewrite("$<.strokeLinejoin = $1", strokeLinejoin)
}

func (*CSSStyleDeclaration) GetStrokeMiterlimit() (strokeMiterlimit string) {
	js.Rewrite("$<.strokeMiterlimit")
	return strokeMiterlimit
}

func (*CSSStyleDeclaration) SetStrokeMiterlimit(strokeMiterlimit string) {
	js.Rewrite("$<.strokeMiterlimit = $1", strokeMiterlimit)
}

func (*CSSStyleDeclaration) GetStrokeOpacity() (strokeOpacity string) {
	js.Rewrite("$<.strokeOpacity")
	return strokeOpacity
}

func (*CSSStyleDeclaration) SetStrokeOpacity(strokeOpacity string) {
	js.Rewrite("$<.strokeOpacity = $1", strokeOpacity)
}

func (*CSSStyleDeclaration) GetStrokeWidth() (strokeWidth string) {
	js.Rewrite("$<.strokeWidth")
	return strokeWidth
}

func (*CSSStyleDeclaration) SetStrokeWidth(strokeWidth string) {
	js.Rewrite("$<.strokeWidth = $1", strokeWidth)
}

func (*CSSStyleDeclaration) GetTableLayout() (tableLayout string) {
	js.Rewrite("$<.tableLayout")
	return tableLayout
}

func (*CSSStyleDeclaration) SetTableLayout(tableLayout string) {
	js.Rewrite("$<.tableLayout = $1", tableLayout)
}

func (*CSSStyleDeclaration) GetTextAlign() (textAlign string) {
	js.Rewrite("$<.textAlign")
	return textAlign
}

func (*CSSStyleDeclaration) SetTextAlign(textAlign string) {
	js.Rewrite("$<.textAlign = $1", textAlign)
}

func (*CSSStyleDeclaration) GetTextAlignLast() (textAlignLast string) {
	js.Rewrite("$<.textAlignLast")
	return textAlignLast
}

func (*CSSStyleDeclaration) SetTextAlignLast(textAlignLast string) {
	js.Rewrite("$<.textAlignLast = $1", textAlignLast)
}

func (*CSSStyleDeclaration) GetTextAnchor() (textAnchor string) {
	js.Rewrite("$<.textAnchor")
	return textAnchor
}

func (*CSSStyleDeclaration) SetTextAnchor(textAnchor string) {
	js.Rewrite("$<.textAnchor = $1", textAnchor)
}

func (*CSSStyleDeclaration) GetTextDecoration() (textDecoration string) {
	js.Rewrite("$<.textDecoration")
	return textDecoration
}

func (*CSSStyleDeclaration) SetTextDecoration(textDecoration string) {
	js.Rewrite("$<.textDecoration = $1", textDecoration)
}

func (*CSSStyleDeclaration) GetTextIndent() (textIndent string) {
	js.Rewrite("$<.textIndent")
	return textIndent
}

func (*CSSStyleDeclaration) SetTextIndent(textIndent string) {
	js.Rewrite("$<.textIndent = $1", textIndent)
}

func (*CSSStyleDeclaration) GetTextJustify() (textJustify string) {
	js.Rewrite("$<.textJustify")
	return textJustify
}

func (*CSSStyleDeclaration) SetTextJustify(textJustify string) {
	js.Rewrite("$<.textJustify = $1", textJustify)
}

func (*CSSStyleDeclaration) GetTextKashida() (textKashida string) {
	js.Rewrite("$<.textKashida")
	return textKashida
}

func (*CSSStyleDeclaration) SetTextKashida(textKashida string) {
	js.Rewrite("$<.textKashida = $1", textKashida)
}

func (*CSSStyleDeclaration) GetTextKashidaSpace() (textKashidaSpace string) {
	js.Rewrite("$<.textKashidaSpace")
	return textKashidaSpace
}

func (*CSSStyleDeclaration) SetTextKashidaSpace(textKashidaSpace string) {
	js.Rewrite("$<.textKashidaSpace = $1", textKashidaSpace)
}

func (*CSSStyleDeclaration) GetTextOverflow() (textOverflow string) {
	js.Rewrite("$<.textOverflow")
	return textOverflow
}

func (*CSSStyleDeclaration) SetTextOverflow(textOverflow string) {
	js.Rewrite("$<.textOverflow = $1", textOverflow)
}

func (*CSSStyleDeclaration) GetTextShadow() (textShadow string) {
	js.Rewrite("$<.textShadow")
	return textShadow
}

func (*CSSStyleDeclaration) SetTextShadow(textShadow string) {
	js.Rewrite("$<.textShadow = $1", textShadow)
}

func (*CSSStyleDeclaration) GetTextTransform() (textTransform string) {
	js.Rewrite("$<.textTransform")
	return textTransform
}

func (*CSSStyleDeclaration) SetTextTransform(textTransform string) {
	js.Rewrite("$<.textTransform = $1", textTransform)
}

func (*CSSStyleDeclaration) GetTextUnderlinePosition() (textUnderlinePosition string) {
	js.Rewrite("$<.textUnderlinePosition")
	return textUnderlinePosition
}

func (*CSSStyleDeclaration) SetTextUnderlinePosition(textUnderlinePosition string) {
	js.Rewrite("$<.textUnderlinePosition = $1", textUnderlinePosition)
}

func (*CSSStyleDeclaration) GetTop() (top string) {
	js.Rewrite("$<.top")
	return top
}

func (*CSSStyleDeclaration) SetTop(top string) {
	js.Rewrite("$<.top = $1", top)
}

func (*CSSStyleDeclaration) GetTouchAction() (touchAction string) {
	js.Rewrite("$<.touchAction")
	return touchAction
}

func (*CSSStyleDeclaration) SetTouchAction(touchAction string) {
	js.Rewrite("$<.touchAction = $1", touchAction)
}

func (*CSSStyleDeclaration) GetTransform() (transform string) {
	js.Rewrite("$<.transform")
	return transform
}

func (*CSSStyleDeclaration) SetTransform(transform string) {
	js.Rewrite("$<.transform = $1", transform)
}

func (*CSSStyleDeclaration) GetTransformOrigin() (transformOrigin string) {
	js.Rewrite("$<.transformOrigin")
	return transformOrigin
}

func (*CSSStyleDeclaration) SetTransformOrigin(transformOrigin string) {
	js.Rewrite("$<.transformOrigin = $1", transformOrigin)
}

func (*CSSStyleDeclaration) GetTransformStyle() (transformStyle string) {
	js.Rewrite("$<.transformStyle")
	return transformStyle
}

func (*CSSStyleDeclaration) SetTransformStyle(transformStyle string) {
	js.Rewrite("$<.transformStyle = $1", transformStyle)
}

func (*CSSStyleDeclaration) GetTransition() (transition string) {
	js.Rewrite("$<.transition")
	return transition
}

func (*CSSStyleDeclaration) SetTransition(transition string) {
	js.Rewrite("$<.transition = $1", transition)
}

func (*CSSStyleDeclaration) GetTransitionDelay() (transitionDelay string) {
	js.Rewrite("$<.transitionDelay")
	return transitionDelay
}

func (*CSSStyleDeclaration) SetTransitionDelay(transitionDelay string) {
	js.Rewrite("$<.transitionDelay = $1", transitionDelay)
}

func (*CSSStyleDeclaration) GetTransitionDuration() (transitionDuration string) {
	js.Rewrite("$<.transitionDuration")
	return transitionDuration
}

func (*CSSStyleDeclaration) SetTransitionDuration(transitionDuration string) {
	js.Rewrite("$<.transitionDuration = $1", transitionDuration)
}

func (*CSSStyleDeclaration) GetTransitionProperty() (transitionProperty string) {
	js.Rewrite("$<.transitionProperty")
	return transitionProperty
}

func (*CSSStyleDeclaration) SetTransitionProperty(transitionProperty string) {
	js.Rewrite("$<.transitionProperty = $1", transitionProperty)
}

func (*CSSStyleDeclaration) GetTransitionTimingFunction() (transitionTimingFunction string) {
	js.Rewrite("$<.transitionTimingFunction")
	return transitionTimingFunction
}

func (*CSSStyleDeclaration) SetTransitionTimingFunction(transitionTimingFunction string) {
	js.Rewrite("$<.transitionTimingFunction = $1", transitionTimingFunction)
}

func (*CSSStyleDeclaration) GetTranslate() (translate string) {
	js.Rewrite("$<.translate")
	return translate
}

func (*CSSStyleDeclaration) SetTranslate(translate string) {
	js.Rewrite("$<.translate = $1", translate)
}

func (*CSSStyleDeclaration) GetUnicodeBidi() (unicodeBidi string) {
	js.Rewrite("$<.unicodeBidi")
	return unicodeBidi
}

func (*CSSStyleDeclaration) SetUnicodeBidi(unicodeBidi string) {
	js.Rewrite("$<.unicodeBidi = $1", unicodeBidi)
}

func (*CSSStyleDeclaration) GetVerticalAlign() (verticalAlign string) {
	js.Rewrite("$<.verticalAlign")
	return verticalAlign
}

func (*CSSStyleDeclaration) SetVerticalAlign(verticalAlign string) {
	js.Rewrite("$<.verticalAlign = $1", verticalAlign)
}

func (*CSSStyleDeclaration) GetVisibility() (visibility string) {
	js.Rewrite("$<.visibility")
	return visibility
}

func (*CSSStyleDeclaration) SetVisibility(visibility string) {
	js.Rewrite("$<.visibility = $1", visibility)
}

func (*CSSStyleDeclaration) GetWebkitAlignContent() (webkitAlignContent string) {
	js.Rewrite("$<.webkitAlignContent")
	return webkitAlignContent
}

func (*CSSStyleDeclaration) SetWebkitAlignContent(webkitAlignContent string) {
	js.Rewrite("$<.webkitAlignContent = $1", webkitAlignContent)
}

func (*CSSStyleDeclaration) GetWebkitAlignItems() (webkitAlignItems string) {
	js.Rewrite("$<.webkitAlignItems")
	return webkitAlignItems
}

func (*CSSStyleDeclaration) SetWebkitAlignItems(webkitAlignItems string) {
	js.Rewrite("$<.webkitAlignItems = $1", webkitAlignItems)
}

func (*CSSStyleDeclaration) GetWebkitAlignSelf() (webkitAlignSelf string) {
	js.Rewrite("$<.webkitAlignSelf")
	return webkitAlignSelf
}

func (*CSSStyleDeclaration) SetWebkitAlignSelf(webkitAlignSelf string) {
	js.Rewrite("$<.webkitAlignSelf = $1", webkitAlignSelf)
}

func (*CSSStyleDeclaration) GetWebkitAnimation() (webkitAnimation string) {
	js.Rewrite("$<.webkitAnimation")
	return webkitAnimation
}

func (*CSSStyleDeclaration) SetWebkitAnimation(webkitAnimation string) {
	js.Rewrite("$<.webkitAnimation = $1", webkitAnimation)
}

func (*CSSStyleDeclaration) GetWebkitAnimationDelay() (webkitAnimationDelay string) {
	js.Rewrite("$<.webkitAnimationDelay")
	return webkitAnimationDelay
}

func (*CSSStyleDeclaration) SetWebkitAnimationDelay(webkitAnimationDelay string) {
	js.Rewrite("$<.webkitAnimationDelay = $1", webkitAnimationDelay)
}

func (*CSSStyleDeclaration) GetWebkitAnimationDirection() (webkitAnimationDirection string) {
	js.Rewrite("$<.webkitAnimationDirection")
	return webkitAnimationDirection
}

func (*CSSStyleDeclaration) SetWebkitAnimationDirection(webkitAnimationDirection string) {
	js.Rewrite("$<.webkitAnimationDirection = $1", webkitAnimationDirection)
}

func (*CSSStyleDeclaration) GetWebkitAnimationDuration() (webkitAnimationDuration string) {
	js.Rewrite("$<.webkitAnimationDuration")
	return webkitAnimationDuration
}

func (*CSSStyleDeclaration) SetWebkitAnimationDuration(webkitAnimationDuration string) {
	js.Rewrite("$<.webkitAnimationDuration = $1", webkitAnimationDuration)
}

func (*CSSStyleDeclaration) GetWebkitAnimationFillMode() (webkitAnimationFillMode string) {
	js.Rewrite("$<.webkitAnimationFillMode")
	return webkitAnimationFillMode
}

func (*CSSStyleDeclaration) SetWebkitAnimationFillMode(webkitAnimationFillMode string) {
	js.Rewrite("$<.webkitAnimationFillMode = $1", webkitAnimationFillMode)
}

func (*CSSStyleDeclaration) GetWebkitAnimationIterationCount() (webkitAnimationIterationCount string) {
	js.Rewrite("$<.webkitAnimationIterationCount")
	return webkitAnimationIterationCount
}

func (*CSSStyleDeclaration) SetWebkitAnimationIterationCount(webkitAnimationIterationCount string) {
	js.Rewrite("$<.webkitAnimationIterationCount = $1", webkitAnimationIterationCount)
}

func (*CSSStyleDeclaration) GetWebkitAnimationName() (webkitAnimationName string) {
	js.Rewrite("$<.webkitAnimationName")
	return webkitAnimationName
}

func (*CSSStyleDeclaration) SetWebkitAnimationName(webkitAnimationName string) {
	js.Rewrite("$<.webkitAnimationName = $1", webkitAnimationName)
}

func (*CSSStyleDeclaration) GetWebkitAnimationPlayState() (webkitAnimationPlayState string) {
	js.Rewrite("$<.webkitAnimationPlayState")
	return webkitAnimationPlayState
}

func (*CSSStyleDeclaration) SetWebkitAnimationPlayState(webkitAnimationPlayState string) {
	js.Rewrite("$<.webkitAnimationPlayState = $1", webkitAnimationPlayState)
}

func (*CSSStyleDeclaration) GetWebkitAnimationTimingFunction() (webkitAnimationTimingFunction string) {
	js.Rewrite("$<.webkitAnimationTimingFunction")
	return webkitAnimationTimingFunction
}

func (*CSSStyleDeclaration) SetWebkitAnimationTimingFunction(webkitAnimationTimingFunction string) {
	js.Rewrite("$<.webkitAnimationTimingFunction = $1", webkitAnimationTimingFunction)
}

func (*CSSStyleDeclaration) GetWebkitAppearance() (webkitAppearance string) {
	js.Rewrite("$<.webkitAppearance")
	return webkitAppearance
}

func (*CSSStyleDeclaration) SetWebkitAppearance(webkitAppearance string) {
	js.Rewrite("$<.webkitAppearance = $1", webkitAppearance)
}

func (*CSSStyleDeclaration) GetWebkitBackfaceVisibility() (webkitBackfaceVisibility string) {
	js.Rewrite("$<.webkitBackfaceVisibility")
	return webkitBackfaceVisibility
}

func (*CSSStyleDeclaration) SetWebkitBackfaceVisibility(webkitBackfaceVisibility string) {
	js.Rewrite("$<.webkitBackfaceVisibility = $1", webkitBackfaceVisibility)
}

func (*CSSStyleDeclaration) GetWebkitBackgroundClip() (webkitBackgroundClip string) {
	js.Rewrite("$<.webkitBackgroundClip")
	return webkitBackgroundClip
}

func (*CSSStyleDeclaration) SetWebkitBackgroundClip(webkitBackgroundClip string) {
	js.Rewrite("$<.webkitBackgroundClip = $1", webkitBackgroundClip)
}

func (*CSSStyleDeclaration) GetWebkitBackgroundOrigin() (webkitBackgroundOrigin string) {
	js.Rewrite("$<.webkitBackgroundOrigin")
	return webkitBackgroundOrigin
}

func (*CSSStyleDeclaration) SetWebkitBackgroundOrigin(webkitBackgroundOrigin string) {
	js.Rewrite("$<.webkitBackgroundOrigin = $1", webkitBackgroundOrigin)
}

func (*CSSStyleDeclaration) GetWebkitBackgroundSize() (webkitBackgroundSize string) {
	js.Rewrite("$<.webkitBackgroundSize")
	return webkitBackgroundSize
}

func (*CSSStyleDeclaration) SetWebkitBackgroundSize(webkitBackgroundSize string) {
	js.Rewrite("$<.webkitBackgroundSize = $1", webkitBackgroundSize)
}

func (*CSSStyleDeclaration) GetWebkitBorderBottomLeftRadius() (webkitBorderBottomLeftRadius string) {
	js.Rewrite("$<.webkitBorderBottomLeftRadius")
	return webkitBorderBottomLeftRadius
}

func (*CSSStyleDeclaration) SetWebkitBorderBottomLeftRadius(webkitBorderBottomLeftRadius string) {
	js.Rewrite("$<.webkitBorderBottomLeftRadius = $1", webkitBorderBottomLeftRadius)
}

func (*CSSStyleDeclaration) GetWebkitBorderBottomRightRadius() (webkitBorderBottomRightRadius string) {
	js.Rewrite("$<.webkitBorderBottomRightRadius")
	return webkitBorderBottomRightRadius
}

func (*CSSStyleDeclaration) SetWebkitBorderBottomRightRadius(webkitBorderBottomRightRadius string) {
	js.Rewrite("$<.webkitBorderBottomRightRadius = $1", webkitBorderBottomRightRadius)
}

func (*CSSStyleDeclaration) GetWebkitBorderImage() (webkitBorderImage string) {
	js.Rewrite("$<.webkitBorderImage")
	return webkitBorderImage
}

func (*CSSStyleDeclaration) SetWebkitBorderImage(webkitBorderImage string) {
	js.Rewrite("$<.webkitBorderImage = $1", webkitBorderImage)
}

func (*CSSStyleDeclaration) GetWebkitBorderRadius() (webkitBorderRadius string) {
	js.Rewrite("$<.webkitBorderRadius")
	return webkitBorderRadius
}

func (*CSSStyleDeclaration) SetWebkitBorderRadius(webkitBorderRadius string) {
	js.Rewrite("$<.webkitBorderRadius = $1", webkitBorderRadius)
}

func (*CSSStyleDeclaration) GetWebkitBorderTopLeftRadius() (webkitBorderTopLeftRadius string) {
	js.Rewrite("$<.webkitBorderTopLeftRadius")
	return webkitBorderTopLeftRadius
}

func (*CSSStyleDeclaration) SetWebkitBorderTopLeftRadius(webkitBorderTopLeftRadius string) {
	js.Rewrite("$<.webkitBorderTopLeftRadius = $1", webkitBorderTopLeftRadius)
}

func (*CSSStyleDeclaration) GetWebkitBorderTopRightRadius() (webkitBorderTopRightRadius string) {
	js.Rewrite("$<.webkitBorderTopRightRadius")
	return webkitBorderTopRightRadius
}

func (*CSSStyleDeclaration) SetWebkitBorderTopRightRadius(webkitBorderTopRightRadius string) {
	js.Rewrite("$<.webkitBorderTopRightRadius = $1", webkitBorderTopRightRadius)
}

func (*CSSStyleDeclaration) GetWebkitBoxAlign() (webkitBoxAlign string) {
	js.Rewrite("$<.webkitBoxAlign")
	return webkitBoxAlign
}

func (*CSSStyleDeclaration) SetWebkitBoxAlign(webkitBoxAlign string) {
	js.Rewrite("$<.webkitBoxAlign = $1", webkitBoxAlign)
}

func (*CSSStyleDeclaration) GetWebkitBoxDirection() (webkitBoxDirection string) {
	js.Rewrite("$<.webkitBoxDirection")
	return webkitBoxDirection
}

func (*CSSStyleDeclaration) SetWebkitBoxDirection(webkitBoxDirection string) {
	js.Rewrite("$<.webkitBoxDirection = $1", webkitBoxDirection)
}

func (*CSSStyleDeclaration) GetWebkitBoxFlex() (webkitBoxFlex string) {
	js.Rewrite("$<.webkitBoxFlex")
	return webkitBoxFlex
}

func (*CSSStyleDeclaration) SetWebkitBoxFlex(webkitBoxFlex string) {
	js.Rewrite("$<.webkitBoxFlex = $1", webkitBoxFlex)
}

func (*CSSStyleDeclaration) GetWebkitBoxOrdinalGroup() (webkitBoxOrdinalGroup string) {
	js.Rewrite("$<.webkitBoxOrdinalGroup")
	return webkitBoxOrdinalGroup
}

func (*CSSStyleDeclaration) SetWebkitBoxOrdinalGroup(webkitBoxOrdinalGroup string) {
	js.Rewrite("$<.webkitBoxOrdinalGroup = $1", webkitBoxOrdinalGroup)
}

func (*CSSStyleDeclaration) GetWebkitBoxOrient() (webkitBoxOrient string) {
	js.Rewrite("$<.webkitBoxOrient")
	return webkitBoxOrient
}

func (*CSSStyleDeclaration) SetWebkitBoxOrient(webkitBoxOrient string) {
	js.Rewrite("$<.webkitBoxOrient = $1", webkitBoxOrient)
}

func (*CSSStyleDeclaration) GetWebkitBoxPack() (webkitBoxPack string) {
	js.Rewrite("$<.webkitBoxPack")
	return webkitBoxPack
}

func (*CSSStyleDeclaration) SetWebkitBoxPack(webkitBoxPack string) {
	js.Rewrite("$<.webkitBoxPack = $1", webkitBoxPack)
}

func (*CSSStyleDeclaration) GetWebkitBoxSizing() (webkitBoxSizing string) {
	js.Rewrite("$<.webkitBoxSizing")
	return webkitBoxSizing
}

func (*CSSStyleDeclaration) SetWebkitBoxSizing(webkitBoxSizing string) {
	js.Rewrite("$<.webkitBoxSizing = $1", webkitBoxSizing)
}

func (*CSSStyleDeclaration) GetWebkitColumnBreakAfter() (webkitColumnBreakAfter string) {
	js.Rewrite("$<.webkitColumnBreakAfter")
	return webkitColumnBreakAfter
}

func (*CSSStyleDeclaration) SetWebkitColumnBreakAfter(webkitColumnBreakAfter string) {
	js.Rewrite("$<.webkitColumnBreakAfter = $1", webkitColumnBreakAfter)
}

func (*CSSStyleDeclaration) GetWebkitColumnBreakBefore() (webkitColumnBreakBefore string) {
	js.Rewrite("$<.webkitColumnBreakBefore")
	return webkitColumnBreakBefore
}

func (*CSSStyleDeclaration) SetWebkitColumnBreakBefore(webkitColumnBreakBefore string) {
	js.Rewrite("$<.webkitColumnBreakBefore = $1", webkitColumnBreakBefore)
}

func (*CSSStyleDeclaration) GetWebkitColumnBreakInside() (webkitColumnBreakInside string) {
	js.Rewrite("$<.webkitColumnBreakInside")
	return webkitColumnBreakInside
}

func (*CSSStyleDeclaration) SetWebkitColumnBreakInside(webkitColumnBreakInside string) {
	js.Rewrite("$<.webkitColumnBreakInside = $1", webkitColumnBreakInside)
}

func (*CSSStyleDeclaration) GetWebkitColumnCount() (webkitColumnCount interface{}) {
	js.Rewrite("$<.webkitColumnCount")
	return webkitColumnCount
}

func (*CSSStyleDeclaration) SetWebkitColumnCount(webkitColumnCount interface{}) {
	js.Rewrite("$<.webkitColumnCount = $1", webkitColumnCount)
}

func (*CSSStyleDeclaration) GetWebkitColumnGap() (webkitColumnGap interface{}) {
	js.Rewrite("$<.webkitColumnGap")
	return webkitColumnGap
}

func (*CSSStyleDeclaration) SetWebkitColumnGap(webkitColumnGap interface{}) {
	js.Rewrite("$<.webkitColumnGap = $1", webkitColumnGap)
}

func (*CSSStyleDeclaration) GetWebkitColumnRule() (webkitColumnRule string) {
	js.Rewrite("$<.webkitColumnRule")
	return webkitColumnRule
}

func (*CSSStyleDeclaration) SetWebkitColumnRule(webkitColumnRule string) {
	js.Rewrite("$<.webkitColumnRule = $1", webkitColumnRule)
}

func (*CSSStyleDeclaration) GetWebkitColumnRuleColor() (webkitColumnRuleColor interface{}) {
	js.Rewrite("$<.webkitColumnRuleColor")
	return webkitColumnRuleColor
}

func (*CSSStyleDeclaration) SetWebkitColumnRuleColor(webkitColumnRuleColor interface{}) {
	js.Rewrite("$<.webkitColumnRuleColor = $1", webkitColumnRuleColor)
}

func (*CSSStyleDeclaration) GetWebkitColumnRuleStyle() (webkitColumnRuleStyle string) {
	js.Rewrite("$<.webkitColumnRuleStyle")
	return webkitColumnRuleStyle
}

func (*CSSStyleDeclaration) SetWebkitColumnRuleStyle(webkitColumnRuleStyle string) {
	js.Rewrite("$<.webkitColumnRuleStyle = $1", webkitColumnRuleStyle)
}

func (*CSSStyleDeclaration) GetWebkitColumnRuleWidth() (webkitColumnRuleWidth interface{}) {
	js.Rewrite("$<.webkitColumnRuleWidth")
	return webkitColumnRuleWidth
}

func (*CSSStyleDeclaration) SetWebkitColumnRuleWidth(webkitColumnRuleWidth interface{}) {
	js.Rewrite("$<.webkitColumnRuleWidth = $1", webkitColumnRuleWidth)
}

func (*CSSStyleDeclaration) GetWebkitColumns() (webkitColumns string) {
	js.Rewrite("$<.webkitColumns")
	return webkitColumns
}

func (*CSSStyleDeclaration) SetWebkitColumns(webkitColumns string) {
	js.Rewrite("$<.webkitColumns = $1", webkitColumns)
}

func (*CSSStyleDeclaration) GetWebkitColumnSpan() (webkitColumnSpan string) {
	js.Rewrite("$<.webkitColumnSpan")
	return webkitColumnSpan
}

func (*CSSStyleDeclaration) SetWebkitColumnSpan(webkitColumnSpan string) {
	js.Rewrite("$<.webkitColumnSpan = $1", webkitColumnSpan)
}

func (*CSSStyleDeclaration) GetWebkitColumnWidth() (webkitColumnWidth interface{}) {
	js.Rewrite("$<.webkitColumnWidth")
	return webkitColumnWidth
}

func (*CSSStyleDeclaration) SetWebkitColumnWidth(webkitColumnWidth interface{}) {
	js.Rewrite("$<.webkitColumnWidth = $1", webkitColumnWidth)
}

func (*CSSStyleDeclaration) GetWebkitFilter() (webkitFilter string) {
	js.Rewrite("$<.webkitFilter")
	return webkitFilter
}

func (*CSSStyleDeclaration) SetWebkitFilter(webkitFilter string) {
	js.Rewrite("$<.webkitFilter = $1", webkitFilter)
}

func (*CSSStyleDeclaration) GetWebkitFlex() (webkitFlex string) {
	js.Rewrite("$<.webkitFlex")
	return webkitFlex
}

func (*CSSStyleDeclaration) SetWebkitFlex(webkitFlex string) {
	js.Rewrite("$<.webkitFlex = $1", webkitFlex)
}

func (*CSSStyleDeclaration) GetWebkitFlexBasis() (webkitFlexBasis string) {
	js.Rewrite("$<.webkitFlexBasis")
	return webkitFlexBasis
}

func (*CSSStyleDeclaration) SetWebkitFlexBasis(webkitFlexBasis string) {
	js.Rewrite("$<.webkitFlexBasis = $1", webkitFlexBasis)
}

func (*CSSStyleDeclaration) GetWebkitFlexDirection() (webkitFlexDirection string) {
	js.Rewrite("$<.webkitFlexDirection")
	return webkitFlexDirection
}

func (*CSSStyleDeclaration) SetWebkitFlexDirection(webkitFlexDirection string) {
	js.Rewrite("$<.webkitFlexDirection = $1", webkitFlexDirection)
}

func (*CSSStyleDeclaration) GetWebkitFlexFlow() (webkitFlexFlow string) {
	js.Rewrite("$<.webkitFlexFlow")
	return webkitFlexFlow
}

func (*CSSStyleDeclaration) SetWebkitFlexFlow(webkitFlexFlow string) {
	js.Rewrite("$<.webkitFlexFlow = $1", webkitFlexFlow)
}

func (*CSSStyleDeclaration) GetWebkitFlexGrow() (webkitFlexGrow string) {
	js.Rewrite("$<.webkitFlexGrow")
	return webkitFlexGrow
}

func (*CSSStyleDeclaration) SetWebkitFlexGrow(webkitFlexGrow string) {
	js.Rewrite("$<.webkitFlexGrow = $1", webkitFlexGrow)
}

func (*CSSStyleDeclaration) GetWebkitFlexShrink() (webkitFlexShrink string) {
	js.Rewrite("$<.webkitFlexShrink")
	return webkitFlexShrink
}

func (*CSSStyleDeclaration) SetWebkitFlexShrink(webkitFlexShrink string) {
	js.Rewrite("$<.webkitFlexShrink = $1", webkitFlexShrink)
}

func (*CSSStyleDeclaration) GetWebkitFlexWrap() (webkitFlexWrap string) {
	js.Rewrite("$<.webkitFlexWrap")
	return webkitFlexWrap
}

func (*CSSStyleDeclaration) SetWebkitFlexWrap(webkitFlexWrap string) {
	js.Rewrite("$<.webkitFlexWrap = $1", webkitFlexWrap)
}

func (*CSSStyleDeclaration) GetWebkitJustifyContent() (webkitJustifyContent string) {
	js.Rewrite("$<.webkitJustifyContent")
	return webkitJustifyContent
}

func (*CSSStyleDeclaration) SetWebkitJustifyContent(webkitJustifyContent string) {
	js.Rewrite("$<.webkitJustifyContent = $1", webkitJustifyContent)
}

func (*CSSStyleDeclaration) GetWebkitOrder() (webkitOrder string) {
	js.Rewrite("$<.webkitOrder")
	return webkitOrder
}

func (*CSSStyleDeclaration) SetWebkitOrder(webkitOrder string) {
	js.Rewrite("$<.webkitOrder = $1", webkitOrder)
}

func (*CSSStyleDeclaration) GetWebkitPerspective() (webkitPerspective string) {
	js.Rewrite("$<.webkitPerspective")
	return webkitPerspective
}

func (*CSSStyleDeclaration) SetWebkitPerspective(webkitPerspective string) {
	js.Rewrite("$<.webkitPerspective = $1", webkitPerspective)
}

func (*CSSStyleDeclaration) GetWebkitPerspectiveOrigin() (webkitPerspectiveOrigin string) {
	js.Rewrite("$<.webkitPerspectiveOrigin")
	return webkitPerspectiveOrigin
}

func (*CSSStyleDeclaration) SetWebkitPerspectiveOrigin(webkitPerspectiveOrigin string) {
	js.Rewrite("$<.webkitPerspectiveOrigin = $1", webkitPerspectiveOrigin)
}

func (*CSSStyleDeclaration) GetWebkitTapHighlightColor() (webkitTapHighlightColor string) {
	js.Rewrite("$<.webkitTapHighlightColor")
	return webkitTapHighlightColor
}

func (*CSSStyleDeclaration) SetWebkitTapHighlightColor(webkitTapHighlightColor string) {
	js.Rewrite("$<.webkitTapHighlightColor = $1", webkitTapHighlightColor)
}

func (*CSSStyleDeclaration) GetWebkitTextFillColor() (webkitTextFillColor string) {
	js.Rewrite("$<.webkitTextFillColor")
	return webkitTextFillColor
}

func (*CSSStyleDeclaration) SetWebkitTextFillColor(webkitTextFillColor string) {
	js.Rewrite("$<.webkitTextFillColor = $1", webkitTextFillColor)
}

func (*CSSStyleDeclaration) GetWebkitTextSizeAdjust() (webkitTextSizeAdjust interface{}) {
	js.Rewrite("$<.webkitTextSizeAdjust")
	return webkitTextSizeAdjust
}

func (*CSSStyleDeclaration) SetWebkitTextSizeAdjust(webkitTextSizeAdjust interface{}) {
	js.Rewrite("$<.webkitTextSizeAdjust = $1", webkitTextSizeAdjust)
}

func (*CSSStyleDeclaration) GetWebkitTextStroke() (webkitTextStroke string) {
	js.Rewrite("$<.webkitTextStroke")
	return webkitTextStroke
}

func (*CSSStyleDeclaration) SetWebkitTextStroke(webkitTextStroke string) {
	js.Rewrite("$<.webkitTextStroke = $1", webkitTextStroke)
}

func (*CSSStyleDeclaration) GetWebkitTextStrokeColor() (webkitTextStrokeColor string) {
	js.Rewrite("$<.webkitTextStrokeColor")
	return webkitTextStrokeColor
}

func (*CSSStyleDeclaration) SetWebkitTextStrokeColor(webkitTextStrokeColor string) {
	js.Rewrite("$<.webkitTextStrokeColor = $1", webkitTextStrokeColor)
}

func (*CSSStyleDeclaration) GetWebkitTextStrokeWidth() (webkitTextStrokeWidth string) {
	js.Rewrite("$<.webkitTextStrokeWidth")
	return webkitTextStrokeWidth
}

func (*CSSStyleDeclaration) SetWebkitTextStrokeWidth(webkitTextStrokeWidth string) {
	js.Rewrite("$<.webkitTextStrokeWidth = $1", webkitTextStrokeWidth)
}

func (*CSSStyleDeclaration) GetWebkitTransform() (webkitTransform string) {
	js.Rewrite("$<.webkitTransform")
	return webkitTransform
}

func (*CSSStyleDeclaration) SetWebkitTransform(webkitTransform string) {
	js.Rewrite("$<.webkitTransform = $1", webkitTransform)
}

func (*CSSStyleDeclaration) GetWebkitTransformOrigin() (webkitTransformOrigin string) {
	js.Rewrite("$<.webkitTransformOrigin")
	return webkitTransformOrigin
}

func (*CSSStyleDeclaration) SetWebkitTransformOrigin(webkitTransformOrigin string) {
	js.Rewrite("$<.webkitTransformOrigin = $1", webkitTransformOrigin)
}

func (*CSSStyleDeclaration) GetWebkitTransformStyle() (webkitTransformStyle string) {
	js.Rewrite("$<.webkitTransformStyle")
	return webkitTransformStyle
}

func (*CSSStyleDeclaration) SetWebkitTransformStyle(webkitTransformStyle string) {
	js.Rewrite("$<.webkitTransformStyle = $1", webkitTransformStyle)
}

func (*CSSStyleDeclaration) GetWebkitTransition() (webkitTransition string) {
	js.Rewrite("$<.webkitTransition")
	return webkitTransition
}

func (*CSSStyleDeclaration) SetWebkitTransition(webkitTransition string) {
	js.Rewrite("$<.webkitTransition = $1", webkitTransition)
}

func (*CSSStyleDeclaration) GetWebkitTransitionDelay() (webkitTransitionDelay string) {
	js.Rewrite("$<.webkitTransitionDelay")
	return webkitTransitionDelay
}

func (*CSSStyleDeclaration) SetWebkitTransitionDelay(webkitTransitionDelay string) {
	js.Rewrite("$<.webkitTransitionDelay = $1", webkitTransitionDelay)
}

func (*CSSStyleDeclaration) GetWebkitTransitionDuration() (webkitTransitionDuration string) {
	js.Rewrite("$<.webkitTransitionDuration")
	return webkitTransitionDuration
}

func (*CSSStyleDeclaration) SetWebkitTransitionDuration(webkitTransitionDuration string) {
	js.Rewrite("$<.webkitTransitionDuration = $1", webkitTransitionDuration)
}

func (*CSSStyleDeclaration) GetWebkitTransitionProperty() (webkitTransitionProperty string) {
	js.Rewrite("$<.webkitTransitionProperty")
	return webkitTransitionProperty
}

func (*CSSStyleDeclaration) SetWebkitTransitionProperty(webkitTransitionProperty string) {
	js.Rewrite("$<.webkitTransitionProperty = $1", webkitTransitionProperty)
}

func (*CSSStyleDeclaration) GetWebkitTransitionTimingFunction() (webkitTransitionTimingFunction string) {
	js.Rewrite("$<.webkitTransitionTimingFunction")
	return webkitTransitionTimingFunction
}

func (*CSSStyleDeclaration) SetWebkitTransitionTimingFunction(webkitTransitionTimingFunction string) {
	js.Rewrite("$<.webkitTransitionTimingFunction = $1", webkitTransitionTimingFunction)
}

func (*CSSStyleDeclaration) GetWebkitUserModify() (webkitUserModify string) {
	js.Rewrite("$<.webkitUserModify")
	return webkitUserModify
}

func (*CSSStyleDeclaration) SetWebkitUserModify(webkitUserModify string) {
	js.Rewrite("$<.webkitUserModify = $1", webkitUserModify)
}

func (*CSSStyleDeclaration) GetWebkitUserSelect() (webkitUserSelect string) {
	js.Rewrite("$<.webkitUserSelect")
	return webkitUserSelect
}

func (*CSSStyleDeclaration) SetWebkitUserSelect(webkitUserSelect string) {
	js.Rewrite("$<.webkitUserSelect = $1", webkitUserSelect)
}

func (*CSSStyleDeclaration) GetWebkitWritingMode() (webkitWritingMode string) {
	js.Rewrite("$<.webkitWritingMode")
	return webkitWritingMode
}

func (*CSSStyleDeclaration) SetWebkitWritingMode(webkitWritingMode string) {
	js.Rewrite("$<.webkitWritingMode = $1", webkitWritingMode)
}

func (*CSSStyleDeclaration) GetWhiteSpace() (whiteSpace string) {
	js.Rewrite("$<.whiteSpace")
	return whiteSpace
}

func (*CSSStyleDeclaration) SetWhiteSpace(whiteSpace string) {
	js.Rewrite("$<.whiteSpace = $1", whiteSpace)
}

func (*CSSStyleDeclaration) GetWidows() (widows string) {
	js.Rewrite("$<.widows")
	return widows
}

func (*CSSStyleDeclaration) SetWidows(widows string) {
	js.Rewrite("$<.widows = $1", widows)
}

func (*CSSStyleDeclaration) GetWidth() (width string) {
	js.Rewrite("$<.width")
	return width
}

func (*CSSStyleDeclaration) SetWidth(width string) {
	js.Rewrite("$<.width = $1", width)
}

func (*CSSStyleDeclaration) GetWordBreak() (wordBreak string) {
	js.Rewrite("$<.wordBreak")
	return wordBreak
}

func (*CSSStyleDeclaration) SetWordBreak(wordBreak string) {
	js.Rewrite("$<.wordBreak = $1", wordBreak)
}

func (*CSSStyleDeclaration) GetWordSpacing() (wordSpacing string) {
	js.Rewrite("$<.wordSpacing")
	return wordSpacing
}

func (*CSSStyleDeclaration) SetWordSpacing(wordSpacing string) {
	js.Rewrite("$<.wordSpacing = $1", wordSpacing)
}

func (*CSSStyleDeclaration) GetWordWrap() (wordWrap string) {
	js.Rewrite("$<.wordWrap")
	return wordWrap
}

func (*CSSStyleDeclaration) SetWordWrap(wordWrap string) {
	js.Rewrite("$<.wordWrap = $1", wordWrap)
}

func (*CSSStyleDeclaration) GetWritingMode() (writingMode string) {
	js.Rewrite("$<.writingMode")
	return writingMode
}

func (*CSSStyleDeclaration) SetWritingMode(writingMode string) {
	js.Rewrite("$<.writingMode = $1", writingMode)
}

func (*CSSStyleDeclaration) GetZIndex() (zIndex string) {
	js.Rewrite("$<.zIndex")
	return zIndex
}

func (*CSSStyleDeclaration) SetZIndex(zIndex string) {
	js.Rewrite("$<.zIndex = $1", zIndex)
}

func (*CSSStyleDeclaration) GetZoom() (zoom string) {
	js.Rewrite("$<.zoom")
	return zoom
}

func (*CSSStyleDeclaration) SetZoom(zoom string) {
	js.Rewrite("$<.zoom = $1", zoom)
}
