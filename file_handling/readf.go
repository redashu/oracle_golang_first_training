package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func FileRead(filename string) {
	fmt.Println("Reading file using go ..")
	
	// package for reading file 
	data , err := ioutil.ReadFile(filename)
	if err != nil {
		log.Panicf("fail in reading file  : %s\n",err)
	}
	fmt.Printf("my file %s \n",filename)
	fmt.Printf("length of data : %d \n",len(data))
	fmt.Printf("data of file is : %s \n",data)

}