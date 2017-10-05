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
	println("%j", user.FirstName)
	println("%j", user.LastName)
	println("%j", user.Age)
	println("%j", user.Phone)
	println("%j", user.Phone.Number)
}
