package main

import "fmt"

func main() {

	myArray := []string{"A","B","A","B","A","B"}

	for i:=0;i<len(myArray);i++ {
		if myArray[i]=="B" {
			myArray[i] = myArray[len(myArray)-1]
			myArray[len(myArray)-1] = ""
			myArray = myArray[:len(myArray)-1]
		}
	}
	fmt.Println(myArray)
	fmt.Println(len(myArray))


}

func containsTest (word string,char uint8) bool{
	for i:=0;i<len(word) ;i++  {
		if word[i] == char {
			return true
		}
	}
	return false
}
