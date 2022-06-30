package main

import (
	"fmt"
)

func main() {

	c := make(chan string)

	go func(input chan string) {
		input <- "hello1"
		input <- "hello2"
		input <- "hello3"
		input <- "hello4"
	}(c)

	for greeting := range c {
		fmt.Println(greeting)
	}
	//fmt.Println(greeting)
	//go HelloWorld()
}

func HelloWorld() {
	fmt.Println("Hello world!")
}
