package main

import (
  "fmt"
  "html/template"
  "net/http"
)

func main() {
  // We're creating a file handler, here.
  fs := http.FileServer(http.Dir("/html/assets/img/about/"))

  http.HandleFunc("/html/assets/img/about/", images)

  // We're binding the handler to the `/images` route, here.
  http.Handle("/html/assets/img/about/", http.StripPrefix("/html/assets/img/about/", fs))

  http.ListenAndServe(":8080", nil)
}

func images(w http.ResponseWriter, r *http.Request) {
  t, err := template.ParseFiles("/test.html")
  if err != nil {
    fmt.Fprintf(w, err.Error())
    return
  }

  t.ExecuteTemplate(w, "html", nil)
}
