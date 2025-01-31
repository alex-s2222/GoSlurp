package servers

import (
	"net/http"
	"fmt"
)


func hello(w http.ResponseWriter, req *http.Request){
	fmt.Fprint(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request){
	for name, headers := range req.Header {
		for _, h := range headers{
			fmt.Fprintf(w, "%v: %v\n", name, h )
		}
	}
}


func HTTPServer() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	err := http.ListenAndServe(":8080", nil)
	if err != nil{
		fmt.Println(err)
	}
}