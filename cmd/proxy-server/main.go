package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Println("error :", err)
		os.Exit(1)
	}
}

const port = ":3000"

func run() error {
	log.Print("starting proxy server")
	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   "localhost:3001",
	})

	if err := http.ListenAndServe(port, proxy); err != nil {
		return err
	}

	return nil
}
