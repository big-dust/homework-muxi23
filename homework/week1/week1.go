package main

import (
	"fmt"
	"strings"
)

type Builder[T any] struct {
	data []T
}

func NewBuilder[T any]() *Builder[T] {
	return &Builder[T]{}
}

func (b *Builder[T]) Add(item T) {
	b.data = append(b.data, item)
}

func (b *Builder[T]) String() string {
	var result strings.Builder
	for _, item := range b.data {
		fmt.Fprintf(&result, "%v", item)
	}
	return result.String()
}

func main() {
	intBuilder := NewBuilder[int]()
	intBuilder.Add(1)
	intBuilder.Add(2)
	intBuilder.Add(3)
	fmt.Println(intBuilder.String())
	strBuilder := NewBuilder[string]()
	strBuilder.Add("Hello, ")
	strBuilder.Add("World!")
	fmt.Println(strBuilder.String())
}
