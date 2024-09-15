package main

import "fmt"

const (
	PassStatus = "pass"
	FailStatus = "fail"
)

type HealhCheck struct{
	serviceId int 
	status string
}

func GenerateCheck() []HealhCheck{
	result := make([]HealhCheck, 0, 5)

	for i := 0; i <= 5; i++ {
		switch{
			case i % 2 == 0:
				result = append(result, HealhCheck{i, PassStatus})
			default:
				result = append(result, HealhCheck{i, FailStatus})
		}
	}
	return result
}

func main(){
	fmt.Println("Тут будет выведен индефикатор")
	results := GenerateCheck()

	for _, element := range(results){
		fmt.Printf("id:%v status:%v\n", element.serviceId, element.status)
	}

	fmt.Println("________Good_status_______")
	for _, element := range(results){
		id := element.serviceId
		status := element.status

		if status == "pass"{
			fmt.Println("id:", id, "status:", status)
		}
	}
}
