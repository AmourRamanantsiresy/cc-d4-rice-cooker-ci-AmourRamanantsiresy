package utils

import (
	"fmt"
	"os"
	"strings"
)

func Show(path string, toReplace []string, toReplaceLength int) {
	filePath := "./assets/" + path
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	replacePattern := "xxx"
	repeatPattern := "═"

	fileContentAsString := string(fileContent)

	if toReplace != nil && len(toReplace) > 0 {
		for _, value := range toReplace {
			fileContentAsString = strings.Replace(fileContentAsString, replacePattern, value, 1)
		}
		if toReplaceLength != 0 {
			value := strings.Repeat(repeatPattern, toReplaceLength)
			fileContentAsString = strings.Replace(fileContentAsString, replacePattern, value, 1)
		} else {
			value := strings.Repeat(replacePattern, len(toReplace[0]))
			fileContentAsString = strings.Replace(fileContentAsString, replacePattern, value, 1)
		}
	}

	fmt.Println(fileContentAsString)
}
