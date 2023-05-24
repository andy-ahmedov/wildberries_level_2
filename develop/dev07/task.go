package main

import (
	"fmt"
	"sync"
	"time"
)

func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

func or(channels ...<-chan interface{}) <-chan interface{} {
	result := make(chan interface{})
	var wg sync.WaitGroup
	wg.Add(1)
	for _, channel := range channels {
		go func(channel <-chan interface{}) {
			for x := range channel {
				result <- x
			}
			wg.Done()
		}(channel)
	}
	wg.Wait()
	close(result)
	return result
}

func main() {
	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("fone after %v", time.Since(start))
}