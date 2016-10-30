package main

import (
	"context"
	"log"
	"sync"
	"time"
)

func ExmapleContextRelationship() {
	ctx1, cancel1 := context.WithCancel(context.Background())
	ctx2, cancel2 := context.WithCancel(ctx1)
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()

		log.Printf("waiting for ctx1...")
		<-ctx1.Done()
		log.Printf("ctx1.Done() returns")
	}()

	go func() {
		defer wg.Done()

		log.Printf("waiting for ctx2...")
		<-ctx2.Done()
		log.Printf("ctx2.Done() returns")
	}()

	time.AfterFunc(time.Second*5, cancel2)
	time.AfterFunc(time.Second*1, cancel1)

	wg.Wait()
}

func main() {
	ExmapleContextRelationship()
}
