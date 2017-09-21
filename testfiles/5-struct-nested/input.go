package main

// User struct
type User struct {
	Name  string
	Phone Phone
}

// Phone struct
type Phone struct {
	Number string
	Main   bool
}

func main() {
	user := User{
		Name: "Matt",
		Phone: Phone{
			Number: "1234511",
			Main:   false,
		},
	}

	println(user.Phone.Number)
}
