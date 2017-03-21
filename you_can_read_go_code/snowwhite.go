package main

import (
	"fmt"
	"time"
)

func main() {
	var sumOfHeight int
	for i := 0; i < 7; i++ {
		sumOfHeight += 160
	}
	fmt.Println("sumOfHeight =", sumOfHeight)

	// Will wait forever
	kissed := false
	for !kissed {
		fmt.Println("Prince not comming yet?")
		time.Sleep(1 * time.Second)
	}
	// END OMIT
}
