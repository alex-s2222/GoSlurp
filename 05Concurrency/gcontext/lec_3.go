package main

import (
	"context"
	"fmt"
	"time"
)


func runWithContext() {
	// Создаем родительский контекст с отменой
	ctx, cancel := context.WithCancel(context.Background())

	// Количество горутин
	numGoroutines := 5

	// Запускаем несколько горутин
	for i := 1; i <= numGoroutines; i++ {
		go func(id int, ctx context.Context) {
			// Горутина будет работать, пока не получит сигнал на завершение
			for {
				select {
				case <-ctx.Done():
					// Контекст был отменен, завершаем горутину
					fmt.Printf("Горутина %d завершена\n", id)
					return
				default:
					// Симулируем работу горутины
					fmt.Printf("Горутина %d работает...\n", id)
					time.Sleep(1 * time.Second) // Имитация работы
				}
			}
		}(i, ctx)
	}

	// Ждем 5 секунд перед отменой контекста
	time.Sleep(5 * time.Second)
	fmt.Println("Отмена контекста, все горутины должны завершиться")

	// Отменяем контекст, что сигнализирует горутинам завершиться
	cancel()

	// Ждем немного, чтобы убедиться, что все горутины завершились
	time.Sleep(1 * time.Second)
}

func main() {
	runWithContext()
}
