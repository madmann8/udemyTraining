package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	s1 rand.Source
	r1 *rand.Rand
)

type Payload struct {
	Timestamp time.Time
	Work      int
}

func init() {
	s1 = rand.NewSource(time.Now().UnixNano())
	r1 = rand.New(s1)
}

func main() {

	work := make(chan Payload)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			randDelay()
			work <- Payload{
				Timestamp: time.Now(),
				Work:      i,
			}
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			randDelay()
			work <- Payload{
				Timestamp: time.Now(),
				Work:      i,
			}
		}
		done <- true
	}()

	go func() {

		<-done

		<-done

		close(work)
	}()

	for n := range work {
		elapsed := time.Since(n.Timestamp)
		fmt.Printf("%d finished in %q time\n", n.Work, elapsed)
	}
}

func randDelay() {

	randomValue := r1.Intn(1000)
	time.Sleep(time.Millisecond * time.Duration(randomValue))
}