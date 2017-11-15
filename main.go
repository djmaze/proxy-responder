package main

import (
  "fmt"
  "log"
  "net"

  "github.com/armon/go-proxyproto"
)

func main() {
  ln, err := net.Listen("tcp", ":8080")
  if err != nil {
    log.Fatal(err)
  }
  proxyList := &proxyproto.Listener{Listener: ln}
  for {
    conn, err := proxyList.Accept()
    if err != nil {
      log.Fatal(err)
    }
    go handleConnection(conn)
  }
}

func handleConnection(conn net.Conn) {
    text := fmt.Sprintf("Hello, %s\n", conn.RemoteAddr().String())
    conn.Write([]byte(text))
    conn.Close()
}
