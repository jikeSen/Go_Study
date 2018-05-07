package main

import (
	"sort"
	"fmt"
)

func main() {
	//定义int类型数组
	a := []int{1, 14, 3, 34, -3}
	sort.Ints(a)
	/*for _, v := range a {
		fmt.Println(v)
	}*/

	fmt.Println(eval(3, 4, "-"))
}

//归并排序
func sortB() {

}

func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("not support op" + op)

	}
}
