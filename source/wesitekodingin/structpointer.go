package main

import "fmt"

type (
	profil struct {
		ID      int
		name    string
		Age     int
		Address string
	}
)

func main() {
	// profil := &profil{
	// 	ID:      1,
	// 	name:    "sam",
	// 	Age:     17,
	// 	Address: "lepau daryen",
	// }
	// fmt.Println(*profil)

	// property lebih dari satu
	profil2 := []profil{{
		ID:      1,
		name:    "dev",
		Age:     18,
		Address: "tinggam",
	},
		{
			ID:   2,
			name: "sammi",
			Age:  20,
			Address: "tinggam	",
		},
	}

	for key, val := range profil2 {
		fmt.Println(key, val)
	}
}
