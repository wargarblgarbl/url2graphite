package main

import (
    "fmt"
    "net/http"
    "strings"
    "log"
)


func sayhelloName(w http.ResponseWriter, r *http.Request) {
	//Turn into slice
	x := strings.Split(r.URL.Path, "/")
	//last portion of the slice
	l := x[len(x)-1]
    	//beginning of the slice
	f := x[:len(x)-1]
	//let's turn f into a string
	z := strings.Join(f[:],".")
	//Minor cleanup, want to get rid of that first .
	u := strings.Replace(z, ".", "", 1)	
	fmt.Println(u, l)   
	fmt.Fprintf(w, r.URL.Path) 
}

func main() {
    http.HandleFunc("/", sayhelloName) // set router
    err := http.ListenAndServe(":9090", nil) // set listen port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
