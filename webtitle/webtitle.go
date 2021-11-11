package webtitle

import (
	// "encoding/json"
	"afscan/portscan2"
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/goscraper/goscraper"
	//"reflect"
)

// func main() {
// 	ports := "443"
// 	ip := "210.76.75.192/24"
// 	Webtitle(ports, ip)
// }
func Webtitle(ip string, ports string, t int, timeout int64) {
	var wg sync.WaitGroup
	workers := 10
	addr := portscan2.PortScan(ip, ports, t, timeout)
	ip_and_port := make(chan string, len(addr))
	for c := 0; c < len(addr); c++ {
		wg.Add(1)
		ip_and_port <- addr[c]
	}
	// addr := portscan.Second_Main(ports, ip)
	fmt.Println(len(addr), len(ip_and_port))

	if len(addr) > 0 {
		for i := 0; i < workers; i++ {
			go func() {
				for ips := range ip_and_port {
					//wg.Add(1)
					//fmt.Println(ips)
					go title_scan(ips, &wg)
				}
			}()
			//wg.Wait()
		}
		wg.Wait()
	}

}

func title_scan(ips string, wg *sync.WaitGroup) {
	//var a [][]string
	//忽略https的校验
	/*
		var tr = &http.Transport{
			MaxIdleConns:      30,
			IdleConnTimeout:   time.Second,
			DisableKeepAlives: true,
			TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
			DialContext: (&net.Dialer{
				Timeout:   time.Second,
				KeepAlive: time.Second,
			}).DialContext,
		}*/

	defer wg.Done()
	var url string
	//client := &http.Client{Transport: tr, Timeout: time.Second}
	countSplit := strings.Split(ips, ":")
	HTTPS_PORT := []string{"443", "9443", "7443", "8443", "6443"}
	if in(countSplit[1], HTTPS_PORT) == true {
		url = "https://" + ips
	} else {
		url = "http://" + ips
	}
	s, err := goscraper.Scrape(url, 1)
	if err != nil {
		//fmt.Println(err)
		return
	}

	fmt.Printf("%s %-20s\n", url, strings.TrimSpace(s.Preview.Title))
}

func in(target string, str_array []string) bool {
	sort.Strings(str_array)
	index := sort.SearchStrings(str_array, target)
	if index < len(str_array) && str_array[index] == target {
		return true
	}
	return false
}
