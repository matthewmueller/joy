package script

// Script struct
type Script struct {
	name   string
	path   string
	source string
}

// New script
func New(name, path, source string) *Script {
	return &Script{name, path, source}
}

// Name fn
func (s *Script) Name() string {
	return s.name
}

// Path fn
func (s *Script) Path() string {
	return s.path
}

// Source fn
func (s *Script) Source() string {
	return s.source
}
