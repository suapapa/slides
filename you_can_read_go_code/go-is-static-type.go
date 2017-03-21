package main

import (
	"fmt"
)

func main() {
	// ok will be bool type
	ok := false

	// this is not valid!
	ok = "true"

	fmt.Println("ok =", ok)
}
