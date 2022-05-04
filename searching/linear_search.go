package main

import "fmt"

func main() {

	arr := []int{2, 5, 8, 0, 7,4,5,6, 3, 1}
	exist := linearSearch(arr, 1)
	fmt.Println("element is exist:", exist)

}
func linearSearch(arr []int, element int) (key int) {

	key = -1
	for i, value := range arr {
		if value == element {
			key = i
		}
	}

	return key

}
