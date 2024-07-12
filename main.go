package main

import "net/http"

func main(){
	serveMux := http.NewServeMux()
	server := &http.Server{Addr: "localhost:8080", Handler: serveMux}
	serveMux.Handle("/app/*", http.StripPrefix("/app", http.FileServer(http.Dir("."))))
	serveMux.HandleFunc("/healthz", func(w http.ResponseWriter, req *http.Request){
		w.Header().Add("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(http.StatusText(http.StatusOK)))
	})
	server.ListenAndServe()
}
