package main

import (
    "fmt"
    "time"

    "github.com/Akilan1999/p2p-rendering-computation/p2p/frp"
)

// Trigger the FRP implementation
func main() {
    port, err := EscapeNATStijarna()
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Stjarna address: 139.59.162.154:" + port)

    for {

    }
}

func EscapeNATStijarna() (SSHport string, err error) {
    // Get free port from P2PRC server node
    serverPort, err := frp.GetFRPServerPort("http://139.59.162.154:8078")

    if err != nil {
        return
    }

    time.Sleep(5 * time.Second)

    //port for the barrierKVM server
    SSHport, err = frp.StartFRPClientForServer("139.59.162.154", serverPort, "5000", "8080")
    if err != nil {
        return
    }

    return
}
