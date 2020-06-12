package main

import (
	"fmt"
)

func cetak(ch chan<- int) {
	for index := 0; index < 10; index++ {
		ch <- index
	}
	close(ch)
}

func main() {
	ch := make(chan int)

	go cetak(ch)

	for {
		data, ok := <-ch

		if ok == false {
			break
		}
		fmt.Printf("Data di terima %v\n", data)
	}
}
