package main

import "fmt"

func main() {
	for i, n := range []uint64{1, 2, 3, 5} {
		fmt.Println(i, n)
	}
}
