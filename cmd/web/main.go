package main

import (
	"embed"
	"io/fs"
	"log"
  "net/http"
)

//go:embed static/*
var static embed.FS

func main() {
	public, err := fs.Sub(static, "static")
	if err != nil {
		panic(err)
	}
  http.Handle("/", http.FileServer(http.FS(public)))
	log.Printf("Server listening on http://localhost:%s/", "8080")
	http.ListenAndServe(":8080", nil)
}
