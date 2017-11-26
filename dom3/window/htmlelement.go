package window

type HTMLElement interface {
	Element

	QuerySelector(selector string) Element
}

