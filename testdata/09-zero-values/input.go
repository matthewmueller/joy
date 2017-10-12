package main

// User struct
type User struct {
	FirstName string
	LastName  string
	Age       int
	Phone     Phone
}

// Phone struct
type Phone struct {
	Type   string
	Number string
}

func main() {
	user := User{FirstName: "Tobi"}
	println(user.FirstName)
	println(user.LastName)
	println(user.Age)
	println(user.Phone.Type)
	println(user.Phone.Number)
}
