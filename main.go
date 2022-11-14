package main

import (
	"log"
	"net/http"
)
import "core_demo/framework"

func main() {
	server := &http.Server{
		Addr:    ":8080",
		Handler: framework.NewCore(),
	}

	log.Fatal(server.ListenAndServe())
}
