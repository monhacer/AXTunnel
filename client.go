package main

import (
  "crypto/tls"
  "encoding/json"
  "fmt"
  "io"
  "net"
  "os"
)

type Config struct {
  RemoteIP    string
  RemotePorts []string
  LocalPort   string
  Password    string
}

func main() {
  var cfg Config
  f, err := os.Open("config.json")
  if err != nil {
    fmt.Println("Cannot read config:", err)
    return
  }
  json.NewDecoder(f).Decode(&cfg)

  ln, err := net.Listen("tcp", ":"+cfg.LocalPort)
  if err != nil {
    fmt.Println("Cannot listen on local port:", err)
    return
  }

  fmt.Println("AXTunnel client listening on port", cfg.LocalPort)

  for {
    localConn, err := ln.Accept()
    if err == nil {
      go handle(localConn, cfg)
    }
  }
}

func handle(local net.Conn, cfg Config) {
  var remote net.Conn
  for _, port := range cfg.RemotePorts {
    fmt.Println("Trying remote port:", port)
    tlsConfig := &tls.Config{ServerName: "www.google.com"}
    r, err := tls.Dial("tcp", cfg.RemoteIP+":"+port, tlsConfig)
    if err == nil {
      remote = r
      break
    }
  }

  if remote == nil {
    fmt.Println("No remote port available")
    local.Close()
    return
  }

  defer remote.Close()
  defer local.Close()

  go io.Copy(remote, local)
  io.Copy(local, remote)
}
