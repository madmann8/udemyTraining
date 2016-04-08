package main

import (
"time"
"fmt"
"sync"
)

var wg sync.WaitGroup


func alphabet() {
	for i := 65; i < 91; i++ {
		fmt.Print(string(i))
		time.Sleep(5*time.Millisecond)
	}
	wg.Done()
}

func numbers() {
	for i:=0; i<26;i++ {
		fmt.Print(i)
		time.Sleep(1*time.Millisecond)
	}
	wg.Done()
}
func main() {
	wg.Add(2)
	go alphabet()
	go numbers()
	wg.Wait()
}