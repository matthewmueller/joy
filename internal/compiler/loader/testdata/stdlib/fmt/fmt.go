package fmt

// Println fn
func Println(a ...interface{}) (n int, err error) {
	println(a)
	return 0, nil
}
