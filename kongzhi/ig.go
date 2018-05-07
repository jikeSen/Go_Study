package main

import (
	"fmt"
)

func main() {
	forFunc()
}

const FLAG = 1

const (
	sex_boy  = 1
	sex_girl = 0
)

//if 条件和语句
func ifFuc() {
	if FLAG < 0 {
		fmt.Println("小于0")
	} else {
		fmt.Println("大于0")
	}
}

//选择语句
func switchFuc() {
	switch FLAG {
	case sex_boy:
		fmt.Println("boy")
	case sex_girl:
		fmt.Println("girl")
	}
}

func switchFunc() {

	var aa interface{}
	aa = 1
	switch aa.(type) {
	case int:
		println("整型")
	case string:
		println("string")
	default:
		fmt.Println("not match")
	}
}

func selectFunc() {
	select {}
}

//循环语句
func forFunc() {
	/*for i := 1; i < 12; i++ {
		fmt.Println(i)
	}*/

	a := []string{"baba", "apple", "bear"}
	for _, value := range a {
		//fmt.Println(key)
		fmt.Println(value)
	}
}
