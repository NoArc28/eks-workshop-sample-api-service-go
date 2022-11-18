package main
 
import (
    "net/http"
)
 
func main() {   
    http.Handle("/", http.FileServer(http.Dir("html")))
    //http.Handle("/static", http.FileServer(http.Dir("wwwroot")))
    http.ListenAndServe(":8080", nil)
}
