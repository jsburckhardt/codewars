package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	sum := 0
	intCh := make(chan int, 1)
	outCh := make(chan int, 1)

	for wID := 1; wID <= 5; wID++ {
		go addI(wID, intCh, outCh)
	}

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		intCh <- i
		sum = sum + <-outCh
	}
	wg.Wait()
	log.Printf("sum = %d", sum)
}

func addI(wID int, intCh chan int, outCh chan int) {
	for {
		i := <-intCh
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		time.Sleep(time.Duration(r.Intn(5)) * time.Second)
		log.Printf("wID = %d. i=%d", wID, i)
		outCh <- wID
		wg.Done()
	}
}
