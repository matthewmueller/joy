package msfilesaver

// js:"MSFileSaver,omit"
type MSFileSaver interface {

	// MsSaveBlob
	MsSaveBlob(blob interface{}, defaultName *string) (b bool)

	// MsSaveOrOpenBlob
	MsSaveOrOpenBlob(blob interface{}, defaultName *string) (b bool)
}
