package main

import "fmt"

func sendNubers(ch chan int){
	for i := 1; i < 6; i++{
		ch <- i
	}
	close(ch)
}

func main(){
	ch := make(chan int )
	go sendNubers(ch)

	for val := range(ch){
		fmt.Printf("Value: %d\n", val)
	}

}