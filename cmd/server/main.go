package main

import (
	"log"

	"github.com/kathy7890/distributedServices/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}