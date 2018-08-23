package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	const filanem = "aa.in"
	const length = 30
	var f *os.File

	w := bufio.NewWriter(f)

	res, err := w.ReadFrom(f)
	

	if (err != nil) {
		panic("sss")
	}
	fmt.Printf("写入 %d 个字节n", res)
	w.Flush()

}
