package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	http := &http.Client{
		Timeout: time.Second,
	}

	res, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
