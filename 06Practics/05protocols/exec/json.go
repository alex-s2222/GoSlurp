package exec

import (
	"encoding/json"
	"fmt"
)


type Dimensions struct {
	Height int 
	Width int
}

type Bird struct {
	Species string
	Description string
	Dimensions Dimensions
}

func ParseJson() {
	birdJson := `{"species": "pigeon", "description": "likes to perch on rocks", "dimensions":{"height":24, "width":10}}`
	var bird Bird
	err := json.Unmarshal([]byte(birdJson), &bird)

	if err != nil{
		panic(err)
	}

	fmt.Println(bird)
}

func CreateJson(){
	bird := Bird{
		Species: "Eagly",
		Description: "Cool eagly",
		Dimensions: Dimensions{
			Height:100,
			Width: 50,
		},
	}

	// data, _ := json.Marshal(bird)
	data, _ := json.MarshalIndent(bird, "", "  ")
	fmt.Println(string(data))
}


func ParseJsonArrays() {
	arraysJson := `["one", "two", "theee"]`

	var nums []string
	json.Unmarshal([]byte(arraysJson), &nums)

	fmt.Println(nums)
}