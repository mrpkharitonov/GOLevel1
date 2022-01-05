package main

import (
	"fmt"
	"os"
)

func main() {
	var x, y, res, x0, y0 float32
	var op string

	fmt.Println("введите значение 1")
	fmt.Scanln(&x)
	fmt.Println("введите оператор")
	fmt.Scanln(&op)
	fmt.Println("введите значение 2")
	fmt.Scanln(&y)

	if x == x0 {
		if y == y0 {
			if len(op) < 1 {
				fmt.Println("не ввели оператор")
				os.Exit(1)
				//почему-то не срабатывает, выпадает с ошибкой в проверке в default, где проверяется доступность функции
			}
			fmt.Println("не ввели вторую цифру")
			os.Exit(1)
		}
		fmt.Println("не ввели первую цифру")
		os.Exit(1)
	}

	switch op {
	case "+":
		res = x + y
	case "-":
		res = x - y
	case "*":
		res = x * y
	case "/":
		if y == 0 {
			fmt.Println("делить на ноль нельзя")
			os.Exit(1)
		} else {
			res = x / y
		}
	case "%":
		res = (x / 100) * y

		// остальные функции не сообразил как выразить без интерфейса просто вводом в консоли
	default:
		fmt.Println("операция недоступна")
		fmt.Println(len(op))
		os.Exit(1)
	}
	fmt.Printf("Результат выполнения операции: %f\n", res)

	var n, mark int
	fmt.Println("введите число N:")
	fmt.Scanln(&n)
	for i := 2; i < n; i++ { //тк простое число это то которе делится на 1 и на само себя перебирать с 0 нет необходимости
		mark = 0
		for j := 2; j < n; j++ {
			if i%j == 0 {
				mark++
			}
		}
		if mark == 1 {
			fmt.Println(i)
		}
	}

}
