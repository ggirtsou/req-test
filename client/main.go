package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	maxConcurrency = 800
	maxIdle        = 1000
)

func main() {
	var i int

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

var (
	tr     = http.Transport{MaxIdleConns: maxIdle, MaxIdleConnsPerHost: maxIdle}
	client = &http.Client{Transport: &tr}
)

func makeReq() {
	resp, err := client.Get("http://localhost:8080")
	if err != nil {
		panic(err)
	}

	// if you remove this, sockets won't be released
	io.Copy(ioutil.Discard, resp.Body)

	// do not remove this
	defer resp.Body.Close()
}
