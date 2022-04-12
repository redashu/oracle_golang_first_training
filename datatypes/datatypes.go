package main
import (
	"fmt"
	"math"
	"time"
) 

func dtype() {
	var x float64 
	fmt.Println("enter a number : ")
	fmt.Scan(&x)
	fmt.Println("square root is getting calculated for : ",x)
	time.Sleep(2*time.Second)
	// math package using 
	result := math.Sqrt(x) // flloat64 type value is supported 
	fmt.Println("calculation is ",result)

}