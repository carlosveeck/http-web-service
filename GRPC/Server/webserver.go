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
    // Fill reply pointer to send the data back
    *reply = time.Now().Format(time.RFC1123)
    return nil
}
func main() {
    // Create a new RPC server
    timeserver := new(TimeServer)
    // Register RPC server
    rpc.Register(timeserver)
    rpc.HandleHTTP()
    // Listen for requests on port 1234
    l, e := net.Listen("tcp", ":2233")
    if e != nil {
        log.Fatal("listen error:", e)
    }
    http.Serve(l, nil)
}