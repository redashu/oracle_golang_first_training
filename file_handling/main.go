package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main(){
	// user input for file name  
	fmt.Println("enter file name : ")
	var filename string
	fmt.Scan(&filename)
	// user input for data 
	fmt.Println("enter data : ")
	inputReader := bufio.NewReader(os.Stdin)
	input,_ := inputReader.ReadString('\n')
	// calling function 
	FileCreate(filename,input) 
	fmt.Println("plz wait file is in reading phase ...")
	fmt.Println("_______________________________")
	time.Sleep(3*time.Second)
	
	// reading file fucntion 
	FileRead(filename)
}
