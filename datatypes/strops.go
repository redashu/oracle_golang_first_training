package main
import (
	"fmt"
	"time"
) 

func stops() {
	var st1 string = "Hello"
	var st2 string 
	fmt.Println("enter any string : ")
	fmt.Scan(&st2)
	time.Sleep(1*time.Second)
	fmt.Println("You have entered : ",st2)
	// adding string 
	fmt.Println(st1+" "+st2)
	fmt.Println(st1,st2) // this is automatically adding space 
	// calculate length 
	fmt.Println("length of given string ",st2," is",len(st2))
	fmt.Println(st2[1]) // ascii value of char
	fmt.Println(string(st2[1])) // value of index 
	fmt.Println(string(st2[1:4])) // range of string char


}