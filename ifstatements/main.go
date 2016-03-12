package main

import "fmt"

func main() {
	for i:=0; i<1000000; i++ {
		if i%42==1 {
			fmt.Println(i)
		}
	}
}