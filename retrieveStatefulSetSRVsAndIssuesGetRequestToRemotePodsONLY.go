// Issues GET request to all remote stateful pods by probing DNS SRV record through headless service
// and responds to query locally directly
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

// retrieves k8s all pods FQDN from headless service
func retrievePodSRVs(serviceName string) []*net.SRV {
	_, srvs, err := net.LookupSRV("", "", serviceName)
	if err != nil {
		log.Fatal(err)
	}
	return srvs
}

// issues a get request to single pod given its FQDN or IP address on port 8080
func getPodData(url string) string {
	resp, err := http.Get("http://" + url + ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return string(body)
}

// You need to type your headless service down below
func main() {
	srvs := retrievePodSRVs("TYPE_YOUR_HEADLESS_SERVICE_NAME_HERE")
	name, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	for _, srv := range srvs {
		// If service is local, just respond directly
		if strings.Contains(srv.Target, string(name)) {
			// print pod's hostname
			fmt.Println("You've hit " + name)
			// Otherwise, if pod is remote, issue a GET request and print response
		} else {
			respBody := getPodData(srv.Target)
			fmt.Println(respBody)
		}
	}
}

