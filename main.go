package main
 
import (
    "net/http"
    "github.com/NoArc28/eks-workshop-sample-api-service-go/html"
)
 
func main() {   
    http.Handle("/", http.FileServer(http.Dir("html/")))
    //http.Handle("/static", http.FileServer(http.Dir("wwwroot")))
    http.ListenAndServe(":8080", nil)
}
