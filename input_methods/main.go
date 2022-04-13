package main

import (
	"bufio"
	"fmt"
	"os"
)

// go compile this is entrypoint
func main() {
	/*
	var x1 string 
	fmt.Println("enter any data : ")
	fmt.Scan(&x1)
	fmt.Println("you have entered : ",x1)
	*/
	// using bufio 
	reader := bufio.NewReader(os.Stdin) // for os STDIn is keyboard 
	fmt.Println("enter anything : ")
	input,_ := reader.ReadString('\n') // keep reading input till new line is not found 
	fmt.Println("thanks for your input : ",input)

	// inline input in golang 
	cmd_input1 := os.Args
	cmd_input2 := os.Args[1:] // except first element 
	fmt.Println(cmd_input1)
	fmt.Println(cmd_input2)
	fmt.Printf("type of data %T \n",cmd_input2)

}