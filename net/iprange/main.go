package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	if ok, _ := testsIP("132.17.2.56/28", "132.17.2.63"); ok {
		fmt.Println("Include")
	} else {
		fmt.Println("Not include")
	}
}

func testsIP(ipRange, testsIP string) (bool, error) {
	if isCIDR(ipRange) {
		return testsIPForCIDR(ipRange, testsIP)
	}
	return ipRange == testsIP, nil
}

func testsIPForCIDR(ipRange, testsIP string) (contains bool, err error) {
	_, ipNet, err := net.ParseCIDR(ipRange)
	if err != nil {
		return
	}
	ip := net.ParseIP(testsIP)
	if ip == nil {
		err = fmt.Errorf("could not parse to ip address: %v", testsIP)
		return
	}
	return ipNet.Contains(ip), nil
}

func isCIDR(ip string) bool {
	return strings.IndexRune(ip, '/') > 0
}
