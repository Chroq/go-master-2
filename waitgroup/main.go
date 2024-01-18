package main

import (
	"fmt"
	"sync"
)

func printMessage(message string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(message)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go printMessage("Hello", &wg)
	go printMessage("World", &wg)

	wg.Wait()
}
