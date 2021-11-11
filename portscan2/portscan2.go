package portscan2

import (
	"afscan/icmpalive"
	"fmt"
	"net"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Addr struct {
	ip   string
	port int
}


func PortScan(ip string, ports string, t int, timeout int64) []string {
	var addresses []string

	var begin = time.Now()
	hostslist := icmpalive.Main(t, ip) //先写死500线程
	probePorts := parsePort(ports)
	workers := 600
	Addrs := make(chan Addr, len(hostslist)*len(probePorts))
	//results := make(chan string, len(hostslist)*len(probePorts))
	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		go func() {
			for addr := range Addrs {
				a := PortConnect(addr, timeout, &wg)
				//PortConnect(addr, results, timeout, &wg)
				wg.Done()
				if a != "" {
					addresses = append(addresses, a)
				}
			}
		}()
	}

	//添加扫描目标
	for _, port := range probePorts {
		for _, host := range hostslist {
			wg.Add(1)
			Addrs <- Addr{host, port}
		}
	}
	wg.Wait()
	close(Addrs)
	//close(results)
	var elapseTime = time.Now().Sub(begin)
	fmt.Println("耗时:", elapseTime, "目标端口开放数量 :", len(addresses))
	return addresses
}

func PortConnect(addr Addr, Timeout int64, wg *sync.WaitGroup) string {
	host, port := addr.ip, addr.port
	conn, err := net.DialTimeout("tcp4", fmt.Sprintf("%s:%v", host, port), time.Duration(Timeout)*time.Second)
	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()
	if err != nil {
		return ""
	}
	address := host + ":" + strconv.Itoa(port)
	//addresses = append(addresses, address)
	fmt.Println("open:", address)
	return address
}

func parsePort(ports string) []int {
	var scanPorts []int
	countSplit := strings.Split(ports, ",")
	for _, port := range countSplit {
		port = strings.Trim(port, " ")
		upper := port
		if strings.Contains(port, "-") {
			ranges := strings.Split(port, "-")
			if len(ranges) < 2 {
				continue
			}
			sort.Strings(ranges)
			port = ranges[0]
			upper = ranges[1]
		}
		start, _ := strconv.Atoi(port)
		end, _ := strconv.Atoi(upper)
		for i := start; i <= end; i++ {
			scanPorts = append(scanPorts, i)
		}
	}
	//fmt.Println(scanPorts)
	return scanPorts
}
