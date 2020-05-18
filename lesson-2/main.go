package main

import "fmt"

func main() {
	fmt.Println(calc(5, 5))

	fmt.Println("shift", multiplication(1, 1)>>1)

	fmt.Println(multiplication(5, 5))

	_, x := getTwo()
	println(x)
}

func calc(x int, y int) int {
	return x + y
}

func getTwo() (int, int) {
	return 1, 2
}

func multiplication(x int, y int) (result int) {
	result = x * y

	return
}
