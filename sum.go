package main

import "fmt"

func main() {
	fmt.Println(soma(2, 2))
}

func soma(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func times(a int, b int) int {
	return a * b
}

func sumX(a int, b int) int {
	return a + b + a
}