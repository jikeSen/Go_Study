package main

import (
	"fmt"
	"reflect"
)

var a int
var b int = 67

var (
	qq  string = "jike"
	zhf int16  = 666
)
//var aa, bb, cc int = 2, 3, 4

func init() {

}

func main() {
	//只能用在函数体内
	j, k, l := 6, 3.44, 8

	a = 89
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(j)
	fmt.Println(reflect.TypeOf(k))
	fmt.Println(l)
	fmt.Println(reflect.TypeOf(qq))
	fmt.Println(zhf)
}
