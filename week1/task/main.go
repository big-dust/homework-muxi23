package main

import "fmt"

func add[T int | float64 | string](a []T) []T {
	var b []T = a
	return b
}

func expand[T int | float64 | string](tar *[]T, exp []T) {
	*tar = append(*tar, exp...)
}

func del[T int | float64 | string](a *[]T) {
	*a = (*a)[:0]
}

func output[T int | float64 | string](a []T) {
	if len(a) == 0 {
		fmt.Println("empty")
		return
	}

	for k, v := range a {
		fmt.Printf("[%v]:%v ", k, v)
	}
	fmt.Printf("\n")
}

func main() {
	tmp := add([]int{1, 2, 3})
	output(tmp)
	expand(&tmp, []int{4, 5})
	output(tmp)
	del(&tmp)
	output(tmp)
}
