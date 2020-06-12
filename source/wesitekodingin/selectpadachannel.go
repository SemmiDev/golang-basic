package main

import "fmt"

func cetakHuruf1(ch chan string) {
	ch <- "sedang belajar golang"
}
func cetakHuruf2(ch chan string) {
	ch <- "sedang belajar java"
}

func main() {
	cHuruf1 := make(chan string)
	cHuruf2 := make(chan string)

	go func() {
		cetakHuruf1(cHuruf1)
	}()

	go func() {
		cetakHuruf2(cHuruf2)
	}()

	select {
	case c1 := <-cHuruf1:
		fmt.Println(c1)
	case c2 := <-cHuruf2:
		fmt.Println(c2)
	}
}
