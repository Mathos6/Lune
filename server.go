package main

import (
  "fmt"
  "context"
  "net"
  "log"
  "io"
  "io/ioutil"
  "github.com/quic-go/quic-go"
  "crypto/tls"
)


func main() {

// l'adresse
/* ResolveUDP retourne un UDPAddr et une erreur*/
addr, err := net.ResolveUDPAddr("udp4", "127.0.0.1:8000")
if err != nil {
  log.Fatal("Erreur lors de la résolution DNS: ", err)
}

/* ListenUDP retourne un *UDPConn et une erreur*/
conn, err := net.ListenUDP("udp4", addr)
if err != nil {
  log.Fatal("Erreur lors de l'établissement de la connexion: ", err)
}

trp := quic.Transport{Conn:conn}

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

lsr, err := trp.Listen(tlsConf, quicConf)
for {
  sess, err := lsr.Accept(context.Background())
  if err != nil {
    continue
  go func(sess quic.Session) {
    for {
      stream, err := sess.AcceptStream(context.Background())
      if err != nil {
        fmt.Println(err)
        continue }
      go handleStream(stream)
    } }(sess)
}
fmt.Println("Serveur lancé")
}
}

func handleStream(stream *quic.Stream) {
  buf := make([]byte, 4)
  io.ReadFull(stream, buf)
  fmt.Println(buf)
}
