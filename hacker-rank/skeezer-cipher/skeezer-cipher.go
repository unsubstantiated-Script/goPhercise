package skeezer_cipher

import (
	"fmt"
)

func RollSkeezerCipher() {
	var length, delta int
	var input string
	fmt.Scanf("%d\n", &length)
	fmt.Scanf("%s\n", &input)
	fmt.Scanf("%d\n", &delta)

	var ret []rune
	for _, ch := range input {
		ret = append(ret, cipher(ch, delta))
	}

	fmt.Println(string(ret))

}

func cipher(r rune, delta int) rune {
	//Between A and Z runes
	if r >= 'A' && r <= 'Z' {
		return rotate(r, 'A', delta)
	}

	if r >= 'a' && r <= 'z' {
		return rotate(r, 'a', delta)
	}

	return r
}

func rotate(r rune, base, delta int) rune {
	//Rebasing as if A is zero
	tmp := int(r) - base
	tmp = (tmp + delta) % 26
	return rune(tmp + base)
}
