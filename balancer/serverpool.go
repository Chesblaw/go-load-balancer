package balancer

import (
	"sync/atomic"
)

type ServerPool struct {
	Servers []*Server
	current uint64
}

func NewServerPool() *ServerPool {
	return &ServerPool{
		Servers: make([]*Server, 0),
	}
}

func (p *ServerPool) AddServer(rawURL string) {
	server := NewServer(rawURL)
	p.Servers = append(p.Servers, server)
}

func (p *ServerPool) GetNextAliveServer() *Server {
	length := len(p.Servers)
	for i := 0; i < length; i++ {
		index := int(atomic.AddUint64(&p.current, 1)) % length
		if p.Servers[index].IsAlive() {
			return p.Servers[index]
		}
	}
	return nil // No alive servers
}
