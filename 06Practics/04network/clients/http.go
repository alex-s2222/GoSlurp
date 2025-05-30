package clients

import (
	"bufio"
	"fmt"
	"net/http"
)


func HTTPClientHeaderGet() {
	client := http.Client{}
	req, _ := http.NewRequest("GET", "http://127.0.0.1:8080/headers", nil)
	req.Header.Add("test", "test")
	
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("error %v", err.Error())
	}
	defer resp.Body.Close()
	
	for name, headers := range resp.Header{
		for _, h := range headers {
			fmt.Printf("H: %v: %v\n", name, h)
		}
	}
	fmt.Println("Response status: ", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil{
		fmt.Printf("Error: %v", err.Error())
	}
}