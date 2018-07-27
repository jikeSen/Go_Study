package main

import "fmt"

func main() {
	createMap()
}

func createMap() {
	ages := make(map[string]string)
	ages["name"] = "jikeSen"
	ages["sex"] = "男"
	ages["money"] = "asd"
	ages["test"] = "asd"

	testName,ok := ages["test"]
	aaa,ok := ages["aaa"]
	fmt.Println(testName,ok)
	fmt.Println(aaa,ok)
}
