package msfilesaver

// MSFileSaver interface
// js:"MSFileSaver"
type MSFileSaver interface {

	// MsSaveBlob
	// js:"msSaveBlob"
	MsSaveBlob(blob interface{}, defaultName *string) (b bool)

	// MsSaveOrOpenBlob
	// js:"msSaveOrOpenBlob"
	MsSaveOrOpenBlob(blob interface{}, defaultName *string) (b bool)
}
