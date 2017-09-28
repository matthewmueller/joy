package main

func main() {
	for i := 0; i < 5; i++ {
		println(i)
	}

	j := 5
	k := 0
	for {
		j--
		if j < 0 {
			break
		} else {
			k++
		}
	}
	println(k)
}
