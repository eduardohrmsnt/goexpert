package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	c := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.google.com", nil)

	req.Header.Set("Content-Type", "application/json")

	req.Header.Set("Authorization", "Bearer 1234567890")
	if err != nil {
		panic(err)
	}

	res, err := c.Do(req)

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
