package main

import (
	"fmt"
	"sync"
	"time"
	"errors"
)

func foo() error {
	var wg sync.WaitGroup

	done := make(chan int)

	numGorutines := 5 

	wg.Add(numGorutines)

	for i := 0; i < numGorutines; i++{
		go func(id int){
			defer wg.Done()
			time.Sleep(time.Duration(id) * time.Second)
			fmt.Printf("Горутина %d завершила работу \n", id)
		}(i)
	}

	go func(){
		wg.Wait()
		close(done)
	}()

	select{
	case <- done:
		fmt.Println("Все горутины успешно завершились")
		return nil
	case <-time.After(2 * time.Second):
		// Время ожидания вышло
		return errors.New("превышено время ожидания: более 2 секунд")
	}

}

func main(){
	if err := foo(); err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Задача выполнена успешно")
	}
}