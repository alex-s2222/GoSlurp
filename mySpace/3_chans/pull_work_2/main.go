package main

import (
	"fmt"
	"sync"
)


func main() {
	jobs := make(chan int)    
	results := make(chan int) 
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range jobs {
				results <- job * 2
			}
		}()
	}

	go func() {
		for i := 1; i <= 10; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	// Чтение результатов
	for res := range results {
		fmt.Println(res)
	}
}