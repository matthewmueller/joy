package file

// File struct
type File struct {
	name   string
	source string
}

// New file
func New(name, source string) *File {
	return &File{
		name:   name,
		source: source,
	}
}

// Name of file
func (f *File) Name() string {
	return f.name
}

// Source of file
func (f *File) Source() string {
	return f.source
}
