package main

import (
	"fmt"
	"time"
)

// func printNum() {
// 	angka := []string{"iye ke?"}
// 	for x := range angka {
// 		time.Sleep(500 * time.Millisecond)
// 		fmt.Println(angka[x])
// 	}
// }

func one() {
	huruf := []string{"Malam,", "he", "he", "he :)", " A", "k", "u", " ", "r", "i", "n", "d", "u", " ", "n", "", "i", "h"}
	for x := range huruf {
		time.Sleep(200 * time.Millisecond)
		fmt.Print(huruf[x])
	}
}

func main() {
	// go printNum()
	fmt.Print("Aku : ")
	go one()
	// go printNum()

	time.Sleep(4100 * time.Millisecond)
	fmt.Println("")
	fmt.Println("Dia : you has been blocked :(")
}
