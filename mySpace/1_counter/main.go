package main

import (
	"fmt"
	"sync"
)

func worker(counter *int, wg *sync.WaitGroup, mu *sync.Mutex){
	defer wg.Done()
	
	for i := 0; i < 100; i++{
		mu.Lock()
		*counter++
 		mu.Unlock()
	}
}

func main(){
	var wg sync.WaitGroup
	var mu sync.Mutex
	counter := 0

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go worker(&counter, &wg, &mu)
	}
	wg.Wait()
	fmt.Printf("Result count after gorut: %d", counter)
}
