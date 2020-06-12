package main

import "fmt"

func main() {
	nextValue := tambah()

	fmt.Println(nextValue())

}

func tambah(value int) func() int {
	x := 10
	return func() int {
		x += 1
		return x
	}
}
