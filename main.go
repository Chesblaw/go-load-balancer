package main

import (
	"log"
	"net/http"
	"time"

	"go-load-balancer/balancer" // <- matches your go.mod module name
)

func main() {
	servers := []string{
		"http://localhost:8081",
		"http://localhost:8082",
	}

	pool := balancer.NewServerPool()
	for _, addr := range servers {
		pool.AddServer(addr)
	}

	//health checking
	checker := balancer.NewHealthChecker(pool, 10*time.Second, 2*time.Second)
	go checker.Start()

	//load balancer HTTP server
	lb := &balancer.LoadBalancer{Pool: pool}
	log.Println("Load balancer running at :8080")
	if err := http.ListenAndServe(":8080", lb); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
