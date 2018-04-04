package main

import (
	"log"
	"net/http"

	"github.com/rakyll/statik/fs"

	_ "github.com/krishnasrinivas/vmware-app/statik" // TODO: Replace with the absolute import path
)

func main() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", http.FileServer(statikFS))
	http.ListenAndServe(":8080", nil)
}
