package main

import (
	"net/http"
	"log"
	"net/http/httputil"
	"net/url"
)

func main() {
	targetURL, err := url.Parse("https://nextendo.online")
	if err != nil {
		log.Fatal(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	})

	log.Println("Proxy server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
