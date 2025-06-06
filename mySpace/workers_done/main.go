package main

import (
	"fmt"
	"sync"
	"time"
)


func worker(id int, wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Printf("Worker %d started\n", id)
	time.Sleep(time.Millisecond * time.Duration(id) * 100)
	fmt.Printf("Worker %d done\n", id)
}


func main(){
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++{
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("All workers done")
}