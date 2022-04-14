// package is manadatory
package main

// declaration of things we are using under code
import (
	"fmt"
	"os"
	"runtime"
	"time"
)
var x int
// writing init funtion 
func init() {
	fmt.Println("i am the one who comes even b4 main function ")
	x = 1000 
}
// go compiler will start the code from main function only 
func main() {
	fmt.Println("hello world welcome to Golang  ..")
	fmt.Println("using variable ",x)
	
}

func init() {
	fmt.Println(" Order does not matter for me bcz my name is init ")
	// host os check to run the entire code 
	if runtime.GOOS == "linux" {
		fmt.Println("yes then code will be executed ")
	} else {
		fmt.Println("you are running your code on wrong platform ")
		time.Sleep(2*time.Second)
		fmt.Println("Your current platform is : ",runtime.GOOS)
		os.Exit(1) // closing code using exit function 
	}
}


