package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Employee struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

func encodeJSON(data []Employee) []byte {
	dataJSON, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	return dataJSON
}

func main() {
	emp := []Employee{{
		FirstName: "Sammi",
		LastName:  "Dev",
		Username:  "sammidev",
		Password:  "12345",
	},
		{
			FirstName: "Sammi2",
			LastName:  "Dev2",
			Username:  "sammidev2",
			Password:  "12345893473984",
		},
	}

	dataJSON := encodeJSON(emp)
	fmt.Println(string(dataJSON))
}
