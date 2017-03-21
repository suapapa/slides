package main

import (
	"fmt"
	"github.com/suapapa/go_hangul"
)

func josa(o string) string {
	// Convert string(utf-8) to []rune(ucs-4)
	r := []rune(string(o))

	// Check if last character has jongsung
	_, _, f := hangul.Split(r[len(r)-1])
	if f == 0 {
		return "야"
	}
	return "아"
}

func printHelloTo(name string) {
	fmt.Printf("%s%s 안녕\n", name, josa(name))
}

// BEFORE_MAIN OMIT

func main() {
	friends := []string{
		"뽀로로", "크롱", "에디",
		"포비", "해리", "패티", "루피",
	}

	for _, f := range friends {
		printHelloTo(f)
	}
}
