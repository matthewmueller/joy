package window

import "github.com/matthewmueller/golly/js"

// CSSStyleDeclaration struct
// js:"CSSStyleDeclaration,omit"
type CSSStyleDeclaration struct {
}

// GetPropertyPriority fn
// js:"getPropertyPriority"
func (*CSSStyleDeclaration) GetPropertyPriority(propertyName string) (s string) {
	js.Rewrite("$_.getPropertyPriority($1)", propertyName)
	return s
}

// GetPropertyValue fn
// js:"getPropertyValue"
func (*CSSStyleDeclaration) GetPropertyValue(propertyName string) (s string) {
	js.Rewrite("$_.getPropertyValue($1)", propertyName)
	return s
}

// Item fn
// js:"item"
func (*CSSStyleDeclaration) Item(index uint) (s string) {
	js.Rewrite("$_.item($1)", index)
	return s
}

// RemoveProperty fn
// js:"removeProperty"
func (*CSSStyleDeclaration) RemoveProperty(propertyName string) (s string) {
	js.Rewrite("$_.removeProperty($1)", propertyName)
	return s
}

// SetProperty fn
// js:"setProperty"
func (*CSSStyleDeclaration) SetProperty(propertyName string, value string, priority *string) {
	js.Rewrite("$_.setProperty($1, $2, $3)", propertyName, value, priority)
}

// AlignContent prop
// js:"alignContent"
func (*CSSStyleDeclaration) AlignContent() (alignContent string) {
	js.Rewrite("$_.alignContent")
	return alignContent
}

// SetAlignContent prop
// js:"alignContent"
func (*CSSStyleDeclaration) SetAlignContent(alignContent string) {
	js.Rewrite("$_.alignContent = $1", alignContent)
}

// AlignItems prop
// js:"alignItems"
func (*CSSStyleDeclaration) AlignItems() (alignItems string) {
	js.Rewrite("$_.alignItems")
	return alignItems
}

// SetAlignItems prop
// js:"alignItems"
func (*CSSStyleDeclaration) SetAlignItems(alignItems string) {
	js.Rewrite("$_.alignItems = $1", alignItems)
}

// AlignmentBaseline prop
// js:"alignmentBaseline"
func (*CSSStyleDeclaration) AlignmentBaseline() (alignmentBaseline string) {
	js.Rewrite("$_.alignmentBaseline")
	return alignmentBaseline
}

// SetAlignmentBaseline prop
// js:"alignmentBaseline"
func (*CSSStyleDeclaration) SetAlignmentBaseline(alignmentBaseline string) {
	js.Rewrite("$_.alignmentBaseline = $1", alignmentBaseline)
}

// AlignSelf prop
// js:"alignSelf"
func (*CSSStyleDeclaration) AlignSelf() (alignSelf string) {
	js.Rewrite("$_.alignSelf")
	return alignSelf
}

// SetAlignSelf prop
// js:"alignSelf"
func (*CSSStyleDeclaration) SetAlignSelf(alignSelf string) {
	js.Rewrite("$_.alignSelf = $1", alignSelf)
}

// Animation prop
// js:"animation"
func (*CSSStyleDeclaration) Animation() (animation string) {
	js.Rewrite("$_.animation")
	return animation
}

// SetAnimation prop
// js:"animation"
func (*CSSStyleDeclaration) SetAnimation(animation string) {
	js.Rewrite("$_.animation = $1", animation)
}

// AnimationDelay prop
// js:"animationDelay"
func (*CSSStyleDeclaration) AnimationDelay() (animationDelay string) {
	js.Rewrite("$_.animationDelay")
	return animationDelay
}

// SetAnimationDelay prop
// js:"animationDelay"
func (*CSSStyleDeclaration) SetAnimationDelay(animationDelay string) {
	js.Rewrite("$_.animationDelay = $1", animationDelay)
}

// AnimationDirection prop
// js:"animationDirection"
func (*CSSStyleDeclaration) AnimationDirection() (animationDirection string) {
	js.Rewrite("$_.animationDirection")
	return animationDirection
}

// SetAnimationDirection prop
// js:"animationDirection"
func (*CSSStyleDeclaration) SetAnimationDirection(animationDirection string) {
	js.Rewrite("$_.animationDirection = $1", animationDirection)
}

// AnimationDuration prop
// js:"animationDuration"
func (*CSSStyleDeclaration) AnimationDuration() (animationDuration string) {
	js.Rewrite("$_.animationDuration")
	return animationDuration
}

// SetAnimationDuration prop
// js:"animationDuration"
func (*CSSStyleDeclaration) SetAnimationDuration(animationDuration string) {
	js.Rewrite("$_.animationDuration = $1", animationDuration)
}

// AnimationFillMode prop
// js:"animationFillMode"
func (*CSSStyleDeclaration) AnimationFillMode() (animationFillMode string) {
	js.Rewrite("$_.animationFillMode")
	return animationFillMode
}

// SetAnimationFillMode prop
// js:"animationFillMode"
func (*CSSStyleDeclaration) SetAnimationFillMode(animationFillMode string) {
	js.Rewrite("$_.animationFillMode = $1", animationFillMode)
}

// AnimationIterationCount prop
// js:"animationIterationCount"
func (*CSSStyleDeclaration) AnimationIterationCount() (animationIterationCount string) {
	js.Rewrite("$_.animationIterationCount")
	return animationIterationCount
}

// SetAnimationIterationCount prop
// js:"animationIterationCount"
func (*CSSStyleDeclaration) SetAnimationIterationCount(animationIterationCount string) {
	js.Rewrite("$_.animationIterationCount = $1", animationIterationCount)
}

// AnimationName prop
// js:"animationName"
func (*CSSStyleDeclaration) AnimationName() (animationName string) {
	js.Rewrite("$_.animationName")
	return animationName
}

// SetAnimationName prop
// js:"animationName"
func (*CSSStyleDeclaration) SetAnimationName(animationName string) {
	js.Rewrite("$_.animationName = $1", animationName)
}

// AnimationPlayState prop
// js:"animationPlayState"
func (*CSSStyleDeclaration) AnimationPlayState() (animationPlayState string) {
	js.Rewrite("$_.animationPlayState")
	return animationPlayState
}

// SetAnimationPlayState prop
// js:"animationPlayState"
func (*CSSStyleDeclaration) SetAnimationPlayState(animationPlayState string) {
	js.Rewrite("$_.animationPlayState = $1", animationPlayState)
}

// AnimationTimingFunction prop
// js:"animationTimingFunction"
func (*CSSStyleDeclaration) AnimationTimingFunction() (animationTimingFunction string) {
	js.Rewrite("$_.animationTimingFunction")
	return animationTimingFunction
}

// SetAnimationTimingFunction prop
// js:"animationTimingFunction"
func (*CSSStyleDeclaration) SetAnimationTimingFunction(animationTimingFunction string) {
	js.Rewrite("$_.animationTimingFunction = $1", animationTimingFunction)
}

// BackfaceVisibility prop
// js:"backfaceVisibility"
func (*CSSStyleDeclaration) BackfaceVisibility() (backfaceVisibility string) {
	js.Rewrite("$_.backfaceVisibility")
	return backfaceVisibility
}

// SetBackfaceVisibility prop
// js:"backfaceVisibility"
func (*CSSStyleDeclaration) SetBackfaceVisibility(backfaceVisibility string) {
	js.Rewrite("$_.backfaceVisibility = $1", backfaceVisibility)
}

// Background prop
// js:"background"
func (*CSSStyleDeclaration) Background() (background string) {
	js.Rewrite("$_.background")
	return background
}

// SetBackground prop
// js:"background"
func (*CSSStyleDeclaration) SetBackground(background string) {
	js.Rewrite("$_.background = $1", background)
}

// BackgroundAttachment prop
// js:"backgroundAttachment"
func (*CSSStyleDeclaration) BackgroundAttachment() (backgroundAttachment string) {
	js.Rewrite("$_.backgroundAttachment")
	return backgroundAttachment
}

// SetBackgroundAttachment prop
// js:"backgroundAttachment"
func (*CSSStyleDeclaration) SetBackgroundAttachment(backgroundAttachment string) {
	js.Rewrite("$_.backgroundAttachment = $1", backgroundAttachment)
}

// BackgroundClip prop
// js:"backgroundClip"
func (*CSSStyleDeclaration) BackgroundClip() (backgroundClip string) {
	js.Rewrite("$_.backgroundClip")
	return backgroundClip
}

// SetBackgroundClip prop
// js:"backgroundClip"
func (*CSSStyleDeclaration) SetBackgroundClip(backgroundClip string) {
	js.Rewrite("$_.backgroundClip = $1", backgroundClip)
}

// BackgroundColor prop
// js:"backgroundColor"
func (*CSSStyleDeclaration) BackgroundColor() (backgroundColor string) {
	js.Rewrite("$_.backgroundColor")
	return backgroundColor
}

// SetBackgroundColor prop
// js:"backgroundColor"
func (*CSSStyleDeclaration) SetBackgroundColor(backgroundColor string) {
	js.Rewrite("$_.backgroundColor = $1", backgroundColor)
}

// BackgroundImage prop
// js:"backgroundImage"
func (*CSSStyleDeclaration) BackgroundImage() (backgroundImage string) {
	js.Rewrite("$_.backgroundImage")
	return backgroundImage
}

// SetBackgroundImage prop
// js:"backgroundImage"
func (*CSSStyleDeclaration) SetBackgroundImage(backgroundImage string) {
	js.Rewrite("$_.backgroundImage = $1", backgroundImage)
}

// BackgroundOrigin prop
// js:"backgroundOrigin"
func (*CSSStyleDeclaration) BackgroundOrigin() (backgroundOrigin string) {
	js.Rewrite("$_.backgroundOrigin")
	return backgroundOrigin
}

// SetBackgroundOrigin prop
// js:"backgroundOrigin"
func (*CSSStyleDeclaration) SetBackgroundOrigin(backgroundOrigin string) {
	js.Rewrite("$_.backgroundOrigin = $1", backgroundOrigin)
}

// BackgroundPosition prop
// js:"backgroundPosition"
func (*CSSStyleDeclaration) BackgroundPosition() (backgroundPosition string) {
	js.Rewrite("$_.backgroundPosition")
	return backgroundPosition
}

// SetBackgroundPosition prop
// js:"backgroundPosition"
func (*CSSStyleDeclaration) SetBackgroundPosition(backgroundPosition string) {
	js.Rewrite("$_.backgroundPosition = $1", backgroundPosition)
}

// BackgroundPositionX prop
// js:"backgroundPositionX"
func (*CSSStyleDeclaration) BackgroundPositionX() (backgroundPositionX string) {
	js.Rewrite("$_.backgroundPositionX")
	return backgroundPositionX
}

// SetBackgroundPositionX prop
// js:"backgroundPositionX"
func (*CSSStyleDeclaration) SetBackgroundPositionX(backgroundPositionX string) {
	js.Rewrite("$_.backgroundPositionX = $1", backgroundPositionX)
}

// BackgroundPositionY prop
// js:"backgroundPositionY"
func (*CSSStyleDeclaration) BackgroundPositionY() (backgroundPositionY string) {
	js.Rewrite("$_.backgroundPositionY")
	return backgroundPositionY
}

// SetBackgroundPositionY prop
// js:"backgroundPositionY"
func (*CSSStyleDeclaration) SetBackgroundPositionY(backgroundPositionY string) {
	js.Rewrite("$_.backgroundPositionY = $1", backgroundPositionY)
}

// BackgroundRepeat prop
// js:"backgroundRepeat"
func (*CSSStyleDeclaration) BackgroundRepeat() (backgroundRepeat string) {
	js.Rewrite("$_.backgroundRepeat")
	return backgroundRepeat
}

// SetBackgroundRepeat prop
// js:"backgroundRepeat"
func (*CSSStyleDeclaration) SetBackgroundRepeat(backgroundRepeat string) {
	js.Rewrite("$_.backgroundRepeat = $1", backgroundRepeat)
}

// BackgroundSize prop
// js:"backgroundSize"
func (*CSSStyleDeclaration) BackgroundSize() (backgroundSize string) {
	js.Rewrite("$_.backgroundSize")
	return backgroundSize
}

// SetBackgroundSize prop
// js:"backgroundSize"
func (*CSSStyleDeclaration) SetBackgroundSize(backgroundSize string) {
	js.Rewrite("$_.backgroundSize = $1", backgroundSize)
}

// BaselineShift prop
// js:"baselineShift"
func (*CSSStyleDeclaration) BaselineShift() (baselineShift string) {
	js.Rewrite("$_.baselineShift")
	return baselineShift
}

// SetBaselineShift prop
// js:"baselineShift"
func (*CSSStyleDeclaration) SetBaselineShift(baselineShift string) {
	js.Rewrite("$_.baselineShift = $1", baselineShift)
}

// Border prop
// js:"border"
func (*CSSStyleDeclaration) Border() (border string) {
	js.Rewrite("$_.border")
	return border
}

// SetBorder prop
// js:"border"
func (*CSSStyleDeclaration) SetBorder(border string) {
	js.Rewrite("$_.border = $1", border)
}

// BorderBottom prop
// js:"borderBottom"
func (*CSSStyleDeclaration) BorderBottom() (borderBottom string) {
	js.Rewrite("$_.borderBottom")
	return borderBottom
}

// SetBorderBottom prop
// js:"borderBottom"
func (*CSSStyleDeclaration) SetBorderBottom(borderBottom string) {
	js.Rewrite("$_.borderBottom = $1", borderBottom)
}

// BorderBottomColor prop
// js:"borderBottomColor"
func (*CSSStyleDeclaration) BorderBottomColor() (borderBottomColor string) {
	js.Rewrite("$_.borderBottomColor")
	return borderBottomColor
}

// SetBorderBottomColor prop
// js:"borderBottomColor"
func (*CSSStyleDeclaration) SetBorderBottomColor(borderBottomColor string) {
	js.Rewrite("$_.borderBottomColor = $1", borderBottomColor)
}

// BorderBottomLeftRadius prop
// js:"borderBottomLeftRadius"
func (*CSSStyleDeclaration) BorderBottomLeftRadius() (borderBottomLeftRadius string) {
	js.Rewrite("$_.borderBottomLeftRadius")
	return borderBottomLeftRadius
}

// SetBorderBottomLeftRadius prop
// js:"borderBottomLeftRadius"
func (*CSSStyleDeclaration) SetBorderBottomLeftRadius(borderBottomLeftRadius string) {
	js.Rewrite("$_.borderBottomLeftRadius = $1", borderBottomLeftRadius)
}

// BorderBottomRightRadius prop
// js:"borderBottomRightRadius"
func (*CSSStyleDeclaration) BorderBottomRightRadius() (borderBottomRightRadius string) {
	js.Rewrite("$_.borderBottomRightRadius")
	return borderBottomRightRadius
}

// SetBorderBottomRightRadius prop
// js:"borderBottomRightRadius"
func (*CSSStyleDeclaration) SetBorderBottomRightRadius(borderBottomRightRadius string) {
	js.Rewrite("$_.borderBottomRightRadius = $1", borderBottomRightRadius)
}

// BorderBottomStyle prop
// js:"borderBottomStyle"
func (*CSSStyleDeclaration) BorderBottomStyle() (borderBottomStyle string) {
	js.Rewrite("$_.borderBottomStyle")
	return borderBottomStyle
}

// SetBorderBottomStyle prop
// js:"borderBottomStyle"
func (*CSSStyleDeclaration) SetBorderBottomStyle(borderBottomStyle string) {
	js.Rewrite("$_.borderBottomStyle = $1", borderBottomStyle)
}

// BorderBottomWidth prop
// js:"borderBottomWidth"
func (*CSSStyleDeclaration) BorderBottomWidth() (borderBottomWidth string) {
	js.Rewrite("$_.borderBottomWidth")
	return borderBottomWidth
}

// SetBorderBottomWidth prop
// js:"borderBottomWidth"
func (*CSSStyleDeclaration) SetBorderBottomWidth(borderBottomWidth string) {
	js.Rewrite("$_.borderBottomWidth = $1", borderBottomWidth)
}

// BorderCollapse prop
// js:"borderCollapse"
func (*CSSStyleDeclaration) BorderCollapse() (borderCollapse string) {
	js.Rewrite("$_.borderCollapse")
	return borderCollapse
}

// SetBorderCollapse prop
// js:"borderCollapse"
func (*CSSStyleDeclaration) SetBorderCollapse(borderCollapse string) {
	js.Rewrite("$_.borderCollapse = $1", borderCollapse)
}

// BorderColor prop
// js:"borderColor"
func (*CSSStyleDeclaration) BorderColor() (borderColor string) {
	js.Rewrite("$_.borderColor")
	return borderColor
}

// SetBorderColor prop
// js:"borderColor"
func (*CSSStyleDeclaration) SetBorderColor(borderColor string) {
	js.Rewrite("$_.borderColor = $1", borderColor)
}

// BorderImage prop
// js:"borderImage"
func (*CSSStyleDeclaration) BorderImage() (borderImage string) {
	js.Rewrite("$_.borderImage")
	return borderImage
}

// SetBorderImage prop
// js:"borderImage"
func (*CSSStyleDeclaration) SetBorderImage(borderImage string) {
	js.Rewrite("$_.borderImage = $1", borderImage)
}

// BorderImageOutset prop
// js:"borderImageOutset"
func (*CSSStyleDeclaration) BorderImageOutset() (borderImageOutset string) {
	js.Rewrite("$_.borderImageOutset")
	return borderImageOutset
}

// SetBorderImageOutset prop
// js:"borderImageOutset"
func (*CSSStyleDeclaration) SetBorderImageOutset(borderImageOutset string) {
	js.Rewrite("$_.borderImageOutset = $1", borderImageOutset)
}

// BorderImageRepeat prop
// js:"borderImageRepeat"
func (*CSSStyleDeclaration) BorderImageRepeat() (borderImageRepeat string) {
	js.Rewrite("$_.borderImageRepeat")
	return borderImageRepeat
}

// SetBorderImageRepeat prop
// js:"borderImageRepeat"
func (*CSSStyleDeclaration) SetBorderImageRepeat(borderImageRepeat string) {
	js.Rewrite("$_.borderImageRepeat = $1", borderImageRepeat)
}

// BorderImageSlice prop
// js:"borderImageSlice"
func (*CSSStyleDeclaration) BorderImageSlice() (borderImageSlice string) {
	js.Rewrite("$_.borderImageSlice")
	return borderImageSlice
}

// SetBorderImageSlice prop
// js:"borderImageSlice"
func (*CSSStyleDeclaration) SetBorderImageSlice(borderImageSlice string) {
	js.Rewrite("$_.borderImageSlice = $1", borderImageSlice)
}

// BorderImageSource prop
// js:"borderImageSource"
func (*CSSStyleDeclaration) BorderImageSource() (borderImageSource string) {
	js.Rewrite("$_.borderImageSource")
	return borderImageSource
}

// SetBorderImageSource prop
// js:"borderImageSource"
func (*CSSStyleDeclaration) SetBorderImageSource(borderImageSource string) {
	js.Rewrite("$_.borderImageSource = $1", borderImageSource)
}

// BorderImageWidth prop
// js:"borderImageWidth"
func (*CSSStyleDeclaration) BorderImageWidth() (borderImageWidth string) {
	js.Rewrite("$_.borderImageWidth")
	return borderImageWidth
}

// SetBorderImageWidth prop
// js:"borderImageWidth"
func (*CSSStyleDeclaration) SetBorderImageWidth(borderImageWidth string) {
	js.Rewrite("$_.borderImageWidth = $1", borderImageWidth)
}

// BorderLeft prop
// js:"borderLeft"
func (*CSSStyleDeclaration) BorderLeft() (borderLeft string) {
	js.Rewrite("$_.borderLeft")
	return borderLeft
}

// SetBorderLeft prop
// js:"borderLeft"
func (*CSSStyleDeclaration) SetBorderLeft(borderLeft string) {
	js.Rewrite("$_.borderLeft = $1", borderLeft)
}

// BorderLeftColor prop
// js:"borderLeftColor"
func (*CSSStyleDeclaration) BorderLeftColor() (borderLeftColor string) {
	js.Rewrite("$_.borderLeftColor")
	return borderLeftColor
}

// SetBorderLeftColor prop
// js:"borderLeftColor"
func (*CSSStyleDeclaration) SetBorderLeftColor(borderLeftColor string) {
	js.Rewrite("$_.borderLeftColor = $1", borderLeftColor)
}

// BorderLeftStyle prop
// js:"borderLeftStyle"
func (*CSSStyleDeclaration) BorderLeftStyle() (borderLeftStyle string) {
	js.Rewrite("$_.borderLeftStyle")
	return borderLeftStyle
}

// SetBorderLeftStyle prop
// js:"borderLeftStyle"
func (*CSSStyleDeclaration) SetBorderLeftStyle(borderLeftStyle string) {
	js.Rewrite("$_.borderLeftStyle = $1", borderLeftStyle)
}

// BorderLeftWidth prop
// js:"borderLeftWidth"
func (*CSSStyleDeclaration) BorderLeftWidth() (borderLeftWidth string) {
	js.Rewrite("$_.borderLeftWidth")
	return borderLeftWidth
}

// SetBorderLeftWidth prop
// js:"borderLeftWidth"
func (*CSSStyleDeclaration) SetBorderLeftWidth(borderLeftWidth string) {
	js.Rewrite("$_.borderLeftWidth = $1", borderLeftWidth)
}

// BorderRadius prop
// js:"borderRadius"
func (*CSSStyleDeclaration) BorderRadius() (borderRadius string) {
	js.Rewrite("$_.borderRadius")
	return borderRadius
}

// SetBorderRadius prop
// js:"borderRadius"
func (*CSSStyleDeclaration) SetBorderRadius(borderRadius string) {
	js.Rewrite("$_.borderRadius = $1", borderRadius)
}

// BorderRight prop
// js:"borderRight"
func (*CSSStyleDeclaration) BorderRight() (borderRight string) {
	js.Rewrite("$_.borderRight")
	return borderRight
}

// SetBorderRight prop
// js:"borderRight"
func (*CSSStyleDeclaration) SetBorderRight(borderRight string) {
	js.Rewrite("$_.borderRight = $1", borderRight)
}

// BorderRightColor prop
// js:"borderRightColor"
func (*CSSStyleDeclaration) BorderRightColor() (borderRightColor string) {
	js.Rewrite("$_.borderRightColor")
	return borderRightColor
}

// SetBorderRightColor prop
// js:"borderRightColor"
func (*CSSStyleDeclaration) SetBorderRightColor(borderRightColor string) {
	js.Rewrite("$_.borderRightColor = $1", borderRightColor)
}

// BorderRightStyle prop
// js:"borderRightStyle"
func (*CSSStyleDeclaration) BorderRightStyle() (borderRightStyle string) {
	js.Rewrite("$_.borderRightStyle")
	return borderRightStyle
}

// SetBorderRightStyle prop
// js:"borderRightStyle"
func (*CSSStyleDeclaration) SetBorderRightStyle(borderRightStyle string) {
	js.Rewrite("$_.borderRightStyle = $1", borderRightStyle)
}

// BorderRightWidth prop
// js:"borderRightWidth"
func (*CSSStyleDeclaration) BorderRightWidth() (borderRightWidth string) {
	js.Rewrite("$_.borderRightWidth")
	return borderRightWidth
}

// SetBorderRightWidth prop
// js:"borderRightWidth"
func (*CSSStyleDeclaration) SetBorderRightWidth(borderRightWidth string) {
	js.Rewrite("$_.borderRightWidth = $1", borderRightWidth)
}

// BorderSpacing prop
// js:"borderSpacing"
func (*CSSStyleDeclaration) BorderSpacing() (borderSpacing string) {
	js.Rewrite("$_.borderSpacing")
	return borderSpacing
}

// SetBorderSpacing prop
// js:"borderSpacing"
func (*CSSStyleDeclaration) SetBorderSpacing(borderSpacing string) {
	js.Rewrite("$_.borderSpacing = $1", borderSpacing)
}

// BorderStyle prop
// js:"borderStyle"
func (*CSSStyleDeclaration) BorderStyle() (borderStyle string) {
	js.Rewrite("$_.borderStyle")
	return borderStyle
}

// SetBorderStyle prop
// js:"borderStyle"
func (*CSSStyleDeclaration) SetBorderStyle(borderStyle string) {
	js.Rewrite("$_.borderStyle = $1", borderStyle)
}

// BorderTop prop
// js:"borderTop"
func (*CSSStyleDeclaration) BorderTop() (borderTop string) {
	js.Rewrite("$_.borderTop")
	return borderTop
}

// SetBorderTop prop
// js:"borderTop"
func (*CSSStyleDeclaration) SetBorderTop(borderTop string) {
	js.Rewrite("$_.borderTop = $1", borderTop)
}

// BorderTopColor prop
// js:"borderTopColor"
func (*CSSStyleDeclaration) BorderTopColor() (borderTopColor string) {
	js.Rewrite("$_.borderTopColor")
	return borderTopColor
}

// SetBorderTopColor prop
// js:"borderTopColor"
func (*CSSStyleDeclaration) SetBorderTopColor(borderTopColor string) {
	js.Rewrite("$_.borderTopColor = $1", borderTopColor)
}

// BorderTopLeftRadius prop
// js:"borderTopLeftRadius"
func (*CSSStyleDeclaration) BorderTopLeftRadius() (borderTopLeftRadius string) {
	js.Rewrite("$_.borderTopLeftRadius")
	return borderTopLeftRadius
}

// SetBorderTopLeftRadius prop
// js:"borderTopLeftRadius"
func (*CSSStyleDeclaration) SetBorderTopLeftRadius(borderTopLeftRadius string) {
	js.Rewrite("$_.borderTopLeftRadius = $1", borderTopLeftRadius)
}

// BorderTopRightRadius prop
// js:"borderTopRightRadius"
func (*CSSStyleDeclaration) BorderTopRightRadius() (borderTopRightRadius string) {
	js.Rewrite("$_.borderTopRightRadius")
	return borderTopRightRadius
}

// SetBorderTopRightRadius prop
// js:"borderTopRightRadius"
func (*CSSStyleDeclaration) SetBorderTopRightRadius(borderTopRightRadius string) {
	js.Rewrite("$_.borderTopRightRadius = $1", borderTopRightRadius)
}

// BorderTopStyle prop
// js:"borderTopStyle"
func (*CSSStyleDeclaration) BorderTopStyle() (borderTopStyle string) {
	js.Rewrite("$_.borderTopStyle")
	return borderTopStyle
}

// SetBorderTopStyle prop
// js:"borderTopStyle"
func (*CSSStyleDeclaration) SetBorderTopStyle(borderTopStyle string) {
	js.Rewrite("$_.borderTopStyle = $1", borderTopStyle)
}

// BorderTopWidth prop
// js:"borderTopWidth"
func (*CSSStyleDeclaration) BorderTopWidth() (borderTopWidth string) {
	js.Rewrite("$_.borderTopWidth")
	return borderTopWidth
}

// SetBorderTopWidth prop
// js:"borderTopWidth"
func (*CSSStyleDeclaration) SetBorderTopWidth(borderTopWidth string) {
	js.Rewrite("$_.borderTopWidth = $1", borderTopWidth)
}

// BorderWidth prop
// js:"borderWidth"
func (*CSSStyleDeclaration) BorderWidth() (borderWidth string) {
	js.Rewrite("$_.borderWidth")
	return borderWidth
}

// SetBorderWidth prop
// js:"borderWidth"
func (*CSSStyleDeclaration) SetBorderWidth(borderWidth string) {
	js.Rewrite("$_.borderWidth = $1", borderWidth)
}

// Bottom prop
// js:"bottom"
func (*CSSStyleDeclaration) Bottom() (bottom string) {
	js.Rewrite("$_.bottom")
	return bottom
}

// SetBottom prop
// js:"bottom"
func (*CSSStyleDeclaration) SetBottom(bottom string) {
	js.Rewrite("$_.bottom = $1", bottom)
}

// BoxShadow prop
// js:"boxShadow"
func (*CSSStyleDeclaration) BoxShadow() (boxShadow string) {
	js.Rewrite("$_.boxShadow")
	return boxShadow
}

// SetBoxShadow prop
// js:"boxShadow"
func (*CSSStyleDeclaration) SetBoxShadow(boxShadow string) {
	js.Rewrite("$_.boxShadow = $1", boxShadow)
}

// BoxSizing prop
// js:"boxSizing"
func (*CSSStyleDeclaration) BoxSizing() (boxSizing string) {
	js.Rewrite("$_.boxSizing")
	return boxSizing
}

// SetBoxSizing prop
// js:"boxSizing"
func (*CSSStyleDeclaration) SetBoxSizing(boxSizing string) {
	js.Rewrite("$_.boxSizing = $1", boxSizing)
}

// BreakAfter prop
// js:"breakAfter"
func (*CSSStyleDeclaration) BreakAfter() (breakAfter string) {
	js.Rewrite("$_.breakAfter")
	return breakAfter
}

// SetBreakAfter prop
// js:"breakAfter"
func (*CSSStyleDeclaration) SetBreakAfter(breakAfter string) {
	js.Rewrite("$_.breakAfter = $1", breakAfter)
}

// BreakBefore prop
// js:"breakBefore"
func (*CSSStyleDeclaration) BreakBefore() (breakBefore string) {
	js.Rewrite("$_.breakBefore")
	return breakBefore
}

// SetBreakBefore prop
// js:"breakBefore"
func (*CSSStyleDeclaration) SetBreakBefore(breakBefore string) {
	js.Rewrite("$_.breakBefore = $1", breakBefore)
}

// BreakInside prop
// js:"breakInside"
func (*CSSStyleDeclaration) BreakInside() (breakInside string) {
	js.Rewrite("$_.breakInside")
	return breakInside
}

// SetBreakInside prop
// js:"breakInside"
func (*CSSStyleDeclaration) SetBreakInside(breakInside string) {
	js.Rewrite("$_.breakInside = $1", breakInside)
}

// CaptionSide prop
// js:"captionSide"
func (*CSSStyleDeclaration) CaptionSide() (captionSide string) {
	js.Rewrite("$_.captionSide")
	return captionSide
}

// SetCaptionSide prop
// js:"captionSide"
func (*CSSStyleDeclaration) SetCaptionSide(captionSide string) {
	js.Rewrite("$_.captionSide = $1", captionSide)
}

// Clear prop
// js:"clear"
func (*CSSStyleDeclaration) Clear() (clear string) {
	js.Rewrite("$_.clear")
	return clear
}

// SetClear prop
// js:"clear"
func (*CSSStyleDeclaration) SetClear(clear string) {
	js.Rewrite("$_.clear = $1", clear)
}

// Clip prop
// js:"clip"
func (*CSSStyleDeclaration) Clip() (clip string) {
	js.Rewrite("$_.clip")
	return clip
}

// SetClip prop
// js:"clip"
func (*CSSStyleDeclaration) SetClip(clip string) {
	js.Rewrite("$_.clip = $1", clip)
}

// ClipPath prop
// js:"clipPath"
func (*CSSStyleDeclaration) ClipPath() (clipPath string) {
	js.Rewrite("$_.clipPath")
	return clipPath
}

// SetClipPath prop
// js:"clipPath"
func (*CSSStyleDeclaration) SetClipPath(clipPath string) {
	js.Rewrite("$_.clipPath = $1", clipPath)
}

// ClipRule prop
// js:"clipRule"
func (*CSSStyleDeclaration) ClipRule() (clipRule string) {
	js.Rewrite("$_.clipRule")
	return clipRule
}

// SetClipRule prop
// js:"clipRule"
func (*CSSStyleDeclaration) SetClipRule(clipRule string) {
	js.Rewrite("$_.clipRule = $1", clipRule)
}

// Color prop
// js:"color"
func (*CSSStyleDeclaration) Color() (color string) {
	js.Rewrite("$_.color")
	return color
}

// SetColor prop
// js:"color"
func (*CSSStyleDeclaration) SetColor(color string) {
	js.Rewrite("$_.color = $1", color)
}

// ColorInterpolationFilters prop
// js:"colorInterpolationFilters"
func (*CSSStyleDeclaration) ColorInterpolationFilters() (colorInterpolationFilters string) {
	js.Rewrite("$_.colorInterpolationFilters")
	return colorInterpolationFilters
}

// SetColorInterpolationFilters prop
// js:"colorInterpolationFilters"
func (*CSSStyleDeclaration) SetColorInterpolationFilters(colorInterpolationFilters string) {
	js.Rewrite("$_.colorInterpolationFilters = $1", colorInterpolationFilters)
}

// ColumnCount prop
// js:"columnCount"
func (*CSSStyleDeclaration) ColumnCount() (columnCount interface{}) {
	js.Rewrite("$_.columnCount")
	return columnCount
}

// SetColumnCount prop
// js:"columnCount"
func (*CSSStyleDeclaration) SetColumnCount(columnCount interface{}) {
	js.Rewrite("$_.columnCount = $1", columnCount)
}

// ColumnFill prop
// js:"columnFill"
func (*CSSStyleDeclaration) ColumnFill() (columnFill string) {
	js.Rewrite("$_.columnFill")
	return columnFill
}

// SetColumnFill prop
// js:"columnFill"
func (*CSSStyleDeclaration) SetColumnFill(columnFill string) {
	js.Rewrite("$_.columnFill = $1", columnFill)
}

// ColumnGap prop
// js:"columnGap"
func (*CSSStyleDeclaration) ColumnGap() (columnGap interface{}) {
	js.Rewrite("$_.columnGap")
	return columnGap
}

// SetColumnGap prop
// js:"columnGap"
func (*CSSStyleDeclaration) SetColumnGap(columnGap interface{}) {
	js.Rewrite("$_.columnGap = $1", columnGap)
}

// ColumnRule prop
// js:"columnRule"
func (*CSSStyleDeclaration) ColumnRule() (columnRule string) {
	js.Rewrite("$_.columnRule")
	return columnRule
}

// SetColumnRule prop
// js:"columnRule"
func (*CSSStyleDeclaration) SetColumnRule(columnRule string) {
	js.Rewrite("$_.columnRule = $1", columnRule)
}

// ColumnRuleColor prop
// js:"columnRuleColor"
func (*CSSStyleDeclaration) ColumnRuleColor() (columnRuleColor interface{}) {
	js.Rewrite("$_.columnRuleColor")
	return columnRuleColor
}

// SetColumnRuleColor prop
// js:"columnRuleColor"
func (*CSSStyleDeclaration) SetColumnRuleColor(columnRuleColor interface{}) {
	js.Rewrite("$_.columnRuleColor = $1", columnRuleColor)
}

// ColumnRuleStyle prop
// js:"columnRuleStyle"
func (*CSSStyleDeclaration) ColumnRuleStyle() (columnRuleStyle string) {
	js.Rewrite("$_.columnRuleStyle")
	return columnRuleStyle
}

// SetColumnRuleStyle prop
// js:"columnRuleStyle"
func (*CSSStyleDeclaration) SetColumnRuleStyle(columnRuleStyle string) {
	js.Rewrite("$_.columnRuleStyle = $1", columnRuleStyle)
}

// ColumnRuleWidth prop
// js:"columnRuleWidth"
func (*CSSStyleDeclaration) ColumnRuleWidth() (columnRuleWidth interface{}) {
	js.Rewrite("$_.columnRuleWidth")
	return columnRuleWidth
}

// SetColumnRuleWidth prop
// js:"columnRuleWidth"
func (*CSSStyleDeclaration) SetColumnRuleWidth(columnRuleWidth interface{}) {
	js.Rewrite("$_.columnRuleWidth = $1", columnRuleWidth)
}

// Columns prop
// js:"columns"
func (*CSSStyleDeclaration) Columns() (columns string) {
	js.Rewrite("$_.columns")
	return columns
}

// SetColumns prop
// js:"columns"
func (*CSSStyleDeclaration) SetColumns(columns string) {
	js.Rewrite("$_.columns = $1", columns)
}

// ColumnSpan prop
// js:"columnSpan"
func (*CSSStyleDeclaration) ColumnSpan() (columnSpan string) {
	js.Rewrite("$_.columnSpan")
	return columnSpan
}

// SetColumnSpan prop
// js:"columnSpan"
func (*CSSStyleDeclaration) SetColumnSpan(columnSpan string) {
	js.Rewrite("$_.columnSpan = $1", columnSpan)
}

// ColumnWidth prop
// js:"columnWidth"
func (*CSSStyleDeclaration) ColumnWidth() (columnWidth interface{}) {
	js.Rewrite("$_.columnWidth")
	return columnWidth
}

// SetColumnWidth prop
// js:"columnWidth"
func (*CSSStyleDeclaration) SetColumnWidth(columnWidth interface{}) {
	js.Rewrite("$_.columnWidth = $1", columnWidth)
}

// Content prop
// js:"content"
func (*CSSStyleDeclaration) Content() (content string) {
	js.Rewrite("$_.content")
	return content
}

// SetContent prop
// js:"content"
func (*CSSStyleDeclaration) SetContent(content string) {
	js.Rewrite("$_.content = $1", content)
}

// CounterIncrement prop
// js:"counterIncrement"
func (*CSSStyleDeclaration) CounterIncrement() (counterIncrement string) {
	js.Rewrite("$_.counterIncrement")
	return counterIncrement
}

// SetCounterIncrement prop
// js:"counterIncrement"
func (*CSSStyleDeclaration) SetCounterIncrement(counterIncrement string) {
	js.Rewrite("$_.counterIncrement = $1", counterIncrement)
}

// CounterReset prop
// js:"counterReset"
func (*CSSStyleDeclaration) CounterReset() (counterReset string) {
	js.Rewrite("$_.counterReset")
	return counterReset
}

// SetCounterReset prop
// js:"counterReset"
func (*CSSStyleDeclaration) SetCounterReset(counterReset string) {
	js.Rewrite("$_.counterReset = $1", counterReset)
}

// CSSFloat prop
// js:"cssFloat"
func (*CSSStyleDeclaration) CSSFloat() (cssFloat string) {
	js.Rewrite("$_.cssFloat")
	return cssFloat
}

// SetCSSFloat prop
// js:"cssFloat"
func (*CSSStyleDeclaration) SetCSSFloat(cssFloat string) {
	js.Rewrite("$_.cssFloat = $1", cssFloat)
}

// CSSText prop
// js:"cssText"
func (*CSSStyleDeclaration) CSSText() (cssText string) {
	js.Rewrite("$_.cssText")
	return cssText
}

// SetCSSText prop
// js:"cssText"
func (*CSSStyleDeclaration) SetCSSText(cssText string) {
	js.Rewrite("$_.cssText = $1", cssText)
}

// Cursor prop
// js:"cursor"
func (*CSSStyleDeclaration) Cursor() (cursor string) {
	js.Rewrite("$_.cursor")
	return cursor
}

// SetCursor prop
// js:"cursor"
func (*CSSStyleDeclaration) SetCursor(cursor string) {
	js.Rewrite("$_.cursor = $1", cursor)
}

// Direction prop
// js:"direction"
func (*CSSStyleDeclaration) Direction() (direction string) {
	js.Rewrite("$_.direction")
	return direction
}

// SetDirection prop
// js:"direction"
func (*CSSStyleDeclaration) SetDirection(direction string) {
	js.Rewrite("$_.direction = $1", direction)
}

// Display prop
// js:"display"
func (*CSSStyleDeclaration) Display() (display string) {
	js.Rewrite("$_.display")
	return display
}

// SetDisplay prop
// js:"display"
func (*CSSStyleDeclaration) SetDisplay(display string) {
	js.Rewrite("$_.display = $1", display)
}

// DominantBaseline prop
// js:"dominantBaseline"
func (*CSSStyleDeclaration) DominantBaseline() (dominantBaseline string) {
	js.Rewrite("$_.dominantBaseline")
	return dominantBaseline
}

// SetDominantBaseline prop
// js:"dominantBaseline"
func (*CSSStyleDeclaration) SetDominantBaseline(dominantBaseline string) {
	js.Rewrite("$_.dominantBaseline = $1", dominantBaseline)
}

// EmptyCells prop
// js:"emptyCells"
func (*CSSStyleDeclaration) EmptyCells() (emptyCells string) {
	js.Rewrite("$_.emptyCells")
	return emptyCells
}

// SetEmptyCells prop
// js:"emptyCells"
func (*CSSStyleDeclaration) SetEmptyCells(emptyCells string) {
	js.Rewrite("$_.emptyCells = $1", emptyCells)
}

// EnableBackground prop
// js:"enableBackground"
func (*CSSStyleDeclaration) EnableBackground() (enableBackground string) {
	js.Rewrite("$_.enableBackground")
	return enableBackground
}

// SetEnableBackground prop
// js:"enableBackground"
func (*CSSStyleDeclaration) SetEnableBackground(enableBackground string) {
	js.Rewrite("$_.enableBackground = $1", enableBackground)
}

// Fill prop
// js:"fill"
func (*CSSStyleDeclaration) Fill() (fill string) {
	js.Rewrite("$_.fill")
	return fill
}

// SetFill prop
// js:"fill"
func (*CSSStyleDeclaration) SetFill(fill string) {
	js.Rewrite("$_.fill = $1", fill)
}

// FillOpacity prop
// js:"fillOpacity"
func (*CSSStyleDeclaration) FillOpacity() (fillOpacity string) {
	js.Rewrite("$_.fillOpacity")
	return fillOpacity
}

// SetFillOpacity prop
// js:"fillOpacity"
func (*CSSStyleDeclaration) SetFillOpacity(fillOpacity string) {
	js.Rewrite("$_.fillOpacity = $1", fillOpacity)
}

// FillRule prop
// js:"fillRule"
func (*CSSStyleDeclaration) FillRule() (fillRule string) {
	js.Rewrite("$_.fillRule")
	return fillRule
}

// SetFillRule prop
// js:"fillRule"
func (*CSSStyleDeclaration) SetFillRule(fillRule string) {
	js.Rewrite("$_.fillRule = $1", fillRule)
}

// Filter prop
// js:"filter"
func (*CSSStyleDeclaration) Filter() (filter string) {
	js.Rewrite("$_.filter")
	return filter
}

// SetFilter prop
// js:"filter"
func (*CSSStyleDeclaration) SetFilter(filter string) {
	js.Rewrite("$_.filter = $1", filter)
}

// Flex prop
// js:"flex"
func (*CSSStyleDeclaration) Flex() (flex string) {
	js.Rewrite("$_.flex")
	return flex
}

// SetFlex prop
// js:"flex"
func (*CSSStyleDeclaration) SetFlex(flex string) {
	js.Rewrite("$_.flex = $1", flex)
}

// FlexBasis prop
// js:"flexBasis"
func (*CSSStyleDeclaration) FlexBasis() (flexBasis string) {
	js.Rewrite("$_.flexBasis")
	return flexBasis
}

// SetFlexBasis prop
// js:"flexBasis"
func (*CSSStyleDeclaration) SetFlexBasis(flexBasis string) {
	js.Rewrite("$_.flexBasis = $1", flexBasis)
}

// FlexDirection prop
// js:"flexDirection"
func (*CSSStyleDeclaration) FlexDirection() (flexDirection string) {
	js.Rewrite("$_.flexDirection")
	return flexDirection
}

// SetFlexDirection prop
// js:"flexDirection"
func (*CSSStyleDeclaration) SetFlexDirection(flexDirection string) {
	js.Rewrite("$_.flexDirection = $1", flexDirection)
}

// FlexFlow prop
// js:"flexFlow"
func (*CSSStyleDeclaration) FlexFlow() (flexFlow string) {
	js.Rewrite("$_.flexFlow")
	return flexFlow
}

// SetFlexFlow prop
// js:"flexFlow"
func (*CSSStyleDeclaration) SetFlexFlow(flexFlow string) {
	js.Rewrite("$_.flexFlow = $1", flexFlow)
}

// FlexGrow prop
// js:"flexGrow"
func (*CSSStyleDeclaration) FlexGrow() (flexGrow string) {
	js.Rewrite("$_.flexGrow")
	return flexGrow
}

// SetFlexGrow prop
// js:"flexGrow"
func (*CSSStyleDeclaration) SetFlexGrow(flexGrow string) {
	js.Rewrite("$_.flexGrow = $1", flexGrow)
}

// FlexShrink prop
// js:"flexShrink"
func (*CSSStyleDeclaration) FlexShrink() (flexShrink string) {
	js.Rewrite("$_.flexShrink")
	return flexShrink
}

// SetFlexShrink prop
// js:"flexShrink"
func (*CSSStyleDeclaration) SetFlexShrink(flexShrink string) {
	js.Rewrite("$_.flexShrink = $1", flexShrink)
}

// FlexWrap prop
// js:"flexWrap"
func (*CSSStyleDeclaration) FlexWrap() (flexWrap string) {
	js.Rewrite("$_.flexWrap")
	return flexWrap
}

// SetFlexWrap prop
// js:"flexWrap"
func (*CSSStyleDeclaration) SetFlexWrap(flexWrap string) {
	js.Rewrite("$_.flexWrap = $1", flexWrap)
}

// FloodColor prop
// js:"floodColor"
func (*CSSStyleDeclaration) FloodColor() (floodColor string) {
	js.Rewrite("$_.floodColor")
	return floodColor
}

// SetFloodColor prop
// js:"floodColor"
func (*CSSStyleDeclaration) SetFloodColor(floodColor string) {
	js.Rewrite("$_.floodColor = $1", floodColor)
}

// FloodOpacity prop
// js:"floodOpacity"
func (*CSSStyleDeclaration) FloodOpacity() (floodOpacity string) {
	js.Rewrite("$_.floodOpacity")
	return floodOpacity
}

// SetFloodOpacity prop
// js:"floodOpacity"
func (*CSSStyleDeclaration) SetFloodOpacity(floodOpacity string) {
	js.Rewrite("$_.floodOpacity = $1", floodOpacity)
}

// Font prop
// js:"font"
func (*CSSStyleDeclaration) Font() (font string) {
	js.Rewrite("$_.font")
	return font
}

// SetFont prop
// js:"font"
func (*CSSStyleDeclaration) SetFont(font string) {
	js.Rewrite("$_.font = $1", font)
}

// FontFamily prop
// js:"fontFamily"
func (*CSSStyleDeclaration) FontFamily() (fontFamily string) {
	js.Rewrite("$_.fontFamily")
	return fontFamily
}

// SetFontFamily prop
// js:"fontFamily"
func (*CSSStyleDeclaration) SetFontFamily(fontFamily string) {
	js.Rewrite("$_.fontFamily = $1", fontFamily)
}

// FontFeatureSettings prop
// js:"fontFeatureSettings"
func (*CSSStyleDeclaration) FontFeatureSettings() (fontFeatureSettings string) {
	js.Rewrite("$_.fontFeatureSettings")
	return fontFeatureSettings
}

// SetFontFeatureSettings prop
// js:"fontFeatureSettings"
func (*CSSStyleDeclaration) SetFontFeatureSettings(fontFeatureSettings string) {
	js.Rewrite("$_.fontFeatureSettings = $1", fontFeatureSettings)
}

// FontSize prop
// js:"fontSize"
func (*CSSStyleDeclaration) FontSize() (fontSize string) {
	js.Rewrite("$_.fontSize")
	return fontSize
}

// SetFontSize prop
// js:"fontSize"
func (*CSSStyleDeclaration) SetFontSize(fontSize string) {
	js.Rewrite("$_.fontSize = $1", fontSize)
}

// FontSizeAdjust prop
// js:"fontSizeAdjust"
func (*CSSStyleDeclaration) FontSizeAdjust() (fontSizeAdjust string) {
	js.Rewrite("$_.fontSizeAdjust")
	return fontSizeAdjust
}

// SetFontSizeAdjust prop
// js:"fontSizeAdjust"
func (*CSSStyleDeclaration) SetFontSizeAdjust(fontSizeAdjust string) {
	js.Rewrite("$_.fontSizeAdjust = $1", fontSizeAdjust)
}

// FontStretch prop
// js:"fontStretch"
func (*CSSStyleDeclaration) FontStretch() (fontStretch string) {
	js.Rewrite("$_.fontStretch")
	return fontStretch
}

// SetFontStretch prop
// js:"fontStretch"
func (*CSSStyleDeclaration) SetFontStretch(fontStretch string) {
	js.Rewrite("$_.fontStretch = $1", fontStretch)
}

// FontStyle prop
// js:"fontStyle"
func (*CSSStyleDeclaration) FontStyle() (fontStyle string) {
	js.Rewrite("$_.fontStyle")
	return fontStyle
}

// SetFontStyle prop
// js:"fontStyle"
func (*CSSStyleDeclaration) SetFontStyle(fontStyle string) {
	js.Rewrite("$_.fontStyle = $1", fontStyle)
}

// FontVariant prop
// js:"fontVariant"
func (*CSSStyleDeclaration) FontVariant() (fontVariant string) {
	js.Rewrite("$_.fontVariant")
	return fontVariant
}

// SetFontVariant prop
// js:"fontVariant"
func (*CSSStyleDeclaration) SetFontVariant(fontVariant string) {
	js.Rewrite("$_.fontVariant = $1", fontVariant)
}

// FontWeight prop
// js:"fontWeight"
func (*CSSStyleDeclaration) FontWeight() (fontWeight string) {
	js.Rewrite("$_.fontWeight")
	return fontWeight
}

// SetFontWeight prop
// js:"fontWeight"
func (*CSSStyleDeclaration) SetFontWeight(fontWeight string) {
	js.Rewrite("$_.fontWeight = $1", fontWeight)
}

// GlyphOrientationHorizontal prop
// js:"glyphOrientationHorizontal"
func (*CSSStyleDeclaration) GlyphOrientationHorizontal() (glyphOrientationHorizontal string) {
	js.Rewrite("$_.glyphOrientationHorizontal")
	return glyphOrientationHorizontal
}

// SetGlyphOrientationHorizontal prop
// js:"glyphOrientationHorizontal"
func (*CSSStyleDeclaration) SetGlyphOrientationHorizontal(glyphOrientationHorizontal string) {
	js.Rewrite("$_.glyphOrientationHorizontal = $1", glyphOrientationHorizontal)
}

// GlyphOrientationVertical prop
// js:"glyphOrientationVertical"
func (*CSSStyleDeclaration) GlyphOrientationVertical() (glyphOrientationVertical string) {
	js.Rewrite("$_.glyphOrientationVertical")
	return glyphOrientationVertical
}

// SetGlyphOrientationVertical prop
// js:"glyphOrientationVertical"
func (*CSSStyleDeclaration) SetGlyphOrientationVertical(glyphOrientationVertical string) {
	js.Rewrite("$_.glyphOrientationVertical = $1", glyphOrientationVertical)
}

// Height prop
// js:"height"
func (*CSSStyleDeclaration) Height() (height string) {
	js.Rewrite("$_.height")
	return height
}

// SetHeight prop
// js:"height"
func (*CSSStyleDeclaration) SetHeight(height string) {
	js.Rewrite("$_.height = $1", height)
}

// ImeMode prop
// js:"imeMode"
func (*CSSStyleDeclaration) ImeMode() (imeMode string) {
	js.Rewrite("$_.imeMode")
	return imeMode
}

// SetImeMode prop
// js:"imeMode"
func (*CSSStyleDeclaration) SetImeMode(imeMode string) {
	js.Rewrite("$_.imeMode = $1", imeMode)
}

// JustifyContent prop
// js:"justifyContent"
func (*CSSStyleDeclaration) JustifyContent() (justifyContent string) {
	js.Rewrite("$_.justifyContent")
	return justifyContent
}

// SetJustifyContent prop
// js:"justifyContent"
func (*CSSStyleDeclaration) SetJustifyContent(justifyContent string) {
	js.Rewrite("$_.justifyContent = $1", justifyContent)
}

// Kerning prop
// js:"kerning"
func (*CSSStyleDeclaration) Kerning() (kerning string) {
	js.Rewrite("$_.kerning")
	return kerning
}

// SetKerning prop
// js:"kerning"
func (*CSSStyleDeclaration) SetKerning(kerning string) {
	js.Rewrite("$_.kerning = $1", kerning)
}

// LayoutGrid prop
// js:"layoutGrid"
func (*CSSStyleDeclaration) LayoutGrid() (layoutGrid string) {
	js.Rewrite("$_.layoutGrid")
	return layoutGrid
}

// SetLayoutGrid prop
// js:"layoutGrid"
func (*CSSStyleDeclaration) SetLayoutGrid(layoutGrid string) {
	js.Rewrite("$_.layoutGrid = $1", layoutGrid)
}

// LayoutGridChar prop
// js:"layoutGridChar"
func (*CSSStyleDeclaration) LayoutGridChar() (layoutGridChar string) {
	js.Rewrite("$_.layoutGridChar")
	return layoutGridChar
}

// SetLayoutGridChar prop
// js:"layoutGridChar"
func (*CSSStyleDeclaration) SetLayoutGridChar(layoutGridChar string) {
	js.Rewrite("$_.layoutGridChar = $1", layoutGridChar)
}

// LayoutGridLine prop
// js:"layoutGridLine"
func (*CSSStyleDeclaration) LayoutGridLine() (layoutGridLine string) {
	js.Rewrite("$_.layoutGridLine")
	return layoutGridLine
}

// SetLayoutGridLine prop
// js:"layoutGridLine"
func (*CSSStyleDeclaration) SetLayoutGridLine(layoutGridLine string) {
	js.Rewrite("$_.layoutGridLine = $1", layoutGridLine)
}

// LayoutGridMode prop
// js:"layoutGridMode"
func (*CSSStyleDeclaration) LayoutGridMode() (layoutGridMode string) {
	js.Rewrite("$_.layoutGridMode")
	return layoutGridMode
}

// SetLayoutGridMode prop
// js:"layoutGridMode"
func (*CSSStyleDeclaration) SetLayoutGridMode(layoutGridMode string) {
	js.Rewrite("$_.layoutGridMode = $1", layoutGridMode)
}

// LayoutGridType prop
// js:"layoutGridType"
func (*CSSStyleDeclaration) LayoutGridType() (layoutGridType string) {
	js.Rewrite("$_.layoutGridType")
	return layoutGridType
}

// SetLayoutGridType prop
// js:"layoutGridType"
func (*CSSStyleDeclaration) SetLayoutGridType(layoutGridType string) {
	js.Rewrite("$_.layoutGridType = $1", layoutGridType)
}

// Left prop
// js:"left"
func (*CSSStyleDeclaration) Left() (left string) {
	js.Rewrite("$_.left")
	return left
}

// SetLeft prop
// js:"left"
func (*CSSStyleDeclaration) SetLeft(left string) {
	js.Rewrite("$_.left = $1", left)
}

// Length prop
// js:"length"
func (*CSSStyleDeclaration) Length() (length uint) {
	js.Rewrite("$_.length")
	return length
}

// LetterSpacing prop
// js:"letterSpacing"
func (*CSSStyleDeclaration) LetterSpacing() (letterSpacing string) {
	js.Rewrite("$_.letterSpacing")
	return letterSpacing
}

// SetLetterSpacing prop
// js:"letterSpacing"
func (*CSSStyleDeclaration) SetLetterSpacing(letterSpacing string) {
	js.Rewrite("$_.letterSpacing = $1", letterSpacing)
}

// LightingColor prop
// js:"lightingColor"
func (*CSSStyleDeclaration) LightingColor() (lightingColor string) {
	js.Rewrite("$_.lightingColor")
	return lightingColor
}

// SetLightingColor prop
// js:"lightingColor"
func (*CSSStyleDeclaration) SetLightingColor(lightingColor string) {
	js.Rewrite("$_.lightingColor = $1", lightingColor)
}

// LineBreak prop
// js:"lineBreak"
func (*CSSStyleDeclaration) LineBreak() (lineBreak string) {
	js.Rewrite("$_.lineBreak")
	return lineBreak
}

// SetLineBreak prop
// js:"lineBreak"
func (*CSSStyleDeclaration) SetLineBreak(lineBreak string) {
	js.Rewrite("$_.lineBreak = $1", lineBreak)
}

// LineHeight prop
// js:"lineHeight"
func (*CSSStyleDeclaration) LineHeight() (lineHeight string) {
	js.Rewrite("$_.lineHeight")
	return lineHeight
}

// SetLineHeight prop
// js:"lineHeight"
func (*CSSStyleDeclaration) SetLineHeight(lineHeight string) {
	js.Rewrite("$_.lineHeight = $1", lineHeight)
}

// ListStyle prop
// js:"listStyle"
func (*CSSStyleDeclaration) ListStyle() (listStyle string) {
	js.Rewrite("$_.listStyle")
	return listStyle
}

// SetListStyle prop
// js:"listStyle"
func (*CSSStyleDeclaration) SetListStyle(listStyle string) {
	js.Rewrite("$_.listStyle = $1", listStyle)
}

// ListStyleImage prop
// js:"listStyleImage"
func (*CSSStyleDeclaration) ListStyleImage() (listStyleImage string) {
	js.Rewrite("$_.listStyleImage")
	return listStyleImage
}

// SetListStyleImage prop
// js:"listStyleImage"
func (*CSSStyleDeclaration) SetListStyleImage(listStyleImage string) {
	js.Rewrite("$_.listStyleImage = $1", listStyleImage)
}

// ListStylePosition prop
// js:"listStylePosition"
func (*CSSStyleDeclaration) ListStylePosition() (listStylePosition string) {
	js.Rewrite("$_.listStylePosition")
	return listStylePosition
}

// SetListStylePosition prop
// js:"listStylePosition"
func (*CSSStyleDeclaration) SetListStylePosition(listStylePosition string) {
	js.Rewrite("$_.listStylePosition = $1", listStylePosition)
}

// ListStyleType prop
// js:"listStyleType"
func (*CSSStyleDeclaration) ListStyleType() (listStyleType string) {
	js.Rewrite("$_.listStyleType")
	return listStyleType
}

// SetListStyleType prop
// js:"listStyleType"
func (*CSSStyleDeclaration) SetListStyleType(listStyleType string) {
	js.Rewrite("$_.listStyleType = $1", listStyleType)
}

// Margin prop
// js:"margin"
func (*CSSStyleDeclaration) Margin() (margin string) {
	js.Rewrite("$_.margin")
	return margin
}

// SetMargin prop
// js:"margin"
func (*CSSStyleDeclaration) SetMargin(margin string) {
	js.Rewrite("$_.margin = $1", margin)
}

// MarginBottom prop
// js:"marginBottom"
func (*CSSStyleDeclaration) MarginBottom() (marginBottom string) {
	js.Rewrite("$_.marginBottom")
	return marginBottom
}

// SetMarginBottom prop
// js:"marginBottom"
func (*CSSStyleDeclaration) SetMarginBottom(marginBottom string) {
	js.Rewrite("$_.marginBottom = $1", marginBottom)
}

// MarginLeft prop
// js:"marginLeft"
func (*CSSStyleDeclaration) MarginLeft() (marginLeft string) {
	js.Rewrite("$_.marginLeft")
	return marginLeft
}

// SetMarginLeft prop
// js:"marginLeft"
func (*CSSStyleDeclaration) SetMarginLeft(marginLeft string) {
	js.Rewrite("$_.marginLeft = $1", marginLeft)
}

// MarginRight prop
// js:"marginRight"
func (*CSSStyleDeclaration) MarginRight() (marginRight string) {
	js.Rewrite("$_.marginRight")
	return marginRight
}

// SetMarginRight prop
// js:"marginRight"
func (*CSSStyleDeclaration) SetMarginRight(marginRight string) {
	js.Rewrite("$_.marginRight = $1", marginRight)
}

// MarginTop prop
// js:"marginTop"
func (*CSSStyleDeclaration) MarginTop() (marginTop string) {
	js.Rewrite("$_.marginTop")
	return marginTop
}

// SetMarginTop prop
// js:"marginTop"
func (*CSSStyleDeclaration) SetMarginTop(marginTop string) {
	js.Rewrite("$_.marginTop = $1", marginTop)
}

// Marker prop
// js:"marker"
func (*CSSStyleDeclaration) Marker() (marker string) {
	js.Rewrite("$_.marker")
	return marker
}

// SetMarker prop
// js:"marker"
func (*CSSStyleDeclaration) SetMarker(marker string) {
	js.Rewrite("$_.marker = $1", marker)
}

// MarkerEnd prop
// js:"markerEnd"
func (*CSSStyleDeclaration) MarkerEnd() (markerEnd string) {
	js.Rewrite("$_.markerEnd")
	return markerEnd
}

// SetMarkerEnd prop
// js:"markerEnd"
func (*CSSStyleDeclaration) SetMarkerEnd(markerEnd string) {
	js.Rewrite("$_.markerEnd = $1", markerEnd)
}

// MarkerMid prop
// js:"markerMid"
func (*CSSStyleDeclaration) MarkerMid() (markerMid string) {
	js.Rewrite("$_.markerMid")
	return markerMid
}

// SetMarkerMid prop
// js:"markerMid"
func (*CSSStyleDeclaration) SetMarkerMid(markerMid string) {
	js.Rewrite("$_.markerMid = $1", markerMid)
}

// MarkerStart prop
// js:"markerStart"
func (*CSSStyleDeclaration) MarkerStart() (markerStart string) {
	js.Rewrite("$_.markerStart")
	return markerStart
}

// SetMarkerStart prop
// js:"markerStart"
func (*CSSStyleDeclaration) SetMarkerStart(markerStart string) {
	js.Rewrite("$_.markerStart = $1", markerStart)
}

// Mask prop
// js:"mask"
func (*CSSStyleDeclaration) Mask() (mask string) {
	js.Rewrite("$_.mask")
	return mask
}

// SetMask prop
// js:"mask"
func (*CSSStyleDeclaration) SetMask(mask string) {
	js.Rewrite("$_.mask = $1", mask)
}

// MaxHeight prop
// js:"maxHeight"
func (*CSSStyleDeclaration) MaxHeight() (maxHeight string) {
	js.Rewrite("$_.maxHeight")
	return maxHeight
}

// SetMaxHeight prop
// js:"maxHeight"
func (*CSSStyleDeclaration) SetMaxHeight(maxHeight string) {
	js.Rewrite("$_.maxHeight = $1", maxHeight)
}

// MaxWidth prop
// js:"maxWidth"
func (*CSSStyleDeclaration) MaxWidth() (maxWidth string) {
	js.Rewrite("$_.maxWidth")
	return maxWidth
}

// SetMaxWidth prop
// js:"maxWidth"
func (*CSSStyleDeclaration) SetMaxWidth(maxWidth string) {
	js.Rewrite("$_.maxWidth = $1", maxWidth)
}

// MinHeight prop
// js:"minHeight"
func (*CSSStyleDeclaration) MinHeight() (minHeight string) {
	js.Rewrite("$_.minHeight")
	return minHeight
}

// SetMinHeight prop
// js:"minHeight"
func (*CSSStyleDeclaration) SetMinHeight(minHeight string) {
	js.Rewrite("$_.minHeight = $1", minHeight)
}

// MinWidth prop
// js:"minWidth"
func (*CSSStyleDeclaration) MinWidth() (minWidth string) {
	js.Rewrite("$_.minWidth")
	return minWidth
}

// SetMinWidth prop
// js:"minWidth"
func (*CSSStyleDeclaration) SetMinWidth(minWidth string) {
	js.Rewrite("$_.minWidth = $1", minWidth)
}

// MsContentZoomChaining prop
// js:"msContentZoomChaining"
func (*CSSStyleDeclaration) MsContentZoomChaining() (msContentZoomChaining string) {
	js.Rewrite("$_.msContentZoomChaining")
	return msContentZoomChaining
}

// SetMsContentZoomChaining prop
// js:"msContentZoomChaining"
func (*CSSStyleDeclaration) SetMsContentZoomChaining(msContentZoomChaining string) {
	js.Rewrite("$_.msContentZoomChaining = $1", msContentZoomChaining)
}

// MsContentZooming prop
// js:"msContentZooming"
func (*CSSStyleDeclaration) MsContentZooming() (msContentZooming string) {
	js.Rewrite("$_.msContentZooming")
	return msContentZooming
}

// SetMsContentZooming prop
// js:"msContentZooming"
func (*CSSStyleDeclaration) SetMsContentZooming(msContentZooming string) {
	js.Rewrite("$_.msContentZooming = $1", msContentZooming)
}

// MsContentZoomLimit prop
// js:"msContentZoomLimit"
func (*CSSStyleDeclaration) MsContentZoomLimit() (msContentZoomLimit string) {
	js.Rewrite("$_.msContentZoomLimit")
	return msContentZoomLimit
}

// SetMsContentZoomLimit prop
// js:"msContentZoomLimit"
func (*CSSStyleDeclaration) SetMsContentZoomLimit(msContentZoomLimit string) {
	js.Rewrite("$_.msContentZoomLimit = $1", msContentZoomLimit)
}

// MsContentZoomLimitMax prop
// js:"msContentZoomLimitMax"
func (*CSSStyleDeclaration) MsContentZoomLimitMax() (msContentZoomLimitMax interface{}) {
	js.Rewrite("$_.msContentZoomLimitMax")
	return msContentZoomLimitMax
}

// SetMsContentZoomLimitMax prop
// js:"msContentZoomLimitMax"
func (*CSSStyleDeclaration) SetMsContentZoomLimitMax(msContentZoomLimitMax interface{}) {
	js.Rewrite("$_.msContentZoomLimitMax = $1", msContentZoomLimitMax)
}

// MsContentZoomLimitMin prop
// js:"msContentZoomLimitMin"
func (*CSSStyleDeclaration) MsContentZoomLimitMin() (msContentZoomLimitMin interface{}) {
	js.Rewrite("$_.msContentZoomLimitMin")
	return msContentZoomLimitMin
}

// SetMsContentZoomLimitMin prop
// js:"msContentZoomLimitMin"
func (*CSSStyleDeclaration) SetMsContentZoomLimitMin(msContentZoomLimitMin interface{}) {
	js.Rewrite("$_.msContentZoomLimitMin = $1", msContentZoomLimitMin)
}

// MsContentZoomSnap prop
// js:"msContentZoomSnap"
func (*CSSStyleDeclaration) MsContentZoomSnap() (msContentZoomSnap string) {
	js.Rewrite("$_.msContentZoomSnap")
	return msContentZoomSnap
}

// SetMsContentZoomSnap prop
// js:"msContentZoomSnap"
func (*CSSStyleDeclaration) SetMsContentZoomSnap(msContentZoomSnap string) {
	js.Rewrite("$_.msContentZoomSnap = $1", msContentZoomSnap)
}

// MsContentZoomSnapPoints prop
// js:"msContentZoomSnapPoints"
func (*CSSStyleDeclaration) MsContentZoomSnapPoints() (msContentZoomSnapPoints string) {
	js.Rewrite("$_.msContentZoomSnapPoints")
	return msContentZoomSnapPoints
}

// SetMsContentZoomSnapPoints prop
// js:"msContentZoomSnapPoints"
func (*CSSStyleDeclaration) SetMsContentZoomSnapPoints(msContentZoomSnapPoints string) {
	js.Rewrite("$_.msContentZoomSnapPoints = $1", msContentZoomSnapPoints)
}

// MsContentZoomSnapType prop
// js:"msContentZoomSnapType"
func (*CSSStyleDeclaration) MsContentZoomSnapType() (msContentZoomSnapType string) {
	js.Rewrite("$_.msContentZoomSnapType")
	return msContentZoomSnapType
}

// SetMsContentZoomSnapType prop
// js:"msContentZoomSnapType"
func (*CSSStyleDeclaration) SetMsContentZoomSnapType(msContentZoomSnapType string) {
	js.Rewrite("$_.msContentZoomSnapType = $1", msContentZoomSnapType)
}

// MsFlowFrom prop
// js:"msFlowFrom"
func (*CSSStyleDeclaration) MsFlowFrom() (msFlowFrom string) {
	js.Rewrite("$_.msFlowFrom")
	return msFlowFrom
}

// SetMsFlowFrom prop
// js:"msFlowFrom"
func (*CSSStyleDeclaration) SetMsFlowFrom(msFlowFrom string) {
	js.Rewrite("$_.msFlowFrom = $1", msFlowFrom)
}

// MsFlowInto prop
// js:"msFlowInto"
func (*CSSStyleDeclaration) MsFlowInto() (msFlowInto string) {
	js.Rewrite("$_.msFlowInto")
	return msFlowInto
}

// SetMsFlowInto prop
// js:"msFlowInto"
func (*CSSStyleDeclaration) SetMsFlowInto(msFlowInto string) {
	js.Rewrite("$_.msFlowInto = $1", msFlowInto)
}

// MsFontFeatureSettings prop
// js:"msFontFeatureSettings"
func (*CSSStyleDeclaration) MsFontFeatureSettings() (msFontFeatureSettings string) {
	js.Rewrite("$_.msFontFeatureSettings")
	return msFontFeatureSettings
}

// SetMsFontFeatureSettings prop
// js:"msFontFeatureSettings"
func (*CSSStyleDeclaration) SetMsFontFeatureSettings(msFontFeatureSettings string) {
	js.Rewrite("$_.msFontFeatureSettings = $1", msFontFeatureSettings)
}

// MsGridColumn prop
// js:"msGridColumn"
func (*CSSStyleDeclaration) MsGridColumn() (msGridColumn interface{}) {
	js.Rewrite("$_.msGridColumn")
	return msGridColumn
}

// SetMsGridColumn prop
// js:"msGridColumn"
func (*CSSStyleDeclaration) SetMsGridColumn(msGridColumn interface{}) {
	js.Rewrite("$_.msGridColumn = $1", msGridColumn)
}

// MsGridColumnAlign prop
// js:"msGridColumnAlign"
func (*CSSStyleDeclaration) MsGridColumnAlign() (msGridColumnAlign string) {
	js.Rewrite("$_.msGridColumnAlign")
	return msGridColumnAlign
}

// SetMsGridColumnAlign prop
// js:"msGridColumnAlign"
func (*CSSStyleDeclaration) SetMsGridColumnAlign(msGridColumnAlign string) {
	js.Rewrite("$_.msGridColumnAlign = $1", msGridColumnAlign)
}

// MsGridColumns prop
// js:"msGridColumns"
func (*CSSStyleDeclaration) MsGridColumns() (msGridColumns string) {
	js.Rewrite("$_.msGridColumns")
	return msGridColumns
}

// SetMsGridColumns prop
// js:"msGridColumns"
func (*CSSStyleDeclaration) SetMsGridColumns(msGridColumns string) {
	js.Rewrite("$_.msGridColumns = $1", msGridColumns)
}

// MsGridColumnSpan prop
// js:"msGridColumnSpan"
func (*CSSStyleDeclaration) MsGridColumnSpan() (msGridColumnSpan interface{}) {
	js.Rewrite("$_.msGridColumnSpan")
	return msGridColumnSpan
}

// SetMsGridColumnSpan prop
// js:"msGridColumnSpan"
func (*CSSStyleDeclaration) SetMsGridColumnSpan(msGridColumnSpan interface{}) {
	js.Rewrite("$_.msGridColumnSpan = $1", msGridColumnSpan)
}

// MsGridRow prop
// js:"msGridRow"
func (*CSSStyleDeclaration) MsGridRow() (msGridRow interface{}) {
	js.Rewrite("$_.msGridRow")
	return msGridRow
}

// SetMsGridRow prop
// js:"msGridRow"
func (*CSSStyleDeclaration) SetMsGridRow(msGridRow interface{}) {
	js.Rewrite("$_.msGridRow = $1", msGridRow)
}

// MsGridRowAlign prop
// js:"msGridRowAlign"
func (*CSSStyleDeclaration) MsGridRowAlign() (msGridRowAlign string) {
	js.Rewrite("$_.msGridRowAlign")
	return msGridRowAlign
}

// SetMsGridRowAlign prop
// js:"msGridRowAlign"
func (*CSSStyleDeclaration) SetMsGridRowAlign(msGridRowAlign string) {
	js.Rewrite("$_.msGridRowAlign = $1", msGridRowAlign)
}

// MsGridRows prop
// js:"msGridRows"
func (*CSSStyleDeclaration) MsGridRows() (msGridRows string) {
	js.Rewrite("$_.msGridRows")
	return msGridRows
}

// SetMsGridRows prop
// js:"msGridRows"
func (*CSSStyleDeclaration) SetMsGridRows(msGridRows string) {
	js.Rewrite("$_.msGridRows = $1", msGridRows)
}

// MsGridRowSpan prop
// js:"msGridRowSpan"
func (*CSSStyleDeclaration) MsGridRowSpan() (msGridRowSpan interface{}) {
	js.Rewrite("$_.msGridRowSpan")
	return msGridRowSpan
}

// SetMsGridRowSpan prop
// js:"msGridRowSpan"
func (*CSSStyleDeclaration) SetMsGridRowSpan(msGridRowSpan interface{}) {
	js.Rewrite("$_.msGridRowSpan = $1", msGridRowSpan)
}

// MsHighContrastAdjust prop
// js:"msHighContrastAdjust"
func (*CSSStyleDeclaration) MsHighContrastAdjust() (msHighContrastAdjust string) {
	js.Rewrite("$_.msHighContrastAdjust")
	return msHighContrastAdjust
}

// SetMsHighContrastAdjust prop
// js:"msHighContrastAdjust"
func (*CSSStyleDeclaration) SetMsHighContrastAdjust(msHighContrastAdjust string) {
	js.Rewrite("$_.msHighContrastAdjust = $1", msHighContrastAdjust)
}

// MsHyphenateLimitChars prop
// js:"msHyphenateLimitChars"
func (*CSSStyleDeclaration) MsHyphenateLimitChars() (msHyphenateLimitChars string) {
	js.Rewrite("$_.msHyphenateLimitChars")
	return msHyphenateLimitChars
}

// SetMsHyphenateLimitChars prop
// js:"msHyphenateLimitChars"
func (*CSSStyleDeclaration) SetMsHyphenateLimitChars(msHyphenateLimitChars string) {
	js.Rewrite("$_.msHyphenateLimitChars = $1", msHyphenateLimitChars)
}

// MsHyphenateLimitLines prop
// js:"msHyphenateLimitLines"
func (*CSSStyleDeclaration) MsHyphenateLimitLines() (msHyphenateLimitLines interface{}) {
	js.Rewrite("$_.msHyphenateLimitLines")
	return msHyphenateLimitLines
}

// SetMsHyphenateLimitLines prop
// js:"msHyphenateLimitLines"
func (*CSSStyleDeclaration) SetMsHyphenateLimitLines(msHyphenateLimitLines interface{}) {
	js.Rewrite("$_.msHyphenateLimitLines = $1", msHyphenateLimitLines)
}

// MsHyphenateLimitZone prop
// js:"msHyphenateLimitZone"
func (*CSSStyleDeclaration) MsHyphenateLimitZone() (msHyphenateLimitZone interface{}) {
	js.Rewrite("$_.msHyphenateLimitZone")
	return msHyphenateLimitZone
}

// SetMsHyphenateLimitZone prop
// js:"msHyphenateLimitZone"
func (*CSSStyleDeclaration) SetMsHyphenateLimitZone(msHyphenateLimitZone interface{}) {
	js.Rewrite("$_.msHyphenateLimitZone = $1", msHyphenateLimitZone)
}

// MsHyphens prop
// js:"msHyphens"
func (*CSSStyleDeclaration) MsHyphens() (msHyphens string) {
	js.Rewrite("$_.msHyphens")
	return msHyphens
}

// SetMsHyphens prop
// js:"msHyphens"
func (*CSSStyleDeclaration) SetMsHyphens(msHyphens string) {
	js.Rewrite("$_.msHyphens = $1", msHyphens)
}

// MsImeAlign prop
// js:"msImeAlign"
func (*CSSStyleDeclaration) MsImeAlign() (msImeAlign string) {
	js.Rewrite("$_.msImeAlign")
	return msImeAlign
}

// SetMsImeAlign prop
// js:"msImeAlign"
func (*CSSStyleDeclaration) SetMsImeAlign(msImeAlign string) {
	js.Rewrite("$_.msImeAlign = $1", msImeAlign)
}

// MsOverflowStyle prop
// js:"msOverflowStyle"
func (*CSSStyleDeclaration) MsOverflowStyle() (msOverflowStyle string) {
	js.Rewrite("$_.msOverflowStyle")
	return msOverflowStyle
}

// SetMsOverflowStyle prop
// js:"msOverflowStyle"
func (*CSSStyleDeclaration) SetMsOverflowStyle(msOverflowStyle string) {
	js.Rewrite("$_.msOverflowStyle = $1", msOverflowStyle)
}

// MsScrollChaining prop
// js:"msScrollChaining"
func (*CSSStyleDeclaration) MsScrollChaining() (msScrollChaining string) {
	js.Rewrite("$_.msScrollChaining")
	return msScrollChaining
}

// SetMsScrollChaining prop
// js:"msScrollChaining"
func (*CSSStyleDeclaration) SetMsScrollChaining(msScrollChaining string) {
	js.Rewrite("$_.msScrollChaining = $1", msScrollChaining)
}

// MsScrollLimit prop
// js:"msScrollLimit"
func (*CSSStyleDeclaration) MsScrollLimit() (msScrollLimit string) {
	js.Rewrite("$_.msScrollLimit")
	return msScrollLimit
}

// SetMsScrollLimit prop
// js:"msScrollLimit"
func (*CSSStyleDeclaration) SetMsScrollLimit(msScrollLimit string) {
	js.Rewrite("$_.msScrollLimit = $1", msScrollLimit)
}

// MsScrollLimitXMax prop
// js:"msScrollLimitXMax"
func (*CSSStyleDeclaration) MsScrollLimitXMax() (msScrollLimitXMax interface{}) {
	js.Rewrite("$_.msScrollLimitXMax")
	return msScrollLimitXMax
}

// SetMsScrollLimitXMax prop
// js:"msScrollLimitXMax"
func (*CSSStyleDeclaration) SetMsScrollLimitXMax(msScrollLimitXMax interface{}) {
	js.Rewrite("$_.msScrollLimitXMax = $1", msScrollLimitXMax)
}

// MsScrollLimitXMin prop
// js:"msScrollLimitXMin"
func (*CSSStyleDeclaration) MsScrollLimitXMin() (msScrollLimitXMin interface{}) {
	js.Rewrite("$_.msScrollLimitXMin")
	return msScrollLimitXMin
}

// SetMsScrollLimitXMin prop
// js:"msScrollLimitXMin"
func (*CSSStyleDeclaration) SetMsScrollLimitXMin(msScrollLimitXMin interface{}) {
	js.Rewrite("$_.msScrollLimitXMin = $1", msScrollLimitXMin)
}

// MsScrollLimitYMax prop
// js:"msScrollLimitYMax"
func (*CSSStyleDeclaration) MsScrollLimitYMax() (msScrollLimitYMax interface{}) {
	js.Rewrite("$_.msScrollLimitYMax")
	return msScrollLimitYMax
}

// SetMsScrollLimitYMax prop
// js:"msScrollLimitYMax"
func (*CSSStyleDeclaration) SetMsScrollLimitYMax(msScrollLimitYMax interface{}) {
	js.Rewrite("$_.msScrollLimitYMax = $1", msScrollLimitYMax)
}

// MsScrollLimitYMin prop
// js:"msScrollLimitYMin"
func (*CSSStyleDeclaration) MsScrollLimitYMin() (msScrollLimitYMin interface{}) {
	js.Rewrite("$_.msScrollLimitYMin")
	return msScrollLimitYMin
}

// SetMsScrollLimitYMin prop
// js:"msScrollLimitYMin"
func (*CSSStyleDeclaration) SetMsScrollLimitYMin(msScrollLimitYMin interface{}) {
	js.Rewrite("$_.msScrollLimitYMin = $1", msScrollLimitYMin)
}

// MsScrollRails prop
// js:"msScrollRails"
func (*CSSStyleDeclaration) MsScrollRails() (msScrollRails string) {
	js.Rewrite("$_.msScrollRails")
	return msScrollRails
}

// SetMsScrollRails prop
// js:"msScrollRails"
func (*CSSStyleDeclaration) SetMsScrollRails(msScrollRails string) {
	js.Rewrite("$_.msScrollRails = $1", msScrollRails)
}

// MsScrollSnapPointsX prop
// js:"msScrollSnapPointsX"
func (*CSSStyleDeclaration) MsScrollSnapPointsX() (msScrollSnapPointsX string) {
	js.Rewrite("$_.msScrollSnapPointsX")
	return msScrollSnapPointsX
}

// SetMsScrollSnapPointsX prop
// js:"msScrollSnapPointsX"
func (*CSSStyleDeclaration) SetMsScrollSnapPointsX(msScrollSnapPointsX string) {
	js.Rewrite("$_.msScrollSnapPointsX = $1", msScrollSnapPointsX)
}

// MsScrollSnapPointsY prop
// js:"msScrollSnapPointsY"
func (*CSSStyleDeclaration) MsScrollSnapPointsY() (msScrollSnapPointsY string) {
	js.Rewrite("$_.msScrollSnapPointsY")
	return msScrollSnapPointsY
}

// SetMsScrollSnapPointsY prop
// js:"msScrollSnapPointsY"
func (*CSSStyleDeclaration) SetMsScrollSnapPointsY(msScrollSnapPointsY string) {
	js.Rewrite("$_.msScrollSnapPointsY = $1", msScrollSnapPointsY)
}

// MsScrollSnapType prop
// js:"msScrollSnapType"
func (*CSSStyleDeclaration) MsScrollSnapType() (msScrollSnapType string) {
	js.Rewrite("$_.msScrollSnapType")
	return msScrollSnapType
}

// SetMsScrollSnapType prop
// js:"msScrollSnapType"
func (*CSSStyleDeclaration) SetMsScrollSnapType(msScrollSnapType string) {
	js.Rewrite("$_.msScrollSnapType = $1", msScrollSnapType)
}

// MsScrollSnapX prop
// js:"msScrollSnapX"
func (*CSSStyleDeclaration) MsScrollSnapX() (msScrollSnapX string) {
	js.Rewrite("$_.msScrollSnapX")
	return msScrollSnapX
}

// SetMsScrollSnapX prop
// js:"msScrollSnapX"
func (*CSSStyleDeclaration) SetMsScrollSnapX(msScrollSnapX string) {
	js.Rewrite("$_.msScrollSnapX = $1", msScrollSnapX)
}

// MsScrollSnapY prop
// js:"msScrollSnapY"
func (*CSSStyleDeclaration) MsScrollSnapY() (msScrollSnapY string) {
	js.Rewrite("$_.msScrollSnapY")
	return msScrollSnapY
}

// SetMsScrollSnapY prop
// js:"msScrollSnapY"
func (*CSSStyleDeclaration) SetMsScrollSnapY(msScrollSnapY string) {
	js.Rewrite("$_.msScrollSnapY = $1", msScrollSnapY)
}

// MsScrollTranslation prop
// js:"msScrollTranslation"
func (*CSSStyleDeclaration) MsScrollTranslation() (msScrollTranslation string) {
	js.Rewrite("$_.msScrollTranslation")
	return msScrollTranslation
}

// SetMsScrollTranslation prop
// js:"msScrollTranslation"
func (*CSSStyleDeclaration) SetMsScrollTranslation(msScrollTranslation string) {
	js.Rewrite("$_.msScrollTranslation = $1", msScrollTranslation)
}

// MsTextCombineHorizontal prop
// js:"msTextCombineHorizontal"
func (*CSSStyleDeclaration) MsTextCombineHorizontal() (msTextCombineHorizontal string) {
	js.Rewrite("$_.msTextCombineHorizontal")
	return msTextCombineHorizontal
}

// SetMsTextCombineHorizontal prop
// js:"msTextCombineHorizontal"
func (*CSSStyleDeclaration) SetMsTextCombineHorizontal(msTextCombineHorizontal string) {
	js.Rewrite("$_.msTextCombineHorizontal = $1", msTextCombineHorizontal)
}

// MsTextSizeAdjust prop
// js:"msTextSizeAdjust"
func (*CSSStyleDeclaration) MsTextSizeAdjust() (msTextSizeAdjust interface{}) {
	js.Rewrite("$_.msTextSizeAdjust")
	return msTextSizeAdjust
}

// SetMsTextSizeAdjust prop
// js:"msTextSizeAdjust"
func (*CSSStyleDeclaration) SetMsTextSizeAdjust(msTextSizeAdjust interface{}) {
	js.Rewrite("$_.msTextSizeAdjust = $1", msTextSizeAdjust)
}

// MsTouchAction prop
// js:"msTouchAction"
func (*CSSStyleDeclaration) MsTouchAction() (msTouchAction string) {
	js.Rewrite("$_.msTouchAction")
	return msTouchAction
}

// SetMsTouchAction prop
// js:"msTouchAction"
func (*CSSStyleDeclaration) SetMsTouchAction(msTouchAction string) {
	js.Rewrite("$_.msTouchAction = $1", msTouchAction)
}

// MsTouchSelect prop
// js:"msTouchSelect"
func (*CSSStyleDeclaration) MsTouchSelect() (msTouchSelect string) {
	js.Rewrite("$_.msTouchSelect")
	return msTouchSelect
}

// SetMsTouchSelect prop
// js:"msTouchSelect"
func (*CSSStyleDeclaration) SetMsTouchSelect(msTouchSelect string) {
	js.Rewrite("$_.msTouchSelect = $1", msTouchSelect)
}

// MsUserSelect prop
// js:"msUserSelect"
func (*CSSStyleDeclaration) MsUserSelect() (msUserSelect string) {
	js.Rewrite("$_.msUserSelect")
	return msUserSelect
}

// SetMsUserSelect prop
// js:"msUserSelect"
func (*CSSStyleDeclaration) SetMsUserSelect(msUserSelect string) {
	js.Rewrite("$_.msUserSelect = $1", msUserSelect)
}

// MsWrapFlow prop
// js:"msWrapFlow"
func (*CSSStyleDeclaration) MsWrapFlow() (msWrapFlow string) {
	js.Rewrite("$_.msWrapFlow")
	return msWrapFlow
}

// SetMsWrapFlow prop
// js:"msWrapFlow"
func (*CSSStyleDeclaration) SetMsWrapFlow(msWrapFlow string) {
	js.Rewrite("$_.msWrapFlow = $1", msWrapFlow)
}

// MsWrapMargin prop
// js:"msWrapMargin"
func (*CSSStyleDeclaration) MsWrapMargin() (msWrapMargin interface{}) {
	js.Rewrite("$_.msWrapMargin")
	return msWrapMargin
}

// SetMsWrapMargin prop
// js:"msWrapMargin"
func (*CSSStyleDeclaration) SetMsWrapMargin(msWrapMargin interface{}) {
	js.Rewrite("$_.msWrapMargin = $1", msWrapMargin)
}

// MsWrapThrough prop
// js:"msWrapThrough"
func (*CSSStyleDeclaration) MsWrapThrough() (msWrapThrough string) {
	js.Rewrite("$_.msWrapThrough")
	return msWrapThrough
}

// SetMsWrapThrough prop
// js:"msWrapThrough"
func (*CSSStyleDeclaration) SetMsWrapThrough(msWrapThrough string) {
	js.Rewrite("$_.msWrapThrough = $1", msWrapThrough)
}

// Opacity prop
// js:"opacity"
func (*CSSStyleDeclaration) Opacity() (opacity string) {
	js.Rewrite("$_.opacity")
	return opacity
}

// SetOpacity prop
// js:"opacity"
func (*CSSStyleDeclaration) SetOpacity(opacity string) {
	js.Rewrite("$_.opacity = $1", opacity)
}

// Order prop
// js:"order"
func (*CSSStyleDeclaration) Order() (order string) {
	js.Rewrite("$_.order")
	return order
}

// SetOrder prop
// js:"order"
func (*CSSStyleDeclaration) SetOrder(order string) {
	js.Rewrite("$_.order = $1", order)
}

// Orphans prop
// js:"orphans"
func (*CSSStyleDeclaration) Orphans() (orphans string) {
	js.Rewrite("$_.orphans")
	return orphans
}

// SetOrphans prop
// js:"orphans"
func (*CSSStyleDeclaration) SetOrphans(orphans string) {
	js.Rewrite("$_.orphans = $1", orphans)
}

// Outline prop
// js:"outline"
func (*CSSStyleDeclaration) Outline() (outline string) {
	js.Rewrite("$_.outline")
	return outline
}

// SetOutline prop
// js:"outline"
func (*CSSStyleDeclaration) SetOutline(outline string) {
	js.Rewrite("$_.outline = $1", outline)
}

// OutlineColor prop
// js:"outlineColor"
func (*CSSStyleDeclaration) OutlineColor() (outlineColor string) {
	js.Rewrite("$_.outlineColor")
	return outlineColor
}

// SetOutlineColor prop
// js:"outlineColor"
func (*CSSStyleDeclaration) SetOutlineColor(outlineColor string) {
	js.Rewrite("$_.outlineColor = $1", outlineColor)
}

// OutlineOffset prop
// js:"outlineOffset"
func (*CSSStyleDeclaration) OutlineOffset() (outlineOffset string) {
	js.Rewrite("$_.outlineOffset")
	return outlineOffset
}

// SetOutlineOffset prop
// js:"outlineOffset"
func (*CSSStyleDeclaration) SetOutlineOffset(outlineOffset string) {
	js.Rewrite("$_.outlineOffset = $1", outlineOffset)
}

// OutlineStyle prop
// js:"outlineStyle"
func (*CSSStyleDeclaration) OutlineStyle() (outlineStyle string) {
	js.Rewrite("$_.outlineStyle")
	return outlineStyle
}

// SetOutlineStyle prop
// js:"outlineStyle"
func (*CSSStyleDeclaration) SetOutlineStyle(outlineStyle string) {
	js.Rewrite("$_.outlineStyle = $1", outlineStyle)
}

// OutlineWidth prop
// js:"outlineWidth"
func (*CSSStyleDeclaration) OutlineWidth() (outlineWidth string) {
	js.Rewrite("$_.outlineWidth")
	return outlineWidth
}

// SetOutlineWidth prop
// js:"outlineWidth"
func (*CSSStyleDeclaration) SetOutlineWidth(outlineWidth string) {
	js.Rewrite("$_.outlineWidth = $1", outlineWidth)
}

// Overflow prop
// js:"overflow"
func (*CSSStyleDeclaration) Overflow() (overflow string) {
	js.Rewrite("$_.overflow")
	return overflow
}

// SetOverflow prop
// js:"overflow"
func (*CSSStyleDeclaration) SetOverflow(overflow string) {
	js.Rewrite("$_.overflow = $1", overflow)
}

// OverflowX prop
// js:"overflowX"
func (*CSSStyleDeclaration) OverflowX() (overflowX string) {
	js.Rewrite("$_.overflowX")
	return overflowX
}

// SetOverflowX prop
// js:"overflowX"
func (*CSSStyleDeclaration) SetOverflowX(overflowX string) {
	js.Rewrite("$_.overflowX = $1", overflowX)
}

// OverflowY prop
// js:"overflowY"
func (*CSSStyleDeclaration) OverflowY() (overflowY string) {
	js.Rewrite("$_.overflowY")
	return overflowY
}

// SetOverflowY prop
// js:"overflowY"
func (*CSSStyleDeclaration) SetOverflowY(overflowY string) {
	js.Rewrite("$_.overflowY = $1", overflowY)
}

// Padding prop
// js:"padding"
func (*CSSStyleDeclaration) Padding() (padding string) {
	js.Rewrite("$_.padding")
	return padding
}

// SetPadding prop
// js:"padding"
func (*CSSStyleDeclaration) SetPadding(padding string) {
	js.Rewrite("$_.padding = $1", padding)
}

// PaddingBottom prop
// js:"paddingBottom"
func (*CSSStyleDeclaration) PaddingBottom() (paddingBottom string) {
	js.Rewrite("$_.paddingBottom")
	return paddingBottom
}

// SetPaddingBottom prop
// js:"paddingBottom"
func (*CSSStyleDeclaration) SetPaddingBottom(paddingBottom string) {
	js.Rewrite("$_.paddingBottom = $1", paddingBottom)
}

// PaddingLeft prop
// js:"paddingLeft"
func (*CSSStyleDeclaration) PaddingLeft() (paddingLeft string) {
	js.Rewrite("$_.paddingLeft")
	return paddingLeft
}

// SetPaddingLeft prop
// js:"paddingLeft"
func (*CSSStyleDeclaration) SetPaddingLeft(paddingLeft string) {
	js.Rewrite("$_.paddingLeft = $1", paddingLeft)
}

// PaddingRight prop
// js:"paddingRight"
func (*CSSStyleDeclaration) PaddingRight() (paddingRight string) {
	js.Rewrite("$_.paddingRight")
	return paddingRight
}

// SetPaddingRight prop
// js:"paddingRight"
func (*CSSStyleDeclaration) SetPaddingRight(paddingRight string) {
	js.Rewrite("$_.paddingRight = $1", paddingRight)
}

// PaddingTop prop
// js:"paddingTop"
func (*CSSStyleDeclaration) PaddingTop() (paddingTop string) {
	js.Rewrite("$_.paddingTop")
	return paddingTop
}

// SetPaddingTop prop
// js:"paddingTop"
func (*CSSStyleDeclaration) SetPaddingTop(paddingTop string) {
	js.Rewrite("$_.paddingTop = $1", paddingTop)
}

// PageBreakAfter prop
// js:"pageBreakAfter"
func (*CSSStyleDeclaration) PageBreakAfter() (pageBreakAfter string) {
	js.Rewrite("$_.pageBreakAfter")
	return pageBreakAfter
}

// SetPageBreakAfter prop
// js:"pageBreakAfter"
func (*CSSStyleDeclaration) SetPageBreakAfter(pageBreakAfter string) {
	js.Rewrite("$_.pageBreakAfter = $1", pageBreakAfter)
}

// PageBreakBefore prop
// js:"pageBreakBefore"
func (*CSSStyleDeclaration) PageBreakBefore() (pageBreakBefore string) {
	js.Rewrite("$_.pageBreakBefore")
	return pageBreakBefore
}

// SetPageBreakBefore prop
// js:"pageBreakBefore"
func (*CSSStyleDeclaration) SetPageBreakBefore(pageBreakBefore string) {
	js.Rewrite("$_.pageBreakBefore = $1", pageBreakBefore)
}

// PageBreakInside prop
// js:"pageBreakInside"
func (*CSSStyleDeclaration) PageBreakInside() (pageBreakInside string) {
	js.Rewrite("$_.pageBreakInside")
	return pageBreakInside
}

// SetPageBreakInside prop
// js:"pageBreakInside"
func (*CSSStyleDeclaration) SetPageBreakInside(pageBreakInside string) {
	js.Rewrite("$_.pageBreakInside = $1", pageBreakInside)
}

// ParentRule prop
// js:"parentRule"
func (*CSSStyleDeclaration) ParentRule() (parentRule CSSRule) {
	js.Rewrite("$_.parentRule")
	return parentRule
}

// Perspective prop
// js:"perspective"
func (*CSSStyleDeclaration) Perspective() (perspective string) {
	js.Rewrite("$_.perspective")
	return perspective
}

// SetPerspective prop
// js:"perspective"
func (*CSSStyleDeclaration) SetPerspective(perspective string) {
	js.Rewrite("$_.perspective = $1", perspective)
}

// PerspectiveOrigin prop
// js:"perspectiveOrigin"
func (*CSSStyleDeclaration) PerspectiveOrigin() (perspectiveOrigin string) {
	js.Rewrite("$_.perspectiveOrigin")
	return perspectiveOrigin
}

// SetPerspectiveOrigin prop
// js:"perspectiveOrigin"
func (*CSSStyleDeclaration) SetPerspectiveOrigin(perspectiveOrigin string) {
	js.Rewrite("$_.perspectiveOrigin = $1", perspectiveOrigin)
}

// PointerEvents prop
// js:"pointerEvents"
func (*CSSStyleDeclaration) PointerEvents() (pointerEvents string) {
	js.Rewrite("$_.pointerEvents")
	return pointerEvents
}

// SetPointerEvents prop
// js:"pointerEvents"
func (*CSSStyleDeclaration) SetPointerEvents(pointerEvents string) {
	js.Rewrite("$_.pointerEvents = $1", pointerEvents)
}

// Position prop
// js:"position"
func (*CSSStyleDeclaration) Position() (position string) {
	js.Rewrite("$_.position")
	return position
}

// SetPosition prop
// js:"position"
func (*CSSStyleDeclaration) SetPosition(position string) {
	js.Rewrite("$_.position = $1", position)
}

// Quotes prop
// js:"quotes"
func (*CSSStyleDeclaration) Quotes() (quotes string) {
	js.Rewrite("$_.quotes")
	return quotes
}

// SetQuotes prop
// js:"quotes"
func (*CSSStyleDeclaration) SetQuotes(quotes string) {
	js.Rewrite("$_.quotes = $1", quotes)
}

// Right prop
// js:"right"
func (*CSSStyleDeclaration) Right() (right string) {
	js.Rewrite("$_.right")
	return right
}

// SetRight prop
// js:"right"
func (*CSSStyleDeclaration) SetRight(right string) {
	js.Rewrite("$_.right = $1", right)
}

// Rotate prop
// js:"rotate"
func (*CSSStyleDeclaration) Rotate() (rotate string) {
	js.Rewrite("$_.rotate")
	return rotate
}

// SetRotate prop
// js:"rotate"
func (*CSSStyleDeclaration) SetRotate(rotate string) {
	js.Rewrite("$_.rotate = $1", rotate)
}

// RubyAlign prop
// js:"rubyAlign"
func (*CSSStyleDeclaration) RubyAlign() (rubyAlign string) {
	js.Rewrite("$_.rubyAlign")
	return rubyAlign
}

// SetRubyAlign prop
// js:"rubyAlign"
func (*CSSStyleDeclaration) SetRubyAlign(rubyAlign string) {
	js.Rewrite("$_.rubyAlign = $1", rubyAlign)
}

// RubyOverhang prop
// js:"rubyOverhang"
func (*CSSStyleDeclaration) RubyOverhang() (rubyOverhang string) {
	js.Rewrite("$_.rubyOverhang")
	return rubyOverhang
}

// SetRubyOverhang prop
// js:"rubyOverhang"
func (*CSSStyleDeclaration) SetRubyOverhang(rubyOverhang string) {
	js.Rewrite("$_.rubyOverhang = $1", rubyOverhang)
}

// RubyPosition prop
// js:"rubyPosition"
func (*CSSStyleDeclaration) RubyPosition() (rubyPosition string) {
	js.Rewrite("$_.rubyPosition")
	return rubyPosition
}

// SetRubyPosition prop
// js:"rubyPosition"
func (*CSSStyleDeclaration) SetRubyPosition(rubyPosition string) {
	js.Rewrite("$_.rubyPosition = $1", rubyPosition)
}

// Scale prop
// js:"scale"
func (*CSSStyleDeclaration) Scale() (scale string) {
	js.Rewrite("$_.scale")
	return scale
}

// SetScale prop
// js:"scale"
func (*CSSStyleDeclaration) SetScale(scale string) {
	js.Rewrite("$_.scale = $1", scale)
}

// StopColor prop
// js:"stopColor"
func (*CSSStyleDeclaration) StopColor() (stopColor string) {
	js.Rewrite("$_.stopColor")
	return stopColor
}

// SetStopColor prop
// js:"stopColor"
func (*CSSStyleDeclaration) SetStopColor(stopColor string) {
	js.Rewrite("$_.stopColor = $1", stopColor)
}

// StopOpacity prop
// js:"stopOpacity"
func (*CSSStyleDeclaration) StopOpacity() (stopOpacity string) {
	js.Rewrite("$_.stopOpacity")
	return stopOpacity
}

// SetStopOpacity prop
// js:"stopOpacity"
func (*CSSStyleDeclaration) SetStopOpacity(stopOpacity string) {
	js.Rewrite("$_.stopOpacity = $1", stopOpacity)
}

// Stroke prop
// js:"stroke"
func (*CSSStyleDeclaration) Stroke() (stroke string) {
	js.Rewrite("$_.stroke")
	return stroke
}

// SetStroke prop
// js:"stroke"
func (*CSSStyleDeclaration) SetStroke(stroke string) {
	js.Rewrite("$_.stroke = $1", stroke)
}

// StrokeDasharray prop
// js:"strokeDasharray"
func (*CSSStyleDeclaration) StrokeDasharray() (strokeDasharray string) {
	js.Rewrite("$_.strokeDasharray")
	return strokeDasharray
}

// SetStrokeDasharray prop
// js:"strokeDasharray"
func (*CSSStyleDeclaration) SetStrokeDasharray(strokeDasharray string) {
	js.Rewrite("$_.strokeDasharray = $1", strokeDasharray)
}

// StrokeDashoffset prop
// js:"strokeDashoffset"
func (*CSSStyleDeclaration) StrokeDashoffset() (strokeDashoffset string) {
	js.Rewrite("$_.strokeDashoffset")
	return strokeDashoffset
}

// SetStrokeDashoffset prop
// js:"strokeDashoffset"
func (*CSSStyleDeclaration) SetStrokeDashoffset(strokeDashoffset string) {
	js.Rewrite("$_.strokeDashoffset = $1", strokeDashoffset)
}

// StrokeLinecap prop
// js:"strokeLinecap"
func (*CSSStyleDeclaration) StrokeLinecap() (strokeLinecap string) {
	js.Rewrite("$_.strokeLinecap")
	return strokeLinecap
}

// SetStrokeLinecap prop
// js:"strokeLinecap"
func (*CSSStyleDeclaration) SetStrokeLinecap(strokeLinecap string) {
	js.Rewrite("$_.strokeLinecap = $1", strokeLinecap)
}

// StrokeLinejoin prop
// js:"strokeLinejoin"
func (*CSSStyleDeclaration) StrokeLinejoin() (strokeLinejoin string) {
	js.Rewrite("$_.strokeLinejoin")
	return strokeLinejoin
}

// SetStrokeLinejoin prop
// js:"strokeLinejoin"
func (*CSSStyleDeclaration) SetStrokeLinejoin(strokeLinejoin string) {
	js.Rewrite("$_.strokeLinejoin = $1", strokeLinejoin)
}

// StrokeMiterlimit prop
// js:"strokeMiterlimit"
func (*CSSStyleDeclaration) StrokeMiterlimit() (strokeMiterlimit string) {
	js.Rewrite("$_.strokeMiterlimit")
	return strokeMiterlimit
}

// SetStrokeMiterlimit prop
// js:"strokeMiterlimit"
func (*CSSStyleDeclaration) SetStrokeMiterlimit(strokeMiterlimit string) {
	js.Rewrite("$_.strokeMiterlimit = $1", strokeMiterlimit)
}

// StrokeOpacity prop
// js:"strokeOpacity"
func (*CSSStyleDeclaration) StrokeOpacity() (strokeOpacity string) {
	js.Rewrite("$_.strokeOpacity")
	return strokeOpacity
}

// SetStrokeOpacity prop
// js:"strokeOpacity"
func (*CSSStyleDeclaration) SetStrokeOpacity(strokeOpacity string) {
	js.Rewrite("$_.strokeOpacity = $1", strokeOpacity)
}

// StrokeWidth prop
// js:"strokeWidth"
func (*CSSStyleDeclaration) StrokeWidth() (strokeWidth string) {
	js.Rewrite("$_.strokeWidth")
	return strokeWidth
}

// SetStrokeWidth prop
// js:"strokeWidth"
func (*CSSStyleDeclaration) SetStrokeWidth(strokeWidth string) {
	js.Rewrite("$_.strokeWidth = $1", strokeWidth)
}

// TableLayout prop
// js:"tableLayout"
func (*CSSStyleDeclaration) TableLayout() (tableLayout string) {
	js.Rewrite("$_.tableLayout")
	return tableLayout
}

// SetTableLayout prop
// js:"tableLayout"
func (*CSSStyleDeclaration) SetTableLayout(tableLayout string) {
	js.Rewrite("$_.tableLayout = $1", tableLayout)
}

// TextAlign prop
// js:"textAlign"
func (*CSSStyleDeclaration) TextAlign() (textAlign string) {
	js.Rewrite("$_.textAlign")
	return textAlign
}

// SetTextAlign prop
// js:"textAlign"
func (*CSSStyleDeclaration) SetTextAlign(textAlign string) {
	js.Rewrite("$_.textAlign = $1", textAlign)
}

// TextAlignLast prop
// js:"textAlignLast"
func (*CSSStyleDeclaration) TextAlignLast() (textAlignLast string) {
	js.Rewrite("$_.textAlignLast")
	return textAlignLast
}

// SetTextAlignLast prop
// js:"textAlignLast"
func (*CSSStyleDeclaration) SetTextAlignLast(textAlignLast string) {
	js.Rewrite("$_.textAlignLast = $1", textAlignLast)
}

// TextAnchor prop
// js:"textAnchor"
func (*CSSStyleDeclaration) TextAnchor() (textAnchor string) {
	js.Rewrite("$_.textAnchor")
	return textAnchor
}

// SetTextAnchor prop
// js:"textAnchor"
func (*CSSStyleDeclaration) SetTextAnchor(textAnchor string) {
	js.Rewrite("$_.textAnchor = $1", textAnchor)
}

// TextDecoration prop
// js:"textDecoration"
func (*CSSStyleDeclaration) TextDecoration() (textDecoration string) {
	js.Rewrite("$_.textDecoration")
	return textDecoration
}

// SetTextDecoration prop
// js:"textDecoration"
func (*CSSStyleDeclaration) SetTextDecoration(textDecoration string) {
	js.Rewrite("$_.textDecoration = $1", textDecoration)
}

// TextIndent prop
// js:"textIndent"
func (*CSSStyleDeclaration) TextIndent() (textIndent string) {
	js.Rewrite("$_.textIndent")
	return textIndent
}

// SetTextIndent prop
// js:"textIndent"
func (*CSSStyleDeclaration) SetTextIndent(textIndent string) {
	js.Rewrite("$_.textIndent = $1", textIndent)
}

// TextJustify prop
// js:"textJustify"
func (*CSSStyleDeclaration) TextJustify() (textJustify string) {
	js.Rewrite("$_.textJustify")
	return textJustify
}

// SetTextJustify prop
// js:"textJustify"
func (*CSSStyleDeclaration) SetTextJustify(textJustify string) {
	js.Rewrite("$_.textJustify = $1", textJustify)
}

// TextKashida prop
// js:"textKashida"
func (*CSSStyleDeclaration) TextKashida() (textKashida string) {
	js.Rewrite("$_.textKashida")
	return textKashida
}

// SetTextKashida prop
// js:"textKashida"
func (*CSSStyleDeclaration) SetTextKashida(textKashida string) {
	js.Rewrite("$_.textKashida = $1", textKashida)
}

// TextKashidaSpace prop
// js:"textKashidaSpace"
func (*CSSStyleDeclaration) TextKashidaSpace() (textKashidaSpace string) {
	js.Rewrite("$_.textKashidaSpace")
	return textKashidaSpace
}

// SetTextKashidaSpace prop
// js:"textKashidaSpace"
func (*CSSStyleDeclaration) SetTextKashidaSpace(textKashidaSpace string) {
	js.Rewrite("$_.textKashidaSpace = $1", textKashidaSpace)
}

// TextOverflow prop
// js:"textOverflow"
func (*CSSStyleDeclaration) TextOverflow() (textOverflow string) {
	js.Rewrite("$_.textOverflow")
	return textOverflow
}

// SetTextOverflow prop
// js:"textOverflow"
func (*CSSStyleDeclaration) SetTextOverflow(textOverflow string) {
	js.Rewrite("$_.textOverflow = $1", textOverflow)
}

// TextShadow prop
// js:"textShadow"
func (*CSSStyleDeclaration) TextShadow() (textShadow string) {
	js.Rewrite("$_.textShadow")
	return textShadow
}

// SetTextShadow prop
// js:"textShadow"
func (*CSSStyleDeclaration) SetTextShadow(textShadow string) {
	js.Rewrite("$_.textShadow = $1", textShadow)
}

// TextTransform prop
// js:"textTransform"
func (*CSSStyleDeclaration) TextTransform() (textTransform string) {
	js.Rewrite("$_.textTransform")
	return textTransform
}

// SetTextTransform prop
// js:"textTransform"
func (*CSSStyleDeclaration) SetTextTransform(textTransform string) {
	js.Rewrite("$_.textTransform = $1", textTransform)
}

// TextUnderlinePosition prop
// js:"textUnderlinePosition"
func (*CSSStyleDeclaration) TextUnderlinePosition() (textUnderlinePosition string) {
	js.Rewrite("$_.textUnderlinePosition")
	return textUnderlinePosition
}

// SetTextUnderlinePosition prop
// js:"textUnderlinePosition"
func (*CSSStyleDeclaration) SetTextUnderlinePosition(textUnderlinePosition string) {
	js.Rewrite("$_.textUnderlinePosition = $1", textUnderlinePosition)
}

// Top prop
// js:"top"
func (*CSSStyleDeclaration) Top() (top string) {
	js.Rewrite("$_.top")
	return top
}

// SetTop prop
// js:"top"
func (*CSSStyleDeclaration) SetTop(top string) {
	js.Rewrite("$_.top = $1", top)
}

// TouchAction prop
// js:"touchAction"
func (*CSSStyleDeclaration) TouchAction() (touchAction string) {
	js.Rewrite("$_.touchAction")
	return touchAction
}

// SetTouchAction prop
// js:"touchAction"
func (*CSSStyleDeclaration) SetTouchAction(touchAction string) {
	js.Rewrite("$_.touchAction = $1", touchAction)
}

// Transform prop
// js:"transform"
func (*CSSStyleDeclaration) Transform() (transform string) {
	js.Rewrite("$_.transform")
	return transform
}

// SetTransform prop
// js:"transform"
func (*CSSStyleDeclaration) SetTransform(transform string) {
	js.Rewrite("$_.transform = $1", transform)
}

// TransformOrigin prop
// js:"transformOrigin"
func (*CSSStyleDeclaration) TransformOrigin() (transformOrigin string) {
	js.Rewrite("$_.transformOrigin")
	return transformOrigin
}

// SetTransformOrigin prop
// js:"transformOrigin"
func (*CSSStyleDeclaration) SetTransformOrigin(transformOrigin string) {
	js.Rewrite("$_.transformOrigin = $1", transformOrigin)
}

// TransformStyle prop
// js:"transformStyle"
func (*CSSStyleDeclaration) TransformStyle() (transformStyle string) {
	js.Rewrite("$_.transformStyle")
	return transformStyle
}

// SetTransformStyle prop
// js:"transformStyle"
func (*CSSStyleDeclaration) SetTransformStyle(transformStyle string) {
	js.Rewrite("$_.transformStyle = $1", transformStyle)
}

// Transition prop
// js:"transition"
func (*CSSStyleDeclaration) Transition() (transition string) {
	js.Rewrite("$_.transition")
	return transition
}

// SetTransition prop
// js:"transition"
func (*CSSStyleDeclaration) SetTransition(transition string) {
	js.Rewrite("$_.transition = $1", transition)
}

// TransitionDelay prop
// js:"transitionDelay"
func (*CSSStyleDeclaration) TransitionDelay() (transitionDelay string) {
	js.Rewrite("$_.transitionDelay")
	return transitionDelay
}

// SetTransitionDelay prop
// js:"transitionDelay"
func (*CSSStyleDeclaration) SetTransitionDelay(transitionDelay string) {
	js.Rewrite("$_.transitionDelay = $1", transitionDelay)
}

// TransitionDuration prop
// js:"transitionDuration"
func (*CSSStyleDeclaration) TransitionDuration() (transitionDuration string) {
	js.Rewrite("$_.transitionDuration")
	return transitionDuration
}

// SetTransitionDuration prop
// js:"transitionDuration"
func (*CSSStyleDeclaration) SetTransitionDuration(transitionDuration string) {
	js.Rewrite("$_.transitionDuration = $1", transitionDuration)
}

// TransitionProperty prop
// js:"transitionProperty"
func (*CSSStyleDeclaration) TransitionProperty() (transitionProperty string) {
	js.Rewrite("$_.transitionProperty")
	return transitionProperty
}

// SetTransitionProperty prop
// js:"transitionProperty"
func (*CSSStyleDeclaration) SetTransitionProperty(transitionProperty string) {
	js.Rewrite("$_.transitionProperty = $1", transitionProperty)
}

// TransitionTimingFunction prop
// js:"transitionTimingFunction"
func (*CSSStyleDeclaration) TransitionTimingFunction() (transitionTimingFunction string) {
	js.Rewrite("$_.transitionTimingFunction")
	return transitionTimingFunction
}

// SetTransitionTimingFunction prop
// js:"transitionTimingFunction"
func (*CSSStyleDeclaration) SetTransitionTimingFunction(transitionTimingFunction string) {
	js.Rewrite("$_.transitionTimingFunction = $1", transitionTimingFunction)
}

// Translate prop
// js:"translate"
func (*CSSStyleDeclaration) Translate() (translate string) {
	js.Rewrite("$_.translate")
	return translate
}

// SetTranslate prop
// js:"translate"
func (*CSSStyleDeclaration) SetTranslate(translate string) {
	js.Rewrite("$_.translate = $1", translate)
}

// UnicodeBidi prop
// js:"unicodeBidi"
func (*CSSStyleDeclaration) UnicodeBidi() (unicodeBidi string) {
	js.Rewrite("$_.unicodeBidi")
	return unicodeBidi
}

// SetUnicodeBidi prop
// js:"unicodeBidi"
func (*CSSStyleDeclaration) SetUnicodeBidi(unicodeBidi string) {
	js.Rewrite("$_.unicodeBidi = $1", unicodeBidi)
}

// VerticalAlign prop
// js:"verticalAlign"
func (*CSSStyleDeclaration) VerticalAlign() (verticalAlign string) {
	js.Rewrite("$_.verticalAlign")
	return verticalAlign
}

// SetVerticalAlign prop
// js:"verticalAlign"
func (*CSSStyleDeclaration) SetVerticalAlign(verticalAlign string) {
	js.Rewrite("$_.verticalAlign = $1", verticalAlign)
}

// Visibility prop
// js:"visibility"
func (*CSSStyleDeclaration) Visibility() (visibility string) {
	js.Rewrite("$_.visibility")
	return visibility
}

// SetVisibility prop
// js:"visibility"
func (*CSSStyleDeclaration) SetVisibility(visibility string) {
	js.Rewrite("$_.visibility = $1", visibility)
}

// WebkitAlignContent prop
// js:"webkitAlignContent"
func (*CSSStyleDeclaration) WebkitAlignContent() (webkitAlignContent string) {
	js.Rewrite("$_.webkitAlignContent")
	return webkitAlignContent
}

// SetWebkitAlignContent prop
// js:"webkitAlignContent"
func (*CSSStyleDeclaration) SetWebkitAlignContent(webkitAlignContent string) {
	js.Rewrite("$_.webkitAlignContent = $1", webkitAlignContent)
}

// WebkitAlignItems prop
// js:"webkitAlignItems"
func (*CSSStyleDeclaration) WebkitAlignItems() (webkitAlignItems string) {
	js.Rewrite("$_.webkitAlignItems")
	return webkitAlignItems
}

// SetWebkitAlignItems prop
// js:"webkitAlignItems"
func (*CSSStyleDeclaration) SetWebkitAlignItems(webkitAlignItems string) {
	js.Rewrite("$_.webkitAlignItems = $1", webkitAlignItems)
}

// WebkitAlignSelf prop
// js:"webkitAlignSelf"
func (*CSSStyleDeclaration) WebkitAlignSelf() (webkitAlignSelf string) {
	js.Rewrite("$_.webkitAlignSelf")
	return webkitAlignSelf
}

// SetWebkitAlignSelf prop
// js:"webkitAlignSelf"
func (*CSSStyleDeclaration) SetWebkitAlignSelf(webkitAlignSelf string) {
	js.Rewrite("$_.webkitAlignSelf = $1", webkitAlignSelf)
}

// WebkitAnimation prop
// js:"webkitAnimation"
func (*CSSStyleDeclaration) WebkitAnimation() (webkitAnimation string) {
	js.Rewrite("$_.webkitAnimation")
	return webkitAnimation
}

// SetWebkitAnimation prop
// js:"webkitAnimation"
func (*CSSStyleDeclaration) SetWebkitAnimation(webkitAnimation string) {
	js.Rewrite("$_.webkitAnimation = $1", webkitAnimation)
}

// WebkitAnimationDelay prop
// js:"webkitAnimationDelay"
func (*CSSStyleDeclaration) WebkitAnimationDelay() (webkitAnimationDelay string) {
	js.Rewrite("$_.webkitAnimationDelay")
	return webkitAnimationDelay
}

// SetWebkitAnimationDelay prop
// js:"webkitAnimationDelay"
func (*CSSStyleDeclaration) SetWebkitAnimationDelay(webkitAnimationDelay string) {
	js.Rewrite("$_.webkitAnimationDelay = $1", webkitAnimationDelay)
}

// WebkitAnimationDirection prop
// js:"webkitAnimationDirection"
func (*CSSStyleDeclaration) WebkitAnimationDirection() (webkitAnimationDirection string) {
	js.Rewrite("$_.webkitAnimationDirection")
	return webkitAnimationDirection
}

// SetWebkitAnimationDirection prop
// js:"webkitAnimationDirection"
func (*CSSStyleDeclaration) SetWebkitAnimationDirection(webkitAnimationDirection string) {
	js.Rewrite("$_.webkitAnimationDirection = $1", webkitAnimationDirection)
}

// WebkitAnimationDuration prop
// js:"webkitAnimationDuration"
func (*CSSStyleDeclaration) WebkitAnimationDuration() (webkitAnimationDuration string) {
	js.Rewrite("$_.webkitAnimationDuration")
	return webkitAnimationDuration
}

// SetWebkitAnimationDuration prop
// js:"webkitAnimationDuration"
func (*CSSStyleDeclaration) SetWebkitAnimationDuration(webkitAnimationDuration string) {
	js.Rewrite("$_.webkitAnimationDuration = $1", webkitAnimationDuration)
}

// WebkitAnimationFillMode prop
// js:"webkitAnimationFillMode"
func (*CSSStyleDeclaration) WebkitAnimationFillMode() (webkitAnimationFillMode string) {
	js.Rewrite("$_.webkitAnimationFillMode")
	return webkitAnimationFillMode
}

// SetWebkitAnimationFillMode prop
// js:"webkitAnimationFillMode"
func (*CSSStyleDeclaration) SetWebkitAnimationFillMode(webkitAnimationFillMode string) {
	js.Rewrite("$_.webkitAnimationFillMode = $1", webkitAnimationFillMode)
}

// WebkitAnimationIterationCount prop
// js:"webkitAnimationIterationCount"
func (*CSSStyleDeclaration) WebkitAnimationIterationCount() (webkitAnimationIterationCount string) {
	js.Rewrite("$_.webkitAnimationIterationCount")
	return webkitAnimationIterationCount
}

// SetWebkitAnimationIterationCount prop
// js:"webkitAnimationIterationCount"
func (*CSSStyleDeclaration) SetWebkitAnimationIterationCount(webkitAnimationIterationCount string) {
	js.Rewrite("$_.webkitAnimationIterationCount = $1", webkitAnimationIterationCount)
}

// WebkitAnimationName prop
// js:"webkitAnimationName"
func (*CSSStyleDeclaration) WebkitAnimationName() (webkitAnimationName string) {
	js.Rewrite("$_.webkitAnimationName")
	return webkitAnimationName
}

// SetWebkitAnimationName prop
// js:"webkitAnimationName"
func (*CSSStyleDeclaration) SetWebkitAnimationName(webkitAnimationName string) {
	js.Rewrite("$_.webkitAnimationName = $1", webkitAnimationName)
}

// WebkitAnimationPlayState prop
// js:"webkitAnimationPlayState"
func (*CSSStyleDeclaration) WebkitAnimationPlayState() (webkitAnimationPlayState string) {
	js.Rewrite("$_.webkitAnimationPlayState")
	return webkitAnimationPlayState
}

// SetWebkitAnimationPlayState prop
// js:"webkitAnimationPlayState"
func (*CSSStyleDeclaration) SetWebkitAnimationPlayState(webkitAnimationPlayState string) {
	js.Rewrite("$_.webkitAnimationPlayState = $1", webkitAnimationPlayState)
}

// WebkitAnimationTimingFunction prop
// js:"webkitAnimationTimingFunction"
func (*CSSStyleDeclaration) WebkitAnimationTimingFunction() (webkitAnimationTimingFunction string) {
	js.Rewrite("$_.webkitAnimationTimingFunction")
	return webkitAnimationTimingFunction
}

// SetWebkitAnimationTimingFunction prop
// js:"webkitAnimationTimingFunction"
func (*CSSStyleDeclaration) SetWebkitAnimationTimingFunction(webkitAnimationTimingFunction string) {
	js.Rewrite("$_.webkitAnimationTimingFunction = $1", webkitAnimationTimingFunction)
}

// WebkitAppearance prop
// js:"webkitAppearance"
func (*CSSStyleDeclaration) WebkitAppearance() (webkitAppearance string) {
	js.Rewrite("$_.webkitAppearance")
	return webkitAppearance
}

// SetWebkitAppearance prop
// js:"webkitAppearance"
func (*CSSStyleDeclaration) SetWebkitAppearance(webkitAppearance string) {
	js.Rewrite("$_.webkitAppearance = $1", webkitAppearance)
}

// WebkitBackfaceVisibility prop
// js:"webkitBackfaceVisibility"
func (*CSSStyleDeclaration) WebkitBackfaceVisibility() (webkitBackfaceVisibility string) {
	js.Rewrite("$_.webkitBackfaceVisibility")
	return webkitBackfaceVisibility
}

// SetWebkitBackfaceVisibility prop
// js:"webkitBackfaceVisibility"
func (*CSSStyleDeclaration) SetWebkitBackfaceVisibility(webkitBackfaceVisibility string) {
	js.Rewrite("$_.webkitBackfaceVisibility = $1", webkitBackfaceVisibility)
}

// WebkitBackgroundClip prop
// js:"webkitBackgroundClip"
func (*CSSStyleDeclaration) WebkitBackgroundClip() (webkitBackgroundClip string) {
	js.Rewrite("$_.webkitBackgroundClip")
	return webkitBackgroundClip
}

// SetWebkitBackgroundClip prop
// js:"webkitBackgroundClip"
func (*CSSStyleDeclaration) SetWebkitBackgroundClip(webkitBackgroundClip string) {
	js.Rewrite("$_.webkitBackgroundClip = $1", webkitBackgroundClip)
}

// WebkitBackgroundOrigin prop
// js:"webkitBackgroundOrigin"
func (*CSSStyleDeclaration) WebkitBackgroundOrigin() (webkitBackgroundOrigin string) {
	js.Rewrite("$_.webkitBackgroundOrigin")
	return webkitBackgroundOrigin
}

// SetWebkitBackgroundOrigin prop
// js:"webkitBackgroundOrigin"
func (*CSSStyleDeclaration) SetWebkitBackgroundOrigin(webkitBackgroundOrigin string) {
	js.Rewrite("$_.webkitBackgroundOrigin = $1", webkitBackgroundOrigin)
}

// WebkitBackgroundSize prop
// js:"webkitBackgroundSize"
func (*CSSStyleDeclaration) WebkitBackgroundSize() (webkitBackgroundSize string) {
	js.Rewrite("$_.webkitBackgroundSize")
	return webkitBackgroundSize
}

// SetWebkitBackgroundSize prop
// js:"webkitBackgroundSize"
func (*CSSStyleDeclaration) SetWebkitBackgroundSize(webkitBackgroundSize string) {
	js.Rewrite("$_.webkitBackgroundSize = $1", webkitBackgroundSize)
}

// WebkitBorderBottomLeftRadius prop
// js:"webkitBorderBottomLeftRadius"
func (*CSSStyleDeclaration) WebkitBorderBottomLeftRadius() (webkitBorderBottomLeftRadius string) {
	js.Rewrite("$_.webkitBorderBottomLeftRadius")
	return webkitBorderBottomLeftRadius
}

// SetWebkitBorderBottomLeftRadius prop
// js:"webkitBorderBottomLeftRadius"
func (*CSSStyleDeclaration) SetWebkitBorderBottomLeftRadius(webkitBorderBottomLeftRadius string) {
	js.Rewrite("$_.webkitBorderBottomLeftRadius = $1", webkitBorderBottomLeftRadius)
}

// WebkitBorderBottomRightRadius prop
// js:"webkitBorderBottomRightRadius"
func (*CSSStyleDeclaration) WebkitBorderBottomRightRadius() (webkitBorderBottomRightRadius string) {
	js.Rewrite("$_.webkitBorderBottomRightRadius")
	return webkitBorderBottomRightRadius
}

// SetWebkitBorderBottomRightRadius prop
// js:"webkitBorderBottomRightRadius"
func (*CSSStyleDeclaration) SetWebkitBorderBottomRightRadius(webkitBorderBottomRightRadius string) {
	js.Rewrite("$_.webkitBorderBottomRightRadius = $1", webkitBorderBottomRightRadius)
}

// WebkitBorderImage prop
// js:"webkitBorderImage"
func (*CSSStyleDeclaration) WebkitBorderImage() (webkitBorderImage string) {
	js.Rewrite("$_.webkitBorderImage")
	return webkitBorderImage
}

// SetWebkitBorderImage prop
// js:"webkitBorderImage"
func (*CSSStyleDeclaration) SetWebkitBorderImage(webkitBorderImage string) {
	js.Rewrite("$_.webkitBorderImage = $1", webkitBorderImage)
}

// WebkitBorderRadius prop
// js:"webkitBorderRadius"
func (*CSSStyleDeclaration) WebkitBorderRadius() (webkitBorderRadius string) {
	js.Rewrite("$_.webkitBorderRadius")
	return webkitBorderRadius
}

// SetWebkitBorderRadius prop
// js:"webkitBorderRadius"
func (*CSSStyleDeclaration) SetWebkitBorderRadius(webkitBorderRadius string) {
	js.Rewrite("$_.webkitBorderRadius = $1", webkitBorderRadius)
}

// WebkitBorderTopLeftRadius prop
// js:"webkitBorderTopLeftRadius"
func (*CSSStyleDeclaration) WebkitBorderTopLeftRadius() (webkitBorderTopLeftRadius string) {
	js.Rewrite("$_.webkitBorderTopLeftRadius")
	return webkitBorderTopLeftRadius
}

// SetWebkitBorderTopLeftRadius prop
// js:"webkitBorderTopLeftRadius"
func (*CSSStyleDeclaration) SetWebkitBorderTopLeftRadius(webkitBorderTopLeftRadius string) {
	js.Rewrite("$_.webkitBorderTopLeftRadius = $1", webkitBorderTopLeftRadius)
}

// WebkitBorderTopRightRadius prop
// js:"webkitBorderTopRightRadius"
func (*CSSStyleDeclaration) WebkitBorderTopRightRadius() (webkitBorderTopRightRadius string) {
	js.Rewrite("$_.webkitBorderTopRightRadius")
	return webkitBorderTopRightRadius
}

// SetWebkitBorderTopRightRadius prop
// js:"webkitBorderTopRightRadius"
func (*CSSStyleDeclaration) SetWebkitBorderTopRightRadius(webkitBorderTopRightRadius string) {
	js.Rewrite("$_.webkitBorderTopRightRadius = $1", webkitBorderTopRightRadius)
}

// WebkitBoxAlign prop
// js:"webkitBoxAlign"
func (*CSSStyleDeclaration) WebkitBoxAlign() (webkitBoxAlign string) {
	js.Rewrite("$_.webkitBoxAlign")
	return webkitBoxAlign
}

// SetWebkitBoxAlign prop
// js:"webkitBoxAlign"
func (*CSSStyleDeclaration) SetWebkitBoxAlign(webkitBoxAlign string) {
	js.Rewrite("$_.webkitBoxAlign = $1", webkitBoxAlign)
}

// WebkitBoxDirection prop
// js:"webkitBoxDirection"
func (*CSSStyleDeclaration) WebkitBoxDirection() (webkitBoxDirection string) {
	js.Rewrite("$_.webkitBoxDirection")
	return webkitBoxDirection
}

// SetWebkitBoxDirection prop
// js:"webkitBoxDirection"
func (*CSSStyleDeclaration) SetWebkitBoxDirection(webkitBoxDirection string) {
	js.Rewrite("$_.webkitBoxDirection = $1", webkitBoxDirection)
}

// WebkitBoxFlex prop
// js:"webkitBoxFlex"
func (*CSSStyleDeclaration) WebkitBoxFlex() (webkitBoxFlex string) {
	js.Rewrite("$_.webkitBoxFlex")
	return webkitBoxFlex
}

// SetWebkitBoxFlex prop
// js:"webkitBoxFlex"
func (*CSSStyleDeclaration) SetWebkitBoxFlex(webkitBoxFlex string) {
	js.Rewrite("$_.webkitBoxFlex = $1", webkitBoxFlex)
}

// WebkitBoxOrdinalGroup prop
// js:"webkitBoxOrdinalGroup"
func (*CSSStyleDeclaration) WebkitBoxOrdinalGroup() (webkitBoxOrdinalGroup string) {
	js.Rewrite("$_.webkitBoxOrdinalGroup")
	return webkitBoxOrdinalGroup
}

// SetWebkitBoxOrdinalGroup prop
// js:"webkitBoxOrdinalGroup"
func (*CSSStyleDeclaration) SetWebkitBoxOrdinalGroup(webkitBoxOrdinalGroup string) {
	js.Rewrite("$_.webkitBoxOrdinalGroup = $1", webkitBoxOrdinalGroup)
}

// WebkitBoxOrient prop
// js:"webkitBoxOrient"
func (*CSSStyleDeclaration) WebkitBoxOrient() (webkitBoxOrient string) {
	js.Rewrite("$_.webkitBoxOrient")
	return webkitBoxOrient
}

// SetWebkitBoxOrient prop
// js:"webkitBoxOrient"
func (*CSSStyleDeclaration) SetWebkitBoxOrient(webkitBoxOrient string) {
	js.Rewrite("$_.webkitBoxOrient = $1", webkitBoxOrient)
}

// WebkitBoxPack prop
// js:"webkitBoxPack"
func (*CSSStyleDeclaration) WebkitBoxPack() (webkitBoxPack string) {
	js.Rewrite("$_.webkitBoxPack")
	return webkitBoxPack
}

// SetWebkitBoxPack prop
// js:"webkitBoxPack"
func (*CSSStyleDeclaration) SetWebkitBoxPack(webkitBoxPack string) {
	js.Rewrite("$_.webkitBoxPack = $1", webkitBoxPack)
}

// WebkitBoxSizing prop
// js:"webkitBoxSizing"
func (*CSSStyleDeclaration) WebkitBoxSizing() (webkitBoxSizing string) {
	js.Rewrite("$_.webkitBoxSizing")
	return webkitBoxSizing
}

// SetWebkitBoxSizing prop
// js:"webkitBoxSizing"
func (*CSSStyleDeclaration) SetWebkitBoxSizing(webkitBoxSizing string) {
	js.Rewrite("$_.webkitBoxSizing = $1", webkitBoxSizing)
}

// WebkitColumnBreakAfter prop
// js:"webkitColumnBreakAfter"
func (*CSSStyleDeclaration) WebkitColumnBreakAfter() (webkitColumnBreakAfter string) {
	js.Rewrite("$_.webkitColumnBreakAfter")
	return webkitColumnBreakAfter
}

// SetWebkitColumnBreakAfter prop
// js:"webkitColumnBreakAfter"
func (*CSSStyleDeclaration) SetWebkitColumnBreakAfter(webkitColumnBreakAfter string) {
	js.Rewrite("$_.webkitColumnBreakAfter = $1", webkitColumnBreakAfter)
}

// WebkitColumnBreakBefore prop
// js:"webkitColumnBreakBefore"
func (*CSSStyleDeclaration) WebkitColumnBreakBefore() (webkitColumnBreakBefore string) {
	js.Rewrite("$_.webkitColumnBreakBefore")
	return webkitColumnBreakBefore
}

// SetWebkitColumnBreakBefore prop
// js:"webkitColumnBreakBefore"
func (*CSSStyleDeclaration) SetWebkitColumnBreakBefore(webkitColumnBreakBefore string) {
	js.Rewrite("$_.webkitColumnBreakBefore = $1", webkitColumnBreakBefore)
}

// WebkitColumnBreakInside prop
// js:"webkitColumnBreakInside"
func (*CSSStyleDeclaration) WebkitColumnBreakInside() (webkitColumnBreakInside string) {
	js.Rewrite("$_.webkitColumnBreakInside")
	return webkitColumnBreakInside
}

// SetWebkitColumnBreakInside prop
// js:"webkitColumnBreakInside"
func (*CSSStyleDeclaration) SetWebkitColumnBreakInside(webkitColumnBreakInside string) {
	js.Rewrite("$_.webkitColumnBreakInside = $1", webkitColumnBreakInside)
}

// WebkitColumnCount prop
// js:"webkitColumnCount"
func (*CSSStyleDeclaration) WebkitColumnCount() (webkitColumnCount interface{}) {
	js.Rewrite("$_.webkitColumnCount")
	return webkitColumnCount
}

// SetWebkitColumnCount prop
// js:"webkitColumnCount"
func (*CSSStyleDeclaration) SetWebkitColumnCount(webkitColumnCount interface{}) {
	js.Rewrite("$_.webkitColumnCount = $1", webkitColumnCount)
}

// WebkitColumnGap prop
// js:"webkitColumnGap"
func (*CSSStyleDeclaration) WebkitColumnGap() (webkitColumnGap interface{}) {
	js.Rewrite("$_.webkitColumnGap")
	return webkitColumnGap
}

// SetWebkitColumnGap prop
// js:"webkitColumnGap"
func (*CSSStyleDeclaration) SetWebkitColumnGap(webkitColumnGap interface{}) {
	js.Rewrite("$_.webkitColumnGap = $1", webkitColumnGap)
}

// WebkitColumnRule prop
// js:"webkitColumnRule"
func (*CSSStyleDeclaration) WebkitColumnRule() (webkitColumnRule string) {
	js.Rewrite("$_.webkitColumnRule")
	return webkitColumnRule
}

// SetWebkitColumnRule prop
// js:"webkitColumnRule"
func (*CSSStyleDeclaration) SetWebkitColumnRule(webkitColumnRule string) {
	js.Rewrite("$_.webkitColumnRule = $1", webkitColumnRule)
}

// WebkitColumnRuleColor prop
// js:"webkitColumnRuleColor"
func (*CSSStyleDeclaration) WebkitColumnRuleColor() (webkitColumnRuleColor interface{}) {
	js.Rewrite("$_.webkitColumnRuleColor")
	return webkitColumnRuleColor
}

// SetWebkitColumnRuleColor prop
// js:"webkitColumnRuleColor"
func (*CSSStyleDeclaration) SetWebkitColumnRuleColor(webkitColumnRuleColor interface{}) {
	js.Rewrite("$_.webkitColumnRuleColor = $1", webkitColumnRuleColor)
}

// WebkitColumnRuleStyle prop
// js:"webkitColumnRuleStyle"
func (*CSSStyleDeclaration) WebkitColumnRuleStyle() (webkitColumnRuleStyle string) {
	js.Rewrite("$_.webkitColumnRuleStyle")
	return webkitColumnRuleStyle
}

// SetWebkitColumnRuleStyle prop
// js:"webkitColumnRuleStyle"
func (*CSSStyleDeclaration) SetWebkitColumnRuleStyle(webkitColumnRuleStyle string) {
	js.Rewrite("$_.webkitColumnRuleStyle = $1", webkitColumnRuleStyle)
}

// WebkitColumnRuleWidth prop
// js:"webkitColumnRuleWidth"
func (*CSSStyleDeclaration) WebkitColumnRuleWidth() (webkitColumnRuleWidth interface{}) {
	js.Rewrite("$_.webkitColumnRuleWidth")
	return webkitColumnRuleWidth
}

// SetWebkitColumnRuleWidth prop
// js:"webkitColumnRuleWidth"
func (*CSSStyleDeclaration) SetWebkitColumnRuleWidth(webkitColumnRuleWidth interface{}) {
	js.Rewrite("$_.webkitColumnRuleWidth = $1", webkitColumnRuleWidth)
}

// WebkitColumns prop
// js:"webkitColumns"
func (*CSSStyleDeclaration) WebkitColumns() (webkitColumns string) {
	js.Rewrite("$_.webkitColumns")
	return webkitColumns
}

// SetWebkitColumns prop
// js:"webkitColumns"
func (*CSSStyleDeclaration) SetWebkitColumns(webkitColumns string) {
	js.Rewrite("$_.webkitColumns = $1", webkitColumns)
}

// WebkitColumnSpan prop
// js:"webkitColumnSpan"
func (*CSSStyleDeclaration) WebkitColumnSpan() (webkitColumnSpan string) {
	js.Rewrite("$_.webkitColumnSpan")
	return webkitColumnSpan
}

// SetWebkitColumnSpan prop
// js:"webkitColumnSpan"
func (*CSSStyleDeclaration) SetWebkitColumnSpan(webkitColumnSpan string) {
	js.Rewrite("$_.webkitColumnSpan = $1", webkitColumnSpan)
}

// WebkitColumnWidth prop
// js:"webkitColumnWidth"
func (*CSSStyleDeclaration) WebkitColumnWidth() (webkitColumnWidth interface{}) {
	js.Rewrite("$_.webkitColumnWidth")
	return webkitColumnWidth
}

// SetWebkitColumnWidth prop
// js:"webkitColumnWidth"
func (*CSSStyleDeclaration) SetWebkitColumnWidth(webkitColumnWidth interface{}) {
	js.Rewrite("$_.webkitColumnWidth = $1", webkitColumnWidth)
}

// WebkitFilter prop
// js:"webkitFilter"
func (*CSSStyleDeclaration) WebkitFilter() (webkitFilter string) {
	js.Rewrite("$_.webkitFilter")
	return webkitFilter
}

// SetWebkitFilter prop
// js:"webkitFilter"
func (*CSSStyleDeclaration) SetWebkitFilter(webkitFilter string) {
	js.Rewrite("$_.webkitFilter = $1", webkitFilter)
}

// WebkitFlex prop
// js:"webkitFlex"
func (*CSSStyleDeclaration) WebkitFlex() (webkitFlex string) {
	js.Rewrite("$_.webkitFlex")
	return webkitFlex
}

// SetWebkitFlex prop
// js:"webkitFlex"
func (*CSSStyleDeclaration) SetWebkitFlex(webkitFlex string) {
	js.Rewrite("$_.webkitFlex = $1", webkitFlex)
}

// WebkitFlexBasis prop
// js:"webkitFlexBasis"
func (*CSSStyleDeclaration) WebkitFlexBasis() (webkitFlexBasis string) {
	js.Rewrite("$_.webkitFlexBasis")
	return webkitFlexBasis
}

// SetWebkitFlexBasis prop
// js:"webkitFlexBasis"
func (*CSSStyleDeclaration) SetWebkitFlexBasis(webkitFlexBasis string) {
	js.Rewrite("$_.webkitFlexBasis = $1", webkitFlexBasis)
}

// WebkitFlexDirection prop
// js:"webkitFlexDirection"
func (*CSSStyleDeclaration) WebkitFlexDirection() (webkitFlexDirection string) {
	js.Rewrite("$_.webkitFlexDirection")
	return webkitFlexDirection
}

// SetWebkitFlexDirection prop
// js:"webkitFlexDirection"
func (*CSSStyleDeclaration) SetWebkitFlexDirection(webkitFlexDirection string) {
	js.Rewrite("$_.webkitFlexDirection = $1", webkitFlexDirection)
}

// WebkitFlexFlow prop
// js:"webkitFlexFlow"
func (*CSSStyleDeclaration) WebkitFlexFlow() (webkitFlexFlow string) {
	js.Rewrite("$_.webkitFlexFlow")
	return webkitFlexFlow
}

// SetWebkitFlexFlow prop
// js:"webkitFlexFlow"
func (*CSSStyleDeclaration) SetWebkitFlexFlow(webkitFlexFlow string) {
	js.Rewrite("$_.webkitFlexFlow = $1", webkitFlexFlow)
}

// WebkitFlexGrow prop
// js:"webkitFlexGrow"
func (*CSSStyleDeclaration) WebkitFlexGrow() (webkitFlexGrow string) {
	js.Rewrite("$_.webkitFlexGrow")
	return webkitFlexGrow
}

// SetWebkitFlexGrow prop
// js:"webkitFlexGrow"
func (*CSSStyleDeclaration) SetWebkitFlexGrow(webkitFlexGrow string) {
	js.Rewrite("$_.webkitFlexGrow = $1", webkitFlexGrow)
}

// WebkitFlexShrink prop
// js:"webkitFlexShrink"
func (*CSSStyleDeclaration) WebkitFlexShrink() (webkitFlexShrink string) {
	js.Rewrite("$_.webkitFlexShrink")
	return webkitFlexShrink
}

// SetWebkitFlexShrink prop
// js:"webkitFlexShrink"
func (*CSSStyleDeclaration) SetWebkitFlexShrink(webkitFlexShrink string) {
	js.Rewrite("$_.webkitFlexShrink = $1", webkitFlexShrink)
}

// WebkitFlexWrap prop
// js:"webkitFlexWrap"
func (*CSSStyleDeclaration) WebkitFlexWrap() (webkitFlexWrap string) {
	js.Rewrite("$_.webkitFlexWrap")
	return webkitFlexWrap
}

// SetWebkitFlexWrap prop
// js:"webkitFlexWrap"
func (*CSSStyleDeclaration) SetWebkitFlexWrap(webkitFlexWrap string) {
	js.Rewrite("$_.webkitFlexWrap = $1", webkitFlexWrap)
}

// WebkitJustifyContent prop
// js:"webkitJustifyContent"
func (*CSSStyleDeclaration) WebkitJustifyContent() (webkitJustifyContent string) {
	js.Rewrite("$_.webkitJustifyContent")
	return webkitJustifyContent
}

// SetWebkitJustifyContent prop
// js:"webkitJustifyContent"
func (*CSSStyleDeclaration) SetWebkitJustifyContent(webkitJustifyContent string) {
	js.Rewrite("$_.webkitJustifyContent = $1", webkitJustifyContent)
}

// WebkitOrder prop
// js:"webkitOrder"
func (*CSSStyleDeclaration) WebkitOrder() (webkitOrder string) {
	js.Rewrite("$_.webkitOrder")
	return webkitOrder
}

// SetWebkitOrder prop
// js:"webkitOrder"
func (*CSSStyleDeclaration) SetWebkitOrder(webkitOrder string) {
	js.Rewrite("$_.webkitOrder = $1", webkitOrder)
}

// WebkitPerspective prop
// js:"webkitPerspective"
func (*CSSStyleDeclaration) WebkitPerspective() (webkitPerspective string) {
	js.Rewrite("$_.webkitPerspective")
	return webkitPerspective
}

// SetWebkitPerspective prop
// js:"webkitPerspective"
func (*CSSStyleDeclaration) SetWebkitPerspective(webkitPerspective string) {
	js.Rewrite("$_.webkitPerspective = $1", webkitPerspective)
}

// WebkitPerspectiveOrigin prop
// js:"webkitPerspectiveOrigin"
func (*CSSStyleDeclaration) WebkitPerspectiveOrigin() (webkitPerspectiveOrigin string) {
	js.Rewrite("$_.webkitPerspectiveOrigin")
	return webkitPerspectiveOrigin
}

// SetWebkitPerspectiveOrigin prop
// js:"webkitPerspectiveOrigin"
func (*CSSStyleDeclaration) SetWebkitPerspectiveOrigin(webkitPerspectiveOrigin string) {
	js.Rewrite("$_.webkitPerspectiveOrigin = $1", webkitPerspectiveOrigin)
}

// WebkitTapHighlightColor prop
// js:"webkitTapHighlightColor"
func (*CSSStyleDeclaration) WebkitTapHighlightColor() (webkitTapHighlightColor string) {
	js.Rewrite("$_.webkitTapHighlightColor")
	return webkitTapHighlightColor
}

// SetWebkitTapHighlightColor prop
// js:"webkitTapHighlightColor"
func (*CSSStyleDeclaration) SetWebkitTapHighlightColor(webkitTapHighlightColor string) {
	js.Rewrite("$_.webkitTapHighlightColor = $1", webkitTapHighlightColor)
}

// WebkitTextFillColor prop
// js:"webkitTextFillColor"
func (*CSSStyleDeclaration) WebkitTextFillColor() (webkitTextFillColor string) {
	js.Rewrite("$_.webkitTextFillColor")
	return webkitTextFillColor
}

// SetWebkitTextFillColor prop
// js:"webkitTextFillColor"
func (*CSSStyleDeclaration) SetWebkitTextFillColor(webkitTextFillColor string) {
	js.Rewrite("$_.webkitTextFillColor = $1", webkitTextFillColor)
}

// WebkitTextSizeAdjust prop
// js:"webkitTextSizeAdjust"
func (*CSSStyleDeclaration) WebkitTextSizeAdjust() (webkitTextSizeAdjust interface{}) {
	js.Rewrite("$_.webkitTextSizeAdjust")
	return webkitTextSizeAdjust
}

// SetWebkitTextSizeAdjust prop
// js:"webkitTextSizeAdjust"
func (*CSSStyleDeclaration) SetWebkitTextSizeAdjust(webkitTextSizeAdjust interface{}) {
	js.Rewrite("$_.webkitTextSizeAdjust = $1", webkitTextSizeAdjust)
}

// WebkitTextStroke prop
// js:"webkitTextStroke"
func (*CSSStyleDeclaration) WebkitTextStroke() (webkitTextStroke string) {
	js.Rewrite("$_.webkitTextStroke")
	return webkitTextStroke
}

// SetWebkitTextStroke prop
// js:"webkitTextStroke"
func (*CSSStyleDeclaration) SetWebkitTextStroke(webkitTextStroke string) {
	js.Rewrite("$_.webkitTextStroke = $1", webkitTextStroke)
}

// WebkitTextStrokeColor prop
// js:"webkitTextStrokeColor"
func (*CSSStyleDeclaration) WebkitTextStrokeColor() (webkitTextStrokeColor string) {
	js.Rewrite("$_.webkitTextStrokeColor")
	return webkitTextStrokeColor
}

// SetWebkitTextStrokeColor prop
// js:"webkitTextStrokeColor"
func (*CSSStyleDeclaration) SetWebkitTextStrokeColor(webkitTextStrokeColor string) {
	js.Rewrite("$_.webkitTextStrokeColor = $1", webkitTextStrokeColor)
}

// WebkitTextStrokeWidth prop
// js:"webkitTextStrokeWidth"
func (*CSSStyleDeclaration) WebkitTextStrokeWidth() (webkitTextStrokeWidth string) {
	js.Rewrite("$_.webkitTextStrokeWidth")
	return webkitTextStrokeWidth
}

// SetWebkitTextStrokeWidth prop
// js:"webkitTextStrokeWidth"
func (*CSSStyleDeclaration) SetWebkitTextStrokeWidth(webkitTextStrokeWidth string) {
	js.Rewrite("$_.webkitTextStrokeWidth = $1", webkitTextStrokeWidth)
}

// WebkitTransform prop
// js:"webkitTransform"
func (*CSSStyleDeclaration) WebkitTransform() (webkitTransform string) {
	js.Rewrite("$_.webkitTransform")
	return webkitTransform
}

// SetWebkitTransform prop
// js:"webkitTransform"
func (*CSSStyleDeclaration) SetWebkitTransform(webkitTransform string) {
	js.Rewrite("$_.webkitTransform = $1", webkitTransform)
}

// WebkitTransformOrigin prop
// js:"webkitTransformOrigin"
func (*CSSStyleDeclaration) WebkitTransformOrigin() (webkitTransformOrigin string) {
	js.Rewrite("$_.webkitTransformOrigin")
	return webkitTransformOrigin
}

// SetWebkitTransformOrigin prop
// js:"webkitTransformOrigin"
func (*CSSStyleDeclaration) SetWebkitTransformOrigin(webkitTransformOrigin string) {
	js.Rewrite("$_.webkitTransformOrigin = $1", webkitTransformOrigin)
}

// WebkitTransformStyle prop
// js:"webkitTransformStyle"
func (*CSSStyleDeclaration) WebkitTransformStyle() (webkitTransformStyle string) {
	js.Rewrite("$_.webkitTransformStyle")
	return webkitTransformStyle
}

// SetWebkitTransformStyle prop
// js:"webkitTransformStyle"
func (*CSSStyleDeclaration) SetWebkitTransformStyle(webkitTransformStyle string) {
	js.Rewrite("$_.webkitTransformStyle = $1", webkitTransformStyle)
}

// WebkitTransition prop
// js:"webkitTransition"
func (*CSSStyleDeclaration) WebkitTransition() (webkitTransition string) {
	js.Rewrite("$_.webkitTransition")
	return webkitTransition
}

// SetWebkitTransition prop
// js:"webkitTransition"
func (*CSSStyleDeclaration) SetWebkitTransition(webkitTransition string) {
	js.Rewrite("$_.webkitTransition = $1", webkitTransition)
}

// WebkitTransitionDelay prop
// js:"webkitTransitionDelay"
func (*CSSStyleDeclaration) WebkitTransitionDelay() (webkitTransitionDelay string) {
	js.Rewrite("$_.webkitTransitionDelay")
	return webkitTransitionDelay
}

// SetWebkitTransitionDelay prop
// js:"webkitTransitionDelay"
func (*CSSStyleDeclaration) SetWebkitTransitionDelay(webkitTransitionDelay string) {
	js.Rewrite("$_.webkitTransitionDelay = $1", webkitTransitionDelay)
}

// WebkitTransitionDuration prop
// js:"webkitTransitionDuration"
func (*CSSStyleDeclaration) WebkitTransitionDuration() (webkitTransitionDuration string) {
	js.Rewrite("$_.webkitTransitionDuration")
	return webkitTransitionDuration
}

// SetWebkitTransitionDuration prop
// js:"webkitTransitionDuration"
func (*CSSStyleDeclaration) SetWebkitTransitionDuration(webkitTransitionDuration string) {
	js.Rewrite("$_.webkitTransitionDuration = $1", webkitTransitionDuration)
}

// WebkitTransitionProperty prop
// js:"webkitTransitionProperty"
func (*CSSStyleDeclaration) WebkitTransitionProperty() (webkitTransitionProperty string) {
	js.Rewrite("$_.webkitTransitionProperty")
	return webkitTransitionProperty
}

// SetWebkitTransitionProperty prop
// js:"webkitTransitionProperty"
func (*CSSStyleDeclaration) SetWebkitTransitionProperty(webkitTransitionProperty string) {
	js.Rewrite("$_.webkitTransitionProperty = $1", webkitTransitionProperty)
}

// WebkitTransitionTimingFunction prop
// js:"webkitTransitionTimingFunction"
func (*CSSStyleDeclaration) WebkitTransitionTimingFunction() (webkitTransitionTimingFunction string) {
	js.Rewrite("$_.webkitTransitionTimingFunction")
	return webkitTransitionTimingFunction
}

// SetWebkitTransitionTimingFunction prop
// js:"webkitTransitionTimingFunction"
func (*CSSStyleDeclaration) SetWebkitTransitionTimingFunction(webkitTransitionTimingFunction string) {
	js.Rewrite("$_.webkitTransitionTimingFunction = $1", webkitTransitionTimingFunction)
}

// WebkitUserModify prop
// js:"webkitUserModify"
func (*CSSStyleDeclaration) WebkitUserModify() (webkitUserModify string) {
	js.Rewrite("$_.webkitUserModify")
	return webkitUserModify
}

// SetWebkitUserModify prop
// js:"webkitUserModify"
func (*CSSStyleDeclaration) SetWebkitUserModify(webkitUserModify string) {
	js.Rewrite("$_.webkitUserModify = $1", webkitUserModify)
}

// WebkitUserSelect prop
// js:"webkitUserSelect"
func (*CSSStyleDeclaration) WebkitUserSelect() (webkitUserSelect string) {
	js.Rewrite("$_.webkitUserSelect")
	return webkitUserSelect
}

// SetWebkitUserSelect prop
// js:"webkitUserSelect"
func (*CSSStyleDeclaration) SetWebkitUserSelect(webkitUserSelect string) {
	js.Rewrite("$_.webkitUserSelect = $1", webkitUserSelect)
}

// WebkitWritingMode prop
// js:"webkitWritingMode"
func (*CSSStyleDeclaration) WebkitWritingMode() (webkitWritingMode string) {
	js.Rewrite("$_.webkitWritingMode")
	return webkitWritingMode
}

// SetWebkitWritingMode prop
// js:"webkitWritingMode"
func (*CSSStyleDeclaration) SetWebkitWritingMode(webkitWritingMode string) {
	js.Rewrite("$_.webkitWritingMode = $1", webkitWritingMode)
}

// WhiteSpace prop
// js:"whiteSpace"
func (*CSSStyleDeclaration) WhiteSpace() (whiteSpace string) {
	js.Rewrite("$_.whiteSpace")
	return whiteSpace
}

// SetWhiteSpace prop
// js:"whiteSpace"
func (*CSSStyleDeclaration) SetWhiteSpace(whiteSpace string) {
	js.Rewrite("$_.whiteSpace = $1", whiteSpace)
}

// Widows prop
// js:"widows"
func (*CSSStyleDeclaration) Widows() (widows string) {
	js.Rewrite("$_.widows")
	return widows
}

// SetWidows prop
// js:"widows"
func (*CSSStyleDeclaration) SetWidows(widows string) {
	js.Rewrite("$_.widows = $1", widows)
}

// Width prop
// js:"width"
func (*CSSStyleDeclaration) Width() (width string) {
	js.Rewrite("$_.width")
	return width
}

// SetWidth prop
// js:"width"
func (*CSSStyleDeclaration) SetWidth(width string) {
	js.Rewrite("$_.width = $1", width)
}

// WordBreak prop
// js:"wordBreak"
func (*CSSStyleDeclaration) WordBreak() (wordBreak string) {
	js.Rewrite("$_.wordBreak")
	return wordBreak
}

// SetWordBreak prop
// js:"wordBreak"
func (*CSSStyleDeclaration) SetWordBreak(wordBreak string) {
	js.Rewrite("$_.wordBreak = $1", wordBreak)
}

// WordSpacing prop
// js:"wordSpacing"
func (*CSSStyleDeclaration) WordSpacing() (wordSpacing string) {
	js.Rewrite("$_.wordSpacing")
	return wordSpacing
}

// SetWordSpacing prop
// js:"wordSpacing"
func (*CSSStyleDeclaration) SetWordSpacing(wordSpacing string) {
	js.Rewrite("$_.wordSpacing = $1", wordSpacing)
}

// WordWrap prop
// js:"wordWrap"
func (*CSSStyleDeclaration) WordWrap() (wordWrap string) {
	js.Rewrite("$_.wordWrap")
	return wordWrap
}

// SetWordWrap prop
// js:"wordWrap"
func (*CSSStyleDeclaration) SetWordWrap(wordWrap string) {
	js.Rewrite("$_.wordWrap = $1", wordWrap)
}

// WritingMode prop
// js:"writingMode"
func (*CSSStyleDeclaration) WritingMode() (writingMode string) {
	js.Rewrite("$_.writingMode")
	return writingMode
}

// SetWritingMode prop
// js:"writingMode"
func (*CSSStyleDeclaration) SetWritingMode(writingMode string) {
	js.Rewrite("$_.writingMode = $1", writingMode)
}

// ZIndex prop
// js:"zIndex"
func (*CSSStyleDeclaration) ZIndex() (zIndex string) {
	js.Rewrite("$_.zIndex")
	return zIndex
}

// SetZIndex prop
// js:"zIndex"
func (*CSSStyleDeclaration) SetZIndex(zIndex string) {
	js.Rewrite("$_.zIndex = $1", zIndex)
}

// Zoom prop
// js:"zoom"
func (*CSSStyleDeclaration) Zoom() (zoom string) {
	js.Rewrite("$_.zoom")
	return zoom
}

// SetZoom prop
// js:"zoom"
func (*CSSStyleDeclaration) SetZoom(zoom string) {
	js.Rewrite("$_.zoom = $1", zoom)
}
