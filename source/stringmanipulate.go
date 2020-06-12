package main

import (
	"fmt"
	"strings"
)

func main() {

	var myString = "Sammi aldhi yanto,sammidev,programmer,developer"
	myString2 := "golang"

	fmt.Println("lower : " + myString)

	myUpperString := strings.ToUpper(myString)
	fmt.Println("biasa yg di upper : " + myUpperString)

	myLowerString := strings.ToLower(myUpperString)
	fmt.Println("yg diupper di lower kembali : " + myLowerString)

	containsString := strings.Contains(myString2, "golang") // "substring" | bersifat case sensitive
	fmt.Println(containsString)                             // true

	resultSplit := strings.Split(myString, ",")
	for _, out := range resultSplit {
		fmt.Println(out)
	}
}
