package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func getServices(serviceName string) []*net.SRV {
	_, srvs, err := net.LookupSRV("", "", serviceName)
	if err != nil {
		log.Fatal(err)
	}
	return srvs
}

func getIP(serviceName string) []net.IP {
	listOfIps, err := net.LookupIP(serviceName)
	if err != nil {
		log.Println("LookUp failed")
		log.Fatal(err)
	}
	return listOfIps
}

func printIP(listOfIps []net.IP, srv string) {
	for _, ip := range listOfIps {
		fmt.Println(srv+" => ", ip.String())
	}
}

func main() {
	name, err := os.Hostname()
	if err != nil {
		log.Println("pod's hostname cannot be retrieved...")
		log.Fatal(err)
	}
	srvs := getServices("kubia")
	for _, srv := range srvs {
		if name != srv.Target[:len(name)] {
			ips := getIP(srv.Target)
			printIP(ips, srv.Target)
		}
	}
}

