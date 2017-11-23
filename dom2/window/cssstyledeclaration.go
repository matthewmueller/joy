package window

import "github.com/matthewmueller/golly/js"

// js:"CSSStyleDeclaration,omit"
type CSSStyleDeclaration struct {
}

// GetPropertyPriority
func (*CSSStyleDeclaration) GetPropertyPriority(propertyName string) (s string) {
	js.Rewrite("$<.GetPropertyPriority($1)", propertyName)
	return s
}

// GetPropertyValue
func (*CSSStyleDeclaration) GetPropertyValue(propertyName string) (s string) {
	js.Rewrite("$<.GetPropertyValue($1)", propertyName)
	return s
}

// Item
func (*CSSStyleDeclaration) Item(index uint) (s string) {
	js.Rewrite("$<.Item($1)", index)
	return s
}

// RemoveProperty
func (*CSSStyleDeclaration) RemoveProperty(propertyName string) (s string) {
	js.Rewrite("$<.RemoveProperty($1)", propertyName)
	return s
}

// SetProperty
func (*CSSStyleDeclaration) SetProperty(propertyName string, value string, priority *string) {
	js.Rewrite("$<.SetProperty($1, $2, $3)", propertyName, value, priority)
}

// AlignContent
func (*CSSStyleDeclaration) AlignContent() (alignContent string) {
	js.Rewrite("$<.AlignContent")
	return alignContent
}

// AlignContent
func (*CSSStyleDeclaration) SetAlignContent(alignContent string) {
	js.Rewrite("$<.AlignContent = $1", alignContent)
}

// AlignItems
func (*CSSStyleDeclaration) AlignItems() (alignItems string) {
	js.Rewrite("$<.AlignItems")
	return alignItems
}

// AlignItems
func (*CSSStyleDeclaration) SetAlignItems(alignItems string) {
	js.Rewrite("$<.AlignItems = $1", alignItems)
}

// AlignmentBaseline
func (*CSSStyleDeclaration) AlignmentBaseline() (alignmentBaseline string) {
	js.Rewrite("$<.AlignmentBaseline")
	return alignmentBaseline
}

// AlignmentBaseline
func (*CSSStyleDeclaration) SetAlignmentBaseline(alignmentBaseline string) {
	js.Rewrite("$<.AlignmentBaseline = $1", alignmentBaseline)
}

// AlignSelf
func (*CSSStyleDeclaration) AlignSelf() (alignSelf string) {
	js.Rewrite("$<.AlignSelf")
	return alignSelf
}

// AlignSelf
func (*CSSStyleDeclaration) SetAlignSelf(alignSelf string) {
	js.Rewrite("$<.AlignSelf = $1", alignSelf)
}

// Animation
func (*CSSStyleDeclaration) Animation() (animation string) {
	js.Rewrite("$<.Animation")
	return animation
}

// Animation
func (*CSSStyleDeclaration) SetAnimation(animation string) {
	js.Rewrite("$<.Animation = $1", animation)
}

// AnimationDelay
func (*CSSStyleDeclaration) AnimationDelay() (animationDelay string) {
	js.Rewrite("$<.AnimationDelay")
	return animationDelay
}

// AnimationDelay
func (*CSSStyleDeclaration) SetAnimationDelay(animationDelay string) {
	js.Rewrite("$<.AnimationDelay = $1", animationDelay)
}

// AnimationDirection
func (*CSSStyleDeclaration) AnimationDirection() (animationDirection string) {
	js.Rewrite("$<.AnimationDirection")
	return animationDirection
}

// AnimationDirection
func (*CSSStyleDeclaration) SetAnimationDirection(animationDirection string) {
	js.Rewrite("$<.AnimationDirection = $1", animationDirection)
}

// AnimationDuration
func (*CSSStyleDeclaration) AnimationDuration() (animationDuration string) {
	js.Rewrite("$<.AnimationDuration")
	return animationDuration
}

// AnimationDuration
func (*CSSStyleDeclaration) SetAnimationDuration(animationDuration string) {
	js.Rewrite("$<.AnimationDuration = $1", animationDuration)
}

// AnimationFillMode
func (*CSSStyleDeclaration) AnimationFillMode() (animationFillMode string) {
	js.Rewrite("$<.AnimationFillMode")
	return animationFillMode
}

// AnimationFillMode
func (*CSSStyleDeclaration) SetAnimationFillMode(animationFillMode string) {
	js.Rewrite("$<.AnimationFillMode = $1", animationFillMode)
}

// AnimationIterationCount
func (*CSSStyleDeclaration) AnimationIterationCount() (animationIterationCount string) {
	js.Rewrite("$<.AnimationIterationCount")
	return animationIterationCount
}

// AnimationIterationCount
func (*CSSStyleDeclaration) SetAnimationIterationCount(animationIterationCount string) {
	js.Rewrite("$<.AnimationIterationCount = $1", animationIterationCount)
}

// AnimationName
func (*CSSStyleDeclaration) AnimationName() (animationName string) {
	js.Rewrite("$<.AnimationName")
	return animationName
}

// AnimationName
func (*CSSStyleDeclaration) SetAnimationName(animationName string) {
	js.Rewrite("$<.AnimationName = $1", animationName)
}

// AnimationPlayState
func (*CSSStyleDeclaration) AnimationPlayState() (animationPlayState string) {
	js.Rewrite("$<.AnimationPlayState")
	return animationPlayState
}

// AnimationPlayState
func (*CSSStyleDeclaration) SetAnimationPlayState(animationPlayState string) {
	js.Rewrite("$<.AnimationPlayState = $1", animationPlayState)
}

// AnimationTimingFunction
func (*CSSStyleDeclaration) AnimationTimingFunction() (animationTimingFunction string) {
	js.Rewrite("$<.AnimationTimingFunction")
	return animationTimingFunction
}

// AnimationTimingFunction
func (*CSSStyleDeclaration) SetAnimationTimingFunction(animationTimingFunction string) {
	js.Rewrite("$<.AnimationTimingFunction = $1", animationTimingFunction)
}

// BackfaceVisibility
func (*CSSStyleDeclaration) BackfaceVisibility() (backfaceVisibility string) {
	js.Rewrite("$<.BackfaceVisibility")
	return backfaceVisibility
}

// BackfaceVisibility
func (*CSSStyleDeclaration) SetBackfaceVisibility(backfaceVisibility string) {
	js.Rewrite("$<.BackfaceVisibility = $1", backfaceVisibility)
}

// Background
func (*CSSStyleDeclaration) Background() (background string) {
	js.Rewrite("$<.Background")
	return background
}

// Background
func (*CSSStyleDeclaration) SetBackground(background string) {
	js.Rewrite("$<.Background = $1", background)
}

// BackgroundAttachment
func (*CSSStyleDeclaration) BackgroundAttachment() (backgroundAttachment string) {
	js.Rewrite("$<.BackgroundAttachment")
	return backgroundAttachment
}

// BackgroundAttachment
func (*CSSStyleDeclaration) SetBackgroundAttachment(backgroundAttachment string) {
	js.Rewrite("$<.BackgroundAttachment = $1", backgroundAttachment)
}

// BackgroundClip
func (*CSSStyleDeclaration) BackgroundClip() (backgroundClip string) {
	js.Rewrite("$<.BackgroundClip")
	return backgroundClip
}

// BackgroundClip
func (*CSSStyleDeclaration) SetBackgroundClip(backgroundClip string) {
	js.Rewrite("$<.BackgroundClip = $1", backgroundClip)
}

// BackgroundColor
func (*CSSStyleDeclaration) BackgroundColor() (backgroundColor string) {
	js.Rewrite("$<.BackgroundColor")
	return backgroundColor
}

// BackgroundColor
func (*CSSStyleDeclaration) SetBackgroundColor(backgroundColor string) {
	js.Rewrite("$<.BackgroundColor = $1", backgroundColor)
}

// BackgroundImage
func (*CSSStyleDeclaration) BackgroundImage() (backgroundImage string) {
	js.Rewrite("$<.BackgroundImage")
	return backgroundImage
}

// BackgroundImage
func (*CSSStyleDeclaration) SetBackgroundImage(backgroundImage string) {
	js.Rewrite("$<.BackgroundImage = $1", backgroundImage)
}

// BackgroundOrigin
func (*CSSStyleDeclaration) BackgroundOrigin() (backgroundOrigin string) {
	js.Rewrite("$<.BackgroundOrigin")
	return backgroundOrigin
}

// BackgroundOrigin
func (*CSSStyleDeclaration) SetBackgroundOrigin(backgroundOrigin string) {
	js.Rewrite("$<.BackgroundOrigin = $1", backgroundOrigin)
}

// BackgroundPosition
func (*CSSStyleDeclaration) BackgroundPosition() (backgroundPosition string) {
	js.Rewrite("$<.BackgroundPosition")
	return backgroundPosition
}

// BackgroundPosition
func (*CSSStyleDeclaration) SetBackgroundPosition(backgroundPosition string) {
	js.Rewrite("$<.BackgroundPosition = $1", backgroundPosition)
}

// BackgroundPositionX
func (*CSSStyleDeclaration) BackgroundPositionX() (backgroundPositionX string) {
	js.Rewrite("$<.BackgroundPositionX")
	return backgroundPositionX
}

// BackgroundPositionX
func (*CSSStyleDeclaration) SetBackgroundPositionX(backgroundPositionX string) {
	js.Rewrite("$<.BackgroundPositionX = $1", backgroundPositionX)
}

// BackgroundPositionY
func (*CSSStyleDeclaration) BackgroundPositionY() (backgroundPositionY string) {
	js.Rewrite("$<.BackgroundPositionY")
	return backgroundPositionY
}

// BackgroundPositionY
func (*CSSStyleDeclaration) SetBackgroundPositionY(backgroundPositionY string) {
	js.Rewrite("$<.BackgroundPositionY = $1", backgroundPositionY)
}

// BackgroundRepeat
func (*CSSStyleDeclaration) BackgroundRepeat() (backgroundRepeat string) {
	js.Rewrite("$<.BackgroundRepeat")
	return backgroundRepeat
}

// BackgroundRepeat
func (*CSSStyleDeclaration) SetBackgroundRepeat(backgroundRepeat string) {
	js.Rewrite("$<.BackgroundRepeat = $1", backgroundRepeat)
}

// BackgroundSize
func (*CSSStyleDeclaration) BackgroundSize() (backgroundSize string) {
	js.Rewrite("$<.BackgroundSize")
	return backgroundSize
}

// BackgroundSize
func (*CSSStyleDeclaration) SetBackgroundSize(backgroundSize string) {
	js.Rewrite("$<.BackgroundSize = $1", backgroundSize)
}

// BaselineShift
func (*CSSStyleDeclaration) BaselineShift() (baselineShift string) {
	js.Rewrite("$<.BaselineShift")
	return baselineShift
}

// BaselineShift
func (*CSSStyleDeclaration) SetBaselineShift(baselineShift string) {
	js.Rewrite("$<.BaselineShift = $1", baselineShift)
}

// Border
func (*CSSStyleDeclaration) Border() (border string) {
	js.Rewrite("$<.Border")
	return border
}

// Border
func (*CSSStyleDeclaration) SetBorder(border string) {
	js.Rewrite("$<.Border = $1", border)
}

// BorderBottom
func (*CSSStyleDeclaration) BorderBottom() (borderBottom string) {
	js.Rewrite("$<.BorderBottom")
	return borderBottom
}

// BorderBottom
func (*CSSStyleDeclaration) SetBorderBottom(borderBottom string) {
	js.Rewrite("$<.BorderBottom = $1", borderBottom)
}

// BorderBottomColor
func (*CSSStyleDeclaration) BorderBottomColor() (borderBottomColor string) {
	js.Rewrite("$<.BorderBottomColor")
	return borderBottomColor
}

// BorderBottomColor
func (*CSSStyleDeclaration) SetBorderBottomColor(borderBottomColor string) {
	js.Rewrite("$<.BorderBottomColor = $1", borderBottomColor)
}

// BorderBottomLeftRadius
func (*CSSStyleDeclaration) BorderBottomLeftRadius() (borderBottomLeftRadius string) {
	js.Rewrite("$<.BorderBottomLeftRadius")
	return borderBottomLeftRadius
}

// BorderBottomLeftRadius
func (*CSSStyleDeclaration) SetBorderBottomLeftRadius(borderBottomLeftRadius string) {
	js.Rewrite("$<.BorderBottomLeftRadius = $1", borderBottomLeftRadius)
}

// BorderBottomRightRadius
func (*CSSStyleDeclaration) BorderBottomRightRadius() (borderBottomRightRadius string) {
	js.Rewrite("$<.BorderBottomRightRadius")
	return borderBottomRightRadius
}

// BorderBottomRightRadius
func (*CSSStyleDeclaration) SetBorderBottomRightRadius(borderBottomRightRadius string) {
	js.Rewrite("$<.BorderBottomRightRadius = $1", borderBottomRightRadius)
}

// BorderBottomStyle
func (*CSSStyleDeclaration) BorderBottomStyle() (borderBottomStyle string) {
	js.Rewrite("$<.BorderBottomStyle")
	return borderBottomStyle
}

// BorderBottomStyle
func (*CSSStyleDeclaration) SetBorderBottomStyle(borderBottomStyle string) {
	js.Rewrite("$<.BorderBottomStyle = $1", borderBottomStyle)
}

// BorderBottomWidth
func (*CSSStyleDeclaration) BorderBottomWidth() (borderBottomWidth string) {
	js.Rewrite("$<.BorderBottomWidth")
	return borderBottomWidth
}

// BorderBottomWidth
func (*CSSStyleDeclaration) SetBorderBottomWidth(borderBottomWidth string) {
	js.Rewrite("$<.BorderBottomWidth = $1", borderBottomWidth)
}

// BorderCollapse
func (*CSSStyleDeclaration) BorderCollapse() (borderCollapse string) {
	js.Rewrite("$<.BorderCollapse")
	return borderCollapse
}

// BorderCollapse
func (*CSSStyleDeclaration) SetBorderCollapse(borderCollapse string) {
	js.Rewrite("$<.BorderCollapse = $1", borderCollapse)
}

// BorderColor
func (*CSSStyleDeclaration) BorderColor() (borderColor string) {
	js.Rewrite("$<.BorderColor")
	return borderColor
}

// BorderColor
func (*CSSStyleDeclaration) SetBorderColor(borderColor string) {
	js.Rewrite("$<.BorderColor = $1", borderColor)
}

// BorderImage
func (*CSSStyleDeclaration) BorderImage() (borderImage string) {
	js.Rewrite("$<.BorderImage")
	return borderImage
}

// BorderImage
func (*CSSStyleDeclaration) SetBorderImage(borderImage string) {
	js.Rewrite("$<.BorderImage = $1", borderImage)
}

// BorderImageOutset
func (*CSSStyleDeclaration) BorderImageOutset() (borderImageOutset string) {
	js.Rewrite("$<.BorderImageOutset")
	return borderImageOutset
}

// BorderImageOutset
func (*CSSStyleDeclaration) SetBorderImageOutset(borderImageOutset string) {
	js.Rewrite("$<.BorderImageOutset = $1", borderImageOutset)
}

// BorderImageRepeat
func (*CSSStyleDeclaration) BorderImageRepeat() (borderImageRepeat string) {
	js.Rewrite("$<.BorderImageRepeat")
	return borderImageRepeat
}

// BorderImageRepeat
func (*CSSStyleDeclaration) SetBorderImageRepeat(borderImageRepeat string) {
	js.Rewrite("$<.BorderImageRepeat = $1", borderImageRepeat)
}

// BorderImageSlice
func (*CSSStyleDeclaration) BorderImageSlice() (borderImageSlice string) {
	js.Rewrite("$<.BorderImageSlice")
	return borderImageSlice
}

// BorderImageSlice
func (*CSSStyleDeclaration) SetBorderImageSlice(borderImageSlice string) {
	js.Rewrite("$<.BorderImageSlice = $1", borderImageSlice)
}

// BorderImageSource
func (*CSSStyleDeclaration) BorderImageSource() (borderImageSource string) {
	js.Rewrite("$<.BorderImageSource")
	return borderImageSource
}

// BorderImageSource
func (*CSSStyleDeclaration) SetBorderImageSource(borderImageSource string) {
	js.Rewrite("$<.BorderImageSource = $1", borderImageSource)
}

// BorderImageWidth
func (*CSSStyleDeclaration) BorderImageWidth() (borderImageWidth string) {
	js.Rewrite("$<.BorderImageWidth")
	return borderImageWidth
}

// BorderImageWidth
func (*CSSStyleDeclaration) SetBorderImageWidth(borderImageWidth string) {
	js.Rewrite("$<.BorderImageWidth = $1", borderImageWidth)
}

// BorderLeft
func (*CSSStyleDeclaration) BorderLeft() (borderLeft string) {
	js.Rewrite("$<.BorderLeft")
	return borderLeft
}

// BorderLeft
func (*CSSStyleDeclaration) SetBorderLeft(borderLeft string) {
	js.Rewrite("$<.BorderLeft = $1", borderLeft)
}

// BorderLeftColor
func (*CSSStyleDeclaration) BorderLeftColor() (borderLeftColor string) {
	js.Rewrite("$<.BorderLeftColor")
	return borderLeftColor
}

// BorderLeftColor
func (*CSSStyleDeclaration) SetBorderLeftColor(borderLeftColor string) {
	js.Rewrite("$<.BorderLeftColor = $1", borderLeftColor)
}

// BorderLeftStyle
func (*CSSStyleDeclaration) BorderLeftStyle() (borderLeftStyle string) {
	js.Rewrite("$<.BorderLeftStyle")
	return borderLeftStyle
}

// BorderLeftStyle
func (*CSSStyleDeclaration) SetBorderLeftStyle(borderLeftStyle string) {
	js.Rewrite("$<.BorderLeftStyle = $1", borderLeftStyle)
}

// BorderLeftWidth
func (*CSSStyleDeclaration) BorderLeftWidth() (borderLeftWidth string) {
	js.Rewrite("$<.BorderLeftWidth")
	return borderLeftWidth
}

// BorderLeftWidth
func (*CSSStyleDeclaration) SetBorderLeftWidth(borderLeftWidth string) {
	js.Rewrite("$<.BorderLeftWidth = $1", borderLeftWidth)
}

// BorderRadius
func (*CSSStyleDeclaration) BorderRadius() (borderRadius string) {
	js.Rewrite("$<.BorderRadius")
	return borderRadius
}

// BorderRadius
func (*CSSStyleDeclaration) SetBorderRadius(borderRadius string) {
	js.Rewrite("$<.BorderRadius = $1", borderRadius)
}

// BorderRight
func (*CSSStyleDeclaration) BorderRight() (borderRight string) {
	js.Rewrite("$<.BorderRight")
	return borderRight
}

// BorderRight
func (*CSSStyleDeclaration) SetBorderRight(borderRight string) {
	js.Rewrite("$<.BorderRight = $1", borderRight)
}

// BorderRightColor
func (*CSSStyleDeclaration) BorderRightColor() (borderRightColor string) {
	js.Rewrite("$<.BorderRightColor")
	return borderRightColor
}

// BorderRightColor
func (*CSSStyleDeclaration) SetBorderRightColor(borderRightColor string) {
	js.Rewrite("$<.BorderRightColor = $1", borderRightColor)
}

// BorderRightStyle
func (*CSSStyleDeclaration) BorderRightStyle() (borderRightStyle string) {
	js.Rewrite("$<.BorderRightStyle")
	return borderRightStyle
}

// BorderRightStyle
func (*CSSStyleDeclaration) SetBorderRightStyle(borderRightStyle string) {
	js.Rewrite("$<.BorderRightStyle = $1", borderRightStyle)
}

// BorderRightWidth
func (*CSSStyleDeclaration) BorderRightWidth() (borderRightWidth string) {
	js.Rewrite("$<.BorderRightWidth")
	return borderRightWidth
}

// BorderRightWidth
func (*CSSStyleDeclaration) SetBorderRightWidth(borderRightWidth string) {
	js.Rewrite("$<.BorderRightWidth = $1", borderRightWidth)
}

// BorderSpacing
func (*CSSStyleDeclaration) BorderSpacing() (borderSpacing string) {
	js.Rewrite("$<.BorderSpacing")
	return borderSpacing
}

// BorderSpacing
func (*CSSStyleDeclaration) SetBorderSpacing(borderSpacing string) {
	js.Rewrite("$<.BorderSpacing = $1", borderSpacing)
}

// BorderStyle
func (*CSSStyleDeclaration) BorderStyle() (borderStyle string) {
	js.Rewrite("$<.BorderStyle")
	return borderStyle
}

// BorderStyle
func (*CSSStyleDeclaration) SetBorderStyle(borderStyle string) {
	js.Rewrite("$<.BorderStyle = $1", borderStyle)
}

// BorderTop
func (*CSSStyleDeclaration) BorderTop() (borderTop string) {
	js.Rewrite("$<.BorderTop")
	return borderTop
}

// BorderTop
func (*CSSStyleDeclaration) SetBorderTop(borderTop string) {
	js.Rewrite("$<.BorderTop = $1", borderTop)
}

// BorderTopColor
func (*CSSStyleDeclaration) BorderTopColor() (borderTopColor string) {
	js.Rewrite("$<.BorderTopColor")
	return borderTopColor
}

// BorderTopColor
func (*CSSStyleDeclaration) SetBorderTopColor(borderTopColor string) {
	js.Rewrite("$<.BorderTopColor = $1", borderTopColor)
}

// BorderTopLeftRadius
func (*CSSStyleDeclaration) BorderTopLeftRadius() (borderTopLeftRadius string) {
	js.Rewrite("$<.BorderTopLeftRadius")
	return borderTopLeftRadius
}

// BorderTopLeftRadius
func (*CSSStyleDeclaration) SetBorderTopLeftRadius(borderTopLeftRadius string) {
	js.Rewrite("$<.BorderTopLeftRadius = $1", borderTopLeftRadius)
}

// BorderTopRightRadius
func (*CSSStyleDeclaration) BorderTopRightRadius() (borderTopRightRadius string) {
	js.Rewrite("$<.BorderTopRightRadius")
	return borderTopRightRadius
}

// BorderTopRightRadius
func (*CSSStyleDeclaration) SetBorderTopRightRadius(borderTopRightRadius string) {
	js.Rewrite("$<.BorderTopRightRadius = $1", borderTopRightRadius)
}

// BorderTopStyle
func (*CSSStyleDeclaration) BorderTopStyle() (borderTopStyle string) {
	js.Rewrite("$<.BorderTopStyle")
	return borderTopStyle
}

// BorderTopStyle
func (*CSSStyleDeclaration) SetBorderTopStyle(borderTopStyle string) {
	js.Rewrite("$<.BorderTopStyle = $1", borderTopStyle)
}

// BorderTopWidth
func (*CSSStyleDeclaration) BorderTopWidth() (borderTopWidth string) {
	js.Rewrite("$<.BorderTopWidth")
	return borderTopWidth
}

// BorderTopWidth
func (*CSSStyleDeclaration) SetBorderTopWidth(borderTopWidth string) {
	js.Rewrite("$<.BorderTopWidth = $1", borderTopWidth)
}

// BorderWidth
func (*CSSStyleDeclaration) BorderWidth() (borderWidth string) {
	js.Rewrite("$<.BorderWidth")
	return borderWidth
}

// BorderWidth
func (*CSSStyleDeclaration) SetBorderWidth(borderWidth string) {
	js.Rewrite("$<.BorderWidth = $1", borderWidth)
}

// Bottom
func (*CSSStyleDeclaration) Bottom() (bottom string) {
	js.Rewrite("$<.Bottom")
	return bottom
}

// Bottom
func (*CSSStyleDeclaration) SetBottom(bottom string) {
	js.Rewrite("$<.Bottom = $1", bottom)
}

// BoxShadow
func (*CSSStyleDeclaration) BoxShadow() (boxShadow string) {
	js.Rewrite("$<.BoxShadow")
	return boxShadow
}

// BoxShadow
func (*CSSStyleDeclaration) SetBoxShadow(boxShadow string) {
	js.Rewrite("$<.BoxShadow = $1", boxShadow)
}

// BoxSizing
func (*CSSStyleDeclaration) BoxSizing() (boxSizing string) {
	js.Rewrite("$<.BoxSizing")
	return boxSizing
}

// BoxSizing
func (*CSSStyleDeclaration) SetBoxSizing(boxSizing string) {
	js.Rewrite("$<.BoxSizing = $1", boxSizing)
}

// BreakAfter
func (*CSSStyleDeclaration) BreakAfter() (breakAfter string) {
	js.Rewrite("$<.BreakAfter")
	return breakAfter
}

// BreakAfter
func (*CSSStyleDeclaration) SetBreakAfter(breakAfter string) {
	js.Rewrite("$<.BreakAfter = $1", breakAfter)
}

// BreakBefore
func (*CSSStyleDeclaration) BreakBefore() (breakBefore string) {
	js.Rewrite("$<.BreakBefore")
	return breakBefore
}

// BreakBefore
func (*CSSStyleDeclaration) SetBreakBefore(breakBefore string) {
	js.Rewrite("$<.BreakBefore = $1", breakBefore)
}

// BreakInside
func (*CSSStyleDeclaration) BreakInside() (breakInside string) {
	js.Rewrite("$<.BreakInside")
	return breakInside
}

// BreakInside
func (*CSSStyleDeclaration) SetBreakInside(breakInside string) {
	js.Rewrite("$<.BreakInside = $1", breakInside)
}

// CaptionSide
func (*CSSStyleDeclaration) CaptionSide() (captionSide string) {
	js.Rewrite("$<.CaptionSide")
	return captionSide
}

// CaptionSide
func (*CSSStyleDeclaration) SetCaptionSide(captionSide string) {
	js.Rewrite("$<.CaptionSide = $1", captionSide)
}

// Clear
func (*CSSStyleDeclaration) Clear() (clear string) {
	js.Rewrite("$<.Clear")
	return clear
}

// Clear
func (*CSSStyleDeclaration) SetClear(clear string) {
	js.Rewrite("$<.Clear = $1", clear)
}

// Clip
func (*CSSStyleDeclaration) Clip() (clip string) {
	js.Rewrite("$<.Clip")
	return clip
}

// Clip
func (*CSSStyleDeclaration) SetClip(clip string) {
	js.Rewrite("$<.Clip = $1", clip)
}

// ClipPath
func (*CSSStyleDeclaration) ClipPath() (clipPath string) {
	js.Rewrite("$<.ClipPath")
	return clipPath
}

// ClipPath
func (*CSSStyleDeclaration) SetClipPath(clipPath string) {
	js.Rewrite("$<.ClipPath = $1", clipPath)
}

// ClipRule
func (*CSSStyleDeclaration) ClipRule() (clipRule string) {
	js.Rewrite("$<.ClipRule")
	return clipRule
}

// ClipRule
func (*CSSStyleDeclaration) SetClipRule(clipRule string) {
	js.Rewrite("$<.ClipRule = $1", clipRule)
}

// Color
func (*CSSStyleDeclaration) Color() (color string) {
	js.Rewrite("$<.Color")
	return color
}

// Color
func (*CSSStyleDeclaration) SetColor(color string) {
	js.Rewrite("$<.Color = $1", color)
}

// ColorInterpolationFilters
func (*CSSStyleDeclaration) ColorInterpolationFilters() (colorInterpolationFilters string) {
	js.Rewrite("$<.ColorInterpolationFilters")
	return colorInterpolationFilters
}

// ColorInterpolationFilters
func (*CSSStyleDeclaration) SetColorInterpolationFilters(colorInterpolationFilters string) {
	js.Rewrite("$<.ColorInterpolationFilters = $1", colorInterpolationFilters)
}

// ColumnCount
func (*CSSStyleDeclaration) ColumnCount() (columnCount interface{}) {
	js.Rewrite("$<.ColumnCount")
	return columnCount
}

// ColumnCount
func (*CSSStyleDeclaration) SetColumnCount(columnCount interface{}) {
	js.Rewrite("$<.ColumnCount = $1", columnCount)
}

// ColumnFill
func (*CSSStyleDeclaration) ColumnFill() (columnFill string) {
	js.Rewrite("$<.ColumnFill")
	return columnFill
}

// ColumnFill
func (*CSSStyleDeclaration) SetColumnFill(columnFill string) {
	js.Rewrite("$<.ColumnFill = $1", columnFill)
}

// ColumnGap
func (*CSSStyleDeclaration) ColumnGap() (columnGap interface{}) {
	js.Rewrite("$<.ColumnGap")
	return columnGap
}

// ColumnGap
func (*CSSStyleDeclaration) SetColumnGap(columnGap interface{}) {
	js.Rewrite("$<.ColumnGap = $1", columnGap)
}

// ColumnRule
func (*CSSStyleDeclaration) ColumnRule() (columnRule string) {
	js.Rewrite("$<.ColumnRule")
	return columnRule
}

// ColumnRule
func (*CSSStyleDeclaration) SetColumnRule(columnRule string) {
	js.Rewrite("$<.ColumnRule = $1", columnRule)
}

// ColumnRuleColor
func (*CSSStyleDeclaration) ColumnRuleColor() (columnRuleColor interface{}) {
	js.Rewrite("$<.ColumnRuleColor")
	return columnRuleColor
}

// ColumnRuleColor
func (*CSSStyleDeclaration) SetColumnRuleColor(columnRuleColor interface{}) {
	js.Rewrite("$<.ColumnRuleColor = $1", columnRuleColor)
}

// ColumnRuleStyle
func (*CSSStyleDeclaration) ColumnRuleStyle() (columnRuleStyle string) {
	js.Rewrite("$<.ColumnRuleStyle")
	return columnRuleStyle
}

// ColumnRuleStyle
func (*CSSStyleDeclaration) SetColumnRuleStyle(columnRuleStyle string) {
	js.Rewrite("$<.ColumnRuleStyle = $1", columnRuleStyle)
}

// ColumnRuleWidth
func (*CSSStyleDeclaration) ColumnRuleWidth() (columnRuleWidth interface{}) {
	js.Rewrite("$<.ColumnRuleWidth")
	return columnRuleWidth
}

// ColumnRuleWidth
func (*CSSStyleDeclaration) SetColumnRuleWidth(columnRuleWidth interface{}) {
	js.Rewrite("$<.ColumnRuleWidth = $1", columnRuleWidth)
}

// Columns
func (*CSSStyleDeclaration) Columns() (columns string) {
	js.Rewrite("$<.Columns")
	return columns
}

// Columns
func (*CSSStyleDeclaration) SetColumns(columns string) {
	js.Rewrite("$<.Columns = $1", columns)
}

// ColumnSpan
func (*CSSStyleDeclaration) ColumnSpan() (columnSpan string) {
	js.Rewrite("$<.ColumnSpan")
	return columnSpan
}

// ColumnSpan
func (*CSSStyleDeclaration) SetColumnSpan(columnSpan string) {
	js.Rewrite("$<.ColumnSpan = $1", columnSpan)
}

// ColumnWidth
func (*CSSStyleDeclaration) ColumnWidth() (columnWidth interface{}) {
	js.Rewrite("$<.ColumnWidth")
	return columnWidth
}

// ColumnWidth
func (*CSSStyleDeclaration) SetColumnWidth(columnWidth interface{}) {
	js.Rewrite("$<.ColumnWidth = $1", columnWidth)
}

// Content
func (*CSSStyleDeclaration) Content() (content string) {
	js.Rewrite("$<.Content")
	return content
}

// Content
func (*CSSStyleDeclaration) SetContent(content string) {
	js.Rewrite("$<.Content = $1", content)
}

// CounterIncrement
func (*CSSStyleDeclaration) CounterIncrement() (counterIncrement string) {
	js.Rewrite("$<.CounterIncrement")
	return counterIncrement
}

// CounterIncrement
func (*CSSStyleDeclaration) SetCounterIncrement(counterIncrement string) {
	js.Rewrite("$<.CounterIncrement = $1", counterIncrement)
}

// CounterReset
func (*CSSStyleDeclaration) CounterReset() (counterReset string) {
	js.Rewrite("$<.CounterReset")
	return counterReset
}

// CounterReset
func (*CSSStyleDeclaration) SetCounterReset(counterReset string) {
	js.Rewrite("$<.CounterReset = $1", counterReset)
}

// CSSFloat
func (*CSSStyleDeclaration) CSSFloat() (cssFloat string) {
	js.Rewrite("$<.CSSFloat")
	return cssFloat
}

// CSSFloat
func (*CSSStyleDeclaration) SetCSSFloat(cssFloat string) {
	js.Rewrite("$<.CSSFloat = $1", cssFloat)
}

// CSSText
func (*CSSStyleDeclaration) CSSText() (cssText string) {
	js.Rewrite("$<.CSSText")
	return cssText
}

// CSSText
func (*CSSStyleDeclaration) SetCSSText(cssText string) {
	js.Rewrite("$<.CSSText = $1", cssText)
}

// Cursor
func (*CSSStyleDeclaration) Cursor() (cursor string) {
	js.Rewrite("$<.Cursor")
	return cursor
}

// Cursor
func (*CSSStyleDeclaration) SetCursor(cursor string) {
	js.Rewrite("$<.Cursor = $1", cursor)
}

// Direction
func (*CSSStyleDeclaration) Direction() (direction string) {
	js.Rewrite("$<.Direction")
	return direction
}

// Direction
func (*CSSStyleDeclaration) SetDirection(direction string) {
	js.Rewrite("$<.Direction = $1", direction)
}

// Display
func (*CSSStyleDeclaration) Display() (display string) {
	js.Rewrite("$<.Display")
	return display
}

// Display
func (*CSSStyleDeclaration) SetDisplay(display string) {
	js.Rewrite("$<.Display = $1", display)
}

// DominantBaseline
func (*CSSStyleDeclaration) DominantBaseline() (dominantBaseline string) {
	js.Rewrite("$<.DominantBaseline")
	return dominantBaseline
}

// DominantBaseline
func (*CSSStyleDeclaration) SetDominantBaseline(dominantBaseline string) {
	js.Rewrite("$<.DominantBaseline = $1", dominantBaseline)
}

// EmptyCells
func (*CSSStyleDeclaration) EmptyCells() (emptyCells string) {
	js.Rewrite("$<.EmptyCells")
	return emptyCells
}

// EmptyCells
func (*CSSStyleDeclaration) SetEmptyCells(emptyCells string) {
	js.Rewrite("$<.EmptyCells = $1", emptyCells)
}

// EnableBackground
func (*CSSStyleDeclaration) EnableBackground() (enableBackground string) {
	js.Rewrite("$<.EnableBackground")
	return enableBackground
}

// EnableBackground
func (*CSSStyleDeclaration) SetEnableBackground(enableBackground string) {
	js.Rewrite("$<.EnableBackground = $1", enableBackground)
}

// Fill
func (*CSSStyleDeclaration) Fill() (fill string) {
	js.Rewrite("$<.Fill")
	return fill
}

// Fill
func (*CSSStyleDeclaration) SetFill(fill string) {
	js.Rewrite("$<.Fill = $1", fill)
}

// FillOpacity
func (*CSSStyleDeclaration) FillOpacity() (fillOpacity string) {
	js.Rewrite("$<.FillOpacity")
	return fillOpacity
}

// FillOpacity
func (*CSSStyleDeclaration) SetFillOpacity(fillOpacity string) {
	js.Rewrite("$<.FillOpacity = $1", fillOpacity)
}

// FillRule
func (*CSSStyleDeclaration) FillRule() (fillRule string) {
	js.Rewrite("$<.FillRule")
	return fillRule
}

// FillRule
func (*CSSStyleDeclaration) SetFillRule(fillRule string) {
	js.Rewrite("$<.FillRule = $1", fillRule)
}

// Filter
func (*CSSStyleDeclaration) Filter() (filter string) {
	js.Rewrite("$<.Filter")
	return filter
}

// Filter
func (*CSSStyleDeclaration) SetFilter(filter string) {
	js.Rewrite("$<.Filter = $1", filter)
}

// Flex
func (*CSSStyleDeclaration) Flex() (flex string) {
	js.Rewrite("$<.Flex")
	return flex
}

// Flex
func (*CSSStyleDeclaration) SetFlex(flex string) {
	js.Rewrite("$<.Flex = $1", flex)
}

// FlexBasis
func (*CSSStyleDeclaration) FlexBasis() (flexBasis string) {
	js.Rewrite("$<.FlexBasis")
	return flexBasis
}

// FlexBasis
func (*CSSStyleDeclaration) SetFlexBasis(flexBasis string) {
	js.Rewrite("$<.FlexBasis = $1", flexBasis)
}

// FlexDirection
func (*CSSStyleDeclaration) FlexDirection() (flexDirection string) {
	js.Rewrite("$<.FlexDirection")
	return flexDirection
}

// FlexDirection
func (*CSSStyleDeclaration) SetFlexDirection(flexDirection string) {
	js.Rewrite("$<.FlexDirection = $1", flexDirection)
}

// FlexFlow
func (*CSSStyleDeclaration) FlexFlow() (flexFlow string) {
	js.Rewrite("$<.FlexFlow")
	return flexFlow
}

// FlexFlow
func (*CSSStyleDeclaration) SetFlexFlow(flexFlow string) {
	js.Rewrite("$<.FlexFlow = $1", flexFlow)
}

// FlexGrow
func (*CSSStyleDeclaration) FlexGrow() (flexGrow string) {
	js.Rewrite("$<.FlexGrow")
	return flexGrow
}

// FlexGrow
func (*CSSStyleDeclaration) SetFlexGrow(flexGrow string) {
	js.Rewrite("$<.FlexGrow = $1", flexGrow)
}

// FlexShrink
func (*CSSStyleDeclaration) FlexShrink() (flexShrink string) {
	js.Rewrite("$<.FlexShrink")
	return flexShrink
}

// FlexShrink
func (*CSSStyleDeclaration) SetFlexShrink(flexShrink string) {
	js.Rewrite("$<.FlexShrink = $1", flexShrink)
}

// FlexWrap
func (*CSSStyleDeclaration) FlexWrap() (flexWrap string) {
	js.Rewrite("$<.FlexWrap")
	return flexWrap
}

// FlexWrap
func (*CSSStyleDeclaration) SetFlexWrap(flexWrap string) {
	js.Rewrite("$<.FlexWrap = $1", flexWrap)
}

// FloodColor
func (*CSSStyleDeclaration) FloodColor() (floodColor string) {
	js.Rewrite("$<.FloodColor")
	return floodColor
}

// FloodColor
func (*CSSStyleDeclaration) SetFloodColor(floodColor string) {
	js.Rewrite("$<.FloodColor = $1", floodColor)
}

// FloodOpacity
func (*CSSStyleDeclaration) FloodOpacity() (floodOpacity string) {
	js.Rewrite("$<.FloodOpacity")
	return floodOpacity
}

// FloodOpacity
func (*CSSStyleDeclaration) SetFloodOpacity(floodOpacity string) {
	js.Rewrite("$<.FloodOpacity = $1", floodOpacity)
}

// Font
func (*CSSStyleDeclaration) Font() (font string) {
	js.Rewrite("$<.Font")
	return font
}

// Font
func (*CSSStyleDeclaration) SetFont(font string) {
	js.Rewrite("$<.Font = $1", font)
}

// FontFamily
func (*CSSStyleDeclaration) FontFamily() (fontFamily string) {
	js.Rewrite("$<.FontFamily")
	return fontFamily
}

// FontFamily
func (*CSSStyleDeclaration) SetFontFamily(fontFamily string) {
	js.Rewrite("$<.FontFamily = $1", fontFamily)
}

// FontFeatureSettings
func (*CSSStyleDeclaration) FontFeatureSettings() (fontFeatureSettings string) {
	js.Rewrite("$<.FontFeatureSettings")
	return fontFeatureSettings
}

// FontFeatureSettings
func (*CSSStyleDeclaration) SetFontFeatureSettings(fontFeatureSettings string) {
	js.Rewrite("$<.FontFeatureSettings = $1", fontFeatureSettings)
}

// FontSize
func (*CSSStyleDeclaration) FontSize() (fontSize string) {
	js.Rewrite("$<.FontSize")
	return fontSize
}

// FontSize
func (*CSSStyleDeclaration) SetFontSize(fontSize string) {
	js.Rewrite("$<.FontSize = $1", fontSize)
}

// FontSizeAdjust
func (*CSSStyleDeclaration) FontSizeAdjust() (fontSizeAdjust string) {
	js.Rewrite("$<.FontSizeAdjust")
	return fontSizeAdjust
}

// FontSizeAdjust
func (*CSSStyleDeclaration) SetFontSizeAdjust(fontSizeAdjust string) {
	js.Rewrite("$<.FontSizeAdjust = $1", fontSizeAdjust)
}

// FontStretch
func (*CSSStyleDeclaration) FontStretch() (fontStretch string) {
	js.Rewrite("$<.FontStretch")
	return fontStretch
}

// FontStretch
func (*CSSStyleDeclaration) SetFontStretch(fontStretch string) {
	js.Rewrite("$<.FontStretch = $1", fontStretch)
}

// FontStyle
func (*CSSStyleDeclaration) FontStyle() (fontStyle string) {
	js.Rewrite("$<.FontStyle")
	return fontStyle
}

// FontStyle
func (*CSSStyleDeclaration) SetFontStyle(fontStyle string) {
	js.Rewrite("$<.FontStyle = $1", fontStyle)
}

// FontVariant
func (*CSSStyleDeclaration) FontVariant() (fontVariant string) {
	js.Rewrite("$<.FontVariant")
	return fontVariant
}

// FontVariant
func (*CSSStyleDeclaration) SetFontVariant(fontVariant string) {
	js.Rewrite("$<.FontVariant = $1", fontVariant)
}

// FontWeight
func (*CSSStyleDeclaration) FontWeight() (fontWeight string) {
	js.Rewrite("$<.FontWeight")
	return fontWeight
}

// FontWeight
func (*CSSStyleDeclaration) SetFontWeight(fontWeight string) {
	js.Rewrite("$<.FontWeight = $1", fontWeight)
}

// GlyphOrientationHorizontal
func (*CSSStyleDeclaration) GlyphOrientationHorizontal() (glyphOrientationHorizontal string) {
	js.Rewrite("$<.GlyphOrientationHorizontal")
	return glyphOrientationHorizontal
}

// GlyphOrientationHorizontal
func (*CSSStyleDeclaration) SetGlyphOrientationHorizontal(glyphOrientationHorizontal string) {
	js.Rewrite("$<.GlyphOrientationHorizontal = $1", glyphOrientationHorizontal)
}

// GlyphOrientationVertical
func (*CSSStyleDeclaration) GlyphOrientationVertical() (glyphOrientationVertical string) {
	js.Rewrite("$<.GlyphOrientationVertical")
	return glyphOrientationVertical
}

// GlyphOrientationVertical
func (*CSSStyleDeclaration) SetGlyphOrientationVertical(glyphOrientationVertical string) {
	js.Rewrite("$<.GlyphOrientationVertical = $1", glyphOrientationVertical)
}

// Height
func (*CSSStyleDeclaration) Height() (height string) {
	js.Rewrite("$<.Height")
	return height
}

// Height
func (*CSSStyleDeclaration) SetHeight(height string) {
	js.Rewrite("$<.Height = $1", height)
}

// ImeMode
func (*CSSStyleDeclaration) ImeMode() (imeMode string) {
	js.Rewrite("$<.ImeMode")
	return imeMode
}

// ImeMode
func (*CSSStyleDeclaration) SetImeMode(imeMode string) {
	js.Rewrite("$<.ImeMode = $1", imeMode)
}

// JustifyContent
func (*CSSStyleDeclaration) JustifyContent() (justifyContent string) {
	js.Rewrite("$<.JustifyContent")
	return justifyContent
}

// JustifyContent
func (*CSSStyleDeclaration) SetJustifyContent(justifyContent string) {
	js.Rewrite("$<.JustifyContent = $1", justifyContent)
}

// Kerning
func (*CSSStyleDeclaration) Kerning() (kerning string) {
	js.Rewrite("$<.Kerning")
	return kerning
}

// Kerning
func (*CSSStyleDeclaration) SetKerning(kerning string) {
	js.Rewrite("$<.Kerning = $1", kerning)
}

// LayoutGrid
func (*CSSStyleDeclaration) LayoutGrid() (layoutGrid string) {
	js.Rewrite("$<.LayoutGrid")
	return layoutGrid
}

// LayoutGrid
func (*CSSStyleDeclaration) SetLayoutGrid(layoutGrid string) {
	js.Rewrite("$<.LayoutGrid = $1", layoutGrid)
}

// LayoutGridChar
func (*CSSStyleDeclaration) LayoutGridChar() (layoutGridChar string) {
	js.Rewrite("$<.LayoutGridChar")
	return layoutGridChar
}

// LayoutGridChar
func (*CSSStyleDeclaration) SetLayoutGridChar(layoutGridChar string) {
	js.Rewrite("$<.LayoutGridChar = $1", layoutGridChar)
}

// LayoutGridLine
func (*CSSStyleDeclaration) LayoutGridLine() (layoutGridLine string) {
	js.Rewrite("$<.LayoutGridLine")
	return layoutGridLine
}

// LayoutGridLine
func (*CSSStyleDeclaration) SetLayoutGridLine(layoutGridLine string) {
	js.Rewrite("$<.LayoutGridLine = $1", layoutGridLine)
}

// LayoutGridMode
func (*CSSStyleDeclaration) LayoutGridMode() (layoutGridMode string) {
	js.Rewrite("$<.LayoutGridMode")
	return layoutGridMode
}

// LayoutGridMode
func (*CSSStyleDeclaration) SetLayoutGridMode(layoutGridMode string) {
	js.Rewrite("$<.LayoutGridMode = $1", layoutGridMode)
}

// LayoutGridType
func (*CSSStyleDeclaration) LayoutGridType() (layoutGridType string) {
	js.Rewrite("$<.LayoutGridType")
	return layoutGridType
}

// LayoutGridType
func (*CSSStyleDeclaration) SetLayoutGridType(layoutGridType string) {
	js.Rewrite("$<.LayoutGridType = $1", layoutGridType)
}

// Left
func (*CSSStyleDeclaration) Left() (left string) {
	js.Rewrite("$<.Left")
	return left
}

// Left
func (*CSSStyleDeclaration) SetLeft(left string) {
	js.Rewrite("$<.Left = $1", left)
}

// Length
func (*CSSStyleDeclaration) Length() (length uint) {
	js.Rewrite("$<.Length")
	return length
}

// LetterSpacing
func (*CSSStyleDeclaration) LetterSpacing() (letterSpacing string) {
	js.Rewrite("$<.LetterSpacing")
	return letterSpacing
}

// LetterSpacing
func (*CSSStyleDeclaration) SetLetterSpacing(letterSpacing string) {
	js.Rewrite("$<.LetterSpacing = $1", letterSpacing)
}

// LightingColor
func (*CSSStyleDeclaration) LightingColor() (lightingColor string) {
	js.Rewrite("$<.LightingColor")
	return lightingColor
}

// LightingColor
func (*CSSStyleDeclaration) SetLightingColor(lightingColor string) {
	js.Rewrite("$<.LightingColor = $1", lightingColor)
}

// LineBreak
func (*CSSStyleDeclaration) LineBreak() (lineBreak string) {
	js.Rewrite("$<.LineBreak")
	return lineBreak
}

// LineBreak
func (*CSSStyleDeclaration) SetLineBreak(lineBreak string) {
	js.Rewrite("$<.LineBreak = $1", lineBreak)
}

// LineHeight
func (*CSSStyleDeclaration) LineHeight() (lineHeight string) {
	js.Rewrite("$<.LineHeight")
	return lineHeight
}

// LineHeight
func (*CSSStyleDeclaration) SetLineHeight(lineHeight string) {
	js.Rewrite("$<.LineHeight = $1", lineHeight)
}

// ListStyle
func (*CSSStyleDeclaration) ListStyle() (listStyle string) {
	js.Rewrite("$<.ListStyle")
	return listStyle
}

// ListStyle
func (*CSSStyleDeclaration) SetListStyle(listStyle string) {
	js.Rewrite("$<.ListStyle = $1", listStyle)
}

// ListStyleImage
func (*CSSStyleDeclaration) ListStyleImage() (listStyleImage string) {
	js.Rewrite("$<.ListStyleImage")
	return listStyleImage
}

// ListStyleImage
func (*CSSStyleDeclaration) SetListStyleImage(listStyleImage string) {
	js.Rewrite("$<.ListStyleImage = $1", listStyleImage)
}

// ListStylePosition
func (*CSSStyleDeclaration) ListStylePosition() (listStylePosition string) {
	js.Rewrite("$<.ListStylePosition")
	return listStylePosition
}

// ListStylePosition
func (*CSSStyleDeclaration) SetListStylePosition(listStylePosition string) {
	js.Rewrite("$<.ListStylePosition = $1", listStylePosition)
}

// ListStyleType
func (*CSSStyleDeclaration) ListStyleType() (listStyleType string) {
	js.Rewrite("$<.ListStyleType")
	return listStyleType
}

// ListStyleType
func (*CSSStyleDeclaration) SetListStyleType(listStyleType string) {
	js.Rewrite("$<.ListStyleType = $1", listStyleType)
}

// Margin
func (*CSSStyleDeclaration) Margin() (margin string) {
	js.Rewrite("$<.Margin")
	return margin
}

// Margin
func (*CSSStyleDeclaration) SetMargin(margin string) {
	js.Rewrite("$<.Margin = $1", margin)
}

// MarginBottom
func (*CSSStyleDeclaration) MarginBottom() (marginBottom string) {
	js.Rewrite("$<.MarginBottom")
	return marginBottom
}

// MarginBottom
func (*CSSStyleDeclaration) SetMarginBottom(marginBottom string) {
	js.Rewrite("$<.MarginBottom = $1", marginBottom)
}

// MarginLeft
func (*CSSStyleDeclaration) MarginLeft() (marginLeft string) {
	js.Rewrite("$<.MarginLeft")
	return marginLeft
}

// MarginLeft
func (*CSSStyleDeclaration) SetMarginLeft(marginLeft string) {
	js.Rewrite("$<.MarginLeft = $1", marginLeft)
}

// MarginRight
func (*CSSStyleDeclaration) MarginRight() (marginRight string) {
	js.Rewrite("$<.MarginRight")
	return marginRight
}

// MarginRight
func (*CSSStyleDeclaration) SetMarginRight(marginRight string) {
	js.Rewrite("$<.MarginRight = $1", marginRight)
}

// MarginTop
func (*CSSStyleDeclaration) MarginTop() (marginTop string) {
	js.Rewrite("$<.MarginTop")
	return marginTop
}

// MarginTop
func (*CSSStyleDeclaration) SetMarginTop(marginTop string) {
	js.Rewrite("$<.MarginTop = $1", marginTop)
}

// Marker
func (*CSSStyleDeclaration) Marker() (marker string) {
	js.Rewrite("$<.Marker")
	return marker
}

// Marker
func (*CSSStyleDeclaration) SetMarker(marker string) {
	js.Rewrite("$<.Marker = $1", marker)
}

// MarkerEnd
func (*CSSStyleDeclaration) MarkerEnd() (markerEnd string) {
	js.Rewrite("$<.MarkerEnd")
	return markerEnd
}

// MarkerEnd
func (*CSSStyleDeclaration) SetMarkerEnd(markerEnd string) {
	js.Rewrite("$<.MarkerEnd = $1", markerEnd)
}

// MarkerMid
func (*CSSStyleDeclaration) MarkerMid() (markerMid string) {
	js.Rewrite("$<.MarkerMid")
	return markerMid
}

// MarkerMid
func (*CSSStyleDeclaration) SetMarkerMid(markerMid string) {
	js.Rewrite("$<.MarkerMid = $1", markerMid)
}

// MarkerStart
func (*CSSStyleDeclaration) MarkerStart() (markerStart string) {
	js.Rewrite("$<.MarkerStart")
	return markerStart
}

// MarkerStart
func (*CSSStyleDeclaration) SetMarkerStart(markerStart string) {
	js.Rewrite("$<.MarkerStart = $1", markerStart)
}

// Mask
func (*CSSStyleDeclaration) Mask() (mask string) {
	js.Rewrite("$<.Mask")
	return mask
}

// Mask
func (*CSSStyleDeclaration) SetMask(mask string) {
	js.Rewrite("$<.Mask = $1", mask)
}

// MaxHeight
func (*CSSStyleDeclaration) MaxHeight() (maxHeight string) {
	js.Rewrite("$<.MaxHeight")
	return maxHeight
}

// MaxHeight
func (*CSSStyleDeclaration) SetMaxHeight(maxHeight string) {
	js.Rewrite("$<.MaxHeight = $1", maxHeight)
}

// MaxWidth
func (*CSSStyleDeclaration) MaxWidth() (maxWidth string) {
	js.Rewrite("$<.MaxWidth")
	return maxWidth
}

// MaxWidth
func (*CSSStyleDeclaration) SetMaxWidth(maxWidth string) {
	js.Rewrite("$<.MaxWidth = $1", maxWidth)
}

// MinHeight
func (*CSSStyleDeclaration) MinHeight() (minHeight string) {
	js.Rewrite("$<.MinHeight")
	return minHeight
}

// MinHeight
func (*CSSStyleDeclaration) SetMinHeight(minHeight string) {
	js.Rewrite("$<.MinHeight = $1", minHeight)
}

// MinWidth
func (*CSSStyleDeclaration) MinWidth() (minWidth string) {
	js.Rewrite("$<.MinWidth")
	return minWidth
}

// MinWidth
func (*CSSStyleDeclaration) SetMinWidth(minWidth string) {
	js.Rewrite("$<.MinWidth = $1", minWidth)
}

// MsContentZoomChaining
func (*CSSStyleDeclaration) MsContentZoomChaining() (msContentZoomChaining string) {
	js.Rewrite("$<.MsContentZoomChaining")
	return msContentZoomChaining
}

// MsContentZoomChaining
func (*CSSStyleDeclaration) SetMsContentZoomChaining(msContentZoomChaining string) {
	js.Rewrite("$<.MsContentZoomChaining = $1", msContentZoomChaining)
}

// MsContentZooming
func (*CSSStyleDeclaration) MsContentZooming() (msContentZooming string) {
	js.Rewrite("$<.MsContentZooming")
	return msContentZooming
}

// MsContentZooming
func (*CSSStyleDeclaration) SetMsContentZooming(msContentZooming string) {
	js.Rewrite("$<.MsContentZooming = $1", msContentZooming)
}

// MsContentZoomLimit
func (*CSSStyleDeclaration) MsContentZoomLimit() (msContentZoomLimit string) {
	js.Rewrite("$<.MsContentZoomLimit")
	return msContentZoomLimit
}

// MsContentZoomLimit
func (*CSSStyleDeclaration) SetMsContentZoomLimit(msContentZoomLimit string) {
	js.Rewrite("$<.MsContentZoomLimit = $1", msContentZoomLimit)
}

// MsContentZoomLimitMax
func (*CSSStyleDeclaration) MsContentZoomLimitMax() (msContentZoomLimitMax interface{}) {
	js.Rewrite("$<.MsContentZoomLimitMax")
	return msContentZoomLimitMax
}

// MsContentZoomLimitMax
func (*CSSStyleDeclaration) SetMsContentZoomLimitMax(msContentZoomLimitMax interface{}) {
	js.Rewrite("$<.MsContentZoomLimitMax = $1", msContentZoomLimitMax)
}

// MsContentZoomLimitMin
func (*CSSStyleDeclaration) MsContentZoomLimitMin() (msContentZoomLimitMin interface{}) {
	js.Rewrite("$<.MsContentZoomLimitMin")
	return msContentZoomLimitMin
}

// MsContentZoomLimitMin
func (*CSSStyleDeclaration) SetMsContentZoomLimitMin(msContentZoomLimitMin interface{}) {
	js.Rewrite("$<.MsContentZoomLimitMin = $1", msContentZoomLimitMin)
}

// MsContentZoomSnap
func (*CSSStyleDeclaration) MsContentZoomSnap() (msContentZoomSnap string) {
	js.Rewrite("$<.MsContentZoomSnap")
	return msContentZoomSnap
}

// MsContentZoomSnap
func (*CSSStyleDeclaration) SetMsContentZoomSnap(msContentZoomSnap string) {
	js.Rewrite("$<.MsContentZoomSnap = $1", msContentZoomSnap)
}

// MsContentZoomSnapPoints
func (*CSSStyleDeclaration) MsContentZoomSnapPoints() (msContentZoomSnapPoints string) {
	js.Rewrite("$<.MsContentZoomSnapPoints")
	return msContentZoomSnapPoints
}

// MsContentZoomSnapPoints
func (*CSSStyleDeclaration) SetMsContentZoomSnapPoints(msContentZoomSnapPoints string) {
	js.Rewrite("$<.MsContentZoomSnapPoints = $1", msContentZoomSnapPoints)
}

// MsContentZoomSnapType
func (*CSSStyleDeclaration) MsContentZoomSnapType() (msContentZoomSnapType string) {
	js.Rewrite("$<.MsContentZoomSnapType")
	return msContentZoomSnapType
}

// MsContentZoomSnapType
func (*CSSStyleDeclaration) SetMsContentZoomSnapType(msContentZoomSnapType string) {
	js.Rewrite("$<.MsContentZoomSnapType = $1", msContentZoomSnapType)
}

// MsFlowFrom
func (*CSSStyleDeclaration) MsFlowFrom() (msFlowFrom string) {
	js.Rewrite("$<.MsFlowFrom")
	return msFlowFrom
}

// MsFlowFrom
func (*CSSStyleDeclaration) SetMsFlowFrom(msFlowFrom string) {
	js.Rewrite("$<.MsFlowFrom = $1", msFlowFrom)
}

// MsFlowInto
func (*CSSStyleDeclaration) MsFlowInto() (msFlowInto string) {
	js.Rewrite("$<.MsFlowInto")
	return msFlowInto
}

// MsFlowInto
func (*CSSStyleDeclaration) SetMsFlowInto(msFlowInto string) {
	js.Rewrite("$<.MsFlowInto = $1", msFlowInto)
}

// MsFontFeatureSettings
func (*CSSStyleDeclaration) MsFontFeatureSettings() (msFontFeatureSettings string) {
	js.Rewrite("$<.MsFontFeatureSettings")
	return msFontFeatureSettings
}

// MsFontFeatureSettings
func (*CSSStyleDeclaration) SetMsFontFeatureSettings(msFontFeatureSettings string) {
	js.Rewrite("$<.MsFontFeatureSettings = $1", msFontFeatureSettings)
}

// MsGridColumn
func (*CSSStyleDeclaration) MsGridColumn() (msGridColumn interface{}) {
	js.Rewrite("$<.MsGridColumn")
	return msGridColumn
}

// MsGridColumn
func (*CSSStyleDeclaration) SetMsGridColumn(msGridColumn interface{}) {
	js.Rewrite("$<.MsGridColumn = $1", msGridColumn)
}

// MsGridColumnAlign
func (*CSSStyleDeclaration) MsGridColumnAlign() (msGridColumnAlign string) {
	js.Rewrite("$<.MsGridColumnAlign")
	return msGridColumnAlign
}

// MsGridColumnAlign
func (*CSSStyleDeclaration) SetMsGridColumnAlign(msGridColumnAlign string) {
	js.Rewrite("$<.MsGridColumnAlign = $1", msGridColumnAlign)
}

// MsGridColumns
func (*CSSStyleDeclaration) MsGridColumns() (msGridColumns string) {
	js.Rewrite("$<.MsGridColumns")
	return msGridColumns
}

// MsGridColumns
func (*CSSStyleDeclaration) SetMsGridColumns(msGridColumns string) {
	js.Rewrite("$<.MsGridColumns = $1", msGridColumns)
}

// MsGridColumnSpan
func (*CSSStyleDeclaration) MsGridColumnSpan() (msGridColumnSpan interface{}) {
	js.Rewrite("$<.MsGridColumnSpan")
	return msGridColumnSpan
}

// MsGridColumnSpan
func (*CSSStyleDeclaration) SetMsGridColumnSpan(msGridColumnSpan interface{}) {
	js.Rewrite("$<.MsGridColumnSpan = $1", msGridColumnSpan)
}

// MsGridRow
func (*CSSStyleDeclaration) MsGridRow() (msGridRow interface{}) {
	js.Rewrite("$<.MsGridRow")
	return msGridRow
}

// MsGridRow
func (*CSSStyleDeclaration) SetMsGridRow(msGridRow interface{}) {
	js.Rewrite("$<.MsGridRow = $1", msGridRow)
}

// MsGridRowAlign
func (*CSSStyleDeclaration) MsGridRowAlign() (msGridRowAlign string) {
	js.Rewrite("$<.MsGridRowAlign")
	return msGridRowAlign
}

// MsGridRowAlign
func (*CSSStyleDeclaration) SetMsGridRowAlign(msGridRowAlign string) {
	js.Rewrite("$<.MsGridRowAlign = $1", msGridRowAlign)
}

// MsGridRows
func (*CSSStyleDeclaration) MsGridRows() (msGridRows string) {
	js.Rewrite("$<.MsGridRows")
	return msGridRows
}

// MsGridRows
func (*CSSStyleDeclaration) SetMsGridRows(msGridRows string) {
	js.Rewrite("$<.MsGridRows = $1", msGridRows)
}

// MsGridRowSpan
func (*CSSStyleDeclaration) MsGridRowSpan() (msGridRowSpan interface{}) {
	js.Rewrite("$<.MsGridRowSpan")
	return msGridRowSpan
}

// MsGridRowSpan
func (*CSSStyleDeclaration) SetMsGridRowSpan(msGridRowSpan interface{}) {
	js.Rewrite("$<.MsGridRowSpan = $1", msGridRowSpan)
}

// MsHighContrastAdjust
func (*CSSStyleDeclaration) MsHighContrastAdjust() (msHighContrastAdjust string) {
	js.Rewrite("$<.MsHighContrastAdjust")
	return msHighContrastAdjust
}

// MsHighContrastAdjust
func (*CSSStyleDeclaration) SetMsHighContrastAdjust(msHighContrastAdjust string) {
	js.Rewrite("$<.MsHighContrastAdjust = $1", msHighContrastAdjust)
}

// MsHyphenateLimitChars
func (*CSSStyleDeclaration) MsHyphenateLimitChars() (msHyphenateLimitChars string) {
	js.Rewrite("$<.MsHyphenateLimitChars")
	return msHyphenateLimitChars
}

// MsHyphenateLimitChars
func (*CSSStyleDeclaration) SetMsHyphenateLimitChars(msHyphenateLimitChars string) {
	js.Rewrite("$<.MsHyphenateLimitChars = $1", msHyphenateLimitChars)
}

// MsHyphenateLimitLines
func (*CSSStyleDeclaration) MsHyphenateLimitLines() (msHyphenateLimitLines interface{}) {
	js.Rewrite("$<.MsHyphenateLimitLines")
	return msHyphenateLimitLines
}

// MsHyphenateLimitLines
func (*CSSStyleDeclaration) SetMsHyphenateLimitLines(msHyphenateLimitLines interface{}) {
	js.Rewrite("$<.MsHyphenateLimitLines = $1", msHyphenateLimitLines)
}

// MsHyphenateLimitZone
func (*CSSStyleDeclaration) MsHyphenateLimitZone() (msHyphenateLimitZone interface{}) {
	js.Rewrite("$<.MsHyphenateLimitZone")
	return msHyphenateLimitZone
}

// MsHyphenateLimitZone
func (*CSSStyleDeclaration) SetMsHyphenateLimitZone(msHyphenateLimitZone interface{}) {
	js.Rewrite("$<.MsHyphenateLimitZone = $1", msHyphenateLimitZone)
}

// MsHyphens
func (*CSSStyleDeclaration) MsHyphens() (msHyphens string) {
	js.Rewrite("$<.MsHyphens")
	return msHyphens
}

// MsHyphens
func (*CSSStyleDeclaration) SetMsHyphens(msHyphens string) {
	js.Rewrite("$<.MsHyphens = $1", msHyphens)
}

// MsImeAlign
func (*CSSStyleDeclaration) MsImeAlign() (msImeAlign string) {
	js.Rewrite("$<.MsImeAlign")
	return msImeAlign
}

// MsImeAlign
func (*CSSStyleDeclaration) SetMsImeAlign(msImeAlign string) {
	js.Rewrite("$<.MsImeAlign = $1", msImeAlign)
}

// MsOverflowStyle
func (*CSSStyleDeclaration) MsOverflowStyle() (msOverflowStyle string) {
	js.Rewrite("$<.MsOverflowStyle")
	return msOverflowStyle
}

// MsOverflowStyle
func (*CSSStyleDeclaration) SetMsOverflowStyle(msOverflowStyle string) {
	js.Rewrite("$<.MsOverflowStyle = $1", msOverflowStyle)
}

// MsScrollChaining
func (*CSSStyleDeclaration) MsScrollChaining() (msScrollChaining string) {
	js.Rewrite("$<.MsScrollChaining")
	return msScrollChaining
}

// MsScrollChaining
func (*CSSStyleDeclaration) SetMsScrollChaining(msScrollChaining string) {
	js.Rewrite("$<.MsScrollChaining = $1", msScrollChaining)
}

// MsScrollLimit
func (*CSSStyleDeclaration) MsScrollLimit() (msScrollLimit string) {
	js.Rewrite("$<.MsScrollLimit")
	return msScrollLimit
}

// MsScrollLimit
func (*CSSStyleDeclaration) SetMsScrollLimit(msScrollLimit string) {
	js.Rewrite("$<.MsScrollLimit = $1", msScrollLimit)
}

// MsScrollLimitXMax
func (*CSSStyleDeclaration) MsScrollLimitXMax() (msScrollLimitXMax interface{}) {
	js.Rewrite("$<.MsScrollLimitXMax")
	return msScrollLimitXMax
}

// MsScrollLimitXMax
func (*CSSStyleDeclaration) SetMsScrollLimitXMax(msScrollLimitXMax interface{}) {
	js.Rewrite("$<.MsScrollLimitXMax = $1", msScrollLimitXMax)
}

// MsScrollLimitXMin
func (*CSSStyleDeclaration) MsScrollLimitXMin() (msScrollLimitXMin interface{}) {
	js.Rewrite("$<.MsScrollLimitXMin")
	return msScrollLimitXMin
}

// MsScrollLimitXMin
func (*CSSStyleDeclaration) SetMsScrollLimitXMin(msScrollLimitXMin interface{}) {
	js.Rewrite("$<.MsScrollLimitXMin = $1", msScrollLimitXMin)
}

// MsScrollLimitYMax
func (*CSSStyleDeclaration) MsScrollLimitYMax() (msScrollLimitYMax interface{}) {
	js.Rewrite("$<.MsScrollLimitYMax")
	return msScrollLimitYMax
}

// MsScrollLimitYMax
func (*CSSStyleDeclaration) SetMsScrollLimitYMax(msScrollLimitYMax interface{}) {
	js.Rewrite("$<.MsScrollLimitYMax = $1", msScrollLimitYMax)
}

// MsScrollLimitYMin
func (*CSSStyleDeclaration) MsScrollLimitYMin() (msScrollLimitYMin interface{}) {
	js.Rewrite("$<.MsScrollLimitYMin")
	return msScrollLimitYMin
}

// MsScrollLimitYMin
func (*CSSStyleDeclaration) SetMsScrollLimitYMin(msScrollLimitYMin interface{}) {
	js.Rewrite("$<.MsScrollLimitYMin = $1", msScrollLimitYMin)
}

// MsScrollRails
func (*CSSStyleDeclaration) MsScrollRails() (msScrollRails string) {
	js.Rewrite("$<.MsScrollRails")
	return msScrollRails
}

// MsScrollRails
func (*CSSStyleDeclaration) SetMsScrollRails(msScrollRails string) {
	js.Rewrite("$<.MsScrollRails = $1", msScrollRails)
}

// MsScrollSnapPointsX
func (*CSSStyleDeclaration) MsScrollSnapPointsX() (msScrollSnapPointsX string) {
	js.Rewrite("$<.MsScrollSnapPointsX")
	return msScrollSnapPointsX
}

// MsScrollSnapPointsX
func (*CSSStyleDeclaration) SetMsScrollSnapPointsX(msScrollSnapPointsX string) {
	js.Rewrite("$<.MsScrollSnapPointsX = $1", msScrollSnapPointsX)
}

// MsScrollSnapPointsY
func (*CSSStyleDeclaration) MsScrollSnapPointsY() (msScrollSnapPointsY string) {
	js.Rewrite("$<.MsScrollSnapPointsY")
	return msScrollSnapPointsY
}

// MsScrollSnapPointsY
func (*CSSStyleDeclaration) SetMsScrollSnapPointsY(msScrollSnapPointsY string) {
	js.Rewrite("$<.MsScrollSnapPointsY = $1", msScrollSnapPointsY)
}

// MsScrollSnapType
func (*CSSStyleDeclaration) MsScrollSnapType() (msScrollSnapType string) {
	js.Rewrite("$<.MsScrollSnapType")
	return msScrollSnapType
}

// MsScrollSnapType
func (*CSSStyleDeclaration) SetMsScrollSnapType(msScrollSnapType string) {
	js.Rewrite("$<.MsScrollSnapType = $1", msScrollSnapType)
}

// MsScrollSnapX
func (*CSSStyleDeclaration) MsScrollSnapX() (msScrollSnapX string) {
	js.Rewrite("$<.MsScrollSnapX")
	return msScrollSnapX
}

// MsScrollSnapX
func (*CSSStyleDeclaration) SetMsScrollSnapX(msScrollSnapX string) {
	js.Rewrite("$<.MsScrollSnapX = $1", msScrollSnapX)
}

// MsScrollSnapY
func (*CSSStyleDeclaration) MsScrollSnapY() (msScrollSnapY string) {
	js.Rewrite("$<.MsScrollSnapY")
	return msScrollSnapY
}

// MsScrollSnapY
func (*CSSStyleDeclaration) SetMsScrollSnapY(msScrollSnapY string) {
	js.Rewrite("$<.MsScrollSnapY = $1", msScrollSnapY)
}

// MsScrollTranslation
func (*CSSStyleDeclaration) MsScrollTranslation() (msScrollTranslation string) {
	js.Rewrite("$<.MsScrollTranslation")
	return msScrollTranslation
}

// MsScrollTranslation
func (*CSSStyleDeclaration) SetMsScrollTranslation(msScrollTranslation string) {
	js.Rewrite("$<.MsScrollTranslation = $1", msScrollTranslation)
}

// MsTextCombineHorizontal
func (*CSSStyleDeclaration) MsTextCombineHorizontal() (msTextCombineHorizontal string) {
	js.Rewrite("$<.MsTextCombineHorizontal")
	return msTextCombineHorizontal
}

// MsTextCombineHorizontal
func (*CSSStyleDeclaration) SetMsTextCombineHorizontal(msTextCombineHorizontal string) {
	js.Rewrite("$<.MsTextCombineHorizontal = $1", msTextCombineHorizontal)
}

// MsTextSizeAdjust
func (*CSSStyleDeclaration) MsTextSizeAdjust() (msTextSizeAdjust interface{}) {
	js.Rewrite("$<.MsTextSizeAdjust")
	return msTextSizeAdjust
}

// MsTextSizeAdjust
func (*CSSStyleDeclaration) SetMsTextSizeAdjust(msTextSizeAdjust interface{}) {
	js.Rewrite("$<.MsTextSizeAdjust = $1", msTextSizeAdjust)
}

// MsTouchAction
func (*CSSStyleDeclaration) MsTouchAction() (msTouchAction string) {
	js.Rewrite("$<.MsTouchAction")
	return msTouchAction
}

// MsTouchAction
func (*CSSStyleDeclaration) SetMsTouchAction(msTouchAction string) {
	js.Rewrite("$<.MsTouchAction = $1", msTouchAction)
}

// MsTouchSelect
func (*CSSStyleDeclaration) MsTouchSelect() (msTouchSelect string) {
	js.Rewrite("$<.MsTouchSelect")
	return msTouchSelect
}

// MsTouchSelect
func (*CSSStyleDeclaration) SetMsTouchSelect(msTouchSelect string) {
	js.Rewrite("$<.MsTouchSelect = $1", msTouchSelect)
}

// MsUserSelect
func (*CSSStyleDeclaration) MsUserSelect() (msUserSelect string) {
	js.Rewrite("$<.MsUserSelect")
	return msUserSelect
}

// MsUserSelect
func (*CSSStyleDeclaration) SetMsUserSelect(msUserSelect string) {
	js.Rewrite("$<.MsUserSelect = $1", msUserSelect)
}

// MsWrapFlow
func (*CSSStyleDeclaration) MsWrapFlow() (msWrapFlow string) {
	js.Rewrite("$<.MsWrapFlow")
	return msWrapFlow
}

// MsWrapFlow
func (*CSSStyleDeclaration) SetMsWrapFlow(msWrapFlow string) {
	js.Rewrite("$<.MsWrapFlow = $1", msWrapFlow)
}

// MsWrapMargin
func (*CSSStyleDeclaration) MsWrapMargin() (msWrapMargin interface{}) {
	js.Rewrite("$<.MsWrapMargin")
	return msWrapMargin
}

// MsWrapMargin
func (*CSSStyleDeclaration) SetMsWrapMargin(msWrapMargin interface{}) {
	js.Rewrite("$<.MsWrapMargin = $1", msWrapMargin)
}

// MsWrapThrough
func (*CSSStyleDeclaration) MsWrapThrough() (msWrapThrough string) {
	js.Rewrite("$<.MsWrapThrough")
	return msWrapThrough
}

// MsWrapThrough
func (*CSSStyleDeclaration) SetMsWrapThrough(msWrapThrough string) {
	js.Rewrite("$<.MsWrapThrough = $1", msWrapThrough)
}

// Opacity
func (*CSSStyleDeclaration) Opacity() (opacity string) {
	js.Rewrite("$<.Opacity")
	return opacity
}

// Opacity
func (*CSSStyleDeclaration) SetOpacity(opacity string) {
	js.Rewrite("$<.Opacity = $1", opacity)
}

// Order
func (*CSSStyleDeclaration) Order() (order string) {
	js.Rewrite("$<.Order")
	return order
}

// Order
func (*CSSStyleDeclaration) SetOrder(order string) {
	js.Rewrite("$<.Order = $1", order)
}

// Orphans
func (*CSSStyleDeclaration) Orphans() (orphans string) {
	js.Rewrite("$<.Orphans")
	return orphans
}

// Orphans
func (*CSSStyleDeclaration) SetOrphans(orphans string) {
	js.Rewrite("$<.Orphans = $1", orphans)
}

// Outline
func (*CSSStyleDeclaration) Outline() (outline string) {
	js.Rewrite("$<.Outline")
	return outline
}

// Outline
func (*CSSStyleDeclaration) SetOutline(outline string) {
	js.Rewrite("$<.Outline = $1", outline)
}

// OutlineColor
func (*CSSStyleDeclaration) OutlineColor() (outlineColor string) {
	js.Rewrite("$<.OutlineColor")
	return outlineColor
}

// OutlineColor
func (*CSSStyleDeclaration) SetOutlineColor(outlineColor string) {
	js.Rewrite("$<.OutlineColor = $1", outlineColor)
}

// OutlineOffset
func (*CSSStyleDeclaration) OutlineOffset() (outlineOffset string) {
	js.Rewrite("$<.OutlineOffset")
	return outlineOffset
}

// OutlineOffset
func (*CSSStyleDeclaration) SetOutlineOffset(outlineOffset string) {
	js.Rewrite("$<.OutlineOffset = $1", outlineOffset)
}

// OutlineStyle
func (*CSSStyleDeclaration) OutlineStyle() (outlineStyle string) {
	js.Rewrite("$<.OutlineStyle")
	return outlineStyle
}

// OutlineStyle
func (*CSSStyleDeclaration) SetOutlineStyle(outlineStyle string) {
	js.Rewrite("$<.OutlineStyle = $1", outlineStyle)
}

// OutlineWidth
func (*CSSStyleDeclaration) OutlineWidth() (outlineWidth string) {
	js.Rewrite("$<.OutlineWidth")
	return outlineWidth
}

// OutlineWidth
func (*CSSStyleDeclaration) SetOutlineWidth(outlineWidth string) {
	js.Rewrite("$<.OutlineWidth = $1", outlineWidth)
}

// Overflow
func (*CSSStyleDeclaration) Overflow() (overflow string) {
	js.Rewrite("$<.Overflow")
	return overflow
}

// Overflow
func (*CSSStyleDeclaration) SetOverflow(overflow string) {
	js.Rewrite("$<.Overflow = $1", overflow)
}

// OverflowX
func (*CSSStyleDeclaration) OverflowX() (overflowX string) {
	js.Rewrite("$<.OverflowX")
	return overflowX
}

// OverflowX
func (*CSSStyleDeclaration) SetOverflowX(overflowX string) {
	js.Rewrite("$<.OverflowX = $1", overflowX)
}

// OverflowY
func (*CSSStyleDeclaration) OverflowY() (overflowY string) {
	js.Rewrite("$<.OverflowY")
	return overflowY
}

// OverflowY
func (*CSSStyleDeclaration) SetOverflowY(overflowY string) {
	js.Rewrite("$<.OverflowY = $1", overflowY)
}

// Padding
func (*CSSStyleDeclaration) Padding() (padding string) {
	js.Rewrite("$<.Padding")
	return padding
}

// Padding
func (*CSSStyleDeclaration) SetPadding(padding string) {
	js.Rewrite("$<.Padding = $1", padding)
}

// PaddingBottom
func (*CSSStyleDeclaration) PaddingBottom() (paddingBottom string) {
	js.Rewrite("$<.PaddingBottom")
	return paddingBottom
}

// PaddingBottom
func (*CSSStyleDeclaration) SetPaddingBottom(paddingBottom string) {
	js.Rewrite("$<.PaddingBottom = $1", paddingBottom)
}

// PaddingLeft
func (*CSSStyleDeclaration) PaddingLeft() (paddingLeft string) {
	js.Rewrite("$<.PaddingLeft")
	return paddingLeft
}

// PaddingLeft
func (*CSSStyleDeclaration) SetPaddingLeft(paddingLeft string) {
	js.Rewrite("$<.PaddingLeft = $1", paddingLeft)
}

// PaddingRight
func (*CSSStyleDeclaration) PaddingRight() (paddingRight string) {
	js.Rewrite("$<.PaddingRight")
	return paddingRight
}

// PaddingRight
func (*CSSStyleDeclaration) SetPaddingRight(paddingRight string) {
	js.Rewrite("$<.PaddingRight = $1", paddingRight)
}

// PaddingTop
func (*CSSStyleDeclaration) PaddingTop() (paddingTop string) {
	js.Rewrite("$<.PaddingTop")
	return paddingTop
}

// PaddingTop
func (*CSSStyleDeclaration) SetPaddingTop(paddingTop string) {
	js.Rewrite("$<.PaddingTop = $1", paddingTop)
}

// PageBreakAfter
func (*CSSStyleDeclaration) PageBreakAfter() (pageBreakAfter string) {
	js.Rewrite("$<.PageBreakAfter")
	return pageBreakAfter
}

// PageBreakAfter
func (*CSSStyleDeclaration) SetPageBreakAfter(pageBreakAfter string) {
	js.Rewrite("$<.PageBreakAfter = $1", pageBreakAfter)
}

// PageBreakBefore
func (*CSSStyleDeclaration) PageBreakBefore() (pageBreakBefore string) {
	js.Rewrite("$<.PageBreakBefore")
	return pageBreakBefore
}

// PageBreakBefore
func (*CSSStyleDeclaration) SetPageBreakBefore(pageBreakBefore string) {
	js.Rewrite("$<.PageBreakBefore = $1", pageBreakBefore)
}

// PageBreakInside
func (*CSSStyleDeclaration) PageBreakInside() (pageBreakInside string) {
	js.Rewrite("$<.PageBreakInside")
	return pageBreakInside
}

// PageBreakInside
func (*CSSStyleDeclaration) SetPageBreakInside(pageBreakInside string) {
	js.Rewrite("$<.PageBreakInside = $1", pageBreakInside)
}

// ParentRule
func (*CSSStyleDeclaration) ParentRule() (parentRule CSSRule) {
	js.Rewrite("$<.ParentRule")
	return parentRule
}

// Perspective
func (*CSSStyleDeclaration) Perspective() (perspective string) {
	js.Rewrite("$<.Perspective")
	return perspective
}

// Perspective
func (*CSSStyleDeclaration) SetPerspective(perspective string) {
	js.Rewrite("$<.Perspective = $1", perspective)
}

// PerspectiveOrigin
func (*CSSStyleDeclaration) PerspectiveOrigin() (perspectiveOrigin string) {
	js.Rewrite("$<.PerspectiveOrigin")
	return perspectiveOrigin
}

// PerspectiveOrigin
func (*CSSStyleDeclaration) SetPerspectiveOrigin(perspectiveOrigin string) {
	js.Rewrite("$<.PerspectiveOrigin = $1", perspectiveOrigin)
}

// PointerEvents
func (*CSSStyleDeclaration) PointerEvents() (pointerEvents string) {
	js.Rewrite("$<.PointerEvents")
	return pointerEvents
}

// PointerEvents
func (*CSSStyleDeclaration) SetPointerEvents(pointerEvents string) {
	js.Rewrite("$<.PointerEvents = $1", pointerEvents)
}

// Position
func (*CSSStyleDeclaration) Position() (position string) {
	js.Rewrite("$<.Position")
	return position
}

// Position
func (*CSSStyleDeclaration) SetPosition(position string) {
	js.Rewrite("$<.Position = $1", position)
}

// Quotes
func (*CSSStyleDeclaration) Quotes() (quotes string) {
	js.Rewrite("$<.Quotes")
	return quotes
}

// Quotes
func (*CSSStyleDeclaration) SetQuotes(quotes string) {
	js.Rewrite("$<.Quotes = $1", quotes)
}

// Right
func (*CSSStyleDeclaration) Right() (right string) {
	js.Rewrite("$<.Right")
	return right
}

// Right
func (*CSSStyleDeclaration) SetRight(right string) {
	js.Rewrite("$<.Right = $1", right)
}

// Rotate
func (*CSSStyleDeclaration) Rotate() (rotate string) {
	js.Rewrite("$<.Rotate")
	return rotate
}

// Rotate
func (*CSSStyleDeclaration) SetRotate(rotate string) {
	js.Rewrite("$<.Rotate = $1", rotate)
}

// RubyAlign
func (*CSSStyleDeclaration) RubyAlign() (rubyAlign string) {
	js.Rewrite("$<.RubyAlign")
	return rubyAlign
}

// RubyAlign
func (*CSSStyleDeclaration) SetRubyAlign(rubyAlign string) {
	js.Rewrite("$<.RubyAlign = $1", rubyAlign)
}

// RubyOverhang
func (*CSSStyleDeclaration) RubyOverhang() (rubyOverhang string) {
	js.Rewrite("$<.RubyOverhang")
	return rubyOverhang
}

// RubyOverhang
func (*CSSStyleDeclaration) SetRubyOverhang(rubyOverhang string) {
	js.Rewrite("$<.RubyOverhang = $1", rubyOverhang)
}

// RubyPosition
func (*CSSStyleDeclaration) RubyPosition() (rubyPosition string) {
	js.Rewrite("$<.RubyPosition")
	return rubyPosition
}

// RubyPosition
func (*CSSStyleDeclaration) SetRubyPosition(rubyPosition string) {
	js.Rewrite("$<.RubyPosition = $1", rubyPosition)
}

// Scale
func (*CSSStyleDeclaration) Scale() (scale string) {
	js.Rewrite("$<.Scale")
	return scale
}

// Scale
func (*CSSStyleDeclaration) SetScale(scale string) {
	js.Rewrite("$<.Scale = $1", scale)
}

// StopColor
func (*CSSStyleDeclaration) StopColor() (stopColor string) {
	js.Rewrite("$<.StopColor")
	return stopColor
}

// StopColor
func (*CSSStyleDeclaration) SetStopColor(stopColor string) {
	js.Rewrite("$<.StopColor = $1", stopColor)
}

// StopOpacity
func (*CSSStyleDeclaration) StopOpacity() (stopOpacity string) {
	js.Rewrite("$<.StopOpacity")
	return stopOpacity
}

// StopOpacity
func (*CSSStyleDeclaration) SetStopOpacity(stopOpacity string) {
	js.Rewrite("$<.StopOpacity = $1", stopOpacity)
}

// Stroke
func (*CSSStyleDeclaration) Stroke() (stroke string) {
	js.Rewrite("$<.Stroke")
	return stroke
}

// Stroke
func (*CSSStyleDeclaration) SetStroke(stroke string) {
	js.Rewrite("$<.Stroke = $1", stroke)
}

// StrokeDasharray
func (*CSSStyleDeclaration) StrokeDasharray() (strokeDasharray string) {
	js.Rewrite("$<.StrokeDasharray")
	return strokeDasharray
}

// StrokeDasharray
func (*CSSStyleDeclaration) SetStrokeDasharray(strokeDasharray string) {
	js.Rewrite("$<.StrokeDasharray = $1", strokeDasharray)
}

// StrokeDashoffset
func (*CSSStyleDeclaration) StrokeDashoffset() (strokeDashoffset string) {
	js.Rewrite("$<.StrokeDashoffset")
	return strokeDashoffset
}

// StrokeDashoffset
func (*CSSStyleDeclaration) SetStrokeDashoffset(strokeDashoffset string) {
	js.Rewrite("$<.StrokeDashoffset = $1", strokeDashoffset)
}

// StrokeLinecap
func (*CSSStyleDeclaration) StrokeLinecap() (strokeLinecap string) {
	js.Rewrite("$<.StrokeLinecap")
	return strokeLinecap
}

// StrokeLinecap
func (*CSSStyleDeclaration) SetStrokeLinecap(strokeLinecap string) {
	js.Rewrite("$<.StrokeLinecap = $1", strokeLinecap)
}

// StrokeLinejoin
func (*CSSStyleDeclaration) StrokeLinejoin() (strokeLinejoin string) {
	js.Rewrite("$<.StrokeLinejoin")
	return strokeLinejoin
}

// StrokeLinejoin
func (*CSSStyleDeclaration) SetStrokeLinejoin(strokeLinejoin string) {
	js.Rewrite("$<.StrokeLinejoin = $1", strokeLinejoin)
}

// StrokeMiterlimit
func (*CSSStyleDeclaration) StrokeMiterlimit() (strokeMiterlimit string) {
	js.Rewrite("$<.StrokeMiterlimit")
	return strokeMiterlimit
}

// StrokeMiterlimit
func (*CSSStyleDeclaration) SetStrokeMiterlimit(strokeMiterlimit string) {
	js.Rewrite("$<.StrokeMiterlimit = $1", strokeMiterlimit)
}

// StrokeOpacity
func (*CSSStyleDeclaration) StrokeOpacity() (strokeOpacity string) {
	js.Rewrite("$<.StrokeOpacity")
	return strokeOpacity
}

// StrokeOpacity
func (*CSSStyleDeclaration) SetStrokeOpacity(strokeOpacity string) {
	js.Rewrite("$<.StrokeOpacity = $1", strokeOpacity)
}

// StrokeWidth
func (*CSSStyleDeclaration) StrokeWidth() (strokeWidth string) {
	js.Rewrite("$<.StrokeWidth")
	return strokeWidth
}

// StrokeWidth
func (*CSSStyleDeclaration) SetStrokeWidth(strokeWidth string) {
	js.Rewrite("$<.StrokeWidth = $1", strokeWidth)
}

// TableLayout
func (*CSSStyleDeclaration) TableLayout() (tableLayout string) {
	js.Rewrite("$<.TableLayout")
	return tableLayout
}

// TableLayout
func (*CSSStyleDeclaration) SetTableLayout(tableLayout string) {
	js.Rewrite("$<.TableLayout = $1", tableLayout)
}

// TextAlign
func (*CSSStyleDeclaration) TextAlign() (textAlign string) {
	js.Rewrite("$<.TextAlign")
	return textAlign
}

// TextAlign
func (*CSSStyleDeclaration) SetTextAlign(textAlign string) {
	js.Rewrite("$<.TextAlign = $1", textAlign)
}

// TextAlignLast
func (*CSSStyleDeclaration) TextAlignLast() (textAlignLast string) {
	js.Rewrite("$<.TextAlignLast")
	return textAlignLast
}

// TextAlignLast
func (*CSSStyleDeclaration) SetTextAlignLast(textAlignLast string) {
	js.Rewrite("$<.TextAlignLast = $1", textAlignLast)
}

// TextAnchor
func (*CSSStyleDeclaration) TextAnchor() (textAnchor string) {
	js.Rewrite("$<.TextAnchor")
	return textAnchor
}

// TextAnchor
func (*CSSStyleDeclaration) SetTextAnchor(textAnchor string) {
	js.Rewrite("$<.TextAnchor = $1", textAnchor)
}

// TextDecoration
func (*CSSStyleDeclaration) TextDecoration() (textDecoration string) {
	js.Rewrite("$<.TextDecoration")
	return textDecoration
}

// TextDecoration
func (*CSSStyleDeclaration) SetTextDecoration(textDecoration string) {
	js.Rewrite("$<.TextDecoration = $1", textDecoration)
}

// TextIndent
func (*CSSStyleDeclaration) TextIndent() (textIndent string) {
	js.Rewrite("$<.TextIndent")
	return textIndent
}

// TextIndent
func (*CSSStyleDeclaration) SetTextIndent(textIndent string) {
	js.Rewrite("$<.TextIndent = $1", textIndent)
}

// TextJustify
func (*CSSStyleDeclaration) TextJustify() (textJustify string) {
	js.Rewrite("$<.TextJustify")
	return textJustify
}

// TextJustify
func (*CSSStyleDeclaration) SetTextJustify(textJustify string) {
	js.Rewrite("$<.TextJustify = $1", textJustify)
}

// TextKashida
func (*CSSStyleDeclaration) TextKashida() (textKashida string) {
	js.Rewrite("$<.TextKashida")
	return textKashida
}

// TextKashida
func (*CSSStyleDeclaration) SetTextKashida(textKashida string) {
	js.Rewrite("$<.TextKashida = $1", textKashida)
}

// TextKashidaSpace
func (*CSSStyleDeclaration) TextKashidaSpace() (textKashidaSpace string) {
	js.Rewrite("$<.TextKashidaSpace")
	return textKashidaSpace
}

// TextKashidaSpace
func (*CSSStyleDeclaration) SetTextKashidaSpace(textKashidaSpace string) {
	js.Rewrite("$<.TextKashidaSpace = $1", textKashidaSpace)
}

// TextOverflow
func (*CSSStyleDeclaration) TextOverflow() (textOverflow string) {
	js.Rewrite("$<.TextOverflow")
	return textOverflow
}

// TextOverflow
func (*CSSStyleDeclaration) SetTextOverflow(textOverflow string) {
	js.Rewrite("$<.TextOverflow = $1", textOverflow)
}

// TextShadow
func (*CSSStyleDeclaration) TextShadow() (textShadow string) {
	js.Rewrite("$<.TextShadow")
	return textShadow
}

// TextShadow
func (*CSSStyleDeclaration) SetTextShadow(textShadow string) {
	js.Rewrite("$<.TextShadow = $1", textShadow)
}

// TextTransform
func (*CSSStyleDeclaration) TextTransform() (textTransform string) {
	js.Rewrite("$<.TextTransform")
	return textTransform
}

// TextTransform
func (*CSSStyleDeclaration) SetTextTransform(textTransform string) {
	js.Rewrite("$<.TextTransform = $1", textTransform)
}

// TextUnderlinePosition
func (*CSSStyleDeclaration) TextUnderlinePosition() (textUnderlinePosition string) {
	js.Rewrite("$<.TextUnderlinePosition")
	return textUnderlinePosition
}

// TextUnderlinePosition
func (*CSSStyleDeclaration) SetTextUnderlinePosition(textUnderlinePosition string) {
	js.Rewrite("$<.TextUnderlinePosition = $1", textUnderlinePosition)
}

// Top
func (*CSSStyleDeclaration) Top() (top string) {
	js.Rewrite("$<.Top")
	return top
}

// Top
func (*CSSStyleDeclaration) SetTop(top string) {
	js.Rewrite("$<.Top = $1", top)
}

// TouchAction
func (*CSSStyleDeclaration) TouchAction() (touchAction string) {
	js.Rewrite("$<.TouchAction")
	return touchAction
}

// TouchAction
func (*CSSStyleDeclaration) SetTouchAction(touchAction string) {
	js.Rewrite("$<.TouchAction = $1", touchAction)
}

// Transform
func (*CSSStyleDeclaration) Transform() (transform string) {
	js.Rewrite("$<.Transform")
	return transform
}

// Transform
func (*CSSStyleDeclaration) SetTransform(transform string) {
	js.Rewrite("$<.Transform = $1", transform)
}

// TransformOrigin
func (*CSSStyleDeclaration) TransformOrigin() (transformOrigin string) {
	js.Rewrite("$<.TransformOrigin")
	return transformOrigin
}

// TransformOrigin
func (*CSSStyleDeclaration) SetTransformOrigin(transformOrigin string) {
	js.Rewrite("$<.TransformOrigin = $1", transformOrigin)
}

// TransformStyle
func (*CSSStyleDeclaration) TransformStyle() (transformStyle string) {
	js.Rewrite("$<.TransformStyle")
	return transformStyle
}

// TransformStyle
func (*CSSStyleDeclaration) SetTransformStyle(transformStyle string) {
	js.Rewrite("$<.TransformStyle = $1", transformStyle)
}

// Transition
func (*CSSStyleDeclaration) Transition() (transition string) {
	js.Rewrite("$<.Transition")
	return transition
}

// Transition
func (*CSSStyleDeclaration) SetTransition(transition string) {
	js.Rewrite("$<.Transition = $1", transition)
}

// TransitionDelay
func (*CSSStyleDeclaration) TransitionDelay() (transitionDelay string) {
	js.Rewrite("$<.TransitionDelay")
	return transitionDelay
}

// TransitionDelay
func (*CSSStyleDeclaration) SetTransitionDelay(transitionDelay string) {
	js.Rewrite("$<.TransitionDelay = $1", transitionDelay)
}

// TransitionDuration
func (*CSSStyleDeclaration) TransitionDuration() (transitionDuration string) {
	js.Rewrite("$<.TransitionDuration")
	return transitionDuration
}

// TransitionDuration
func (*CSSStyleDeclaration) SetTransitionDuration(transitionDuration string) {
	js.Rewrite("$<.TransitionDuration = $1", transitionDuration)
}

// TransitionProperty
func (*CSSStyleDeclaration) TransitionProperty() (transitionProperty string) {
	js.Rewrite("$<.TransitionProperty")
	return transitionProperty
}

// TransitionProperty
func (*CSSStyleDeclaration) SetTransitionProperty(transitionProperty string) {
	js.Rewrite("$<.TransitionProperty = $1", transitionProperty)
}

// TransitionTimingFunction
func (*CSSStyleDeclaration) TransitionTimingFunction() (transitionTimingFunction string) {
	js.Rewrite("$<.TransitionTimingFunction")
	return transitionTimingFunction
}

// TransitionTimingFunction
func (*CSSStyleDeclaration) SetTransitionTimingFunction(transitionTimingFunction string) {
	js.Rewrite("$<.TransitionTimingFunction = $1", transitionTimingFunction)
}

// Translate
func (*CSSStyleDeclaration) Translate() (translate string) {
	js.Rewrite("$<.Translate")
	return translate
}

// Translate
func (*CSSStyleDeclaration) SetTranslate(translate string) {
	js.Rewrite("$<.Translate = $1", translate)
}

// UnicodeBidi
func (*CSSStyleDeclaration) UnicodeBidi() (unicodeBidi string) {
	js.Rewrite("$<.UnicodeBidi")
	return unicodeBidi
}

// UnicodeBidi
func (*CSSStyleDeclaration) SetUnicodeBidi(unicodeBidi string) {
	js.Rewrite("$<.UnicodeBidi = $1", unicodeBidi)
}

// VerticalAlign
func (*CSSStyleDeclaration) VerticalAlign() (verticalAlign string) {
	js.Rewrite("$<.VerticalAlign")
	return verticalAlign
}

// VerticalAlign
func (*CSSStyleDeclaration) SetVerticalAlign(verticalAlign string) {
	js.Rewrite("$<.VerticalAlign = $1", verticalAlign)
}

// Visibility
func (*CSSStyleDeclaration) Visibility() (visibility string) {
	js.Rewrite("$<.Visibility")
	return visibility
}

// Visibility
func (*CSSStyleDeclaration) SetVisibility(visibility string) {
	js.Rewrite("$<.Visibility = $1", visibility)
}

// WebkitAlignContent
func (*CSSStyleDeclaration) WebkitAlignContent() (webkitAlignContent string) {
	js.Rewrite("$<.WebkitAlignContent")
	return webkitAlignContent
}

// WebkitAlignContent
func (*CSSStyleDeclaration) SetWebkitAlignContent(webkitAlignContent string) {
	js.Rewrite("$<.WebkitAlignContent = $1", webkitAlignContent)
}

// WebkitAlignItems
func (*CSSStyleDeclaration) WebkitAlignItems() (webkitAlignItems string) {
	js.Rewrite("$<.WebkitAlignItems")
	return webkitAlignItems
}

// WebkitAlignItems
func (*CSSStyleDeclaration) SetWebkitAlignItems(webkitAlignItems string) {
	js.Rewrite("$<.WebkitAlignItems = $1", webkitAlignItems)
}

// WebkitAlignSelf
func (*CSSStyleDeclaration) WebkitAlignSelf() (webkitAlignSelf string) {
	js.Rewrite("$<.WebkitAlignSelf")
	return webkitAlignSelf
}

// WebkitAlignSelf
func (*CSSStyleDeclaration) SetWebkitAlignSelf(webkitAlignSelf string) {
	js.Rewrite("$<.WebkitAlignSelf = $1", webkitAlignSelf)
}

// WebkitAnimation
func (*CSSStyleDeclaration) WebkitAnimation() (webkitAnimation string) {
	js.Rewrite("$<.WebkitAnimation")
	return webkitAnimation
}

// WebkitAnimation
func (*CSSStyleDeclaration) SetWebkitAnimation(webkitAnimation string) {
	js.Rewrite("$<.WebkitAnimation = $1", webkitAnimation)
}

// WebkitAnimationDelay
func (*CSSStyleDeclaration) WebkitAnimationDelay() (webkitAnimationDelay string) {
	js.Rewrite("$<.WebkitAnimationDelay")
	return webkitAnimationDelay
}

// WebkitAnimationDelay
func (*CSSStyleDeclaration) SetWebkitAnimationDelay(webkitAnimationDelay string) {
	js.Rewrite("$<.WebkitAnimationDelay = $1", webkitAnimationDelay)
}

// WebkitAnimationDirection
func (*CSSStyleDeclaration) WebkitAnimationDirection() (webkitAnimationDirection string) {
	js.Rewrite("$<.WebkitAnimationDirection")
	return webkitAnimationDirection
}

// WebkitAnimationDirection
func (*CSSStyleDeclaration) SetWebkitAnimationDirection(webkitAnimationDirection string) {
	js.Rewrite("$<.WebkitAnimationDirection = $1", webkitAnimationDirection)
}

// WebkitAnimationDuration
func (*CSSStyleDeclaration) WebkitAnimationDuration() (webkitAnimationDuration string) {
	js.Rewrite("$<.WebkitAnimationDuration")
	return webkitAnimationDuration
}

// WebkitAnimationDuration
func (*CSSStyleDeclaration) SetWebkitAnimationDuration(webkitAnimationDuration string) {
	js.Rewrite("$<.WebkitAnimationDuration = $1", webkitAnimationDuration)
}

// WebkitAnimationFillMode
func (*CSSStyleDeclaration) WebkitAnimationFillMode() (webkitAnimationFillMode string) {
	js.Rewrite("$<.WebkitAnimationFillMode")
	return webkitAnimationFillMode
}

// WebkitAnimationFillMode
func (*CSSStyleDeclaration) SetWebkitAnimationFillMode(webkitAnimationFillMode string) {
	js.Rewrite("$<.WebkitAnimationFillMode = $1", webkitAnimationFillMode)
}

// WebkitAnimationIterationCount
func (*CSSStyleDeclaration) WebkitAnimationIterationCount() (webkitAnimationIterationCount string) {
	js.Rewrite("$<.WebkitAnimationIterationCount")
	return webkitAnimationIterationCount
}

// WebkitAnimationIterationCount
func (*CSSStyleDeclaration) SetWebkitAnimationIterationCount(webkitAnimationIterationCount string) {
	js.Rewrite("$<.WebkitAnimationIterationCount = $1", webkitAnimationIterationCount)
}

// WebkitAnimationName
func (*CSSStyleDeclaration) WebkitAnimationName() (webkitAnimationName string) {
	js.Rewrite("$<.WebkitAnimationName")
	return webkitAnimationName
}

// WebkitAnimationName
func (*CSSStyleDeclaration) SetWebkitAnimationName(webkitAnimationName string) {
	js.Rewrite("$<.WebkitAnimationName = $1", webkitAnimationName)
}

// WebkitAnimationPlayState
func (*CSSStyleDeclaration) WebkitAnimationPlayState() (webkitAnimationPlayState string) {
	js.Rewrite("$<.WebkitAnimationPlayState")
	return webkitAnimationPlayState
}

// WebkitAnimationPlayState
func (*CSSStyleDeclaration) SetWebkitAnimationPlayState(webkitAnimationPlayState string) {
	js.Rewrite("$<.WebkitAnimationPlayState = $1", webkitAnimationPlayState)
}

// WebkitAnimationTimingFunction
func (*CSSStyleDeclaration) WebkitAnimationTimingFunction() (webkitAnimationTimingFunction string) {
	js.Rewrite("$<.WebkitAnimationTimingFunction")
	return webkitAnimationTimingFunction
}

// WebkitAnimationTimingFunction
func (*CSSStyleDeclaration) SetWebkitAnimationTimingFunction(webkitAnimationTimingFunction string) {
	js.Rewrite("$<.WebkitAnimationTimingFunction = $1", webkitAnimationTimingFunction)
}

// WebkitAppearance
func (*CSSStyleDeclaration) WebkitAppearance() (webkitAppearance string) {
	js.Rewrite("$<.WebkitAppearance")
	return webkitAppearance
}

// WebkitAppearance
func (*CSSStyleDeclaration) SetWebkitAppearance(webkitAppearance string) {
	js.Rewrite("$<.WebkitAppearance = $1", webkitAppearance)
}

// WebkitBackfaceVisibility
func (*CSSStyleDeclaration) WebkitBackfaceVisibility() (webkitBackfaceVisibility string) {
	js.Rewrite("$<.WebkitBackfaceVisibility")
	return webkitBackfaceVisibility
}

// WebkitBackfaceVisibility
func (*CSSStyleDeclaration) SetWebkitBackfaceVisibility(webkitBackfaceVisibility string) {
	js.Rewrite("$<.WebkitBackfaceVisibility = $1", webkitBackfaceVisibility)
}

// WebkitBackgroundClip
func (*CSSStyleDeclaration) WebkitBackgroundClip() (webkitBackgroundClip string) {
	js.Rewrite("$<.WebkitBackgroundClip")
	return webkitBackgroundClip
}

// WebkitBackgroundClip
func (*CSSStyleDeclaration) SetWebkitBackgroundClip(webkitBackgroundClip string) {
	js.Rewrite("$<.WebkitBackgroundClip = $1", webkitBackgroundClip)
}

// WebkitBackgroundOrigin
func (*CSSStyleDeclaration) WebkitBackgroundOrigin() (webkitBackgroundOrigin string) {
	js.Rewrite("$<.WebkitBackgroundOrigin")
	return webkitBackgroundOrigin
}

// WebkitBackgroundOrigin
func (*CSSStyleDeclaration) SetWebkitBackgroundOrigin(webkitBackgroundOrigin string) {
	js.Rewrite("$<.WebkitBackgroundOrigin = $1", webkitBackgroundOrigin)
}

// WebkitBackgroundSize
func (*CSSStyleDeclaration) WebkitBackgroundSize() (webkitBackgroundSize string) {
	js.Rewrite("$<.WebkitBackgroundSize")
	return webkitBackgroundSize
}

// WebkitBackgroundSize
func (*CSSStyleDeclaration) SetWebkitBackgroundSize(webkitBackgroundSize string) {
	js.Rewrite("$<.WebkitBackgroundSize = $1", webkitBackgroundSize)
}

// WebkitBorderBottomLeftRadius
func (*CSSStyleDeclaration) WebkitBorderBottomLeftRadius() (webkitBorderBottomLeftRadius string) {
	js.Rewrite("$<.WebkitBorderBottomLeftRadius")
	return webkitBorderBottomLeftRadius
}

// WebkitBorderBottomLeftRadius
func (*CSSStyleDeclaration) SetWebkitBorderBottomLeftRadius(webkitBorderBottomLeftRadius string) {
	js.Rewrite("$<.WebkitBorderBottomLeftRadius = $1", webkitBorderBottomLeftRadius)
}

// WebkitBorderBottomRightRadius
func (*CSSStyleDeclaration) WebkitBorderBottomRightRadius() (webkitBorderBottomRightRadius string) {
	js.Rewrite("$<.WebkitBorderBottomRightRadius")
	return webkitBorderBottomRightRadius
}

// WebkitBorderBottomRightRadius
func (*CSSStyleDeclaration) SetWebkitBorderBottomRightRadius(webkitBorderBottomRightRadius string) {
	js.Rewrite("$<.WebkitBorderBottomRightRadius = $1", webkitBorderBottomRightRadius)
}

// WebkitBorderImage
func (*CSSStyleDeclaration) WebkitBorderImage() (webkitBorderImage string) {
	js.Rewrite("$<.WebkitBorderImage")
	return webkitBorderImage
}

// WebkitBorderImage
func (*CSSStyleDeclaration) SetWebkitBorderImage(webkitBorderImage string) {
	js.Rewrite("$<.WebkitBorderImage = $1", webkitBorderImage)
}

// WebkitBorderRadius
func (*CSSStyleDeclaration) WebkitBorderRadius() (webkitBorderRadius string) {
	js.Rewrite("$<.WebkitBorderRadius")
	return webkitBorderRadius
}

// WebkitBorderRadius
func (*CSSStyleDeclaration) SetWebkitBorderRadius(webkitBorderRadius string) {
	js.Rewrite("$<.WebkitBorderRadius = $1", webkitBorderRadius)
}

// WebkitBorderTopLeftRadius
func (*CSSStyleDeclaration) WebkitBorderTopLeftRadius() (webkitBorderTopLeftRadius string) {
	js.Rewrite("$<.WebkitBorderTopLeftRadius")
	return webkitBorderTopLeftRadius
}

// WebkitBorderTopLeftRadius
func (*CSSStyleDeclaration) SetWebkitBorderTopLeftRadius(webkitBorderTopLeftRadius string) {
	js.Rewrite("$<.WebkitBorderTopLeftRadius = $1", webkitBorderTopLeftRadius)
}

// WebkitBorderTopRightRadius
func (*CSSStyleDeclaration) WebkitBorderTopRightRadius() (webkitBorderTopRightRadius string) {
	js.Rewrite("$<.WebkitBorderTopRightRadius")
	return webkitBorderTopRightRadius
}

// WebkitBorderTopRightRadius
func (*CSSStyleDeclaration) SetWebkitBorderTopRightRadius(webkitBorderTopRightRadius string) {
	js.Rewrite("$<.WebkitBorderTopRightRadius = $1", webkitBorderTopRightRadius)
}

// WebkitBoxAlign
func (*CSSStyleDeclaration) WebkitBoxAlign() (webkitBoxAlign string) {
	js.Rewrite("$<.WebkitBoxAlign")
	return webkitBoxAlign
}

// WebkitBoxAlign
func (*CSSStyleDeclaration) SetWebkitBoxAlign(webkitBoxAlign string) {
	js.Rewrite("$<.WebkitBoxAlign = $1", webkitBoxAlign)
}

// WebkitBoxDirection
func (*CSSStyleDeclaration) WebkitBoxDirection() (webkitBoxDirection string) {
	js.Rewrite("$<.WebkitBoxDirection")
	return webkitBoxDirection
}

// WebkitBoxDirection
func (*CSSStyleDeclaration) SetWebkitBoxDirection(webkitBoxDirection string) {
	js.Rewrite("$<.WebkitBoxDirection = $1", webkitBoxDirection)
}

// WebkitBoxFlex
func (*CSSStyleDeclaration) WebkitBoxFlex() (webkitBoxFlex string) {
	js.Rewrite("$<.WebkitBoxFlex")
	return webkitBoxFlex
}

// WebkitBoxFlex
func (*CSSStyleDeclaration) SetWebkitBoxFlex(webkitBoxFlex string) {
	js.Rewrite("$<.WebkitBoxFlex = $1", webkitBoxFlex)
}

// WebkitBoxOrdinalGroup
func (*CSSStyleDeclaration) WebkitBoxOrdinalGroup() (webkitBoxOrdinalGroup string) {
	js.Rewrite("$<.WebkitBoxOrdinalGroup")
	return webkitBoxOrdinalGroup
}

// WebkitBoxOrdinalGroup
func (*CSSStyleDeclaration) SetWebkitBoxOrdinalGroup(webkitBoxOrdinalGroup string) {
	js.Rewrite("$<.WebkitBoxOrdinalGroup = $1", webkitBoxOrdinalGroup)
}

// WebkitBoxOrient
func (*CSSStyleDeclaration) WebkitBoxOrient() (webkitBoxOrient string) {
	js.Rewrite("$<.WebkitBoxOrient")
	return webkitBoxOrient
}

// WebkitBoxOrient
func (*CSSStyleDeclaration) SetWebkitBoxOrient(webkitBoxOrient string) {
	js.Rewrite("$<.WebkitBoxOrient = $1", webkitBoxOrient)
}

// WebkitBoxPack
func (*CSSStyleDeclaration) WebkitBoxPack() (webkitBoxPack string) {
	js.Rewrite("$<.WebkitBoxPack")
	return webkitBoxPack
}

// WebkitBoxPack
func (*CSSStyleDeclaration) SetWebkitBoxPack(webkitBoxPack string) {
	js.Rewrite("$<.WebkitBoxPack = $1", webkitBoxPack)
}

// WebkitBoxSizing
func (*CSSStyleDeclaration) WebkitBoxSizing() (webkitBoxSizing string) {
	js.Rewrite("$<.WebkitBoxSizing")
	return webkitBoxSizing
}

// WebkitBoxSizing
func (*CSSStyleDeclaration) SetWebkitBoxSizing(webkitBoxSizing string) {
	js.Rewrite("$<.WebkitBoxSizing = $1", webkitBoxSizing)
}

// WebkitColumnBreakAfter
func (*CSSStyleDeclaration) WebkitColumnBreakAfter() (webkitColumnBreakAfter string) {
	js.Rewrite("$<.WebkitColumnBreakAfter")
	return webkitColumnBreakAfter
}

// WebkitColumnBreakAfter
func (*CSSStyleDeclaration) SetWebkitColumnBreakAfter(webkitColumnBreakAfter string) {
	js.Rewrite("$<.WebkitColumnBreakAfter = $1", webkitColumnBreakAfter)
}

// WebkitColumnBreakBefore
func (*CSSStyleDeclaration) WebkitColumnBreakBefore() (webkitColumnBreakBefore string) {
	js.Rewrite("$<.WebkitColumnBreakBefore")
	return webkitColumnBreakBefore
}

// WebkitColumnBreakBefore
func (*CSSStyleDeclaration) SetWebkitColumnBreakBefore(webkitColumnBreakBefore string) {
	js.Rewrite("$<.WebkitColumnBreakBefore = $1", webkitColumnBreakBefore)
}

// WebkitColumnBreakInside
func (*CSSStyleDeclaration) WebkitColumnBreakInside() (webkitColumnBreakInside string) {
	js.Rewrite("$<.WebkitColumnBreakInside")
	return webkitColumnBreakInside
}

// WebkitColumnBreakInside
func (*CSSStyleDeclaration) SetWebkitColumnBreakInside(webkitColumnBreakInside string) {
	js.Rewrite("$<.WebkitColumnBreakInside = $1", webkitColumnBreakInside)
}

// WebkitColumnCount
func (*CSSStyleDeclaration) WebkitColumnCount() (webkitColumnCount interface{}) {
	js.Rewrite("$<.WebkitColumnCount")
	return webkitColumnCount
}

// WebkitColumnCount
func (*CSSStyleDeclaration) SetWebkitColumnCount(webkitColumnCount interface{}) {
	js.Rewrite("$<.WebkitColumnCount = $1", webkitColumnCount)
}

// WebkitColumnGap
func (*CSSStyleDeclaration) WebkitColumnGap() (webkitColumnGap interface{}) {
	js.Rewrite("$<.WebkitColumnGap")
	return webkitColumnGap
}

// WebkitColumnGap
func (*CSSStyleDeclaration) SetWebkitColumnGap(webkitColumnGap interface{}) {
	js.Rewrite("$<.WebkitColumnGap = $1", webkitColumnGap)
}

// WebkitColumnRule
func (*CSSStyleDeclaration) WebkitColumnRule() (webkitColumnRule string) {
	js.Rewrite("$<.WebkitColumnRule")
	return webkitColumnRule
}

// WebkitColumnRule
func (*CSSStyleDeclaration) SetWebkitColumnRule(webkitColumnRule string) {
	js.Rewrite("$<.WebkitColumnRule = $1", webkitColumnRule)
}

// WebkitColumnRuleColor
func (*CSSStyleDeclaration) WebkitColumnRuleColor() (webkitColumnRuleColor interface{}) {
	js.Rewrite("$<.WebkitColumnRuleColor")
	return webkitColumnRuleColor
}

// WebkitColumnRuleColor
func (*CSSStyleDeclaration) SetWebkitColumnRuleColor(webkitColumnRuleColor interface{}) {
	js.Rewrite("$<.WebkitColumnRuleColor = $1", webkitColumnRuleColor)
}

// WebkitColumnRuleStyle
func (*CSSStyleDeclaration) WebkitColumnRuleStyle() (webkitColumnRuleStyle string) {
	js.Rewrite("$<.WebkitColumnRuleStyle")
	return webkitColumnRuleStyle
}

// WebkitColumnRuleStyle
func (*CSSStyleDeclaration) SetWebkitColumnRuleStyle(webkitColumnRuleStyle string) {
	js.Rewrite("$<.WebkitColumnRuleStyle = $1", webkitColumnRuleStyle)
}

// WebkitColumnRuleWidth
func (*CSSStyleDeclaration) WebkitColumnRuleWidth() (webkitColumnRuleWidth interface{}) {
	js.Rewrite("$<.WebkitColumnRuleWidth")
	return webkitColumnRuleWidth
}

// WebkitColumnRuleWidth
func (*CSSStyleDeclaration) SetWebkitColumnRuleWidth(webkitColumnRuleWidth interface{}) {
	js.Rewrite("$<.WebkitColumnRuleWidth = $1", webkitColumnRuleWidth)
}

// WebkitColumns
func (*CSSStyleDeclaration) WebkitColumns() (webkitColumns string) {
	js.Rewrite("$<.WebkitColumns")
	return webkitColumns
}

// WebkitColumns
func (*CSSStyleDeclaration) SetWebkitColumns(webkitColumns string) {
	js.Rewrite("$<.WebkitColumns = $1", webkitColumns)
}

// WebkitColumnSpan
func (*CSSStyleDeclaration) WebkitColumnSpan() (webkitColumnSpan string) {
	js.Rewrite("$<.WebkitColumnSpan")
	return webkitColumnSpan
}

// WebkitColumnSpan
func (*CSSStyleDeclaration) SetWebkitColumnSpan(webkitColumnSpan string) {
	js.Rewrite("$<.WebkitColumnSpan = $1", webkitColumnSpan)
}

// WebkitColumnWidth
func (*CSSStyleDeclaration) WebkitColumnWidth() (webkitColumnWidth interface{}) {
	js.Rewrite("$<.WebkitColumnWidth")
	return webkitColumnWidth
}

// WebkitColumnWidth
func (*CSSStyleDeclaration) SetWebkitColumnWidth(webkitColumnWidth interface{}) {
	js.Rewrite("$<.WebkitColumnWidth = $1", webkitColumnWidth)
}

// WebkitFilter
func (*CSSStyleDeclaration) WebkitFilter() (webkitFilter string) {
	js.Rewrite("$<.WebkitFilter")
	return webkitFilter
}

// WebkitFilter
func (*CSSStyleDeclaration) SetWebkitFilter(webkitFilter string) {
	js.Rewrite("$<.WebkitFilter = $1", webkitFilter)
}

// WebkitFlex
func (*CSSStyleDeclaration) WebkitFlex() (webkitFlex string) {
	js.Rewrite("$<.WebkitFlex")
	return webkitFlex
}

// WebkitFlex
func (*CSSStyleDeclaration) SetWebkitFlex(webkitFlex string) {
	js.Rewrite("$<.WebkitFlex = $1", webkitFlex)
}

// WebkitFlexBasis
func (*CSSStyleDeclaration) WebkitFlexBasis() (webkitFlexBasis string) {
	js.Rewrite("$<.WebkitFlexBasis")
	return webkitFlexBasis
}

// WebkitFlexBasis
func (*CSSStyleDeclaration) SetWebkitFlexBasis(webkitFlexBasis string) {
	js.Rewrite("$<.WebkitFlexBasis = $1", webkitFlexBasis)
}

// WebkitFlexDirection
func (*CSSStyleDeclaration) WebkitFlexDirection() (webkitFlexDirection string) {
	js.Rewrite("$<.WebkitFlexDirection")
	return webkitFlexDirection
}

// WebkitFlexDirection
func (*CSSStyleDeclaration) SetWebkitFlexDirection(webkitFlexDirection string) {
	js.Rewrite("$<.WebkitFlexDirection = $1", webkitFlexDirection)
}

// WebkitFlexFlow
func (*CSSStyleDeclaration) WebkitFlexFlow() (webkitFlexFlow string) {
	js.Rewrite("$<.WebkitFlexFlow")
	return webkitFlexFlow
}

// WebkitFlexFlow
func (*CSSStyleDeclaration) SetWebkitFlexFlow(webkitFlexFlow string) {
	js.Rewrite("$<.WebkitFlexFlow = $1", webkitFlexFlow)
}

// WebkitFlexGrow
func (*CSSStyleDeclaration) WebkitFlexGrow() (webkitFlexGrow string) {
	js.Rewrite("$<.WebkitFlexGrow")
	return webkitFlexGrow
}

// WebkitFlexGrow
func (*CSSStyleDeclaration) SetWebkitFlexGrow(webkitFlexGrow string) {
	js.Rewrite("$<.WebkitFlexGrow = $1", webkitFlexGrow)
}

// WebkitFlexShrink
func (*CSSStyleDeclaration) WebkitFlexShrink() (webkitFlexShrink string) {
	js.Rewrite("$<.WebkitFlexShrink")
	return webkitFlexShrink
}

// WebkitFlexShrink
func (*CSSStyleDeclaration) SetWebkitFlexShrink(webkitFlexShrink string) {
	js.Rewrite("$<.WebkitFlexShrink = $1", webkitFlexShrink)
}

// WebkitFlexWrap
func (*CSSStyleDeclaration) WebkitFlexWrap() (webkitFlexWrap string) {
	js.Rewrite("$<.WebkitFlexWrap")
	return webkitFlexWrap
}

// WebkitFlexWrap
func (*CSSStyleDeclaration) SetWebkitFlexWrap(webkitFlexWrap string) {
	js.Rewrite("$<.WebkitFlexWrap = $1", webkitFlexWrap)
}

// WebkitJustifyContent
func (*CSSStyleDeclaration) WebkitJustifyContent() (webkitJustifyContent string) {
	js.Rewrite("$<.WebkitJustifyContent")
	return webkitJustifyContent
}

// WebkitJustifyContent
func (*CSSStyleDeclaration) SetWebkitJustifyContent(webkitJustifyContent string) {
	js.Rewrite("$<.WebkitJustifyContent = $1", webkitJustifyContent)
}

// WebkitOrder
func (*CSSStyleDeclaration) WebkitOrder() (webkitOrder string) {
	js.Rewrite("$<.WebkitOrder")
	return webkitOrder
}

// WebkitOrder
func (*CSSStyleDeclaration) SetWebkitOrder(webkitOrder string) {
	js.Rewrite("$<.WebkitOrder = $1", webkitOrder)
}

// WebkitPerspective
func (*CSSStyleDeclaration) WebkitPerspective() (webkitPerspective string) {
	js.Rewrite("$<.WebkitPerspective")
	return webkitPerspective
}

// WebkitPerspective
func (*CSSStyleDeclaration) SetWebkitPerspective(webkitPerspective string) {
	js.Rewrite("$<.WebkitPerspective = $1", webkitPerspective)
}

// WebkitPerspectiveOrigin
func (*CSSStyleDeclaration) WebkitPerspectiveOrigin() (webkitPerspectiveOrigin string) {
	js.Rewrite("$<.WebkitPerspectiveOrigin")
	return webkitPerspectiveOrigin
}

// WebkitPerspectiveOrigin
func (*CSSStyleDeclaration) SetWebkitPerspectiveOrigin(webkitPerspectiveOrigin string) {
	js.Rewrite("$<.WebkitPerspectiveOrigin = $1", webkitPerspectiveOrigin)
}

// WebkitTapHighlightColor
func (*CSSStyleDeclaration) WebkitTapHighlightColor() (webkitTapHighlightColor string) {
	js.Rewrite("$<.WebkitTapHighlightColor")
	return webkitTapHighlightColor
}

// WebkitTapHighlightColor
func (*CSSStyleDeclaration) SetWebkitTapHighlightColor(webkitTapHighlightColor string) {
	js.Rewrite("$<.WebkitTapHighlightColor = $1", webkitTapHighlightColor)
}

// WebkitTextFillColor
func (*CSSStyleDeclaration) WebkitTextFillColor() (webkitTextFillColor string) {
	js.Rewrite("$<.WebkitTextFillColor")
	return webkitTextFillColor
}

// WebkitTextFillColor
func (*CSSStyleDeclaration) SetWebkitTextFillColor(webkitTextFillColor string) {
	js.Rewrite("$<.WebkitTextFillColor = $1", webkitTextFillColor)
}

// WebkitTextSizeAdjust
func (*CSSStyleDeclaration) WebkitTextSizeAdjust() (webkitTextSizeAdjust interface{}) {
	js.Rewrite("$<.WebkitTextSizeAdjust")
	return webkitTextSizeAdjust
}

// WebkitTextSizeAdjust
func (*CSSStyleDeclaration) SetWebkitTextSizeAdjust(webkitTextSizeAdjust interface{}) {
	js.Rewrite("$<.WebkitTextSizeAdjust = $1", webkitTextSizeAdjust)
}

// WebkitTextStroke
func (*CSSStyleDeclaration) WebkitTextStroke() (webkitTextStroke string) {
	js.Rewrite("$<.WebkitTextStroke")
	return webkitTextStroke
}

// WebkitTextStroke
func (*CSSStyleDeclaration) SetWebkitTextStroke(webkitTextStroke string) {
	js.Rewrite("$<.WebkitTextStroke = $1", webkitTextStroke)
}

// WebkitTextStrokeColor
func (*CSSStyleDeclaration) WebkitTextStrokeColor() (webkitTextStrokeColor string) {
	js.Rewrite("$<.WebkitTextStrokeColor")
	return webkitTextStrokeColor
}

// WebkitTextStrokeColor
func (*CSSStyleDeclaration) SetWebkitTextStrokeColor(webkitTextStrokeColor string) {
	js.Rewrite("$<.WebkitTextStrokeColor = $1", webkitTextStrokeColor)
}

// WebkitTextStrokeWidth
func (*CSSStyleDeclaration) WebkitTextStrokeWidth() (webkitTextStrokeWidth string) {
	js.Rewrite("$<.WebkitTextStrokeWidth")
	return webkitTextStrokeWidth
}

// WebkitTextStrokeWidth
func (*CSSStyleDeclaration) SetWebkitTextStrokeWidth(webkitTextStrokeWidth string) {
	js.Rewrite("$<.WebkitTextStrokeWidth = $1", webkitTextStrokeWidth)
}

// WebkitTransform
func (*CSSStyleDeclaration) WebkitTransform() (webkitTransform string) {
	js.Rewrite("$<.WebkitTransform")
	return webkitTransform
}

// WebkitTransform
func (*CSSStyleDeclaration) SetWebkitTransform(webkitTransform string) {
	js.Rewrite("$<.WebkitTransform = $1", webkitTransform)
}

// WebkitTransformOrigin
func (*CSSStyleDeclaration) WebkitTransformOrigin() (webkitTransformOrigin string) {
	js.Rewrite("$<.WebkitTransformOrigin")
	return webkitTransformOrigin
}

// WebkitTransformOrigin
func (*CSSStyleDeclaration) SetWebkitTransformOrigin(webkitTransformOrigin string) {
	js.Rewrite("$<.WebkitTransformOrigin = $1", webkitTransformOrigin)
}

// WebkitTransformStyle
func (*CSSStyleDeclaration) WebkitTransformStyle() (webkitTransformStyle string) {
	js.Rewrite("$<.WebkitTransformStyle")
	return webkitTransformStyle
}

// WebkitTransformStyle
func (*CSSStyleDeclaration) SetWebkitTransformStyle(webkitTransformStyle string) {
	js.Rewrite("$<.WebkitTransformStyle = $1", webkitTransformStyle)
}

// WebkitTransition
func (*CSSStyleDeclaration) WebkitTransition() (webkitTransition string) {
	js.Rewrite("$<.WebkitTransition")
	return webkitTransition
}

// WebkitTransition
func (*CSSStyleDeclaration) SetWebkitTransition(webkitTransition string) {
	js.Rewrite("$<.WebkitTransition = $1", webkitTransition)
}

// WebkitTransitionDelay
func (*CSSStyleDeclaration) WebkitTransitionDelay() (webkitTransitionDelay string) {
	js.Rewrite("$<.WebkitTransitionDelay")
	return webkitTransitionDelay
}

// WebkitTransitionDelay
func (*CSSStyleDeclaration) SetWebkitTransitionDelay(webkitTransitionDelay string) {
	js.Rewrite("$<.WebkitTransitionDelay = $1", webkitTransitionDelay)
}

// WebkitTransitionDuration
func (*CSSStyleDeclaration) WebkitTransitionDuration() (webkitTransitionDuration string) {
	js.Rewrite("$<.WebkitTransitionDuration")
	return webkitTransitionDuration
}

// WebkitTransitionDuration
func (*CSSStyleDeclaration) SetWebkitTransitionDuration(webkitTransitionDuration string) {
	js.Rewrite("$<.WebkitTransitionDuration = $1", webkitTransitionDuration)
}

// WebkitTransitionProperty
func (*CSSStyleDeclaration) WebkitTransitionProperty() (webkitTransitionProperty string) {
	js.Rewrite("$<.WebkitTransitionProperty")
	return webkitTransitionProperty
}

// WebkitTransitionProperty
func (*CSSStyleDeclaration) SetWebkitTransitionProperty(webkitTransitionProperty string) {
	js.Rewrite("$<.WebkitTransitionProperty = $1", webkitTransitionProperty)
}

// WebkitTransitionTimingFunction
func (*CSSStyleDeclaration) WebkitTransitionTimingFunction() (webkitTransitionTimingFunction string) {
	js.Rewrite("$<.WebkitTransitionTimingFunction")
	return webkitTransitionTimingFunction
}

// WebkitTransitionTimingFunction
func (*CSSStyleDeclaration) SetWebkitTransitionTimingFunction(webkitTransitionTimingFunction string) {
	js.Rewrite("$<.WebkitTransitionTimingFunction = $1", webkitTransitionTimingFunction)
}

// WebkitUserModify
func (*CSSStyleDeclaration) WebkitUserModify() (webkitUserModify string) {
	js.Rewrite("$<.WebkitUserModify")
	return webkitUserModify
}

// WebkitUserModify
func (*CSSStyleDeclaration) SetWebkitUserModify(webkitUserModify string) {
	js.Rewrite("$<.WebkitUserModify = $1", webkitUserModify)
}

// WebkitUserSelect
func (*CSSStyleDeclaration) WebkitUserSelect() (webkitUserSelect string) {
	js.Rewrite("$<.WebkitUserSelect")
	return webkitUserSelect
}

// WebkitUserSelect
func (*CSSStyleDeclaration) SetWebkitUserSelect(webkitUserSelect string) {
	js.Rewrite("$<.WebkitUserSelect = $1", webkitUserSelect)
}

// WebkitWritingMode
func (*CSSStyleDeclaration) WebkitWritingMode() (webkitWritingMode string) {
	js.Rewrite("$<.WebkitWritingMode")
	return webkitWritingMode
}

// WebkitWritingMode
func (*CSSStyleDeclaration) SetWebkitWritingMode(webkitWritingMode string) {
	js.Rewrite("$<.WebkitWritingMode = $1", webkitWritingMode)
}

// WhiteSpace
func (*CSSStyleDeclaration) WhiteSpace() (whiteSpace string) {
	js.Rewrite("$<.WhiteSpace")
	return whiteSpace
}

// WhiteSpace
func (*CSSStyleDeclaration) SetWhiteSpace(whiteSpace string) {
	js.Rewrite("$<.WhiteSpace = $1", whiteSpace)
}

// Widows
func (*CSSStyleDeclaration) Widows() (widows string) {
	js.Rewrite("$<.Widows")
	return widows
}

// Widows
func (*CSSStyleDeclaration) SetWidows(widows string) {
	js.Rewrite("$<.Widows = $1", widows)
}

// Width
func (*CSSStyleDeclaration) Width() (width string) {
	js.Rewrite("$<.Width")
	return width
}

// Width
func (*CSSStyleDeclaration) SetWidth(width string) {
	js.Rewrite("$<.Width = $1", width)
}

// WordBreak
func (*CSSStyleDeclaration) WordBreak() (wordBreak string) {
	js.Rewrite("$<.WordBreak")
	return wordBreak
}

// WordBreak
func (*CSSStyleDeclaration) SetWordBreak(wordBreak string) {
	js.Rewrite("$<.WordBreak = $1", wordBreak)
}

// WordSpacing
func (*CSSStyleDeclaration) WordSpacing() (wordSpacing string) {
	js.Rewrite("$<.WordSpacing")
	return wordSpacing
}

// WordSpacing
func (*CSSStyleDeclaration) SetWordSpacing(wordSpacing string) {
	js.Rewrite("$<.WordSpacing = $1", wordSpacing)
}

// WordWrap
func (*CSSStyleDeclaration) WordWrap() (wordWrap string) {
	js.Rewrite("$<.WordWrap")
	return wordWrap
}

// WordWrap
func (*CSSStyleDeclaration) SetWordWrap(wordWrap string) {
	js.Rewrite("$<.WordWrap = $1", wordWrap)
}

// WritingMode
func (*CSSStyleDeclaration) WritingMode() (writingMode string) {
	js.Rewrite("$<.WritingMode")
	return writingMode
}

// WritingMode
func (*CSSStyleDeclaration) SetWritingMode(writingMode string) {
	js.Rewrite("$<.WritingMode = $1", writingMode)
}

// ZIndex
func (*CSSStyleDeclaration) ZIndex() (zIndex string) {
	js.Rewrite("$<.ZIndex")
	return zIndex
}

// ZIndex
func (*CSSStyleDeclaration) SetZIndex(zIndex string) {
	js.Rewrite("$<.ZIndex = $1", zIndex)
}

// Zoom
func (*CSSStyleDeclaration) Zoom() (zoom string) {
	js.Rewrite("$<.Zoom")
	return zoom
}

// Zoom
func (*CSSStyleDeclaration) SetZoom(zoom string) {
	js.Rewrite("$<.Zoom = $1", zoom)
}
