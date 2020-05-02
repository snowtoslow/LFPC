package main

import (
	"bufio"
	"fmt"
	"os"
)

type mapsWithDuplicate struct {
	symbols string
	values  string
}

func main() {

	var myarray []mapsWithDuplicate
	createMyMap(readLines("varianta21.txt"), &myarray)
	addNewStartingSymbol(&myarray)
	removeEpsilonProduction(&myarray)
	makeUnitSubstitution(&myarray)
	fmt.Print(myarray)

}

func readLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func createMyMap(lines []string, myarray *[]mapsWithDuplicate) *[]mapsWithDuplicate {
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] == '-' {
				*myarray = append(*myarray, mapsWithDuplicate{lines[i][:j], lines[i][j+1:]})
			}
		}
	}
	return myarray
}

//remove epsilon productions
func removeEpsilonProduction(myArray *[]mapsWithDuplicate) *[]mapsWithDuplicate {
	var arrayToRemove []mapsWithDuplicate
	for i := 0; i < len(*myArray); i++ {
		if (*myArray)[i].values == "ε" {
			arrayToRemove = append(arrayToRemove, (*myArray)[i])
			makeSubstitution(myArray, findEpsilonProductionsSymbols(myArray))
			deleteMultipleElements1(myArray, arrayToRemove)
		}
	}
	return myArray
}

func makeSubstitution(myArray *[]mapsWithDuplicate, stringsArray []string) *[]mapsWithDuplicate {
	for i := 0; i < len(*myArray); i++ {
		for j := 0; j < len(stringsArray); j++ {
			/*fmt.Printf("%s--->%s==%t\n",myArray[i].values,stringsArray[j],contains(myArray[i].values,stringsArray[j]))*/
			if contains((*myArray)[i].values, stringsArray[j]) {
				*myArray = append(*myArray, mapsWithDuplicate{(*myArray)[i].symbols, myTrimFunc((*myArray)[i].values, stringsArray[j])})
			}
		}
	}
	return myArray
}

func findEpsilonProductionsSymbols(myArray *[]mapsWithDuplicate) []string {
	var epsilonArray []string
	for i := 0; i < len(*myArray); i++ {
		if (*myArray)[i].values == "ε" {
			epsilonArray = append(epsilonArray, (*myArray)[i].symbols)
		}
	}
	return epsilonArray

}

//unit productions
func makeUnitSubstitution(myarray *[]mapsWithDuplicate) *[]mapsWithDuplicate {
	var arrayToDelete []mapsWithDuplicate
	for i := len(*myarray) - 1; i > 0; i-- {
		if isUnitProduction((*myarray)[i]) {
			arrayToDelete = append(arrayToDelete, mapsWithDuplicate{(*myarray)[i].symbols, (*myarray)[i].values})
			for j := 0; j < len(getProductionBySymbol(*myarray, (*myarray)[i].values)); j++ {
				*myarray = append(*myarray, mapsWithDuplicate{(*myarray)[i].symbols, getProductionBySymbol(*myarray, (*myarray)[i].values)[j]})
			}
			deleteMultipleElements1(myarray, arrayToDelete)
			continue
		}
	}
	return myarray
}

func isUnitProduction(myarray mapsWithDuplicate) bool {
	if (len(myarray.symbols) == 1 && len(myarray.values) == 1) &&
		(isNonTerminal(myarray.symbols[0]) && isNonTerminal(myarray.values[0])) {
		return true
	}

	return false
}

func getProductionBySymbol(myarray []mapsWithDuplicate, symbol string) []string {
	var arrayProduction []string
	for i := 0; i < len(myarray); i++ {
		if myarray[i].symbols == symbol {
			arrayProduction = append(arrayProduction, myarray[i].values)
		}
	}
	return arrayProduction
}

//utils
func deleteMultipleElements1(baseArray *[]mapsWithDuplicate, arrayToDelete []mapsWithDuplicate) *[]mapsWithDuplicate {
	for i := 0; i < len(*baseArray); i++ {
		url := (*baseArray)[i]
		for _, rem := range arrayToDelete {
			if url == rem {
				*baseArray = append((*baseArray)[:i], (*baseArray)[i+1:]...)
				i--
				break
			}
		}
	}
	return baseArray

}

func myTrimFunc(word string, charToTrim string) string {

	myRunes := []rune(word)
	for i := 0; i < len(myRunes); i++ {
		if string(myRunes[i]) == charToTrim && len(string(myRunes)) > 1 {
			myRunes := append(myRunes[:i], myRunes[i+1:]...)
			return string(myRunes)
		} else if word == charToTrim { //be carefull here also can be len(word)==1
			myRunes[i] = 949
			return string(myRunes)
		}
	}
	return "WORD DOES NOT CONTAIN CHARACTER"
}

func contains(word string, char string) bool {
	for i := 0; i < len(word); i++ {
		if string(word[i]) == char {
			return true
		}
	}
	return false
}

func isNonTerminal(char uint8) bool {
	if char >= 65 && char <= 90 {
		return true
	}
	return false
}

func findStartingSymbol(myMap *[]mapsWithDuplicate) string {
	return (*myMap)[0].symbols
}

func addNewStartingSymbol(myMap *[]mapsWithDuplicate) {
	for i := 0; i < len(*myMap); i++ {
		if contains((*myMap)[i].values, findStartingSymbol(myMap)) {
			*myMap = append(*myMap, mapsWithDuplicate{"W", findStartingSymbol(myMap)})
			break
		}
	}
}
