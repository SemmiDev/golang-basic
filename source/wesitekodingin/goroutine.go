package main

import (
	"fmt"
	"time"
)

func say() {
	fmt.Println("Selamat Malam")
}
func main() {
	go say()
	time.Sleep(1 * time.Second) // ini harus ada, testing lah
	fmt.Println("rindu ye")
}
