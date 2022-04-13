package main

import (
	"context"
	"fmt"
	"time"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)
func docker1() {
	ctx1 := context.Background() // creating docker context in golang 
	cli,err := client.NewClientWithOpts(client.FromEnv,client.WithAPIVersionNegotiation())
	// above line of code to initiate connection with docker using apiNeg , reading date from ENV
	if err != nil {
		panic(err) // if error found then panic will terminate programe 
	}
	// docker images kind of operations 
	myimages,err := cli.ImageList(ctx1,types.ImageListOptions{})
	if err != nil {
		panic(err) // if error found then panic will terminate programe 
	}
	for _,image := range myimages {
		fmt.Println(image.ID)
		time.Sleep(2*time.Second)
		fmt.Println(image.RepoTags)
	}
}
