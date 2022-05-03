
package main

import (
	"fmt"
)

func getPivotIndex(arr []int, startIndex, endIndex int) int {
	pivotIndex := startIndex - 1
	pivotElement := arr[endIndex]

	for i := startIndex; i < endIndex; i++ {
		if arr[i] < pivotElement {
			pivotIndex++
			arr[i], arr[pivotIndex] = arr[pivotIndex], arr[i]
		}
	}
	arr[pivotIndex+1], arr[endIndex] = arr[endIndex], arr[pivotIndex+1]
	return pivotIndex + 1
}

func quickSort(arr []int, startIndex, endIndex int) {
	if startIndex < endIndex {
		pivotIndex := getPivotIndex(arr,
			startIndex, endIndex)
		quickSort(arr, startIndex, pivotIndex-1)
		quickSort(arr, pivotIndex+1, endIndex)
	}

}

func main() {
	arr := []int{10, 0, 8, 2, 6, 5, 1, 3, 9, 4, 7}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}