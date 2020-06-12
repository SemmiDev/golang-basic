package main

import "fmt"

func cetak(ch chan int) {
	fmt.Println("ini adalah goroutine")
	ch <- 10
}

func main() {
	chdata := make(chan int)
	go cetak(chdata)
	nilai := <-chdata
	fmt.Println("integer's value : ", nilai)
	fmt.Println("main func")
}
