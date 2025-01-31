package servers

import (
	"fmt"
	"net"
)

func UDPServer() {
	p := make([]byte, 2048)
	addr := net.UDPAddr{
		Port: 1234,
		IP: net.ParseIP("127.0.0.1"),
	}

	ser, err := net.ListenUDP("udp", &addr)
	if err != nil{
		fmt.Println("Some error: ", err.Error())
		return
	}
	fmt.Println("Listening on localhost:1234")

	for {
		_, remodeAddr, err := ser.ReadFromUDP(p)
		fmt.Printf("Read a message from %v, %v\n", remodeAddr, string(p))
		if err != nil{
			fmt.Println("some error", err.Error())
			continue
		}
		go sendResponse(ser, remodeAddr)
	}
}


func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
	_, err := conn.WriteToUDP([]byte("From server: Hello I got your message "), addr)
	if err != nil {
		fmt.Println("Couldn't send response: ", err.Error())
	}
}
