package main 

import "fmt"

func main (){

    arr:=[]int{3,8,1,6,0,7,2,5,4,9}
    afterSorting:= mergeSort(arr)
    fmt.Println(afterSorting)
}
func mergeSort(elements []int) []int{
    if len(elements)<2{
       return elements
    }
    arrPart1:= mergeSort(elements[: len(elements)/2])
    arrPart2:= mergeSort(elements[len(elements)/2:])

       return merge(arrPart1 , arrPart2)
}
func merge(arr1 []int,arr2 []int)[]int{
    sortArr:=[]int{}
    i:=0
    j:=0 
   for i<len(arr1) && j<len(arr2){
       if arr1[i]<arr2[j]{
        sortArr=append(sortArr,arr1[i])
        i++
       }else{
        sortArr=append(sortArr,arr2[j])
        j++
   }
}
   for ;i<len(arr1);i++{
      sortArr=append(sortArr,arr1[i])

   }
   for ;j<len(arr2);j++{
    sortArr=append(sortArr,arr2[j])
   }
    return sortArr
}