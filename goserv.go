package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
	"net"
	"time"
	"flag"
	"strconv"
	"os"
)

var graphURL = flag.String("gurl", "192.168.1.138", "url of graphite server")
var graphPort = flag.String("gport", "2003", "graphite port")
var listenPort = flag.String("lport", "9090", "local server listen port")
var listenAddress = flag.String("laddress", "", "local server address")
var getTrue = flag.Bool("get", true, "set server to to parse GET requests via URL (classic functionality)")
var postTrue = flag.Bool("post", false, "set default router to parse POST")

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
	flag.Parse()
	conn, err := net.Dial("tcp", *graphURL)
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

func info(w http.ResponseWriter, r *http.Request){
	flag.Parse()
	lAdd := *listenAddress
	if lAdd == "" {
		lAdd = "localhost"
	}
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "error could not find hostname"
	}
	text := "server running on: "+lAdd+" ("+hostname+")"+`
server running on port: `+*listenPort+`
graphite server address: `+*graphURL+
`
graphite server port: `+*graphPort
	fmt.Fprint(w, text)
}

func main() {
	flag.Parse()

	//make defaults turn off 
	if *postTrue {
		*getTrue = false
	}
	
	switch {
	case *getTrue:
		go http.HandleFunc("/", sayhelloName) 
	case *postTrue:
		//Stub this out for now with the info route
		go http.HandleFunc("/", info)
	}
	
	go http.HandleFunc("/info/", info)
	err := http.ListenAndServe(*listenAddress+":"+*listenPort, nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
