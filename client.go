package main

import (
  "log"
  "net"
  "crypto/tls"
  "io/ioutil"
  "github.com/quic-go/quic-go"
  "context"
  "time"
)

func main() {
ctx, cancel := context.WithTimeout(context.Background(),
  3*time.Second)
defer cancel()
certPEM, err := ioutil.ReadFile("cert.pem")
if err != nil {
  log.Fatal("Problème lors de la lecture de cert.pem: ", err)
}
keyPEM, err := ioutil.ReadFile("key.pem")
if err != nil {
  log.Fatal("Problème lors de la lecture de key.pem: ", err)
}
cert, err := tls.X509KeyPair(certPEM, keyPEM)
if err != nil {
  log.Fatal("Error dealing with it", err)
}
tlsConf := &tls.Config{
  Certificates : []tls.Certificate{cert},
  NextProtos : []string{"my-quic-protocol"},
  }
quicConf := &quic.Config{}
addr, err := net.ResolveUDPAddr("udp4", "127.0.0.1:8000")
if err != nil {
  log.Fatal("Erreur DNS: ", err)
}
conn, err := net.ListenUDP("udp", addr)
if err != nil {
  log.Fatal("Erreur Listen: ", err)
}
trp := quic.Transport{Conn:conn}
sess, err := trp.Dial(ctx, addr, tlsConf, quicConf)
if err != nil {
  log.Fatal("Error when dialing server: ", err)
}
for {
  stream, err := sess.OpenStreamSync(ctx)
  if err != nil {
    log.Fatal("Error when opening a stream: ", err)
  }
  go pingServer(stream)
}

}

func pingServer(stream *quic.Stream) {
  defer stream.Close()
  msg := make([]byte, 4)
  msg = []byte("ping")
  _, err := stream.Write(msg)
  if err != nil {
    log.Fatal("Error while writing: ", err)
  }
}
