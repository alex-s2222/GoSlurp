package main

import (
	"fmt"
	"math"
)


type Area interface{
	Area() error 
}

type Type interface{
	Type() string
}

type Rectangle struct{
	name string
	side int 
}

type Circle struct{
	name string
	radius int
}

func (r *Rectangle) Area() error{
	result := math.Pow(float64(r.side), 2)
	fmt.Printf("У %v Площадь: %v\n", r.Type(), result)
	return nil
} 

func (r *Rectangle) Type() string{
	return r.name
}

func (c *Circle) Area() error{
	name := GetType(c)
	result := math.Pi * math.Pow(float64(c.radius),2)

	fmt.Printf("У %v Площадь: %v\n", name, result)
	return nil 
}

func (c *Circle) Type() string{
	return c.name
}

func CalculateArea(a Area){
	// err := a.Area()
	// _, ok := a.(*Circle)	
	switch a.(type){
		case *Circle:
			a.Area()
		case *Rectangle:
			a.Area()
	}
}

func GetType(t Type) string{
	return t.Type()
}


func main(){
	rectangle := &Rectangle{"прямоугольник",3}
	CalculateArea(rectangle)
	circle := &Circle{"круг",5}
	CalculateArea(circle)
}