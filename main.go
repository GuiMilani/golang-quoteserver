package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"golang_quoteserver/quote_generator"
)

func getKanye(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got kanye request\n")
	io.WriteString(w, quote_generator.GetQuote("kanye"))
}

func getTaylor(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got taylor request\n")
	io.WriteString(w, quote_generator.GetQuote("taylor"))
}


func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/kanye", getKanye)
	mux.HandleFunc("/taylor", getTaylor)

	err := http.ListenAndServe(":3333", mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}