package main

import (
	"fmt"
)

func lamarPT1(ch chan string) {
	ch <- "berhasil melamar di PT 1"
}
func lamarPT2(ch chan string) {
	ch <- "berhasil melamar di PT 2"
}

func main() {
	lamar1 := make(chan string)
	lamar2 := make(chan string)

	go lamarPT1(lamar1)
	go lamarPT2(lamar2)

	for i := 0; i < 2; i++ {
		select {
		case l1 := <-lamar1:
			fmt.Println(l1)
		case l2 := <-lamar2:
			fmt.Println(l2)
		}
	}
}
