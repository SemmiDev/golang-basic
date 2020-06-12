package main

import "fmt"

func main() {

	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			// jika habis di modulo 2 maka akan di skip
			continue
		} else {
			fmt.Println(i)
		}
	}
	fmt.Println("\n\n")

	for i := 1; i <= 20; i++ {
		if i == 10 {
			// jika pas di perulangan ke 20 makan akan di hentikan
			break
		} else {
			fmt.Println(i)
		}
	}
	fmt.Println("program selesai")

}
