package main

import (
	"fmt"
)

func getChars() []byte {
	return []byte{
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'_', '@',
	}
}

func findPasswordNaive(password string) string {
	word := ""
	for k := 0; k < len(getChars()); k++ {
		word = string(getChars()[k])
		if word == password {
			return word
		}
		for i := range getChars() {
			word = string(getChars()[k]) + string(getChars()[i])
			if word == password {
				return word
			}
			for j := range getChars() {
				word = string(getChars()[k]) + string(getChars()[i]) + string(getChars()[j])
				if word == password {
					return word
				}
				for l := range getChars() {
					word = string(getChars()[k]) + string(getChars()[i]) + string(getChars()[j]) + string(getChars()[l])
					if word == password {
						return word
					}
					for m := range getChars() {
						word = string(getChars()[k]) + string(getChars()[i]) + string(getChars()[j]) + string(getChars()[l]) + string(getChars()[m])
						if word == password {
							return word
						}
					}
				}
			}
		}
	}

	return ""
}

func main() {
	fmt.Println("Password found: ", findPasswordNaive("@kAl1"))
}
