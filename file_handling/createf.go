package main
import (
	"fmt"
	"log"
	"os"
)
func FileCreate(filename , text string) {
	fmt.Println("We are creating file : ")
	// creating file always creating new file 
	file_obj,err  := os.Create(filename)
	if err != nil {
		log.Fatalf("creating file is showing this error %s \n",err)
	}
	defer file_obj.Close() // closing file 

	// want to write some content in the same 
	len1 , err := file_obj.WriteString(text)
	if err != nil {
		log.Fatalf("fail to write data %s \n",err)
	}
	fmt.Printf("name of file : %s \n",file_obj.Name())
	fmt.Printf("leng in bytes %d \n",len1)
}