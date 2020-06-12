package main

import "fmt"

func main() {
	fmt.Println("belajar golang")      // string
	fmt.Println("belajar " + "golang") // concat
	fmt.Println(len("hello"))          // length

	fmt.Println("belajar"[0])   // get character index 0 berupa byte
	fmt.Println("belajar"[0:4]) // get string from index 0 -> 3
	fmt.Println("SammiDev"[5:]) // get data from index 4 -> last
}
