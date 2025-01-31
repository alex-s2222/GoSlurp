package clients

import (
	"bufio"
	"fmt"
	"net"
)

func UDPClient() {
	p := make([]byte, 1024)
	conn, err := net.Dial("udp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println("Dial error: ", err.Error())
		return
	}
	fmt.Fprintf(conn, "HI UDP Server, How are you?")
	_, err = bufio.NewReader(conn).Read(p)

	if err == nil {
		fmt.Println(string(p))
	} else {
		fmt.Println("Error:", err.Error())
	}
	conn.Close()
}