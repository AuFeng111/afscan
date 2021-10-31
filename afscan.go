package main

import (
	//"pac"
	"flag"
	"fmt"

	//"log"
	//"icmp"
	"afscan/icmpalive"
	"afscan/portscan"
	"afscan/smbcrack"
	"afscan/sshcrack"
)

var model, ip, port, user, pass string
var t int
var y int

//sshPassword := []string{"123456", "admin", "admin123", "root", "", "pass123", "pass@123", "password", "123123", "654321", "111111", "123", "1", "admin@123", "Admin@123", "admin123!@#", "P@ssw0rd!", "P@ssw0rd", "Passw0rd", "qwe123", "12345678", "test", "test123", "123qwe!@#", "123456789", "123321", "666666", "a123456.", "123456~a", "123456!a", "000000", "1234567890", "8888888", "!QAZ2wsx", "1qaz2wsx", "abc123", "abc123456", "1qaz@WSX", "a11111", "a12345", "Aa1234", "Aa1234.", "Aa12345", "a123456", "a123123", "Aa123123", "Aa123456", "Aa12345.", "sysadmin", "system", "1qaz!QAZ", "2wsx@WSX", "qwe123!@#", "Aa123456!", "A123456s!", "sa123456", "1q2w3e"}

func init() {
	Password := "123456,admin,admin123,root,,pass123,pass@123,password,123123,654321,111111,123,1,admin@123,Admin@123,admin123!@#,P@ssw0rd!,P@ssw0rd,Passw0rd,qwe123,12345678,test,test123,123qwe!@#,123456789,123321,666666,a123456.,123456~a,123456!a,000000,1234567890,8888888,!QAZ2wsx,1qaz2wsx,abc123,abc123456,1qaz@WSX,a11111,a12345,Aa1234,Aa1234.,Aa12345,a123456,a123123,Aa123123,Aa123456,Aa12345.,sysadmin,system,1qaz!QAZ,2wsx@WSX,qwe123!@#,Aa123456!,A123456s!,sa123456,1q2w3e"

	flag.StringVar(&model, "model", "", "icmpalive\nportscan\nsshcrack\nsmbcrack")
	flag.StringVar(&ip, "ip", "127.0.0.1", "192.168.1-255.1-10\n192.168.1.1/24\n192.168.1.*")
	flag.IntVar(&t, "t", 500, "thread")
	flag.StringVar(&port, "port", "22,80,445,3389,8000-9000", "")
	flag.StringVar(&user, "user", "ubuntu,admin,test,user,root", "")
	//flag.StringVar(&pass, "pass", "123456,admin,admin123,root,,pass123,pass@123,password,123123,654321,111111,123,1,admin@123,Admin@123,admin123!@#,P@ssw0rd!,P@ssw0rd,Passw0rd,qwe123,12345678,test,test123,123qwe!@#,123456789,123321,666666,a123456.,123456~a,123456!a,000000,1234567890,8888888,!QAZ2wsx,1qaz2wsx,abc123,abc123456,1qaz@WSX,a11111,a12345,Aa1234,Aa1234.,Aa12345,a123456,a123123,Aa123123,Aa123456,Aa12345.,sysadmin,system,1qaz!QAZ,2wsx@WSX,qwe123!@#,Aa123456!,A123456s!,sa123456,1q2w3e", "")
	flag.StringVar(&pass, "pass", Password, "")
	flag.Usage = func() {
		fmt.Printf("\nUsage: \n./afscan.exe -model icmpalive -ip 192.168.201.1/24\n./afscan.exe -model icmpalive -ip 192.168.201.1/24 -t 500\n./afscan.exe -model portscan -ip 192.168.201.1/24 -port 22,445,1-10000\n./afscan.exe -model sshcrack -ip 192.168.201.1/24 -port 222(使用默认账号密码,指定端口爆破,port参数必须使用。) \n./afscan.exe -model sshcrack -ip 192.168.201.1/24 -port 22,222 -user root,admin -pass root,123456(指定用户名密码爆破，用逗号分割)\n./afscan.exe -model smbcrack -ip 192.168.201.1/24 -port 445 -user administrator,admin,guest,test(可指定密码,不指定默认使用默认密码)\n                       --by au7eng\n\nOptions:\n")
		flag.PrintDefaults() //输出flag
	}
	flag.Parse() //解析flag
}

func main() {
	//基本用法
	switch model {
	case "icmpalive":
		icmpalive.Main(t, ip)
	case "portscan":
		portscan.Main(port, t, ip)
	case "sshcrack":
		sshcrack.Main(ip, port, user, pass)
	case "smbcrack":
		smbcrack.Main(ip, port, user, pass)
	default:
		flag.Usage()
	}
}
