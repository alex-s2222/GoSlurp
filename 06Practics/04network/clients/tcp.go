package clients

import (
	"bufio"
	"fmt"
	"net"
)

func TCPClient() {
	p := make([]byte, 1024)
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Dial error: ", err.Error())
	}
	fmt.Fprintf(conn, "HI TCP Server, How are you?")
	_, err = bufio.NewReader(conn).Read(p)

	if err == nil {
		fmt.Println(string(p))
	} else {
		fmt.Println("Error:", err.Error())
	}
	conn.Close()
}