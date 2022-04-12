package main
import (
	"fmt"
	"time"
)

func ifcase1() {
	fmt.Printf("Press 1 to check current time in your OS : \n")
	fmt.Printf("Press 2 to create a folder  in your OS under current folder : \n")
	fmt.Printf("Press 3 to connect docker engine local/remote : \n")
	var user_input int 
	fmt.Scan(&user_input)

	if user_input == 1 {
		fmt.Println("current time is : ",time.Now().UTC())
	} else if user_input == 2  {
		// taking user input 

	} else {
		fmt.Println("no match found ")
	}
}