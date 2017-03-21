package main

import (
	"fmt"
)

func main() {
	s := "미지의세계"

	fmt.Printf("type = %T, len = %d\n", s[0], len(s))

	b := []byte(s)
	fmt.Printf("type = %T, len = %d\n", b[0], len(b))

	r := []rune(s)
	fmt.Printf("type = %T, len = %d\n", r[0], len(r))
}
