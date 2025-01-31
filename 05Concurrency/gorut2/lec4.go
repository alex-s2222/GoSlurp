package main

import (
	"fmt"
	"sync"
)

// Функция для добавления нечетных чисел в массив
func addOddNumbers(a []int, c *[]int, oddChan, evenChan chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, num := range a {
		<-oddChan           
		*c = append(*c, num)
		evenChan <- struct{}{} 
	}
}

// Функция для добавления четных чисел в массив
func addEvenNumbers(b []int, c *[]int, oddChan, evenChan chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, num := range b {
		<-evenChan          
		*c = append(*c, num) 
		oddChan <- struct{}{} 
	}
}

func main() {
	// Исходные массивы
	a := []int{1, 3, 5, 7, 9}  
	b := []int{2, 4, 6, 8, 10} 

	var c []int

	oddChan := make(chan struct{}, 1)
	evenChan := make(chan struct{}, 1)

	var wg sync.WaitGroup

	// Добавляем в WaitGroup две горутины
	wg.Add(2)

	go addOddNumbers(a, &c, oddChan, evenChan, &wg)
	go addEvenNumbers(b, &c, oddChan, evenChan, &wg)

	// Начинаем с сигнала для добавления нечетного числа
	oddChan <- struct{}{}

	// Ждем завершения всех горутин
	wg.Wait()

	fmt.Println("Результирующий массив:", c)
}
