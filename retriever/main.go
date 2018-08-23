package main

import (
	"Go_Study/retriever/baidu"
	"fmt"
	"Go_Study/retriever/google"
)

type Retrivevr interface {
	Get(url string) string
}

func download(r Retrivevr) string {
	return r.Get("http://www.jd.com")
}

func download2(r Retrivevr) string {
	return r.Get("http://www.jd.com")
}

func main() {
	var r Retrivevr
	r = baidu.Retrivevr{Contents: "this is a fake baidu retrivevr"}
	fmt.Printf("%T  %v\n", r, r)
	r = &google.Retriever{
		UserAgent:"mozilla/5.0",
	}
	fmt.Printf("%T  %v\n", r, r)
	//sstr := download(r)
	//fmt.Println(sstr)
}
