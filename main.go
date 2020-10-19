package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/nobonobo/tinygo-tls/orig/crypto/tls"
	orig "github.com/nobonobo/tinygo-tls/orig/net"
)

const header = `GET / HTTP/1.0
Host: localhostt:8443

`

type stdConn struct {
}

// Read ...
func (c *stdConn) Read(b []byte) (n int, err error) { return os.Stdin.Read(b) }

// Write ...
func (c *stdConn) Write(b []byte) (n int, err error) { return os.Stdout.Write(b) }

// Close closes the connection.
func (c *stdConn) Close() error { return nil }

// LocalAddr ...
func (c *stdConn) LocalAddr() orig.Addr { return nil }

// RemoteAddr ...
func (c *stdConn) RemoteAddr() orig.Addr { return nil }

// SetDeadline ...
func (c *stdConn) SetDeadline(t time.Time) error { return nil }

// SetReadDeadline ...
func (c *stdConn) SetReadDeadline(t time.Time) error { return nil }

// SetWriteDeadline ...
func (c *stdConn) SetWriteDeadline(t time.Time) error { return nil }

func main() {
	log.SetFlags(log.Lshortfile)
	raw := &stdConn{}
	conn := tls.Client(raw, &tls.Config{
		InsecureSkipVerify: true,
	})
	if err := conn.Handshake(); err != nil {
		log.Panic(err)
	}
	fmt.Fprint(conn, header)
	if _, err := io.Copy(os.Stderr, conn); err != nil {
		if err != io.EOF {
			log.Panic(err)
		}
	}
	conn.Close()
}
