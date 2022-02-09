package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func permutationMethodSort(us []int) []int {

	n := len(us)

	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if us[j] < us[min] {
				min = j
			}
		}
		us[i], us[min] = us[min], us[i]
	}
	return us
}

func main() {
	fmt.Print("введите целые числа, для завершения ввода нажмите два раза ввод \n")
	var unsort []int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("введите число: ")
		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {
			fmt.Println(text)
			textN, err := strconv.Atoi(text)
			if err != nil {
				log.Fatal(err)
			}
			unsort = append(unsort, textN)
		} else {
			break
		}

	}
	fmt.Println("\n\nнеотсортированный")
	fmt.Println(unsort)

	fmt.Println("\nотсортированный")
	fmt.Println(permutationMethodSort(unsort))
}
