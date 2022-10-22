package main

import "fmt"

func main(){
	/*var (
		name string = "aysu"
		age int = 23
		number  = 257
		isMarried bool = false
		weight float32 = 50.6
	)
	*/
	// var name, age, number, isMarried, weight = "Aysu", 23, 257, false, 67.8
    // name, age, number, isMarried, weight := "Aysu", 23, 257, false, 67.8  -short-

	// fmt.Println(name)
	// fmt.Println(age)	
	// fmt.Println(number)	
	// fmt.Println(isMarried)	
	// fmt.Println(weight)	 

	var name string
	fmt.Println(name)  // string --> " "

	var age int 
	fmt.Println(age) // numeric --> 0

	var isMarried bool 
	fmt.Println(isMarried) // bool --> false
}
