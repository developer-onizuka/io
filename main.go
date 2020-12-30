package main

import (
	"io"
	"os"
	"fmt"
	"log"
	"net"
)

func main() {
	switch os.Args[1] {
	case "write":
		write()
	case "read":
		read()
	case "copy":
		copy()
	case "net":
		netdial()
	default:
		panic("bad command")
	}
}

// https://golang.org/src/io/io.go
func write() {
	x := "hello world\n"
	fmt.Printf("message: %v", x)

	f, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	y := []byte(x)
	fmt.Printf("message in byte: %v\n", y)
	n, _ := f.Write(y)
	fmt.Printf("n:%v\n", n)

	//stdin  := os.Stdin
	stdout := os.Stdout
	//stderr := os.Stderr
	n, _ = stdout.Write([]byte("Writing finished successfully.\n"))
}

func read() {
	g, err := os.Open(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	defer g.Close()

	z := make([]byte, 128)
	m, _ := g.Read(z)
	fmt.Printf("message in file: %v\n", z)
	u := string(z)
	fmt.Printf("message: %v\n", u)
	fmt.Printf("m:%v\n", m)
	fmt.Printf("z[0:%v]: %v\n",m , z[0:m])
	fmt.Printf("u[0:%v]: %v\n",m , u[0:m])

	//stdin  := os.Stdin
	stdout := os.Stdout
	//stderr := os.Stderr
	m, _ = stdout.Write([]byte("Reading finished successfully.\n"))
}

func copy() {
	src, err := os.Open(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	dst, err := os.Create("copy.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		log.Fatal(err)
	}
}

// http://golang.jp/pkg/net
// https://kobatako.hatenablog.com/entry/2017/11/07/001719
func netdial() {
	url := "yahoo.co.jp"

	// type struct
	ipaddr, err := net.ResolveIPAddr("ip", "www." + url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("address: %v\n", ipaddr.IP)

	// type interface
	conn, err := net.Dial("tcp", url + ":80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	//action := "GET / HTTP/1.0\r\n\r\n"
	action := "HEAD / HTTP/1.0\r\n\r\n"
	conn.Write([]byte(action))

	//stdin  := os.Stdin
	stdout := os.Stdout
	//stderr := os.Stderr
	io.Copy(stdout, conn)
}
