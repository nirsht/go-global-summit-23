package main

import "fmt"

type Gopher struct {
	Name string
}

func main() {
	var prints []func()
	for _, v := range []int{1, 2, 3} {
		prints = append(prints, func() {
			fmt.Println(v)
		})
	}

	for _, print := range prints {
		print()
	}
}
