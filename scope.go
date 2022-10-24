package main

import "fmt"

var packVar = "Package Scope"
// var funcVar = "Func Scope" -occupies memory-
// packVar := "Package Scope" -error-
 var name = "Aysu"
 name,surname := "Aylin","soft"
 fmt.Println(name,surname) //-no error-

var funcVar = "Func(Package) Scope" //print in func main funcVar


func main(){
	// if true {
	// 	var blockVar = "Block Scope"
	// 	fmt.Println(blockVar)
	
	// }
	// fmt.Println(blockVar)  -error-
	var funcvar = "Func Scope"  //occupies space during block operation
	                            //print this funcVar
	funcvar := "Func Scope"


	fmt.Println(funcvar)

	fmt.Println(packVar)

	myFunc()

}

/*func myFunc () {
	fmt.Println(packVar)
	// fmt.Println(funcVar) error

}
*/
