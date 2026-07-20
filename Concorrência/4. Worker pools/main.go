package main

import (
	"fmt"
	"worker_pools/workers"
)

func main() {
	workers.WorkerPool()
	fmt.Println("-------------------------")
	workers.WorkerPoolSemBuffer()
}
