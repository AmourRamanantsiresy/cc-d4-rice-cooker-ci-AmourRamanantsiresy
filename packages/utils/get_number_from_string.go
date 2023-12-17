package utils

import (
	"fmt"
	"strconv"
)

func GetNumberFromString(number string) int {
	res, err := strconv.Atoi(number)

	if err != nil {
		fmt.Print(err)
		return 0
	}

	return res
}
