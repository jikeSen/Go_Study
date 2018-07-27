package main

import "fmt"

func Calls() {
	fmt.Println("我可以被引用")
}

func testCall() {
	fmt.Println("我不能被引用")
}

func main() {
	Calls()
	testCall()
}
