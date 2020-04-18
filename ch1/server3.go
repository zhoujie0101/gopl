package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, req *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "%s %s %s\n", req.Method, req.URL, req.Proto)
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", req.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", req.RemoteAddr)
	if err := req.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range req.Form {
		fmt.Fprintf(w, "Form[%q] = %q", k, v)
	}
}

func counter(w http.ResponseWriter, req *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}