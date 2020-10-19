package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello!")
}

func main() {
	cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	if err != nil {
		log.Fatal(err)
	}
	l, err := tls.Listen("tcp", "localhost:8443", &tls.Config{
		Certificates: []tls.Certificate{cert},
	})
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", index)
	if err := http.Serve(l, nil); err != nil {
		log.Panic(err)
	}
}
