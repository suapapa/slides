package main

import (
	"fmt"
)

func main() {
	s := "bounce 바운스"
	for i, c := range s {
		fmt.Printf("i=%v, c=%c t=%T\n", i, c, c)
	}
}
