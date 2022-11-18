package main

import (
    "log"
    "os"
    "text/template"
)

var tpl *template.Template

func init() {
    tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

    sages := []string{"This", "is", "a", "string", "slice"}

    err := tpl.Execute(os.Stdout, sages)
    if err != nil {
        log.Fatalln(err)
    }
}
