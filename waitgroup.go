package main

import (
	"fmt"
	"sync"
)

func runner1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Print("\nI am first runner")
}

func runner2(wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(1)
	fmt.Print("\nI am second runner")
}

func execute() {
	wg := new(sync.WaitGroup)

	if true {
		wg.Add(1)
		go runner1(wg)

	}

	if true {
		go runner2(wg)
	}

	wg.Wait()
}

func main() {

	// Launching both the runners
	execute()
}
