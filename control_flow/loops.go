package main

import (
	"fmt"
	"time"
)

func loops() {
	fmt.Println("Golang is only Having For in looping concept !! ")
	// method 1 
	fmt.Println("enter your Name : ")
	var x string
	fmt.Scan(&x)
	for i := 0; i < 4 ; i++ {
		fmt.Printf("Hey %v welcome Golang loop %v times \n",x,i)
		time.Sleep(1*time.Second)
	}
	// standard to other languages 
	var y int = 0 
	for ( y < 3 ) {
		fmt.Println("hello")
		y++
	}
	// using string method 
	var st string = "oracle"
	for i, j:= range st {
		fmt.Printf("index number is %v is %v \n",j,i)
	} 
	// using array 
	arr2 := [3]string{"hello","oracle","India"} 
	for i, item := range arr2 {
		fmt.Printf("date is %v and iteration is %v \n",i,item)
	}
}