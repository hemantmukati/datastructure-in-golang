/*package main 

import "fmt"


func main (){

    myArray:=[]int{1,2,3,4,5,6,7,8,9}

    findElement:=binarySearch(myArray,0,len(myArray)-1,6)
    fmt.Println("positin of element is:",findElement)
}
   func binarySearch(arr[]int ,l,r,element int)int{
     
      if r>=l{

       m:= l+(r-l)/2
      if arr[m]==element{
          return m
     }
      if arr[m]<element{
        return binarySearch(arr ,m+1,r,element)

      }else{

      	return binarySearch(arr,l,m-1,element)
      }


      }
      return -1

   }

*/
   package main 

   import "fmt"

   func main (){

       arr:=[]int {1,2,3,4,5,6,7,8,9}

    fmt.Println(binarySearch(arr , 0 ,len(arr)-1 ,3))
   } 
      func binarySearch(arr []int ,l ,r ,element int) int{

      for r>=l{
      
      m:= l+(r-l)/2

      if arr[m]==element{
         return m
      }
      if arr[m]< element{
      	 l= m+1 
      }else{
      	 r=m-1
      }
   }
    return -1
}






