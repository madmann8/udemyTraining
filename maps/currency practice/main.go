package main

import "fmt"

var currency = map[string]float64 {
"Bitcoin": .0024,
"Euro": .89,
"Yuan": 6.51,
"Mark": 1.72,
}

func main () {
	conversion("Bitcoin",5000)
}


func conversion (input string, number float64)  {
	if currency[input] == currency[input] {
		fmt.Println(currency[input] * number, "American Dollars is equvialant to", number, input,"s.")
	}
}