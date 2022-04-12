package main
// declaration of things we are using under code 
import (
	"fmt" 
	"time"
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
	//vartest()
	operators1()

}
func operators1() {
	var x int
	var y int
	fmt.Println("Enter first number: ")
	fmt.Scan(&x) // taking input from user 
	fmt.Println("enter second number : ")
	fmt.Scan(&y)
	// equal operator 
	fmt.Println("plz wait just getting started with Operatros ..")
	time.Sleep(5*time.Second)
	output1 :=  x == y 
	fmt.Println("both values are same ",output1)
	output2:=  x > y 
	fmt.Println("both first number is greator  ",output2)
	sum := x + y 
	fmt.Printf("addition of values %v \n",sum)
	subs :=  x - y 
	fmt.Println("substraction of values are ",subs)

}
