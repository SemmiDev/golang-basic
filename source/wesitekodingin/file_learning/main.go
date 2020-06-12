package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// use Util
	data, err := ioutil.ReadFile("dataset2.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

	// use Os
	file, err2 := os.Open("dataset2.txt")
	if err2 != nil {
		log.Fatal(err2)
	}
	toByte := make([]byte, 1000) // 1k is range
	for {
		count, err := file.Read(toByte)
		if err != nil {
			break
		}
		fmt.Println(string(toByte[0:count]))
	}

}
