package main

import (
	"fmt"
)

func getCb() func(int, int) int {
	cbAdd := func(a, b int) int {
		return a + b
	}
	return cbAdd
}

func calc(cb func(int, int) int, a, b int) int {
	return cb(a, b)
}

func main() {
	sum := calc(getCb(), 3, 4)
	fmt.Println(sum)
}

//END OMIT
