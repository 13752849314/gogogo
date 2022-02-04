package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println(123)
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "Hello World!", r.URL.Path)
	if err != nil {
		return
	}
}
