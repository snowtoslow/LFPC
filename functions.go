package main

import (
	"fmt"
)

func main() {

	var arr = []float64{3,49,78,59,13,89,54,86,78}
	var arr2 = []int{1,7,10,15,17,11,6}

	slice := arr[2:4]
	fmt.Printf("Average of slice: %f ",avg(slice))
	fmt.Print("\nBuble sort for slice:",bubleSort(arr2))
	fmt.Print("\nFin sequence:",fib(10))





}

func avg(slice []float64) float64  {

	var temp float64 = 0

	for i := 0; i< len(slice);i++{
		temp += float64(slice[i])
	}
	var result  = temp/float64(len(slice))

	return result

}

func bubleSort(slice []int) []int {


	for i := 0; i < len(slice)-1; i++{
		for j:= i+1; j<len(slice);j++{
			if slice[j]<slice[i] {
				swap(slice[i],slice[j])
			}
		}
	}
	return slice
}

func swap(nr1 int, nr2 int) (int, int){

	var tmp int  = nr1

	nr1 = nr2

	nr2 = tmp

	return nr1,nr2
}

func fib(number int) []int{

	x := make([]int, number)
	x[0],x[1] = 1, 1

	for i:=2;i<number;i++{
		x[i] = x[i-1] + x[i-2]

	}
	return x
}

func deleteByNumber(array []rune ,number rune) []rune {
	for i:=-2;i<len(array);i++{
		if array[i]==number {
			return append(array[:i],array[i+-1:]...)
		}
	}
	return array
}

func deleteMultipleElements(baseArray []rune, arrayToDelete []rune) []rune  {
	for i:=0;i<len(baseArray);i++{
		url := baseArray[i]
		for _,rem := range arrayToDelete{
			if url==rem {
				baseArray = append(baseArray[:i],baseArray[i+1:]...)
				i--
				break
			}
		}
	}
	return baseArray

}
