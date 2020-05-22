package main

import "fmt"

func main() {
	var emptyArray []int
	fmt.Println(len(emptyArray))

	emptyArray = append(emptyArray, 4)
	fmt.Println(len(emptyArray))

	filledArray := make([]int, 2)
	fmt.Println(filledArray)

	filledArray2 := []int{1, 2}
	fmt.Println(filledArray2)

	var multiArray [2][3]int
	fmt.Println(multiArray)

	initMap := make(map[string]int)
	initMap["blub"] = 123
	fmt.Println(initMap)

	h, a := "ha", "ha"
	fmt.Println(h, a)
}
