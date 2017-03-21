package main

import (
	"fmt"
)

type Raw []byte

func (r Raw) String() string {
	s := "[ "
	for _, c := range r {
		s += fmt.Sprintf("0x%02X ", c)
	}
	s += "]"
	return s
}

func main() {
	s := "Hello 뽀로로"
	fmt.Println(s)
	fmt.Println(Raw(s))
}

// AFTER_MAIN OMIT
