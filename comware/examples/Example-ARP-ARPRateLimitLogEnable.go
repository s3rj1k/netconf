package main

import (
	"log"

	"github.com/exsver/netconf/comware"
)

// CLI equivalent
//  ARPRateLimitLogEnable(true) - "arp rate-limit log enable"
//  ARPRateLimitLogEnable(false) - "undo arp rate-limit log enable"

func main() {
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	err = sw.ARPRateLimitLogEnable(true)
	if err != nil {
		log.Fatalf("%s", err)
	}
}
