package main

import (
	"gorestapi/api"
	"net/http"
)

func main() {

	// Hello word api
	/*
		http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hi There, This is Jai. Welcome to go-lang rest demo api"))
		})
	*/
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}
