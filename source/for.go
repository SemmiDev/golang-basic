package main

import "fmt"

func main() {
	// di permrograman go perulangan hanya ada for
	// cara pertama
	// increment
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	// decrement
	for i := 10; i >= 1; i-- {
		fmt.Println(i)
	}

	// cara kedua
	i := 1
	for i < 10 {
		fmt.Println("angka ke-", i)
		i++
	}

}
