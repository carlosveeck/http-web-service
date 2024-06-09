package main

import (
	"log"
	"net/rpc"
	"time"
)

type Args struct {
}

func main() {
	var reply string
	args := Args{}
    start := time.Now() // Marca o início do tempo de requisição
	client, err := rpc.DialHTTP("tcp", "localhost"+":2233")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	err = client.Call("TimeServer.GiveServerTime", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	elapsed := time.Since(start) // Marca o fim do tempo de requisição

	log.Printf("Resposta do servidor: %s", reply)
	log.Printf("Tempo de resposta: %s", elapsed)
}
