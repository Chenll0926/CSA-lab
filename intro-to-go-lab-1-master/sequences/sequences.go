package main

import "fmt"

func addOne(a int) int {
	return a + 1
}

func square(a int) int {
	return a * a
}

func double(slice []int) []int {
	slice = append(slice, slice...)
	return slice
}

func mapSlice(f func(a int) int, slice []int) {
	for i, v := range slice {
		slice[i] = f(v)
	}
}

func mapArray(f func(a int) int, array [3]int) {
	for i, v := range array {
		array[i] = f(v)
	}
}

func main() {
	intSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(intSlice)
	mapSlice(addOne, intSlice)
	fmt.Println(intSlice)

	//intArray := [3]int{1, 2, 3}
	//fmt.Println(intArray)
	//mapArray(addOne, intArray)
	//fmt.Println(intArray)

	newSlice := intSlice[1:3]
	fmt.Println(newSlice)
	mapSlice(square, newSlice)
	fmt.Println(newSlice)
	fmt.Println(intSlice)

	intSlice = double(intSlice)
	fmt.Println(intSlice)

}
