package main

import (
	"fmt"
	"time"
)

func main() {
	for _, v := range []int{1, 2, 3, 4, 5, 6, 7} {
		go func(value int) {
			printNumber(&value)
		}(v)
	}

	time.Sleep(2 * time.Second)
}

func printNumber(v *int) {
	fmt.Println(*v)
}
