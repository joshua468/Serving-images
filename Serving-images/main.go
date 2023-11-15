package main

import (
	"log"
	"net/http"
)

func main() {
	const FilePath = "."
	port := "8080"

	mux := http.NewServeMux()
	CorsMux := middlewareCors(mux)
	mux.Handle("/", http.FileServer(http.Dir(FilePath)))
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: CorsMux,
	}
	log.Printf("starting server at port %s\n", port)
	log.Fatal(srv.ListenAndServe())
}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
