package main

import "fmt"

func main() {
	x, y := f()
	fmt.Println(x)
	fmt.Println(y)

	// tambah terus
	xs := []int{1, 2, 3, 4}
	fmt.Println(add(xs...))

	fmt.Println(add(1, 2, 3))
}

func f() (int, int) {
	return 5, 6
}

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}
