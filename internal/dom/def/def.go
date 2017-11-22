package def

// Definition interface
type Definition interface {
	ID() string
	Name() string
	Kind() string
	Generate() (string, error)
	Children() ([]Definition, error)
}
