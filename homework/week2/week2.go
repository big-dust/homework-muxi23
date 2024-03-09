package main

import (
	"fmt"
)

func main() {
	str2 := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	str1 := "0123456789"

	ch := make(chan rune)

	go func() {
		for _, char := range str1 {
			ch <- char
		}
	}()

	go func() {
		for _, char := range str2 {
			ch <- char
		}
	}()

	for i := 0; i < len(str1)+len(str2); i++ {
		fmt.Printf("%c", <-ch)
	}
}
