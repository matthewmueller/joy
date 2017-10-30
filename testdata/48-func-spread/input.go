package main

func main() {
	items := test("a", "b", "c", "d")
	for _, item := range items {
		println(item)
	}

	a := aa{}
	items = a.test2("a", "b", "c", "d")
	for _, item := range items {
		println(item)
	}
}

func test(a string, items ...string) []string {
	return items
}

type aa struct{}

func (*aa) test2(a string, items ...string) []string {
	return items
}
