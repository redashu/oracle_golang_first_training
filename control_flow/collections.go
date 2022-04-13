package main

import "fmt"

func array1() {
	fmt.Println("welcome to array in Golang !!")
	// decleare array 
	// method1 
	var arr1[3] int 
	arr1[0] = 10 
	arr1[1]	= 89
	arr1[2]	= 101
	fmt.Println(arr1)
	fmt.Printf("type of array is %T \n",arr1)
	// method 2 
	arr2 := [3]string{"hello","oracle","India"} 
	fmt.Printf("type of arr2 is %T \n",arr2)
	fmt.Println(arr2[1:])
	fmt.Println(arr2[1],arr2[2])
	fmt.Println("_______###@@________")
	// slice 
	// method 1 
	var slice_1 = []int{3,6,9} 
	// method 2
	slice_2 := []string{"hello","oracle","India","Golang"}
	fmt.Printf("value is %v and type is %T \n",slice_1,slice_1)
	fmt.Printf("value is %v and type is %T \n",slice_2,slice_2)
	// from array to slice 
	var slice_3 = arr2[1:]
	fmt.Println(slice_3)
}