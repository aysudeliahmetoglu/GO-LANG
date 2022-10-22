package main

import "fmt"

func main(){
   var name string = "aysu"
   var age int = 23
   var number  = 257
   var isMarried bool = false
   var weight float32 = 50.6

   name = "alice"
   
   fmt.Println(name)
   fmt.Println(age)
   fmt.Println(isMarried)

   fmt.Printf("%T\n",name)
   fmt.Printf("%T\n",age)
   fmt.Printf("%T\n",isMarried)
   fmt.Printf("%T\n",weight)
   fmt.Printf("%T\n",number)

   /*numeric types:
   int uint8 -> unsigned (0 to 255)
	float -> float32
    complex 
    byte
    rune

   string types:

   boolean tpes:
   true-false

   composite types:
   slice-struct-pointer-function -interface-map-channel-array */

}