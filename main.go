package main

import (
	"log"
	"net/http"
)

func main() {

	//This will route the request and assign a handler
	mux := http.NewServeMux()

	rh := http.RedirectHandler("http://example.org", 307)
	mux.Handle("/foo", rh)

	log.Println("Listening...")

	//Note that the signature of ListenAndServe is ListenAndServe(addr string, handler Handler)
	//But mux basically acts as a special handler
	http.ListenAndServe(":3000", mux)
}
