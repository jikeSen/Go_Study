package main

import (
	"fmt"
)

const NAME  =  "jikeSen"

var APP = "iphone"

//声明借口
type stu interface {

}

//结构声明
type Iler struct {

}

//自定义函数
func test()  {
	fmt.Print("hello world")
}

//主函数
func main()  {
	test()
}
