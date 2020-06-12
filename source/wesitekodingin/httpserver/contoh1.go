// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// // please uncomment
// // const port string = "9040"

// func welcome(w http.ResponseWriter, r *http.Request) {
// 	dataHTML := `<h1>Bahasa Pemrograman terkeren 2020</h1>
// 	<ul>
// 		<li>JavaScript</li>
// 		<li>Python</li>
// 		<li>Golang</li>
// 		<li>Java</li>
// 		<li>PHP</li>
// 		<li>Rust</li>
// 	</ul>

// 	Tutorial by : <a href="file:///home/sammidev/Videos/Tutorial/share/Belajar%20Golang%20%2337%20_%20Menjalankan%20HTTP%20Server%20-%20WEB%20_%20Kodingin.mhtml">Check</a>
// 	`

// 	if r.Method == "GET" {
// 		fmt.Fprintf(w, dataHTML)
// 	}
// }
// func main() {
// 	fmt.Println("berjalan di port : " + port)
// 	http.HandleFunc("/", welcome)

// 	if err := http.ListenAndServe(":"+port, nil); err != nil {
// 		log.Fatal(err)
// 	}
// }
