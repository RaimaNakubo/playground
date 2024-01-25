package main

import (
	"fmt"
)

// function spit splits the given number in two numbers in a proportional rate
func split(sum int) (x, y int) {
	x = sum * 1 / 3
	y = sum - x
	return
}

func main() {
	fmt.Println(split(100))
}
