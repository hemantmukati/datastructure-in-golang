package main

import "fmt"

func main (){

   arr:=[]int{7,8,5,2,4,6,3}

   for i:=1; i<len(arr)-1;i++{

   for j:=i+1;j>=1;j--{

   	if arr[j-1]>arr[j]{

      arr[j-1],arr[j] = arr[j],arr[j-1]

   	}
   }

   fmt.Println(arr)

   }
}