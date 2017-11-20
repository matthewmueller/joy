package focusnavigationeventinit

type FocusNavigationEventInit struct {
	*EventInit

	navigationReason *string
	originHeight     *float32
	originLeft       *float32
	originTop        *float32
	originWidth      *float32
}
