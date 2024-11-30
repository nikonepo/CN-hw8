package main

import (
    "flag"
    "fmt"
    "os"
)

func main() {
    ip := flag.String("ip", "127.0.0.1", "IP address")
    mode := flag.String("mode", "tcp-server", "Mode: tcp-server, tcp-client")
    key := flag.String("key", "", "Key file")
    crt := flag.String("crt", "", "Certificate file")

    flag.Parse()

    os.Setenv("SSLKEYLOGFILE", "sslkey.log")

    switch *mode {
    case "tcp-server":
        StartTcpServer(*ip, *key, *crt)
    case "tcp-client":
        StartTcpClient(*ip, *key, *crt)
    default:
        fmt.Println("Illegal mode")
    }
}
