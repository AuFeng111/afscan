package portscan

import (
	"fmt"
	//"gohive"
	"afscan/icmpalive"
	"net"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
	//"golang.org/x/crypto/ssh"
	//"github.com/gpool"
	"iprange-master"
)

var addresses []string
var wg sync.WaitGroup

func Main(ports string, t int, ip string) []string { //先进行icmp探测
	ipalive := icmpalive.Main(t, ip)
	//fmt.Println("ipalive  end")
	var begin = time.Now()

	//ipalive :=[...]string{"10.82.200.91","10.82.200.86","10.82.200.1","10.82.200.98","10.82.200.68","10.82.200.81","10.82.200.75","10.82.200.79","10.82.200.54","10.82.200.11","10.82.200.48","10.82.200.20","10.82.200.61","10.82.200.90","10.82.200.15","10.82.200.40","10.82.200.8","10.82.200.18","10.82.200.103","10.82.200.26","10.82.200.64","10.82.200.89","10.82.200.87","10.82.200.27","10.82.200.100","10.82.200.47","10.82.200.43","10.82.200.2","10.82.200.109","10.82.200.49","10.82.200.51","10.82.200.131"}
	//ipalive :=[...]string{"121.4.236.99","121.4.236.5","121.4.236.9","121.4.236.14","121.4.236.6","121.4.236.11","121.4.236.8","121.4.236.59","121.4.236.96","121.4.236.84","121.4.236.2","121.4.236.10","121.4.236.53","121.4.236.39","121.4.236.62","121.4.236.63"}
	//ipalive :=[...]string{"10.82.10.30","10.82.10.33","10.82.10.103","10.82.10.106","10.82.10.40","10.82.10.21","10.82.10.104","10.82.10.111","10.82.10.110","10.82.10.109","10.82.10.55","10.82.10.108","10.82.10.66","10.82.10.150","10.82.10.102","10.82.10.101","10.82.10.151","10.82.10.100","10.82.10.152","10.82.10.170","10.82.10.160"}
	//ipalive :=[...]string{"192.168.0.3","192.168.0.112","192.168.0.117","192.168.0.1"}
	var scanport = make(chan string, 1000)
	//pool := gpool.New(1000)
	//results := make(chan string, 100)
	//fmt.Println(reflect.TypeOf(ipalive))
	//fmt.Scanf("%s",&str)
	a := parsePort(ports) //分割端口
	//fmt.Println(a)
	//results = squarer(scanport,results)
	go func() {
		//fmt.Println(a)
		//用于生成ip:port,并且存放到地址管道种
		for _, ip := range ipalive {
			for i := 0; i < len(a); i++ {
				var address = fmt.Sprintf("%s:%d", ip, a[i])
				scanport <- address
				//fmt.Println(scanport,' ',i)
			}
		}
		close(scanport)
		//var elapseTime = time.Now().Sub(begin)
	}()
	//fmt.Println("通道获取长度:", len(scanport))

	for i := 0; i < 5000; i++ {
		wg.Add(1)
		//pool.Add(1)
		go worker(scanport)
		//go worker(scanport,pool)

	}
	/*
		for work := 0; work < pool_size; work++ {
			wg.Add(1)
			pool.Submit(worker)
		}*/
	//等待结束
	wg.Wait()
	//fmt.Println("----------------------")

	//pool.Wait()
	//计算时间
	var elapseTime = time.Now().Sub(begin)
	fmt.Println("耗时:", elapseTime, "目标端口开放数量 :", len(addresses))
	return addresses
}

func Second_Main(ports string, ip string) []string { //先进行icmp探测
	ipalive := Iplist(ip)
	//fmt.Println("ipalive  end")
	var begin = time.Now()

	//ipalive :=[...]string{"10.82.200.91","10.82.200.86","10.82.200.1","10.82.200.98","10.82.200.68","10.82.200.81","10.82.200.75","10.82.200.79","10.82.200.54","10.82.200.11","10.82.200.48","10.82.200.20","10.82.200.61","10.82.200.90","10.82.200.15","10.82.200.40","10.82.200.8","10.82.200.18","10.82.200.103","10.82.200.26","10.82.200.64","10.82.200.89","10.82.200.87","10.82.200.27","10.82.200.100","10.82.200.47","10.82.200.43","10.82.200.2","10.82.200.109","10.82.200.49","10.82.200.51","10.82.200.131"}
	//ipalive :=[...]string{"121.4.236.99","121.4.236.5","121.4.236.9","121.4.236.14","121.4.236.6","121.4.236.11","121.4.236.8","121.4.236.59","121.4.236.96","121.4.236.84","121.4.236.2","121.4.236.10","121.4.236.53","121.4.236.39","121.4.236.62","121.4.236.63"}
	//ipalive :=[...]string{"10.82.10.30","10.82.10.33","10.82.10.103","10.82.10.106","10.82.10.40","10.82.10.21","10.82.10.104","10.82.10.111","10.82.10.110","10.82.10.109","10.82.10.55","10.82.10.108","10.82.10.66","10.82.10.150","10.82.10.102","10.82.10.101","10.82.10.151","10.82.10.100","10.82.10.152","10.82.10.170","10.82.10.160"}
	//ipalive :=[...]string{"192.168.0.3","192.168.0.112","192.168.0.117","192.168.0.1"}
	var scanport = make(chan string, 1000)
	//pool := gpool.New(1000)
	//results := make(chan string, 100)
	//fmt.Println(reflect.TypeOf(ipalive))
	//fmt.Scanf("%s",&str)
	a := parsePort(ports) //分割端口
	//fmt.Println(a)
	//results = squarer(scanport,results)
	go func() {
		//fmt.Println(a)
		//用于生成ip:port,并且存放到地址管道种
		for _, ip := range ipalive {
			for i := 0; i < len(a); i++ {
				var address = fmt.Sprintf("%s:%d", ip, a[i])
				scanport <- address
				//fmt.Println(scanport,' ',i)
			}
		}
		close(scanport)
		//var elapseTime = time.Now().Sub(begin)
	}()
	//fmt.Println("通道获取长度:", len(scanport))

	for i := 0; i < 5000; i++ {
		wg.Add(1)
		//pool.Add(1)
		go worker(scanport)
		//go worker(scanport,pool)

	}
	/*
		for work := 0; work < pool_size; work++ {
			wg.Add(1)
			pool.Submit(worker)
		}*/
	//等待结束
	wg.Wait()
	//fmt.Println("----------------------")

	//pool.Wait()
	//计算时间
	var elapseTime = time.Now().Sub(begin)
	fmt.Println("耗时:", elapseTime, "目标端口开放数量 :", len(addresses))
	return addresses
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

/*
func squarer(results chan string, scanport chan string) chan string{   //存入的值转换为取出的值
	for i := range scanport {
		results <- i
	}
	close(results)
	return results
}*/
//工人
func worker(scanport chan string) {
	//函数结束释放连接
	defer wg.Done()
	for {
		address, ok := <-scanport
		if !ok {
			break
		}
		//fmt.Println("address:", address)
		//conn, err := net.Dial("tcp", address)
		conn, err := net.DialTimeout("tcp", address, 1*time.Second)
		if err != nil {
			//fmt.Println("close:", address, err)
			continue
		}
		conn.Close()
		fmt.Println("open:", address)
		addresses = append(addresses, address)
	}
	//pool.Wait()
	//return addresses
}

func Iplist(ip string) []string {
	var a []string

	list, err := iprange.ParseList(ip)
	if err != nil {
		fmt.Printf("error: %s", err)
	}
	fmt.Printf("%+v", list)
	rng := list.Expand()
	for _, i := range rng {
		a = append(a, i.String())
	}
	return a

}
