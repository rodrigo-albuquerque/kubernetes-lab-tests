package main

import (
	"fmt"
	"log"
	"net"
)

// returns DNS SRV records associated with a particular Service
// Useful to retrieve individual pods from headless Service
func retrievePodSRVs(serviceName string) []*net.SRV {
	_, srvs, err := net.LookupSRV("", "", serviceName)
	if err != nil {
		log.Fatal(err)
	}
	return srvs
}

func main() {
	serviceNames := retrievePodSRVs("TYPE_YOUR_SERVICE_NAME_HERE_TO_TEST")
	for _, srv := range serviceNames {
		fmt.Printf("%v\n", srv.Target)
	}
}

