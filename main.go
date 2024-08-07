package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sync"
)

type Server struct {
	IpAddress   string
	CurrentLoad int
}

func (s *Server) incrementLoad() {
	s.CurrentLoad += 1
}

func (s *Server) getCurrentLoad() int {
	return s.CurrentLoad
}

type LoadBalancer struct {
	Servers *MinHeap
}

func (l *LoadBalancer) addServer(ipaddress string, load int) {
	newServer := &Server{IpAddress: ipaddress, CurrentLoad: load}
	heap.Push(l.Servers, newServer)
}

func (l *LoadBalancer) getNextServer() *Server {
	fmt.Println("getting next server")
	if l.Servers.Len() == 0 {
		fmt.Println("No servers available")
		return nil
	}

	minLoadServer := heap.Pop(l.Servers).(*Server)
	minLoadServer.incrementLoad()

	heap.Push(l.Servers, minLoadServer)

	return minLoadServer
}

func InitLoadBalancer(serverCount int) *LoadBalancer {
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	lb := &LoadBalancer{Servers: minHeap}

	for i := 0; i < serverCount; i++ {
		fmt.Println("Enter Server #", i+1, "IP addr:")

		var ipaddress string
		fmt.Scanln(&ipaddress)
		lb.addServer(ipaddress, 0)
	}

	return lb
}

func main() {
	wg := sync.WaitGroup{}
	var serverCount int
	fmt.Println("Enter the number of servers to initialize:")
	fmt.Scanln(&serverCount)

	lb := InitLoadBalancer(serverCount)

	fmt.Println("listening to requests")
	wg.Add(1)
	go func() {
		for {
			input := bufio.NewScanner(os.Stdin)
			input.Scan()
			if input.Text() == "exit" {
				wg.Done()
			}
			server := lb.getNextServer()
			fmt.Println("request routed to Server:", server.IpAddress)
		}
	}()

	wg.Wait()
	fmt.Println("exiting")
}
