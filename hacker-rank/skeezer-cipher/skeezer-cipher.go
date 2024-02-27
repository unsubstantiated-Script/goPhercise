package skeezer_cipher

import "fmt"

func RollSkeezerCipher() {
	var length, delta int
	var input string
	fmt.Scanf("%d\n", &length)
	fmt.Scanf("%s\n", &input)
	fmt.Scanf("%d\n", &delta)

	//fmt.Printf("length: %d\n", length)
	//fmt.Printf("input: %s\n", input)
	//fmt.Printf("delta: %d\n", delta)

	//Converts this to a slice of runes
	alphabet := []rune("abcdefghijklmnopqrstuvwxyz")

	newRune := rotate('z', 2, alphabet)
	fmt.Println(string(newRune))
}

func rotate(s rune, delta int, key []rune) rune {
	//We don't know where we are yet
	idx := -1
	//looping through the slice of char runes
	for i, r := range key {
		//If the value is equal to the rune, we're going to set that index.
		if r == s {
			idx = i
			break
		}
	}
	if idx < 0 {
		panic("idx < 0")
	}

	for i := 0; i < delta; i++ {
		idx++
		if idx >= len(key) {
			idx = 0
		}
	}
	return key[idx]
}
