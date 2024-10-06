package main

import (
	"log"

	"github.com/halladj/dis-log/internal/server"
)

func main() {

	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
