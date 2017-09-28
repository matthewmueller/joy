package main

func main() {
	a := "a"
	b, c, d := "b", "c", "d"

	var e = 3

	f, g := func() (string, string) {
		return "f", "g"
	}()

	var (
		h    = "h"
		i    = "i"
		j, k = "j", "k"
	)

	println(a, b, c, d, e, f, g, h, i, j, k)
}
