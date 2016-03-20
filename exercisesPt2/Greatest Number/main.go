package main

import "fmt"

func greatest(n...int)int {
	var largest int
	for _,v:= range n {
		if v>largest {
			largest=v
		}
	}
	return largest
}

func main () {
	max := greatest(1,2,3,4,5,6,32,45,2,3,2,3,32,3222222222,22,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2)
	fmt.Println(max)
}
