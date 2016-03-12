package main

import "fmt"

myName := "LM"

func main() {
	switch len(myName) {
	case 1: {
		fmt.Println("The Length is 1")
	}
	case 2: {
		fmt.Println("The length is 2")
	}

	}
}
