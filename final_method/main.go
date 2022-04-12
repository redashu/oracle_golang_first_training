// package is manadatory 
package main
// declaration of things we are using under code 
import (
	"fmt" 
)
// package level variable declaration 
var (
	x1 int = 500 
	name string = "ashutoshh"
)
// go compiler will start the code from main function only 
func main() {
	fmt.Println("hello world welcome to Golang  ..")
	fmt.Println(x1)
	fmt.Printf("hello My name is : %v \n",name)
	// defining one constant which will be having a fixed value 
	const total_exam = 50 
	fmt.Println(total_exam)
	// calling function 
	vartest()
	operators1()

}


