package main

func main() {
	doc := New()
	el := doc.GetElementByID("canvas")
	println("got element " + el)
}

// New document
func New() Document {
	return Document{
		Version: 1,
	}
}

// Document struct
type Document struct {
	Version int `js:"version"`
}

// GetElementByID lookup element by it's id
func (*Document) GetElementByID(id string) string {
	return "#" + id
}
