package msfilesaver

// MSFileSaver interface
// js:"MSFileSaver"
type MSFileSaver interface {

	// MsSaveBlob
	// js:"msSaveBlob"
	// jsrewrite:"$_.msSaveBlob($1, $2)"
	MsSaveBlob(blob interface{}, defaultName *string) (b bool)

	// MsSaveOrOpenBlob
	// js:"msSaveOrOpenBlob"
	// jsrewrite:"$_.msSaveOrOpenBlob($1, $2)"
	MsSaveOrOpenBlob(blob interface{}, defaultName *string) (b bool)
}
