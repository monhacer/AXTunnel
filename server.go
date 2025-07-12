package main

import (
  "crypto/tls"
  "io"
  "net"
  "fmt"
)

func main() {
  cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
  if err != nil {
    fmt.Println("Certificate loading error:", err)
    return
  }

  config := &tls.Config{Certificates: []tls.Certificate{cert}}
  ln, _ := tls.Listen("tcp", ":443", config)
  fmt.Println("AXTunnel server listening on port 443")

  for {
    conn, _ := ln.Accept()
    go handle(conn)
  }
}

func handle(c net.Conn) {
  target, _ := net.Dial("tcp", "127.0.0.1:22") // Change destination port as needed
  go io.Copy(target, c)
  io.Copy(c, target)
}
