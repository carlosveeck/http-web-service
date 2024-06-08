package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
)

func main() {
    // Faz uma requisição GET para o servidor
	resp, err := http.Get("http://localhost:5555/time")
    if err != nil {
        log.Fatalf("Erro ao fazer a requisição: %v", err)
    }
    defer resp.Body.Close()

    // Lê o corpo da resposta
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatalf("Erro ao ler a resposta: %v", err)
    }

    // Imprime a resposta
    fmt.Printf("Resposta do servidor: %s\n", body)
}