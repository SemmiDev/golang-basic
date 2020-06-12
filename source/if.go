package main

import "fmt"

func main() {
	// for i := 1; i <= 10; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Println(i, " : genap")
	// 	} else if i%2 != 0 {
	// 		fmt.Println(i, " : ganjil")
	// 	} else {
	// 		fmt.Println("default")
	// 	}

	// }

	for i := 1; i < 30; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fooBar")
		} else if i%3 == 0 {
			fmt.Println("foo")
		} else if i%5 == 0 {
			fmt.Println("bar")
		} else {
			fmt.Println(i)
		}
	}

}
