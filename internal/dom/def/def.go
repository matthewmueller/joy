package def

// Definition interface
type Definition interface {
	ID() string
	Name() string
	Kind() string
	Dependencies() ([]Definition, error)
	Generate() (string, error)
}
