package main

import (
	"fmt"
	"time"
)

var fibonacciCache = make(map[uint]uint)

func fibbonachi(n uint) uint {
	var value uint

	if fibonacciCache[n] != 0 {
		value = fibonacciCache[n]
	} else if n < 2 {
		value = n
	} else {
		value = fibbonachi(n-1) + fibbonachi(n-2)
	}
	fibonacciCache[n] = value
	return value
}

func main() {
	val := 56
	fmt.Println(time.Now())
	fmt.Println(fibbonachi(uint(val)))
	fmt.Println(time.Now())
	fmt.Println(fibbonachi(uint(val)))
	fmt.Println(time.Now())
}
