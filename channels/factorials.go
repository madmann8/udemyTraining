package main

import "fmt"

func main() {
	n:= fact(20)
	for r:= range(n) {
		fmt.Println(r)
	}
}

func fact (n int)chan int {
	out:=make(chan int)
	go func(){
		total:= 1
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()
	return out
}
