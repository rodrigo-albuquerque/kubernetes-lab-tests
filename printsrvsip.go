package main

import (
	"fmt"
	"log"
	"net"
)

func retrievePodSRVs(serviceName string) []*net.SRV {
	_, srvs, err := net.LookupSRV("", "", serviceName)
	if err != nil {
		log.Fatal(err)
	}
	return srvs
}

func getIP(serviceName string) []net.IP {
	ip, err := net.LookupIP(serviceName)
	if err != nil {
		log.Println("LookUp failed")
		log.Println(err)
	}
	return ip
}

func main() {
	srvs := retrievePodSRVs("kubia")
	for _, srv := range srvs {
		ips := getIP(srv.Target)
		for _, ip := range ips {
			fmt.Println(srv.Target+" => ", ip.String())
		}
	}
}

