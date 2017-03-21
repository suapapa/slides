package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("euckr.xml")
	defer f.Close()

	b := make([]byte, 10*1024)
	c, _ := f.Read(b)

	fmt.Println(c)
	fmt.Printf("%s", b[:c])
}
