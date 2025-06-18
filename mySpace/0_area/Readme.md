## Интерфейсы и структуры

### Задача: “Фигуры и площадь”

* Создай интерфейс Shape с методом Area() float64.
* Создай структуры Circle и Rectangle, реализующие этот интерфейс.
* Напиши функцию, принимающую Shape и печатающую площадь.

#### Пример использования:

```go
shapes := []Shape{
    Circle{Radius: 5},
    Rectangle{Width: 4, Height: 3},
}
for _, shape := range shapes {
    PrintArea(shape)
}
```