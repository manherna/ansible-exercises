package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please input at least 1 IP")
		os.Exit(1)
	}
	// Expects a single host IP or a CIDR network designation
	ip := os.Args[1]

	if strings.Contains(ip, "/") {
		ips := GetNetworkRange(ip)
		for _, i := range ips {
			fmt.Println(i)
			ScanAllPorts(i)
		}
	} else {
		ScanAllPorts(ip)
	}

}

func GetNetworkRange(cidrRange string) []string {
	// convert string to IPNet struct
	_, ipv4Net, err := net.ParseCIDR(cidrRange)
	if err != nil {
		fmt.Println(err)
	}

	// convert IPNet struct mask and address to uint32
	mask := binary.BigEndian.Uint32(ipv4Net.Mask)
	start := binary.BigEndian.Uint32(ipv4Net.IP)
	// find the final address
	finish := (start & mask) | (mask ^ 0xffffffff)

	var result []string
	// loop through addresses as uint32
	for i := start; i <= finish; i++ {
		// convert back to net.IP
		ip := make(net.IP, 4)
		binary.BigEndian.PutUint32(ip, i)
		result = append(result, ip.String())
	}
	return result
}

func ScanAllPorts(ip string) {
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
	fmt.Printf("* %v/tcp\topen\n", port)
}
