package main

import "fmt"

var currency = map[string]float64 {
"Bitcoin": .0024,
"Euro": .89,
"Yuan": 6.51,
"Mark": 1.72,
}

func main () {
	fmt.Println(conversion(toconvert,number))
}


func conversion (input string, number float64) float64 {
	for k:= range currency {
		if currency[input] != currency[k] {
			return 0
		}
	}
	return currency[input]*number
}
