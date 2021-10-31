package webtitle

import (
	// "encoding/json"
	//"afscan/portscan"
	"crypto/tls"
	//"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/axgle/mahonia"
	"golang.org/x/net/html/charset"
	//"reflect"
)

func main() {
	ports := "443"
	ip := "210.76.75.192/24"
	Webtitle(ports, ip)
}

var wg sync.WaitGroup

func Webtitle(ports string, ip string) {
	addr := portscan.Second_Main(ports, ip)
	if len(addr) > 0 {
		for _, ips := range addr {
			wg.Add(1)
			//fmt.Println(ips)
			go title_scan(ips)
		}
		wg.Wait()
	}

}

func title_scan(ips string) {
	var a [][]string
	//忽略https的校验
	var tr = &http.Transport{
		MaxIdleConns:      30,
		IdleConnTimeout:   time.Second,
		DisableKeepAlives: true,
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		DialContext: (&net.Dialer{
			Timeout:   time.Second,
			KeepAlive: time.Second,
		}).DialContext,
	}

	defer wg.Done()
	var url string
	client := &http.Client{Transport: tr, Timeout: time.Second}
	countSplit := strings.Split(ips, ":")
	HTTPS_PORT := []string{"443", "9443", "7443", "8443", "6443"}
	if in(countSplit[1], HTTPS_PORT) == true {
		url = "https://" + ips
	} else {
		url = "http://" + ips
	}
	//fmt.Printf("\n ips=%v \n", ips)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		//fmt.Println(err)
	}
	//设置request的header
	response, err := client.Do(request)
	if err != nil {
		//fmt.Println(err)
		return
	}
	defer response.Body.Close()
	reg := regexp.MustCompile(`<title>(?s:(.*?))</title>`)
	if response.StatusCode == 200 {

		r, err := ioutil.ReadAll(response.Body)
		if err != nil {
			//fmt.Println(err)
		}

		/*
			utf8Reader := transform.NewReader(response.Body, simplifiedchinese.GBK.NewDecoder())
			bodyData, _ := ioutil.ReadAll(utf8Reader)
		*/
		//bodystr := mahonia.NewDecoder("utf-8").ConvertString(string(r))
		if DetermineEncoding(string(r)) == "gbk" {
			a = reg.FindAllStringSubmatch(mahonia.NewDecoder("gbk").ConvertString(string(r)), -1)
		} else {
			a = reg.FindAllStringSubmatch(mahonia.NewDecoder("utf-8").ConvertString(string(r)), -1)
		}

		if len(a) > 0 {
			//fmt.Println(ips + "\t" + strings.TrimSpace(a[0][1]))
			fmt.Printf("%s %5s\n", ips, strings.TrimSpace(a[0][1]))
			//fmt.Println('\n')
		}
	} else {
		//fmt.Print("not code 200")
		return
	}
}

func in(target string, str_array []string) bool {
	sort.Strings(str_array)
	index := sort.SearchStrings(str_array, target)
	if index < len(str_array) && str_array[index] == target {
		return true
	}
	return false
}

func DetermineEncoding(html string) string {
	_, name, _ := charset.DetermineEncoding([]byte(html), "")
	return name
}

/*
func Encoding(html string, ct string) string {
	e, name := DetermineEncoding(html)
	if name != "utf-8" {
		html = ConvertToStr(html, "gbk", "utf-8")
		e = unicode.UTF8
	}
	r := strings.NewReader(html)

	utf8Reader := transform.NewReader(r, e.NewDecoder())
	//将其他编码的reader转换为常用的utf8reader
	all, _ := ioutil.ReadAll(utf8Reader)
	log.Println(string(all))
	return string(all)
}
*/
