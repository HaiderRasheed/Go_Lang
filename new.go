package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	var prefix, suffix string = "i", "n"
	var inner = "a"

	fmt.Println("Please, enter a string: ")
	fmt.Scan(&str)

	str = strings.ToLower(str)
	condPre := strings.HasPrefix(str, prefix)
	condSuf := strings.HasSuffix(str, suffix)
	condCon := strings.Contains(str, inner)

	if condPre && condSuf && condCon {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
