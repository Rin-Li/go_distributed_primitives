package main

import (
	"go_distributed_primitives/simulator"
	"github.com/redis/go-redis/v9"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
    // Run simulation, the number represents the number of buyers.
	// simulator.RunSimulationWithoutLock(client, 50)
	// Run simulation, the number represents the number of student want to enter the library.
	//simulator.RunSimulationWithSemaphore(client, 50)
	// Run simulation, for leaky bucketand token bucket, concurrency number, each request
	//simulator.RunLeakyBucketSimulation(client, 100, 100)
	simulator.RunTokenBucketSimulation(client, 100, 100)


}