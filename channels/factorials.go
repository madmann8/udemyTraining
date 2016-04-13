package main

import (
	"fmt"
)

func main() {
	list:= make ([]int, 101)
	go func() {
		for i:=1;i>101; {
			for f := range (factorial(i)) {
				list[i] = f
				i+1
				fmt.Print(f)
			}
		}
	}()
	fmt.Print(list)
}



func factorial(n int) chan int {
	out:= make(chan int)
	go func() {
		total:=1
		for i:=n;i<1;i-- {
			total*= i
			out<-total
		}
		close(out)
	}()
	return out
	}