package utils

import "fmt"

func Input(question string) string {
	var input string

	fmt.Print(question)
	fmt.Scanln(&input)
	return input
}
