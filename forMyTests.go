package main

import (
	"fmt"
)

func main() {

	myArray := []string{"Vova","mama","MIsha","masha"}
	fmt.Print(arrayContains2(myArray,"costea"))

}

func arrayContains2(stringArray []string, containingString string) bool{
	for i:=0;i<len(stringArray); i++{
		if stringArray[i] == containingString {
			return true
		}
	}
	return false
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
			if isTerminalChar1(string2[j]) && j==len(string2)-1  {
				continue
			}else if !isTerminalChar1(string2[j]) && j==len(string2)-1{
				fmt.Println("HERE")
				return false
			}
		}
	}
	fmt.Print("THERE")
	return true
}

func isTerminalChar1(char uint8) bool{
	if char >=97 && char<=122{
		return true
	}
	return false
}
/*
func removeNonGeneratingSymbolsTest(myarray *[]mapsWithDuplicate) *[]mapsWithDuplicate{//working for my variant
	var arrayOfNonGeneratingSymbols []mapsWithDuplicate
	for i:=0;i<len(*myarray);i++{
		if len((*myarray)[i].values)==1 && isTerminalChar1((*myarray)[i].values[0]) {
			continue
		}else if len((*myarray)[i].values)>1{
			for j:=0;j<len((*myarray)[i].values);j++ {
				if isTerminalChar1((*myarray)[i].values[j])  {
					continue
				}

			}
		} else {
			arrayOfNonGeneratingSymbols = append(arrayOfNonGeneratingSymbols,mapsWithDuplicate{(*myarray)[i].symbols,(*myarray)[i].values})
		}
	}
	log.Print("ARRAY TO REMOVE:",arrayOfNonGeneratingSymbols)
	return deleteMultipleElements1(myarray,arrayOfNonGeneratingSymbols)
}*/