package main

import (
	"fmt"
	"log"

	"github.com/exsver/netconf/comware"
)

func main() {
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf", "netconf")
	if err != nil {
		log.Fatalf("%s", err)

	}

	agingTime, err := sw.GetMacAgingTime()
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Printf("Min AgingTime: %v seconds Max AgingTime: %v seconds, Current AgingTime: %v seconds\n", agingTime.AgingTimeMin, agingTime.AgingTimeMax, agingTime.AgingTime)

	macSpecification, err := sw.GetMacSpecification()
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Printf("PortLearnMaxNumLimit: %v\n", macSpecification.PortLearnMaxNumLimit)
}
