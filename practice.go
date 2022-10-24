package main

import "fmt"

func main(){

	//first
	/*var studentName string = "John Doe"
	var grade int = 77
	var isPassed bool = true
	*/

	//second
	/*
		var (
		studentName string = "John Doe"
		grade int = 77
		isPassed bool = true
		)
	*/
	
	//third
	/*var studentName  = "John Doe"
	var grade = 77
	var isPassed bool = true
	

	studentName  := "John Doe"
	grade := 77
	isPassed bool := true
	*/

    /*
	var studentName, grade, isPassed ="John Doe",77,true
	studentName, grade, isPassed := "John Doe",77,true

		fmt.Println(studentName)

		fmt.Println(grade)

		fmt.Println(isPassed)
	*/
	
	//declaration
	var studentName string // " "

	//assign 
	var studentName string = "John Doe"

	studentName = "Aysu Deliahmetoğlu"

	//statically typed and dynamically type go and python

	/*
	-python- (dynamically type)
	print(x)
	print(type(x))
	x = "Aysu"
	print(type(x))

	-go- statically type
	var x int  = 100
	fmt.printLn(x)
	var x string  = "aysu"
	fmt.printLn(x) //error no data type change 





	fmt.Println(studentName)



}