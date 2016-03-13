package main

import "fmt"

func main() {
	var largeNumber int
	var smallNumber int
	fmt.Print("Enter a large number: ")
	fmt.Scan(&largeNumber)
	fmt.Print("Enter a smaller number: ")
	fmt.Scan (&smallNumber)
	remainder := largeNumber%smallNumber
	fmt.Println("The remainder of the large number divided by the large number is ", remainder)
}
