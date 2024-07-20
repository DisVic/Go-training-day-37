package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	t := time.Now()
	fmt.Print(findPrimeNumbers(50000), "\n")
	fmt.Print(time.Since(t).String())
}

func findPrimeNumbers(count int) []int {
	primeNumbers := make([]int, 0)
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	for i := 2; i <= count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			div := 0
			for j := 1; j <= i; j++ {
				if i%j == 0 {
					div++
				}
			}
			mu.Lock()
			if div <= 2 {
				primeNumbers = append(primeNumbers, i)
			}
			mu.Unlock()
		}()
	}
	wg.Wait()
	return primeNumbers
}
