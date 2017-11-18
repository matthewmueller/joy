package generator

type index struct {
	CallbackFunctions  map[string]*CallbackFunction
	CallbackInterfaces map[string]*CallbackInterface
	Dictionaries       map[string]*Dictionary
	Enums              map[string]*Enum
	Interfaces         map[string]*Interface
	MixinInterfaces    map[string]*Interface
	TypeDefs           map[string]*Typedef
}

func newIndex(d *dom) *index {
	idx := &index{
		CallbackFunctions:  map[string]*CallbackFunction{},
		CallbackInterfaces: map[string]*CallbackInterface{},
		Dictionaries:       map[string]*Dictionary{},
		Enums:              map[string]*Enum{},
		Interfaces:         map[string]*Interface{},
		MixinInterfaces:    map[string]*Interface{},
		TypeDefs:           map[string]*Typedef{},
	}
	for _, node := range d.CallbackFunctions {
		idx.CallbackFunctions[node.Name] = node
	}
	for _, node := range d.CallbackInterfaces {
		idx.CallbackInterfaces[node.Name] = node
	}
	for _, node := range d.Dictionaries {
		idx.Dictionaries[node.Name] = node
	}
	for _, node := range d.Enums {
		idx.Enums[node.Name] = node
	}
	for _, node := range d.Interfaces {
		idx.Interfaces[node.Name] = node
	}
	for _, node := range d.MixinInterfaces {
		idx.MixinInterfaces[node.Name] = node
	}
	for _, node := range d.TypeDefs {
		idx.TypeDefs[node.NewType] = node
	}

	idx = adjust(idx)
	return idx
}

// manual adjustments
func adjust(idx *index) *index {
	var iface *Interface

	// Remove some Pointer event properties that
	// conflict with methods of the same name
	// and aren't in the spec
	iface = idx.Interfaces["MSPointerEvent"]
	for i, prop := range iface.Properties {
		if prop.Name == "currentPoint" {
			iface.Properties[i] = nil
		}
		if prop.Name == "intermediatePoints" {
			iface.Properties[i] = nil
		}
	}
	iface = idx.Interfaces["PointerEvent"]
	for i, prop := range iface.Properties {
		if prop.Name == "currentPoint" {
			iface.Properties[i] = nil
		}
		if prop.Name == "intermediatePoints" {
			iface.Properties[i] = nil
		}
	}

	return idx
}
