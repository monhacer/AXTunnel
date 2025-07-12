package main

import (
  "crypto/tls"
  "io"
  "net"
  "fmt"
  "os"
)

func main() {
  port := "443" // default
  if len(os.Args) > 1 {
    port = os.Args[1]
  }

  cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
  if err != nil {
    fmt.Println("Certificate error:", err)
    return
  }

  config := &tls.Config{Certificates: []tls.Certificate{cert}}
  listener, err := tls.Listen("tcp", ":"+port, config)
  if err != nil {
    fmt.Println("Listen error:", err)
    return
  }

  fmt.Println("AXTunnel server listening on port", port)

  for {
    conn, err := listener.Accept()
    if err == nil {
      go handle(conn)
    }
  }
}

func handle(c net.Conn) {
  target, err := net.Dial("tcp", "127.0.0.1:22") // local, changab
  if err != nil {
    c.Close()
    return
  }
  go io.Copy(target, c)
  io.Copy(c, target)
}
