package main

import "fmt"

func main() {
	// closhure sederhana
	gol1 := func() {
		fmt.Println("belajar golang")
	}
	gol2 := func() {
		fmt.Println("belajar java")
	}

	gol1()
	gol2()

	// membuat fungsi closure dengan nilai balik
	dataTodo := todo()
	fmt.Println("nilai todo :", dataTodo())

}

func todo() func() int {
	i := 4
	return func() int {
		i++
		return i
	}
}
