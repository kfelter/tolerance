package tolerance

import (
	"log"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var (
	// Threads - change this to increase or decrease redundancy workers
	Threads = runtime.GOMAXPROCS(0)
	// Fault - change this to insert faulty data into the math functions simulating a bit flip
	Fault = 0
)

func Do(nums []float64, mathFunc func(...float64) float64) (float64, int, time.Duration) {
	counter := 0
	start := time.Now()
	for {
		votes := map[float64]int{}
		wg := sync.WaitGroup{}
		voteChan := make(chan float64)

		for i := 0; i < Threads; i++ {
			wg.Add(1)
			// add some Fault
			numsL := nums
			if rand.Intn(100) < Fault {
				numsL = []float64{rand.Float64() * 100, rand.Float64() * 100}
			}
			go func(i int, n []float64) {
				defer wg.Done()
				res := mathFunc(n...)
				log.Printf("worker [%d] votes %v\n", i, res)
				voteChan <- res
			}(i, numsL)
		}
		go func() {
			wg.Add(Threads)
			for v := range voteChan {
				votes[v]++
				wg.Done()
			}
		}()
		wg.Wait()
		close(voteChan)

		for i, v := range votes {
			if v > Threads/2 {
				log.Printf("result %v has %d votes", i, v)
				return i, counter + 1, time.Now().Sub(start)
			}
		}
		counter++
		log.Println("fualt on more than", Threads/2-1, "/", Threads, "redundant operations - retry")
	}
}
