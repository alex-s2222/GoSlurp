package main

import (
	"fmt"
	"sync"
)


func worker(input <-chan int, output chan<- int, wg *sync.WaitGroup){
	defer wg.Done()
	for val := range input{
		output <- val * 2
	}
}


func main(){
	jobs := make(chan int)
	result := make(chan int, 30)

	var wg sync.WaitGroup

	for i := 0;  i < 3; i++{
		wg.Add(1)
		go worker(jobs, result, &wg)
	}
	
	for i := 0; i < 30; i++{
		jobs <- i
	}

	close(jobs)
	wg.Wait()
	close(result)

	for i := range result{
		fmt.Printf("%d\n", i)
	}
}