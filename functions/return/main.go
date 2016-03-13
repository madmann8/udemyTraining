package main

import "fmt"

func main() {
	fmt.Println(greet ("Luke","Mann"))
}

func greet(fname, lnamm string) string {
	return fmt.Sprint(fname, lnamm)
}
