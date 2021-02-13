package main

import (
	"golangweb/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	aboutHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is about page"))
	}

	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/hello", handler.HelloHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/product", handler.ProductHandler)
	mux.HandleFunc("/detail", handler.DetailHandler)
	mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is Profile Page"))
	})
	mux.HandleFunc("/post-get", handler.PostGet)
	mux.HandleFunc("/form", handler.FormHandler)
	mux.HandleFunc("/process", handler.ProcessHandler)

	// static file handling (css, js, and images)
	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting web on port 8080")

	err := http.ListenAndServe("localhost:8080", mux)
	log.Fatal(err)
}
