package main

func main() {
	name := "anki"
	b := true
	if name == "matt" || name == "mark" {
		println("matt or mark")
	} else if name == "anki" && b {
		println("anki")
	} else if b {
		println("truthy")
	} else {
		println("something else")
	}
}
