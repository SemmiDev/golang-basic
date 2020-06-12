package main

import (
	"fmt"
	"sync"
)

func cetak(name string, wg *sync.WaitGroup) {
	wg.Done()
	fmt.Println(name)
}
func main() {
	/*
		Add(), digunakan untuk menambah goroutine yang jalan.
		Wait(), digunakan untuk menunggu proses berjalan.
		Done(), digunakan untuk menandai bahwa semua goroutine sudah selesai.

	*/
	var wg sync.WaitGroup
	dataMhs := []string{"didik", "sam", "gusnur", "adit", "otong"}
	for _, v := range dataMhs {
		wg.Add(1)
		go cetak(v, &wg)
	}
	wg.Wait()
}
