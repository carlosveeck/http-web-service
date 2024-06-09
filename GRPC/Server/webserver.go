package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Args struct{}
type TimeServer string

func (t *TimeServer) GiveServerTime(args *Args, reply *string) error {

	*reply = time.Now().Format(time.RFC1123)
	return nil
}
func main() {

	timeserver := new(TimeServer)

	rpc.Register(timeserver)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":2233")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
}
