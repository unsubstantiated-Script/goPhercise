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
	fmt.Println(s)
}
