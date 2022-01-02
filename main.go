package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float32
	var c, d float64
	var e, f, g, h int

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

	fmt.Print("введите 3х значное число ")
	fmt.Scanln(&e)

	h = e % 10
	g = (e - h) % 100
	f = (e - g - h)

	fmt.Printf("сотни: %d\n", f/100)
	fmt.Printf("десятки: %d\n", g/10)
	fmt.Printf("единицы: %d\n", h)
}
