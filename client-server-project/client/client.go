package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*300)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "localhost:8080", nil)

	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}
	_, err := f.Write([]byte("Dados: {}"))

	io.ReadAll(resp.Body)
}
