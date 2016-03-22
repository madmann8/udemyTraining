package main

import "fmt"

func main() {
	myGreeting := map[int]string{
		1: "Good Morning",
		2: "Hello",
	}
	myGreeting[3] = "hola"
	fmt.Println("test")
	fmt.Println(myGreeting)
}