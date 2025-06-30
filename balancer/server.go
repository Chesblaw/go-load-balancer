package balancer

import (
	"net/http/httputil"
	"net/url"
	"sync"
)

type Server struct {
	URL   *url.URL
	Alive bool
	Mutex sync.RWMutex
	Proxy *httputil.ReverseProxy
}

func NewServer(rawURL string) *Server {
	parsedURL, _ := url.Parse(rawURL)
	return &Server{
		URL:   parsedURL,
		Alive: true,
		Proxy: httputil.NewSingleHostReverseProxy(parsedURL),
	}
}

func (s *Server) SetAlive(alive bool) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	s.Alive = alive
}

func (s *Server) IsAlive() bool {
	s.Mutex.RLock()
	defer s.Mutex.RUnlock()
	return s.Alive
}
