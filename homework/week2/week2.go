package main

import (
	"context"
	"fmt"
)

func mergeStrings(ctx context.Context, letters, numbers string, result chan<- string) {
	defer close(result)

	var merged string
	letterIndex, numberIndex := 0, 0

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context canceled, stopping goroutine")
			return
		default:
			if letterIndex < len(letters) {
				merged += string(letters[letterIndex])
				letterIndex++
			}
			if numberIndex < len(numbers) {
				merged += string(numbers[numberIndex])
				numberIndex++
			}
			result <- merged
			merged = ""
		}
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	result := make(chan string)

	go mergeStrings(ctx, "ABCDEFGHIJKLMNOPQRSTUVWXYZ", "0123456789", result)

	for merged := range result {
		fmt.Println(merged)
	}
}
