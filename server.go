package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
)

var portFlag = flag.String("p", "8080", "Specify directory to serve. (default: 8080)")
var servecount int

func listen() {
	//listen for ctrl-c
	go func() {
		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, os.Interrupt)
		<-sigchan
		log.Printf("Served %d requests", servecount)
		log.Println("Terminating whatismyipserver, Goodbye!")
		os.Exit(0)
	}()
	log.Printf("Listening on port %s", *portFlag)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		servecount = servecount + 1
		fmt.Fprintf(w, "Your IP address is: %s\n", r.RemoteAddr)
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
	})

	log.Fatal(http.ListenAndServe(strings.Join([]string{":", *portFlag}, ""), nil))
}

func init() {
	servecount = 0
}

func main() {
	// Simple static webserver:
	flag.Parse()
	listen()
}
