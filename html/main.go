package main

import (
    "html/template"
    "net/http"
)

func handlerequest(w http.ResponseWriter, r *http.Request) {
    title := "Hello Golang World!"
    t := template.New("index.html")
    t, _ = t.ParseFiles("index.html")
    t.Execute(w, title)
}

func main() {
    http.HandleFunc("/", handlerequest)
    http.ListenAndServe(":8000", nil)
}
