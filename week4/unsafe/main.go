package main

import (
	"fmt"
	"unsafe"
)

func SliceToString(x []byte) string {
	res := *(*string)(unsafe.Pointer(&x))
	return res
}
func main() {
	x := []byte("一个人可以被毁灭，但绝不会被打败")
	s := SliceToString(x)
	fmt.Println(s)
}
