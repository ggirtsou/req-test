package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	var i int
	for {
		makeReq()
		fmt.Printf("made %d requests, wait\n", i)
		i++
		time.Sleep(1 * time.Millisecond)
	}
}

func makeReq() {
	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
}
