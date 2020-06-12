package main

import "fmt"

func main() {
	var mp interface{}
	data := map[string]string{
		"FirstName": "Fauziah \n",
		"LastName":  "Ramadhani \n",
		"NickName":  "Eji \n",
		"Gender":    "Female \n",
		"Age":       "40 \n",
		"Address":   "Her home \n",
	}
	mp = map[string]interface{}{
		"data": data,
	}
	fmt.Println(mp)
}
