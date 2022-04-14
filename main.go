// package is manadatory
package main

import (
	"fmt"
	"time"
)

// creating structure
type employ struct {
	name , job , country string
	salary int
}
type human struct {
	Firstname string
	lastname string
	age int
}
// declaration of things we are using under code
func main(){

	// for employ 
	var emp1 employ 
	fmt.Println(emp1)	
	emp2 := employ{"ashutoshh","tech","india",30}
	fmt.Println(emp2)

	// for human 
	h1 := human {
		Firstname: "dev",
		lastname: "ckk",
		age: 29,
	}
	fmt.Println(h1)

	// using map method1 
	//   name      key   value 
	mymap1 := map[int]string{
		10: "hello",
		99: "oracle",
		1: "golang",
	}
	fmt.Println(mymap1)
	// adding value 
	mymap1[100] = "okgoogle"
	// delete 
	delete(mymap1,99)
	for i , item := range mymap1 {
		fmt.Println(i,item)
		time.Sleep(1*time.Second)
	}
}


