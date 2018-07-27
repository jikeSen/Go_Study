package main

import "fmt"

func main() {
	/*arr := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr)
	for _, v := range arr {
		fmt.Println(v)
	}*/
	arays()
}

func arays() {
	//var arr1 [5]int
	//arr2 := [3]int{1, 3, 4}
	arr3 := [...]int{2, 4, 6, 8, 10} //让编译器计算int 数量
	//var grid [2][4]int               //二维数组
	instanceArr(arr3[:])
	/*fmt.Println(arr1, arr2, arr3, grid)

	for _, k := range arr3 { //循环
		fmt.Println(k)
	}*/
}

func instanceArr(arr []int) {
	for k, v := range arr {
		fmt.Println(k, v)
	}
}
