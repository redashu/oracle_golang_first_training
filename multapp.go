package main

import (
	"fmt"
	"time"
)



func task111(msg string) {
	for i := 0;  i < 10 ; i++ {
		time.Sleep(10*time.Microsecond) // 
		fmt.Println("here is message : ",msg)
	}
}