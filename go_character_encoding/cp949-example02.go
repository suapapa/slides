package main

import (
	"fmt"
	"github.com/suapapa/go_hangul/encoding/cp949"
	"os"
)

func main() {
	f, _ := os.Open("euckr.xml")
	defer f.Close()

	r, _ := cp949.NewReader(f)

	b := make([]byte, 10*1024)
	c, _ := r.Read(b)

	fmt.Println(c)
	fmt.Printf("%s", b[:c])
}
