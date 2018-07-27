package main

import (
	"fmt"
)

var s = []int{0, 1, 2, 3, 4, 5, 6, 7}

func main() {
	//createSlice()
	//fmt.Println(len(s), cap(s))
	appendSlice()
}

func appendSlice() {
	// 声明一个包含5个元素的整型数组
	//var arr = [5] int{1, 2, 3, 4, 5}
	s = append(s, 8, 9, 10)
	fmt.Println(s)
}

func createSlice() {
	//创建一个长度 和cap 都是4 的字符串切片
	slince_ := make([]string, 4, 5)
	fmt.Println(slince_)
	fmt.Println("slince_", len(slince_), cap(slince_))

	newslice := []string{"jike", "apple", "php", "java", "js"}

	newnewslice := newslice[1:3]
	newnewslice[1] = "apple_Tv"

	newslice = append(newslice, "golang")

	newslice1 := append(newslice, "javascript")
	fmt.Println(newslice)
	fmt.Println(newslice1)
}

func delSlice() {

}
