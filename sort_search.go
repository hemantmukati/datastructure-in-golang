package main

import "fmt"

func main() {

	myArray := []int{4, 2, 6, 8, 0, 5, 1, 7, 3}
	fmt.Println("before sorting:", myArray)
	sortedArr := mergeSort(myArray)
	fmt.Println("after sorting:", sortedArr)

	elementIndex := binarySearch(sortedArr, 0, len(sortedArr)-1, 1)

	fmt.Println("index of element:", elementIndex)

}
func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	arr1 := mergeSort(arr[0 : len(arr)/2])
	arr2 := mergeSort(arr[len(arr)/2:])

	return merge(arr1, arr2)
}
func merge(part1 []int, part2 []int) []int {
	emptyArr := []int{}
	i := 0
	j := 0
	for i < len(part1) && j < len(part2) {
		if part1[i] < part2[j] {
			emptyArr = append(emptyArr, part1[i])
			i++
		} else {
			emptyArr = append(emptyArr, part2[j])
			j++
		}
	}
	for ; i < len(part1); i++ {
		emptyArr = append(emptyArr, part1[i])
	}
	for ; j < len(part2); j++ {
		emptyArr = append(emptyArr, part2[j])
	}
	return emptyArr
}
func binarySearch(sortedArr []int, l, r, element int) int {
	if r >= l {
		m:= l + (r-l)/2
		if sortedArr[m] == element {
			return m
		}
		if sortedArr[m] < element {
			return binarySearch(sortedArr ,m+1,r,element)
		} else {
			return binarySearch(sortedArr ,l,m-1,element)
		}
	}
	return -1
} 
