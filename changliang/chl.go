package main

import "fmt"

//常量不支持派生类

//显式定义
const NAME string = "jikesne"

const SEX  = "my name is jikesen"

const (
	cat = "dog"
	mox = "monkey"
)

const apple,banana  string = "xiangjianbg","axoajosxaj"

const aa,bb = 1,"asdphaasdasnfoiasfphwfob"

func test()  {
	const APP = "iphone"
}

func main()  {
	/*fmt.Println(NAME)
	fmt.Println(SEX)
	fmt.Println(cat)
	fmt.Println(mox)
	fmt.Println(apple)
	fmt.Println(banana)
	fmt.Println(reflect.TypeOf(aa))
	fmt.Println(bb)*/
	fmt.Println(icc(apple))
}

func icc(name string)  string{
	return name
}