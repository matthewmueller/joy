package main

// User struct
type User struct {
	Name string `js:"name"`
	Age  int
}

func main() {
	age := 28
	u := User{"matt", age}
	println(u.Name, u.Age)
}
