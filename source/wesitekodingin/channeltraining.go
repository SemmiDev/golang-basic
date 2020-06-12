package main

import "fmt"

func main() {
	chdata := make(chan int)
	data := []int{1, 2, 3, 4, 5, 6, 7}
	for _, y := range data {
		go func(y int) {
			chdata <- y
		}(y)
	}

	for i := 0; i < len(data); i++ {
		fmt.Println(<-chdata)
	}
}
