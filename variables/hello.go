// package is manadatory 
package main
// declaration of things we are using under code 
import (
	"fmt" 
)
var x1 int = 500 // this is package level scope 
// go compiler will start the code from main function only 
func main() {
	fmt.Println("hello world welcome to Golang  ..")
	fmt.Println(x1)
	// calling function 
	vartest()

}
func vartest() {
	// declair variable method 1 
	var x string
	x = "hello"
	x =  "ok Oralce"
	// print without new line 
	fmt.Printf("%v world  !!  \n",x)
	// second method for var declaration 
	var z int = 100 
	// print with new line 
	fmt.Printf("%v world with new line char..\n",z)
	// final method don't declear and give type 
	a := 300 
	fmt.Printf("%v world with new method and type of data %T..\n",a,a)
	fmt.Println(x1)

}