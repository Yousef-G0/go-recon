package modules

import (
	"fmt"
	"net"
	"sort"
	"time"
)

func RunPortScan(host string) {
	fmt.Println("\033[1;34m[*] Port Scan (Top Ports):\033[0m")
	topPorts := []int{21, 22, 23, 25, 53, 80, 110, 139, 143, 443, 445, 3389, 8080, 3306, 1433, 5900}
	sort.Ints(topPorts)
	for _, port := range topPorts {
		address := fmt.Sprintf("%s:%d", host, port)
		conn, err := net.DialTimeout("tcp", address, 2*time.Second)
		if err == nil {
			fmt.Printf("\033[32m[+] Open:\033[0m %d\n", port)
			conn.Close()
		} else {
			fmt.Printf("\033[31m[-] Closed:\033[0m %d\n", port)
		}
	}
	fmt.Println()
}
