package main

import "fmt"

func main() {

	strArray := []string{"kapil", "hemant", "ram", "shyam", "gopal"}

	searchName(strArray, "ramrgrg")
}

func searchName(array []string, name string) {

	for index, value := range array {

		if value == name {

			fmt.Println(index)
			return

		}
	}
	fmt.Println(fmt.Sprintf("%s name does not exist",name))
}
