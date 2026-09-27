package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i == 2 {
			fmt.Println("Yoooo 2 occurred")
			continue
		} else if i == 8 {
			fmt.Println("It's Joeover")
			break
		}
		fmt.Println("Number ", i)
	}

	words := []string{"orange", "banana", "apple"}
	for fruits, fruitsName := range words {
		fmt.Println(fruits, fruitsName)
	}
}
