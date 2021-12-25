package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float32
	var c, d float64

	fmt.Print("введите первую сторону прямоугольника ")
	fmt.Scanln(&a)

	fmt.Print("введите вторую сторону прямоугольника: ")
	fmt.Scanln(&b)

	fmt.Printf("площадь прямоугольника: %f\n", a*b)

	fmt.Print("введите площадь круга ")
	fmt.Scanln(&c)

	d = math.Sqrt(c / math.Pi)

	fmt.Printf("радиус круга: %f\n", d)
	fmt.Printf("диаметр круга: %f\n", d*2)

	//todo: не успел 3.	С клавиатуры вводится трехзначное число. Выведите цифры, соответствующие количество сотен, десятков и единиц  в этом числе. прошу отсрочку

}
