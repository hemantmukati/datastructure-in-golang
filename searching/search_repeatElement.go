package main 

import "fmt"

func main (){

	arr:= []string{"a","b","a","c","a","b"}
     
    newMap:=make(map[string]int)

    for i:=0;i<len(arr);i++{
    
     value:=newMap[arr[i]]
     if value==0{   
         newMap[arr[i]] = 1
      } else{
        newMap[arr[i]] = newMap[arr[i]]+1
      }  
    }

    fmt.Println(newMap)
}