package main

// User struct
type User struct {
	ID string
}

func main() {
	u := User{"a"}
	// names := map[User]string{}
	// names[u] = "a"
	println(u.ID)
}
