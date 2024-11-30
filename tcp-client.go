package main

import (
    "bufio"
    "crypto/tls"
    "fmt"
    "os"
    "strings"
)

func StartTcpClient(ip string, key string, crt string) {

    f, err := os.OpenFile("sslkey.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
    cer, err := tls.LoadX509KeyPair(crt, key)
    config := tls.Config{Certificates: []tls.Certificate{cer}, InsecureSkipVerify: true, KeyLogWriter: f}

    c, err := tls.Dial("tcp", ip+":443", &config)

    if err != nil {
        fmt.Println("Error connecting to TCP server:", err)
        return
    }

    defer c.Close()

    buffer := make([]byte, 1024)
    consoleReader := bufio.NewReader(os.Stdin)

    fmt.Println("Connected to TCP server:", ip)

    for {
        fmt.Println("Enter message to send to server:")
        text, _ := consoleReader.ReadString('\n')
        text = strings.TrimSpace(text)

        if text == "big" {
            largeMessage := strings.Repeat("A", 40000)
            c.Write([]byte(largeMessage))
            fmt.Println("Sent large message to server")
        } else {
            c.Write([]byte(text))
        }

        n, err := c.Read(buffer)
        if err != nil {
            fmt.Println("Error reading from TCP connection:", err)
            return
        }

        message := string(buffer[:n])
        fmt.Println("Message from server:", message)
    }
}
