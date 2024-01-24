package main

import (
	"fmt"
)

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	firstWord, secondWord := swap("Fuck", "You")
	fmt.Println(firstWord, secondWord)
}
