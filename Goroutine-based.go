package main

import (
	"fmt"
	"math"
	"sync"
)

func checkDivisibility(number int, start int, end int, resultChan chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := start; i <= end; i++ {
		if number%i == 0 {
			resultChan <- false
			return
		}
	}
}

func isPrime(number int) bool {
	if number <= 1 {
		return false
	}
	if number <= 3 {
		return true
	}
	if number%2 == 0 || number%3 == 0 {
		return false
	}

	limit := int(math.Sqrt(float64(number)))

	resultChan := make(chan bool, 1)

	numGoroutines := 4
	rangeSize := (limit - 4) / numGoroutines
	if rangeSize < 1 {
		rangeSize = 1
	}

	var wg sync.WaitGroup

	for i := 0; i < numGoroutines; i++ {
		start := 5 + i*rangeSize
		end := start + rangeSize - 1
		if i == numGoroutines-1 {
			end = limit
		}
		wg.Add(1)
		go checkDivisibility(number, start, end, resultChan, &wg)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for result := range resultChan {
		if result == false {
			return false
		}
	}

	return true
}

func main() {
	var number int
	fmt.Print("Enter a number to check if it is prime: ")
	fmt.Scanln(&number)

	if isPrime(number) {
		fmt.Printf("%d is a Prime Number ✅\n", number)
	} else {
		fmt.Printf("%d is NOT a Prime Number ❌\n", number)
	}
}
