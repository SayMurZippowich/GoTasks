package main

import "fmt"

type IPAddr [4]byte

// Add a "String() string" method to IPAddr.
func (adr IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", adr[0], adr[1], adr[2], adr[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
