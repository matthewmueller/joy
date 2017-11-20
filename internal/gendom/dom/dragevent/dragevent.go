package dragevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/datatransfer"
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/file"
	"github.com/matthewmueller/golly/internal/gendom/dom/mouseevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/window"
	"github.com/matthewmueller/golly/js"
)

type DragEvent struct {
	*mouseevent.MouseEvent
}

func (*DragEvent) InitDragEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *window.Window, detailArg int, screenXArg int, screenYArg int, clientXArg int, clientYArg int, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg uint8, relatedTargetArg *eventtarget.EventTarget, dataTransferArg *datatransfer.DataTransfer) {
	js.Rewrite("$<.initDragEvent($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, screenXArg, screenYArg, clientXArg, clientYArg, ctrlKeyArg, altKeyArg, shiftKeyArg, metaKeyArg, buttonArg, relatedTargetArg, dataTransferArg)
}

func (*DragEvent) MsConvertURL(file *file.File, targetType string, targetURL string) {
	js.Rewrite("$<.msConvertURL($1, $2, $3)", file, targetType, targetURL)
}

func (*DragEvent) GetDataTransfer() (dataTransfer *datatransfer.DataTransfer) {
	js.Rewrite("$<.dataTransfer")
	return dataTransfer
}
