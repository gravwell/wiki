package main

import (
	"log"
	"net/http"
)

// Wrap the given handler, adding headers to disable caching
func noCacheLoggingHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
		h.ServeHTTP(w, r)
	})
}

func main() {
	http.Handle("/", noCacheLoggingHandler(http.FileServer(http.Dir("."))))
	log.Fatal(http.ListenAndServe(":3001", nil))
}
