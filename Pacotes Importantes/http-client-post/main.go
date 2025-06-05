package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	http := &http.Client{
		Timeout: time.Second,
	}

	res, err := http.Post("https://www.google.com", "application/json", strings.NewReader(`{"name": "John", "age": 30}`))
	//or
	res, err = http.Post("https://www.google.com", "application/json", bytes.NewBuffer([]byte(`{"name": "John", "age": 30}`)))

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	io.CopyBuffer(os.Stdout, res.Body, nil)
}
