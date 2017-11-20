package def

// Definition interface
type Definition interface {
	ID() string
	Name() string
	Kind() string
	// Parents() []Definition
	// Ancestors() []Definition
	Children() ([]Definition, error)
}
