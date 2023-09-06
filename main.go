package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hola mundo")
	})

	srv := http.Server{
		Addr: ":8080",
	}

	srv.ListenAndServe()
}
