package main

import (
	"fmt"
	"sort"
)

type ipRange struct {
	from, to int
}

func countIPs(listIpRanges []ipRange) int {
	counter := 0
	for _, inputIpRange := range listIpRanges {
		// +1 because inclusive interval
		counter += (inputIpRange.to - inputIpRange.from + 1)
	}
	return counter
}

func GetFreeIPs(blockedIps []ipRange, max int) []ipRange {
	// List of free ip ranges
	freeIps := []ipRange{}

	// Sort ip ranges by from ip
	sort.Slice(blockedIps, func(i, j int) bool {
		if blockedIps[i].from == blockedIps[j].from {
			return blockedIps[i].to < blockedIps[j].to
		}
		return blockedIps[i].from < blockedIps[j].from
	})

	// CurrentMin holds the minimum ip that may be free
	// At first, 0 might be free
	currentMin := 0
	for _, blockedIpRange := range blockedIps {
		// if blocked ip range starts with ip higher than currentMin, then currentMin to that ip is free
		if currentMin < blockedIpRange.from {
			freeIps = append(freeIps, ipRange{from: currentMin, to: blockedIpRange.from - 1})
		}
		// If last in blocked ip range is higher than current min, then currentMin ip to that ip is blocked and currentMin is the next ip
		if blockedIpRange.to+1 > currentMin {
			currentMin = blockedIpRange.to + 1
		}
	}
	if currentMin <= max {
		// Add last ips
		freeIps = append(freeIps, ipRange{from: currentMin, to: max})
	}
	return freeIps
}

func main() {
	blockedIps := []ipRange{}

	for {
		var inputIpRange ipRange
		_, err := fmt.Scanf("%d-%d", &inputIpRange.from, &inputIpRange.to)
		if err != nil {
			break
		}
		blockedIps = append(blockedIps, inputIpRange)
	}
	// Get list of free ip ranges sorted
	freeIps := GetFreeIPs(blockedIps, 4294967295)

	// Part 1
	result1 := freeIps[0].from
	fmt.Println(result1)

	// Part 2
	result2 := countIPs(freeIps)
	fmt.Println(result2)
}
