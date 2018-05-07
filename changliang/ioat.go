package main

import "fmt"

const a  = iota + 1
const b = iota

func main()  {
	fmt.Println(a)
	fmt.Println(b)
}