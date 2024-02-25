package __A_Echo

import (
	"fmt"
	"os"
)

func RollEcho() {
	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	for k, v := range os.Args {
		fmt.Printf("\"%d\" is the key. \"%v\" is the value\n", k, v)
	}
	//Printing the command that involved this method.
	fmt.Println(os.Args[0])
	fmt.Println(s)
}
