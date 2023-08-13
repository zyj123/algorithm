package tcp

import (
	"fmt"
	"net"
)

func server() {
	addr, err := net.ResolveTCPAddr("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	for {
		fmt.Println("pre accept...")
		conn, err := listener.Accept()
		fmt.Println("after accept...")
		if err != nil {
			fmt.Println("accept err:", err)
		}

		go handler(conn)
	}
}

func handler(conn net.Conn) {

}

func client() {
	addr, err := net.ResolveTCPAddr("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	_, err = net.DialTCP("tcp", nil, addr)
	if err != nil {
		panic(err)
	}

	closed := make(chan int)
	for {
		select {
		case <-closed:

		}
	}

}
