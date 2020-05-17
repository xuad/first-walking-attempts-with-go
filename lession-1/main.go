package main

import "fmt"

func main() {
	printString()
	calculate()
	calcComplex()
}

func printString() {
	var name string = "test"

	fmt.Println(name)
}

func calculate() {
	var numberOne int = 5
	numberTwo := 5
	var z float64 = 3.66 + 5

	fmt.Println(numberOne+numberTwo, z)
}

func calcComplex() {
	var c complex128 = 5i + 3
	var z = 2i
	fmt.Println(c + z)
}
