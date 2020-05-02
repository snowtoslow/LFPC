package main

import "fmt"

func main() {
	//myFunc()
	//runeExample()
	//exercise1()
	//exercise2()
	//exercise3()
	exercise5()
}

func myFunc() {

	i := 0
HERE:
	fmt.Print(i)
	i++

	goto HERE


}

func runeExample(){

	var a string
	
	a = "Hello world"

	c := []rune(a)

	fmt.Print(c)
}

func exercise1(){

	for i :=0; i<10; i++  {
		fmt.Print("\n",i)
	}

}

func exercise2(){

	var i int = 0

Loop:
	if i<10 {
		fmt.Printf("%d\n",i)
		i++
		goto  Loop
	}
}

func exercise3()  {

	var array = [10]int{}

	for i :=0;i<10;i++{
		array[i]=i
	}
	fmt.Print(array)

}

func exercise5(){

	var array[100] int

	for i:=0;i<100 ;i++  {
		array[i]=i
	}

	var condition bool

	for _,v := range array{
		condition = false

		if v%3==0 {
			fmt.Print("FIZZ")
			condition = true
		}
		if v % 5 == 0{
			fmt.Print("BUZZ")
			condition = true
		}

		if !condition {
			fmt.Print(v)
		}

	}

}
