package def

// Definition interface
type Definition interface {
	ID() string
	Name() string
	Kind() string
	Type() (string, error)
	Dependencies() ([]Definition, error)
	Generate() (string, error)

	// This comes after the definion is
	// already made is sort of hacky
	// TODO: come up with a better solution
	SetPackage(pkg string)
	GetPackage() string

	SetFile(file string)
	GetFile() string
}
