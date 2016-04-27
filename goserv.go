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
const (
	graphite_ip = "192.168.1.138:2003"
)


func procRequest(input string)(output string) {
	//turn into slice
	x := strings.Split(input, "/")
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
	return packet 
}

func sendTCP(packet string)(output string){
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
	return "sent"

}


func sayhelloName(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/favicon.ico" {
		sendTCP(procRequest(r.URL.Path))

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
