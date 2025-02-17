package main

import (
   "fmt"
   "net/http"
   
   "github.com/NoArc28/eks-workshop-sample-api-service-go"
)

func main() {
   http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintf(w, "Hyeonho aws test success!!")
   })

   http.HandleFunc("/greet/", func(w http.ResponseWriter, r *http.Request) {
      name := r.URL.Path[len("/greet/"):]
      fmt.Fprintf(w, "Hello %s\n", name)
   })

   http.ListenAndServe(":8080", nil)

}
