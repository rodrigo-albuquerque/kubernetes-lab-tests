package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func retrievePodSRVs(serviceName string) []*net.SRV {
	_, srvs, err := net.LookupSRV("", "", serviceName)
	if err != nil {
		log.Fatal(err)
	}
	return srvs
}

func main() {
	if len(os.Args) == 2 {
		myArgs := os.Args[1]
		serviceNames := retrievePodSRVs(myArgs)
		for _, srv := range serviceNames {
			fmt.Printf("%v\n", srv.Target)
		}
	} else {
		log.Fatal("Please, specify ONLY one argument.")
	}
}
