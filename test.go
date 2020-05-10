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
	removeNonGeneratingSymbols(&myarray)
	//removeUnreachebleSymbols(&myarray)
	/*getNewProductionsFromStartingSymbolProductions(getStaringSymbolsProduction(&myarray),&myarray)
	fmt.Println(compareProductions(getMyArrayProductions(&myarray),getStaringSymbolsProduction(&myarray)))*/
	removeUnreachebleSymbols1(&myarray)
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
			/*fmt.Printf("%s--->%s==%t\n",myArray[i].values,stringsArray[j],stringContains(myArray[i].values,stringsArray[j]))*/
			if stringContains((*myArray)[i].values, stringsArray[j]) {
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

func stringContains(word string, char string) bool {
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
		if stringContains((*myMap)[i].values, findStartingSymbol(myMap)) {
			*myMap = append(*myMap, mapsWithDuplicate{"W", findStartingSymbol(myMap)})
			break
		}
	}
}

func hasTerminal(stringArray []string) bool{
	for i:=0;i<len(stringArray);i++ {
		if isTerminalString(stringArray[i]) {
			return true
		}
	}
	return false
}

func isTerminalString(string2 string) bool{
	if len(string2)==1 && isTerminalChar(string2[0]){
		return true
	}else if len(string2)==1 && !isTerminalChar(string2[0]){
		return false
	} else if len(string2)>1 {
		for j:=0;j<len(string2);j++ {
			if isTerminalChar(string2[j]) && j==len(string2)-1  {
				continue
			}else if !isTerminalChar(string2[j]) && j==len(string2)-1{
				return false
			}
		}
	}
	return true
}

func isTerminalChar(char uint8) bool{
	if char >=97 && char<=122{
		return true
	}
	return false
}

func arrayContains(stringArray []string, containingString string) bool{
	for i:=0;i<len(stringArray); {
		if stringArray[i] == containingString {
			return true
		}
	}
	return false
}


// non generating symbols
func removeNonGeneratingSymbols(myarray *[]mapsWithDuplicate) *[]mapsWithDuplicate{
	var arrayOfNonGeneratingSymbols []mapsWithDuplicate
	for i:=0;i<len(*myarray);i++ {
		if hasTerminal(getProductionBySymbol(/*fmt.Println(getProductionBySymbol(*myarray,findStartingSymbol(myarray)))*/*myarray,(*myarray)[i].symbols)) {
			continue
		}
		arrayOfNonGeneratingSymbols = append(arrayOfNonGeneratingSymbols,mapsWithDuplicate{(*myarray)[i].symbols,(*myarray)[i].values})
	}
	return deleteMultipleElements1(myarray,arrayOfNonGeneratingSymbols)
}

//unreacheble symbols
func removeUnreachebleSymbols1(myarray *[]mapsWithDuplicate) *[]mapsWithDuplicate{
	var arrayToDelete []mapsWithDuplicate
	for k := range compareProductions(getMyArrayProductions(myarray),getStaringSymbolsProduction(myarray)){
		if compareProductions(getMyArrayProductions(myarray),getStaringSymbolsProduction(myarray))[k]==false {
			for _,n := range getProductionBySymbol(*myarray,k){
				arrayToDelete = append(arrayToDelete,mapsWithDuplicate{k,n})
			}
		}
	}
	return deleteMultipleElements1(myarray,arrayToDelete)
}


func compareProductions(myMap map[string]bool,startingSymbolMap map[string]bool) map[string]bool{
	for k := range myMap{
		for i :=range startingSymbolMap{
			if k==i {
				myMap[k]=true
			}
		}
	}
	return myMap
}

func getMyArrayProductions(myarray *[]mapsWithDuplicate) map[string]bool{
	myMap := make(map[string]bool)
	for i:=0;i<len(*myarray);i++ {
		myMap[(*myarray)[i].symbols] = false
	}
	return myMap
}


func getStaringSymbolsProduction(myarray *[]mapsWithDuplicate) map[string]bool{
	myMap := make(map[string]bool)
	for i:=0;i<len(getProductionBySymbol(*myarray,"W"));i++ {
		for _,v := range getProductionBySymbol(*myarray,"W")[i]{
			if isNonTerminal(uint8(v)) {
				myMap["W"] = true
				myMap[string(uint8(v))] = true
			}
		}
	}
	return myMap
}

