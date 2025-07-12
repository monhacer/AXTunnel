package main

import (
  "crypto/tls"
  "encoding/json"
  "io"
  "net"
  "os"
)

type Config struct {
  RemoteIP   string
  RemotePort string
  LocalPort  string
  Password   string
}

func main() {
  var cfg Config
  f, _ := os.Open("config.json")
  json.NewDecoder(f).Decode(&cfg)

  ln, _ := net.Listen("tcp", ":"+cfg.LocalPort)
  for {
    localConn, _ := ln.Accept()
    go handle(localConn, cfg)
  }
}

func handle(local net.Conn, cfg Config) {
  tlsConfig := &tls.Config{ServerName: "www.google.com"}
  remote, err := tls.Dial("tcp", cfg.RemoteIP+":"+cfg.RemotePort, tlsConfig)
  if err != nil {
    return
  }
  defer remote.Close()
  go io.Copy(remote, local)
  io.Copy(local, remote)
}
