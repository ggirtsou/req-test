package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	var i int

	maxConcurrency := 100
	tokens := make(chan struct{}, maxConcurrency)
	for i := 0; i < maxConcurrency; i++ {
		tokens <- struct{}{}
	}

	for {
		<-tokens
		go func() {
			defer func() { tokens <- struct{}{} }()
			makeReq()
			fmt.Printf("made %d requests\n", i)
			i++
		}()
	}
}

func makeReq() {
	tr := http.Transport{MaxIdleConns: 1}
	client := &http.Client{Transport: &tr}
	resp, err := client.Get("http://localhost:8080")
	if err != nil {
		panic(err)
	}

	// if you remove this, you end up making more requests
	io.Copy(ioutil.Discard, resp.Body)

	// do not remove this
	defer resp.Body.Close()
}
