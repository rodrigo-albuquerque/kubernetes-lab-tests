package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}
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
	srvs := retrievePodSRVs("TYPE_YOUR_HEADLESS_K8S_SERVICE_NAME_HERE")
	for _, srv := range srvs {
		respBody := getPodData(srv.Target)
		fmt.Println(respBody)
	}
}

