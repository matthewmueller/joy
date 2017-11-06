package main

type thing struct {
}

func (t *thing) Type() int {
	return 1
}

type entity struct{}

func (t *entity) Entity() string {
	return "ent"
}

type animal struct {
	kind string
	*thing
	entity
}

func (a *animal) Name() string {
	return "animal"
}

type dog struct {
	*animal
}

func (d *dog) Legs() int {
	return 4
}

func main() {
	d := dog{
		animal: &animal{
			kind: "sloth",
		},
	}
	println(d.Name(), d.Legs(), d.kind, d.Type(), d.Entity())
}
