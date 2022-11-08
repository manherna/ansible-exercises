package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {

	if len(os.Args) < 2 {
		println("Please input at least 1 IP")
		os.Exit(1)
	}
	// Expects a single host IP or a CIDR network designation
	ip := os.Args[1]

	if strings.Contains(ip, "/") {
		ips := GetNetworkRange(ip)
		println("CIDR range")
		for _, i := range ips {
			println(i)
			ScanAllPorts(i)
		}
	} else {
		ScanAllPorts(ip)
	}

}

func GetNetworkRange(cidrRange string) []string {
	//TODO: Implement network ranges parsing
	return []string{"1.1.1.1", "2.2.2.2", "3.3.3.3", "4.4.4.4"}
}

func ScanAllPorts(ip string) {
	println(ip)
	maxPort := 65555
	for i := 1; i < maxPort; i++ {
		go ScanPort(ip, i, 500*time.Millisecond)
	}
}

// Sourced from tutorial https://medium.com/@KentGruber/building-a-high-performance-port-scanner-with-golang-9976181ec39d
// simplified without ulimits
func ScanPort(ip string, port int, timeout time.Duration) {
	target := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", target, timeout)

	if err != nil {
		if strings.Contains(err.Error(), "too many open files") {
			time.Sleep(timeout)
			ScanPort(ip, port, timeout)
		} else {
			fmt.Printf("* %v/tcp\tclosed\n", port)
		}
		return
	}

	conn.Close()
	fmt.Printf("* %v/tcp\tclosed\n", port)
}
