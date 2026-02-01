package main

import (
  "context"
  "crypto/tls"
  "fmt"
  "log"

  "github.com/quic-go/quic-go"
)

func main() {
  tlsConf := &tls.Config{
    Certificates: []tls.Certificate{mustLoadTLS()},
    NextProtos:   []string{"quic-ping"},
  }

  listener, err := quic.ListenAddr("0.0.0.0:4242", tlsConf, nil)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("QUIC server listening on :4242")

  for {
    conn, err := listener.Accept(context.Background())
    if err != nil {
      log.Println(err)
      continue
    }
    go handleConn(conn)
  }
}

func handleConn(conn *quic.Conn) {
  stream, err := conn.AcceptStream(context.Background())
  if err != nil {
    return
  }

  buf := make([]byte, 1024)
  n, _ := stream.Read(buf)

  fmt.Println("Received:", string(buf[:n]))

  stream.Write([]byte("pong"))
  stream.Close()
}

func mustLoadTLS() tls.Certificate {
  cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
  if err != nil {
    panic(err)
  }
  return cert
}
