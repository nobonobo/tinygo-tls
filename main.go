package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"

	orig "github.com/nobonobo/tls/orig/net"
	"github.com/nobonobo/tls/orig/tls"
)

// Conn ...
type Conn struct {
	net.Conn
}

// LocalAddr ...
func (c *Conn) LocalAddr() orig.Addr {
	return c.Conn.LocalAddr()
}

// RemoteAddr ...
func (c *Conn) RemoteAddr() orig.Addr {
	return c.Conn.RemoteAddr()
}

const header = `GET / HTTP/1.0
Host: localhostt:8443

`

func main() {
	log.SetFlags(log.Lshortfile)
	raw, err := net.Dial("tcp", "localhost:8443")
	if err != nil {
		log.Panic(err)
	}
	conn := tls.Client(&Conn{raw}, &tls.Config{
		InsecureSkipVerify: true,
	})
	if err := conn.Handshake(); err != nil {
		log.Panic(err)
	}
	fmt.Fprint(conn, header)
	if _, err := io.Copy(os.Stdout, conn); err != nil {
		if err != io.EOF {
			log.Panic(err)
		}
	}
	conn.Close()
}
