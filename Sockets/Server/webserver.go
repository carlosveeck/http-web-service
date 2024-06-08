package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "Bem vindo ao meu site!\n")
}
func getTime(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /time request\n")
	currentTime := time.Now().Format(time.RFC1123)
	io.WriteString(w, fmt.Sprintf("Server Time: %s\n", currentTime))
}
func main() {

	http.HandleFunc("/", getRoot)

	http.HandleFunc("/time", getTime)

	err := http.ListenAndServe(":5555", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("servidor fechou\n")
	} else if err != nil {
		fmt.Printf("error ao inicializar o servidor: %s\n", err)
		os.Exit(1)
	}
}