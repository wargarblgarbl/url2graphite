package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
	"net"
	"time"
	"strconv"
)


func sayhelloName(w http.ResponseWriter, r *http.Request) {
	const (
	graphite_ip = "192.168.1.138:2003"
	)
	//turn into slice
	x := strings.Split(r.URL.Path, "/")
	//last portion of the slice
	l := x[len(x)-1]
    	//beginning of the slice
	f := x[:len(x)-1]
	//let's turn f into a string
	z := strings.Join(f[:],".")
	//Minor cleanup, want to get rid of that first .
	u := strings.Replace(z, ".", "", 1)
	//let's get the date
	now :=time.Now()
	date := strconv.Itoa(int(now.Unix()))
	// let's concat this into the one thing we want it to be.(newline is necessary)
	packet := u +" "+ l+ " "+date+"\n"
	fmt.Println("processed:", packet)
	fmt.Fprintf(w, r.URL.Path)
	//let's send this stuff to TCP
	fmt.Println("sent:", packet)
	conn, err := net.Dial("tcp", graphite_ip)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	_, err = fmt.Fprintf(conn, packet)
	if err != nil {
		fmt.Printf("fatal error", err)
		return
	}
	if tcpcon, ok := conn.(*net.TCPConn); ok {
		tcpcon.CloseWrite()
	}
	err = conn.Close()
	if err != nil {
		fmt.Printf("fatal error2", err)
	}


}



func main() {
	const (listen_port=":9090" )
	go http.HandleFunc("/", sayhelloName) // set router
	err := http.ListenAndServe(listen_port, nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
