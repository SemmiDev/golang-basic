package main

import (
	"fmt"
	"time"
)

func tulis(data []string, ch chan string) {
	for _, dt := range data {
		ch <- dt
	}
	close(ch)
}

func cetak(ch chan string) {
	for result := range ch {
		fmt.Println(result)
		time.Sleep(2 * time.Second)
	}
}

func main() {
	data := []string{"JavaScript", "Golang", "C++", "Java"}
	ch := make(chan string, len(data))
	go tulis(data, ch)
	cetak(ch)
}
