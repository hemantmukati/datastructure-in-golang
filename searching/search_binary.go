// package main
// import "fmt"

// func binarySearch(checknum int, arr []int) bool {

// 	low := 0
// 	high := len(arr) - 1

// 	for low <= high{
// 		median := (low + high) / 2

// 		if arr[median] < checknum {
// 			low = median + 1
// 		}else{
// 			high = median - 1
// 		}
// 	}

// 	if low == len(arr) || arr[low] != checknum {
// 		return false
// 	}

// 	return true
// }
// func main(){
// 	arr := []int{1,2, 9, 20, 31, 45, 63, 70, 100}
// 	fmt.Println(binarySearch(63, arr))
// }

package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println("Exist at postion : ", binarySearch(arr, 0, len(arr)-1, 5))

}

func binarySearch(arr []int, l, r, element int)int {

	if l <= r {
		m := l + (r-l)/2

		if arr[m] == element {
			return m
		}

		if arr[m] > element {
			return  binarySearch(arr, l, m-1, element)
		} else {
			return binarySearch(arr, m+1, r, element)
		}

	}
	return -1
}
