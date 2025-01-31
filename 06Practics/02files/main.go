package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(e error){
	if e != nil {
		panic(e)
	}
}

// func read(filepath string){
// 	dat, err := os.ReadFile(filepath)
// 	check(err)
// 	fmt.Print(string(dat))
// }

func read(filepath string) {
	f, err := os.Open(filepath)
	check(err)
	defer f.Close()

	b1 := make([]byte, 7)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	b2 := make([]byte, 4)
	n2, err := f.Read(b2)
	check(err)	
	fmt.Printf("%v\n", string(b2[:n2]))

	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @%d %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, 0)
	check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(15)
	check(err)
	fmt.Printf("%v\n", string(b4))
}


func write(filepath string) {
	f, err := os.Create(filepath)
	check(err)
	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes \n", n2)

	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	_ = f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("byffered\n")
	check(err)
	fmt.Printf("wrote %d bytes \n", n4)

	w.Flush()
}

func main(){
	read("text.txt")
	write("write.txt")
}