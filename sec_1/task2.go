package main

import (
	"fmt"
	"math/rand/v2"
)

func checkArr(arr [9]int) bool{
	hashMap := make(map[int]int)
	for idx, value := range(arr){
		if _, ok := hashMap[value]; ok{
			return true
		}
		hashMap[value] = idx
	}
	return false
}

func createArr() [9]int{
	var result [9]int
	for i := 0; i < 9; i++{	
		num := rand.IntN(10) 
		result[i] = num
	}
	return result
}

func main(){
	fmt.Println("___True____")
	true_arr := [9]int{1,2,3,4,5,6,7,8,8}
	fmt.Println(checkArr(true_arr))

	fmt.Println("___False____")
	false_arr := [9]int{1,2,3,4,5,6,7,8,9}
	fmt.Println(checkArr(false_arr))

	fmt.Println("___Random____")
	rand_arr := createArr()
	fmt.Println(rand_arr, "->", checkArr(rand_arr))
}