package camel_case

import (
	"fmt"
	"unicode"
)

//Run with go run main.go < ./hacker-rank/camel-case/camel.in > ./hacker-rank/camel-case/camel.out

func RollCamelCase() {
	var input string
	fmt.Scanf("%s\n", &input)
	//fmt.Println("Input is:", input)
	answer := 0

	for _, ch := range input {
		if unicode.IsUpper(ch) {
			answer++
		}
	}

	fmt.Println(answer)
}
