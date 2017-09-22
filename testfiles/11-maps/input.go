package main

func main() {
	names := map[string]string{
		"a": "a",
		"b": "b",
	}

	nested := map[string]map[string]string{
		"a": map[string]string{
			"a": "a",
		},
		"b": map[string]string{
			"b": "b",
		},
	}

	println(names["a"])
	println(nested["b"]["b"])
}
