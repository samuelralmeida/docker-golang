package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested the path: %s\n", r.URL.Path)
	})

	_ = http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
