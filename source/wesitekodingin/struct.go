package main

import "fmt"

type profile struct {
	ID      int
	Name    string
	Age     int
	Address string
}

func main() {

	// call model 1
	var profile1 profile
	profile1.ID = 1
	profile1.Name = "Sammi"
	profile1.Age = 19
	profile1.Address = "Tinggam"

	fmt.Println(profile1)

	// call model 2
	profile2 := profile{
		ID:      2,
		Name:    "dev",
		Age:     15,
		Address: "Tinggam",
	}
	fmt.Println(profile2)

	// call anonymous struct
	anony := struct {
		ID      int
		Name    string
		Age     int
		Address string
	}{
		ID:      3,
		Name:    "Sammi Aldhi Yanto",
		Age:     13,
		Address: "lepau daryen",
	}

	fmt.Println(anony)

	// nedted struct
	profileChildreen1 := profileChildreen{
		ID:      4,
		Name:    "anonymous people",
		Age:     40,
		Address: "lepau daryen",
		Education: Education{
			SMA: "SMAN 1 Talamau",
			SMP: "SMP 1 Talamau",
		},
	}

	fmt.Println(profileChildreen1)

}

type (
	Education struct {
		SMA string
		SMP string
	}

	profileChildreen struct {
		ID        int
		Name      string
		Age       int
		Address   string
		Education Education
	}
)
