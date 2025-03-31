package main

import (
	"net/http"
	_ "net/http/pprof"
	"time"
)

// go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30
func consumeCPU() {
	for {
		for i := 0; i < 1000000; i++ {
		}
	}
}

func lessConsumingCPU() {
	for {
		for i := 0; i < 10000; i++ {
		}
	}
}

// go tool pprof http://localhost:6060/debug/pprof/heap
func leakMemory() {
	var leak []string
	for {
		leak = append(leak, "leak memory")
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	// leakMemory()
	go lessConsumingCPU()
	go consumeCPU()

	time.Sleep(60 * time.Minute)
}
