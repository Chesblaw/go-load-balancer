package balancer

import (
	"net/http"
)

// LoadBalancer routes requests to healthy backend servers
type LoadBalancer struct {
    Pool *ServerPool
}

func (lb *LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    target := lb.Pool.GetNextAliveServer()
    if target == nil {
        http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
        return
    }

    target.Proxy.ServeHTTP(w, r)
}
