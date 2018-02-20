package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "response body goes here")
	})
	fmt.Println("up and running")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
