package main
 
import (
    "log"
    "net/http"
    "github.com/api-mysql/config"
    "fmt"
)
 
func main() {
 
    db, e := config.MySQL()
 
    if e != nil {
        log.Fatal(e)
    }
 
    eb := db.Ping()
    if eb != nil {
        panic(eb.Error()) 
    }
 
    fmt.Println("Success")
 
    err := http.ListenAndServe(":7000", nil)
 
    if err != nil {
        log.Fatal(err)
    }
}
