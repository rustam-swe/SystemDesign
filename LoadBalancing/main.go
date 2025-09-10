package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

type SimpleServer struct {
	address string
	proxy   *httputil.ReverseProxy
}

type Server interface {
	Address() string
	IsAlive() bool
	Serve(http.ResponseWriter, *http.Request)
}

func NewSimpleServer(address string) *SimpleServer {
	serverUrl, err := url.Parse(address)
	handleError(err)
	return &SimpleServer{
		address: address,
		proxy:   httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

func handleError(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func (s *SimpleServer) Address() string {
	return s.address
}

func (s *SimpleServer) IsAlive() bool {
	return true
}

func (s *SimpleServer) Serve(w http.ResponseWriter, r *http.Request) {
	s.proxy.ServeHTTP(w, r)
}

type LoadBalancer struct {
	port            string
	roundRobinCount int
	servers         []Server
}

func NewLoadBalancer(port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{
		port:            port,
		roundRobinCount: 0,
		servers:         servers,
	}
}

func (lb *LoadBalancer) getNextAvailableServer() Server {
	server := lb.servers[lb.roundRobinCount%len(lb.servers)]
	lb.roundRobinCount = (lb.roundRobinCount + 1) % len(lb.servers)
	if server.IsAlive() {
		return server
	}

	return lb.getNextAvailableServer()
}

func (lb *LoadBalancer) serveProxy(w http.ResponseWriter, r *http.Request) {
	target := lb.getNextAvailableServer()
	fmt.Printf("Forwarding request to address: %s\n", target.Address())
	target.Serve(w, r)
}

func main() {
	servers := []Server{
		NewSimpleServer("https://daryo.uz"),
		NewSimpleServer("https://kun.uz"),
		NewSimpleServer("https://afisha.uz"),
	}

	lb := NewLoadBalancer(":8000", servers)

	handleRedirect := func(w http.ResponseWriter, r *http.Request) {
		lb.serveProxy(w, r)
	}
	http.HandleFunc("/", handleRedirect)
	fmt.Printf("Load balancer started at localhost%s\n", lb.port)
	handleError(http.ListenAndServe(lb.port, nil))
}
