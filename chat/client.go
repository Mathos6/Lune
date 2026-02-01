package main

import (
  "context"
  "crypto/tls"
  "fmt"
  "time"

  "github.com/quic-go/quic-go"
)

func main() {
  tlsConf := &tls.Config{
    InsecureSkipVerify: true,
    NextProtos:         []string{"quic-ping"},
  }

  start := time.Now()

  conn, err := quic.DialAddr(
    context.Background(),
    "127.0.0.1:4242",
    tlsConf,
    nil,
  )
  if err != nil {
    panic(err)
  }

  stream, _ := conn.OpenStreamSync(context.Background())
  stream.Write([]byte("ping"))

  buf := make([]byte, 1024)
  n, _ := stream.Read(buf)

  elapsed := time.Since(start)

  fmt.Println("Reply:", string(buf[:n]))
  fmt.Println("RTT:", elapsed)

  stream.Close()
  conn.CloseWithError(0, "")
}
