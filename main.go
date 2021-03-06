package main

import (
	"io"
	"os"
	"fmt"
	"log"
	"net"
	"bufio"
	"syscall"
)

func main() {
	switch os.Args[1] {
	case "write":
		write()
	case "read":
		read()
	case "sysread":
		sysread()
	case "newread":
		newread()
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
	_, _ = stdout.Write([]byte("Writing finished successfully.\n"))
}

func read() {
	g, err := os.Open(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	defer g.Close()

	for {
		z := make([]byte, 8)
		m, err := g.Read(z)
		u := string(z)
		if err == io.EOF {
			break
		}
		fmt.Printf("---\nmessage in file: %v\n", z)
		fmt.Printf("message: %v\n", u)
		fmt.Printf("m:%v\n", m)
		fmt.Printf("z[:%v]: %v\n",m , z[:m])
		fmt.Printf("u[:%v]: %v\n",m , u[:m])
	}

	//stdin  := os.Stdin
	stdout := os.Stdout
	//stderr := os.Stderr
	_, _ = stdout.Write([]byte("Reading finished successfully.\n"))
}

func sysread() {
	g, err := syscall.Open(os.Args[2], syscall.O_RDWR, 0) 
	if err != nil {
		log.Fatal(err)
	}
	defer syscall.Close(g)

	for {
		z := make([]byte, 8)
		m, err := syscall.Read(g, z)
		u := string(z)
		if err != nil {
			log.Fatal(err)
		}
		if m == 0 {
			break
		}
		fmt.Printf("---\nmessage in file: %v\n", z)
		fmt.Printf("message: %v\n", u)
		fmt.Printf("m:%v\n", m)
		fmt.Printf("z[:%v]: %v\n",m , z[:m])
		fmt.Printf("u[:%v]: %v\n",m , u[:m])
	}

	//stdin  := os.Stdin
	//stdout := os.Stdout
	//stderr := os.Stderr
	_, _ = syscall.Write(syscall.Stdout, []byte("Reading finished successfully.\n"))
	//https://xn--go-hh0g6u.com/pkg/syscall/
}

func newread() {
	g, err := os.Open(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	defer g.Close()

	newg := bufio.NewReader(g)
	for {
		z := make([]byte, 8)
		m, err := newg.Read(z)
		u := string(z)
		if err == io.EOF {
			break
		}
		fmt.Printf("---\nmessage in file: %v\n", z)
		fmt.Printf("message: %v\n", u)
		fmt.Printf("m:%v\n", m)
		fmt.Printf("z[:%v]: %v\n",m , z[:m])
		fmt.Printf("u[:%v]: %v\n",m , u[:m])
	}

	//stdin  := os.Stdin
	stdout := os.Stdout
	//stderr := os.Stderr
	_, _ = stdout.Write([]byte("Reading finished successfully.\n"))
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
