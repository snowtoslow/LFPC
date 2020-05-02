package utils

import (
	structure "awesomeProject/mystruct"
	"bufio"
	"os"
)

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

func DeleteMultipleElements1(baseArray *[]structure.MapsWithDuplicate, arrayToDelete []structure.MapsWithDuplicate) *[]structure.MapsWithDuplicate {
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

func MyTrimFunc(word string, charToTrim string) string {

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

func addNewStartingSymbol(myMap *[]structure.MapsWithDuplicate) {
	for i := 0; i < len(*myMap); i++ {
		if Contains((*myMap)[i].Values, findStartingSymbol(myMap)) {
			*myMap = append(*myMap, structure.MapsWithDuplicate{Symbols: "W", Values: findStartingSymbol(myMap)})
			break
		}
	}
}

func findStartingSymbol(myMap *[]structure.MapsWithDuplicate) string {
	return (*myMap)[0].Symbols
}

func IsNonTerminal(char uint8) bool {
	if char >= 65 && char <= 90 {
		return true
	}
	return false
}

func Contains(word string, char string) bool {
	for i := 0; i < len(word); i++ {
		if string(word[i]) == char {
			return true
		}
	}
	return false
}
