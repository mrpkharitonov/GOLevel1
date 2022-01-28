package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func CustomSort(us []int) {

	sortUnsort := make([]int, len(us))
	copy(sortUnsort, us)

	n := len(sortUnsort)

	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if sortUnsort[j] < sortUnsort[min] {
				min = j
			}
		}
		sortUnsort[i], sortUnsort[min] = sortUnsort[min], sortUnsort[i]
	}
	fmt.Println(sortUnsort)
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
	CustomSort(unsort)
}
