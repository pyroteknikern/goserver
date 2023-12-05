package main

import (
    "log"
    "html/template"
    "net/http"
    "github.com/vartanbeno/go-reddit/v2/reddit"
)

func handler(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("html/view.html")
    t.Execute(w, nil)
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
