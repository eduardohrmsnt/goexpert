package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	_ "modernc.org/sqlite"
)

type CotacaoResponse struct {
	Cotacao Cotacao `json:"USDBRL"`
}

type Cotacao struct {
	Valor string `json:"bid"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/cotacao", obterCotacao)

	http.ListenAndServe(":8080", mux)
}

func obterCotacao(responseWriter http.ResponseWriter, request *http.Request) {

	cotacao, err := buscarCotacaoApi()

	if err != nil {
		if err == context.DeadlineExceeded {
			log.Println("Request timed out")
		}
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte("500 - Internal Server Error"))
		return
	}

	err = salvarCotacao(cotacao)

	if err != nil {
		if err == context.DeadlineExceeded {
			log.Println("Request timed out")
		}
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte("500 - Internal Server Error"))
		return
	}

	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(cotacao)
}

func salvarCotacao(cotacao *Cotacao) error {
	ctxDb := context.Background()
	ctxDb, cancelDb := context.WithTimeout(ctxDb, time.Millisecond*10)

	defer cancelDb()

	db, err := sql.Open("sqlite", ":memory:")

	if err != nil {
		log.Println(err)
		return err
	}

	defer db.Close()
	_, err = db.ExecContext(ctxDb, `CREATE TABLE IF NOT EXISTS cotacao (
		valor TEXT
	)`)

	_, err = db.ExecContext(ctxDb, "INSERT INTO cotacao (valor) VALUES (?)", &cotacao.Valor)

	if err != nil {
		log.Println(err)
		return err
	}

	return err
}

func buscarCotacaoApi() (*Cotacao, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)

	if err != nil {
		return nil, err
	}

	content, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer content.Body.Close()
	body, err := io.ReadAll(content.Body)

	var cotacao CotacaoResponse
	err = json.Unmarshal(body, &cotacao)

	return &cotacao.Cotacao, nil
}
