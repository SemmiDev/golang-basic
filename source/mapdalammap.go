package main

import "fmt"

func main() {

	employee := map[string]map[string]string{
		"1001": map[string]string{
			"username": "=> Sammi dev",
			"password": "=> thispass",
			"skill":    "=> java, golang",
		},
		"1002": map[string]string{
			"username": "=> Otong surotong",
			"password": "=> otongbesa",
			"skill":    "=> javascript,php",
		},
		"1003": map[string]string{
			"username": "=> ayatullah rj",
			"password": "=> bigEgg",
			"skill":    "=> marketing,financial",
		},
	}

	delete(employee, "1001") // map,key
	fmt.Println(employee["1003"]["username"])
	fmt.Println(employee["1003"]["password"])
	fmt.Println(employee["1003"]["skill"])
}
