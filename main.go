package main

import "net/http"

func main(){
	serveMux := http.NewServeMux()
	server := &http.Server{Addr: "localhost:8080", Handler: serveMux}
	serveMux.Handle("/", http.FileServer(http.Dir(".")))
	server.ListenAndServe()
}
