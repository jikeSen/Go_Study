package main

import (
	"fmt"
	"~/Go/"
)

func main() {
	p := pipline.ArraySo(3, 4, 2, 45, -1)
	for {
		if num, ok := <-p; ok {
			fmt.Println(num)
		}
	}
}
