package main

// Golang Script to expand IP ranges to individual IP addresses, and write them
// to output
//
// Based on work by 'kotakanbe' published @
//     https://gist.github.com/kotakanbe/d3059af990252ba89a82
//

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)

func getIndividualHosts(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}
	// remove network address and broadcast address
	return ips[1 : len(ips)-1], nil
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func main() {

	var wg sync.WaitGroup

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		ipRange := sc.Text()

		if ipRange != "" {
			wg.Add(1)
			go func(ipRange string) {
				defer wg.Done()
				ips, _ := getIndividualHosts(ipRange)
				for _, ip := range ips {
					fmt.Println(ip)
				}
			}(ipRange)
		}
	}
	wg.Wait()
}
