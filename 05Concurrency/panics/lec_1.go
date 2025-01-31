package main

import (
	"errors"
	"fmt"
)

func foo() (err error){ 
	defer func() {
		if r := recover(); r != nil {
			// Преобразуем панику в ошибку
			err = errors.New(fmt.Sprintf("произошла паника: %v", r))
		}
	}()
	// return errors.New(fmt.Sprintf("произошла паника: %v", "fefe"))
	panic("паника СОС!!!")
	
}

func main(){
	fmt.Println("Начинаем")

	if err := foo(); err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("ХУХ Всё прошло успешно")
	}
}