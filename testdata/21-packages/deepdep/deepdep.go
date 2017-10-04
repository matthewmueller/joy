package deepdep

// Dep struct
type Dep struct {
	deep string
}

// New fn
func New() *Dep {
	return &Dep{
		deep: "deep",
	}
}

func (d *Dep) String() string {
	return d.deep
}
