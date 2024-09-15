package main

import "fmt"

func countChars(input string) map[string]int{
	result := make(map[string]int)
	for _, value := range input{
		char := string(value)

		if _, ok := result[char]; ok{
			result[char]++
		}else{
			result[char] = 1
		}
	}
	return result
}

func main(){
	input := "cъешь ещё этих мягких французских булок, да выпей чаю"
	result := countChars(input)
	for key, value := range(result){
		fmt.Println(key, "=", value)
	}
	
}