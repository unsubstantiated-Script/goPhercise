package camel_case

import (
	"fmt"
	"unicode"
)

//Run with go run main.go < ./hacker-rank/camel-case/camel.in > ./hacker-rank/camel-case/camel.out

func RollCamelCase() {
	var input string
	fmt.Scanf("%s\n", &input)

	//Set to one because the first word is guaranteed.
	answer := 1

	for _, ch := range input {
		if unicode.IsUpper(ch) {
			answer++
		}
	}

	fmt.Println(answer)
}
