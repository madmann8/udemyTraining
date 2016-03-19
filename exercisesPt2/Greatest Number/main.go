package main

import "fmt"

func greatest(n...int)int {
	largest:=0
	l:= len(n)
	for i:=0;i>l;i++{
		if i == 0 {
			largest = n(i)
			continue
		}
		if i> largest {
			largest= n(i)
			continue
		}
		return n
	}

}

func main () {
	fmt.Println(grettest(1,2,3,4,5))
}
