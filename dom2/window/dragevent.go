package window

import (
	"github.com/matthewmueller/golly/dom2/file"
	"github.com/matthewmueller/golly/js"
)

// js:"DragEvent,omit"
type DragEvent struct {
	MouseEvent
}

// InitDragEvent
func (*DragEvent) InitDragEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int, screenXArg int, screenYArg int, clientXArg int, clientYArg int, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg uint8, relatedTargetArg EventTarget, dataTransferArg *datatransfer.DataTransfer) {
	js.Rewrite("$<.InitDragEvent($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, screenXArg, screenYArg, clientXArg, clientYArg, ctrlKeyArg, altKeyArg, shiftKeyArg, metaKeyArg, buttonArg, relatedTargetArg, dataTransferArg)
}

// MsConvertURL
func (*DragEvent) MsConvertURL(file *file.File, targetType string, targetURL *string) {
	js.Rewrite("$<.MsConvertURL($1, $2, $3)", file, targetType, targetURL)
}

// DataTransfer
func (*DragEvent) DataTransfer() (dataTransfer *datatransfer.DataTransfer) {
	js.Rewrite("$<.DataTransfer")
	return dataTransfer
}
