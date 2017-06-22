package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan bool)

	go func() {
		time.Sleep(time.Second * 1)
		fmt.Println("One Second")
		time.Sleep(time.Second * 1)
		fmt.Println("One Second")
		c <- true
	}()

	d := <-c
	fmt.Println(d)

}
