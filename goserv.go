package main

import (
    "fmt"
    "net/http"
    "strings"
    "log"
)




func sayhelloName(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.URL.Path)
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
	fmt.Fprintf(w, r.URL.Path)
}

func main() {
    http.HandleFunc("/", sayhelloName) // set router
    err := http.ListenAndServe(":9090", nil) // set listen port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
