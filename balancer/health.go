package balancer

import (
	"log"
	"net/http"
	"time"
)

type HealthChecker struct {
	Pool     *ServerPool
	Interval time.Duration
	Timeout  time.Duration
}

func NewHealthChecker(pool *ServerPool, interval, timeout time.Duration) *HealthChecker {
	return &HealthChecker{
		Pool:     pool,
		Interval: interval,
		Timeout:  timeout,
	}
}

func (hc *HealthChecker) checkServer(s *Server) bool {
	client := http.Client{Timeout: hc.Timeout}
	resp, err := client.Get(s.URL.String())
	if err != nil || resp.StatusCode != http.StatusOK {
		return false
	}
	return true
}

func (hc *HealthChecker) Start() {
	ticker := time.NewTicker(hc.Interval)
	for range ticker.C {
		log.Println("Running health checks...")
		for _, server := range hc.Pool.Servers {
			alive := hc.checkServer(server)
			server.SetAlive(alive)
			log.Printf("Checked %s → Alive: %v\n", server.URL, alive)
		}
	}
}
