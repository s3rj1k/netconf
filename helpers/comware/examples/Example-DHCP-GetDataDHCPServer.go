package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/exsver/netconf/helpers/comware"
	"log"
)

func main() {
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf", "netconf")
	if err != nil {
		log.Fatalf("%s", err)
	}
	dhcpInfo, err := sw.GetDataDHCPServer()
	if err != nil {
		log.Fatalf("%s", err)
	}
	spew.Dump(dhcpInfo)
}
