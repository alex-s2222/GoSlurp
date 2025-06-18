package main

import (
	"fmt"
	"math"
)

type Circle struct{
	Radius int
}

type Rectangle struct{
	Width int
	Height int
}

type Shape interface{ 
	Area() float64
}


func (c *Circle) Area() float64{
	result := math.Pi * math.Pow(float64(c.Radius), 2)
	fmt.Printf("Area Circle %.2f\n", result)
	return result
}

func (r *Rectangle) Area() float64{ 
	result := float64(r.Height * r.Width)
	fmt.Printf("Area Rectangle %.2f\n", result)
	return result
}

func PrintArea(s Shape){
	var size float64

	switch s.(type){
	case *Circle:
		size = s.Area()
	case *Rectangle:
		size = s.Area()
	}
	fmt.Println(size)
}




func main(){
	shapes := []Shape{
				&Circle{Radius: 5},
				&Rectangle{Width: 4, Height: 3},
			}

	for _, shape := range shapes {
		PrintArea(shape)
	}
}