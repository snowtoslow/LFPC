package main

import (
	"fmt"
)

func main() {

	fmt.Print(isTerminalString1("aaV"))



}

func containsTest (word string,char uint8) bool{
	for i:=0;i<len(word) ;i++  {
		if word[i] == char {
			return true
		}
	}
	return false
}

func isTerminalString1(string2 string) bool{
	if len(string2)==1 && isTerminalChar1(string2[0]){
		return true
	}else if len(string2)==1 && !isTerminalChar1(string2[0]){
		return false
	} else if len(string2)>1 {
		for j:=0;j<len(string2);j++ {
			if isTerminalChar1(string2[j]) {
				continue
			}else if !isTerminalChar1(string2[j]){
				fmt.Println("HERE")
				return false
			}
		}
	}
	fmt.Print("THERE")
	return false
}

func isTerminalChar1(char uint8) bool{
	if char >=97 && char<=122{
		return true
	}
	return false
}