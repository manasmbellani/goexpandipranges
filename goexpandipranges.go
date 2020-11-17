package main

// Golang Script to expand IP ranges to individual IP addresses, and write them
// to output
//
// Based on work by 'kotakanbe' published @
//     https://gist.github.com/kotakanbe/d3059af990252ba89a82
//

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"sync"
)

func getIndividualHosts(cidr string, excludeNetworkAddr bool,
	excludeBroadcastAddr bool) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}

	// Remove the network and broadcast addresses, if required
	if len(ips) > 2 {
		if excludeNetworkAddr {
			ips = ips[1:len(ips)]
		}
		if excludeBroadcastAddr {
			ips = ips[0 : len(ips)-1]
		}
	}
	return ips, nil
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
	var numGoRoutines int
	var excludeNetworkAddr bool
	var excludeBroadcastAddr bool

	flag.IntVar(&numGoRoutines, "t", 20, "Number of goroutines to use")
	flag.BoolVar(&excludeNetworkAddr, "en", false, "Exclude the network address")
	flag.BoolVar(&excludeBroadcastAddr, "eb", false, "Exclude the broadcast address")
	flag.Parse()

	var wg sync.WaitGroup

	// Maintain a channel with IP ranges to expand
	ipRanges := make(chan string)

	// Start the goRoutines to expand the IP ranges
	for i := 0; i < numGoRoutines; i++ {

		go func(ipRanges chan string, wg *sync.WaitGroup) {

			for ipRange := range ipRanges {
				defer wg.Done()
				ips, _ := getIndividualHosts(ipRange, excludeNetworkAddr,
					excludeBroadcastAddr)
				for _, ip := range ips {
					fmt.Println(ip)
				}
			}

		}(ipRanges, &wg)
	}

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		ipRange := sc.Text()

		if ipRange != "" {
			ipRanges <- ipRange
			wg.Add(1)

		}
	}

	close(ipRanges)

	wg.Wait()
}
