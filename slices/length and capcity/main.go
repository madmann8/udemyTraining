package main

import "fmt"

func main() {
	foo:= make([]int,5000)
	for i:=0;i<5000;i++{
		foo[i]=i
		fmt.Println(i," - ",foo[i])
	}

}
