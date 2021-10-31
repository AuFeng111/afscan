package icmpalive

import (
	"log"
	"iprange-master"
	"github.com/gpool"
	"time"
	"icmp"
	//"flag"
)

var ipalive []string
var ip string
var t int

func RunTest(i string,pool *gpool.Pool)  {
	a := icmp.Main(i)
	if a !=""{
		ipalive = append(ipalive, a)
	}
	pool.Done()
}

func Main(t int,ip string) []string{

	size:=t
	pool := gpool.New(size) //设置线程池大小
	if ip !=""{
		list, err := iprange.ParseList(ip)
		if err != nil {
			log.Printf("error: %s", err)
		}
		log.Printf("%+v", list)
		rng := list.Expand()

		start := time.Now()  //开始计时
		for _, host := range rng {
			pool.Add(1)
			go RunTest(host.String(),pool)
		}
		pool.Wait()
		elapsed := time.Since(start) //结束计时
		log.Println("该函数执行完成耗时: ", elapsed)
		log.Println("存活的主机数量: ",len(ipalive))
		//return ipalive
	}
	return ipalive
}