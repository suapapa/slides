package main

import (
	"fmt"
)

func printHelloTo(name string) {
	fmt.Printf("%s야 안녕\n", name)
}

func main() {
	friends := []string{
		"뽀로로", "크롱", "에디",
		"포비", "해리", "패티", "루피",
	}

	for _, f := range friends {
		printHelloTo(f)
	}
}
