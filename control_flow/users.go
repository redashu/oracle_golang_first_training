package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)
func main() {
	fmt.Println("creating users in Linux host ")
	user_names := os.Args[1:]
	fmt.Println(user_names)
	fmt.Printf("type of data we give %T \n",user_names)
	// usage range with slice to iterate using for loop 
	for _,itme := range user_names {
		//fmt.Printf(" %v times executing loop and value is %v \n",i,itme)
		fmt.Println(itme)
		time.Sleep(1*time.Second)
		cmd,err  := exec.Command(itme).Output()
		if err != nil {
			panic(err) // for terminating program 
		}
		fmt.Printf("%s\n",cmd)
	}
}