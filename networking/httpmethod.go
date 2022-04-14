package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)
var my_url string = "http://18.214.154.73/"

func http_req1() {
	// sending request to above url 
	response,err := http.Get(my_url) 

	if err != nil {
		log.Fatalln(err)
		panic(err)
	}
	fmt.Println("no error found waiting for data response ")
	time.Sleep(2*time.Second)
	fmt.Printf("type of Response are getting %T \n",response)
	defer response.Body.Close() // closing connection -- client responsiblity 

	// reading data 
	data , err := ioutil.ReadAll(response.Body) 
	if err != nil {
		panic(err)
	}
	fmt.Println(response.StatusCode) // http status code 
	fmt.Println(string(data)) // data of response body 
	fmt.Printf("Response lenght : %d \n",response.ContentLength)
}