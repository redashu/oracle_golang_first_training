package main
import (
	"fmt"
	"time"
)

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