package main

import (
	"io"
	"log"
	"net"
	"os"
	"os/exec"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8443")
	if err != nil {
		log.Panic(err)
	}
	log.Println("exec:", os.Args[1:])
	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	stdin, err := cmd.StdinPipe()
	stdout, err := cmd.StdoutPipe()
	stderr, err := cmd.StderrPipe()
	go io.Copy(stdin, conn)
	go io.Copy(conn, stdout)
	go io.Copy(os.Stderr, stderr)
	if err := cmd.Run(); err != nil {
		log.Panic(err)
	}
}
