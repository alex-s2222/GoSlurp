package main

import (
	"fmt"
)


func generator(in chan int){
	for i := 1; i<10; i++{
		in <- i
	}
	close(in)
}

func squarer(in chan int, out chan int){
	for val := range in{
		out <- val * val
	}
	close(out)
}

func main(){
	input := make(chan int)
	output := make(chan int)

	go generator(input)
	go squarer(input, output)

	for num := range output{
		fmt.Printf("Value: %d \n", num)
	}

}