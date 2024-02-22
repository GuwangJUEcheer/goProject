package main

import (
	"fmt"
)

func main() {
	a,b := split(9)
	fmt.Println(a, b)
}

/*
func swap(x, y string) (string, string) {
	return y, x
}
*/
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
