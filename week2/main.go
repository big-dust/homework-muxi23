package main

import (
	"fmt"
	"time"
)

func PrintString(s string) {
	for i := 0; i < len(s); i++ {
		time.Sleep(50 * time.Millisecond)
		fmt.Printf("%c", s[i])
	}
}
func main() {
	A := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	B := "0123456789"
	go PrintString(A)
	go PrintString(B)
	time.Sleep(2 * time.Second)
}
