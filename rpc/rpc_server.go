package main

// considered harmful
import "./server"
import "fmt"
import "net"
import "net/http"
import "net/rpc"
import "log"

func main() {
    fmt.Println("RPC Server")
    arith := new(server.Arith)
    rpc.Register(arith)
    rpc.HandleHTTP()
    l, e := net.Listen("tcp", ":1234")
    if e != nil {
        log.Fatal("listen error:", e)
    }
    http.Serve(l, nil)
}
