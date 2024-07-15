package main

import (
	"net/http"
	"strconv"
)

type apiConfig struct {
	fileserverHits int
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits++
		next.ServeHTTP(w, r)
	})
}

func (cfg *apiConfig) handleMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("Hits: " + strconv.Itoa(cfg.fileserverHits)))
}

func (cfg *apiConfig) handleReset(w http.ResponseWriter, r *http.Request) {
	cfg.fileserverHits = 0
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("Reset complete"))
}

func handleStatus(w http.ResponseWriter, req *http.Request){
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}

func main(){
	apiCfg := &apiConfig{}
	serveMux := http.NewServeMux()
	server := &http.Server{Addr: "localhost:8080", Handler: serveMux}
	handler := http.StripPrefix("/app", http.FileServer(http.Dir(".")))
	serveMux.Handle("/app/*", apiCfg.middlewareMetricsInc(handler))
	serveMux.HandleFunc("/healthz", handleStatus)
	serveMux.HandleFunc("/metrics", apiCfg.handleMetrics)
	serveMux.HandleFunc("/reset", apiCfg.handleReset)
	server.ListenAndServe()
}
