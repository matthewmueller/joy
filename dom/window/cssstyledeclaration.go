package window

import "github.com/matthewmueller/golly/js"

// CSSStyleDeclaration struct
// js:"CSSStyleDeclaration,omit"
type CSSStyleDeclaration struct {
}

// GetPropertyPriority fn
func (*CSSStyleDeclaration) GetPropertyPriority(propertyName string) (s string) {
	js.Rewrite("$<.getPropertyPriority($1)", propertyName)
	return s
}

// GetPropertyValue fn
func (*CSSStyleDeclaration) GetPropertyValue(propertyName string) (s string) {
	js.Rewrite("$<.getPropertyValue($1)", propertyName)
	return s
}

// Item fn
func (*CSSStyleDeclaration) Item(index uint) (s string) {
	js.Rewrite("$<.item($1)", index)
	return s
}

// RemoveProperty fn
func (*CSSStyleDeclaration) RemoveProperty(propertyName string) (s string) {
	js.Rewrite("$<.removeProperty($1)", propertyName)
	return s
}

// SetProperty fn
func (*CSSStyleDeclaration) SetProperty(propertyName string, value string, priority *string) {
	js.Rewrite("$<.setProperty($1, $2, $3)", propertyName, value, priority)
}

// AlignContent prop
func (*CSSStyleDeclaration) AlignContent() (alignContent string) {
	js.Rewrite("$<.alignContent")
	return alignContent
}

// AlignContent prop
func (*CSSStyleDeclaration) SetAlignContent(alignContent string) {
	js.Rewrite("$<.alignContent = $1", alignContent)
}

// AlignItems prop
func (*CSSStyleDeclaration) AlignItems() (alignItems string) {
	js.Rewrite("$<.alignItems")
	return alignItems
}

// AlignItems prop
func (*CSSStyleDeclaration) SetAlignItems(alignItems string) {
	js.Rewrite("$<.alignItems = $1", alignItems)
}

// AlignmentBaseline prop
func (*CSSStyleDeclaration) AlignmentBaseline() (alignmentBaseline string) {
	js.Rewrite("$<.alignmentBaseline")
	return alignmentBaseline
}

// AlignmentBaseline prop
func (*CSSStyleDeclaration) SetAlignmentBaseline(alignmentBaseline string) {
	js.Rewrite("$<.alignmentBaseline = $1", alignmentBaseline)
}

// AlignSelf prop
func (*CSSStyleDeclaration) AlignSelf() (alignSelf string) {
	js.Rewrite("$<.alignSelf")
	return alignSelf
}

// AlignSelf prop
func (*CSSStyleDeclaration) SetAlignSelf(alignSelf string) {
	js.Rewrite("$<.alignSelf = $1", alignSelf)
}

// Animation prop
func (*CSSStyleDeclaration) Animation() (animation string) {
	js.Rewrite("$<.animation")
	return animation
}

// Animation prop
func (*CSSStyleDeclaration) SetAnimation(animation string) {
	js.Rewrite("$<.animation = $1", animation)
}

// AnimationDelay prop
func (*CSSStyleDeclaration) AnimationDelay() (animationDelay string) {
	js.Rewrite("$<.animationDelay")
	return animationDelay
}

// AnimationDelay prop
func (*CSSStyleDeclaration) SetAnimationDelay(animationDelay string) {
	js.Rewrite("$<.animationDelay = $1", animationDelay)
}

// AnimationDirection prop
func (*CSSStyleDeclaration) AnimationDirection() (animationDirection string) {
	js.Rewrite("$<.animationDirection")
	return animationDirection
}

// AnimationDirection prop
func (*CSSStyleDeclaration) SetAnimationDirection(animationDirection string) {
	js.Rewrite("$<.animationDirection = $1", animationDirection)
}

// AnimationDuration prop
func (*CSSStyleDeclaration) AnimationDuration() (animationDuration string) {
	js.Rewrite("$<.animationDuration")
	return animationDuration
}

// AnimationDuration prop
func (*CSSStyleDeclaration) SetAnimationDuration(animationDuration string) {
	js.Rewrite("$<.animationDuration = $1", animationDuration)
}

// AnimationFillMode prop
func (*CSSStyleDeclaration) AnimationFillMode() (animationFillMode string) {
	js.Rewrite("$<.animationFillMode")
	return animationFillMode
}

// AnimationFillMode prop
func (*CSSStyleDeclaration) SetAnimationFillMode(animationFillMode string) {
	js.Rewrite("$<.animationFillMode = $1", animationFillMode)
}

// AnimationIterationCount prop
func (*CSSStyleDeclaration) AnimationIterationCount() (animationIterationCount string) {
	js.Rewrite("$<.animationIterationCount")
	return animationIterationCount
}

// AnimationIterationCount prop
func (*CSSStyleDeclaration) SetAnimationIterationCount(animationIterationCount string) {
	js.Rewrite("$<.animationIterationCount = $1", animationIterationCount)
}

// AnimationName prop
func (*CSSStyleDeclaration) AnimationName() (animationName string) {
	js.Rewrite("$<.animationName")
	return animationName
}

// AnimationName prop
func (*CSSStyleDeclaration) SetAnimationName(animationName string) {
	js.Rewrite("$<.animationName = $1", animationName)
}

// AnimationPlayState prop
func (*CSSStyleDeclaration) AnimationPlayState() (animationPlayState string) {
	js.Rewrite("$<.animationPlayState")
	return animationPlayState
}

// AnimationPlayState prop
func (*CSSStyleDeclaration) SetAnimationPlayState(animationPlayState string) {
	js.Rewrite("$<.animationPlayState = $1", animationPlayState)
}

// AnimationTimingFunction prop
func (*CSSStyleDeclaration) AnimationTimingFunction() (animationTimingFunction string) {
	js.Rewrite("$<.animationTimingFunction")
	return animationTimingFunction
}

// AnimationTimingFunction prop
func (*CSSStyleDeclaration) SetAnimationTimingFunction(animationTimingFunction string) {
	js.Rewrite("$<.animationTimingFunction = $1", animationTimingFunction)
}

// BackfaceVisibility prop
func (*CSSStyleDeclaration) BackfaceVisibility() (backfaceVisibility string) {
	js.Rewrite("$<.backfaceVisibility")
	return backfaceVisibility
}

// BackfaceVisibility prop
func (*CSSStyleDeclaration) SetBackfaceVisibility(backfaceVisibility string) {
	js.Rewrite("$<.backfaceVisibility = $1", backfaceVisibility)
}

// Background prop
func (*CSSStyleDeclaration) Background() (background string) {
	js.Rewrite("$<.background")
	return background
}

// Background prop
func (*CSSStyleDeclaration) SetBackground(background string) {
	js.Rewrite("$<.background = $1", background)
}

// BackgroundAttachment prop
func (*CSSStyleDeclaration) BackgroundAttachment() (backgroundAttachment string) {
	js.Rewrite("$<.backgroundAttachment")
	return backgroundAttachment
}

// BackgroundAttachment prop
func (*CSSStyleDeclaration) SetBackgroundAttachment(backgroundAttachment string) {
	js.Rewrite("$<.backgroundAttachment = $1", backgroundAttachment)
}

// BackgroundClip prop
func (*CSSStyleDeclaration) BackgroundClip() (backgroundClip string) {
	js.Rewrite("$<.backgroundClip")
	return backgroundClip
}

// BackgroundClip prop
func (*CSSStyleDeclaration) SetBackgroundClip(backgroundClip string) {
	js.Rewrite("$<.backgroundClip = $1", backgroundClip)
}

// BackgroundColor prop
func (*CSSStyleDeclaration) BackgroundColor() (backgroundColor string) {
	js.Rewrite("$<.backgroundColor")
	return backgroundColor
}

// BackgroundColor prop
func (*CSSStyleDeclaration) SetBackgroundColor(backgroundColor string) {
	js.Rewrite("$<.backgroundColor = $1", backgroundColor)
}

// BackgroundImage prop
func (*CSSStyleDeclaration) BackgroundImage() (backgroundImage string) {
	js.Rewrite("$<.backgroundImage")
	return backgroundImage
}

// BackgroundImage prop
func (*CSSStyleDeclaration) SetBackgroundImage(backgroundImage string) {
	js.Rewrite("$<.backgroundImage = $1", backgroundImage)
}

// BackgroundOrigin prop
func (*CSSStyleDeclaration) BackgroundOrigin() (backgroundOrigin string) {
	js.Rewrite("$<.backgroundOrigin")
	return backgroundOrigin
}

// BackgroundOrigin prop
func (*CSSStyleDeclaration) SetBackgroundOrigin(backgroundOrigin string) {
	js.Rewrite("$<.backgroundOrigin = $1", backgroundOrigin)
}

// BackgroundPosition prop
func (*CSSStyleDeclaration) BackgroundPosition() (backgroundPosition string) {
	js.Rewrite("$<.backgroundPosition")
	return backgroundPosition
}

// BackgroundPosition prop
func (*CSSStyleDeclaration) SetBackgroundPosition(backgroundPosition string) {
	js.Rewrite("$<.backgroundPosition = $1", backgroundPosition)
}

// BackgroundPositionX prop
func (*CSSStyleDeclaration) BackgroundPositionX() (backgroundPositionX string) {
	js.Rewrite("$<.backgroundPositionX")
	return backgroundPositionX
}

// BackgroundPositionX prop
func (*CSSStyleDeclaration) SetBackgroundPositionX(backgroundPositionX string) {
	js.Rewrite("$<.backgroundPositionX = $1", backgroundPositionX)
}

// BackgroundPositionY prop
func (*CSSStyleDeclaration) BackgroundPositionY() (backgroundPositionY string) {
	js.Rewrite("$<.backgroundPositionY")
	return backgroundPositionY
}

// BackgroundPositionY prop
func (*CSSStyleDeclaration) SetBackgroundPositionY(backgroundPositionY string) {
	js.Rewrite("$<.backgroundPositionY = $1", backgroundPositionY)
}

// BackgroundRepeat prop
func (*CSSStyleDeclaration) BackgroundRepeat() (backgroundRepeat string) {
	js.Rewrite("$<.backgroundRepeat")
	return backgroundRepeat
}

// BackgroundRepeat prop
func (*CSSStyleDeclaration) SetBackgroundRepeat(backgroundRepeat string) {
	js.Rewrite("$<.backgroundRepeat = $1", backgroundRepeat)
}

// BackgroundSize prop
func (*CSSStyleDeclaration) BackgroundSize() (backgroundSize string) {
	js.Rewrite("$<.backgroundSize")
	return backgroundSize
}

// BackgroundSize prop
func (*CSSStyleDeclaration) SetBackgroundSize(backgroundSize string) {
	js.Rewrite("$<.backgroundSize = $1", backgroundSize)
}

// BaselineShift prop
func (*CSSStyleDeclaration) BaselineShift() (baselineShift string) {
	js.Rewrite("$<.baselineShift")
	return baselineShift
}

// BaselineShift prop
func (*CSSStyleDeclaration) SetBaselineShift(baselineShift string) {
	js.Rewrite("$<.baselineShift = $1", baselineShift)
}

// Border prop
func (*CSSStyleDeclaration) Border() (border string) {
	js.Rewrite("$<.border")
	return border
}

// Border prop
func (*CSSStyleDeclaration) SetBorder(border string) {
	js.Rewrite("$<.border = $1", border)
}

// BorderBottom prop
func (*CSSStyleDeclaration) BorderBottom() (borderBottom string) {
	js.Rewrite("$<.borderBottom")
	return borderBottom
}

// BorderBottom prop
func (*CSSStyleDeclaration) SetBorderBottom(borderBottom string) {
	js.Rewrite("$<.borderBottom = $1", borderBottom)
}

// BorderBottomColor prop
func (*CSSStyleDeclaration) BorderBottomColor() (borderBottomColor string) {
	js.Rewrite("$<.borderBottomColor")
	return borderBottomColor
}

// BorderBottomColor prop
func (*CSSStyleDeclaration) SetBorderBottomColor(borderBottomColor string) {
	js.Rewrite("$<.borderBottomColor = $1", borderBottomColor)
}

// BorderBottomLeftRadius prop
func (*CSSStyleDeclaration) BorderBottomLeftRadius() (borderBottomLeftRadius string) {
	js.Rewrite("$<.borderBottomLeftRadius")
	return borderBottomLeftRadius
}

// BorderBottomLeftRadius prop
func (*CSSStyleDeclaration) SetBorderBottomLeftRadius(borderBottomLeftRadius string) {
	js.Rewrite("$<.borderBottomLeftRadius = $1", borderBottomLeftRadius)
}

// BorderBottomRightRadius prop
func (*CSSStyleDeclaration) BorderBottomRightRadius() (borderBottomRightRadius string) {
	js.Rewrite("$<.borderBottomRightRadius")
	return borderBottomRightRadius
}

// BorderBottomRightRadius prop
func (*CSSStyleDeclaration) SetBorderBottomRightRadius(borderBottomRightRadius string) {
	js.Rewrite("$<.borderBottomRightRadius = $1", borderBottomRightRadius)
}

// BorderBottomStyle prop
func (*CSSStyleDeclaration) BorderBottomStyle() (borderBottomStyle string) {
	js.Rewrite("$<.borderBottomStyle")
	return borderBottomStyle
}

// BorderBottomStyle prop
func (*CSSStyleDeclaration) SetBorderBottomStyle(borderBottomStyle string) {
	js.Rewrite("$<.borderBottomStyle = $1", borderBottomStyle)
}

// BorderBottomWidth prop
func (*CSSStyleDeclaration) BorderBottomWidth() (borderBottomWidth string) {
	js.Rewrite("$<.borderBottomWidth")
	return borderBottomWidth
}

// BorderBottomWidth prop
func (*CSSStyleDeclaration) SetBorderBottomWidth(borderBottomWidth string) {
	js.Rewrite("$<.borderBottomWidth = $1", borderBottomWidth)
}

// BorderCollapse prop
func (*CSSStyleDeclaration) BorderCollapse() (borderCollapse string) {
	js.Rewrite("$<.borderCollapse")
	return borderCollapse
}

// BorderCollapse prop
func (*CSSStyleDeclaration) SetBorderCollapse(borderCollapse string) {
	js.Rewrite("$<.borderCollapse = $1", borderCollapse)
}

// BorderColor prop
func (*CSSStyleDeclaration) BorderColor() (borderColor string) {
	js.Rewrite("$<.borderColor")
	return borderColor
}

// BorderColor prop
func (*CSSStyleDeclaration) SetBorderColor(borderColor string) {
	js.Rewrite("$<.borderColor = $1", borderColor)
}

// BorderImage prop
func (*CSSStyleDeclaration) BorderImage() (borderImage string) {
	js.Rewrite("$<.borderImage")
	return borderImage
}

// BorderImage prop
func (*CSSStyleDeclaration) SetBorderImage(borderImage string) {
	js.Rewrite("$<.borderImage = $1", borderImage)
}

// BorderImageOutset prop
func (*CSSStyleDeclaration) BorderImageOutset() (borderImageOutset string) {
	js.Rewrite("$<.borderImageOutset")
	return borderImageOutset
}

// BorderImageOutset prop
func (*CSSStyleDeclaration) SetBorderImageOutset(borderImageOutset string) {
	js.Rewrite("$<.borderImageOutset = $1", borderImageOutset)
}

// BorderImageRepeat prop
func (*CSSStyleDeclaration) BorderImageRepeat() (borderImageRepeat string) {
	js.Rewrite("$<.borderImageRepeat")
	return borderImageRepeat
}

// BorderImageRepeat prop
func (*CSSStyleDeclaration) SetBorderImageRepeat(borderImageRepeat string) {
	js.Rewrite("$<.borderImageRepeat = $1", borderImageRepeat)
}

// BorderImageSlice prop
func (*CSSStyleDeclaration) BorderImageSlice() (borderImageSlice string) {
	js.Rewrite("$<.borderImageSlice")
	return borderImageSlice
}

// BorderImageSlice prop
func (*CSSStyleDeclaration) SetBorderImageSlice(borderImageSlice string) {
	js.Rewrite("$<.borderImageSlice = $1", borderImageSlice)
}

// BorderImageSource prop
func (*CSSStyleDeclaration) BorderImageSource() (borderImageSource string) {
	js.Rewrite("$<.borderImageSource")
	return borderImageSource
}

// BorderImageSource prop
func (*CSSStyleDeclaration) SetBorderImageSource(borderImageSource string) {
	js.Rewrite("$<.borderImageSource = $1", borderImageSource)
}

// BorderImageWidth prop
func (*CSSStyleDeclaration) BorderImageWidth() (borderImageWidth string) {
	js.Rewrite("$<.borderImageWidth")
	return borderImageWidth
}

// BorderImageWidth prop
func (*CSSStyleDeclaration) SetBorderImageWidth(borderImageWidth string) {
	js.Rewrite("$<.borderImageWidth = $1", borderImageWidth)
}

// BorderLeft prop
func (*CSSStyleDeclaration) BorderLeft() (borderLeft string) {
	js.Rewrite("$<.borderLeft")
	return borderLeft
}

// BorderLeft prop
func (*CSSStyleDeclaration) SetBorderLeft(borderLeft string) {
	js.Rewrite("$<.borderLeft = $1", borderLeft)
}

// BorderLeftColor prop
func (*CSSStyleDeclaration) BorderLeftColor() (borderLeftColor string) {
	js.Rewrite("$<.borderLeftColor")
	return borderLeftColor
}

// BorderLeftColor prop
func (*CSSStyleDeclaration) SetBorderLeftColor(borderLeftColor string) {
	js.Rewrite("$<.borderLeftColor = $1", borderLeftColor)
}

// BorderLeftStyle prop
func (*CSSStyleDeclaration) BorderLeftStyle() (borderLeftStyle string) {
	js.Rewrite("$<.borderLeftStyle")
	return borderLeftStyle
}

// BorderLeftStyle prop
func (*CSSStyleDeclaration) SetBorderLeftStyle(borderLeftStyle string) {
	js.Rewrite("$<.borderLeftStyle = $1", borderLeftStyle)
}

// BorderLeftWidth prop
func (*CSSStyleDeclaration) BorderLeftWidth() (borderLeftWidth string) {
	js.Rewrite("$<.borderLeftWidth")
	return borderLeftWidth
}

// BorderLeftWidth prop
func (*CSSStyleDeclaration) SetBorderLeftWidth(borderLeftWidth string) {
	js.Rewrite("$<.borderLeftWidth = $1", borderLeftWidth)
}

// BorderRadius prop
func (*CSSStyleDeclaration) BorderRadius() (borderRadius string) {
	js.Rewrite("$<.borderRadius")
	return borderRadius
}

// BorderRadius prop
func (*CSSStyleDeclaration) SetBorderRadius(borderRadius string) {
	js.Rewrite("$<.borderRadius = $1", borderRadius)
}

// BorderRight prop
func (*CSSStyleDeclaration) BorderRight() (borderRight string) {
	js.Rewrite("$<.borderRight")
	return borderRight
}

// BorderRight prop
func (*CSSStyleDeclaration) SetBorderRight(borderRight string) {
	js.Rewrite("$<.borderRight = $1", borderRight)
}

// BorderRightColor prop
func (*CSSStyleDeclaration) BorderRightColor() (borderRightColor string) {
	js.Rewrite("$<.borderRightColor")
	return borderRightColor
}

// BorderRightColor prop
func (*CSSStyleDeclaration) SetBorderRightColor(borderRightColor string) {
	js.Rewrite("$<.borderRightColor = $1", borderRightColor)
}

// BorderRightStyle prop
func (*CSSStyleDeclaration) BorderRightStyle() (borderRightStyle string) {
	js.Rewrite("$<.borderRightStyle")
	return borderRightStyle
}

// BorderRightStyle prop
func (*CSSStyleDeclaration) SetBorderRightStyle(borderRightStyle string) {
	js.Rewrite("$<.borderRightStyle = $1", borderRightStyle)
}

// BorderRightWidth prop
func (*CSSStyleDeclaration) BorderRightWidth() (borderRightWidth string) {
	js.Rewrite("$<.borderRightWidth")
	return borderRightWidth
}

// BorderRightWidth prop
func (*CSSStyleDeclaration) SetBorderRightWidth(borderRightWidth string) {
	js.Rewrite("$<.borderRightWidth = $1", borderRightWidth)
}

// BorderSpacing prop
func (*CSSStyleDeclaration) BorderSpacing() (borderSpacing string) {
	js.Rewrite("$<.borderSpacing")
	return borderSpacing
}

// BorderSpacing prop
func (*CSSStyleDeclaration) SetBorderSpacing(borderSpacing string) {
	js.Rewrite("$<.borderSpacing = $1", borderSpacing)
}

// BorderStyle prop
func (*CSSStyleDeclaration) BorderStyle() (borderStyle string) {
	js.Rewrite("$<.borderStyle")
	return borderStyle
}

// BorderStyle prop
func (*CSSStyleDeclaration) SetBorderStyle(borderStyle string) {
	js.Rewrite("$<.borderStyle = $1", borderStyle)
}

// BorderTop prop
func (*CSSStyleDeclaration) BorderTop() (borderTop string) {
	js.Rewrite("$<.borderTop")
	return borderTop
}

// BorderTop prop
func (*CSSStyleDeclaration) SetBorderTop(borderTop string) {
	js.Rewrite("$<.borderTop = $1", borderTop)
}

// BorderTopColor prop
func (*CSSStyleDeclaration) BorderTopColor() (borderTopColor string) {
	js.Rewrite("$<.borderTopColor")
	return borderTopColor
}

// BorderTopColor prop
func (*CSSStyleDeclaration) SetBorderTopColor(borderTopColor string) {
	js.Rewrite("$<.borderTopColor = $1", borderTopColor)
}

// BorderTopLeftRadius prop
func (*CSSStyleDeclaration) BorderTopLeftRadius() (borderTopLeftRadius string) {
	js.Rewrite("$<.borderTopLeftRadius")
	return borderTopLeftRadius
}

// BorderTopLeftRadius prop
func (*CSSStyleDeclaration) SetBorderTopLeftRadius(borderTopLeftRadius string) {
	js.Rewrite("$<.borderTopLeftRadius = $1", borderTopLeftRadius)
}

// BorderTopRightRadius prop
func (*CSSStyleDeclaration) BorderTopRightRadius() (borderTopRightRadius string) {
	js.Rewrite("$<.borderTopRightRadius")
	return borderTopRightRadius
}

// BorderTopRightRadius prop
func (*CSSStyleDeclaration) SetBorderTopRightRadius(borderTopRightRadius string) {
	js.Rewrite("$<.borderTopRightRadius = $1", borderTopRightRadius)
}

// BorderTopStyle prop
func (*CSSStyleDeclaration) BorderTopStyle() (borderTopStyle string) {
	js.Rewrite("$<.borderTopStyle")
	return borderTopStyle
}

// BorderTopStyle prop
func (*CSSStyleDeclaration) SetBorderTopStyle(borderTopStyle string) {
	js.Rewrite("$<.borderTopStyle = $1", borderTopStyle)
}

// BorderTopWidth prop
func (*CSSStyleDeclaration) BorderTopWidth() (borderTopWidth string) {
	js.Rewrite("$<.borderTopWidth")
	return borderTopWidth
}

// BorderTopWidth prop
func (*CSSStyleDeclaration) SetBorderTopWidth(borderTopWidth string) {
	js.Rewrite("$<.borderTopWidth = $1", borderTopWidth)
}

// BorderWidth prop
func (*CSSStyleDeclaration) BorderWidth() (borderWidth string) {
	js.Rewrite("$<.borderWidth")
	return borderWidth
}

// BorderWidth prop
func (*CSSStyleDeclaration) SetBorderWidth(borderWidth string) {
	js.Rewrite("$<.borderWidth = $1", borderWidth)
}

// Bottom prop
func (*CSSStyleDeclaration) Bottom() (bottom string) {
	js.Rewrite("$<.bottom")
	return bottom
}

// Bottom prop
func (*CSSStyleDeclaration) SetBottom(bottom string) {
	js.Rewrite("$<.bottom = $1", bottom)
}

// BoxShadow prop
func (*CSSStyleDeclaration) BoxShadow() (boxShadow string) {
	js.Rewrite("$<.boxShadow")
	return boxShadow
}

// BoxShadow prop
func (*CSSStyleDeclaration) SetBoxShadow(boxShadow string) {
	js.Rewrite("$<.boxShadow = $1", boxShadow)
}

// BoxSizing prop
func (*CSSStyleDeclaration) BoxSizing() (boxSizing string) {
	js.Rewrite("$<.boxSizing")
	return boxSizing
}

// BoxSizing prop
func (*CSSStyleDeclaration) SetBoxSizing(boxSizing string) {
	js.Rewrite("$<.boxSizing = $1", boxSizing)
}

// BreakAfter prop
func (*CSSStyleDeclaration) BreakAfter() (breakAfter string) {
	js.Rewrite("$<.breakAfter")
	return breakAfter
}

// BreakAfter prop
func (*CSSStyleDeclaration) SetBreakAfter(breakAfter string) {
	js.Rewrite("$<.breakAfter = $1", breakAfter)
}

// BreakBefore prop
func (*CSSStyleDeclaration) BreakBefore() (breakBefore string) {
	js.Rewrite("$<.breakBefore")
	return breakBefore
}

// BreakBefore prop
func (*CSSStyleDeclaration) SetBreakBefore(breakBefore string) {
	js.Rewrite("$<.breakBefore = $1", breakBefore)
}

// BreakInside prop
func (*CSSStyleDeclaration) BreakInside() (breakInside string) {
	js.Rewrite("$<.breakInside")
	return breakInside
}

// BreakInside prop
func (*CSSStyleDeclaration) SetBreakInside(breakInside string) {
	js.Rewrite("$<.breakInside = $1", breakInside)
}

// CaptionSide prop
func (*CSSStyleDeclaration) CaptionSide() (captionSide string) {
	js.Rewrite("$<.captionSide")
	return captionSide
}

// CaptionSide prop
func (*CSSStyleDeclaration) SetCaptionSide(captionSide string) {
	js.Rewrite("$<.captionSide = $1", captionSide)
}

// Clear prop
func (*CSSStyleDeclaration) Clear() (clear string) {
	js.Rewrite("$<.clear")
	return clear
}

// Clear prop
func (*CSSStyleDeclaration) SetClear(clear string) {
	js.Rewrite("$<.clear = $1", clear)
}

// Clip prop
func (*CSSStyleDeclaration) Clip() (clip string) {
	js.Rewrite("$<.clip")
	return clip
}

// Clip prop
func (*CSSStyleDeclaration) SetClip(clip string) {
	js.Rewrite("$<.clip = $1", clip)
}

// ClipPath prop
func (*CSSStyleDeclaration) ClipPath() (clipPath string) {
	js.Rewrite("$<.clipPath")
	return clipPath
}

// ClipPath prop
func (*CSSStyleDeclaration) SetClipPath(clipPath string) {
	js.Rewrite("$<.clipPath = $1", clipPath)
}

// ClipRule prop
func (*CSSStyleDeclaration) ClipRule() (clipRule string) {
	js.Rewrite("$<.clipRule")
	return clipRule
}

// ClipRule prop
func (*CSSStyleDeclaration) SetClipRule(clipRule string) {
	js.Rewrite("$<.clipRule = $1", clipRule)
}

// Color prop
func (*CSSStyleDeclaration) Color() (color string) {
	js.Rewrite("$<.color")
	return color
}

// Color prop
func (*CSSStyleDeclaration) SetColor(color string) {
	js.Rewrite("$<.color = $1", color)
}

// ColorInterpolationFilters prop
func (*CSSStyleDeclaration) ColorInterpolationFilters() (colorInterpolationFilters string) {
	js.Rewrite("$<.colorInterpolationFilters")
	return colorInterpolationFilters
}

// ColorInterpolationFilters prop
func (*CSSStyleDeclaration) SetColorInterpolationFilters(colorInterpolationFilters string) {
	js.Rewrite("$<.colorInterpolationFilters = $1", colorInterpolationFilters)
}

// ColumnCount prop
func (*CSSStyleDeclaration) ColumnCount() (columnCount interface{}) {
	js.Rewrite("$<.columnCount")
	return columnCount
}

// ColumnCount prop
func (*CSSStyleDeclaration) SetColumnCount(columnCount interface{}) {
	js.Rewrite("$<.columnCount = $1", columnCount)
}

// ColumnFill prop
func (*CSSStyleDeclaration) ColumnFill() (columnFill string) {
	js.Rewrite("$<.columnFill")
	return columnFill
}

// ColumnFill prop
func (*CSSStyleDeclaration) SetColumnFill(columnFill string) {
	js.Rewrite("$<.columnFill = $1", columnFill)
}

// ColumnGap prop
func (*CSSStyleDeclaration) ColumnGap() (columnGap interface{}) {
	js.Rewrite("$<.columnGap")
	return columnGap
}

// ColumnGap prop
func (*CSSStyleDeclaration) SetColumnGap(columnGap interface{}) {
	js.Rewrite("$<.columnGap = $1", columnGap)
}

// ColumnRule prop
func (*CSSStyleDeclaration) ColumnRule() (columnRule string) {
	js.Rewrite("$<.columnRule")
	return columnRule
}

// ColumnRule prop
func (*CSSStyleDeclaration) SetColumnRule(columnRule string) {
	js.Rewrite("$<.columnRule = $1", columnRule)
}

// ColumnRuleColor prop
func (*CSSStyleDeclaration) ColumnRuleColor() (columnRuleColor interface{}) {
	js.Rewrite("$<.columnRuleColor")
	return columnRuleColor
}

// ColumnRuleColor prop
func (*CSSStyleDeclaration) SetColumnRuleColor(columnRuleColor interface{}) {
	js.Rewrite("$<.columnRuleColor = $1", columnRuleColor)
}

// ColumnRuleStyle prop
func (*CSSStyleDeclaration) ColumnRuleStyle() (columnRuleStyle string) {
	js.Rewrite("$<.columnRuleStyle")
	return columnRuleStyle
}

// ColumnRuleStyle prop
func (*CSSStyleDeclaration) SetColumnRuleStyle(columnRuleStyle string) {
	js.Rewrite("$<.columnRuleStyle = $1", columnRuleStyle)
}

// ColumnRuleWidth prop
func (*CSSStyleDeclaration) ColumnRuleWidth() (columnRuleWidth interface{}) {
	js.Rewrite("$<.columnRuleWidth")
	return columnRuleWidth
}

// ColumnRuleWidth prop
func (*CSSStyleDeclaration) SetColumnRuleWidth(columnRuleWidth interface{}) {
	js.Rewrite("$<.columnRuleWidth = $1", columnRuleWidth)
}

// Columns prop
func (*CSSStyleDeclaration) Columns() (columns string) {
	js.Rewrite("$<.columns")
	return columns
}

// Columns prop
func (*CSSStyleDeclaration) SetColumns(columns string) {
	js.Rewrite("$<.columns = $1", columns)
}

// ColumnSpan prop
func (*CSSStyleDeclaration) ColumnSpan() (columnSpan string) {
	js.Rewrite("$<.columnSpan")
	return columnSpan
}

// ColumnSpan prop
func (*CSSStyleDeclaration) SetColumnSpan(columnSpan string) {
	js.Rewrite("$<.columnSpan = $1", columnSpan)
}

// ColumnWidth prop
func (*CSSStyleDeclaration) ColumnWidth() (columnWidth interface{}) {
	js.Rewrite("$<.columnWidth")
	return columnWidth
}

// ColumnWidth prop
func (*CSSStyleDeclaration) SetColumnWidth(columnWidth interface{}) {
	js.Rewrite("$<.columnWidth = $1", columnWidth)
}

// Content prop
func (*CSSStyleDeclaration) Content() (content string) {
	js.Rewrite("$<.content")
	return content
}

// Content prop
func (*CSSStyleDeclaration) SetContent(content string) {
	js.Rewrite("$<.content = $1", content)
}

// CounterIncrement prop
func (*CSSStyleDeclaration) CounterIncrement() (counterIncrement string) {
	js.Rewrite("$<.counterIncrement")
	return counterIncrement
}

// CounterIncrement prop
func (*CSSStyleDeclaration) SetCounterIncrement(counterIncrement string) {
	js.Rewrite("$<.counterIncrement = $1", counterIncrement)
}

// CounterReset prop
func (*CSSStyleDeclaration) CounterReset() (counterReset string) {
	js.Rewrite("$<.counterReset")
	return counterReset
}

// CounterReset prop
func (*CSSStyleDeclaration) SetCounterReset(counterReset string) {
	js.Rewrite("$<.counterReset = $1", counterReset)
}

// CSSFloat prop
func (*CSSStyleDeclaration) CSSFloat() (cssFloat string) {
	js.Rewrite("$<.cssFloat")
	return cssFloat
}

// CSSFloat prop
func (*CSSStyleDeclaration) SetCSSFloat(cssFloat string) {
	js.Rewrite("$<.cssFloat = $1", cssFloat)
}

// CSSText prop
func (*CSSStyleDeclaration) CSSText() (cssText string) {
	js.Rewrite("$<.cssText")
	return cssText
}

// CSSText prop
func (*CSSStyleDeclaration) SetCSSText(cssText string) {
	js.Rewrite("$<.cssText = $1", cssText)
}

// Cursor prop
func (*CSSStyleDeclaration) Cursor() (cursor string) {
	js.Rewrite("$<.cursor")
	return cursor
}

// Cursor prop
func (*CSSStyleDeclaration) SetCursor(cursor string) {
	js.Rewrite("$<.cursor = $1", cursor)
}

// Direction prop
func (*CSSStyleDeclaration) Direction() (direction string) {
	js.Rewrite("$<.direction")
	return direction
}

// Direction prop
func (*CSSStyleDeclaration) SetDirection(direction string) {
	js.Rewrite("$<.direction = $1", direction)
}

// Display prop
func (*CSSStyleDeclaration) Display() (display string) {
	js.Rewrite("$<.display")
	return display
}

// Display prop
func (*CSSStyleDeclaration) SetDisplay(display string) {
	js.Rewrite("$<.display = $1", display)
}

// DominantBaseline prop
func (*CSSStyleDeclaration) DominantBaseline() (dominantBaseline string) {
	js.Rewrite("$<.dominantBaseline")
	return dominantBaseline
}

// DominantBaseline prop
func (*CSSStyleDeclaration) SetDominantBaseline(dominantBaseline string) {
	js.Rewrite("$<.dominantBaseline = $1", dominantBaseline)
}

// EmptyCells prop
func (*CSSStyleDeclaration) EmptyCells() (emptyCells string) {
	js.Rewrite("$<.emptyCells")
	return emptyCells
}

// EmptyCells prop
func (*CSSStyleDeclaration) SetEmptyCells(emptyCells string) {
	js.Rewrite("$<.emptyCells = $1", emptyCells)
}

// EnableBackground prop
func (*CSSStyleDeclaration) EnableBackground() (enableBackground string) {
	js.Rewrite("$<.enableBackground")
	return enableBackground
}

// EnableBackground prop
func (*CSSStyleDeclaration) SetEnableBackground(enableBackground string) {
	js.Rewrite("$<.enableBackground = $1", enableBackground)
}

// Fill prop
func (*CSSStyleDeclaration) Fill() (fill string) {
	js.Rewrite("$<.fill")
	return fill
}

// Fill prop
func (*CSSStyleDeclaration) SetFill(fill string) {
	js.Rewrite("$<.fill = $1", fill)
}

// FillOpacity prop
func (*CSSStyleDeclaration) FillOpacity() (fillOpacity string) {
	js.Rewrite("$<.fillOpacity")
	return fillOpacity
}

// FillOpacity prop
func (*CSSStyleDeclaration) SetFillOpacity(fillOpacity string) {
	js.Rewrite("$<.fillOpacity = $1", fillOpacity)
}

// FillRule prop
func (*CSSStyleDeclaration) FillRule() (fillRule string) {
	js.Rewrite("$<.fillRule")
	return fillRule
}

// FillRule prop
func (*CSSStyleDeclaration) SetFillRule(fillRule string) {
	js.Rewrite("$<.fillRule = $1", fillRule)
}

// Filter prop
func (*CSSStyleDeclaration) Filter() (filter string) {
	js.Rewrite("$<.filter")
	return filter
}

// Filter prop
func (*CSSStyleDeclaration) SetFilter(filter string) {
	js.Rewrite("$<.filter = $1", filter)
}

// Flex prop
func (*CSSStyleDeclaration) Flex() (flex string) {
	js.Rewrite("$<.flex")
	return flex
}

// Flex prop
func (*CSSStyleDeclaration) SetFlex(flex string) {
	js.Rewrite("$<.flex = $1", flex)
}

// FlexBasis prop
func (*CSSStyleDeclaration) FlexBasis() (flexBasis string) {
	js.Rewrite("$<.flexBasis")
	return flexBasis
}

// FlexBasis prop
func (*CSSStyleDeclaration) SetFlexBasis(flexBasis string) {
	js.Rewrite("$<.flexBasis = $1", flexBasis)
}

// FlexDirection prop
func (*CSSStyleDeclaration) FlexDirection() (flexDirection string) {
	js.Rewrite("$<.flexDirection")
	return flexDirection
}

// FlexDirection prop
func (*CSSStyleDeclaration) SetFlexDirection(flexDirection string) {
	js.Rewrite("$<.flexDirection = $1", flexDirection)
}

// FlexFlow prop
func (*CSSStyleDeclaration) FlexFlow() (flexFlow string) {
	js.Rewrite("$<.flexFlow")
	return flexFlow
}

// FlexFlow prop
func (*CSSStyleDeclaration) SetFlexFlow(flexFlow string) {
	js.Rewrite("$<.flexFlow = $1", flexFlow)
}

// FlexGrow prop
func (*CSSStyleDeclaration) FlexGrow() (flexGrow string) {
	js.Rewrite("$<.flexGrow")
	return flexGrow
}

// FlexGrow prop
func (*CSSStyleDeclaration) SetFlexGrow(flexGrow string) {
	js.Rewrite("$<.flexGrow = $1", flexGrow)
}

// FlexShrink prop
func (*CSSStyleDeclaration) FlexShrink() (flexShrink string) {
	js.Rewrite("$<.flexShrink")
	return flexShrink
}

// FlexShrink prop
func (*CSSStyleDeclaration) SetFlexShrink(flexShrink string) {
	js.Rewrite("$<.flexShrink = $1", flexShrink)
}

// FlexWrap prop
func (*CSSStyleDeclaration) FlexWrap() (flexWrap string) {
	js.Rewrite("$<.flexWrap")
	return flexWrap
}

// FlexWrap prop
func (*CSSStyleDeclaration) SetFlexWrap(flexWrap string) {
	js.Rewrite("$<.flexWrap = $1", flexWrap)
}

// FloodColor prop
func (*CSSStyleDeclaration) FloodColor() (floodColor string) {
	js.Rewrite("$<.floodColor")
	return floodColor
}

// FloodColor prop
func (*CSSStyleDeclaration) SetFloodColor(floodColor string) {
	js.Rewrite("$<.floodColor = $1", floodColor)
}

// FloodOpacity prop
func (*CSSStyleDeclaration) FloodOpacity() (floodOpacity string) {
	js.Rewrite("$<.floodOpacity")
	return floodOpacity
}

// FloodOpacity prop
func (*CSSStyleDeclaration) SetFloodOpacity(floodOpacity string) {
	js.Rewrite("$<.floodOpacity = $1", floodOpacity)
}

// Font prop
func (*CSSStyleDeclaration) Font() (font string) {
	js.Rewrite("$<.font")
	return font
}

// Font prop
func (*CSSStyleDeclaration) SetFont(font string) {
	js.Rewrite("$<.font = $1", font)
}

// FontFamily prop
func (*CSSStyleDeclaration) FontFamily() (fontFamily string) {
	js.Rewrite("$<.fontFamily")
	return fontFamily
}

// FontFamily prop
func (*CSSStyleDeclaration) SetFontFamily(fontFamily string) {
	js.Rewrite("$<.fontFamily = $1", fontFamily)
}

// FontFeatureSettings prop
func (*CSSStyleDeclaration) FontFeatureSettings() (fontFeatureSettings string) {
	js.Rewrite("$<.fontFeatureSettings")
	return fontFeatureSettings
}

// FontFeatureSettings prop
func (*CSSStyleDeclaration) SetFontFeatureSettings(fontFeatureSettings string) {
	js.Rewrite("$<.fontFeatureSettings = $1", fontFeatureSettings)
}

// FontSize prop
func (*CSSStyleDeclaration) FontSize() (fontSize string) {
	js.Rewrite("$<.fontSize")
	return fontSize
}

// FontSize prop
func (*CSSStyleDeclaration) SetFontSize(fontSize string) {
	js.Rewrite("$<.fontSize = $1", fontSize)
}

// FontSizeAdjust prop
func (*CSSStyleDeclaration) FontSizeAdjust() (fontSizeAdjust string) {
	js.Rewrite("$<.fontSizeAdjust")
	return fontSizeAdjust
}

// FontSizeAdjust prop
func (*CSSStyleDeclaration) SetFontSizeAdjust(fontSizeAdjust string) {
	js.Rewrite("$<.fontSizeAdjust = $1", fontSizeAdjust)
}

// FontStretch prop
func (*CSSStyleDeclaration) FontStretch() (fontStretch string) {
	js.Rewrite("$<.fontStretch")
	return fontStretch
}

// FontStretch prop
func (*CSSStyleDeclaration) SetFontStretch(fontStretch string) {
	js.Rewrite("$<.fontStretch = $1", fontStretch)
}

// FontStyle prop
func (*CSSStyleDeclaration) FontStyle() (fontStyle string) {
	js.Rewrite("$<.fontStyle")
	return fontStyle
}

// FontStyle prop
func (*CSSStyleDeclaration) SetFontStyle(fontStyle string) {
	js.Rewrite("$<.fontStyle = $1", fontStyle)
}

// FontVariant prop
func (*CSSStyleDeclaration) FontVariant() (fontVariant string) {
	js.Rewrite("$<.fontVariant")
	return fontVariant
}

// FontVariant prop
func (*CSSStyleDeclaration) SetFontVariant(fontVariant string) {
	js.Rewrite("$<.fontVariant = $1", fontVariant)
}

// FontWeight prop
func (*CSSStyleDeclaration) FontWeight() (fontWeight string) {
	js.Rewrite("$<.fontWeight")
	return fontWeight
}

// FontWeight prop
func (*CSSStyleDeclaration) SetFontWeight(fontWeight string) {
	js.Rewrite("$<.fontWeight = $1", fontWeight)
}

// GlyphOrientationHorizontal prop
func (*CSSStyleDeclaration) GlyphOrientationHorizontal() (glyphOrientationHorizontal string) {
	js.Rewrite("$<.glyphOrientationHorizontal")
	return glyphOrientationHorizontal
}

// GlyphOrientationHorizontal prop
func (*CSSStyleDeclaration) SetGlyphOrientationHorizontal(glyphOrientationHorizontal string) {
	js.Rewrite("$<.glyphOrientationHorizontal = $1", glyphOrientationHorizontal)
}

// GlyphOrientationVertical prop
func (*CSSStyleDeclaration) GlyphOrientationVertical() (glyphOrientationVertical string) {
	js.Rewrite("$<.glyphOrientationVertical")
	return glyphOrientationVertical
}

// GlyphOrientationVertical prop
func (*CSSStyleDeclaration) SetGlyphOrientationVertical(glyphOrientationVertical string) {
	js.Rewrite("$<.glyphOrientationVertical = $1", glyphOrientationVertical)
}

// Height prop
func (*CSSStyleDeclaration) Height() (height string) {
	js.Rewrite("$<.height")
	return height
}

// Height prop
func (*CSSStyleDeclaration) SetHeight(height string) {
	js.Rewrite("$<.height = $1", height)
}

// ImeMode prop
func (*CSSStyleDeclaration) ImeMode() (imeMode string) {
	js.Rewrite("$<.imeMode")
	return imeMode
}

// ImeMode prop
func (*CSSStyleDeclaration) SetImeMode(imeMode string) {
	js.Rewrite("$<.imeMode = $1", imeMode)
}

// JustifyContent prop
func (*CSSStyleDeclaration) JustifyContent() (justifyContent string) {
	js.Rewrite("$<.justifyContent")
	return justifyContent
}

// JustifyContent prop
func (*CSSStyleDeclaration) SetJustifyContent(justifyContent string) {
	js.Rewrite("$<.justifyContent = $1", justifyContent)
}

// Kerning prop
func (*CSSStyleDeclaration) Kerning() (kerning string) {
	js.Rewrite("$<.kerning")
	return kerning
}

// Kerning prop
func (*CSSStyleDeclaration) SetKerning(kerning string) {
	js.Rewrite("$<.kerning = $1", kerning)
}

// LayoutGrid prop
func (*CSSStyleDeclaration) LayoutGrid() (layoutGrid string) {
	js.Rewrite("$<.layoutGrid")
	return layoutGrid
}

// LayoutGrid prop
func (*CSSStyleDeclaration) SetLayoutGrid(layoutGrid string) {
	js.Rewrite("$<.layoutGrid = $1", layoutGrid)
}

// LayoutGridChar prop
func (*CSSStyleDeclaration) LayoutGridChar() (layoutGridChar string) {
	js.Rewrite("$<.layoutGridChar")
	return layoutGridChar
}

// LayoutGridChar prop
func (*CSSStyleDeclaration) SetLayoutGridChar(layoutGridChar string) {
	js.Rewrite("$<.layoutGridChar = $1", layoutGridChar)
}

// LayoutGridLine prop
func (*CSSStyleDeclaration) LayoutGridLine() (layoutGridLine string) {
	js.Rewrite("$<.layoutGridLine")
	return layoutGridLine
}

// LayoutGridLine prop
func (*CSSStyleDeclaration) SetLayoutGridLine(layoutGridLine string) {
	js.Rewrite("$<.layoutGridLine = $1", layoutGridLine)
}

// LayoutGridMode prop
func (*CSSStyleDeclaration) LayoutGridMode() (layoutGridMode string) {
	js.Rewrite("$<.layoutGridMode")
	return layoutGridMode
}

// LayoutGridMode prop
func (*CSSStyleDeclaration) SetLayoutGridMode(layoutGridMode string) {
	js.Rewrite("$<.layoutGridMode = $1", layoutGridMode)
}

// LayoutGridType prop
func (*CSSStyleDeclaration) LayoutGridType() (layoutGridType string) {
	js.Rewrite("$<.layoutGridType")
	return layoutGridType
}

// LayoutGridType prop
func (*CSSStyleDeclaration) SetLayoutGridType(layoutGridType string) {
	js.Rewrite("$<.layoutGridType = $1", layoutGridType)
}

// Left prop
func (*CSSStyleDeclaration) Left() (left string) {
	js.Rewrite("$<.left")
	return left
}

// Left prop
func (*CSSStyleDeclaration) SetLeft(left string) {
	js.Rewrite("$<.left = $1", left)
}

// Length prop
func (*CSSStyleDeclaration) Length() (length uint) {
	js.Rewrite("$<.length")
	return length
}

// LetterSpacing prop
func (*CSSStyleDeclaration) LetterSpacing() (letterSpacing string) {
	js.Rewrite("$<.letterSpacing")
	return letterSpacing
}

// LetterSpacing prop
func (*CSSStyleDeclaration) SetLetterSpacing(letterSpacing string) {
	js.Rewrite("$<.letterSpacing = $1", letterSpacing)
}

// LightingColor prop
func (*CSSStyleDeclaration) LightingColor() (lightingColor string) {
	js.Rewrite("$<.lightingColor")
	return lightingColor
}

// LightingColor prop
func (*CSSStyleDeclaration) SetLightingColor(lightingColor string) {
	js.Rewrite("$<.lightingColor = $1", lightingColor)
}

// LineBreak prop
func (*CSSStyleDeclaration) LineBreak() (lineBreak string) {
	js.Rewrite("$<.lineBreak")
	return lineBreak
}

// LineBreak prop
func (*CSSStyleDeclaration) SetLineBreak(lineBreak string) {
	js.Rewrite("$<.lineBreak = $1", lineBreak)
}

// LineHeight prop
func (*CSSStyleDeclaration) LineHeight() (lineHeight string) {
	js.Rewrite("$<.lineHeight")
	return lineHeight
}

// LineHeight prop
func (*CSSStyleDeclaration) SetLineHeight(lineHeight string) {
	js.Rewrite("$<.lineHeight = $1", lineHeight)
}

// ListStyle prop
func (*CSSStyleDeclaration) ListStyle() (listStyle string) {
	js.Rewrite("$<.listStyle")
	return listStyle
}

// ListStyle prop
func (*CSSStyleDeclaration) SetListStyle(listStyle string) {
	js.Rewrite("$<.listStyle = $1", listStyle)
}

// ListStyleImage prop
func (*CSSStyleDeclaration) ListStyleImage() (listStyleImage string) {
	js.Rewrite("$<.listStyleImage")
	return listStyleImage
}

// ListStyleImage prop
func (*CSSStyleDeclaration) SetListStyleImage(listStyleImage string) {
	js.Rewrite("$<.listStyleImage = $1", listStyleImage)
}

// ListStylePosition prop
func (*CSSStyleDeclaration) ListStylePosition() (listStylePosition string) {
	js.Rewrite("$<.listStylePosition")
	return listStylePosition
}

// ListStylePosition prop
func (*CSSStyleDeclaration) SetListStylePosition(listStylePosition string) {
	js.Rewrite("$<.listStylePosition = $1", listStylePosition)
}

// ListStyleType prop
func (*CSSStyleDeclaration) ListStyleType() (listStyleType string) {
	js.Rewrite("$<.listStyleType")
	return listStyleType
}

// ListStyleType prop
func (*CSSStyleDeclaration) SetListStyleType(listStyleType string) {
	js.Rewrite("$<.listStyleType = $1", listStyleType)
}

// Margin prop
func (*CSSStyleDeclaration) Margin() (margin string) {
	js.Rewrite("$<.margin")
	return margin
}

// Margin prop
func (*CSSStyleDeclaration) SetMargin(margin string) {
	js.Rewrite("$<.margin = $1", margin)
}

// MarginBottom prop
func (*CSSStyleDeclaration) MarginBottom() (marginBottom string) {
	js.Rewrite("$<.marginBottom")
	return marginBottom
}

// MarginBottom prop
func (*CSSStyleDeclaration) SetMarginBottom(marginBottom string) {
	js.Rewrite("$<.marginBottom = $1", marginBottom)
}

// MarginLeft prop
func (*CSSStyleDeclaration) MarginLeft() (marginLeft string) {
	js.Rewrite("$<.marginLeft")
	return marginLeft
}

// MarginLeft prop
func (*CSSStyleDeclaration) SetMarginLeft(marginLeft string) {
	js.Rewrite("$<.marginLeft = $1", marginLeft)
}

// MarginRight prop
func (*CSSStyleDeclaration) MarginRight() (marginRight string) {
	js.Rewrite("$<.marginRight")
	return marginRight
}

// MarginRight prop
func (*CSSStyleDeclaration) SetMarginRight(marginRight string) {
	js.Rewrite("$<.marginRight = $1", marginRight)
}

// MarginTop prop
func (*CSSStyleDeclaration) MarginTop() (marginTop string) {
	js.Rewrite("$<.marginTop")
	return marginTop
}

// MarginTop prop
func (*CSSStyleDeclaration) SetMarginTop(marginTop string) {
	js.Rewrite("$<.marginTop = $1", marginTop)
}

// Marker prop
func (*CSSStyleDeclaration) Marker() (marker string) {
	js.Rewrite("$<.marker")
	return marker
}

// Marker prop
func (*CSSStyleDeclaration) SetMarker(marker string) {
	js.Rewrite("$<.marker = $1", marker)
}

// MarkerEnd prop
func (*CSSStyleDeclaration) MarkerEnd() (markerEnd string) {
	js.Rewrite("$<.markerEnd")
	return markerEnd
}

// MarkerEnd prop
func (*CSSStyleDeclaration) SetMarkerEnd(markerEnd string) {
	js.Rewrite("$<.markerEnd = $1", markerEnd)
}

// MarkerMid prop
func (*CSSStyleDeclaration) MarkerMid() (markerMid string) {
	js.Rewrite("$<.markerMid")
	return markerMid
}

// MarkerMid prop
func (*CSSStyleDeclaration) SetMarkerMid(markerMid string) {
	js.Rewrite("$<.markerMid = $1", markerMid)
}

// MarkerStart prop
func (*CSSStyleDeclaration) MarkerStart() (markerStart string) {
	js.Rewrite("$<.markerStart")
	return markerStart
}

// MarkerStart prop
func (*CSSStyleDeclaration) SetMarkerStart(markerStart string) {
	js.Rewrite("$<.markerStart = $1", markerStart)
}

// Mask prop
func (*CSSStyleDeclaration) Mask() (mask string) {
	js.Rewrite("$<.mask")
	return mask
}

// Mask prop
func (*CSSStyleDeclaration) SetMask(mask string) {
	js.Rewrite("$<.mask = $1", mask)
}

// MaxHeight prop
func (*CSSStyleDeclaration) MaxHeight() (maxHeight string) {
	js.Rewrite("$<.maxHeight")
	return maxHeight
}

// MaxHeight prop
func (*CSSStyleDeclaration) SetMaxHeight(maxHeight string) {
	js.Rewrite("$<.maxHeight = $1", maxHeight)
}

// MaxWidth prop
func (*CSSStyleDeclaration) MaxWidth() (maxWidth string) {
	js.Rewrite("$<.maxWidth")
	return maxWidth
}

// MaxWidth prop
func (*CSSStyleDeclaration) SetMaxWidth(maxWidth string) {
	js.Rewrite("$<.maxWidth = $1", maxWidth)
}

// MinHeight prop
func (*CSSStyleDeclaration) MinHeight() (minHeight string) {
	js.Rewrite("$<.minHeight")
	return minHeight
}

// MinHeight prop
func (*CSSStyleDeclaration) SetMinHeight(minHeight string) {
	js.Rewrite("$<.minHeight = $1", minHeight)
}

// MinWidth prop
func (*CSSStyleDeclaration) MinWidth() (minWidth string) {
	js.Rewrite("$<.minWidth")
	return minWidth
}

// MinWidth prop
func (*CSSStyleDeclaration) SetMinWidth(minWidth string) {
	js.Rewrite("$<.minWidth = $1", minWidth)
}

// MsContentZoomChaining prop
func (*CSSStyleDeclaration) MsContentZoomChaining() (msContentZoomChaining string) {
	js.Rewrite("$<.msContentZoomChaining")
	return msContentZoomChaining
}

// MsContentZoomChaining prop
func (*CSSStyleDeclaration) SetMsContentZoomChaining(msContentZoomChaining string) {
	js.Rewrite("$<.msContentZoomChaining = $1", msContentZoomChaining)
}

// MsContentZooming prop
func (*CSSStyleDeclaration) MsContentZooming() (msContentZooming string) {
	js.Rewrite("$<.msContentZooming")
	return msContentZooming
}

// MsContentZooming prop
func (*CSSStyleDeclaration) SetMsContentZooming(msContentZooming string) {
	js.Rewrite("$<.msContentZooming = $1", msContentZooming)
}

// MsContentZoomLimit prop
func (*CSSStyleDeclaration) MsContentZoomLimit() (msContentZoomLimit string) {
	js.Rewrite("$<.msContentZoomLimit")
	return msContentZoomLimit
}

// MsContentZoomLimit prop
func (*CSSStyleDeclaration) SetMsContentZoomLimit(msContentZoomLimit string) {
	js.Rewrite("$<.msContentZoomLimit = $1", msContentZoomLimit)
}

// MsContentZoomLimitMax prop
func (*CSSStyleDeclaration) MsContentZoomLimitMax() (msContentZoomLimitMax interface{}) {
	js.Rewrite("$<.msContentZoomLimitMax")
	return msContentZoomLimitMax
}

// MsContentZoomLimitMax prop
func (*CSSStyleDeclaration) SetMsContentZoomLimitMax(msContentZoomLimitMax interface{}) {
	js.Rewrite("$<.msContentZoomLimitMax = $1", msContentZoomLimitMax)
}

// MsContentZoomLimitMin prop
func (*CSSStyleDeclaration) MsContentZoomLimitMin() (msContentZoomLimitMin interface{}) {
	js.Rewrite("$<.msContentZoomLimitMin")
	return msContentZoomLimitMin
}

// MsContentZoomLimitMin prop
func (*CSSStyleDeclaration) SetMsContentZoomLimitMin(msContentZoomLimitMin interface{}) {
	js.Rewrite("$<.msContentZoomLimitMin = $1", msContentZoomLimitMin)
}

// MsContentZoomSnap prop
func (*CSSStyleDeclaration) MsContentZoomSnap() (msContentZoomSnap string) {
	js.Rewrite("$<.msContentZoomSnap")
	return msContentZoomSnap
}

// MsContentZoomSnap prop
func (*CSSStyleDeclaration) SetMsContentZoomSnap(msContentZoomSnap string) {
	js.Rewrite("$<.msContentZoomSnap = $1", msContentZoomSnap)
}

// MsContentZoomSnapPoints prop
func (*CSSStyleDeclaration) MsContentZoomSnapPoints() (msContentZoomSnapPoints string) {
	js.Rewrite("$<.msContentZoomSnapPoints")
	return msContentZoomSnapPoints
}

// MsContentZoomSnapPoints prop
func (*CSSStyleDeclaration) SetMsContentZoomSnapPoints(msContentZoomSnapPoints string) {
	js.Rewrite("$<.msContentZoomSnapPoints = $1", msContentZoomSnapPoints)
}

// MsContentZoomSnapType prop
func (*CSSStyleDeclaration) MsContentZoomSnapType() (msContentZoomSnapType string) {
	js.Rewrite("$<.msContentZoomSnapType")
	return msContentZoomSnapType
}

// MsContentZoomSnapType prop
func (*CSSStyleDeclaration) SetMsContentZoomSnapType(msContentZoomSnapType string) {
	js.Rewrite("$<.msContentZoomSnapType = $1", msContentZoomSnapType)
}

// MsFlowFrom prop
func (*CSSStyleDeclaration) MsFlowFrom() (msFlowFrom string) {
	js.Rewrite("$<.msFlowFrom")
	return msFlowFrom
}

// MsFlowFrom prop
func (*CSSStyleDeclaration) SetMsFlowFrom(msFlowFrom string) {
	js.Rewrite("$<.msFlowFrom = $1", msFlowFrom)
}

// MsFlowInto prop
func (*CSSStyleDeclaration) MsFlowInto() (msFlowInto string) {
	js.Rewrite("$<.msFlowInto")
	return msFlowInto
}

// MsFlowInto prop
func (*CSSStyleDeclaration) SetMsFlowInto(msFlowInto string) {
	js.Rewrite("$<.msFlowInto = $1", msFlowInto)
}

// MsFontFeatureSettings prop
func (*CSSStyleDeclaration) MsFontFeatureSettings() (msFontFeatureSettings string) {
	js.Rewrite("$<.msFontFeatureSettings")
	return msFontFeatureSettings
}

// MsFontFeatureSettings prop
func (*CSSStyleDeclaration) SetMsFontFeatureSettings(msFontFeatureSettings string) {
	js.Rewrite("$<.msFontFeatureSettings = $1", msFontFeatureSettings)
}

// MsGridColumn prop
func (*CSSStyleDeclaration) MsGridColumn() (msGridColumn interface{}) {
	js.Rewrite("$<.msGridColumn")
	return msGridColumn
}

// MsGridColumn prop
func (*CSSStyleDeclaration) SetMsGridColumn(msGridColumn interface{}) {
	js.Rewrite("$<.msGridColumn = $1", msGridColumn)
}

// MsGridColumnAlign prop
func (*CSSStyleDeclaration) MsGridColumnAlign() (msGridColumnAlign string) {
	js.Rewrite("$<.msGridColumnAlign")
	return msGridColumnAlign
}

// MsGridColumnAlign prop
func (*CSSStyleDeclaration) SetMsGridColumnAlign(msGridColumnAlign string) {
	js.Rewrite("$<.msGridColumnAlign = $1", msGridColumnAlign)
}

// MsGridColumns prop
func (*CSSStyleDeclaration) MsGridColumns() (msGridColumns string) {
	js.Rewrite("$<.msGridColumns")
	return msGridColumns
}

// MsGridColumns prop
func (*CSSStyleDeclaration) SetMsGridColumns(msGridColumns string) {
	js.Rewrite("$<.msGridColumns = $1", msGridColumns)
}

// MsGridColumnSpan prop
func (*CSSStyleDeclaration) MsGridColumnSpan() (msGridColumnSpan interface{}) {
	js.Rewrite("$<.msGridColumnSpan")
	return msGridColumnSpan
}

// MsGridColumnSpan prop
func (*CSSStyleDeclaration) SetMsGridColumnSpan(msGridColumnSpan interface{}) {
	js.Rewrite("$<.msGridColumnSpan = $1", msGridColumnSpan)
}

// MsGridRow prop
func (*CSSStyleDeclaration) MsGridRow() (msGridRow interface{}) {
	js.Rewrite("$<.msGridRow")
	return msGridRow
}

// MsGridRow prop
func (*CSSStyleDeclaration) SetMsGridRow(msGridRow interface{}) {
	js.Rewrite("$<.msGridRow = $1", msGridRow)
}

// MsGridRowAlign prop
func (*CSSStyleDeclaration) MsGridRowAlign() (msGridRowAlign string) {
	js.Rewrite("$<.msGridRowAlign")
	return msGridRowAlign
}

// MsGridRowAlign prop
func (*CSSStyleDeclaration) SetMsGridRowAlign(msGridRowAlign string) {
	js.Rewrite("$<.msGridRowAlign = $1", msGridRowAlign)
}

// MsGridRows prop
func (*CSSStyleDeclaration) MsGridRows() (msGridRows string) {
	js.Rewrite("$<.msGridRows")
	return msGridRows
}

// MsGridRows prop
func (*CSSStyleDeclaration) SetMsGridRows(msGridRows string) {
	js.Rewrite("$<.msGridRows = $1", msGridRows)
}

// MsGridRowSpan prop
func (*CSSStyleDeclaration) MsGridRowSpan() (msGridRowSpan interface{}) {
	js.Rewrite("$<.msGridRowSpan")
	return msGridRowSpan
}

// MsGridRowSpan prop
func (*CSSStyleDeclaration) SetMsGridRowSpan(msGridRowSpan interface{}) {
	js.Rewrite("$<.msGridRowSpan = $1", msGridRowSpan)
}

// MsHighContrastAdjust prop
func (*CSSStyleDeclaration) MsHighContrastAdjust() (msHighContrastAdjust string) {
	js.Rewrite("$<.msHighContrastAdjust")
	return msHighContrastAdjust
}

// MsHighContrastAdjust prop
func (*CSSStyleDeclaration) SetMsHighContrastAdjust(msHighContrastAdjust string) {
	js.Rewrite("$<.msHighContrastAdjust = $1", msHighContrastAdjust)
}

// MsHyphenateLimitChars prop
func (*CSSStyleDeclaration) MsHyphenateLimitChars() (msHyphenateLimitChars string) {
	js.Rewrite("$<.msHyphenateLimitChars")
	return msHyphenateLimitChars
}

// MsHyphenateLimitChars prop
func (*CSSStyleDeclaration) SetMsHyphenateLimitChars(msHyphenateLimitChars string) {
	js.Rewrite("$<.msHyphenateLimitChars = $1", msHyphenateLimitChars)
}

// MsHyphenateLimitLines prop
func (*CSSStyleDeclaration) MsHyphenateLimitLines() (msHyphenateLimitLines interface{}) {
	js.Rewrite("$<.msHyphenateLimitLines")
	return msHyphenateLimitLines
}

// MsHyphenateLimitLines prop
func (*CSSStyleDeclaration) SetMsHyphenateLimitLines(msHyphenateLimitLines interface{}) {
	js.Rewrite("$<.msHyphenateLimitLines = $1", msHyphenateLimitLines)
}

// MsHyphenateLimitZone prop
func (*CSSStyleDeclaration) MsHyphenateLimitZone() (msHyphenateLimitZone interface{}) {
	js.Rewrite("$<.msHyphenateLimitZone")
	return msHyphenateLimitZone
}

// MsHyphenateLimitZone prop
func (*CSSStyleDeclaration) SetMsHyphenateLimitZone(msHyphenateLimitZone interface{}) {
	js.Rewrite("$<.msHyphenateLimitZone = $1", msHyphenateLimitZone)
}

// MsHyphens prop
func (*CSSStyleDeclaration) MsHyphens() (msHyphens string) {
	js.Rewrite("$<.msHyphens")
	return msHyphens
}

// MsHyphens prop
func (*CSSStyleDeclaration) SetMsHyphens(msHyphens string) {
	js.Rewrite("$<.msHyphens = $1", msHyphens)
}

// MsImeAlign prop
func (*CSSStyleDeclaration) MsImeAlign() (msImeAlign string) {
	js.Rewrite("$<.msImeAlign")
	return msImeAlign
}

// MsImeAlign prop
func (*CSSStyleDeclaration) SetMsImeAlign(msImeAlign string) {
	js.Rewrite("$<.msImeAlign = $1", msImeAlign)
}

// MsOverflowStyle prop
func (*CSSStyleDeclaration) MsOverflowStyle() (msOverflowStyle string) {
	js.Rewrite("$<.msOverflowStyle")
	return msOverflowStyle
}

// MsOverflowStyle prop
func (*CSSStyleDeclaration) SetMsOverflowStyle(msOverflowStyle string) {
	js.Rewrite("$<.msOverflowStyle = $1", msOverflowStyle)
}

// MsScrollChaining prop
func (*CSSStyleDeclaration) MsScrollChaining() (msScrollChaining string) {
	js.Rewrite("$<.msScrollChaining")
	return msScrollChaining
}

// MsScrollChaining prop
func (*CSSStyleDeclaration) SetMsScrollChaining(msScrollChaining string) {
	js.Rewrite("$<.msScrollChaining = $1", msScrollChaining)
}

// MsScrollLimit prop
func (*CSSStyleDeclaration) MsScrollLimit() (msScrollLimit string) {
	js.Rewrite("$<.msScrollLimit")
	return msScrollLimit
}

// MsScrollLimit prop
func (*CSSStyleDeclaration) SetMsScrollLimit(msScrollLimit string) {
	js.Rewrite("$<.msScrollLimit = $1", msScrollLimit)
}

// MsScrollLimitXMax prop
func (*CSSStyleDeclaration) MsScrollLimitXMax() (msScrollLimitXMax interface{}) {
	js.Rewrite("$<.msScrollLimitXMax")
	return msScrollLimitXMax
}

// MsScrollLimitXMax prop
func (*CSSStyleDeclaration) SetMsScrollLimitXMax(msScrollLimitXMax interface{}) {
	js.Rewrite("$<.msScrollLimitXMax = $1", msScrollLimitXMax)
}

// MsScrollLimitXMin prop
func (*CSSStyleDeclaration) MsScrollLimitXMin() (msScrollLimitXMin interface{}) {
	js.Rewrite("$<.msScrollLimitXMin")
	return msScrollLimitXMin
}

// MsScrollLimitXMin prop
func (*CSSStyleDeclaration) SetMsScrollLimitXMin(msScrollLimitXMin interface{}) {
	js.Rewrite("$<.msScrollLimitXMin = $1", msScrollLimitXMin)
}

// MsScrollLimitYMax prop
func (*CSSStyleDeclaration) MsScrollLimitYMax() (msScrollLimitYMax interface{}) {
	js.Rewrite("$<.msScrollLimitYMax")
	return msScrollLimitYMax
}

// MsScrollLimitYMax prop
func (*CSSStyleDeclaration) SetMsScrollLimitYMax(msScrollLimitYMax interface{}) {
	js.Rewrite("$<.msScrollLimitYMax = $1", msScrollLimitYMax)
}

// MsScrollLimitYMin prop
func (*CSSStyleDeclaration) MsScrollLimitYMin() (msScrollLimitYMin interface{}) {
	js.Rewrite("$<.msScrollLimitYMin")
	return msScrollLimitYMin
}

// MsScrollLimitYMin prop
func (*CSSStyleDeclaration) SetMsScrollLimitYMin(msScrollLimitYMin interface{}) {
	js.Rewrite("$<.msScrollLimitYMin = $1", msScrollLimitYMin)
}

// MsScrollRails prop
func (*CSSStyleDeclaration) MsScrollRails() (msScrollRails string) {
	js.Rewrite("$<.msScrollRails")
	return msScrollRails
}

// MsScrollRails prop
func (*CSSStyleDeclaration) SetMsScrollRails(msScrollRails string) {
	js.Rewrite("$<.msScrollRails = $1", msScrollRails)
}

// MsScrollSnapPointsX prop
func (*CSSStyleDeclaration) MsScrollSnapPointsX() (msScrollSnapPointsX string) {
	js.Rewrite("$<.msScrollSnapPointsX")
	return msScrollSnapPointsX
}

// MsScrollSnapPointsX prop
func (*CSSStyleDeclaration) SetMsScrollSnapPointsX(msScrollSnapPointsX string) {
	js.Rewrite("$<.msScrollSnapPointsX = $1", msScrollSnapPointsX)
}

// MsScrollSnapPointsY prop
func (*CSSStyleDeclaration) MsScrollSnapPointsY() (msScrollSnapPointsY string) {
	js.Rewrite("$<.msScrollSnapPointsY")
	return msScrollSnapPointsY
}

// MsScrollSnapPointsY prop
func (*CSSStyleDeclaration) SetMsScrollSnapPointsY(msScrollSnapPointsY string) {
	js.Rewrite("$<.msScrollSnapPointsY = $1", msScrollSnapPointsY)
}

// MsScrollSnapType prop
func (*CSSStyleDeclaration) MsScrollSnapType() (msScrollSnapType string) {
	js.Rewrite("$<.msScrollSnapType")
	return msScrollSnapType
}

// MsScrollSnapType prop
func (*CSSStyleDeclaration) SetMsScrollSnapType(msScrollSnapType string) {
	js.Rewrite("$<.msScrollSnapType = $1", msScrollSnapType)
}

// MsScrollSnapX prop
func (*CSSStyleDeclaration) MsScrollSnapX() (msScrollSnapX string) {
	js.Rewrite("$<.msScrollSnapX")
	return msScrollSnapX
}

// MsScrollSnapX prop
func (*CSSStyleDeclaration) SetMsScrollSnapX(msScrollSnapX string) {
	js.Rewrite("$<.msScrollSnapX = $1", msScrollSnapX)
}

// MsScrollSnapY prop
func (*CSSStyleDeclaration) MsScrollSnapY() (msScrollSnapY string) {
	js.Rewrite("$<.msScrollSnapY")
	return msScrollSnapY
}

// MsScrollSnapY prop
func (*CSSStyleDeclaration) SetMsScrollSnapY(msScrollSnapY string) {
	js.Rewrite("$<.msScrollSnapY = $1", msScrollSnapY)
}

// MsScrollTranslation prop
func (*CSSStyleDeclaration) MsScrollTranslation() (msScrollTranslation string) {
	js.Rewrite("$<.msScrollTranslation")
	return msScrollTranslation
}

// MsScrollTranslation prop
func (*CSSStyleDeclaration) SetMsScrollTranslation(msScrollTranslation string) {
	js.Rewrite("$<.msScrollTranslation = $1", msScrollTranslation)
}

// MsTextCombineHorizontal prop
func (*CSSStyleDeclaration) MsTextCombineHorizontal() (msTextCombineHorizontal string) {
	js.Rewrite("$<.msTextCombineHorizontal")
	return msTextCombineHorizontal
}

// MsTextCombineHorizontal prop
func (*CSSStyleDeclaration) SetMsTextCombineHorizontal(msTextCombineHorizontal string) {
	js.Rewrite("$<.msTextCombineHorizontal = $1", msTextCombineHorizontal)
}

// MsTextSizeAdjust prop
func (*CSSStyleDeclaration) MsTextSizeAdjust() (msTextSizeAdjust interface{}) {
	js.Rewrite("$<.msTextSizeAdjust")
	return msTextSizeAdjust
}

// MsTextSizeAdjust prop
func (*CSSStyleDeclaration) SetMsTextSizeAdjust(msTextSizeAdjust interface{}) {
	js.Rewrite("$<.msTextSizeAdjust = $1", msTextSizeAdjust)
}

// MsTouchAction prop
func (*CSSStyleDeclaration) MsTouchAction() (msTouchAction string) {
	js.Rewrite("$<.msTouchAction")
	return msTouchAction
}

// MsTouchAction prop
func (*CSSStyleDeclaration) SetMsTouchAction(msTouchAction string) {
	js.Rewrite("$<.msTouchAction = $1", msTouchAction)
}

// MsTouchSelect prop
func (*CSSStyleDeclaration) MsTouchSelect() (msTouchSelect string) {
	js.Rewrite("$<.msTouchSelect")
	return msTouchSelect
}

// MsTouchSelect prop
func (*CSSStyleDeclaration) SetMsTouchSelect(msTouchSelect string) {
	js.Rewrite("$<.msTouchSelect = $1", msTouchSelect)
}

// MsUserSelect prop
func (*CSSStyleDeclaration) MsUserSelect() (msUserSelect string) {
	js.Rewrite("$<.msUserSelect")
	return msUserSelect
}

// MsUserSelect prop
func (*CSSStyleDeclaration) SetMsUserSelect(msUserSelect string) {
	js.Rewrite("$<.msUserSelect = $1", msUserSelect)
}

// MsWrapFlow prop
func (*CSSStyleDeclaration) MsWrapFlow() (msWrapFlow string) {
	js.Rewrite("$<.msWrapFlow")
	return msWrapFlow
}

// MsWrapFlow prop
func (*CSSStyleDeclaration) SetMsWrapFlow(msWrapFlow string) {
	js.Rewrite("$<.msWrapFlow = $1", msWrapFlow)
}

// MsWrapMargin prop
func (*CSSStyleDeclaration) MsWrapMargin() (msWrapMargin interface{}) {
	js.Rewrite("$<.msWrapMargin")
	return msWrapMargin
}

// MsWrapMargin prop
func (*CSSStyleDeclaration) SetMsWrapMargin(msWrapMargin interface{}) {
	js.Rewrite("$<.msWrapMargin = $1", msWrapMargin)
}

// MsWrapThrough prop
func (*CSSStyleDeclaration) MsWrapThrough() (msWrapThrough string) {
	js.Rewrite("$<.msWrapThrough")
	return msWrapThrough
}

// MsWrapThrough prop
func (*CSSStyleDeclaration) SetMsWrapThrough(msWrapThrough string) {
	js.Rewrite("$<.msWrapThrough = $1", msWrapThrough)
}

// Opacity prop
func (*CSSStyleDeclaration) Opacity() (opacity string) {
	js.Rewrite("$<.opacity")
	return opacity
}

// Opacity prop
func (*CSSStyleDeclaration) SetOpacity(opacity string) {
	js.Rewrite("$<.opacity = $1", opacity)
}

// Order prop
func (*CSSStyleDeclaration) Order() (order string) {
	js.Rewrite("$<.order")
	return order
}

// Order prop
func (*CSSStyleDeclaration) SetOrder(order string) {
	js.Rewrite("$<.order = $1", order)
}

// Orphans prop
func (*CSSStyleDeclaration) Orphans() (orphans string) {
	js.Rewrite("$<.orphans")
	return orphans
}

// Orphans prop
func (*CSSStyleDeclaration) SetOrphans(orphans string) {
	js.Rewrite("$<.orphans = $1", orphans)
}

// Outline prop
func (*CSSStyleDeclaration) Outline() (outline string) {
	js.Rewrite("$<.outline")
	return outline
}

// Outline prop
func (*CSSStyleDeclaration) SetOutline(outline string) {
	js.Rewrite("$<.outline = $1", outline)
}

// OutlineColor prop
func (*CSSStyleDeclaration) OutlineColor() (outlineColor string) {
	js.Rewrite("$<.outlineColor")
	return outlineColor
}

// OutlineColor prop
func (*CSSStyleDeclaration) SetOutlineColor(outlineColor string) {
	js.Rewrite("$<.outlineColor = $1", outlineColor)
}

// OutlineOffset prop
func (*CSSStyleDeclaration) OutlineOffset() (outlineOffset string) {
	js.Rewrite("$<.outlineOffset")
	return outlineOffset
}

// OutlineOffset prop
func (*CSSStyleDeclaration) SetOutlineOffset(outlineOffset string) {
	js.Rewrite("$<.outlineOffset = $1", outlineOffset)
}

// OutlineStyle prop
func (*CSSStyleDeclaration) OutlineStyle() (outlineStyle string) {
	js.Rewrite("$<.outlineStyle")
	return outlineStyle
}

// OutlineStyle prop
func (*CSSStyleDeclaration) SetOutlineStyle(outlineStyle string) {
	js.Rewrite("$<.outlineStyle = $1", outlineStyle)
}

// OutlineWidth prop
func (*CSSStyleDeclaration) OutlineWidth() (outlineWidth string) {
	js.Rewrite("$<.outlineWidth")
	return outlineWidth
}

// OutlineWidth prop
func (*CSSStyleDeclaration) SetOutlineWidth(outlineWidth string) {
	js.Rewrite("$<.outlineWidth = $1", outlineWidth)
}

// Overflow prop
func (*CSSStyleDeclaration) Overflow() (overflow string) {
	js.Rewrite("$<.overflow")
	return overflow
}

// Overflow prop
func (*CSSStyleDeclaration) SetOverflow(overflow string) {
	js.Rewrite("$<.overflow = $1", overflow)
}

// OverflowX prop
func (*CSSStyleDeclaration) OverflowX() (overflowX string) {
	js.Rewrite("$<.overflowX")
	return overflowX
}

// OverflowX prop
func (*CSSStyleDeclaration) SetOverflowX(overflowX string) {
	js.Rewrite("$<.overflowX = $1", overflowX)
}

// OverflowY prop
func (*CSSStyleDeclaration) OverflowY() (overflowY string) {
	js.Rewrite("$<.overflowY")
	return overflowY
}

// OverflowY prop
func (*CSSStyleDeclaration) SetOverflowY(overflowY string) {
	js.Rewrite("$<.overflowY = $1", overflowY)
}

// Padding prop
func (*CSSStyleDeclaration) Padding() (padding string) {
	js.Rewrite("$<.padding")
	return padding
}

// Padding prop
func (*CSSStyleDeclaration) SetPadding(padding string) {
	js.Rewrite("$<.padding = $1", padding)
}

// PaddingBottom prop
func (*CSSStyleDeclaration) PaddingBottom() (paddingBottom string) {
	js.Rewrite("$<.paddingBottom")
	return paddingBottom
}

// PaddingBottom prop
func (*CSSStyleDeclaration) SetPaddingBottom(paddingBottom string) {
	js.Rewrite("$<.paddingBottom = $1", paddingBottom)
}

// PaddingLeft prop
func (*CSSStyleDeclaration) PaddingLeft() (paddingLeft string) {
	js.Rewrite("$<.paddingLeft")
	return paddingLeft
}

// PaddingLeft prop
func (*CSSStyleDeclaration) SetPaddingLeft(paddingLeft string) {
	js.Rewrite("$<.paddingLeft = $1", paddingLeft)
}

// PaddingRight prop
func (*CSSStyleDeclaration) PaddingRight() (paddingRight string) {
	js.Rewrite("$<.paddingRight")
	return paddingRight
}

// PaddingRight prop
func (*CSSStyleDeclaration) SetPaddingRight(paddingRight string) {
	js.Rewrite("$<.paddingRight = $1", paddingRight)
}

// PaddingTop prop
func (*CSSStyleDeclaration) PaddingTop() (paddingTop string) {
	js.Rewrite("$<.paddingTop")
	return paddingTop
}

// PaddingTop prop
func (*CSSStyleDeclaration) SetPaddingTop(paddingTop string) {
	js.Rewrite("$<.paddingTop = $1", paddingTop)
}

// PageBreakAfter prop
func (*CSSStyleDeclaration) PageBreakAfter() (pageBreakAfter string) {
	js.Rewrite("$<.pageBreakAfter")
	return pageBreakAfter
}

// PageBreakAfter prop
func (*CSSStyleDeclaration) SetPageBreakAfter(pageBreakAfter string) {
	js.Rewrite("$<.pageBreakAfter = $1", pageBreakAfter)
}

// PageBreakBefore prop
func (*CSSStyleDeclaration) PageBreakBefore() (pageBreakBefore string) {
	js.Rewrite("$<.pageBreakBefore")
	return pageBreakBefore
}

// PageBreakBefore prop
func (*CSSStyleDeclaration) SetPageBreakBefore(pageBreakBefore string) {
	js.Rewrite("$<.pageBreakBefore = $1", pageBreakBefore)
}

// PageBreakInside prop
func (*CSSStyleDeclaration) PageBreakInside() (pageBreakInside string) {
	js.Rewrite("$<.pageBreakInside")
	return pageBreakInside
}

// PageBreakInside prop
func (*CSSStyleDeclaration) SetPageBreakInside(pageBreakInside string) {
	js.Rewrite("$<.pageBreakInside = $1", pageBreakInside)
}

// ParentRule prop
func (*CSSStyleDeclaration) ParentRule() (parentRule CSSRule) {
	js.Rewrite("$<.parentRule")
	return parentRule
}

// Perspective prop
func (*CSSStyleDeclaration) Perspective() (perspective string) {
	js.Rewrite("$<.perspective")
	return perspective
}

// Perspective prop
func (*CSSStyleDeclaration) SetPerspective(perspective string) {
	js.Rewrite("$<.perspective = $1", perspective)
}

// PerspectiveOrigin prop
func (*CSSStyleDeclaration) PerspectiveOrigin() (perspectiveOrigin string) {
	js.Rewrite("$<.perspectiveOrigin")
	return perspectiveOrigin
}

// PerspectiveOrigin prop
func (*CSSStyleDeclaration) SetPerspectiveOrigin(perspectiveOrigin string) {
	js.Rewrite("$<.perspectiveOrigin = $1", perspectiveOrigin)
}

// PointerEvents prop
func (*CSSStyleDeclaration) PointerEvents() (pointerEvents string) {
	js.Rewrite("$<.pointerEvents")
	return pointerEvents
}

// PointerEvents prop
func (*CSSStyleDeclaration) SetPointerEvents(pointerEvents string) {
	js.Rewrite("$<.pointerEvents = $1", pointerEvents)
}

// Position prop
func (*CSSStyleDeclaration) Position() (position string) {
	js.Rewrite("$<.position")
	return position
}

// Position prop
func (*CSSStyleDeclaration) SetPosition(position string) {
	js.Rewrite("$<.position = $1", position)
}

// Quotes prop
func (*CSSStyleDeclaration) Quotes() (quotes string) {
	js.Rewrite("$<.quotes")
	return quotes
}

// Quotes prop
func (*CSSStyleDeclaration) SetQuotes(quotes string) {
	js.Rewrite("$<.quotes = $1", quotes)
}

// Right prop
func (*CSSStyleDeclaration) Right() (right string) {
	js.Rewrite("$<.right")
	return right
}

// Right prop
func (*CSSStyleDeclaration) SetRight(right string) {
	js.Rewrite("$<.right = $1", right)
}

// Rotate prop
func (*CSSStyleDeclaration) Rotate() (rotate string) {
	js.Rewrite("$<.rotate")
	return rotate
}

// Rotate prop
func (*CSSStyleDeclaration) SetRotate(rotate string) {
	js.Rewrite("$<.rotate = $1", rotate)
}

// RubyAlign prop
func (*CSSStyleDeclaration) RubyAlign() (rubyAlign string) {
	js.Rewrite("$<.rubyAlign")
	return rubyAlign
}

// RubyAlign prop
func (*CSSStyleDeclaration) SetRubyAlign(rubyAlign string) {
	js.Rewrite("$<.rubyAlign = $1", rubyAlign)
}

// RubyOverhang prop
func (*CSSStyleDeclaration) RubyOverhang() (rubyOverhang string) {
	js.Rewrite("$<.rubyOverhang")
	return rubyOverhang
}

// RubyOverhang prop
func (*CSSStyleDeclaration) SetRubyOverhang(rubyOverhang string) {
	js.Rewrite("$<.rubyOverhang = $1", rubyOverhang)
}

// RubyPosition prop
func (*CSSStyleDeclaration) RubyPosition() (rubyPosition string) {
	js.Rewrite("$<.rubyPosition")
	return rubyPosition
}

// RubyPosition prop
func (*CSSStyleDeclaration) SetRubyPosition(rubyPosition string) {
	js.Rewrite("$<.rubyPosition = $1", rubyPosition)
}

// Scale prop
func (*CSSStyleDeclaration) Scale() (scale string) {
	js.Rewrite("$<.scale")
	return scale
}

// Scale prop
func (*CSSStyleDeclaration) SetScale(scale string) {
	js.Rewrite("$<.scale = $1", scale)
}

// StopColor prop
func (*CSSStyleDeclaration) StopColor() (stopColor string) {
	js.Rewrite("$<.stopColor")
	return stopColor
}

// StopColor prop
func (*CSSStyleDeclaration) SetStopColor(stopColor string) {
	js.Rewrite("$<.stopColor = $1", stopColor)
}

// StopOpacity prop
func (*CSSStyleDeclaration) StopOpacity() (stopOpacity string) {
	js.Rewrite("$<.stopOpacity")
	return stopOpacity
}

// StopOpacity prop
func (*CSSStyleDeclaration) SetStopOpacity(stopOpacity string) {
	js.Rewrite("$<.stopOpacity = $1", stopOpacity)
}

// Stroke prop
func (*CSSStyleDeclaration) Stroke() (stroke string) {
	js.Rewrite("$<.stroke")
	return stroke
}

// Stroke prop
func (*CSSStyleDeclaration) SetStroke(stroke string) {
	js.Rewrite("$<.stroke = $1", stroke)
}

// StrokeDasharray prop
func (*CSSStyleDeclaration) StrokeDasharray() (strokeDasharray string) {
	js.Rewrite("$<.strokeDasharray")
	return strokeDasharray
}

// StrokeDasharray prop
func (*CSSStyleDeclaration) SetStrokeDasharray(strokeDasharray string) {
	js.Rewrite("$<.strokeDasharray = $1", strokeDasharray)
}

// StrokeDashoffset prop
func (*CSSStyleDeclaration) StrokeDashoffset() (strokeDashoffset string) {
	js.Rewrite("$<.strokeDashoffset")
	return strokeDashoffset
}

// StrokeDashoffset prop
func (*CSSStyleDeclaration) SetStrokeDashoffset(strokeDashoffset string) {
	js.Rewrite("$<.strokeDashoffset = $1", strokeDashoffset)
}

// StrokeLinecap prop
func (*CSSStyleDeclaration) StrokeLinecap() (strokeLinecap string) {
	js.Rewrite("$<.strokeLinecap")
	return strokeLinecap
}

// StrokeLinecap prop
func (*CSSStyleDeclaration) SetStrokeLinecap(strokeLinecap string) {
	js.Rewrite("$<.strokeLinecap = $1", strokeLinecap)
}

// StrokeLinejoin prop
func (*CSSStyleDeclaration) StrokeLinejoin() (strokeLinejoin string) {
	js.Rewrite("$<.strokeLinejoin")
	return strokeLinejoin
}

// StrokeLinejoin prop
func (*CSSStyleDeclaration) SetStrokeLinejoin(strokeLinejoin string) {
	js.Rewrite("$<.strokeLinejoin = $1", strokeLinejoin)
}

// StrokeMiterlimit prop
func (*CSSStyleDeclaration) StrokeMiterlimit() (strokeMiterlimit string) {
	js.Rewrite("$<.strokeMiterlimit")
	return strokeMiterlimit
}

// StrokeMiterlimit prop
func (*CSSStyleDeclaration) SetStrokeMiterlimit(strokeMiterlimit string) {
	js.Rewrite("$<.strokeMiterlimit = $1", strokeMiterlimit)
}

// StrokeOpacity prop
func (*CSSStyleDeclaration) StrokeOpacity() (strokeOpacity string) {
	js.Rewrite("$<.strokeOpacity")
	return strokeOpacity
}

// StrokeOpacity prop
func (*CSSStyleDeclaration) SetStrokeOpacity(strokeOpacity string) {
	js.Rewrite("$<.strokeOpacity = $1", strokeOpacity)
}

// StrokeWidth prop
func (*CSSStyleDeclaration) StrokeWidth() (strokeWidth string) {
	js.Rewrite("$<.strokeWidth")
	return strokeWidth
}

// StrokeWidth prop
func (*CSSStyleDeclaration) SetStrokeWidth(strokeWidth string) {
	js.Rewrite("$<.strokeWidth = $1", strokeWidth)
}

// TableLayout prop
func (*CSSStyleDeclaration) TableLayout() (tableLayout string) {
	js.Rewrite("$<.tableLayout")
	return tableLayout
}

// TableLayout prop
func (*CSSStyleDeclaration) SetTableLayout(tableLayout string) {
	js.Rewrite("$<.tableLayout = $1", tableLayout)
}

// TextAlign prop
func (*CSSStyleDeclaration) TextAlign() (textAlign string) {
	js.Rewrite("$<.textAlign")
	return textAlign
}

// TextAlign prop
func (*CSSStyleDeclaration) SetTextAlign(textAlign string) {
	js.Rewrite("$<.textAlign = $1", textAlign)
}

// TextAlignLast prop
func (*CSSStyleDeclaration) TextAlignLast() (textAlignLast string) {
	js.Rewrite("$<.textAlignLast")
	return textAlignLast
}

// TextAlignLast prop
func (*CSSStyleDeclaration) SetTextAlignLast(textAlignLast string) {
	js.Rewrite("$<.textAlignLast = $1", textAlignLast)
}

// TextAnchor prop
func (*CSSStyleDeclaration) TextAnchor() (textAnchor string) {
	js.Rewrite("$<.textAnchor")
	return textAnchor
}

// TextAnchor prop
func (*CSSStyleDeclaration) SetTextAnchor(textAnchor string) {
	js.Rewrite("$<.textAnchor = $1", textAnchor)
}

// TextDecoration prop
func (*CSSStyleDeclaration) TextDecoration() (textDecoration string) {
	js.Rewrite("$<.textDecoration")
	return textDecoration
}

// TextDecoration prop
func (*CSSStyleDeclaration) SetTextDecoration(textDecoration string) {
	js.Rewrite("$<.textDecoration = $1", textDecoration)
}

// TextIndent prop
func (*CSSStyleDeclaration) TextIndent() (textIndent string) {
	js.Rewrite("$<.textIndent")
	return textIndent
}

// TextIndent prop
func (*CSSStyleDeclaration) SetTextIndent(textIndent string) {
	js.Rewrite("$<.textIndent = $1", textIndent)
}

// TextJustify prop
func (*CSSStyleDeclaration) TextJustify() (textJustify string) {
	js.Rewrite("$<.textJustify")
	return textJustify
}

// TextJustify prop
func (*CSSStyleDeclaration) SetTextJustify(textJustify string) {
	js.Rewrite("$<.textJustify = $1", textJustify)
}

// TextKashida prop
func (*CSSStyleDeclaration) TextKashida() (textKashida string) {
	js.Rewrite("$<.textKashida")
	return textKashida
}

// TextKashida prop
func (*CSSStyleDeclaration) SetTextKashida(textKashida string) {
	js.Rewrite("$<.textKashida = $1", textKashida)
}

// TextKashidaSpace prop
func (*CSSStyleDeclaration) TextKashidaSpace() (textKashidaSpace string) {
	js.Rewrite("$<.textKashidaSpace")
	return textKashidaSpace
}

// TextKashidaSpace prop
func (*CSSStyleDeclaration) SetTextKashidaSpace(textKashidaSpace string) {
	js.Rewrite("$<.textKashidaSpace = $1", textKashidaSpace)
}

// TextOverflow prop
func (*CSSStyleDeclaration) TextOverflow() (textOverflow string) {
	js.Rewrite("$<.textOverflow")
	return textOverflow
}

// TextOverflow prop
func (*CSSStyleDeclaration) SetTextOverflow(textOverflow string) {
	js.Rewrite("$<.textOverflow = $1", textOverflow)
}

// TextShadow prop
func (*CSSStyleDeclaration) TextShadow() (textShadow string) {
	js.Rewrite("$<.textShadow")
	return textShadow
}

// TextShadow prop
func (*CSSStyleDeclaration) SetTextShadow(textShadow string) {
	js.Rewrite("$<.textShadow = $1", textShadow)
}

// TextTransform prop
func (*CSSStyleDeclaration) TextTransform() (textTransform string) {
	js.Rewrite("$<.textTransform")
	return textTransform
}

// TextTransform prop
func (*CSSStyleDeclaration) SetTextTransform(textTransform string) {
	js.Rewrite("$<.textTransform = $1", textTransform)
}

// TextUnderlinePosition prop
func (*CSSStyleDeclaration) TextUnderlinePosition() (textUnderlinePosition string) {
	js.Rewrite("$<.textUnderlinePosition")
	return textUnderlinePosition
}

// TextUnderlinePosition prop
func (*CSSStyleDeclaration) SetTextUnderlinePosition(textUnderlinePosition string) {
	js.Rewrite("$<.textUnderlinePosition = $1", textUnderlinePosition)
}

// Top prop
func (*CSSStyleDeclaration) Top() (top string) {
	js.Rewrite("$<.top")
	return top
}

// Top prop
func (*CSSStyleDeclaration) SetTop(top string) {
	js.Rewrite("$<.top = $1", top)
}

// TouchAction prop
func (*CSSStyleDeclaration) TouchAction() (touchAction string) {
	js.Rewrite("$<.touchAction")
	return touchAction
}

// TouchAction prop
func (*CSSStyleDeclaration) SetTouchAction(touchAction string) {
	js.Rewrite("$<.touchAction = $1", touchAction)
}

// Transform prop
func (*CSSStyleDeclaration) Transform() (transform string) {
	js.Rewrite("$<.transform")
	return transform
}

// Transform prop
func (*CSSStyleDeclaration) SetTransform(transform string) {
	js.Rewrite("$<.transform = $1", transform)
}

// TransformOrigin prop
func (*CSSStyleDeclaration) TransformOrigin() (transformOrigin string) {
	js.Rewrite("$<.transformOrigin")
	return transformOrigin
}

// TransformOrigin prop
func (*CSSStyleDeclaration) SetTransformOrigin(transformOrigin string) {
	js.Rewrite("$<.transformOrigin = $1", transformOrigin)
}

// TransformStyle prop
func (*CSSStyleDeclaration) TransformStyle() (transformStyle string) {
	js.Rewrite("$<.transformStyle")
	return transformStyle
}

// TransformStyle prop
func (*CSSStyleDeclaration) SetTransformStyle(transformStyle string) {
	js.Rewrite("$<.transformStyle = $1", transformStyle)
}

// Transition prop
func (*CSSStyleDeclaration) Transition() (transition string) {
	js.Rewrite("$<.transition")
	return transition
}

// Transition prop
func (*CSSStyleDeclaration) SetTransition(transition string) {
	js.Rewrite("$<.transition = $1", transition)
}

// TransitionDelay prop
func (*CSSStyleDeclaration) TransitionDelay() (transitionDelay string) {
	js.Rewrite("$<.transitionDelay")
	return transitionDelay
}

// TransitionDelay prop
func (*CSSStyleDeclaration) SetTransitionDelay(transitionDelay string) {
	js.Rewrite("$<.transitionDelay = $1", transitionDelay)
}

// TransitionDuration prop
func (*CSSStyleDeclaration) TransitionDuration() (transitionDuration string) {
	js.Rewrite("$<.transitionDuration")
	return transitionDuration
}

// TransitionDuration prop
func (*CSSStyleDeclaration) SetTransitionDuration(transitionDuration string) {
	js.Rewrite("$<.transitionDuration = $1", transitionDuration)
}

// TransitionProperty prop
func (*CSSStyleDeclaration) TransitionProperty() (transitionProperty string) {
	js.Rewrite("$<.transitionProperty")
	return transitionProperty
}

// TransitionProperty prop
func (*CSSStyleDeclaration) SetTransitionProperty(transitionProperty string) {
	js.Rewrite("$<.transitionProperty = $1", transitionProperty)
}

// TransitionTimingFunction prop
func (*CSSStyleDeclaration) TransitionTimingFunction() (transitionTimingFunction string) {
	js.Rewrite("$<.transitionTimingFunction")
	return transitionTimingFunction
}

// TransitionTimingFunction prop
func (*CSSStyleDeclaration) SetTransitionTimingFunction(transitionTimingFunction string) {
	js.Rewrite("$<.transitionTimingFunction = $1", transitionTimingFunction)
}

// Translate prop
func (*CSSStyleDeclaration) Translate() (translate string) {
	js.Rewrite("$<.translate")
	return translate
}

// Translate prop
func (*CSSStyleDeclaration) SetTranslate(translate string) {
	js.Rewrite("$<.translate = $1", translate)
}

// UnicodeBidi prop
func (*CSSStyleDeclaration) UnicodeBidi() (unicodeBidi string) {
	js.Rewrite("$<.unicodeBidi")
	return unicodeBidi
}

// UnicodeBidi prop
func (*CSSStyleDeclaration) SetUnicodeBidi(unicodeBidi string) {
	js.Rewrite("$<.unicodeBidi = $1", unicodeBidi)
}

// VerticalAlign prop
func (*CSSStyleDeclaration) VerticalAlign() (verticalAlign string) {
	js.Rewrite("$<.verticalAlign")
	return verticalAlign
}

// VerticalAlign prop
func (*CSSStyleDeclaration) SetVerticalAlign(verticalAlign string) {
	js.Rewrite("$<.verticalAlign = $1", verticalAlign)
}

// Visibility prop
func (*CSSStyleDeclaration) Visibility() (visibility string) {
	js.Rewrite("$<.visibility")
	return visibility
}

// Visibility prop
func (*CSSStyleDeclaration) SetVisibility(visibility string) {
	js.Rewrite("$<.visibility = $1", visibility)
}

// WebkitAlignContent prop
func (*CSSStyleDeclaration) WebkitAlignContent() (webkitAlignContent string) {
	js.Rewrite("$<.webkitAlignContent")
	return webkitAlignContent
}

// WebkitAlignContent prop
func (*CSSStyleDeclaration) SetWebkitAlignContent(webkitAlignContent string) {
	js.Rewrite("$<.webkitAlignContent = $1", webkitAlignContent)
}

// WebkitAlignItems prop
func (*CSSStyleDeclaration) WebkitAlignItems() (webkitAlignItems string) {
	js.Rewrite("$<.webkitAlignItems")
	return webkitAlignItems
}

// WebkitAlignItems prop
func (*CSSStyleDeclaration) SetWebkitAlignItems(webkitAlignItems string) {
	js.Rewrite("$<.webkitAlignItems = $1", webkitAlignItems)
}

// WebkitAlignSelf prop
func (*CSSStyleDeclaration) WebkitAlignSelf() (webkitAlignSelf string) {
	js.Rewrite("$<.webkitAlignSelf")
	return webkitAlignSelf
}

// WebkitAlignSelf prop
func (*CSSStyleDeclaration) SetWebkitAlignSelf(webkitAlignSelf string) {
	js.Rewrite("$<.webkitAlignSelf = $1", webkitAlignSelf)
}

// WebkitAnimation prop
func (*CSSStyleDeclaration) WebkitAnimation() (webkitAnimation string) {
	js.Rewrite("$<.webkitAnimation")
	return webkitAnimation
}

// WebkitAnimation prop
func (*CSSStyleDeclaration) SetWebkitAnimation(webkitAnimation string) {
	js.Rewrite("$<.webkitAnimation = $1", webkitAnimation)
}

// WebkitAnimationDelay prop
func (*CSSStyleDeclaration) WebkitAnimationDelay() (webkitAnimationDelay string) {
	js.Rewrite("$<.webkitAnimationDelay")
	return webkitAnimationDelay
}

// WebkitAnimationDelay prop
func (*CSSStyleDeclaration) SetWebkitAnimationDelay(webkitAnimationDelay string) {
	js.Rewrite("$<.webkitAnimationDelay = $1", webkitAnimationDelay)
}

// WebkitAnimationDirection prop
func (*CSSStyleDeclaration) WebkitAnimationDirection() (webkitAnimationDirection string) {
	js.Rewrite("$<.webkitAnimationDirection")
	return webkitAnimationDirection
}

// WebkitAnimationDirection prop
func (*CSSStyleDeclaration) SetWebkitAnimationDirection(webkitAnimationDirection string) {
	js.Rewrite("$<.webkitAnimationDirection = $1", webkitAnimationDirection)
}

// WebkitAnimationDuration prop
func (*CSSStyleDeclaration) WebkitAnimationDuration() (webkitAnimationDuration string) {
	js.Rewrite("$<.webkitAnimationDuration")
	return webkitAnimationDuration
}

// WebkitAnimationDuration prop
func (*CSSStyleDeclaration) SetWebkitAnimationDuration(webkitAnimationDuration string) {
	js.Rewrite("$<.webkitAnimationDuration = $1", webkitAnimationDuration)
}

// WebkitAnimationFillMode prop
func (*CSSStyleDeclaration) WebkitAnimationFillMode() (webkitAnimationFillMode string) {
	js.Rewrite("$<.webkitAnimationFillMode")
	return webkitAnimationFillMode
}

// WebkitAnimationFillMode prop
func (*CSSStyleDeclaration) SetWebkitAnimationFillMode(webkitAnimationFillMode string) {
	js.Rewrite("$<.webkitAnimationFillMode = $1", webkitAnimationFillMode)
}

// WebkitAnimationIterationCount prop
func (*CSSStyleDeclaration) WebkitAnimationIterationCount() (webkitAnimationIterationCount string) {
	js.Rewrite("$<.webkitAnimationIterationCount")
	return webkitAnimationIterationCount
}

// WebkitAnimationIterationCount prop
func (*CSSStyleDeclaration) SetWebkitAnimationIterationCount(webkitAnimationIterationCount string) {
	js.Rewrite("$<.webkitAnimationIterationCount = $1", webkitAnimationIterationCount)
}

// WebkitAnimationName prop
func (*CSSStyleDeclaration) WebkitAnimationName() (webkitAnimationName string) {
	js.Rewrite("$<.webkitAnimationName")
	return webkitAnimationName
}

// WebkitAnimationName prop
func (*CSSStyleDeclaration) SetWebkitAnimationName(webkitAnimationName string) {
	js.Rewrite("$<.webkitAnimationName = $1", webkitAnimationName)
}

// WebkitAnimationPlayState prop
func (*CSSStyleDeclaration) WebkitAnimationPlayState() (webkitAnimationPlayState string) {
	js.Rewrite("$<.webkitAnimationPlayState")
	return webkitAnimationPlayState
}

// WebkitAnimationPlayState prop
func (*CSSStyleDeclaration) SetWebkitAnimationPlayState(webkitAnimationPlayState string) {
	js.Rewrite("$<.webkitAnimationPlayState = $1", webkitAnimationPlayState)
}

// WebkitAnimationTimingFunction prop
func (*CSSStyleDeclaration) WebkitAnimationTimingFunction() (webkitAnimationTimingFunction string) {
	js.Rewrite("$<.webkitAnimationTimingFunction")
	return webkitAnimationTimingFunction
}

// WebkitAnimationTimingFunction prop
func (*CSSStyleDeclaration) SetWebkitAnimationTimingFunction(webkitAnimationTimingFunction string) {
	js.Rewrite("$<.webkitAnimationTimingFunction = $1", webkitAnimationTimingFunction)
}

// WebkitAppearance prop
func (*CSSStyleDeclaration) WebkitAppearance() (webkitAppearance string) {
	js.Rewrite("$<.webkitAppearance")
	return webkitAppearance
}

// WebkitAppearance prop
func (*CSSStyleDeclaration) SetWebkitAppearance(webkitAppearance string) {
	js.Rewrite("$<.webkitAppearance = $1", webkitAppearance)
}

// WebkitBackfaceVisibility prop
func (*CSSStyleDeclaration) WebkitBackfaceVisibility() (webkitBackfaceVisibility string) {
	js.Rewrite("$<.webkitBackfaceVisibility")
	return webkitBackfaceVisibility
}

// WebkitBackfaceVisibility prop
func (*CSSStyleDeclaration) SetWebkitBackfaceVisibility(webkitBackfaceVisibility string) {
	js.Rewrite("$<.webkitBackfaceVisibility = $1", webkitBackfaceVisibility)
}

// WebkitBackgroundClip prop
func (*CSSStyleDeclaration) WebkitBackgroundClip() (webkitBackgroundClip string) {
	js.Rewrite("$<.webkitBackgroundClip")
	return webkitBackgroundClip
}

// WebkitBackgroundClip prop
func (*CSSStyleDeclaration) SetWebkitBackgroundClip(webkitBackgroundClip string) {
	js.Rewrite("$<.webkitBackgroundClip = $1", webkitBackgroundClip)
}

// WebkitBackgroundOrigin prop
func (*CSSStyleDeclaration) WebkitBackgroundOrigin() (webkitBackgroundOrigin string) {
	js.Rewrite("$<.webkitBackgroundOrigin")
	return webkitBackgroundOrigin
}

// WebkitBackgroundOrigin prop
func (*CSSStyleDeclaration) SetWebkitBackgroundOrigin(webkitBackgroundOrigin string) {
	js.Rewrite("$<.webkitBackgroundOrigin = $1", webkitBackgroundOrigin)
}

// WebkitBackgroundSize prop
func (*CSSStyleDeclaration) WebkitBackgroundSize() (webkitBackgroundSize string) {
	js.Rewrite("$<.webkitBackgroundSize")
	return webkitBackgroundSize
}

// WebkitBackgroundSize prop
func (*CSSStyleDeclaration) SetWebkitBackgroundSize(webkitBackgroundSize string) {
	js.Rewrite("$<.webkitBackgroundSize = $1", webkitBackgroundSize)
}

// WebkitBorderBottomLeftRadius prop
func (*CSSStyleDeclaration) WebkitBorderBottomLeftRadius() (webkitBorderBottomLeftRadius string) {
	js.Rewrite("$<.webkitBorderBottomLeftRadius")
	return webkitBorderBottomLeftRadius
}

// WebkitBorderBottomLeftRadius prop
func (*CSSStyleDeclaration) SetWebkitBorderBottomLeftRadius(webkitBorderBottomLeftRadius string) {
	js.Rewrite("$<.webkitBorderBottomLeftRadius = $1", webkitBorderBottomLeftRadius)
}

// WebkitBorderBottomRightRadius prop
func (*CSSStyleDeclaration) WebkitBorderBottomRightRadius() (webkitBorderBottomRightRadius string) {
	js.Rewrite("$<.webkitBorderBottomRightRadius")
	return webkitBorderBottomRightRadius
}

// WebkitBorderBottomRightRadius prop
func (*CSSStyleDeclaration) SetWebkitBorderBottomRightRadius(webkitBorderBottomRightRadius string) {
	js.Rewrite("$<.webkitBorderBottomRightRadius = $1", webkitBorderBottomRightRadius)
}

// WebkitBorderImage prop
func (*CSSStyleDeclaration) WebkitBorderImage() (webkitBorderImage string) {
	js.Rewrite("$<.webkitBorderImage")
	return webkitBorderImage
}

// WebkitBorderImage prop
func (*CSSStyleDeclaration) SetWebkitBorderImage(webkitBorderImage string) {
	js.Rewrite("$<.webkitBorderImage = $1", webkitBorderImage)
}

// WebkitBorderRadius prop
func (*CSSStyleDeclaration) WebkitBorderRadius() (webkitBorderRadius string) {
	js.Rewrite("$<.webkitBorderRadius")
	return webkitBorderRadius
}

// WebkitBorderRadius prop
func (*CSSStyleDeclaration) SetWebkitBorderRadius(webkitBorderRadius string) {
	js.Rewrite("$<.webkitBorderRadius = $1", webkitBorderRadius)
}

// WebkitBorderTopLeftRadius prop
func (*CSSStyleDeclaration) WebkitBorderTopLeftRadius() (webkitBorderTopLeftRadius string) {
	js.Rewrite("$<.webkitBorderTopLeftRadius")
	return webkitBorderTopLeftRadius
}

// WebkitBorderTopLeftRadius prop
func (*CSSStyleDeclaration) SetWebkitBorderTopLeftRadius(webkitBorderTopLeftRadius string) {
	js.Rewrite("$<.webkitBorderTopLeftRadius = $1", webkitBorderTopLeftRadius)
}

// WebkitBorderTopRightRadius prop
func (*CSSStyleDeclaration) WebkitBorderTopRightRadius() (webkitBorderTopRightRadius string) {
	js.Rewrite("$<.webkitBorderTopRightRadius")
	return webkitBorderTopRightRadius
}

// WebkitBorderTopRightRadius prop
func (*CSSStyleDeclaration) SetWebkitBorderTopRightRadius(webkitBorderTopRightRadius string) {
	js.Rewrite("$<.webkitBorderTopRightRadius = $1", webkitBorderTopRightRadius)
}

// WebkitBoxAlign prop
func (*CSSStyleDeclaration) WebkitBoxAlign() (webkitBoxAlign string) {
	js.Rewrite("$<.webkitBoxAlign")
	return webkitBoxAlign
}

// WebkitBoxAlign prop
func (*CSSStyleDeclaration) SetWebkitBoxAlign(webkitBoxAlign string) {
	js.Rewrite("$<.webkitBoxAlign = $1", webkitBoxAlign)
}

// WebkitBoxDirection prop
func (*CSSStyleDeclaration) WebkitBoxDirection() (webkitBoxDirection string) {
	js.Rewrite("$<.webkitBoxDirection")
	return webkitBoxDirection
}

// WebkitBoxDirection prop
func (*CSSStyleDeclaration) SetWebkitBoxDirection(webkitBoxDirection string) {
	js.Rewrite("$<.webkitBoxDirection = $1", webkitBoxDirection)
}

// WebkitBoxFlex prop
func (*CSSStyleDeclaration) WebkitBoxFlex() (webkitBoxFlex string) {
	js.Rewrite("$<.webkitBoxFlex")
	return webkitBoxFlex
}

// WebkitBoxFlex prop
func (*CSSStyleDeclaration) SetWebkitBoxFlex(webkitBoxFlex string) {
	js.Rewrite("$<.webkitBoxFlex = $1", webkitBoxFlex)
}

// WebkitBoxOrdinalGroup prop
func (*CSSStyleDeclaration) WebkitBoxOrdinalGroup() (webkitBoxOrdinalGroup string) {
	js.Rewrite("$<.webkitBoxOrdinalGroup")
	return webkitBoxOrdinalGroup
}

// WebkitBoxOrdinalGroup prop
func (*CSSStyleDeclaration) SetWebkitBoxOrdinalGroup(webkitBoxOrdinalGroup string) {
	js.Rewrite("$<.webkitBoxOrdinalGroup = $1", webkitBoxOrdinalGroup)
}

// WebkitBoxOrient prop
func (*CSSStyleDeclaration) WebkitBoxOrient() (webkitBoxOrient string) {
	js.Rewrite("$<.webkitBoxOrient")
	return webkitBoxOrient
}

// WebkitBoxOrient prop
func (*CSSStyleDeclaration) SetWebkitBoxOrient(webkitBoxOrient string) {
	js.Rewrite("$<.webkitBoxOrient = $1", webkitBoxOrient)
}

// WebkitBoxPack prop
func (*CSSStyleDeclaration) WebkitBoxPack() (webkitBoxPack string) {
	js.Rewrite("$<.webkitBoxPack")
	return webkitBoxPack
}

// WebkitBoxPack prop
func (*CSSStyleDeclaration) SetWebkitBoxPack(webkitBoxPack string) {
	js.Rewrite("$<.webkitBoxPack = $1", webkitBoxPack)
}

// WebkitBoxSizing prop
func (*CSSStyleDeclaration) WebkitBoxSizing() (webkitBoxSizing string) {
	js.Rewrite("$<.webkitBoxSizing")
	return webkitBoxSizing
}

// WebkitBoxSizing prop
func (*CSSStyleDeclaration) SetWebkitBoxSizing(webkitBoxSizing string) {
	js.Rewrite("$<.webkitBoxSizing = $1", webkitBoxSizing)
}

// WebkitColumnBreakAfter prop
func (*CSSStyleDeclaration) WebkitColumnBreakAfter() (webkitColumnBreakAfter string) {
	js.Rewrite("$<.webkitColumnBreakAfter")
	return webkitColumnBreakAfter
}

// WebkitColumnBreakAfter prop
func (*CSSStyleDeclaration) SetWebkitColumnBreakAfter(webkitColumnBreakAfter string) {
	js.Rewrite("$<.webkitColumnBreakAfter = $1", webkitColumnBreakAfter)
}

// WebkitColumnBreakBefore prop
func (*CSSStyleDeclaration) WebkitColumnBreakBefore() (webkitColumnBreakBefore string) {
	js.Rewrite("$<.webkitColumnBreakBefore")
	return webkitColumnBreakBefore
}

// WebkitColumnBreakBefore prop
func (*CSSStyleDeclaration) SetWebkitColumnBreakBefore(webkitColumnBreakBefore string) {
	js.Rewrite("$<.webkitColumnBreakBefore = $1", webkitColumnBreakBefore)
}

// WebkitColumnBreakInside prop
func (*CSSStyleDeclaration) WebkitColumnBreakInside() (webkitColumnBreakInside string) {
	js.Rewrite("$<.webkitColumnBreakInside")
	return webkitColumnBreakInside
}

// WebkitColumnBreakInside prop
func (*CSSStyleDeclaration) SetWebkitColumnBreakInside(webkitColumnBreakInside string) {
	js.Rewrite("$<.webkitColumnBreakInside = $1", webkitColumnBreakInside)
}

// WebkitColumnCount prop
func (*CSSStyleDeclaration) WebkitColumnCount() (webkitColumnCount interface{}) {
	js.Rewrite("$<.webkitColumnCount")
	return webkitColumnCount
}

// WebkitColumnCount prop
func (*CSSStyleDeclaration) SetWebkitColumnCount(webkitColumnCount interface{}) {
	js.Rewrite("$<.webkitColumnCount = $1", webkitColumnCount)
}

// WebkitColumnGap prop
func (*CSSStyleDeclaration) WebkitColumnGap() (webkitColumnGap interface{}) {
	js.Rewrite("$<.webkitColumnGap")
	return webkitColumnGap
}

// WebkitColumnGap prop
func (*CSSStyleDeclaration) SetWebkitColumnGap(webkitColumnGap interface{}) {
	js.Rewrite("$<.webkitColumnGap = $1", webkitColumnGap)
}

// WebkitColumnRule prop
func (*CSSStyleDeclaration) WebkitColumnRule() (webkitColumnRule string) {
	js.Rewrite("$<.webkitColumnRule")
	return webkitColumnRule
}

// WebkitColumnRule prop
func (*CSSStyleDeclaration) SetWebkitColumnRule(webkitColumnRule string) {
	js.Rewrite("$<.webkitColumnRule = $1", webkitColumnRule)
}

// WebkitColumnRuleColor prop
func (*CSSStyleDeclaration) WebkitColumnRuleColor() (webkitColumnRuleColor interface{}) {
	js.Rewrite("$<.webkitColumnRuleColor")
	return webkitColumnRuleColor
}

// WebkitColumnRuleColor prop
func (*CSSStyleDeclaration) SetWebkitColumnRuleColor(webkitColumnRuleColor interface{}) {
	js.Rewrite("$<.webkitColumnRuleColor = $1", webkitColumnRuleColor)
}

// WebkitColumnRuleStyle prop
func (*CSSStyleDeclaration) WebkitColumnRuleStyle() (webkitColumnRuleStyle string) {
	js.Rewrite("$<.webkitColumnRuleStyle")
	return webkitColumnRuleStyle
}

// WebkitColumnRuleStyle prop
func (*CSSStyleDeclaration) SetWebkitColumnRuleStyle(webkitColumnRuleStyle string) {
	js.Rewrite("$<.webkitColumnRuleStyle = $1", webkitColumnRuleStyle)
}

// WebkitColumnRuleWidth prop
func (*CSSStyleDeclaration) WebkitColumnRuleWidth() (webkitColumnRuleWidth interface{}) {
	js.Rewrite("$<.webkitColumnRuleWidth")
	return webkitColumnRuleWidth
}

// WebkitColumnRuleWidth prop
func (*CSSStyleDeclaration) SetWebkitColumnRuleWidth(webkitColumnRuleWidth interface{}) {
	js.Rewrite("$<.webkitColumnRuleWidth = $1", webkitColumnRuleWidth)
}

// WebkitColumns prop
func (*CSSStyleDeclaration) WebkitColumns() (webkitColumns string) {
	js.Rewrite("$<.webkitColumns")
	return webkitColumns
}

// WebkitColumns prop
func (*CSSStyleDeclaration) SetWebkitColumns(webkitColumns string) {
	js.Rewrite("$<.webkitColumns = $1", webkitColumns)
}

// WebkitColumnSpan prop
func (*CSSStyleDeclaration) WebkitColumnSpan() (webkitColumnSpan string) {
	js.Rewrite("$<.webkitColumnSpan")
	return webkitColumnSpan
}

// WebkitColumnSpan prop
func (*CSSStyleDeclaration) SetWebkitColumnSpan(webkitColumnSpan string) {
	js.Rewrite("$<.webkitColumnSpan = $1", webkitColumnSpan)
}

// WebkitColumnWidth prop
func (*CSSStyleDeclaration) WebkitColumnWidth() (webkitColumnWidth interface{}) {
	js.Rewrite("$<.webkitColumnWidth")
	return webkitColumnWidth
}

// WebkitColumnWidth prop
func (*CSSStyleDeclaration) SetWebkitColumnWidth(webkitColumnWidth interface{}) {
	js.Rewrite("$<.webkitColumnWidth = $1", webkitColumnWidth)
}

// WebkitFilter prop
func (*CSSStyleDeclaration) WebkitFilter() (webkitFilter string) {
	js.Rewrite("$<.webkitFilter")
	return webkitFilter
}

// WebkitFilter prop
func (*CSSStyleDeclaration) SetWebkitFilter(webkitFilter string) {
	js.Rewrite("$<.webkitFilter = $1", webkitFilter)
}

// WebkitFlex prop
func (*CSSStyleDeclaration) WebkitFlex() (webkitFlex string) {
	js.Rewrite("$<.webkitFlex")
	return webkitFlex
}

// WebkitFlex prop
func (*CSSStyleDeclaration) SetWebkitFlex(webkitFlex string) {
	js.Rewrite("$<.webkitFlex = $1", webkitFlex)
}

// WebkitFlexBasis prop
func (*CSSStyleDeclaration) WebkitFlexBasis() (webkitFlexBasis string) {
	js.Rewrite("$<.webkitFlexBasis")
	return webkitFlexBasis
}

// WebkitFlexBasis prop
func (*CSSStyleDeclaration) SetWebkitFlexBasis(webkitFlexBasis string) {
	js.Rewrite("$<.webkitFlexBasis = $1", webkitFlexBasis)
}

// WebkitFlexDirection prop
func (*CSSStyleDeclaration) WebkitFlexDirection() (webkitFlexDirection string) {
	js.Rewrite("$<.webkitFlexDirection")
	return webkitFlexDirection
}

// WebkitFlexDirection prop
func (*CSSStyleDeclaration) SetWebkitFlexDirection(webkitFlexDirection string) {
	js.Rewrite("$<.webkitFlexDirection = $1", webkitFlexDirection)
}

// WebkitFlexFlow prop
func (*CSSStyleDeclaration) WebkitFlexFlow() (webkitFlexFlow string) {
	js.Rewrite("$<.webkitFlexFlow")
	return webkitFlexFlow
}

// WebkitFlexFlow prop
func (*CSSStyleDeclaration) SetWebkitFlexFlow(webkitFlexFlow string) {
	js.Rewrite("$<.webkitFlexFlow = $1", webkitFlexFlow)
}

// WebkitFlexGrow prop
func (*CSSStyleDeclaration) WebkitFlexGrow() (webkitFlexGrow string) {
	js.Rewrite("$<.webkitFlexGrow")
	return webkitFlexGrow
}

// WebkitFlexGrow prop
func (*CSSStyleDeclaration) SetWebkitFlexGrow(webkitFlexGrow string) {
	js.Rewrite("$<.webkitFlexGrow = $1", webkitFlexGrow)
}

// WebkitFlexShrink prop
func (*CSSStyleDeclaration) WebkitFlexShrink() (webkitFlexShrink string) {
	js.Rewrite("$<.webkitFlexShrink")
	return webkitFlexShrink
}

// WebkitFlexShrink prop
func (*CSSStyleDeclaration) SetWebkitFlexShrink(webkitFlexShrink string) {
	js.Rewrite("$<.webkitFlexShrink = $1", webkitFlexShrink)
}

// WebkitFlexWrap prop
func (*CSSStyleDeclaration) WebkitFlexWrap() (webkitFlexWrap string) {
	js.Rewrite("$<.webkitFlexWrap")
	return webkitFlexWrap
}

// WebkitFlexWrap prop
func (*CSSStyleDeclaration) SetWebkitFlexWrap(webkitFlexWrap string) {
	js.Rewrite("$<.webkitFlexWrap = $1", webkitFlexWrap)
}

// WebkitJustifyContent prop
func (*CSSStyleDeclaration) WebkitJustifyContent() (webkitJustifyContent string) {
	js.Rewrite("$<.webkitJustifyContent")
	return webkitJustifyContent
}

// WebkitJustifyContent prop
func (*CSSStyleDeclaration) SetWebkitJustifyContent(webkitJustifyContent string) {
	js.Rewrite("$<.webkitJustifyContent = $1", webkitJustifyContent)
}

// WebkitOrder prop
func (*CSSStyleDeclaration) WebkitOrder() (webkitOrder string) {
	js.Rewrite("$<.webkitOrder")
	return webkitOrder
}

// WebkitOrder prop
func (*CSSStyleDeclaration) SetWebkitOrder(webkitOrder string) {
	js.Rewrite("$<.webkitOrder = $1", webkitOrder)
}

// WebkitPerspective prop
func (*CSSStyleDeclaration) WebkitPerspective() (webkitPerspective string) {
	js.Rewrite("$<.webkitPerspective")
	return webkitPerspective
}

// WebkitPerspective prop
func (*CSSStyleDeclaration) SetWebkitPerspective(webkitPerspective string) {
	js.Rewrite("$<.webkitPerspective = $1", webkitPerspective)
}

// WebkitPerspectiveOrigin prop
func (*CSSStyleDeclaration) WebkitPerspectiveOrigin() (webkitPerspectiveOrigin string) {
	js.Rewrite("$<.webkitPerspectiveOrigin")
	return webkitPerspectiveOrigin
}

// WebkitPerspectiveOrigin prop
func (*CSSStyleDeclaration) SetWebkitPerspectiveOrigin(webkitPerspectiveOrigin string) {
	js.Rewrite("$<.webkitPerspectiveOrigin = $1", webkitPerspectiveOrigin)
}

// WebkitTapHighlightColor prop
func (*CSSStyleDeclaration) WebkitTapHighlightColor() (webkitTapHighlightColor string) {
	js.Rewrite("$<.webkitTapHighlightColor")
	return webkitTapHighlightColor
}

// WebkitTapHighlightColor prop
func (*CSSStyleDeclaration) SetWebkitTapHighlightColor(webkitTapHighlightColor string) {
	js.Rewrite("$<.webkitTapHighlightColor = $1", webkitTapHighlightColor)
}

// WebkitTextFillColor prop
func (*CSSStyleDeclaration) WebkitTextFillColor() (webkitTextFillColor string) {
	js.Rewrite("$<.webkitTextFillColor")
	return webkitTextFillColor
}

// WebkitTextFillColor prop
func (*CSSStyleDeclaration) SetWebkitTextFillColor(webkitTextFillColor string) {
	js.Rewrite("$<.webkitTextFillColor = $1", webkitTextFillColor)
}

// WebkitTextSizeAdjust prop
func (*CSSStyleDeclaration) WebkitTextSizeAdjust() (webkitTextSizeAdjust interface{}) {
	js.Rewrite("$<.webkitTextSizeAdjust")
	return webkitTextSizeAdjust
}

// WebkitTextSizeAdjust prop
func (*CSSStyleDeclaration) SetWebkitTextSizeAdjust(webkitTextSizeAdjust interface{}) {
	js.Rewrite("$<.webkitTextSizeAdjust = $1", webkitTextSizeAdjust)
}

// WebkitTextStroke prop
func (*CSSStyleDeclaration) WebkitTextStroke() (webkitTextStroke string) {
	js.Rewrite("$<.webkitTextStroke")
	return webkitTextStroke
}

// WebkitTextStroke prop
func (*CSSStyleDeclaration) SetWebkitTextStroke(webkitTextStroke string) {
	js.Rewrite("$<.webkitTextStroke = $1", webkitTextStroke)
}

// WebkitTextStrokeColor prop
func (*CSSStyleDeclaration) WebkitTextStrokeColor() (webkitTextStrokeColor string) {
	js.Rewrite("$<.webkitTextStrokeColor")
	return webkitTextStrokeColor
}

// WebkitTextStrokeColor prop
func (*CSSStyleDeclaration) SetWebkitTextStrokeColor(webkitTextStrokeColor string) {
	js.Rewrite("$<.webkitTextStrokeColor = $1", webkitTextStrokeColor)
}

// WebkitTextStrokeWidth prop
func (*CSSStyleDeclaration) WebkitTextStrokeWidth() (webkitTextStrokeWidth string) {
	js.Rewrite("$<.webkitTextStrokeWidth")
	return webkitTextStrokeWidth
}

// WebkitTextStrokeWidth prop
func (*CSSStyleDeclaration) SetWebkitTextStrokeWidth(webkitTextStrokeWidth string) {
	js.Rewrite("$<.webkitTextStrokeWidth = $1", webkitTextStrokeWidth)
}

// WebkitTransform prop
func (*CSSStyleDeclaration) WebkitTransform() (webkitTransform string) {
	js.Rewrite("$<.webkitTransform")
	return webkitTransform
}

// WebkitTransform prop
func (*CSSStyleDeclaration) SetWebkitTransform(webkitTransform string) {
	js.Rewrite("$<.webkitTransform = $1", webkitTransform)
}

// WebkitTransformOrigin prop
func (*CSSStyleDeclaration) WebkitTransformOrigin() (webkitTransformOrigin string) {
	js.Rewrite("$<.webkitTransformOrigin")
	return webkitTransformOrigin
}

// WebkitTransformOrigin prop
func (*CSSStyleDeclaration) SetWebkitTransformOrigin(webkitTransformOrigin string) {
	js.Rewrite("$<.webkitTransformOrigin = $1", webkitTransformOrigin)
}

// WebkitTransformStyle prop
func (*CSSStyleDeclaration) WebkitTransformStyle() (webkitTransformStyle string) {
	js.Rewrite("$<.webkitTransformStyle")
	return webkitTransformStyle
}

// WebkitTransformStyle prop
func (*CSSStyleDeclaration) SetWebkitTransformStyle(webkitTransformStyle string) {
	js.Rewrite("$<.webkitTransformStyle = $1", webkitTransformStyle)
}

// WebkitTransition prop
func (*CSSStyleDeclaration) WebkitTransition() (webkitTransition string) {
	js.Rewrite("$<.webkitTransition")
	return webkitTransition
}

// WebkitTransition prop
func (*CSSStyleDeclaration) SetWebkitTransition(webkitTransition string) {
	js.Rewrite("$<.webkitTransition = $1", webkitTransition)
}

// WebkitTransitionDelay prop
func (*CSSStyleDeclaration) WebkitTransitionDelay() (webkitTransitionDelay string) {
	js.Rewrite("$<.webkitTransitionDelay")
	return webkitTransitionDelay
}

// WebkitTransitionDelay prop
func (*CSSStyleDeclaration) SetWebkitTransitionDelay(webkitTransitionDelay string) {
	js.Rewrite("$<.webkitTransitionDelay = $1", webkitTransitionDelay)
}

// WebkitTransitionDuration prop
func (*CSSStyleDeclaration) WebkitTransitionDuration() (webkitTransitionDuration string) {
	js.Rewrite("$<.webkitTransitionDuration")
	return webkitTransitionDuration
}

// WebkitTransitionDuration prop
func (*CSSStyleDeclaration) SetWebkitTransitionDuration(webkitTransitionDuration string) {
	js.Rewrite("$<.webkitTransitionDuration = $1", webkitTransitionDuration)
}

// WebkitTransitionProperty prop
func (*CSSStyleDeclaration) WebkitTransitionProperty() (webkitTransitionProperty string) {
	js.Rewrite("$<.webkitTransitionProperty")
	return webkitTransitionProperty
}

// WebkitTransitionProperty prop
func (*CSSStyleDeclaration) SetWebkitTransitionProperty(webkitTransitionProperty string) {
	js.Rewrite("$<.webkitTransitionProperty = $1", webkitTransitionProperty)
}

// WebkitTransitionTimingFunction prop
func (*CSSStyleDeclaration) WebkitTransitionTimingFunction() (webkitTransitionTimingFunction string) {
	js.Rewrite("$<.webkitTransitionTimingFunction")
	return webkitTransitionTimingFunction
}

// WebkitTransitionTimingFunction prop
func (*CSSStyleDeclaration) SetWebkitTransitionTimingFunction(webkitTransitionTimingFunction string) {
	js.Rewrite("$<.webkitTransitionTimingFunction = $1", webkitTransitionTimingFunction)
}

// WebkitUserModify prop
func (*CSSStyleDeclaration) WebkitUserModify() (webkitUserModify string) {
	js.Rewrite("$<.webkitUserModify")
	return webkitUserModify
}

// WebkitUserModify prop
func (*CSSStyleDeclaration) SetWebkitUserModify(webkitUserModify string) {
	js.Rewrite("$<.webkitUserModify = $1", webkitUserModify)
}

// WebkitUserSelect prop
func (*CSSStyleDeclaration) WebkitUserSelect() (webkitUserSelect string) {
	js.Rewrite("$<.webkitUserSelect")
	return webkitUserSelect
}

// WebkitUserSelect prop
func (*CSSStyleDeclaration) SetWebkitUserSelect(webkitUserSelect string) {
	js.Rewrite("$<.webkitUserSelect = $1", webkitUserSelect)
}

// WebkitWritingMode prop
func (*CSSStyleDeclaration) WebkitWritingMode() (webkitWritingMode string) {
	js.Rewrite("$<.webkitWritingMode")
	return webkitWritingMode
}

// WebkitWritingMode prop
func (*CSSStyleDeclaration) SetWebkitWritingMode(webkitWritingMode string) {
	js.Rewrite("$<.webkitWritingMode = $1", webkitWritingMode)
}

// WhiteSpace prop
func (*CSSStyleDeclaration) WhiteSpace() (whiteSpace string) {
	js.Rewrite("$<.whiteSpace")
	return whiteSpace
}

// WhiteSpace prop
func (*CSSStyleDeclaration) SetWhiteSpace(whiteSpace string) {
	js.Rewrite("$<.whiteSpace = $1", whiteSpace)
}

// Widows prop
func (*CSSStyleDeclaration) Widows() (widows string) {
	js.Rewrite("$<.widows")
	return widows
}

// Widows prop
func (*CSSStyleDeclaration) SetWidows(widows string) {
	js.Rewrite("$<.widows = $1", widows)
}

// Width prop
func (*CSSStyleDeclaration) Width() (width string) {
	js.Rewrite("$<.width")
	return width
}

// Width prop
func (*CSSStyleDeclaration) SetWidth(width string) {
	js.Rewrite("$<.width = $1", width)
}

// WordBreak prop
func (*CSSStyleDeclaration) WordBreak() (wordBreak string) {
	js.Rewrite("$<.wordBreak")
	return wordBreak
}

// WordBreak prop
func (*CSSStyleDeclaration) SetWordBreak(wordBreak string) {
	js.Rewrite("$<.wordBreak = $1", wordBreak)
}

// WordSpacing prop
func (*CSSStyleDeclaration) WordSpacing() (wordSpacing string) {
	js.Rewrite("$<.wordSpacing")
	return wordSpacing
}

// WordSpacing prop
func (*CSSStyleDeclaration) SetWordSpacing(wordSpacing string) {
	js.Rewrite("$<.wordSpacing = $1", wordSpacing)
}

// WordWrap prop
func (*CSSStyleDeclaration) WordWrap() (wordWrap string) {
	js.Rewrite("$<.wordWrap")
	return wordWrap
}

// WordWrap prop
func (*CSSStyleDeclaration) SetWordWrap(wordWrap string) {
	js.Rewrite("$<.wordWrap = $1", wordWrap)
}

// WritingMode prop
func (*CSSStyleDeclaration) WritingMode() (writingMode string) {
	js.Rewrite("$<.writingMode")
	return writingMode
}

// WritingMode prop
func (*CSSStyleDeclaration) SetWritingMode(writingMode string) {
	js.Rewrite("$<.writingMode = $1", writingMode)
}

// ZIndex prop
func (*CSSStyleDeclaration) ZIndex() (zIndex string) {
	js.Rewrite("$<.zIndex")
	return zIndex
}

// ZIndex prop
func (*CSSStyleDeclaration) SetZIndex(zIndex string) {
	js.Rewrite("$<.zIndex = $1", zIndex)
}

// Zoom prop
func (*CSSStyleDeclaration) Zoom() (zoom string) {
	js.Rewrite("$<.zoom")
	return zoom
}

// Zoom prop
func (*CSSStyleDeclaration) SetZoom(zoom string) {
	js.Rewrite("$<.zoom = $1", zoom)
}
