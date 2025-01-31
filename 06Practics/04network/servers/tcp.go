package servers

import (
	"fmt"
	"net"
	"os"
)

func TCPServer() {
	l, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	defer l.Close()
	fmt.Println("Listening on localhost:1234")
	
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err.Error())
			os.Exit(1)
		}
		
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	buf := make([]byte, 1024)

	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading: ", err)
	}

	fmt.Println(reqLen)
	fmt.Println(string(buf))

	bytesWritten, err := conn.Write([]byte("Messege received"))
	if err != nil{
		fmt.Println("Error writing: ", err.Error())
	}
	fmt.Println("Bytes written ", bytesWritten)

	conn.Close()
}