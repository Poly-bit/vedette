package main

import (
	"context"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

// PortScanner struct holding our ip lock
type PortScanner struct {
	ip   string
	lock *semaphore.Weighted
}

func scanPort(ip string, port int, timeout time.Duration) {
	target := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", target, timeout)
	if err != nil {
		if strings.Contains(err.Error(), "too many open files") {
			time.Sleep(timeout)
			scanPort(ip, port, timeout)
		} else {
			fmt.Println(port, "closed")
		}
		return
	}
	conn.Close()
	fmt.Println(port, "open")
}

// Scan function to actually do the scanning
func (ps *PortScanner) Scan(f, l int, timeout time.Duration) {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	for port := f; port <= l; port++ {
		wg.Add(1)
		ps.lock.Acquire(context.TODO(), 1)
		go func(port int) {
			defer ps.lock.Release(1)
			defer wg.Done()
			scanPort(ps.ip, port, timeout)
		}(port)
	}
}

func main() {
	ps := &PortScanner{
		ip:   "127.0.0.1",
		lock: semaphore.NewWeighted(1048576),
	}
	ps.Scan(1, 65535, 500*time.Millisecond)
}
